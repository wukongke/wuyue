package common

func PageUtil(total int, page int, limit int, docs *[]interface{}) *map[string]interface{} {
	pages := total / limit
	if total%limit > 0 {
		pages = total/limit + 1
	}
	if len(*docs) <= 0 {
		*docs = []interface{}{}
	}
	result := map[string]interface{}{
		"total": total,
		"limit": limit,
		"page":  page,
		"pages": pages,
		"docs":  *docs,
	}
	return &result
}

func PageList(total int, page int, limit int, docs *[]map[string]interface{}) *map[string]interface{} {
	pages := total / limit
	if total%limit > 0 {
		pages = total/limit + 1
	}
	if len(*docs) <= 0 {
		*docs = []map[string]interface{}{}
	}
	result := map[string]interface{}{
		"total": total,
		"limit": limit,
		"page":  page,
		"pages": pages,
		"docs":  *docs,
	}
	return &result
}
