package proxy

import (
	"work-codes/wuyue/app/common"
	"work-codes/wuyue/app/modules/book/models"
)

var BookProxy *bookProxy

type bookProxy struct{}

func (proxy *bookProxy) FindOne(filter interface{}) (*map[string]interface{}, error) {
	var row map[string]interface{}
	err := models.BookVO.Find(filter).One(&row)
	return &row, err
}

func (proxy *bookProxy) List(filter interface{}, page int, limit int, sort []string) (*map[string]interface{}, error) {
	var rows []map[string]interface{}
	skip := (page - 1) * limit
	err := models.BookVO.Find(filter).Sort(sort...).Skip(skip).Limit(limit).All(&rows)

	total, _ := models.BookVO.Find(filter).Count()
	result := common.PageList(total, page, limit, &rows)
	return result, err
}
