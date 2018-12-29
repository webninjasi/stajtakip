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
	const sql string = "INSERT INTO ogrenci (No, Ad, Soyad, Ogretim) VALUES (?, ?, ?, ?);"

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

func StajiTamamOgrenciler(conn *Connection) ([]Ogrenci, error) {
	const sql string = `SELECT o.No, o.Ad, o.Soyad, o.Ogretim
FROM ogrenci AS o, staj AS s LEFT JOIN denkstaj AS d ON s.OgrenciNo = d.OgrenciNo
WHERE o.No = s.OgrenciNo
GROUP BY o.No, o.Ad, o.Soyad, o.Ogretim
HAVING (SUM(s.Kabulgun) >= ? OR SUM(s.Kabulgun) + SUM(d.KabulGun) >= ? OR SUM(d.KabulGun) >= ?)
AND (SUM(s.ToplamGun) >= 60 OR SUM(d.ToplamGun) >= 60 OR SUM(s.ToplamGun) + SUM(d.ToplamGun) >= 60)`

	q, err := conn.db.Query(sql, cfg.GerekenStajGunu(), cfg.GerekenStajGunu(), cfg.GerekenStajGunu())
	if err != nil {
		return nil, err
	}
	defer q.Close()

	liste := []Ogrenci{}
	for q.Next() {
		var ogr Ogrenci

		err = q.Scan(&ogr.No, &ogr.Ad, &ogr.Soyad, &ogr.Ogretim)
		if err != nil {
			return nil, err
		}

		liste = append(liste, ogr)
	}

	return liste, nil
}

// TODO diğer fieldlar
// TODO update, delete fonksiyonları
