package database

type Staj struct {
	No        int
	Sinif     int
	Kurum     string
	Sehir     string
	Konu      string
	Baslangic string
	Bitis     string
}

func (stj *Staj) Insert(conn *Connection) error {
	const sql string = "INSERT INTO stajlar (`OgrenciNo`, `Sinif`, `Kurum`, `Sehir`, `Konu`, `Baslangic`, `Bitis`) VALUES (?, ?, ?, ?, ?, ?, ?);"

	result, err := conn.db.Exec(sql, stj.No, stj.Sinif, stj.Kurum, stj.Sehir, stj.Konu, stj.Baslangic, stj.Bitis)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

// TODO update, delete fonksiyonları
