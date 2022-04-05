package usecase

import "github.com/vitor-paholsky/CodePix-go/domain/model"

type PixUseCase struct {
	PixkeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) Registerkey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := p.PixkeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err

	}

	p.PixkeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, err
	}
	return pixKey, err
}
