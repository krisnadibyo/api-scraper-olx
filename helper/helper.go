package helper

import "strconv"

func GenerateUrl(url string, keyword string, page int) string {
	location := "4000020" // Bekasi

	if page > 0 {
		return url + "&query=" + keyword + "&location=" + location + "&page=" + strconv.Itoa(page)
	} 
	return url + "&query=" + keyword + "&location=" + location
}