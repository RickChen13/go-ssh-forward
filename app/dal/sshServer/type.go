package sshServer

type tableNameConfig struct{}

func (tableNameConfig) TableName() string {
	return "config_ssh_server"
}

type ConfigSshServerC struct {
	tableNameConfig
	Name       string `json:"name"`
	Host       string `json:"host"`
	User       string `json:"user"`
	Pass       string `json:"pass"`
	KeyPath    string `json:"key_path"`
	PassPhrase string `json:"pass_phrase"`
	Sort       uint   `json:"sort"`
}

type ConfigSshServerU struct {
	tableNameConfig
	ConfigSshServerC `gorm:"embedded"`

	ClearPass       bool `json:"clear_pass"`
	ClearPassPhrase bool `json:"clear_pass_phrase"`
}

type ConfigSshServer struct {
	tableNameConfig
	Id               uint `json:"id"`
	ConfigSshServerC `gorm:"embedded"`
}
