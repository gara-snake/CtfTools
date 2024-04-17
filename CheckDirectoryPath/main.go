package main

import (
	"flag"
)

const ans_flag = "VLRAnGNBLnj"

func main() {

	path := flag.String("p", "", "対象のディレクトリパス")

	flag.Parse()

	if *path == "/home/fcs/LEVEL1/cat_file/dir4/sub4_2" {
		println(ans_flag)
		return
	}

	println("パラメータの指定が違います")

}
