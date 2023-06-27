//go:build !go1.17
// +build !go1.17

package base

import (
	"fmt"
	"os"
	"os/exec"
)

// GoInstall go get path.
func GoInstall(path ...string) error {
	for _, p := range path {
		fmt.Printf("go get -u %s\n", p)
		cmd := exec.Command("go", "get", "-u", p)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}

// ⚠️ 看文件的头两行，这里对编译器的版本提出约束，指明这个文件只能在 go1.17 以下的版本的编译器上建时才会参与构建。
// 为啥会有这个约束？因为 go1.17 之前是通过 go get 下载安装可执行文件的。
