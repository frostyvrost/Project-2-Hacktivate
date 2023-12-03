package service

import (
	"project-2/model"
	"project-2/pkg"
	"project-2/repo"

	"github.com/asaskevich/govalidator"
)

type commentServiceRepo interface {
	CreateComment(comment *model.Comment, userID uint) (*model.Comment, pkg.Error)
	GetComment(userID uint) ([]*model.Comment, pkg.Error)
	UpdateComment(*model.CommentUpdate, uint) (*model.Comment, pkg.Error)
	DeleteComment(uint) (*model.Comment, pkg.Error)
}

type commentService struct{}

var CommentService commentServiceRepo = &commentService{}

func (t *commentService) CreateComment(comment *model.Comment, userID uint) (*model.Comment, pkg.Error) {
	comment.UserID = userID

	if _, err := govalidator.ValidateStruct(comment); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	result, err := repo.CommentModel.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *commentService) GetComment(userID uint) ([]*model.Comment, pkg.Error) {

	comment, err := repo.CommentModel.GetComment(userID)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (t *commentService) UpdateComment(message *model.CommentUpdate, commentID uint) (*model.Comment, pkg.Error) {
	if _, err := govalidator.ValidateStruct(message); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	result, err := repo.CommentModel.UpdateComment(message, commentID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *commentService) DeleteComment(commentID uint) (*model.Comment, pkg.Error) {
	Result, err := repo.CommentModel.DeleteComment(commentID)

	if err != nil {
		return nil, err
	}

	return Result, nil
}
