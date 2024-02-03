package utils

import (
	"CARIAD/pkg/models"
	"net/url"
	"sort"
	"strings"
)

func CreateNumberRequest(urls []string) []models.NumberRequest {
	var numberRequests []models.NumberRequest

	for _, urlPath := range urls {
		parsedURL, err := url.ParseRequestURI(urlPath)
		if err == nil {
			numberRequests = append(numberRequests, models.NumberRequest{
				Url:   parsedURL,
				NType: findNumberType(parsedURL),
			})
		}
	}
	return numberRequests
}

func findNumberType(parsedURL *url.URL) models.NumberType {
	path := strings.Trim(parsedURL.Path, "/")
	if models.NumberType(path).ValidateTypes() {
		return models.NumberType(path)
	}

	return models.Err
}

func SortAndRemoveDuplicates(responses []int) []int {
	uniqueMap := make(map[int]struct{})
	for _, value := range responses {
		uniqueMap[value] = struct{}{}
	}

	var result []int
	for key := range uniqueMap {
		result = append(result, key)
	}

	sort.Ints(result)
	return result
}
