package go_learn

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestEnv(t *testing.T) {
	os.Setenv("WEB", "https://coolshell.cn") //设置环境变量
	println(os.Getenv("WEB"))                //读出来
	for _, env := range os.Environ() { //穷举环境变量
		e := strings.Split(env, "=")
		println(e[0], "=", e[1])
	}
}

func TestExec(t *testing.T) {
	cmd := exec.Command("ifconfig")
	out, err := cmd.Output()
	if err != nil {
		println("Command Error!", err.Error())
		return
	}
	fmt.Println(string(out))
}
