package algorithm

import (
	"fmt"
	"strings"
)

/**
Jaccard相似度算法:用于计算两个集合之间的相似度，可用于比较文本中出现的不同单词和词语，计算方式为: 两个集合的交集大小除以他们的并集大小
	优点:
		简单容易实现，适用于短文本
	缺点:
		无法考虑单词/词语的重要性和顺序
*/

// 计算两个字符串的Jaccard相似度
func jaccard(a, b string) float64 {
	setA := make(map[string]bool)
	setB := make(map[string]bool)
	for _, word := range strings.Fields(strings.ToLower(a)) {
		setA[word] = true
	}
	for _, word := range strings.Fields(strings.ToLower(b)) {
		setB[word] = true
	}
	var intersection float64
	for word := range setA {
		if setB[word] {
			intersection++
		}
	}
	union := float64(len(setA)+len(setB)) - intersection
	return intersection / union
}

func TestJaccard() {
	a := "The quick brown fox jumps over the lazy dog"
	b := "THE"
	similarity := jaccard(a, b)
	fmt.Println(similarity)
}
