// Copyright © 2024 The Hot Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logging

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type loggerConfig struct {
	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.  It uses <processname>-lumberjack.log in
	// os.TempDir() if empty.
	FileName string `json:"filename" yaml:"filename"`

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

	// Log Level. The default is to use Info
	Level string `json:"level" yaml:"level"`

	// Env, either prod or dev
	Env string `json:"env" yaml:"env"`
}

type config map[string]loggerConfig

func loadConfig(fileName string) (conf config, err error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return
	}

	if err = json.Unmarshal(bytes, &conf); err != nil {
		return
	}

	return
}

func newLumLogger(c loggerConfig) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   c.FileName,
		MaxSize:    c.MaxSize,
		MaxAge:     c.MaxAge,
		MaxBackups: c.MaxBackups,
		LocalTime:  c.LocalTime,
		Compress:   c.Compress,
	}
}

var errUnmarshalNilLevel = errors.New("can't unmarshal a nil *Env")

type Env int8

const (
	UnknowEnv Env = iota - 1
	ProdEnv
	DevEnv
)

// String returns a lower-case ASCII representation of the env.
func (e Env) String() string {
	switch e {
	case ProdEnv:
		return "prod"
	case DevEnv:
		return "dev"
	default:
		return fmt.Sprintf("unrecognized env: %d", e)
	}
}

func (e *Env) UnmarshalText(text []byte) error {
	if e == nil {
		return errUnmarshalNilLevel
	}
	if !e.unmarshalText(text) {
		return fmt.Errorf("unrecognized env: %s", string(text))
	}
	return nil
}

func (e *Env) unmarshalText(text []byte) bool {
	var env string
	if len(text) == 0 {
		env = "prod"
	} else {
		env = string(bytes.ToLower(text))
	}
	switch env {
	case "prod": // make the zero value useful
		*e = ProdEnv
	case "dev":
		*e = DevEnv
	default:
		return false
	}
	return true
}

// Set sets the level for the flag.Value interface.
func (e *Env) Set(s string) error {
	return e.UnmarshalText([]byte(s))
}

// Get gets the level for the flag.Getter interface.
func (e *Env) Get() interface{} {
	return *e
}

func confile(prefix string) string {
	return filepath.Join(prefix, "etc", "conf", "logging.json")
}
