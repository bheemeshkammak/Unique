package services

import (
	"github.com/bheemeshkammak/Unique/unique/pkg/rest/server/daos"
	"github.com/bheemeshkammak/Unique/unique/pkg/rest/server/models"
)

type PremService struct {
	premDao *daos.PremDao
}

func NewPremService() (*PremService, error) {
	premDao, err := daos.NewPremDao()
	if err != nil {
		return nil, err
	}
	return &PremService{
		premDao: premDao,
	}, nil
}

func (premService *PremService) CreatePrem(prem *models.Prem) (*models.Prem, error) {
	return premService.premDao.CreatePrem(prem)
}

func (premService *PremService) UpdatePrem(id int64, prem *models.Prem) (*models.Prem, error) {
	return premService.premDao.UpdatePrem(id, prem)
}

func (premService *PremService) DeletePrem(id int64) error {
	return premService.premDao.DeletePrem(id)
}

func (premService *PremService) ListPrems() ([]*models.Prem, error) {
	return premService.premDao.ListPrems()
}

func (premService *PremService) GetPrem(id int64) (*models.Prem, error) {
	return premService.premDao.GetPrem(id)
}
