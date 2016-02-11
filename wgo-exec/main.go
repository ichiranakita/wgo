/*
Copyright 2015 Google Inc. All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
The wgo tool is a small wrapper around the go tool. It adds the concept
of a workspace, in addition to that of GOPATH, and several new commands
to manage that workspace.
*/
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/skelterjohn/wgo/workspaces"
)

func main() {
	w, err := workspaces.GetCurrentWorkspace()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: wgo-exec COMMAND [ARG+]")
		os.Exit(1)
	}

	path := os.Getenv("PATH")
	gopath := w.Gopath(true)
	sep := string(os.PathListSeparator)
	for _, p := range strings.Split(gopath, sep) {
		path = filepath.Join(p, "bin") + sep + path
	}

	os.Setenv("GOPATH", gopath)
	os.Setenv("PATH", path)
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
