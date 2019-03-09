package logging

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"github.com/mallbook/commandline"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loggers = make(map[string]*zap.Logger)
)

func init() {
	prefix := commandline.PrefixPath()
	configFile := prefix + "/etc/conf/logging.json"
	logPath := commandline.LogPath()
	config, err := loadConfig(configFile)
	if err != nil {
		fmt.Printf("Load config file(%s) fail, err = %s", configFile, err.Error())
		os.Exit(1)
	}

	for n, c := range config {
		level := zap.NewAtomicLevel()
		err = level.UnmarshalText([]byte(c.Level))
		if err != nil {
			// fail, not return, use InfoLevel
			fmt.Printf("Unmarshal level(%s) fail, err = %s", c.Level, err.Error())
		}

		c.Rotate.Filename = logPath + c.Rotate.Filename

		lumLogger := newLumLogger(c.Rotate)
		w := zapcore.AddSync(lumLogger)
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			w,
			level,
		)
		logger := zap.New(core)
		if _, ok := loggers[n]; !ok {
			loggers[n] = logger
		}
	}
}

// RLogger return root zap.Logger instance
func RLogger() (logger *zap.Logger, ok bool) {
	logger, ok = loggers["root"]
	return
}

// Logger return the named zap.Logger instance
func Logger(name string) (logger *zap.Logger, ok bool) {
	logger, ok = loggers[name]
	return
}

type rotateConfig struct {
	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.  It uses <processname>-lumberjack.log in
	// os.TempDir() if empty.
	Filename string `json:"filename" yaml:"filename"`

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `json:"maxsize" yaml:"maxsize"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `json:"maxage" yaml:"maxage"`

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int `json:"maxbackups" yaml:"maxbackups"`

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool `json:"localtime" yaml:"localtime"`

	// Compress determines if the rotated log files should be compressed
	// using gzip.
	Compress bool `json:"compress" yaml:"compress"`
}

type loggerConfig struct {
	Rotate rotateConfig `json:"rotate" yaml:"rotate"`
	Level  string       `json:"level" yaml:"level"`
}

type config map[string]loggerConfig

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

func newLumLogger(c rotateConfig) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   c.Filename,
		MaxSize:    c.MaxSize,
		MaxAge:     c.MaxAge,
		MaxBackups: c.MaxBackups,
		LocalTime:  c.LocalTime,
		Compress:   c.Compress,
	}
}
