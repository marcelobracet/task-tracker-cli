package utils

type Error struct {
	Name string 
	Status int32 
	Message string
	Data interface{}
}

func NewError(name string, status int32, message string, data interface{}) *Error {
	return &Error{
		Name: name,
		Status: status,
		Message: message,
		Data: data,
	}
}

