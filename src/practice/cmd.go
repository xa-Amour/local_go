package main

import (
	"fmt"
	"os"
	"os/exec"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/13.6.md
执行linux命令和程序
*/
/*func init() {
	fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
}*/
func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("lack param ?func=xxx")
		return
	}

	execute(args[1])
}

func execute(n string) {
	funs := map[string]func(){
		"cmd1" : cmd1,
		"cmd2" : cmd2,
		"cmd3" : cmd3,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}

func cmd1()  {
	env  := os.Environ()
	proc := &os.ProcAttr{
		Env:env,
		Files:[]*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	pid,err := os.StartProcess("/bin/ls",[]string{"ls","-l"},proc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pid)
}

func cmd2()  {
	env  := os.Environ()
	proc := &os.ProcAttr{
		Env:env,
		Files:[]*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	pid,err := os.StartProcess("/bin/ps",[]string{"ps","-e","-opid,ppid,comm"},proc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pid)
}

func cmd3()  {
	cmd := exec.Command("ls","-l")
	out,err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}