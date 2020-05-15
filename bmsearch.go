package strsearch

// BMSearch 基于 BM 的字符串匹配算法
func BMSearch(str, pattern string) (index int) {
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
	// 1. 坏字符算法
	// 当前匹配位置 - 上一次出现位置
	badChar := make([]int, 256)
	for i := range badChar {
		badChar[i] = -1
	}
	for i := 0; i < pL; i++ {
		badChar[pattern[i]] = i
	}

	// 2. 好后缀算法
	// 后缀的最后一个位置 - 前缀的最后一个位置 或者 后缀第一个位置 - 前缀第一个位置
	goodSuffix := make([]int, pL + 1)
	next := make([]int, pL + 1)
	i, j := pL - 1, pL
	next[i+1] = j
	for i >= 0 {
		for j < pL && pattern[j] != pattern[i] {
			j = next[j] + 1
		}
		i--
		j--
		next[i+1] = j
	}
	i, j = pL - 1, pL - 1
	goodSuffix[pL] = 0
	for ; i >= 0; i-- {
		if j < 0 {
			goodSuffix[i] = goodSuffix[i + 1]
			continue
		}
		for ; j >= 0 && next[j] >= i; j-- {
		}
		// next[j + 1] - j
		goodSuffix[i] = next[j + 1] - j
	}
	// 效率有问题...暴力匹配
	// goodSuffix := make([]int, pL)
	// i, j, k := pL - 1, pL - 2, 0
	// goodSuffix[i] = 0
	// i--
	// j--
	// for ; i >= 0; i-- {
	// 	for ; j >= 0 && k < pL - i; j-- {
	// 		if pattern[j] == pattern[pL - 1 - k] {
	// 			k++
	// 		} else {
	// 			k = 0
	// 		}
	// 	}
	// 	// i - (j + 1) + pL - i - k
	// 	// pL - (j + 1) - k
	// 	goodSuffix[i] = pL - (j + 1) -k
	// }

	// 3. 跳转 = max(坏字符， 好后缀)
	i, j = 0, pL - 1
	for i <= sL - pL {
		for ; j >= 0 && pattern[j] == str[i + j]; j-- {
		}
		if j == -1 {
			return i
		}
		i += max(j - badChar[str[i + j]], goodSuffix[j + 1])
		j = pL - 1
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// BMSearch2 基于 BM 的字符串匹配算法
// 优化 next, goodSuffix
func BMSearch2(str, pattern string) (index int) {
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
	// 1. 坏字符算法
	// 当前匹配位置 - 上一次出现位置
	badChar := make([]int, 256)
	for i := range badChar {
		badChar[i] = -1
	}
	for i := 0; i < pL; i++ {
		badChar[pattern[i]] = i
	}

	// 2. 好后缀算法
	// 后缀的最后一个位置 - 前缀的最后一个位置 或者 后缀第一个位置 - 前缀第一个位置
	goodSuffix := make([]int, pL)
	next := make([]int, pL)
	i, j := pL - 1, pL
	next[i] = i
	for i >= 0 {
		for j < pL && pattern[j] != pattern[i] {
			j = next[j] + 1
		}
		i--
		j--
		next[i+1] = j
	}
	i, j = pL - 1, pL - 1
	goodSuffix[i] = 0
	for ; i > 0; i-- {
		if j < 0 {
			goodSuffix[i - 1] = goodSuffix[i]
			continue
		}
		for ; j >= 0 && next[j] >= i; j-- {
		}
		// next[j + 1] - j
		goodSuffix[i - 1] = next[j + 1] - j
	}
	// 3. 跳转 = max(坏字符， 好后缀)
	i, j = 0, pL - 1
	for i <= sL - pL {
		for ; j >= 0 && pattern[j] == str[i + j]; j-- {
		}
		if j == -1 {
			return i
		}
		i += max(j - badChar[str[i + j]], goodSuffix[j])
		j = pL - 1
	}
	return -1
}