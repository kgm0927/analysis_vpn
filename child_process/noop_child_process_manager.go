package childprocess

type NoopchildProcessManager struct{}

func (c NoopchildProcessManager) StartProcess() (StartupErrorCode, error) {
	return 0, nil
}

func (c NoopchildProcessManager) StopProcess(bool) error {
	return nil
}

func (c NoopchildProcessManager) RestartProcess() error {
	return nil
}

func (c NoopchildProcessManager) ProcessStatus() ProcessStatus {
	return NotRunning
}
