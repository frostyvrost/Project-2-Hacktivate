package service

import (
	"project-2/model"
	"project-2/pkg"
	"project-2/repo"

	"github.com/asaskevich/govalidator"
)

type socialMediaServiceRepo interface {
	CreateSocialMedia(*model.SocialMedia) (*model.SocialMedia, pkg.Error)
	GetAllSocialMedias(uint) ([]*model.SocialMedia, pkg.Error)
	UpdateSocialMedia(*model.SocialMediaUpdate, uint) (*model.SocialMedia, pkg.Error)
	DeleteSocialMedia(uint) pkg.Error
}

type socialMediaService struct{}

var SocialMediaService socialMediaServiceRepo = &socialMediaService{}

func (s *socialMediaService) CreateSocialMedia(socialMedia *model.SocialMedia) (*model.SocialMedia, pkg.Error) {
	if _, err := govalidator.ValidateStruct(socialMedia); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	socialMediaResponse, err := repo.SocialMediaRepo.CreateSocialMedia(socialMedia)

	if err != nil {
		return nil, err
	}

	return socialMediaResponse, nil
}

func (s *socialMediaService) GetAllSocialMedias(userId uint) ([]*model.SocialMedia, pkg.Error) {
	socialMedias, err := repo.SocialMediaRepo.GetAllSocialMedias(userId)

	if err != nil {
		return nil, err
	}

	return socialMedias, nil
}

func (s *socialMediaService) UpdateSocialMedia(socialMediaUpdated *model.SocialMediaUpdate, socialMediaId uint) (*model.SocialMedia, pkg.Error) {
	if _, err := govalidator.ValidateStruct(socialMediaUpdated); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	socialMedia, err := repo.SocialMediaRepo.UpdateSocialMedia(socialMediaUpdated, socialMediaId)

	if err != nil {
		return nil, err
	}

	return socialMedia, nil
}

func (s *socialMediaService) DeleteSocialMedia(socialMediaId uint) pkg.Error {
	err := repo.SocialMediaRepo.DeleteSocialMedia(socialMediaId)

	if err != nil {
		return err
	}

	return nil
}
