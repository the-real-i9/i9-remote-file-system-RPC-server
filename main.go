package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("rmdir", "i9")
	cmd.Dir = "i9FSHome"
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
