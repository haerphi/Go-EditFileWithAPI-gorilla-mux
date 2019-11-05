package sh

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

// RunShFile run a sh file at the path.
func RunShFile(path string) {
	cmd := exec.Command("bash", path)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
