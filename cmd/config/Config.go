package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	StorageList []string
	Port        string
}

func New() *Config {
	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// 環境変数MY_ARRAYを取得
	myArrayStr := os.Getenv("STORAGE_LIST")

	// 配列文字列をパース
	myArray := strings.Split(strings.Trim(myArrayStr, "[]"), ",")

	// 配列を整形（不要なスペースを削除）
	for i, val := range myArray {
		myArray[i] = strings.TrimSpace(val)
	}

	return &Config{
		StorageList: myArray,
		Port:        ":" + os.Getenv("APP_PORT"),
	}

}
