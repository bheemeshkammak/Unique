package daos

import (
	"database/sql"
	"errors"
	"github.com/bheemeshkammak/Unique/unique/pkg/rest/server/daos/clients/sqls"
	"github.com/bheemeshkammak/Unique/unique/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type PremDao struct {
	sqlClient *sqls.SQLiteClient
}

func migratePrems(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS prems(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		INR TEXT NOT NULL,
		Trisha TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewPremDao() (*PremDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migratePrems(sqlClient)
	if err != nil {
		return nil, err
	}
	return &PremDao{
		sqlClient,
	}, nil
}

func (premDao *PremDao) CreatePrem(m *models.Prem) (*models.Prem, error) {
	insertQuery := "INSERT INTO prems(INR, Trisha)values(?, ?)"
	res, err := premDao.sqlClient.DB.Exec(insertQuery, m.INR, m.Trisha)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("prem created")
	return m, nil
}

func (premDao *PremDao) UpdatePrem(id int64, m *models.Prem) (*models.Prem, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	prem, err := premDao.GetPrem(id)
	if err != nil {
		return nil, err
	}
	if prem == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE prems SET INR = ?, Trisha = ? WHERE Id = ?"
	res, err := premDao.sqlClient.DB.Exec(updateQuery, m.INR, m.Trisha, id)
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

	log.Debugf("prem updated")
	return m, nil
}

func (premDao *PremDao) DeletePrem(id int64) error {
	deleteQuery := "DELETE FROM prems WHERE Id = ?"
	res, err := premDao.sqlClient.DB.Exec(deleteQuery, id)
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

	log.Debugf("prem deleted")
	return nil
}

func (premDao *PremDao) ListPrems() ([]*models.Prem, error) {
	selectQuery := "SELECT * FROM prems"
	rows, err := premDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var prems []*models.Prem
	for rows.Next() {
		m := models.Prem{}
		if err = rows.Scan(&m.Id, &m.INR, &m.Trisha); err != nil {
			return nil, err
		}
		prems = append(prems, &m)
	}
	if prems == nil {
		prems = []*models.Prem{}
	}

	log.Debugf("prem listed")
	return prems, nil
}

func (premDao *PremDao) GetPrem(id int64) (*models.Prem, error) {
	selectQuery := "SELECT * FROM prems WHERE Id = ?"
	row := premDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Prem{}
	if err := row.Scan(&m.Id, &m.INR, &m.Trisha); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("prem retrieved")
	return &m, nil
}
