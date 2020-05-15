package strsearch

// CommonSearch 朴素字符串匹配算法
func CommonSearch(str, pattern string) (index int) {
	sL := len(str)
	pL := len(pattern)
	if sL < pL {
		return -1
	}
	if sL == pL {
		if str == pattern {
			return 0
		}
		return -1
	}
	if pL == 0 {
		return 0
	}
	for i := 0; i <= sL-pL; i++ {
		j := 0
		for ; j < pL && str[i+j] == pattern[j]; j++ {
		}
		if j == pL {
			return i
		}
	}
	return -1
}
