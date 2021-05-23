# logging

Logging package based on zap and lumberjack

## Quick Start

- Create file ./etc/conf/logging.json and config it.

```json
{
    "default": {
        "filename": "root.log",
        "maxsize": 20,
        "maxage": 7,
        "maxbackups": 50,
        "localtime": true,
        "compress": true,
        "level": "warn",
        "env": "prod"
    },
    "root": {
        "filename": "root.log",
        "maxsize": 20,
        "maxage": 7,
        "maxbackups": 50,
        "localtime": true,
        "compress": true,
        "level": "warn",
        "env": "prod"
    },
    "proxy": {
        "filename": "proxy.log",
        "maxsize": 20,
        "maxage": 7,
        "maxbackups": 50,
        "localtime": true,
        "compress": true,
        "level": "warn"
    }
}
```

The `root` and `proxy` are the names of the loggers from which you can get the logger instance.  

The range of level is `debug`, `info`, `warn`, `error`, `dpanic`, `panic`, `fatal`, and the value is from low to high.  

- Usage

```go
package main

import (
    "github.com/mallbook/logging"
    "go.uber.org/zap"
)

func main() {
    logger := logging.Logger("mylog") // or logger := logging.Logger("root")

    for i := 0; i < 100; i++ {
        logger.Info("hello world", zap.String("key", "value"), zap.Int("age", 20))
        logger.Debug("hello china", zap.String("key", "value"), zap.Int("age", 20))
        logger.Error("hello error", zap.String("key", "value"), zap.Int("age", 30))
    }

    proxyLogger := logging.Logger("proxy") 
    proxyLogger.Info("hello world", zap.String("key", "value"), zap.Int("age", 20))
}
```

Logo Level

```go
// DebugLevel logs are typically voluminous, and are usually disabled in
// production.
DebugLevel Level = iota - 1

// InfoLevel is the default logging priority.
InfoLevel

// WarnLevel logs are more important than Info, but don't need individual
// human review.
WarnLevel

// ErrorLevel logs are high-priority. If an application is running smoothly,
// it shouldn't generate any error-level logs.
ErrorLevel

// DPanicLevel logs are particularly important errors. In development the
// logger panics after writing the message.
DPanicLevel

// PanicLevel logs a message, then panics.
PanicLevel

// FatalLevel logs a message, then calls os.Exit(1).
FatalLevel
```
