package mvn

import (
	"io"
	"os/exec"
)

// Maven 表示本地的Maven
type Maven struct {
	homeDirectory string
}

// NewMaven 根据Home目录创建一个Maven
func NewMaven(homeDirectory string) (*Maven, error) {

}

func (x *Maven) GetRepoDirectory() {

}

func (x *Maven) FindGAV() {

}

func (x *Maven) GetExecPath() (string, error) {

}

type ExecOptions struct {
	Output io.Writer
}

func (x *Maven) Exec(args ...string) {
	path, err := x.GetExecPath()
	if err != nil {
		return err
	}
	command := exec.Command(path, args...)
	command.Stdout =
}
