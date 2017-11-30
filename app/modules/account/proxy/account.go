package proxy

import (
	"work-codes/wuyue/app/common"
	"work-codes/wuyue/app/modules/account/models"
)

var AccountProxy *accountProxy

type accountProxy struct{}

func (proxy *accountProxy) FindOne(filter interface{}) (*map[string]interface{}, error) {
	var row map[string]interface{}
	err := models.AccountVO.Find(filter).One(&row)
	return &row, err
}

func (proxy *accountProxy) List(filter interface{}, page int, limit int, sort []string) *map[string]interface{} {
	var rows []map[string]interface{}
	skip := (page - 1) * limit
	err := models.AccountVO.Find(filter).Sort(sort...).Skip(skip).Limit(limit).All(&rows)
	if err != nil {
		panic(err)
		return nil
	}
	total, _ := models.AccountVO.Find(filter).Count()
	result := common.PageList(total, page, limit, &rows)
	return result
}
