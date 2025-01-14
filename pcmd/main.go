// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package pcmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/uber/prototool/internal/cmd"
)

func GetCommand() *cobra.Command {
	var exitCode int
	RootCommand := cmd.GetRootCoomand(false, &exitCode, os.Args[1:], os.Stdin, os.Stdout, os.Stderr)
	RootCommand.Run = func(c *cobra.Command, args []string) {
		if err := cmd.CheckOS(); err != nil {
			exitCode = cmd.PrintAndGetErrorExitCode(err, os.Stdout)
		}
	}
	RootCommand.PersistentPostRun = func(c *cobra.Command, args []string) {
		if exitCode != 0 {
			os.Exit(exitCode)
		}
	}
	return RootCommand
}
