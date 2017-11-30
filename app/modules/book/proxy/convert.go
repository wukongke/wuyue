package proxy

func Convert2Book(data map[string]interface{}) *map[string]interface{} {
	row := map[string]interface{}{
		"id":          data["_id"],
		"bookNo":      data["bookNo"],
		"name":        data["name"],
		"accountId":   data["accountId"],
		"type":        data["type"],
		"intro":       data["intro"],
		"image":       data["image"],
		"rate":        data["rate"],
		"wordCount":   data["wordCount"],
		"serialize":   data["serialize"],
		"loveCount":   data["loveCount"],
		"likes":       data["likes"],
		"url":         data["url"],
		"lastChapter": data["lastChapter"],
		"status":      data["status"],
		"createdAt":   data["createdAt"],
		"updatedAt":   data["updatedAt"],
	}
	return &row
}

func Convert2Chapter(data map[string]interface{}) *map[string]interface{} {
	row := map[string]interface{}{
		"id":        data["_id"],
		"chapterNo": data["chapterNo"],
		"title":     data["title"],
		"bookNo":    data["bookNo"],
		"content":   data["content"],
		"url":       data["url"],
		"status":    data["status"],
		"createdAt": data["createdAt"],
		"updatedAt": data["updatedAt"],
	}
	return &row
}
