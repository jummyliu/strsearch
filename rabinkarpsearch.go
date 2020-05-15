package strsearch

// const primeRK = 16777619
const primeRK = 256

// RabinKarpSearch 基于 RabinKarp 的字符串匹配算法
func RabinKarpSearch(str, pattern string) (index int) {
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
	// 1. 计算hash
	hash, pow := hashStr(pattern)
	var sHash uint
	for i := 0; i < pL; i++ {
		sHash = sHash*primeRK + uint(str[i])
	}
	if sHash == hash && str[:pL] == pattern {
		return 0
	}
	for i := pL; i < sL; {
		sHash *= primeRK
		sHash -= pow * uint(str[i-pL])
		sHash += uint(str[i])
		i++
		if sHash == hash && str[i-pL:i] == pattern {
			return i-pL
		}
	}
	return -1
}

func hashStr(str string) (hash, pow uint) {
	sLen := len(str)
	for i := 0; i < sLen; i++ {
		hash = hash*primeRK + uint(str[i])
	}
	var sq uint = primeRK
	pow = 1
	for i := sLen; i > 0; i >>= 1 {
		if i%2 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return
}