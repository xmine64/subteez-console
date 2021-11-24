package subteez

type NotFoundError struct{}
type BadRequestError struct{}
type ServerError struct{}

type ClientError struct {
	message string
}

func (*NotFoundError) Error() string {
	return "Requested resource not found."
}

func (*BadRequestError) Error() string {
	return "Bad request"
}

func (*ServerError) Error() string {
	return "Server error"
}

func (e *ClientError) Error() string {
	return "Client error: " + e.message
}
