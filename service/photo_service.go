package service

import (
	"project-2/model"
	"project-2/pkg"
	"project-2/repo"

	"github.com/asaskevich/govalidator"
)

type photoServiceRepo interface {
	CreatePhoto(*model.Photo, uint) (*model.Photo, pkg.Error)
	UpdatePhoto(*model.PhotoUpdate, uint) (*model.Photo, pkg.Error)
	GetAllPhotos(uint) ([]*model.Photo, pkg.Error)
	DeletePhoto(uint) pkg.Error
}

type photoService struct{}

var PhotoService photoServiceRepo = &photoService{}

func (t *photoService) CreatePhoto(photo *model.Photo, userID uint) (*model.Photo, pkg.Error) {
	photo.UserID = userID

	if _, err := govalidator.ValidateStruct(photo); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	result, err := repo.PhotoModel.CreatePhoto(photo)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *photoService) UpdatePhoto(photo *model.PhotoUpdate, photoID uint) (*model.Photo, pkg.Error) {
	if _, err := govalidator.ValidateStruct(photo); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	result, err := repo.PhotoModel.UpdatePhoto(photo, photoID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *photoService) GetAllPhotos(userId uint) ([]*model.Photo, pkg.Error) {
	photos, err := repo.PhotoModel.GetAllPhotos(userId)

	if err != nil {
		return nil, err
	}

	return photos, nil
}

func (p *photoService) DeletePhoto(photoId uint) pkg.Error {
	err := repo.PhotoModel.DeletePhoto(photoId)

	if err != nil {
		return err
	}

	return nil
}
