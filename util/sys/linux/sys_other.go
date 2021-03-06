// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.
// +build !linux

package linux

import (
	"runtime"
	"github.com/pingcap/tidb/trace_util_0"

	// OSVersion returns version info of operation system.
	// for non-linux system will only return os and arch info.
)

func OSVersion() (osVersion string, err error) {
	trace_util_0.Count(_sys_other_00000, 0)
	osVersion = runtime.GOOS + "." + runtime.GOARCH
	return
}

var _sys_other_00000 = "util/sys/linux/sys_other.go"
