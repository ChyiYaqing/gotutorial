package algorithm

import (
	"fmt"
	"math"
	"strings"
)

/**
预先相似度:
	用于计算两个向量之间的相似度，可以用于比较文本中出现的单词(或其他数值类型特征)向量，计算方式为: 两个向量的内积除以他们的模长乘积。

	优点:
		能够考虑单词的重要性和顺序，适用长文本
	缺点:
		无法处理文本中的语义信息
*/

// 计算两个向量之间的余弦相似度
func cosineSimilarity(v1, v2 []float64) float64 {
	var sum1, sum2, dotProduct float64
	for i := 0; i < len(v1); i++ {
		sum1 += v1[i] * v1[i]
		sum2 += v2[i] * v2[i]
		dotProduct += v1[i] * v2[i]
	}
	norm1 := math.Sqrt(sum1)
	norm2 := math.Sqrt(sum2)
	return dotProduct / (norm1 * norm2)
}

// 将文本转换为词向量
func textToVector(text string) []float64 {
	var vector []float64
	wordCount := make(map[string]int)
	totalWords := 0
	for _, word := range strings.Fields(text) {
		wordCount[word]++
		totalWords++
	}

	for _, count := range wordCount {
		tf := float64(count) / float64(totalWords)
		idf := math.Log(float64(totalWords) / float64(count))
		vector = append(vector, tf*idf)
	}
	return vector
}

func TestCosine() {
	text1 := `2023-03-17T15:43:06.849+0800 ERROR Currently new worker [934_934] failed, Err: [try lock: lock: context deadline exceeded] {"uuid": 1678763951, "funcName": "discovery.(*EtcdAllocator).checkAllocate", "caller": "/var/go/pkg/mod/git.hyperchain.cn/blocface/gateway@v1.4.2/pkg/discovery/allocator.go:191"}`
	text2 := "2023-03-17"
	vector1 := textToVector(text1)
	vector2 := textToVector(text2)
	similarity := cosineSimilarity(vector1, vector2)
	fmt.Println(similarity)
}
