package logging

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hotmall/commandline"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	defaultLogger *zap.Logger
	loggers       = make(map[string]*zap.Logger)
)

// Logger return the named zap.Logger instance, if no names return defaultLogger
func Logger(names ...string) *zap.Logger {
	if len(names) == 0 {
		return defaultLogger
	}
	name := names[0]
	if logger, ok := loggers[name]; ok {
		return logger
	}
	return defaultLogger
}

func init() {
	prefix := commandline.PrefixPath()
	configFile := prefix + "/etc/conf/logging.json"
	config, err := loadConfig(configFile)
	if err != nil {
		defaultLogger = initDefaultLogger()
		return
	}

	logPath := commandline.LogPath()
	for n, c := range config {
		level := zap.NewAtomicLevel()
		err = level.UnmarshalText([]byte(c.Level))
		if err != nil {
			// fail, not return, use InfoLevel
			fmt.Printf("Unmarshal level(%s) fail, err = %s", c.Level, err.Error())
			level.SetLevel(zapcore.InfoLevel)
		}

		var env Env
		if ok := env.unmarshalText([]byte(c.Env)); !ok {
			// fail, set prod env
			env.Set("prod")
		}

		var enc zapcore.Encoder
		if env == ProdEnv {
			enc = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		} else {
			enc = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		}

		c.FileName = logPath + "/" + c.FileName
		c.FileName = filepath.Clean(c.FileName)

		lumLogger := newLumLogger(c)
		w := zapcore.AddSync(lumLogger)
		core := zapcore.NewCore(enc, w, level)
		if n == "default" {
			// stdout 与 stderr 由 default 接管
			tee := zapcore.NewTee(core, newStdoutCore(enc), newStderrCore(enc))
			logger := zap.New(tee, zap.AddCaller())
			defaultLogger = logger
			continue
		}
		logger := zap.New(core, zap.AddCaller())
		if _, ok := loggers[n]; !ok {
			loggers[n] = logger
		}
	}

	if defaultLogger == nil {
		defaultLogger = initDefaultLogger()
	}

	// 重定向标准日志库输出日志到 defaultLogger
	zap.RedirectStdLog(defaultLogger)
}

func initDefaultLogger() *zap.Logger {
	logPath := commandline.LogPath()
	procName := commandline.ProcName()
	fileName := filepath.Clean(logPath + "/" + procName + ".log")

	r := loggerConfig{
		FileName:  fileName,
		MaxSize:   20,
		MaxAge:    7,
		LocalTime: true,
		Compress:  true,
		Level:     "info",
		Env:       "prod",
	}
	level := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	lumLogger := newLumLogger(r)
	w := zapcore.AddSync(lumLogger)
	enc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(enc, w, level)
	// 接管 stdout 与 stderr 的日志
	tee := zapcore.NewTee(core, newStdoutCore(enc), newStderrCore(enc))
	return zap.New(tee, zap.AddCaller())
}

func newStdoutCore(enc zapcore.Encoder) zapcore.Core {
	// info level enabler
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.InfoLevel
	})
	// write syncers
	stdoutSyncer := zapcore.Lock(os.Stdout)
	return zapcore.NewCore(enc, stdoutSyncer, infoLevel)
}

func newStderrCore(enc zapcore.Encoder) zapcore.Core {
	// error and fatal level enabler
	errorFatalLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.ErrorLevel || level == zapcore.FatalLevel
	})
	// write syncers
	stderrSyncer := zapcore.Lock(os.Stderr)
	return zapcore.NewCore(enc, stderrSyncer, errorFatalLevel)
}
