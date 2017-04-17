package model

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type UserClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type User struct {
	ID        string `gorm:"primary_key" sql:"type:uuid;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string
	Password  string
	Email     string
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})

	return db
}
