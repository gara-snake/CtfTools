package main

import (
	"flag"
	"os"
	"strings"
)

const ans_flag = "hot dog"

const file_name = "server_file"

func main() {

	name := flag.String("u", "", "あなたの名前")

	flag.Parse()

	file, e := os.Open("./" + file_name)

	if e != nil {
		println("書き換え対象のファイルが見つかりません")
		return
	}

	defer file.Close()

	data := make([]byte, 1024)
	count, err := file.Read(data)
	if err != nil {
		println("ファイルの読み取りに失敗しました")
		return
	}

	ret := string(data[:count])
	ret = strings.ReplaceAll(ret, "\n", "")

	if ret == *name {
		println(ans_flag)
		return
	}

	println(ret)
	println("ファイルの内容がオプションと一致しません")

}
