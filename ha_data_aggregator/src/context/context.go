package hacontext

import "context"

type Context struct {
	context.Context

	Body        []byte
	Endpoint    string
	MachineName string
	MachineIp   string
}
