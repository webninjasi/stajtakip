package database

import (
	dsql "database/sql"
	"github.com/go-sql-driver/mysql"
)

const TarihSaatFormati = "2006-01-02 15:04"
const TarihFormati = "2006-01-02"
const SaatFormati = "15:04"

type MulakatOgrenci struct {
	OgrenciNo      int
	Ad string
	Soyad string
	Ogretim int
	StajBaslangic      string
	Tarih      string
	Saat      string
	KomisyonUye1      string
	KomisyonUye2      string
}

type Mulakat struct {
	OgrenciNo      int
	StajBaslangic      string
	TarihSaat      string
	KomisyonUye1      string
	KomisyonUye2      string
}

func (mul *Mulakat) Update(conn *Connection) error {
	const sql string = `UPDATE mulakat
SET TarihSaat=?, KomisyonUye1=?, KomisyonUye2=?
WHERE OgrenciNo=? AND StajBaslangic=?`

	result, err := conn.db.Exec(sql, mul.TarihSaat, mul.KomisyonUye1, mul.KomisyonUye2, mul.OgrenciNo, mul.StajBaslangic)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

type MulakatSonuc struct {
	OgrenciNo      int
	StajBaslangic string
	PuanDevam      int
	PuanCaba      int
	PuanVakit      int
	PuanAmireDavranis int
	PuanIsArkadasaDavranis int
	PuanProje      int
	PuanDuzen      int
	PuanSunum      int
	PuanIcerik      int
	PuanMulakat      int
}

func (mul *MulakatSonuc) Update(conn *Connection) error {
	const sql string = `UPDATE mulakat
SET PuanDevam=?, PuanCaba=?, PuanVakit=?, PuanAmireDavranis=?,
PuanIsArkadasaDavranis=?, PuanProje=?, PuanDuzen=?, PuanSunum=?, PuanIcerik=?, PuanMulakat=?)
WHERE OgrenciNo=? AND StajBaslangic=?`

	result, err := conn.db.Exec(
    sql, mul.PuanDevam, mul.PuanCaba, mul.PuanVakit, mul.PuanAmireDavranis,
		mul.PuanIsArkadasaDavranis, mul.PuanProje, mul.PuanDuzen, mul.PuanSunum,
		mul.PuanIcerik, mul.PuanMulakat, mul.OgrenciNo, mul.StajBaslangic)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func MulakatListesi(conn *Connection) ([]MulakatOgrenci, error) {
	const sql string = `SELECT m.OgrenciNo, o.Ad, o.Soyad, o.Ogretim,
m.StajBaslangic, m.TarihSaat, m.KomisyonUye1, m.KomisyonUye2
FROM mulakat AS m, ogrenci AS o WHERE m.OgrenciNo=o.No AND m.PuanMulakat IS NULL`

	q, err := conn.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	liste := []MulakatOgrenci{}
	for q.Next() {
		var mul MulakatOgrenci
		var tarihSaat mysql.NullTime
		var kom1, kom2 dsql.NullString

		err = q.Scan(&mul.OgrenciNo, &mul.Ad, &mul.Soyad, &mul.Ogretim, &mul.StajBaslangic,
								 &tarihSaat, &kom1, &kom2)
		if err != nil {
			return nil, err
		}

		if tarihSaat.Valid {
			mul.Tarih = tarihSaat.Time.Format(TarihFormati)
			mul.Saat = tarihSaat.Time.Format(SaatFormati)
		} else {
			mul.Tarih = ""
			mul.Saat = ""
		}

		if kom1.Valid {
			mul.KomisyonUye1 = kom1.String
		} else {
			mul.KomisyonUye1 = "-"
		}

		if kom2.Valid {
			mul.KomisyonUye2 = kom2.String
		} else {
			mul.KomisyonUye2 = "-"
		}

		liste = append(liste, mul)
	}

	return liste, nil
}

func MulakatListesiOlustur(conn *Connection) error {
	const sql string = `INSERT INTO mulakat (OgrenciNo, StajBaslangic)
SELECT s.OgrenciNo, s.Baslangic
FROM staj AS s LEFT JOIN mulakat AS m
ON s.OgrenciNo=m.OgrenciNo AND s.Baslangic=m.StajBaslangic
WHERE m.StajBaslangic IS NULL AND NOT s.Degerlendirildi`

	result, err := conn.db.Exec(sql)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
