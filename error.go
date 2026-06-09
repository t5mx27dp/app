package app

type SystemError struct {
	Code    string
	Message string
}

func NewSystemError(code, message string) *SystemError {
	return &SystemError{
		Code:    code,
		Message: message,
	}
}

type BusinessError struct {
	Code    string
	Message string
}

func NewBusinessError(code, message string) *BusinessError {
	return &BusinessError{
		Code:    code,
		Message: message,
	}
}
