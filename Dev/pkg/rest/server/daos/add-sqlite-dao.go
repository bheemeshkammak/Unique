package daos

import (
	"database/sql"
	"errors"
	"github.com/bheemeshkammak/Unique/dev/pkg/rest/server/daos/clients/sqls"
	"github.com/bheemeshkammak/Unique/dev/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type AddDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateAdds(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS adds(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Mul TEXT NOT NULL,
		Sub TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewAddDao() (*AddDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateAdds(sqlClient)
	if err != nil {
		return nil, err
	}
	return &AddDao{
		sqlClient,
	}, nil
}

func (addDao *AddDao) CreateAdd(m *models.Add) (*models.Add, error) {
	insertQuery := "INSERT INTO adds(Mul, Sub)values(?, ?)"
	res, err := addDao.sqlClient.DB.Exec(insertQuery, m.Mul, m.Sub)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("add created")
	return m, nil
}

func (addDao *AddDao) UpdateAdd(id int64, m *models.Add) (*models.Add, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	add, err := addDao.GetAdd(id)
	if err != nil {
		return nil, err
	}
	if add == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE adds SET Mul = ?, Sub = ? WHERE Id = ?"
	res, err := addDao.sqlClient.DB.Exec(updateQuery, m.Mul, m.Sub, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("add updated")
	return m, nil
}

func (addDao *AddDao) DeleteAdd(id int64) error {
	deleteQuery := "DELETE FROM adds WHERE Id = ?"
	res, err := addDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("add deleted")
	return nil
}

func (addDao *AddDao) ListAdds() ([]*models.Add, error) {
	selectQuery := "SELECT * FROM adds"
	rows, err := addDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var adds []*models.Add
	for rows.Next() {
		m := models.Add{}
		if err = rows.Scan(&m.Id, &m.Mul, &m.Sub); err != nil {
			return nil, err
		}
		adds = append(adds, &m)
	}
	if adds == nil {
		adds = []*models.Add{}
	}

	log.Debugf("add listed")
	return adds, nil
}

func (addDao *AddDao) GetAdd(id int64) (*models.Add, error) {
	selectQuery := "SELECT * FROM adds WHERE Id = ?"
	row := addDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Add{}
	if err := row.Scan(&m.Id, &m.Mul, &m.Sub); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("add retrieved")
	return &m, nil
}
