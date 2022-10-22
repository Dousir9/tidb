package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/yanyiwu/gojieba"
)

var (
	Jieba        = gojieba.NewJieba()
	stopWordSet  = loadStopWords()
	forwardDict  = loadModel("tidb_hackathon_forwards.json")
	backwardDict = loadModel("tidb_hackathon_backwards.json")
	thredhold    = 0.8
	basePath     = "/home/dousir9/TiDB-Hackathon/tidb_hackathon_2022/parser/"
)

func Cutl(sentence string) []string {
	return removeStopwords(Jieba.Cut(sentence, true))
}

func CutForSearch(sentence string) []string {
	return removeStopwords(Jieba.CutForSearch(sentence, true))
}

func removeStopwords(input []string) []string {
	cut_res := make([]string, 0)
	for _, word := range input {
		_, exists := stopWordSet[word]
		if !exists && word != "" && word != " " {
			cut_res = append(cut_res, word)
		}
	}
	return cut_res
}

func loadStopWords() map[string]interface{} {
	set := make(map[string]interface{})

	f, err := os.Open(basePath + "stops.txt")
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

func loadModel(path string) map[string]map[string]float64 {
	var model map[string]map[string]float64
	data, err := ioutil.ReadFile(basePath + path)
	if err != nil {
		fmt.Println(err)
		panic("can't load model file")
	}

	err = json.Unmarshal(data, &model)
	return model
}

func SimWord(word string) []string {
	candidate_set := make([]string, 0)
	// forward
	wordForward, exists := forwardDict[word]
	if exists {
		fwdWord, fwdScore := getSimWord(wordForward)
		if fwdBckDict, exists := backwardDict[fwdWord]; exists {
			fwdBckWord, fwdBckScore := getSimWord(fwdBckDict)
			if fwdBckScore*fwdScore >= 0.8 {
				candidate_set = append(candidate_set, fwdBckWord)
			}
		}
	}
	// backward
	wordBack, exists := backwardDict[word]
	if exists {
		bckWord, backScore := getSimWord(wordBack)
		if bckFwdDict, exists := forwardDict[bckWord]; exists {
			bckFwdWord, bckFwdScore := getSimWord(bckFwdDict)
			if bckFwdScore*backScore > 0. {
				candidate_set = append(candidate_set, bckFwdWord)
			}
		}
	}
	return candidate_set
}

func getSimWord(transDict map[string]float64) (simWord string, simScore float64) {
	for word, score := range transDict {
		if score > simScore {
			simScore = score
			simWord = word
		}
	}
	return
}
