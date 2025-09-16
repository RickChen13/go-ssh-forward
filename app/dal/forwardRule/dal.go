package forwardRule

import (
	"gorm.io/gorm"
)

type ForwardRuleDal struct {
	db *gorm.DB
}

func NewForwardRuleDal(db *gorm.DB) *ForwardRuleDal {
	return &ForwardRuleDal{
		db: db,
	}
}

func (dal *ForwardRuleDal) List(page, number int) (results []ConfigForwardRuleData, err error) {
	skip := (page - 1) * number

	err = dal.db.Table("config_forward_rule AS cfr").
		// 使用 Joins 方法来执行 LEFT JOIN
		Joins("left join config_ssh_server AS css on css.id = cfr.css_id").
		// 使用 Select 方法来指定查询的列，并使用 AS 来创建别名
		Select(`
			cfr.id,
			cfr.name,
			cfr.remote_addr,
			cfr.local_addr,
			cfr.tag,
			cfr.sort,
			cfr.css_id,
			css.name AS ssh_server_name
		`).
		// 使用 Order 方法来指定排序
		Order("cfr.sort DESC, cfr.id DESC").
		// 使用 Limit 和 Offset 来实现分页
		Limit(number).
		Offset(skip).
		// 执行查询并将结果扫描到 Result 结构体切片中
		Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

// 插入一条数据
func (dal *ForwardRuleDal) Insert(data ConfigForwardRuleUpdateData) error {
	return dal.db.Create(&data).Error
}

func (dal *ForwardRuleDal) Delete(id int) error {
	return dal.db.Where("id = ?", id).Delete(&ConfigForwardRule{}).Error
}

func (dal *ForwardRuleDal) Update(id int, data ConfigForwardRuleUpdateData) error {
	return dal.db.Model(&ConfigForwardRule{}).Where("id = ?", id).Updates(data).Error
}

func (dal *ForwardRuleDal) GetInfoById(id int) (ConfigForwardRuleUpdateData, error) {
	var data ConfigForwardRuleUpdateData
	err := dal.db.Where("id = ?", id).First(&data).Error
	return data, err
}

func (dal *ForwardRuleDal) CssIdCount(id int) (int64, error) {
	var count int64
	err := dal.db.Model(&ConfigForwardRule{}).Where("css_id = ?", id).Count(&count).Error
	return count, err
}
