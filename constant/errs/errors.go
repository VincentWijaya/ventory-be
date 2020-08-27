package errs

import (
	"errors"
)

var (
	//errors
	Unauthorized = errors.New("unauthorized")
	NoData       = errors.New("cannot fetch requested data")
	BadRequest   = errors.New("bad Request")
	BadConfig    = errors.New("bad configuration on ventory")
)

var (
	//codes
	Success               = "00"
	UnauthorizedErrorCode = "03"
	NoDataCode            = "-1"
	BadRequestCode        = "04"
	BadConfigCode         = "05"
	UndefinedErrorCode    = "99"

	//messages
	GeneralErrorMessage = "Saat ini sedang terjadi gangguan pada system, silahkan coba beberapa saat lagi"
	NoDataMessage       = "Data tidak di temukan"
)
