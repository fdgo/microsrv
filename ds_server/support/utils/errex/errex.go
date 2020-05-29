package errex

var Errors = make(map[string]*Error)

type Error struct {
	HttpErr int    `json:"http_err"`
	ErrCode string `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

var nilErr *Error

func NewError(httpCode int, code, msg string) *Error {
	err := &Error{httpCode, code, msg}
	Errors[code] = err
	return err
}

func New(err string) *Error {
	return &Error{0, err, ""}
}

func (e Error) Error() string {
	return e.ErrCode
}

func (e Error) HttpCode() int {
	return e.HttpErr
}

func (e Error) Msg() string {
	return e.ErrMsg
}

func (e Error) Equal(e1 error) bool {
	return e.Error() == e1.Error()
}

func Equal(e1, e2 error) bool {
	return e1.Error() == e2.Error()
}

func IsNil(e error) bool {
	if nil == e || nilErr == e {
		return true
	}
	return false
}
