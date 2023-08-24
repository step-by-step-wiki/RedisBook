package resp

type Response struct {
	Code    int
	Message string
	Data    []interface{}
}
