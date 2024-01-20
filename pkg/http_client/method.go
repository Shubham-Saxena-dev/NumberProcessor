package http_client

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
	PATCH  Method = "PATCH"
)

func (m Method) validateMethodType() Method {
	switch m {
	case POST, PUT, DELETE, PATCH:
		return m
	}
	return GET
}
