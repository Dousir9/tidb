package parser

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/yanyiwu/gojieba"
)

var (
	Jieba       = gojieba.NewJieba()
	stopWordSet = loadStopWords()
)

func Cutl(sentence string) []string {
	return removeStopwords(Jieba.Cut(sentence, true))
}

func CutForSearch(sentence string) []string {
	return removeStopwords(Jieba.CutForSearch(sentence, true))
}

func removeStopwords(input []string) []string {
	cut_res := make([]string, 2)
	for _, word := range input {
		_, exists := stopWordSet[word]
		if !exists {
			cut_res = append(cut_res, word)
		}
	}
	return cut_res
}

func loadStopWords() map[string]interface{} {
	set := make(map[string]interface{})

	f, err := os.Open("/home/dousir9/TiDB-Hackathon/tidb_hackathon_2022/parser/stops.txt")
	if err != nil {
		panic("can't open file stops.txt")
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		panic("read stopwords file error")
	}

	for _, word := range strings.Split(string(fd), "\n") {
		set[word] = struct{}{}
	}
	return set
}
