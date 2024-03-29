package config

import "os"

// 環境変数をパースした構造体
type env struct {
	AllowOrigin   string // corsで許可するオリジン
	DbUser        string // GCPのCloudSQLにアクセスするユーザー名
	DbPass        string // 該当ユーザーのパスワード
	DbHost        string // GCPのCloudSQLの接続名
	DbName        string // データベースの名前
	SecretKey     string // jwtの署名に使用するキー
	DbFlag        string // CloudSQLを使うかDockerのMySQLを使うかのフラグ
	COOKIE_SECURE bool   // cookieのsecure属性を有効にするかどうか
}

var Environment *env

func init() {
	cs := true
	if os.Getenv("COOKIE_SECURE") == "false" {
		cs = false
	}
	Environment = &env{
		AllowOrigin:   os.Getenv("ALLOW_ORIGIN"),
		DbUser:        os.Getenv("DB_USER"),
		DbPass:        os.Getenv("DB_PASS"),
		DbHost:        os.Getenv("INSTANCE_CONNECTION_NAME"),
		DbName:        os.Getenv("DB_NAME"),
		SecretKey:     os.Getenv("JWT_SECRET_KEY"),
		DbFlag:        os.Getenv("DB_FLAG"),
		COOKIE_SECURE: cs,
	}
}
