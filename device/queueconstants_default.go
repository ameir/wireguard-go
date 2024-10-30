//go:build !android && !ios && !windows

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package device

// Reduce memory requirements -- copy from iOS
var (
	QueueStagedSize                   = 128
	QueueOutboundSize                 = 1024
	QueueInboundSize                  = 1024
	QueueHandshakeSize                = 1024
	PreallocatedBuffersPerPool uint32 = 1024
)

const MaxSegmentSize = 1700
