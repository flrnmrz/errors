package errors

import (
    "fmt"

    logger "github.com/Sirupsen/logrus"
)

func WitField(key string, value interface{}) (eb ErrorBuilder) {
    str := fmt.Sprintf("%v", value)
    return ErrorBuilder{key, str, nil}
}

// TODO: also want New, Wrap, and Fmt implemented as package-scoped function

type ErrorBuilder struct {
    key     string
    value   string
    wrapped *ErrorBuilder
}

func (eb *ErrorBuilder) WithField(key string, value interface{}) *ErrorBuilder {
    res := ErrorBuilder{}
    res.key = key
    res.value = fmt.Sprintf("%v", value)
    res.wrapped = eb
    return &res
}

// recursively collect key values
func (eb *ErrorBuilder) collectKeyValues(res map[string]string) {
    res[eb.key] = eb.value
    if eb.wrapped != nil {
        eb.wrapped.collectKeyValues(res)
    }
}

func (eb *ErrorBuilder) Wrap(err error, msg string) *WrappingError {
    kv := map[string]string{}
    eb.collectKeyValues(kv)
    return &WrappingError{msg:msg, wrapped:err, keyValue : kv}
}

func (eb *ErrorBuilder) New(msg string) *WrappingError {
    kv := map[string]string{}
    eb.collectKeyValues(kv)
    return &WrappingError{msg: msg, wrapped: nil, keyValue:kv}
}


func (eb *ErrorBuilder) Fmt(format string, args ...interface{}) *WrappingError {
    msg := fmt.Sprintf(format, args...)
    kv := map[string]string{}
    eb.collectKeyValues(kv)
    return &WrappingError{msg: msg, wrapped: nil, keyValue:kv}
}


type WrappingError struct {
    keyValue map[string]string
    msg      string
    wrapped  error
}

func (err *WrappingError) Log(logger logger.Logger) *WrappingError {
    logger.Error(err.Error())

    return err
}

func (w *WrappingError) Error() string {
    // TODO: recursively assemble a string from the error
    return ""
}
