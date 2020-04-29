package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"wxServ/service/db"
)

type Token struct {
	gorm.Model

	OpenID string `gorm:"not null"`
	Token  string `gorm:"unique;not null"`
}

func Init() {
	db.Instance().AutoMigrate(&Token{})
}

func (t *Token) BeforeSave() error {
	if t.Token == "" {
		return errors.New("token is empty")
	}
	if t.OpenID == "" {
		return errors.New("openID is empty")
	}
	return nil
}

func (t *Token) Load() error {
	if t.Token != "" {
		return db.Instance().Where("token = ?", t.Token).Find(t).Error
	}
	if t.OpenID != "" {
		return db.Instance().Where("open_id = ?", t.OpenID).Find(t).Error
	}
	return errors.New("both token and openID are empty")
}

func (t *Token) Update() error {
	err := t.Load()
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return db.Instance().Save(t).Error
}

func GetOpenID(token string) (openID string, err error) {
	tk := Token{Token: token}
	if err := tk.Load(); err != nil {
		return "", err
	}
	return tk.OpenID, nil
}
