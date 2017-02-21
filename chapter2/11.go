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
	fmt.Println("======== Go ========")
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(strings.Replace(s.Text(), "\t", " ", -1))
	}

	// コマンドにより取得
	fmt.Println("======== UNIX Command ========")
	out, err := exec.Command("expand", "-t", "1", path).Output()
	fmt.Println(string(out))
}
