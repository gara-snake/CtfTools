package main

import (
	"os"
)

const ans_flag = "Dwayne Johnson"

const file_name = "secret_file"

func main() {

	file, e := os.Open("./" + file_name)

	if e != nil {
		println(ans_flag)

		// 復旧
		f2, e2 := os.Create("./" + file_name)
		if e2 != nil {
			println(e2.Error())
			return
		}

		defer f2.Close()
		f2.WriteString("秘密\n")

		return
	}

	defer file.Close()

	println("ファイルがまだあります")

}
