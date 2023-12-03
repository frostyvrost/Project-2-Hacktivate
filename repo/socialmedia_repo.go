package repo

import (
	"project-2/database"
	"project-2/model"
	"project-2/pkg"
)

type socialMediaModelRepo interface {
	CreateSocialMedia(*model.SocialMedia) (*model.SocialMedia, pkg.Error)
	GetAllSocialMedias(uint) ([]*model.SocialMedia, pkg.Error)
	UpdateSocialMedia(*model.SocialMediaUpdate, uint) (*model.SocialMedia, pkg.Error)
	DeleteSocialMedia(uint) pkg.Error
}

type socialMediaRepo struct{}

var SocialMediaRepo socialMediaModelRepo = &socialMediaRepo{}

func (s *socialMediaRepo) CreateSocialMedia(socialMedia *model.SocialMedia) (*model.SocialMedia, pkg.Error) {
	db := database.GetDB()

	err := db.Create(&socialMedia).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return socialMedia, nil
}

func (s *socialMediaRepo) GetAllSocialMedias(userId uint) ([]*model.SocialMedia, pkg.Error) {
	db := database.GetDB()
	var socialMedia []*model.SocialMedia

	err := db.Preload("User").Where("user_id", userId).Find(&socialMedia).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	if len(socialMedia) == 0 {
		return nil, pkg.NotFound("Social media data is still empty. Please add data to continue.")
	}

	return socialMedia, nil
}

func (s *socialMediaRepo) UpdateSocialMedia(socialMediaUpdated *model.SocialMediaUpdate, socialMediaId uint) (*model.SocialMedia, pkg.Error) {
	db := database.GetDB()
	var socialMedia model.SocialMedia

	err := db.First(&socialMedia, socialMediaId).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	db.Model(&socialMedia).Updates(socialMediaUpdated)

	return &socialMedia, nil
}

func (s *socialMediaRepo) DeleteSocialMedia(socialMediaId uint) pkg.Error {
	db := database.GetDB()
	var socialMedia model.SocialMedia

	err := db.Where("id = ?", socialMediaId).Delete(&socialMedia).Error

	if err != nil {
		return pkg.ParseError(err)
	}

	return nil
}
