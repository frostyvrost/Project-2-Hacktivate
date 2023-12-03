package repo

import (
	"project-2/database"
	"project-2/model"
	"project-2/pkg"
)

type commentModelRepo interface {
	CreateComment(comment *model.Comment) (*model.Comment, pkg.Error)
	GetComment(userID uint) ([]*model.Comment, pkg.Error)
	UpdateComment(*model.CommentUpdate, uint) (*model.Comment, pkg.Error)
	DeleteComment(uint) (*model.Comment, pkg.Error)
}

type commentModel struct{}

var CommentModel commentModelRepo = &commentModel{}

func (t *commentModel) CreateComment(comment *model.Comment) (*model.Comment, pkg.Error) {
	db := database.GetDB()

	err := db.Create(&comment).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return comment, nil
}

func (t *commentModel) GetComment(userID uint) ([]*model.Comment, pkg.Error) {
	db := database.GetDB()

	var comment []*model.Comment

	err := db.Preload("User").Preload("Photo").Where("user_id = ?", userID).Find(&comment).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return comment, nil
}

func (t *commentModel) UpdateComment(message *model.CommentUpdate, commentID uint) (*model.Comment, pkg.Error) {
	db := database.GetDB()

	var comment model.Comment
	err := db.First(&comment, commentID).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	db.Model(&comment).Updates(message)

	return &comment, nil
}

func (p *commentModel) DeleteComment(commentID uint) (*model.Comment, pkg.Error) {
	db := database.GetDB()

	var comment model.Comment

	err := db.Where("id = ?", commentID).Delete(&comment).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return &comment, nil
}
