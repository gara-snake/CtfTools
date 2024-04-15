package main

import (
	"flag"
	"os"
	"path/filepath"
)

const ans_flag = "WWtgHUWDX"

const error_msg = "正しくディレクトリができていないか、パラメータが誤っています"

func main() {

	name := flag.String("u", "", "あなたの名前")

	flag.Parse()

	path := filepath.Join("/home/fcs/LEVEL1/make_directory/", *name)
	println(path)

	_, e := os.Stat(path)

	if e == nil {
		println(ans_flag)
		return
	}

	println(error_msg)

}
