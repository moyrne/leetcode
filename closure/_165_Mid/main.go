package _165_Mid

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	count1 := strings.Count(version1, ".")
	count2 := strings.Count(version2, ".")

	base := 1
	if count2 > count1 {
		version1, version2 = version2, version1
		base = -1
	}

	v1, v2 := version1, version2
	for len(v1) > 0 {
		subv1, idx1 := item(v1)
		subv2, idx2 := item(v2)

		if subv1 < subv2 {
			return -base
		}
		if subv1 > subv2 {
			return base
		}

		v1 = v1[idx1+1:]
		if len(v2) > 0 {
			v2 = v2[idx2+1:]
		}
	}

	return 0
}

func item(s string) (int, int) {
	if s == "" {
		return 0, 0
	}

	var sub string
	idx := strings.Index(s, ".")
	if idx == -1 {
		idx = len(s) - 1
		sub = s
	} else {
		sub = s[:idx]
	}

	subv, _ := strconv.Atoi(sub)
	return subv, idx
}
