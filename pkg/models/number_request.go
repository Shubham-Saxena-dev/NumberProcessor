package models

import (
	"net/url"
)

type NumberRequest struct {
	Url   *url.URL
	NType NumberType
}
