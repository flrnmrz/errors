package errors

func Wrap(err error, msg string) (eb *ErrorBuilder) {
    return ErrorBuilder{msg:msg, wrapped:err}
}

func New(msg string) {
    return ErrorBuilder{msg: msg}
}


func Fmt(format string, args ..interface{}) (eb *ErrorBuilder) {
    msg, err := fmt.Sprintf(format, args)
    if err != nil {
        // ?
    }

    return ErrorBuilder{ msg: msg }
}


type ErrorBuilder struct {
    msg string
    wrapped error
    fields map[string]interface{}
}

func (eb *ErrorBuilder) Log(logger Logger) (eb *ErrorBuilder) {
    for key, value := range eb.fields {
        logger.WithField(key, value)
    }

    logger.Error()

    return eb
}


func (eb *ErrorBuilder) WithField(key string, value interface{}) (eb *ErrorBuilder) {
    eb.fields[key] = value
    return eb
}

type kvError {
    // how to store stuff
}

func (eb *ErrorBuilder) Error() {
    // build kv error from the error builder
}

