package configs

type RequestError struct {
	Code    int
	Message string
}

func (re RequestError) Error() string {
	return re.Message
}

type ServerError struct {
	Code    int
	Message string
}

func (se ServerError) Error() string {
	return se.Message
}

type ResponseError struct {
	Code    int
	Message string
}

func (rpe ResponseError) Error() string {
	return rpe.Message
}

type DataValidationError struct {
	Field   string
	Message string
}

func (dve DataValidationError) Error() string {
	return dve.Message
}
