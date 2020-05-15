package strsearch

import (
	"testing"
)

type testStruct struct {
	str string
	pattern string
	result int
}

var testArr []testStruct = []testStruct{
	{
		str: "hellobbbbb",
		pattern: "ababa",
		result: -1,
	},
	{
		str: "ab",
		pattern: "abc",
		result: -1,
	},
	{
		str: "abcabdacqqqdc",
		pattern: "abcabdacqqq",
		result: 0,
	},
	{
		str: "abcbcabcbcabcabc",
		pattern: "cabcab",
		result: 9,
	},
	{
		str: "abcabdacqqqdcabcabdace",
		pattern: "abcabdace",
		result: 13,
	},
	{
		str: "hello",
		pattern: "hello",
		result: 0,
	},
	{
		str: "abcabdacdabcabdacf",
		pattern: "abcabdace",
		result: -1,
	},
	{
		str: "abcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcaabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdebcdabcabcabcdabcabcabcdabcabcabcdeabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcd",
		pattern: "abcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcde",
		result: 2881,
	},
	{
		str: "abcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdaabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcaabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdeabcabcabcdebcdabcabcabcdabcabcabcdabcabcabcdeabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcabcdabcabcdabcdaabcabcdabcabcabcdabcabcabcdabcabcabcd",
		pattern: "abcdabcda",
		result: 3143,
	},
	{
		str: "ABACHOIEHGAAABAABAAABAAASDEOAAABAA",
		pattern: "AAABAABAAABAAA",
		result: 10,
	},
	{
		str: "hello",
		pattern: "",
		result: 0,
	},
	{"", "", 0},
	{"", "a", -1},
	{"", "foo", -1},
	{"fo", "foo", -1},
	{"foo", "foo", 0},
	{"oofofoofooo", "f", 2},
	{"oofofoofooo", "foo", 4},
	{"barfoobarfoo", "foo", 3},
	{"foo", "", 0},
	{"foo", "o", 1},
	{"abcABCabc", "A", 3},
	// cases with one byte strings - test special case in Index()
	{"", "a", -1},
	{"x", "a", -1},
	{"x", "x", 0},
	{"abc", "a", 0},
	{"abc", "b", 1},
	{"abc", "c", 2},
	{"abc", "x", -1},
	// test special cases in Index() for short strings
	{"", "ab", -1},
	{"bc", "ab", -1},
	{"ab", "ab", 0},
	{"xab", "ab", 1},
	{"xab"[:2], "ab", -1},
	{"", "abc", -1},
	{"xbc", "abc", -1},
	{"abc", "abc", 0},
	{"xabc", "abc", 1},
	{"xabc"[:3], "abc", -1},
	{"xabxc", "abc", -1},
	{"", "abcd", -1},
	{"xbcd", "abcd", -1},
	{"abcd", "abcd", 0},
	{"xabcd", "abcd", 1},
	{"xyabcd"[:5], "abcd", -1},
	{"xbcqq", "abcqq", -1},
	{"abcqq", "abcqq", 0},
	{"xabcqq", "abcqq", 1},
	{"xyabcqq"[:6], "abcqq", -1},
	{"xabxcqq", "abcqq", -1},
	{"xabcqxq", "abcqq", -1},
	{"", "01234567", -1},
	{"32145678", "01234567", -1},
	{"01234567", "01234567", 0},
	{"x01234567", "01234567", 1},
	{"x0123456x01234567", "01234567", 9},
	{"xx01234567"[:9], "01234567", -1},
	{"", "0123456789", -1},
	{"3214567844", "0123456789", -1},
	{"0123456789", "0123456789", 0},
	{"x0123456789", "0123456789", 1},
	{"x012345678x0123456789", "0123456789", 11},
	{"xyz0123456789"[:12], "0123456789", -1},
	{"x01234567x89", "0123456789", -1},
	{"", "0123456789012345", -1},
	{"3214567889012345", "0123456789012345", -1},
	{"0123456789012345", "0123456789012345", 0},
	{"x0123456789012345", "0123456789012345", 1},
	{"x012345678901234x0123456789012345", "0123456789012345", 17},
	{"", "01234567890123456789", -1},
	{"32145678890123456789", "01234567890123456789", -1},
	{"01234567890123456789", "01234567890123456789", 0},
	{"x01234567890123456789", "01234567890123456789", 1},
	{"x0123456789012345678x01234567890123456789", "01234567890123456789", 21},
	{"xyz01234567890123456789"[:22], "01234567890123456789", -1},
	{"", "0123456789012345678901234567890", -1},
	{"321456788901234567890123456789012345678911", "0123456789012345678901234567890", -1},
	{"0123456789012345678901234567890", "0123456789012345678901234567890", 0},
	{"x0123456789012345678901234567890", "0123456789012345678901234567890", 1},
	{"x012345678901234567890123456789x0123456789012345678901234567890", "0123456789012345678901234567890", 32},
	{"xyz0123456789012345678901234567890"[:33], "0123456789012345678901234567890", -1},
	{"", "01234567890123456789012345678901", -1},
	{"32145678890123456789012345678901234567890211", "01234567890123456789012345678901", -1},
	{"01234567890123456789012345678901", "01234567890123456789012345678901", 0},
	{"x01234567890123456789012345678901", "01234567890123456789012345678901", 1},
	{"x0123456789012345678901234567890x01234567890123456789012345678901", "01234567890123456789012345678901", 33},
	{"xyz01234567890123456789012345678901"[:34], "01234567890123456789012345678901", -1},
	{"xxxxxx012345678901234567890123456789012345678901234567890123456789012", "012345678901234567890123456789012345678901234567890123456789012", 6},
	{"", "0123456789012345678901234567890123456789", -1},
	{"xx012345678901234567890123456789012345678901234567890123456789012", "0123456789012345678901234567890123456789", 2},
	{"xx012345678901234567890123456789012345678901234567890123456789012"[:41], "0123456789012345678901234567890123456789", -1},
	{"xx012345678901234567890123456789012345678901234567890123456789012", "0123456789012345678901234567890123456xxx", -1},
	{"xx0123456789012345678901234567890123456789012345678901234567890120123456789012345678901234567890123456xxx", "0123456789012345678901234567890123456xxx", 65},
	// test fallback to Rabin-Karp.
	{"oxoxoxoxoxoxoxoxoxoxoxoy", "oy", 22},
	{"oxoxoxoxoxoxoxoxoxoxoxox", "oy", -1},
	// 不支持中文
	// {"中文你好", "你好", 2},
}

