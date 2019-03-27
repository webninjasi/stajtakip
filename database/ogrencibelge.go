package database

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	dbsql "database/sql"
)

type OgrenciEk struct {
	OgrenciNo      int
	Dosya      string
}

func (ogr *OgrenciEk) Insert(conn *Connection) (*dbsql.Tx, error) {
	const sql string = "INSERT INTO ogrenciek (OgrenciNo, Dosya) VALUES (?, ?);"

	tx, err := conn.db.Begin()
	if err != nil {
		return nil, err
	}

	if _, err := tx.Exec(sql, ogr.OgrenciNo, ogr.Dosya); err == nil {
		return tx, nil
	}

	tx.Rollback()

	me, ok := err.(*mysql.MySQLError)
	if ok {
		if me.Number == 1062 {
			return nil, errors.New("Bu öğrenci için ek zaten var!")
		}
		if me.Number == 1406 {
			return nil, errors.New("Dosya ismi çok uzun!")
		}
	}

	logrus.WithFields(logrus.Fields{
		"err": err,
	}).Error("Öğrenci eki eklenirken veritabanında bir hata oluştu!")

	return nil, errors.New("Veritabanında bir hata oluştu!")
}

func OgrenciEkBul(conn *Connection, no int) (*OgrenciEk, error) {
	const sql string = `SELECT OgrenciNo, Dosya FROM ogrenciek WHERE OgrenciNo=?`

	q, err := conn.db.Query(sql, no)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	if !q.Next() {
		return nil, ErrVeriBulunamadi
	}

	var ogr OgrenciEk
	if err = q.Scan(&ogr.OgrenciNo, &ogr.Dosya); err != nil {
		return nil, err
	}

	return &ogr, nil
}
