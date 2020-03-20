package bingContent

type InvalidArgumentError struct {
	message string
}

func (e *InvalidArgumentError) Error() string {
	return e.message
}
