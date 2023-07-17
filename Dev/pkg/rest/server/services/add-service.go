package services

import (
	"github.com/bheemeshkammak/Unique/dev/pkg/rest/server/daos"
	"github.com/bheemeshkammak/Unique/dev/pkg/rest/server/models"
)

type AddService struct {
	addDao *daos.AddDao
}

func NewAddService() (*AddService, error) {
	addDao, err := daos.NewAddDao()
	if err != nil {
		return nil, err
	}
	return &AddService{
		addDao: addDao,
	}, nil
}

func (addService *AddService) CreateAdd(add *models.Add) (*models.Add, error) {
	return addService.addDao.CreateAdd(add)
}

func (addService *AddService) UpdateAdd(id int64, add *models.Add) (*models.Add, error) {
	return addService.addDao.UpdateAdd(id, add)
}

func (addService *AddService) DeleteAdd(id int64) error {
	return addService.addDao.DeleteAdd(id)
}

func (addService *AddService) ListAdds() ([]*models.Add, error) {
	return addService.addDao.ListAdds()
}

func (addService *AddService) GetAdd(id int64) (*models.Add, error) {
	return addService.addDao.GetAdd(id)
}
