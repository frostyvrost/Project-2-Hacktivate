package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"not null;unique" json:"email" valid:"required,email"`
	Username     string    `gorm:"not null;unique" json:"username" valid:"required"`
	Password     string    `gorm:"not null" json:"password" valid:"required,minstringlength(6)"`
	Age          int       `gorm:"not null" json:"age" valid:"required,range(8|150)"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Photos       []Photo
	Comments     []Comment
	SocialMedias []SocialMedia
}

type UserUpdate struct {
	Email    string `gorm:"not null;unique" json:"email" valid:"required,email"`
	Username string `gorm:"not null;unique" json:"username" valid:"required"`
}

type LoginCredential struct {
	Email    string `gorm:"not null;unique" json:"email" valid:"required,email"`
	Password string `gorm:"not null" json:"password" valid:"required,minstringlength(6)"`
}

type SocialMedia struct {
	ID             uint      `json:"id,omitempty" gorm:"primaryKey" `
	Name           string    `json:"name" gorm:"not null" valid:"required~Name is required"`
	SocialMediaURL string    `json:"social_media_url" gorm:"not null" valid:"required~Social media url is required"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           *User
}

type SocialMediaCreate struct {
	Name           string `json:"name" gorm:"not null" valid:"required~Name is required"`
	SocialMediaURL string `json:"social_media_url" gorm:"not null" valid:"required~Social media url is required"`
}

type SocialMediaUpdate struct {
	Name           string `json:"name" gorm:"not null" valid:"required~Name is required"`
	SocialMediaURL string `json:"social_media_url" gorm:"not null" valid:"required~Social media url is required"`
}

type Photo struct {
	ID        uint      `json:"id" gorm:"primaryKey" `
	Title     string    `json:"title" gorm:"not null" valid:"required"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url" gorm:"not null" valid:"required"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User
	Comments  []Comment
}

type PhotoUpdate struct {
	Title    string `json:"title" gorm:"not null" valid:"required"`
	Caption  string `json:"caption" valid:"required"`
	PhotoURL string `json:"photo_url" gorm:"not null" valid:"required"`
}

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey" `
	Message   string    `json:"message" gorm:"not null" valid:"required"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"photo_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User
	Photo     *Photo
}

type CommentUpdate struct {
	Message string `json:"message" gorm:"not null" valid:"required"`
}

func (comment *Comment) Validate() error {
	_, err := govalidator.ValidateStruct(comment)

	if err != nil {
		return err
	}

	return nil
}
