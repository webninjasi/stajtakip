package database

type Komisyon struct {
	AdSoyad string
	Dahil   bool
}

func (kom *Komisyon) Insert(conn *Connection) error {
	const sql string = "INSERT INTO komisyon (AdSoyad) VALUES (?);"

	result, err := conn.db.Exec(sql, kom.AdSoyad)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func KomisyonListesi(conn *Connection) ([]string, error) {
	const sql string = `SELECT AdSoyad FROM komisyon WHERE Dahil=1`

	q, err := conn.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	liste := []string{}
	for q.Next() {
		var AdSoyad string

		err = q.Scan(&AdSoyad)
		if err != nil {
			return nil, err
		}

		liste = append(liste, AdSoyad)
	}

	return liste, nil
}
