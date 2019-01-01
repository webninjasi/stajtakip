package database

type OgrenciEk struct {
	OgrenciNo      int
	Dosya      string
}

func (ogr *OgrenciEk) Insert(conn *Connection) error {
	const sql string = "INSERT INTO ogrenciek (OgrenciNo, Dosya) VALUES (?, ?);"

	result, err := conn.db.Exec(sql, ogr.OgrenciNo, ogr.Dosya)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
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
