package govirtlib

type vmOption string

const (
	// VERSION govirtlib version number
	VERSION string = "0.0.1-dev"
	// QEMUSystem connects to a QEMU system mode daemon
	QEMUSystem string = "qemu:///system"

	// disconnectedTimeout is how long to wait for disconnect cleanup to
	// complete
	// disconnectTimeout = 5 * time.Second

	nostate     = "NOSTATE"
	running     = "RUNNING"
	blocked     = "BLOCKED"
	paused      = "PAUSED "
	shutdown    = "SHUTDOWN"
	crashed     = "CRASHED "
	pmSuspended = "PMSUSPENDED"
	shutOff     = "SHUTOFF"

	// VMOptUUID ...
	VMOptUUID vmOption = "UUID"
	// VMOptName ...
	VMOptName vmOption = "Name"
)
