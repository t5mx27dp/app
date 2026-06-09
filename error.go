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

func (e *SystemError) Error() string {
	if e.Message == "" {
		return e.Code
	}
	return e.Code + ": " + e.Message
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

func (e *BusinessError) Error() string {
	if e.Message == "" {
		return e.Code
	}
	return e.Code + ": " + e.Message
}
