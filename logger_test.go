package logging_test

import (
	"testing"

	"github.com/mallbook/logging"
)

func TestRLogger(t *testing.T) {
	logger, ok := logging.RLogger()
	if !ok {
		t.Error("Expect ok is true, but false")
	}
	logger.Info("Hello world")
}

func TestLogger(t *testing.T) {
	_, ok := logging.Logger("mylog")
	if ok {
		t.Error("Get mylog logger, expect return not ok, but return ok")
	}
}
