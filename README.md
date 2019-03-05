# logging
Logging package based on zap and lumberjack
## Quick Start

1. Create file ./etc/conf/zap.json and config it.
```json
{
    "rotate": {
        "filename": "var/log/foo.log",
        "maxsize": 20,
        "maxage": 7,
        "maxbackups": 50,
        "localtime": true,
        "compress": true
    },
    "level": "warn"
}
```

The range of level is `debug`, `info`, `warn`, `error`, `dpanic`, `panic`, `fatal`, and the value is from low to high.

2. Usage
```go
package main

import (
	"github.com/mallbook/logging"
	"go.uber.org/zap"
)

func main() {
	logger := logging.Logger()

	for i := 0; i < 100; i++ {
		logger.Info("helloworld", zap.String("key", "value"), zap.Int("age", 20))
		logger.Debug("hello china", zap.String("key", "value"), zap.Int("age", 20))
		logger.Error("hello error", zap.String("key", "value"), zap.Int("age", 30))
	}
}
```
