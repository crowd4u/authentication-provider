package infra

import (
	"fmt"
	"notchman8600/authentication-provider/interfaces/env"
	"os"
)

type EnvHandler struct {
}

func NewEnvHandler() env.EnvHandler {
	envHandler := new(EnvHandler)
	return envHandler
}

func (handler *EnvHandler) ReadEnv(statement string) (string, error) {
	// TODO:これは削除すること
	// 演習室環境用にSQL_URLをハードコーディングの上でコンパイル時に叩き込む
	const SQL_URL = "s2113591:tsukuba@localhost:3306/s2113591?charset=utf8"
	if statement == "SQL_URL" {
		return SQL_URL, nil
	}

	envs_as_string := os.Getenv(statement)
	if len(envs_as_string) < 1 {
		//エラー処理
		return "No Value", fmt.Errorf("error, there is no value in env")
	}
	return envs_as_string, nil
}
