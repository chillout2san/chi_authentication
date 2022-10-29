package user

import (
	"chi_sample/common/utils"
	"errors"
)

type Password struct {
	Value string
}

// 新しいパスワードを作成するファクトリ関数
// セキュリティ的にuserモデルと分離してある
func NewPassword(pass string) (Password, error) {
	if pass == "" {
		return Password{}, errors.New("パスワードが空です。")
	}

	hashedPass := utils.CreateHash(pass)

	return Password{Value: hashedPass}, nil
}
