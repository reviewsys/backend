package database

import (
	"context"

	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/jinzhu/gorm"
	pb "github.com/reviewsys/backend/app/interface/rpc/v1.0/protocol"
	log "github.com/sirupsen/logrus"

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

func (r *userRepository) Store(ctx context.Context, user *pb.User) (*pb.User, error) {
	user, err := pb.DefaultCreateUser(ctx, user, r.Conn)
	if err != nil {
		log.Errorf("database error: %v", err)
	}
	return user, err
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
