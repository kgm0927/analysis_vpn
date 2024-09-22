package childprocess

// Package child_process contains common utilities for running NordVPN helper apps
// (eg. fileshare and norduser) as child process, rather than a system daemon.

// 패키지 childe_process는 데몬 시스템보다는 자식 프로세스로서, NorcVPN의 도우미 응용 프로그램을 돌리기 위해
// 일반적인 유틸리티(기능)을 저장해 놓는다.

type StartupErrorCode int

const (
	CodeAlreadyRunning StartupErrorCode = iota + 1
	CodeAlreadyRunningForOtherUser
	CodeFailedToCreateUnixSocket
	CodeMeshnetNotEnabled
	CodeAddressAlreadyInUse
	CodeFailedToEnable
	CodeUserNotInGroup
)

type ProcessStatus int

const (
	Running ProcessStatus = iota
	RunningForOtherUser
	NotRunning
)

type ChildProcessManager interface {

	// StartProcess starts the process
	// StartProcess는 프로세스를 시작시킨다.
	StartProcess() (StartupErrorCode, error)

	// StopProcess stops the process
	// StopProcess는 프로세스를 정지한다.
	StopProcess(disable bool) error

	// RestartProcess restarts the process
	// RestartProcess는 프로세스를 재시작한다.
	RestartProcess() error

	// ProcessStatus checks the status of process
	ProcessStatus() ProcessStatus
}
