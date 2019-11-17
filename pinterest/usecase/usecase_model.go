package usecase

import (
	"github.com/ValeryBMSTU/DB_TP/pinterest/repository"
	"sync"
)


type UseStruct struct {
	PRepository repository.ReposInterface
	Mu          *sync.Mutex
}

func (USC *UseStruct) NewUseCase(mu *sync.Mutex, rep repository.ReposInterface) error {
	USC.Mu = mu
	USC.PRepository = rep
	return nil
}

type UseInterface interface {
	//SetJSONData(data interface{}, token string, infMsg string) models.OutJSON

}
