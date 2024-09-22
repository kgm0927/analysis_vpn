package childprocess

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"google.golang.org/grpc/status"
)

type ProcessClient interface {
	Ping(nowait bool) error
	Stop(disable bool) error
	Restart() error
}

type GRPCChildProcessManager struct {
	processClient     ProcessClient
	processBinaryPath string
}

func NewGRPCChildProcessManager(processClient ProcessClient, processBinaryPath string) *GRPCChildProcessManager {

	return &GRPCChildProcessManager{
		processClient:     processClient,
		processBinaryPath: processBinaryPath,
	}
}

func (g *GRPCChildProcessManager) StartProcess() (StartupErrorCode, error) {
	errChan := make(chan error)

	go func() {
		// #nosec G204 -- arg values are known before even running the program
		//				  인자 값은 프로그램 실행 전에 정의됨.

		err := exec.Command(g.processBinaryPath).Run()
		errChan <- err
	}()

	pingChan := make(chan error)
	// Start another goroutine where we ping the WaitForReady option, so that server has time to start up before we run
	// the actual command.
	// WaitForReady 옵션에 핑을 날릴 때 다른 고루틴을 시작해서, 우리가 코맨드를 실행하기 전 서버가 시작할 시간을 갖는다.

	go func() {
		err := g.processClient.Ping(false)
		pingChan <- err
	}()

	select {
	case err := <-errChan:
		if err == nil {
			return 0, fmt.Errorf("process finished unexpectedly")
		}
		var exiterr *exec.ExitError
		// 커맨드를 통해 성공적으로 exit 하지 못했다고 알림.

		if errors.As(err, &exiterr) {
			exitCode := StartupErrorCode(exiterr.ExitCode())
			return exitCode, nil
		}
		return 0, fmt.Errorf("failed to start the process: %w", err)

	case err := <-pingChan:
		if err != nil {
			return 0, fmt.Errorf("failed to ping the process after starting: %w", err)
		}

		// Process was started and pinged successfilly.
		// 프로세스는 시작되었고 성공적으로 핑 됨.
		return 0, nil

	}
}

func (g *GRPCChildProcessManager) StopProcess(disable bool) error {
	err := g.processClient.Restart()
	if err != nil {
		return fmt.Errorf("restarting process: %w", err)
	}
	return nil
}

func (g *GRPCChildProcessManager) ProcessStatus() ProcessStatus {
	err := g.processClient.Ping(true)
	if err != nil {
		if strings.Contains(status.Convert(err).Message(), "permission denied") {
			return RunningForOtherUser
		}
		return NotRunning
	}
	return Running
}
