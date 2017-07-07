# errors
Wrap-style error handling with built-in logrus-style logging support.

## Examples

The intended usage of this package is as follows:

```go
import (
    "github.com/flrnmrz/errors"
    logger "github.com/Sirupsen/logrus"
)

func doSomething() error {
    errors := errors.WithField("function","doSomething")

    err := doSomethingElse(value)
    if err != nil {
        return errors.WithField("arg", value).Wrap(err, "failed doing something else")
    }

    err = doSomethingDifferent(value)
    if err != nil {
        // Log here, but also attach to the returned error
        // will not be displaced by default with the returned error however
        return errors.WithField("arg", value).Wrap(err, "failed doing something different").Log(logger)
    }

    _, ok := doSomethingEvenMoreDifferent(value)
    if !ok {
        return errors.New("Failed doing something even more different")
    }

    _, ok := foobar(arg)
    if !ok {
        errors.WithField("arg", arg).Fmt("failed processing %s", something)
    }

    return nil
}
```
