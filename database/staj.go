package database

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

func KonuListesi(conn *Connection) ([]string, error) {
	const sql string = `SELECT Baslik FROM Konu WHERE Aktif=1`

	q, err := conn.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	liste := []string{}
	for q.Next() {
		var baslik string

		err = q.Scan(&baslik)
		if err != nil {
			return nil, err
		}

		liste = append(liste, baslik)
	}

	return liste, nil
}

// TODO update, delete fonksiyonları
