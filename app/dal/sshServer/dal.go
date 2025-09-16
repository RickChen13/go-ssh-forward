package sshServer

import (
	"gorm.io/gorm"
)

type SshServerDal struct {
	db *gorm.DB
}

func NewSshServerDal(db *gorm.DB) *SshServerDal {
	return &SshServerDal{db: db}
}

func (dal *SshServerDal) List(page, number int) (results []ConfigSshServer, err error) {
	err = dal.db.Offset((page - 1) * number).Limit(number).Order("sort DESC, id DESC").Find(&results).Error
	return
}

// 插入一条数据
func (dal *SshServerDal) Insert(data ConfigSshServerC) error {
	return dal.db.Create(&data).Error
}

func (dal *SshServerDal) Delete(id int) error {
	return dal.db.Where("id = ?", id).Delete(&ConfigSshServer{}).Error
}

func (dal *SshServerDal) Update(id int, data map[string]interface{}) error {
	return dal.db.Model(&ConfigSshServer{}).Where("id = ?", id).Updates(data).Error
}

func (dal *SshServerDal) GetInfoById(id int) (ConfigSshServerC, error) {
	var data ConfigSshServerC
	err := dal.db.Where("id = ?", id).First(&data).Error
	return data, err
}
