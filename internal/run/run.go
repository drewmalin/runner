package run

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/drewmalin/runner/internal/cmdcontext"
)

func run(context *cmdcontext.Context) error {
	return execCommand("ls", []string{"-l", "-a", "-h"}, context)
}

// execCommand executes the provided command, passing the provided arguments to the call. Any
// output will be continuously streamed back to the standard out of the 'runner' parent command
// until the executed process completes.
func execCommand(command string, args []string, context *cmdcontext.Context) error {
	cmd := exec.Command(command, args...)
	
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmdScanner := bufio.NewScanner(cmdReader)

	done := make(chan bool)
	go func() {
		for cmdScanner.Scan() {
			fmt.Fprintln(context.VerboseOut, cmdScanner.Text())
		}
		done <- true
	}()

	cmd.Start()
	if err != nil {
		return err
	}
	fmt.Fprintln(context.Out, fmt.Sprintf("[PID: %d] - Process started", cmd.Process.Pid))

	if err := cmd.Wait(); err != nil {
		return err
	}
	<- done

	fmt.Fprintln(context.Out, fmt.Sprintf("[PID: %d] - Process ended", cmd.Process.Pid))
	return nil
}