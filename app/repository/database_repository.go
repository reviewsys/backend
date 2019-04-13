package repository

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/reviewsys/backend/app/models"
)

type databaseUserRepository struct {
	Conn *gorm.DB
}

func NewDatabaseUserRepository(Conn *gorm.DB) UserRepository {

	return &databaseUserRepository{Conn}
}

func (r *databaseUserRepository) GetByID(id int64) (*models.User, error) {
	user := models.User{}
	err := r.Conn.First(&user, id).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &user, nil
}

func (r *databaseUserRepository) GetByTeamID(teamID int64) (*models.User, error) {
	user := models.User{}
	err := r.Conn.First(&user, teamID).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &user, nil
}

func (r *databaseUserRepository) Store(user *models.User) error {
	err := r.Conn.Create(&user).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *databaseUserRepository) Delete(user *models.User) error {
	err := r.Conn.Delete(&user).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *databaseUserRepository) Update(user *models.User) error {
	err := r.Conn.Save(&user).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
