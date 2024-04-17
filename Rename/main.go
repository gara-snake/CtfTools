package main

import (
	"bufio"
	"flag"
	"log"
	"math/rand"
	"os"
	"strings"
)

const ans_flag = "T.M.Revolution"

func main() {

	file := flag.String("f", "", "ファイル名")

	flag.Parse()

	if *file != "" {
		f, e := os.Open(*file)

		if e != nil {
			println("ファイルオープンエラーです")
			return
		}

		defer f.Close()

		var str string

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			str = str + scanner.Text()
		}

		str = strings.ReplaceAll(str, "\n", "")
		str = strings.ReplaceAll(str, " ", "")

		fname := f.Name()

		if str == fname {
			println(ans_flag)

			os.Remove(*file)
		} else {
			println("ファイル名と内容が一致しません")
		}

		return

	} else {
		println("対象ファイルを作成します")

		name := makeRandomStr(10)

		f, err := os.Create(name)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		// 書き込むテキスト
		text := makeRandomStr(rand.Intn(10))

		// ファイルに書き込む
		_, err = f.WriteString(text + "\n")
		if err != nil {
			log.Fatal(err)
		}

		println(name + " を作成しました")

	}

}

// TODO 共通関数に直したい
func makeRandomStr(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
