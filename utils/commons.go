package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func SanitizeServerUrl(tsaServer string) string {
	tsaServeSanitized := strings.ReplaceAll(tsaServer, "/", "_")
	tsaServeSanitized = strings.ReplaceAll(tsaServeSanitized, ".", "_")
	tsaServeSanitized = strings.ReplaceAll(tsaServeSanitized, ":", "_")
	return tsaServeSanitized
}

func RunCommand(cmdName string, directory string) (string, error) {
	split := strings.Split(cmdName, " ")
	cmd := exec.Command(split[0], split[1:len(split)]...)
	if directory != "" {
		cmd.Dir = directory
	}
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
