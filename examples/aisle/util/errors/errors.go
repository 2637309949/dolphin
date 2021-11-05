package errors

import "github.com/go-errors/errors"

type Error interface {
	Error() string
	Code() uint32
	Status() string
}

type err struct {
	err    string
	code   uint32
	status string
}

func (e *err) Error() string {
	return e.err
}

func (e *err) Code() uint32 {
	return e.code
}

func (e *err) Status() string {
	return e.status
}

func New(code uint32, msg, status string) error {
	return &err{code: code, err: msg, status: status}
}

func NewWithError(code uint32, e error, status string) error {
	return &err{code: code, err: e.Error(), status: status}
}

func Wrap(e interface{}, skip int) *errors.Error {
	return errors.Wrap(e, skip)
}

func Code(u uint32) uint32 {
	_, ok := fix[u]
	if ok {
		panic(New(100001, "code already exists", "code already exists"))
	}
	fix[u] = u
	return u
}

var fix = map[uint32]uint32{}
