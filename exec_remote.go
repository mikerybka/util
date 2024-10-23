package util

import (
	"fmt"
	"os/exec"
)

func ExecRemote(user, host, cmd string) ([]byte, error) {
	c := exec.Command("ssh", "-o", "StrictHostKeyChecking=no", fmt.Sprintf("%s@%s", user, host), cmd)
	return c.CombinedOutput()
	// err := AddToKnownHosts(user, host)
	// if err != nil {
	// 	return nil, err
	// }
	// config, err := SSHClientConfig(user)
	// if err != nil {
	// 	return nil, err
	// }
	// return RunSSHCommandWithKnownHostsCheck(host+":22", user, cmd, config.Auth)
}
