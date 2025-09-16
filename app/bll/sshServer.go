package bll

import (
	"go-ssh-forward/app/common"
	"go-ssh-forward/app/dal/sshServer"
)

type SshServerBll struct {
	ssDal *sshServer.SshServerDal
}

func NewSshServerBll(ssDal *sshServer.SshServerDal) *SshServerBll {
	return &SshServerBll{
		ssDal: ssDal,
	}
}

func (b *SshServerBll) List(page, number int) Result {
	results, err := b.ssDal.List(page, number)
	if err != nil {
		return Result{
			Result: false,
			Msg:    err.Error(),
			Data:   nil,
		}
	}
	return Result{
		Result: true,
		Msg:    "success",
		Data:   results,
	}
}

func (b *SshServerBll) Add(data sshServer.ConfigSshServerC) Result {
	if data.Pass != "" {
		data.Pass = common.Encode(data.Pass)
	}
	if data.PassPhrase != "" {
		data.PassPhrase = common.Encode(data.PassPhrase)
	}
	err := b.ssDal.Insert(data)
	return NilDataErr(err)
}

func (b *SshServerBll) Delete(id int) Result {
	err := b.ssDal.Delete(id)
	return NilDataErr(err)
}

func (b *SshServerBll) Update(id int, data sshServer.ConfigSshServerU) Result {
	updates := map[string]interface{}{
		"name":     data.Name,
		"host":     data.Host,
		"user":     data.User,
		"pass":     data.Pass,
		"key_path": data.KeyPath,
		"sort":     data.Sort,
	}

	if data.Pass != "" {
		updates["pass"] = common.Encode(data.Pass)
	}
	if data.ClearPass {
		updates["pass"] = ""
	}

	if data.PassPhrase != "" {
		updates["pass_phrase"] = common.Encode(data.PassPhrase)
	}
	if data.ClearPassPhrase {
		updates["pass_phrase"] = ""
	}
	err := b.ssDal.Update(id, updates)
	return NilDataErr(err)
}

func (b *SshServerBll) GetInfoById(id int) Result {
	data, err := b.ssDal.GetInfoById(id)
	return DataErr(data, err)
}
