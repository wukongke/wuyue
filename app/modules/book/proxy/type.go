package proxy

import (
	"work-codes/wuyue/app/common"
	"work-codes/wuyue/app/modules/book/models"
)

var TypeProxy *typeProxy

type typeProxy struct{}

func (proxy *typeProxy) FindOne(filter interface{}) (*map[string]interface{}, error) {
	var row map[string]interface{}
	err := models.TypeVO.Find(filter).One(&row)
	// result := Convert2Hotel(row)
	return &row, err
}

func (proxy *typeProxy) List(filter interface{}, page int, limit int, sort []string) *map[string]interface{} {
	var rows []map[string]interface{}
	skip := (page - 1) * limit
	err := models.TypeVO.Find(filter).Sort(sort...).Skip(skip).Limit(limit).All(&rows)
	if err != nil {
		panic(err)
		return nil
	}
	total, _ := models.TypeVO.Find(filter).Count()
	result := common.PageList(total, page, limit, &rows)
	return result
}
