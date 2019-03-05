package logging

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mallbook/commandline"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.Logger
)

func init() {
	prefix := commandline.GetPrefixPath()
	configFile := prefix + "/etc/conf/zap.json"
	config, err := loadConfig(configFile)
	if err != nil {
		fmt.Printf("Load config file(%s) fail, err = %s", configFile, err.Error())
		os.Exit(1)
	}

	level := zap.NewAtomicLevel()
	err = level.UnmarshalText([]byte(config.Level))
	if err != nil {
		// fail, not return, use InfoLevel
		fmt.Printf("Unmarshal level(%s) fail, err = %s", config.Level, err.Error())
	}

	w := zapcore.AddSync(&config.Rotate)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		level,
	)
	logger = zap.New(core)
}

// Logger return zap.Logger instance
func Logger() *zap.Logger {
	return logger
}

type config struct {
	Rotate lumberjack.Logger `json:"rotate"`
	Level  string            `json:"level"`
}

func loadConfig(fileName string) (conf config, err error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	if err = json.Unmarshal(bytes, &conf); err != nil {
		return
	}

	return
}
