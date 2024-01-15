package myerrors

import "fmt"

type MyErros struct {
	code int
	msg  string
}

func (m *MyErros) Error() string {
	return fmt.Sprintf("MyError - code: %d msg: %s\n", m.code, m.msg)
}

func NewMyErrors(code int, msg string) error {
	return &MyErros{
		code: code,
		msg:  msg,
	}
}
