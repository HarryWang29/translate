package util

// StringInSlice 检查字符串 a 是否在 list 中
func StringInSlice(a string, list []string) bool {
	if len(list) == 0 {
		return false
	}
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
