package core

type Error interface {
	Error() string
	Code() uint32
	Status() string
}
