package utils

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

	"github.com/mattn/go-runewidth"
)

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return lines, err

}

func GetUserAtHost() ([]string, error) {
	shell_command := exec.Command("whoami")
	user, err := shell_command.Output()
	if err != nil {
		return nil, err
	}

	shell_command = exec.Command("hostname")
	host, err := shell_command.Output()
	if err != nil {
		return nil, err
	}

	user_host := []string{strings.TrimSpace(string(user)), strings.TrimSpace(string(host))}

	return user_host, nil
}

func GetDistroName() (string, error) {
	text, err := ReadFile("/etc/os-release")
	if err != nil {
		return "", err
	}

	distroName := "unknown"
	for _, line := range text {
		if strings.Contains(line, "PRETTY_NAME") {
			distroName = strings.Split(line, "=")[1]
		}
	}

	return strings.Trim(distroName, "\""), nil

}

func GetKernelVers() (string, error) {
	shell_command := exec.Command("uname", "-r")
	version, err := shell_command.Output()
	if err != nil {
		return "unknown", err
	}

	return strings.TrimSpace(string(version)), nil

}

// func GetShell() string {
// 	shell := os.Getenv("SHELL")
// 	if shell == "" {
// 		return "unknown"
// 	}
//
// 	return shell
//
// }
// TODO log errors...
func GetShell() (string, error) {

  shell_command := exec.Command("/bin/sh", "-c", "echo $SHELL")
  
  shell, err := shell_command.Output()
  if err != nil {
    return "unknown", err
  }
	return strings.TrimSpace(string(shell)), nil
}



func GetUpTime() (string, error) {
	shell_command := exec.Command("uptime", "-p")
	raw_uptime, err := shell_command.Output()
	if err != nil {
		return "", err
	}

	raw_uptime_str := strings.TrimSpace(string(raw_uptime))
	uptime := strings.Split(raw_uptime_str, "up ")[1]

	return uptime, nil
}

// taking care of multibyte runes
func GetMaxLogoLineLenght(lines []string) int {
	max := 0
	for _, line := range lines {
		width := runewidth.StringWidth(line)
		if width > max {
			max = width
		}
	}

	return max
}

func LenLogoLine(line string) int {
	return runewidth.StringWidth(line)
}

// thanks person who did this so chatgpt could steal it
// i'd like to give credit but i don't know who you are
func StripAnsiEscapeSeq(line string) string {

	return strings.Map(func(char rune) rune {
		if char >= 32 && char <= 126 {
			return char
		}
		return -1
	}, line)

}

func MapStrSl(strSl []string, f func(string) string) []string {
	stSlMapped := make([]string, len(strSl))
	for i, str := range strSl {
		stSlMapped[i] = f(str)
	}
	return stSlMapped
}

func ClearTerm() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
