package errors

import (
    "fmt"

    logger "github.com/Sirupsen/logrus"
)

func New(msg string) *WrappingError {
    return &WrappingError{msg:msg, wrapped:nil, keyValue:map[string]string{}}
}

func Fmt(format string, args ...interface{}) *WrappingError {
    msg := fmt.Sprintf(format, args...)
    return &WrappingError{msg: msg, wrapped: nil, keyValue:map[string]string{}}
}

func Wrap(err error, msg string) *WrappingError {
    kv := map[string]string{}
    return &WrappingError{msg:msg, wrapped:err, keyValue : kv}
}

func WithField(key string, value interface{}) (eb *ErrorBuilder) {
    str := fmt.Sprintf("%v", value)
    return &ErrorBuilder{key, str, nil}
}

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

func (err *WrappingError) Error() string {
    // TODO: recursively assemble a string from the error
    msg := err.msg + " ("
    idx := 0
    for key, value := range err.keyValue {
        if idx != 0 {
            msg += ", " + key + ": " + value
        } else {
            msg += key + ": " + value
        }
    }

    return msg + ")"
}
