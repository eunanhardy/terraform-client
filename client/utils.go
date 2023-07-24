package client

import (
	"fmt"
	"os"
	"os/exec"
)

func CreateVarArgs(vars []string) []string {
	var_list := []string{}
	for _, v := range vars {
		var_list = append(var_list, fmt.Sprintf("-var=%s", v))
	}

	return var_list
}

func RunCommand(args ...string) (err error) {
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err1 := cmd.Run()
	if err1 != nil {
		return err1
	}

	return nil
}

func RunCommandWithOutput(args ...string) (string, error) {
	cmd := exec.Command(args[0], args[1:]...)

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
