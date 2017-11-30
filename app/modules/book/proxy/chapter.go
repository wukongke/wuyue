package proxy

import (
	"work-codes/wuyue/app/common"
	"work-codes/wuyue/app/modules/book/models"
)

var ChapterProxy *chapterProxy

type chapterProxy struct{}

func (proxy *chapterProxy) Find(filter interface{}) ([]*map[string]interface{}, error) {
	var rows []*map[string]interface{}
	err := models.ChapterVO.Find(filter).All(&rows)
	return rows, err
}

func (proxy *chapterProxy) FindOne(filter interface{}) (*map[string]interface{}, error) {
	var row map[string]interface{}
	err := models.ChapterVO.Find(filter).One(&row)
	return &row, err
}

func (proxy *chapterProxy) List(filter interface{}, page int, limit int, sort []string) *map[string]interface{} {
	var rows []map[string]interface{}
	skip := (page - 1) * limit
	err := models.ChapterVO.Find(filter).Sort(sort...).Skip(skip).Limit(limit).All(&rows)
	if err != nil {
		panic(err)
		return nil
	}
	total, _ := models.ChapterVO.Find(filter).Count()
	result := common.PageList(total, page, limit, &rows)
	return result
}
