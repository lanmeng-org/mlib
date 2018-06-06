package utils

import (
	"strings"
	"strconv"
)

func IP2long(ip string) (ipLong uint32) {
	ips := strings.Split(ip, ".")
	if len(ips) != 4 {
		return 0
	}

	for i:=0;i<4;i++ {
		ipInt, err := strconv.Atoi(ips[i])
		if err != nil || ipInt < 0 || ipInt >255 {
			return 0
		}

		offset := byte((3 - i) * 8)
		ipLong += uint32(ipInt * (0x1 << offset))
	}

	return ipLong
}

func Long2IP(ipInt uint32) string {
	var ipParts [4]string
	for i:=0;i<4;i++ {
		ipPart := ipInt << byte(i * 8) >> byte(3 * 8)

		ipParts[i] = strconv.Itoa(int(ipPart))
	}

	return strings.Join(ipParts[:], ".")
}