package common

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"sync"
	"time"
)

// note:
// c.CombinedOutput() 适合
// c.StdinPipe()  & c.StdoutPipe() c.StderrPipe() 适合多次输入输出的命令
func Demo() {

	c := exec.Command("/bin/bash", "-c", "python3")
	stdin, _ := c.StdinPipe()
	stdout, _ := c.StdoutPipe()

	if err := c.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return
	}

	stdin.Write([]byte(`print("hello")`))
	stdin.Close()

	var out []byte
	out, _ = ioutil.ReadAll(stdout)
	stdout.Close()

	if err := c.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return
	}

	fmt.Println("Execute fanish:", string(out))

}

func Exe() {
	var wg sync.WaitGroup

	ctx, cancelFunc := context.WithCancel(context.TODO())

	wg.Add(1)
	go func() {
		var (
			output []byte
			err    error
		)

		cmd := exec.CommandContext(ctx, "bash", "-c", "sleep 2; pwd")

		if output, err = cmd.CombinedOutput(); err != nil {
			log.Println(err)
		}

		fmt.Println(string(output))

		wg.Done()
	}()

	time.Sleep(1 * time.Second)
	cancelFunc()

	wg.Wait()
}
