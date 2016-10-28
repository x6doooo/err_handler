package err_handler

import (
    "errors"
    "strconv"
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

func (me *CommonError) Error() {
    codeStr := strconv.Itoa(me.Code)
    return "[" + codeStr + "] " + me.Msg
}

func GetErr(e interface{}) (err CommonError) {

    er, ok := e.(CommonError)
    if ok {
        err = er
        return
    }

    switch e.(type) {
    case error:
        err = CommonError{
            Code: 1,
            Msg: e.(error).Error(),
            SrcErr: e,
        }
    case string:
        err = CommonError{
            Code: 1,
            Msg: e.(string),
        }
    default:
        err = CommonError{
            Code: 2,
            Msg: "",
            SrcErr: e,
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