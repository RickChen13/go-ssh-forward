package bll

import (
	"context"
	"fmt"
	"go-ssh-forward/app/common/flog"
	"go-ssh-forward/app/common/forward"
	"go-ssh-forward/app/common/sqlite3"
	"go-ssh-forward/app/dal/forwardRule"
	"go-ssh-forward/app/dal/sshServer"
	"sync"
)

type Status struct {
	name       string
	cancelFunc context.CancelFunc
}

type forwardBll struct {
	ch     chan forward.Config
	status map[string]Status
	mu     sync.Mutex

	frDal *forwardRule.ForwardRuleDal
	ssDal *sshServer.SshServerDal
}

func NewForwardDal() *forwardBll {
	forwar := &forwardBll{
		ch:     make(chan forward.Config, 4),
		status: make(map[string]Status),
		frDal:  forwardRule.NewForwardRuleDal(sqlite3.Db),
		ssDal:  sshServer.NewSshServerDal(sqlite3.Db),
	}
	go forwar.start()
	return forwar
}

func (f *forwardBll) start() {
	for config := range f.ch {
		f.mu.Lock()
		_, ok := f.status[fmt.Sprint(config.Forward.Id)]
		if !ok {
			ctx, cancel := context.WithCancel(context.Background())
			f.status[fmt.Sprint(config.Forward.Id)] = Status{
				name:       config.Forward.Name,
				cancelFunc: cancel,
			}
			go (func() {
				forward.Run(config, ctx)
				f.mu.Lock()
				_, exits := f.status[fmt.Sprint(config.Forward.Id)]
				if exits {
					delete(f.status, fmt.Sprint(config.Forward.Id))
				}
				f.mu.Unlock()
				flog.Debug(config.Forward.Name + " stop")
			})()
		} else {
			flog.Debug("Already running: " + config.Forward.Name)
		}
		f.mu.Unlock()
	}
}

func (f *forwardBll) Stop(id string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if status, ok := f.status[id]; ok {
		if status.cancelFunc != nil {
			status.cancelFunc()
		}
		delete(f.status, id)
	}
}

func (f *forwardBll) Run(config forward.Config) {
	f.ch <- config
}

func (f *forwardBll) RunById(id int) {
	forwardInfo, err := f.frDal.GetInfoById(id)
	if err != nil {
		flog.Debug("GetInfoById err:" + err.Error())
		return
	}

	ss, err := f.ssDal.GetInfoById(forwardInfo.CssId)
	if err != nil {
		flog.Debug("GetInfoById err:" + err.Error())
		return
	}
	config := forward.Config{
		Forward: forward.ForwardConfig{
			Id:         id,
			Name:       forwardInfo.Name,
			RemoteAddr: forwardInfo.RemoteAddr,
			LocalAddr:  forwardInfo.LocalAddr,
			Tag:        forwardInfo.Tag,
		},
		SshServer: forward.SshServerConfig{
			Host:       ss.Host,
			User:       ss.User,
			Pass:       ss.Pass,
			KeyPath:    ss.KeyPath,
			PassPhrase: ss.PassPhrase,
		},
	}
	f.ch <- config
}

func (f *forwardBll) Status() map[string]bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	res := make(map[string]bool)
	for id, status := range f.status {
		res[id] = status.cancelFunc != nil
	}
	return res
}
