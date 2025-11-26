package custom_errors

type UserNotFoundError struct {
	message string
}

func NewUserNotFoundError(message string) UserNotFoundError {
	return UserNotFoundError{
		message: message,
	}
}

func (e UserNotFoundError) Error() string {
	return e.message
}
