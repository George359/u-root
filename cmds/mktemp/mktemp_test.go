// Copyright 2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/u-root/u-root/pkg/testutil"
)

type test struct {
	flags      []string
	out        string
	stdErr     string
	exitStatus int
}

func TestMkTemp(t *testing.T) {

	var tests = []test{
		{
			flags:      []string{},
			out:        "/tmp/",
			stdErr:     "",
			exitStatus: 0,
		},
		{
			flags:      []string{"-d"},
			out:        "/tmp",
			stdErr:     "",
			exitStatus: 0,
		},
		{
			flags:      []string{"-d", "-p", "foo"},
			out:        "/tmp/foo",
			stdErr:     "",
			exitStatus: 0,
		},
		{
			flags:      []string{"foo"},
			out:        "",
			stdErr:     "",
			exitStatus: 1,
		},
	}

	// Table-driven testing
	for _, tt := range tests {
		var out, stdErr bytes.Buffer
		cmd := testutil.Command(t, tt.flags...)
		cmd.Stdout = &out
		cmd.Stderr = &stdErr
		err := cmd.Run()

		if !strings.Contains(out.String(), tt.out) {
			t.Errorf("stdout got:\n%s\nwant:\n%s", out.String(), tt.out)
		}

		if !strings.Contains(stdErr.String(), tt.stdErr) {
			t.Errorf("stderr got:\n%s\nwant:\n%s", stdErr.String(), tt.stdErr)
		}

		if tt.exitStatus == 0 && err != nil {
			t.Errorf("expected to exit with %d, but exited with err %s", tt.exitStatus, err)
		}
	}
}

func TestMain(m *testing.M) {
	testutil.Run(m, main)
}
