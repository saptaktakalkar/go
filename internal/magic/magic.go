// Copyright © 2017 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package magic

const (
	ext234SMagicOffL = 0x438
	ext234SMagicOffM = 0x439
	ext234SMagicValL = 0x53
	ext234SMagicValM = 0xef

	ext234SUUIDOff = 0x468
	ext234SUUIDLen = 16
)

func IdentifyPartitionMap(sniff []byte) string {
	return ""
}

func IdentifyPartition(sniff []byte) string {
	if sniff[ext234SMagicOffL] == ext234SMagicValL &&
		sniff[ext234SMagicOffM] == ext234SMagicValM {
		return "ext234"
	}
	return ""
}

func IdentifyFile(sniff []byte) string {
	return ""
}
