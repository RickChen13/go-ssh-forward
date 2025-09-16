package bll

import "go-ssh-forward/app/dal/forwardRule"

type ForwardRuleBll struct {
	frDal *forwardRule.ForwardRuleDal
}

func NewForwardRuleBll(frDal *forwardRule.ForwardRuleDal) *ForwardRuleBll {
	return &ForwardRuleBll{
		frDal: frDal,
	}
}

func (b *ForwardRuleBll) List(page, number int) Result {
	results, err := b.frDal.List(page, number)
	return DataErr(results, err)
}

func (b *ForwardRuleBll) Add(data forwardRule.ConfigForwardRuleUpdateData) Result {
	err := b.frDal.Insert(data)
	return NilDataErr(err)
}

func (b *ForwardRuleBll) Delete(id int) Result {
	err := b.frDal.Delete(id)
	return NilDataErr(err)
}

func (b *ForwardRuleBll) Update(id int, data forwardRule.ConfigForwardRuleUpdateData) Result {
	err := b.frDal.Update(id, data)
	return NilDataErr(err)
}

func (b *ForwardRuleBll) GetInfoById(id int) Result {
	data, err := b.frDal.GetInfoById(id)
	return DataErr(data, err)
}

func (b *ForwardRuleBll) CssIdCount(id int) Result {
	count, err := b.frDal.CssIdCount(id)
	return DataErr(count, err)
}
