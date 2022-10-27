package register

import (
	"chi_sample/domain/model/user"
	ru "chi_sample/domain/repository/user"
)

type accountInteractor struct {
	userRepository ru.IUserRepository
}

// accountのinteractorを返却する
func NewAccountInteractor(ui ru.IUserRepository) accountInteractor {
	return accountInteractor{userRepository: ui}
}

func (ai accountInteractor) Interact(i InputDto) OutputDto {
	u, err := user.NewUser(i.Name, i.Mail, i.ImagePath)

	if err != nil {
		return OutputDto{
			IsRegistered: false,
			ErrMessage:   err.Error(),
		}
	}

	p, err := user.NewPassword(i.Password)

	if err != nil {
		return OutputDto{
			IsRegistered: false,
			ErrMessage:   err.Error(),
		}
	}

	err = ai.userRepository.Create(u.Id, u.Name, u.Mail, u.ImagePath, p.Value)

	if err != nil {
		return OutputDto{
			IsRegistered: false,
			ErrMessage:   err.Error(),
		}
	}

	return OutputDto{
		IsRegistered: true,
		ErrMessage:   "",
	}
}
