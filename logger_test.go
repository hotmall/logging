package logging

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestLogger(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestRLogger(c *C) {

	l1 := Logger("root")
	c.Assert(l1, NotNil)
	l2 := Logger("ROOT")
	c.Assert(l2, NotNil)
}

func (s *MySuite) TestLogger(c *C) {
	l1 := Logger("mylog")
	c.Assert(l1, NotNil)

	l2 := Logger("MYLOG")
	c.Assert(l2, NotNil)

	l3 := Logger("mylog1")
	c.Assert(l3, NotNil)
}
