package database

type DenkStaj struct {
	OgrenciNo       int
	KurumAdi        string
	OncekiOkul           string
	ToplamGun       int
	KabulGun        int
}

func (stj *DenkStaj) Insert(conn *Connection) error {
	const sql string = "INSERT INTO denkstaj (OgrenciNo, KurumAdi, OncekiOkul, ToplamGun, KabulGun) VALUES (?, ?, ?, ?, ?);"

	result, err := conn.db.Exec(sql, stj.OgrenciNo, stj.KurumAdi, stj.OncekiOkul, stj.ToplamGun, stj.KabulGun)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func OgrenciDenkStajListesi(conn *Connection, no int) ([]DenkStaj, error) {
	const sql string = `SELECT KurumAdi, OncekiOkul, ToplamGun, KabulGun
FROM denkstaj WHERE OgrenciNo=?`

	q, err := conn.db.Query(sql, no)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	liste := []DenkStaj{}
	for q.Next() {
		var stj DenkStaj

		err = q.Scan(&stj.KurumAdi, &stj.OncekiOkul, &stj.ToplamGun, &stj.KabulGun)
		if err != nil {
			return nil, err
		}

		liste = append(liste, stj)
	}

	return liste, nil
}
