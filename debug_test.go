package dump_test

import (
	"testing"

	dump "github.com/178inaba/go.dump"
)

func TestDump(t *testing.T) {
	var v struct {
		num int
		msg string
	}
	v.num = 100
	v.msg = "Dump!"

	dump.Dump(v)
}
