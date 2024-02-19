package main

import (
	"flag"
	"os"
	"path/filepath"
)

const ans_flag = "WWtgHUWDX"

const error_msg = "正しくディレクトリができていないようです"

func main() {

	name := flag.String("u", "", "あなたの名前")
	path := filepath.Join("/home/fcs/", *name)

	p, e := os.Stat(path)

	if e != nil {
		println(e.Error())
		return
	}

	if p.IsDir() {
		println(ans_flag)
	}

	println(error_msg)
}
