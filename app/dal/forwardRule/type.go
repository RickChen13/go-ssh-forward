package forwardRule

type tableNameConfig struct{}

func (tableNameConfig) TableName() string {
	return "config_forward_rule"
}

type ConfigForwardRuleUpdateData struct {
	tableNameConfig
	Name       string `json:"name"`
	RemoteAddr string `json:"remote_addr"`
	LocalAddr  string `json:"local_addr"`
	Tag        string `json:"tag"`
	Sort       int    `json:"sort"`
	CssId      int    `json:"css_id"`
}

type ConfigForwardRule struct {
	tableNameConfig
	Id                          uint `json:"id"`
	ConfigForwardRuleUpdateData `gorm:"embedded"`
}

type ConfigForwardRuleData struct {
	tableNameConfig
	ConfigForwardRule `gorm:"embedded"`
	SshServerName     string `json:"ssh_server_name"`
}
