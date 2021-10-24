package exception

import (
	"errors"
)

type Duplicate struct {
	Err error
	Code int
}

type NotFound struct {
	Err error
	Code int
}

type Forbidden struct {
	Err error
	Code int
}

func PanicForbidden(err error) {
	if err != nil {
		panic(Forbidden{
			Err:  err,
			Code: 403,
		})
	}
}

func PanicNotFound(err error) {
	if err != nil {
		panic(NotFound{
			Err:  errors.New("data tidak ditemukan"),
			Code: 404,
		})
	}
}

func PanicDuplicate(err error) {
	if err != nil {
		panic(Duplicate{
			Err:  err,
			Code: 409,
		})
	}
}