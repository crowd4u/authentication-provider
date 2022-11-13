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
	envs_as_string := os.Getenv(statement)
	if len(envs_as_string) < 1 {
		//エラー処理
		return "No Value", fmt.Errorf("error, there is no value in env")
	}
	return envs_as_string, nil
}
