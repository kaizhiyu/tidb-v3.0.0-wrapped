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

package ddl

import (
	"sync"

	"github.com/ngaut/pools"
	"github.com/pingcap/errors"
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/tidb/sessionctx"
	"github.com/pingcap/tidb/trace_util_0"
	"github.com/pingcap/tidb/util/logutil"
	"github.com/pingcap/tidb/util/mock"
)

// sessionPool is used to new session.
type sessionPool struct {
	mu struct {
		sync.Mutex
		closed bool
	}
	resPool *pools.ResourcePool
}

func newSessionPool(resPool *pools.ResourcePool) *sessionPool {
	trace_util_0.Count(_session_pool_00000, 0)
	return &sessionPool{resPool: resPool}
}

// get gets sessionctx from context resource pool.
// Please remember to call put after you finished using sessionctx.
func (sg *sessionPool) get() (sessionctx.Context, error) {
	trace_util_0.Count(_session_pool_00000, 1)
	if sg.resPool == nil {
		trace_util_0.Count(_session_pool_00000, 5)
		return mock.NewContext(), nil
	}

	trace_util_0.Count(_session_pool_00000, 2)
	sg.mu.Lock()
	if sg.mu.closed {
		trace_util_0.Count(_session_pool_00000, 6)
		sg.mu.Unlock()
		return nil, errors.Errorf("sessionPool is closed.")
	}
	trace_util_0.Count(_session_pool_00000, 3)
	sg.mu.Unlock()

	// no need to protect sg.resPool
	resource, err := sg.resPool.Get()
	if err != nil {
		trace_util_0.Count(_session_pool_00000, 7)
		return nil, errors.Trace(err)
	}

	trace_util_0.Count(_session_pool_00000, 4)
	ctx := resource.(sessionctx.Context)
	ctx.GetSessionVars().SetStatusFlag(mysql.ServerStatusAutocommit, true)
	ctx.GetSessionVars().InRestrictedSQL = true
	return ctx, nil
}

// put returns sessionctx to context resource pool.
func (sg *sessionPool) put(ctx sessionctx.Context) {
	trace_util_0.Count(_session_pool_00000, 8)
	if sg.resPool == nil {
		trace_util_0.Count(_session_pool_00000, 10)
		return
	}

	// no need to protect sg.resPool, even the sg.resPool is closed, the ctx still need to
	// put into resPool, because when resPool is closing, it will wait all the ctx returns, then resPool finish closing.
	trace_util_0.Count(_session_pool_00000, 9)
	sg.resPool.Put(ctx.(pools.Resource))
}

// close clean up the sessionPool.
func (sg *sessionPool) close() {
	trace_util_0.Count(_session_pool_00000, 11)
	sg.mu.Lock()
	defer sg.mu.Unlock()
	// prevent closing resPool twice.
	if sg.mu.closed || sg.resPool == nil {
		trace_util_0.Count(_session_pool_00000, 13)
		return
	}
	trace_util_0.Count(_session_pool_00000, 12)
	logutil.Logger(ddlLogCtx).Info("[ddl] closing sessionPool")
	sg.resPool.Close()
	sg.mu.closed = true
}

var _session_pool_00000 = "ddl/session_pool.go"
