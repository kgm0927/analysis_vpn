package distro

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
Package disto provides information about the current Linux distribution
*/

// 패키지 distro는 현 리눅스 구분에 관한 정보를 준다.

const (
	kernelName    = "Linux"
	osReleaseFile = "/etc/os-release"
)

// osRelease represents contents of /etc/os-release
// osRelease는 /etc/os-release의 내용을 반환한다.

type osRelease struct {
	Name       string
	PrettyName string
}

func (o *osRelease) UnmarshalText(text []byte) error {
	for _, line := range bytes.Split(bytes.TrimSpace(text), []byte("\n")) {
		key, value, ok := bytes.Cut(line, []byte("="))
		if !ok {
			continue
		}

		isDoubleQuote := func(r rune) bool {
			return r == '"'
		}

		switch string(key) {
		case "NAME":
			o.Name = string(bytes.TrimFunc(value, isDoubleQuote))

		case "PRETTY_NAME":
			o.PrettyName = string(bytes.TrimFunc(value, isDoubleQuote))

		default:
			// ignore undefined fields
		}
	}
	return nil
}

// ReleaseName of the currently running distribution.
// ReleaseName은 현재 구분을 하는 중이다.

func ReleaseName() (string, error) {
	data, err := os.ReadFile(osReleaseFile)
	if err != nil {
		return "", err
	}

	var release osRelease
	if err := (&release).UnmarshalText(data); err != nil {
		return "", err
	}

	return release.Name, nil

}

// ReleasePrettyName of the currently running distribution.
func ReleasePrettyName() (string, error) {
	data, err := os.ReadFile(osReleaseFile)
	if err != nil {
		return "", err
	}

	var release osRelease
	if err := (&release).UnmarshalText(data); err != nil {
		return "", err
	}

	return release.PrettyName, nil
}

// KernalName of the currently running kernel.
func KernalName() string {
	return
}

// KernalFull name of the currently running kernel.
func KernalFull() string {
	return
}

// uname returns operating system information from uname executable
// uname은 실행 가능한 uname 운영체제 정보를 반환한다.
func uname(flags string) string {
	// #nosec G204 -- input is known before running the program
	// 					입력은 프로그램을 실행하기 전부터 정해져 있음.

	out, _ := exec.Command("sh", "-c", fmt.Sprintf("uname %s", flags)).Output()
	trimmed := strings.Trim
}
