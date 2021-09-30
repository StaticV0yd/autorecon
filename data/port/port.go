package port

import (
	"github.com/StaticV0yd/autorecon/data/service"
)

type Port struct {
	protocol string
	portid   int
	state    string
	service  service.Service
}
