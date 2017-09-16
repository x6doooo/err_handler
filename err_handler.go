package err_handler

import (
    "strconv"
    "errors"
    "fmt"
)

type CommonError struct {
    Code   int
    Msg    string
    SrcErr error
}

func NewCommonError(code int, msg string, err error) CommonError {
    return CommonError{
        Code: code,
        Msg: msg,
        SrcErr: err,
    }
}

func (me *CommonError) Error() string {
    codeStr := strconv.Itoa(me.Code)
    return "[" + codeStr + "] " + me.Msg
}

func GetErr(e interface{}) (err *CommonError) {

    cerr, ok := e.(CommonError)
    if ok {
        err = &cerr
        return
    }

    switch e.(type) {
    case error:
        er := e.(error)
        err = &CommonError{
            Code: 1,
            Msg: er.Error(),
            SrcErr: er,
        }
    case string:
        er := e.(string)
        err = &CommonError{
            Code: 1,
            Msg: er,
            SrcErr: errors.New(er),
        }
    default:
        er := fmt.Sprintf("meta error: %v", e)
        err = &CommonError{
            Code: 2,
            Msg: er,
            SrcErr: errors.New(er),
        }
    }
    return
}

func Recover(err *error, cb func()) {
    if e := recover(); e != nil {
        *err = GetErr(e)
    }
    cb()
}

