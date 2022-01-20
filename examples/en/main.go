package main

import (
	"fmt"

	"github.com/go-ego/gse"
)

func main() {
	seg, err := gse.New("zh,../../testdata/test_dict3.txt", "alpha")
	fmt.Println("new gse error: ", err)
	// var seg gse.Segmenter
	// seg.AlphaNum = true
	// seg.LoadDict("zh, ../../testdata/test_dict3.txt")
	seg.AddToken("winter is coming!", 100, "n")

	freq, pos, ok := seg.Find("hello")
	fmt.Println(freq, pos, ok)

	freq, pos, ok = seg.Find("world")
	fmt.Println(freq, pos, ok)

	text := "Helloworld, winter is coming! 你好世界."

	tx := seg.Cut(text)
	fmt.Println(tx)

	tx = seg.Cut(text, true)
	fmt.Println(tx)

	tx = seg.Trim(tx)
	fmt.Println(tx)

	a := seg.Analyze(tx, text)
	fmt.Println(a)

	for i := 140010; i <= 140020; i++ {
		fmt.Printf("The token text is %v\n The token seg is %v\n",
			seg.Dict.Tokens[i].Text(), seg.Dict.Tokens[i].Segments())
		for _, v := range seg.Dict.Tokens[i].Segments() {
			fmt.Printf("the seg is %v the related text is %v\n", *v, (*v).Token().Text())
		}
		fmt.Println()
	}

}
