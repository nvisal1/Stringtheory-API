package shared

type errorTypes struct{
	BadRequest string
	InvalidAccess string
	Forbidden string
	NotFound string
}

var ErrorTypes = errorTypes{
	BadRequest:    "BadRequest",
	InvalidAccess: "InvalidAccess",
	Forbidden:     "Forbidden",
	NotFound:      "NotFound",
}

type ServiceError struct {
	err    string
	errType  string
}

func (errorTypes errorTypes) mapErrorToHttpStatus(errorType string) int {
	switch errorType {
	case ErrorTypes.BadRequest:
		return 400
	case ErrorTypes.InvalidAccess:
		return 401
	case ErrorTypes.Forbidden:
		return 403
	case ErrorTypes.NotFound:
		return 404
	default:
		return 500
	}
}

