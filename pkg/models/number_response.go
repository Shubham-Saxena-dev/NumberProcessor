package models

var EmptyResponse = NumbersResponse{}

type NumbersResponse struct {
	Numbers []int `json:"numbers"`
}
