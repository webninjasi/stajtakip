package database

type RaporSehir struct {
	Sehir       string
	BasariOrani float32
}
type RaporKonu struct {
	Konu        string
	BasariOrani float32
}

func RaporSehirler(conn *Connection, year int) ([]RaporSehir, error) {
	const sql string = `SELECT 
Sehir, 
ROUND((SUM(KabulGun)*1.0/SUM(ToplamGun))*100,2)
as "%BasariOrani" 
FROM staj 
Where Year(Bitis) = (?) 
Group BY Sehir;`

	q, err := conn.db.Query(sql, year)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	liste := []RaporSehir{}
	for q.Next() {
		var Sehir string
		var BasariOrani float32
		err := q.Scan(&Sehir, &BasariOrani)
		if err != nil {
			return nil, err
		}
		liste = append(liste, RaporSehir{Sehir, BasariOrani})
	}

	return liste, nil
}

func RaporKonular(conn *Connection, year int) ([]RaporKonu, error) {
	const sql string = `SELECT 
KonuBaslik, 
ROUND((SUM(KabulGun)*1.0/SUM(ToplamGun))*100,2)
as "%BasariOrani" 
FROM staj 
Where Year(Bitis) = (?) 
Group BY KonuBaslik;`

	q, err := conn.db.Query(sql, year)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	liste := []RaporKonu{}
	for q.Next() {
		var Konu string
		var BasariOrani float32
		err := q.Scan(&Konu, &BasariOrani)
		if err != nil {
			return nil, err
		}
		liste = append(liste, RaporKonu{Konu, BasariOrani})
	}

	return liste, nil
}
