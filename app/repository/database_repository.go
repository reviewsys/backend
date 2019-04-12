package repository

import (
	"database/sql"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/reviewsys/backend/app/models"
)

type databaseUserRepository struct {
	Conn *sql.DB
}

func NewDatabaseUserRepository(Conn *sql.DB) UserRepository {

	return &databaseUserRepository{Conn}
}

func (m *databaseUserRepository) fetch(query string, args ...interface{}) ([]*models.User, error) {

	rows, err := m.Conn.Query(query, args...)

	if err != nil {
		log.Error(err)
		return nil, models.INTERNAL_SERVER_ERROR
	}
	defer rows.Close()
	result := make([]*models.User, 0)
	for rows.Next() {
		t := new(models.User)
		err = rows.Scan(
			&t.ID,
			&t.TeamID,
			&t.Name,
			&t.IsAdmin,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			log.Error(err)
			return nil, models.INTERNAL_SERVER_ERROR
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *databaseUserRepository) Fetch(cursor string, num int64) ([]*models.User, error) {

	query := `SELECT id,team_id,name,is_admin,updated_at, created_at
  						FROM user WHERE ID > ? LIMIT ?`

	return m.fetch(query, cursor, num)

}
func (m *databaseUserRepository) GetByID(id int64) (*models.User, error) {
	query := `SELECT id,team_id,name,is_admin,updated_at, created_at
  						FROM user WHERE ID = ?`

	list, err := m.fetch(query, id)
	if err != nil {

		return nil, err
	}

	a := &models.User{}
	if len(list) > 0 {
		a = list[0]
	} else {

		return nil, models.NOT_FOUND_ERROR
	}

	return a, nil
}

func (m *databaseUserRepository) GetByTeamID(team_id int64) (*models.User, error) {
	query := `SELECT id,team_id,name,is_admin,updated_at, created_at
  						FROM user WHERE title = ?`

	list, err := m.fetch(query, team_id)
	if err != nil {
		return nil, err
	}

	a := &models.User{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, models.NOT_FOUND_ERROR
	}
	return a, nil
}

func (m *databaseUserRepository) Store(u *models.User) (int64, error) {

	query := `INSERT user SET team_id=? , name=? , is_admin=? , updated_at=? , created_at=?`
	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		log.Error(err)
		return 0, models.INTERNAL_SERVER_ERROR
	}
	res, err := stmt.Exec(u.TeamID, u.Name, u.IsAdmin, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		log.Error(err)
		return 0, models.INTERNAL_SERVER_ERROR
	}
	return res.LastInsertId()
}

func (m *databaseUserRepository) Delete(id int64) (bool, error) {
	query := "DELETE FROM user WHERE id = ?"

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		log.Error(err)
		return false, models.INTERNAL_SERVER_ERROR
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Error(err)
		return false, models.INTERNAL_SERVER_ERROR
	}
	rowsAfected, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return false, models.INTERNAL_SERVER_ERROR
	}
	if rowsAfected <= 0 {
		return false, models.INTERNAL_SERVER_ERROR
	}

	return true, nil
}

func (m *databaseUserRepository) Update(u *models.User) (*models.User, error) {
	query := `UPDATE user set term_id=?, name=?, is_admin=? , updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	res, err := stmt.Exec(u.TeamID, u.Name, u.IsAdmin, u.UpdatedAt, u.ID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if affect < 1 {
		return nil, errors.New("Nothing Affected. Make sure your user is exist in DB")
	}

	return u, nil
}
