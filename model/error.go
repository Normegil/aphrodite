package model

import (
	"net/url"
	"time"
	"encoding/csv"
	"strconv"
	"bytes"
	"github.com/normegil/aphrodite/ressources"
	"fmt"
)

type ErrWithCode interface {
	Code() int
	error
}

type ErrWithCodeImpl struct {
	error
	code int
}

func (e ErrWithCodeImpl) Code() int {
	return e.code
}

func NewErrWithCode(code int, e error) ErrWithCode {
	return &ErrWithCodeImpl{
		code: code,
		error: e,
	}
}

var errors []Error

func init() {
	fmt.Print(ressources.AssetNames())
	assetBytes, err := ressources.Asset("model/errors.csv")
	if nil != err {
		panic(err)
	}
	content, err := csv.NewReader(bytes.NewReader(assetBytes)).ReadAll()
	if nil != err {
		panic(err)
	}
	for _, row := range content {
		code, err := strconv.Atoi(row[0])
		if nil != err {
			panic(err)
		}
		httpStatus, err := strconv.Atoi(row[1])
		if nil != err {
			panic(err)
		}
		url, err := url.Parse(row[2])
		if nil != err {
			panic(err)
		}
		errors = append(errors, Error{errorDetails{
			Code: code,
			HTTPStatus: httpStatus,
			MoreInfo: *url,
			Message: row[3],
			DeveloperMessage: row[4],
		}})
	}
}

type errorDetails struct {
	HTTPStatus       int
	Code             int
	Message          string
	DeveloperMessage string
	MoreInfo         url.URL
	Time             time.Time
	Err              error
}

type Error struct {
	details errorDetails
}

const DEFAULT_CODE = 500

func NewError(e error) *Error {
	code := DEFAULT_CODE
	if eWithCode, ok := e.(ErrWithCode); ok {
		code = eWithCode.Code()
	}

	for _, err := range errors {
		if (code == err.Code()) {
			generatedError := fromError(err).WithTime(time.Now()).WithError(e)
			return &generatedError
		}
	}
	return &Error{}
}

func fromError(e Error) Error {
	return Error{errorDetails{
		Code: e.Code(),
		HTTPStatus: e.HTTPStatus(),
		Message: e.Message(),
		DeveloperMessage: e.DeveloperMessage(),
		Err: e.SourceError(),
		MoreInfo: e.URL(),
		Time: e.Time(),
	}}
}

func (e Error) HTTPStatus() int {
	return e.details.HTTPStatus
}

func (e Error) Code() int {
	return e.details.Code
}

func (e Error) Message() string {
	return e.details.Message
}

func (e Error) DeveloperMessage() string {
	return e.details.DeveloperMessage
}

func (e Error) URL() url.URL {
	return e.details.MoreInfo
}

func (e Error) Time() time.Time {
	return e.details.Time
}

func (e Error) WithTime(date time.Time) Error {
	err := fromError(e)
	err.details.Time = date
	return err
}

func (e Error) SourceError() error {
	return e.details.Err
}

func (e Error) WithError(err error) Error {
	newError := fromError(e)
	newError.details.Err = err
	return newError
}

func (e Error) String() string {
	errURL := e.URL()
	u := &errURL
	return "[Status HTTP:" + strconv.Itoa(e.HTTPStatus()) + ";Code:" + strconv.Itoa(e.Code()) + ";URL:" + u.String() + ";Time:" + e.Time().String() + ";Msg:" + e.Message() + ";DevMsg:" + e.DeveloperMessage() + ";Err:" + e.SourceError().Error() + "]"
}