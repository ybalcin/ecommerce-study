package common

import "strconv"

// ValueOfSlice returns empty if index out of range
func ValueOfSlice(index int, slice []string) string {
	for i, v := range slice {
		if i == index {
			return v
		}
	}
	return ""
}

// StringToInt returns zero if val invalid
func StringToInt(val string) int {
	v, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}

	return v
}
