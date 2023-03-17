package algorithm

import (
	"fmt"
	"math"
	"strings"
)

// 用于存储文档中每个单词出现的次数
type wordCount struct {
	word  string
	count int
}

// 计算一个文档中每个单词出现的次数
func getWordCounts(text string) []wordCount {
	words := strings.Fields(text)
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	var result []wordCount
	for word, count := range counts {
		result = append(result, wordCount{word, count})
	}
	return result
}

// 计算一个单词在所有文档中出现的次数
func getDocCount(word string, docs []string) int {
	count := 0
	for _, doc := range docs {
		if strings.Contains(doc, word) {
			count++
		}
	}
	return count
}

// 计算一个文档中每个单词的TF值
func getTf(text string) map[string]float64 {
	tf := make(map[string]float64)
	wordCounts := getWordCounts(text)
	totalWords := len(strings.Fields(text))
	for _, wc := range wordCounts {
		tf[wc.word] = float64(wc.count) / float64(totalWords)
	}
	return tf
}

// 计算一个文档中每个单词的IDF值
func getIdf(word string, docs []string) float64 {
	n := float64(len(docs))
	df := float64(getDocCount(word, docs))
	return math.Log(n / (1.0 + df))
}

// 计算一个文档中每个单词的TF-IDF值，并返回一个map
func getTfIdf(text string, docs []string) map[string]float64 {
	tfidf := make(map[string]float64)
	tf := getTf(text)
	for word, tfValue := range tf {
		idf := getIdf(word, docs)
		tfidf[word] = tfValue * idf
	}
	return tfidf
}

func TestTF_IDF() {
	docs := []string{
		`2023-03-17T15:43:06.849+0800 ERROR Currently new worker [934_934] failed, Err: [try lock: lock: context deadline exceeded] {"uuid": 1678763951, "funcName": "discovery.(*EtcdAllocator).checkAllocate", "caller": "/var/go/pkg/mod/git.hyperchain.cn/blocface/gateway@v1.4.2/pkg/discovery/allocator.go:191"}`,
	}
	text := "2023-03-17T15 ERROR new worker"
	tfidf := getTfIdf(text, docs)
	fmt.Println(tfidf)
}
