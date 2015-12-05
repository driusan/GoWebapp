package URLHandler

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
