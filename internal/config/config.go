package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var GlobalConfig *Config

type Store map[string]string

type Config struct {
	Store
}

// NOTE: Config周りが適当な実装なので、ライブラリ等入れていい感じにする。

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(".envファイルが読み込めませんでした。")
	}

	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		fmt.Println("警告: POSTGRES_USERが設定されていません。デフォルト値を使用します。")
		user = "postgres"
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		fmt.Println("警告: POSTGRES_PASSWORDが設定されていません。デフォルト値を使用します。")
		password = "postgres"
	}

	dbName := os.Getenv("POSTGRES_DB")
	if dbName == "" {
		fmt.Println("警告: POSTGRES_DBが設定されていません。デフォルト値を使用します。")
		dbName = "postgres"
	}

	GlobalConfig = &Config{
		Store: Store{
			"database.user":     user,
			"database.password": password,
			"database.dbname":   dbName,
		},
	}
}
