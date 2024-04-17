package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {

	rand.New(rand.NewSource(time.Now().UnixNano()))

	baseDir := "random_structure"
	depth := 5
	numDirs := 5

	os.RemoveAll(baseDir)
	os.Mkdir(baseDir, os.ModePerm)

	createDirsRecursively(baseDir, depth, numDirs)
}

func createDirsRecursively(parentDir string, depth, numDirs int) {
	if depth == 0 {
		return
	}

	for i := 0; i < numDirs; i++ {
		dirName := fmt.Sprintf("dir%d", i)
		dirPath := filepath.Join(parentDir, dirName)
		os.Mkdir(dirPath, os.ModePerm)

		fileNum := rand.Intn(6)

		for i := 0; i < fileNum; i++ {

			fPath := filepath.Join(dirPath, "code_"+fmt.Sprint(i))

			f, err := os.Create(fPath)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			// 書き込むテキスト
			text := makeRandomStr(rand.Intn(50))

			// ファイルに書き込む
			_, err = f.WriteString(text + "\n")
			if err != nil {
				log.Fatal(err)
			}

		}

		createDirsRecursively(dirPath, depth-1, rand.Intn(5))
	}
}

func makeRandomStr(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
