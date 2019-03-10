package logging

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestLogger(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestRLogger(c *C) {
	_, ok := RLogger()
	c.Check(ok, Equals, true)

	_, ok = Logger("root")
	c.Check(ok, Equals, true)

	_, ok = Logger("ROOT")
	c.Check(ok, Equals, false)
}

func (s *MySuite) TestLogger(c *C) {
	_, ok := Logger("mylog")
	c.Check(ok, Equals, true)

	_, ok = Logger("MYLOG")
	c.Check(ok, Equals, false)

	_, ok = Logger("mylog1")
	c.Check(ok, Equals, false)
}
