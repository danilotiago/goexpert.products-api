package database

import (
	"testing"

	"github.com/backendengineerark/products-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&entity.User{})

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("john", "j@j.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&entity.User{})

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("john", "john3@john3.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	var count int64
	db.Model(&entity.User{}).Count(&count)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), count)

	println(count)

	userFound, err := userDB.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.Equal(t, user.ID.String(), userFound.ID.String())
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
