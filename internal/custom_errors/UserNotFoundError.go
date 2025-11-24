package custom_errors

const userNotFoundMessage = "user not found"

type UserNotFoundError struct {
}

func (e UserNotFoundError) Error() string {
	return userNotFoundMessage
}
