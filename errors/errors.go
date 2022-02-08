package errors

/**
 * This struct creates a error info from the two parts: Code and Msg,
 * , which is uint16 and String, respectively.
 *
 * @param Code           the code of error as a string
 * @param Msg            the message of error as a string
 */
type ErrorInfo struct {
	Code uint16 `json:"code"`
	Msg  string `json:"msg"`
}

func (e ErrorInfo) IsNil() bool {
	return ErrorInfo{} == e
}

var Ok = ErrorInfo{}

// Invalid input
var InvalidInputError = ErrorInfo{
	Code: 400,
	Msg:  "Invalid input",
}

var NotFoundError = ErrorInfo{
	Code: 404,
	Msg:  "Not found",
}

//Inter server error means unexpected error
var InternalServerError = ErrorInfo{
	Code: 500,
	Msg:  "Internal Server Error",
}
