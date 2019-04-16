package database

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/reviewsys/backend/app/domain/model"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) *userRepository {

	return &userRepository{Conn: Conn}
}

func (r *userRepository) GetByID(id *resource.Identifier) (*model.User, error) {
	user := model.User{}
	err := r.Conn.First(&user, id).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByTeamID(teamID int64) (*model.User, error) {
	user := model.User{}
	err := r.Conn.First(&user, teamID).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Store(user *model.User) error {
	err := r.Conn.Create(&user).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *userRepository) Delete(user *model.User) error {
	err := r.Conn.Delete(&user).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *userRepository) Update(user *model.User) error {
	err := r.Conn.Save(&user).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
