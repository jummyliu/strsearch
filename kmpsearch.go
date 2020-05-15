package strsearch

// KMPSearch 基于 KMP 的字符串匹配算法
func KMPSearch(str, pattern string) (index int) {
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
	i, j := 0, -1
	next := make([]int, pL + 1)
	next[i] = j
	for i < pL {
		for j >= 0 && pattern[j] != pattern[i] {
			j = next[j]
		}
		i++
		j++
		next[i] = j
	}

	i, j = 0, 0
	for i < sL {
		for j >= 0 && pattern[j] != str[i] {
			j = next[j]
		}
		i++
		j++
		if j == pL {
			return i - j
		}
	}
	return -1
}

// KMPSearch2 基于 KMP 的字符串匹配算法
// 优化 next 数组
func KMPSearch2(str, pattern string) (index int) {
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
	i, j := 0, -1
	next := make([]int, pL + 1)
	next[i] = j
	for i < pL {
		for j >= 0 && pattern[j] != pattern[i] {
			j = next[j]
		}
		i++
		j++
		// 优化，如果当前位置需要匹配的字符，跟上一个位置的字符是一样的，则 next[i] = next[j]
		if i < pL && pattern[j] == pattern[i] {
			next[i] = next[j]
		} else {
			next[i] = j
		}
	}

	i, j = 0, 0
	for i < sL {
		for j >= 0 && pattern[j] != str[i] {
			j = next[j]
		}
		i++
		j++
		if j == pL {
			return i - j
		}
	}
	return -1
}