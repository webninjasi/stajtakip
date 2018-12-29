package database

type OgrenciBelge struct {
	OgrenciNo      int
	Dosya      string
}

func (ogr *OgrenciBelge) Insert(conn *Connection) error {
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
