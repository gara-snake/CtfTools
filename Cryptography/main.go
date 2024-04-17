package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const ans_flag = "Chikawa"

const test_txt = "fSLXH6PydAfK34qglX6KO4HPAIUDANpGw5yPYqMYjvL3rTIxnBVsAsc28GJETHSL0ysLNRNHG3LdudMOakSrdRjXx0BB3e3XfJJbUe2sC8vpAr28wJ2s6QjLDgAKHs6MhkKvdiGNNRjsd7QxsoHMOsUMrllyysQxXlXuEpOFwDVUSgsDYAxDXH8TrKQCKf3nu1cAcnEmEU5YCCdYTeMOw1odCueJyc8VgKI3M0UQtS1jpM4n5HwIf2Nvin3jjeMosk1Lfltv8xlNquyKk6NqiwnKO8mK0nynJwKOend0RMAk0AuT8VpLptkxJRhOm3yUfxc0my1XMdQVUl6BVi7wWxUucra4KmOFwwtI7DdU7jT18FvlwpGVjDERjbM6JIN3UatJw84k8gm50xyKoona0MIOqaMNsWJV03WKhMwk5Lp78vh8MSHhV08RbfRsQOrAf8c4QMs5Am5dqpbnLEnrAdgBsWUIQ7TPIYeLqd2CSG0dd1f8TGj1"
const flg_txt = "EgG3MPjYTk8GFfClKHLCea14CSkHJGvC8Hc4MBEB4FPAG5J0ybagBFi1EDYbQLQHdkwKEbxykNmjeXoshHQ53S4UWaUywUR8XxOfOEvUY07cXG4WNuA47pmmxsi7kqFPosVwIv2BcPF26hJdQwrbjmmyPw7XeaUaaF5sNFvvshK5ERJsXcD885jfwmrQ1Wj5NLSxsgq3EIIQX74nJs7EfKuAeuwSYd2tlGOUdNuSye52NJUbJmOlwoacs1VHls0UMit8uV81BLI6PO7wJuNf6sHVU4P3mc6IHI5oEfh3cmKFfWRO2ys6eKy4sAfPqsDI6erQDR8TIn0nhdKQ8KrkPvhaElLXCSYmb0F3LgW3CBjrjHj7chFtmXmRiWtGao5Iw4MRFFejAIafiOrhAKHsL43JcF1FcXjFR1ebHdGTlmqfXm2EAPI3K6fmyas8RAnv7kIOD2f3jaoFMHNmc3evyVc0dM8SGHfOUlxbdbulutNAB0JpNAms"

func main() {

	inputData, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}

	// 標準出力に書き込み
	fmt.Println("入力されたデータ:", string(inputData))

	str := string(inputData)
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, " ", "")

	if str == test_txt {
		println("正常にテスト用ファイルを解析しました。")
		return
	}

	if str == flg_txt {
		println(ans_flag)
		return
	}

	println(str)
	println("ファイルの内容が違います")

}
