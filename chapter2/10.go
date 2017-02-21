package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var path string

func main() {
	flag.StringVar(&path, "f", "", "入力ファイル")
	flag.Parse()

	f, err := os.Open(path)
	if err != nil {
		fmt.Println("ファイルを開けません")
		return
	}

	// Scannerにより取得
	s := bufio.NewScanner(f)
	i := 0
	for s.Scan() {
		i += 1
	}
	fmt.Println("go      : ", i, "lines")

	// コマンドラインにより取得
	out, err := exec.Command("wc", "-l", path).Output()

	fmt.Println("command : ", strings.Fields(string(out))[0], "lines")
}
