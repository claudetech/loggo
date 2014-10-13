# loggo [![Build Status](https://travis-ci.org/claudetech/loggo.svg?branch=master)](https://travis-ci.org/claudetech/loggo)

An easy to use and configurable logger for Go.

## Basic usage

`loggo/default` exposes a `Log` instance to use for
really simple cases.
It logs levels `<= Info` to `stdout` and level `>= Warning`
to stderr.

```go
package main

import (
  "github.com/claudetech/loggo"
  . "github.com/claudetech/loggo/default"
)

func main() {
  Log.SetLevel(loggo.Trace)
  Log.Trace("foobar")
  Log.Debug("foobar")
  Log.Info("foobar")
  Log.Warning("foobar")
  Log.Error("foobar")
  Log.Fatal("foobar")
}
```

will give the following output

![loggo-output](http://res.cloudinary.com/dtdu3sqtl/image/upload/v1413182891/loggo_b2iw6n.png)


## Normal usage

In most use cases, you will first create a logger and
add some appenders, eventually with some filters.

For example, this logger will log everything to stdout and color
the output.

```go
package main

import (
  "github.com/claudetech/loggo"
)

func main() {
  logger := loggo.New("logger name")
  logger.SetLevel(loggo.Info)
  stdoutAppender := loggo.NewStdoutAppender()
  logger.AddAppender(stdoutAppender, loggo.Color)
  logger.Infof("logger started with level %s", logger.Level().String())
  logger.Debug("this is debug")
  logger.Fatal("this is fatal")
}
```

The `Debug` is be ignored as the level is set to `Info`.

### Appenders

Appenders work more or less like in log4j and company.
An `Appender` is only an interface with an `Append` method
that takes a `Message` and returns nothing.
