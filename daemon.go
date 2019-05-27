package lust

import "time"

type DaemonHook func()

type Daemon struct {
	Interval time.Duration
	hooks    []DaemonHook
	//grpcService ?
}

func NewDaemonServer() {

}

func NewDaemonClient() {

}
