// Copyright (c) 2020 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin,!cgo

package interfaces

func init() {
	likelyHomeRouterIP = likelyHomeRouterIPDarwinExec
}
