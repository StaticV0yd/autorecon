package host

import (
	"github.com/StaticV0yd/autorecon/data/port"
)

type Host struct {
	address string
	ports   []port.Port
}