func TestCommonSearch(t *testing.T) {
	for _, val := range testArr {
		if result := CommonSearch(val.str, val.pattern); val.result != result {
			t.Fatalf("Test common search failure, str: %s, pattern: %s, need %d but got %d", val.str, val.pattern, val.result, result)
		}
	}
}

func TestKMPSearch(t *testing.T) {
	for _, val := range testArr {
		if result := KMPSearch(val.str, val.pattern); val.result != result {
			t.Fatalf("Test KMP search failure, str: %s, pattern: %s, need %d but got %d", val.str, val.pattern, val.result, result)
		}
	}
}

func TestKMPSearch2(t *testing.T) {
	for _, val := range testArr {
		if result := KMPSearch2(val.str, val.pattern); val.result != result {
			t.Fatalf("Test KMP2 search failure, str: %s, pattern: %s, need %d but got %d", val.str, val.pattern, val.result, result)
		}
	}
}

func TestBMSearch(t *testing.T) {
	for _, val := range testArr {
		if result := BMSearch(val.str, val.pattern); val.result != result {
			t.Fatalf("Test BM search failure, str: %s, pattern: %s, need %d but got %d", val.str, val.pattern, val.result, result)
		}
	}
}

func TestBMSearch2(t *testing.T) {
	for _, val := range testArr {
		if result := BMSearch2(val.str, val.pattern); val.result != result {
			t.Fatalf("Test BM2 search failure, str: %s, pattern: %s, need %d but got %d", val.str, val.pattern, val.result, result)
		}
	}
}

func TestRKSearch2(t *testing.T) {
	for _, val := range testArr {
		if result := RabinKarpSearch(val.str, val.pattern); val.result != result {
			t.Fatalf("Test RabinKarp search failure, str: %s, pattern: %s, need %d but got %d", val.str, val.pattern, val.result, result)
		}
	}
}

func BenchmarkCommonSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// val := testArr[9]
		// CommonSearch(val.str, val.pattern)
		for _, val := range testArr {
			CommonSearch(val.str, val.pattern)
		}
	}
}

func BenchmarkKMPSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// val := testArr[9]
		// KMPSearch(val.str, val.pattern)
		for _, val := range testArr {
			KMPSearch(val.str, val.pattern)
		}
	}
}

func BenchmarkKMPSearch2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// val := testArr[9]
		// KMPSearch2(val.str, val.pattern)
		for _, val := range testArr {
			KMPSearch2(val.str, val.pattern)
		}
	}
}

func BenchmarkBMSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// val := testArr[9]
		// BMSearch(val.str, val.pattern)
		for _, val := range testArr {
			BMSearch(val.str, val.pattern)
		}
	}
}

func BenchmarkBMSearch2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// val := testArr[9]
		// BMSearch2(val.str, val.pattern)
		for _, val := range testArr {
			BMSearch2(val.str, val.pattern)
		}
	}
}

func BenchmarkRKSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// val := testArr[9]
		// RabinKarpSearch(val.str, val.pattern)
		for _, val := range testArr {
			RabinKarpSearch(val.str, val.pattern)
		}
	}
}