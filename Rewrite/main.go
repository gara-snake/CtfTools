package main

import (
	"flag"
	"os"
)

const ans_flag = "hot dog"

const file_name = "server_file"

func main() {

	name := flag.String("u", "", "あなたの名前")

	flag.Parse()

	file, e := os.Stat("./" + file_name)

	if e != nil {
		println("書き換え対象のファイルが見つかりません")
		return
	}

}
