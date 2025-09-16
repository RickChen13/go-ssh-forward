package forward

import (
	"context"
	"fmt"
	"go-ssh-forward/app/common"
	"io"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

type SshServerConfig struct {
	Host       string `json:"host"`
	User       string `json:"user"`
	Pass       string `json:"pass"`
	KeyPath    string `json:"key_path"`
	PassPhrase string `json:"pass_phrase"`
}

type ForwardConfig struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	RemoteAddr string `json:"remote_addr"`
	LocalAddr  string `json:"local_addr"`
	Tag        string `json:"tag"`
}

type Config struct {
	Forward   ForwardConfig   `json:"forward"`
	SshServer SshServerConfig `json:"ssh_server"`
}

func Run(config Config, ctx context.Context) {
	sshClient, err := getSshClient(config.SshServer)
	if err != nil {
		log.Println(err)
		return
	}
	defer sshClient.Close()
	ctx1, cancel := context.WithCancel(context.Background())
	listener, err := forward(ctx1, config.Forward, sshClient)
	if err != nil {
		log.Println(err)
		cancel()
		return
	}

	// 出错后不会执行到这里，无需担心会 ctx.Done() 会阻塞
	defer listener.Close()
	if config.Forward.Tag == "R" {
		log.Println(config.Forward.Name, "forwarding", config.Forward.LocalAddr+"(local)", "to", config.Forward.RemoteAddr+"(remote)")
	} else {
		log.Println(config.Forward.Name, "forwarding", config.Forward.RemoteAddr+"(remote)", "to", config.Forward.LocalAddr+"(local)")
	}

	for range ctx.Done() {
	}
	cancel()
}

func getSshClient(sshServerConfig SshServerConfig) (*ssh.Client, error) {
	getSshConfig := func() (*ssh.ClientConfig, error) {

		if sshServerConfig.KeyPath != "" {
			var (
				signer ssh.Signer
			)
			key, err := os.ReadFile(sshServerConfig.KeyPath)
			if err != nil {
				return nil, err
			}
			if sshServerConfig.PassPhrase != "" {
				passPhrase := []byte(common.Decode(sshServerConfig.PassPhrase))
				// 使用带密码的私钥解析
				signer, err = ssh.ParsePrivateKeyWithPassphrase(key, passPhrase)
				if err != nil {
					return nil, err
				}
			} else {
				signer, err = ssh.ParsePrivateKey(key)
				if err != nil {
					return nil, err
				}
			}
			// SSH 配置
			config := &ssh.ClientConfig{
				User: sshServerConfig.User,
				Auth: []ssh.AuthMethod{
					ssh.PublicKeys(signer),
				},

				HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 不校验公钥
			}
			return config, nil
		}
		pass := ""
		if sshServerConfig.Pass != "" {
			pass = common.Decode(sshServerConfig.Pass)
		}
		config := &ssh.ClientConfig{
			User: sshServerConfig.User,
			Auth: []ssh.AuthMethod{
				ssh.Password(pass),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 不校验公钥
		}
		return config, nil
	}

	config, err := getSshConfig()
	if err != nil {
		return nil, err
	}
	// 连接 SSH
	sshClient, err := ssh.Dial("tcp", sshServerConfig.Host, config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial SSH server: %v", err)
	}
	return sshClient, nil
}

func forward(ctx context.Context, forwardConfig ForwardConfig, sshConn *ssh.Client) (net.Listener, error) {
	var (
		listenerAddr string
		sendAddr     string

		listener net.Listener
		err      error
	)
	if forwardConfig.Tag == "R" {
		listenerAddr = forwardConfig.RemoteAddr
		sendAddr = forwardConfig.LocalAddr
		listener, err = sshConn.Listen("tcp", listenerAddr)
	} else {
		listenerAddr = forwardConfig.LocalAddr
		sendAddr = forwardConfig.RemoteAddr
		listener, err = net.Listen("tcp", listenerAddr)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to listen on remote port:%w", err)
	}
	go (func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:

			}
			listenerConn, err := listener.Accept()
			if err != nil {
				if opErr, ok := err.(*net.OpError); ok && opErr.Err.Error() == "use of closed network connection" {
					return
				}
				//log.Println("Failed to accept connection:", err)
				continue
			}

			go func() {
				var (
					sendConn net.Conn
					err      error
				)
				if forwardConfig.Tag == "R" {
					sendConn, err = net.Dial("tcp", sendAddr)
				} else {
					sendConn, err = sshConn.Dial("tcp", sendAddr)
				}
				if err != nil {
					log.Println("Failed to dial:", err)
					return
				}
				// 双向转发
				go io.Copy(sendConn, listenerConn)
				go io.Copy(listenerConn, sendConn)
			}()
		}
	})()
	return listener, nil
}
