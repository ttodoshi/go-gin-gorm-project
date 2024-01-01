package errors

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

type BodyMappingError struct {
	Message string
}

func (e *BodyMappingError) Error() string {
	return e.Message
}

type MappingError struct {
	Message string
}

func (e *MappingError) Error() string {
	return e.Message
}
