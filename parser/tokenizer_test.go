package parser

import (
	"fmt"
	"testing"
)

func TestTokenizerInit(t *testing.T) {
	fmt.Println("initing")
	Jieba.Cut("this is a test case", true)
}

func TestSimWord(t *testing.T) {
	fmt.Println(SimWord("快乐"))
}
