package URLHandler

type BadRequestError struct{}

func (r BadRequestError) Error() string {
	return "400 Bad Request"
}

type ForbiddenError struct{}

func (r ForbiddenError) Error() string {
	return "403 Forbidden"
}

type NotFoundError struct{}

func (r NotFoundError) Error() string {
	return "404 Not Found"
}

type InvalidMethodError struct{}

func (r InvalidMethodError) Error() string {
	return "405 Method Not Allowed"
}
