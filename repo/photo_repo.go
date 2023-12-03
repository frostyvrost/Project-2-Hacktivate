package repo

import (
	"project-2/database"
	"project-2/model"
	"project-2/pkg"
)

type photoModelRepo interface {
	CreatePhoto(*model.Photo) (*model.Photo, pkg.Error)
	UpdatePhoto(*model.PhotoUpdate, uint) (*model.Photo, pkg.Error)
	GetAllPhotos(uint) ([]*model.Photo, pkg.Error)
	DeletePhoto(uint) pkg.Error
}

type photoModel struct{}

var PhotoModel photoModelRepo = &photoModel{}

func (t *photoModel) CreatePhoto(photo *model.Photo) (*model.Photo, pkg.Error) {
	db := database.GetDB()

	err := db.Create(&photo).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return photo, nil
}

func (t *photoModel) UpdatePhoto(input *model.PhotoUpdate, photoID uint) (*model.Photo, pkg.Error) {
	db := database.GetDB()

	var photo model.Photo
	err := db.First(&photo, photoID).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	db.Model(&photo).Updates(input)

	return &photo, nil
}

func (p *photoModel) GetAllPhotos(userId uint) ([]*model.Photo, pkg.Error) {
	db := database.GetDB()
	var photos []*model.Photo

	err := db.Preload("User").Where("user_id = ?", userId).Find(&photos).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return photos, nil
}

func (p *photoModel) DeletePhoto(photoId uint) pkg.Error {
	db := database.GetDB()

	var photo model.Photo

	err := db.Where("id = ?", photoId).Delete(&photo).Error

	if err != nil {
		return pkg.ParseError(err)
	}

	return nil
}
