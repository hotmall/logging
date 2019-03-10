# logging
Logging package based on zap and lumberjack
## Quick Start

1. Create file ./etc/conf/logging.json and config it.
```json
{
    "root" : {
        "rotate": {
            "filename": "var/log/root.log",
            "maxsize": 20,
            "maxage": 7,
            "maxbackups": 50,
            "localtime": true,
            "compress": true
        },
        "level": "warn"
    },
    "proxy" : {
        "rotate": {
            "filename": "var/log/proxy.log",
            "maxsize": 20,
            "maxage": 7,
            "maxbackups": 50,
            "localtime": true,
            "compress": true
        },
        "level": "warn"
    }
}
```
Root and proxy are the names of the loggers from which you can get the logger instance.  

The range of level is `debug`, `info`, `warn`, `error`, `dpanic`, `panic`, `fatal`, and the value is from low to high.  

2. Usage
```go
package main

import (
	"github.com/mallbook/logging"
	"go.uber.org/zap"
)

func main() {
	logger := logging.RLogger() // or logger := logging.Logger("root")

	for i := 0; i < 100; i++ {
		logger.Info("hello world", zap.String("key", "value"), zap.Int("age", 20))
		logger.Debug("hello china", zap.String("key", "value"), zap.Int("age", 20))
		logger.Error("hello error", zap.String("key", "value"), zap.Int("age", 30))
	}

    proxyLogger : logging.Logger("proxy")
	logger.Info("hello world", zap.String("key", "value"), zap.Int("age", 20))
}
```
