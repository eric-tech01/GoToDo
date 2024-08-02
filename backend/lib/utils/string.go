package utils

import "strings"

func StringSliceContainsAny(arr []string, sarr ...string) bool {
	for _, a := range arr {
		for _, s := range sarr {
			if strings.Compare(a, s) == 0 {
				return true
			}
		}
	}
	return false
}
