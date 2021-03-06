// Copyright 2013 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author：polaris	studygolang@gmail.com

package process

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ExecutableDir() (string, error) {
	var (
		dir, p string
		err    error
	)
	pid := os.Getpid()
	lnk := "/proc/" + strconv.Itoa(pid) + "/exe"
	p, err = os.Readlink(lnk)
	if err != nil {
		return "", err
	}
	dir = filepath.Dir(p)
	dir = strings.Replace(dir, "\\", "/", -1)
	return dir, nil
}
