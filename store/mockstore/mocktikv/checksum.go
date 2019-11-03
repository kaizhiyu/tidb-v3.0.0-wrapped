// Copyright 2018 PingCAP, Inc.
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

package mocktikv

import (
	"fmt"

	"github.com/pingcap/kvproto/pkg/coprocessor"
	"github.com/pingcap/tidb/trace_util_0"
	"github.com/pingcap/tipb/go-tipb"
)

func (h *rpcHandler) handleCopChecksumRequest(req *coprocessor.Request) *coprocessor.Response {
	trace_util_0.Count(_checksum_00000, 0)
	resp := &tipb.ChecksumResponse{
		Checksum:   1,
		TotalKvs:   1,
		TotalBytes: 1,
	}
	data, err := resp.Marshal()
	if err != nil {
		trace_util_0.Count(_checksum_00000, 2)
		return &coprocessor.Response{OtherError: fmt.Sprintf("marshal checksum response error: %v", err)}
	}
	trace_util_0.Count(_checksum_00000, 1)
	return &coprocessor.Response{Data: data}
}

var _checksum_00000 = "store/mockstore/mocktikv/checksum.go"