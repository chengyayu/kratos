//go:build go1.17
// +build go1.17

package base

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GoInstall go get path.
func GoInstall(path ...string) error {
	for _, p := range path {
		if !strings.Contains(p, "@") {
			p += "@latest"
		}
		fmt.Printf("go install %s\n", p)
		cmd := exec.Command("go", "install", p)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}

// ⚠️ 看文件的头两行，这里对编译器的版本提出约束，指明这个文件只能在 go1.17 及更新的版本的编译器上建时才会参与构建。
// 为啥会有这个约束？因为 go1.17 开始 go install 被指定为下载安装可执行文件的指令。
