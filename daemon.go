package lust

import "time"

type DaemonHook func()

type Daemon struct {
	Interval time.Duration
	hooks    []DaemonHook
}
