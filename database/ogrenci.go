package database

import (
	"stajtakip/cfg"
)

type Ogrenci struct {
	No      int
	Ad      string
	Soyad   string
	Ogretim int
}

func (ogr *Ogrenci) Insert(conn *Connection) error {
	const sql string = "INSERT INTO ogrenciler (`No`, `Ad`, `Soyad`, `Ogretim`) VALUES (?, ?, ?, ?);"

	result, err := conn.db.Exec(sql, ogr.No, ogr.Ad, ogr.Soyad, ogr.Ogretim)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func OgrenciListesi(conn *Connection) ([]*Ogrenci, error) {
	const sql string = "SELECT * FROM ogrenciler WHERE KabulEdilen >= ?"

	q, err := conn.db.Query(sql, cfg.GerekenStajGunu())
	if err != nil {
		return nil, err
	}
	defer q.Close()

	liste := []*Ogrenci{}
	for q.Next() {
		var ogr Ogrenci

		err = q.Scan(&ogr.No, &ogr.Ad, &ogr.Soyad, &ogr.Ogretim)
		if err != nil {
			return nil, err
		}

		liste = append(liste, &ogr)
	}

	return liste, nil
}

// TODO diğer fieldlar
// TODO update, delete fonksiyonları
