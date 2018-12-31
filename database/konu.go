package database

type Konu struct {
	Baslik string
	Aktif   bool
}

func (konu *Konu) Insert(conn *Connection) error {
	const sql string = "INSERT INTO konu (Baslik) VALUES (?);"

	result, err := conn.db.Exec(sql, konu.Baslik)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func(konu *Konu) Update(conn *Connection) error {
	const sql string = "UPDATE konu SET Aktif = (?) WHERE Baslik = (?);"

	_, err := conn.db.Exec(sql, konu.Aktif, konu.Baslik)
	if err != nil {
		return err
	}

	return nil
}

func(konu *Konu) Delete(conn *Connection) error {
	const sql string = "DELETE FROM konu WHERE Baslik = (?);"

	_, err := conn.db.Exec(sql, konu.Baslik)
	if err != nil {
		return err
	}

	return nil
}

func KonuListesi(conn *Connection) ([]Konu, error) {
	const sql string = `SELECT Baslik, Aktif FROM konu`

	q, err := conn.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	liste := []Konu{}
	for q.Next() {
		var Baslik string
		var Aktif bool
		err = q.Scan(&Baslik, &Aktif)
		if err != nil {
			return nil, err
		}

		liste = append(liste, Konu{Baslik,Aktif})
	}

	return liste, nil
}
