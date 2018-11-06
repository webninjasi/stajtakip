package database

type Ogrenci struct {
	No      int
	Ad      string
	Soyad   string
	Ogretim int
}

func (ogr *Ogrenci) Insert(db *StajDatabase) error {
	const sql string = "INSERT INTO `stajtest`.`ogrenciler` (`No`, `Ad`, `Soyad`, `Ogretim`) VALUES (?, ?, ?, ?);"

	result, err := db.db.Exec(sql, ogr.No, ogr.Ad, ogr.Soyad, ogr.Ogretim)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

// TODO diğer fieldlar
// TODO update, delete fonksiyonları
