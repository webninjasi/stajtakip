package database

import (
	dsql "database/sql"
)

type Staj struct {
	OgrenciNo       int
	KurumAdi        string
	Sehir           string
	KonuBaslik      string
	Baslangic       string
	Bitis           string
	Sinif           int
	ToplamGun       int
	KabulGun        int
	Degerlendirildi bool
}

func (stj *Staj) Insert(conn *Connection) error {
	const sql string = "INSERT INTO staj (OgrenciNo, KurumAdi, Sehir, KonuBaslik, Baslangic, Bitis, Sinif, ToplamGun) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"

	result, err := conn.db.Exec(sql, stj.OgrenciNo, stj.KurumAdi, stj.Sehir, stj.KonuBaslik, stj.Baslangic, stj.Bitis, stj.Sinif, stj.ToplamGun)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func OgrenciStajListesi(conn *Connection, no int) ([]Staj, error) {
	const sql string = `SELECT KurumAdi, Sehir, KonuBaslik, Baslangic, Bitis, Sinif, ToplamGun, KabulGun, Degerlendirildi
FROM Staj WHERE OgrenciNo=?`

	q, err := conn.db.Query(sql, no)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	liste := []Staj{}
	for q.Next() {
		var stj Staj
		var kabul dsql.NullInt64

		err = q.Scan(&stj.KurumAdi, &stj.Sehir, &stj.KonuBaslik, &stj.Baslangic,
			&stj.Bitis, &stj.Sinif, &stj.ToplamGun, &kabul, &stj.Degerlendirildi)
		if err != nil {
			return nil, err
		}
		if kabul.Valid {
			stj.KabulGun = int(kabul.Int64)
		}

		liste = append(liste, stj)
	}

	return liste, nil
}
