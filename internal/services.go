package internal

import "github.com/sqids/sqids-go"

func shortenUrl(url string) (string, error) {
	newUrl := URL{Url: url}
	result := DB.Create(&newUrl)
	if result.Error != nil {
		return "", result.Error
	} else {
		return *newUrl.Code, nil
	}
}

func decodeUrl(code string) (string, error) {
	s, err := sqids.New()
	if err != nil {
		return "", err
	}

	decodedId := s.Decode(code)

	var result URL
	res := DB.First(&result, decodedId[0])
	if res.Error != nil {
		return "", res.Error
	}

	return result.Url, nil
}
