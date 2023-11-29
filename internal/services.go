package internal

func shortenUrl(url string) (string, error) {
	newUrl := URL{Url: url}
	result := DB.Create(&newUrl)
	if result.Error != nil {
		return "", result.Error
	} else {
		return *newUrl.Code, nil
	}
}
