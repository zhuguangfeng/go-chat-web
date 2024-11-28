package utils

func GetPageCount(count, pageSize int) int {
	if pageSize == 0 {
		pageSize = 10
	}
	pages := count / pageSize
	if count%pageSize != 0 {
		pages++
	}
	return pages
}
