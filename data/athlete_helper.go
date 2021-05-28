package data

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func GenerateID(length int, charset string) string {
	id := make([]byte, length)
	for i := range id {
		id[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(id)
}

func GetDate() string {
	current := time.Now()
	date := time.Date(current.Year(), current.Month(), current.Day(), 0, 0, 0, 0, current.Location())
	return date.String()
}

func AddAthlete(a *Athlete) error {
	id := GenerateID(10, charset)
	current_date := GetDate()
	sqlStatement := `INSERT INTO public.athlete VALUES($1, $2, $3, $4, $5);`
	_, err := DBConn.Exec(sqlStatement, id, a.First_Name, a.Last_Name, a.Age, current_date)
	return err
}

func GetAthlete(id string) (Athlete, error) {
	var athlete Athlete
	sqlStatement := `SELECT * FROM public.athlete WHERE Athlete_ID = $1;`
	row := DBConn.QueryRow(sqlStatement, id)
	err := row.Scan(&athlete.Athlete_ID, &athlete.First_Name, &athlete.Last_Name, &athlete.Age, &athlete.Joined)

	if err != nil {
		return Athlete{}, err
	}

	return athlete, nil
}

func UpdateAthlete(a *Athlete) error {
	sqlStatement := `UPDATE public.athlete 
					SET First_Name = $1,
						Last_Name = $2,
						Age = $3
						WHERE Athlete_ID = $4;`
	_, err := DBConn.Exec(sqlStatement, a.First_Name, a.Last_Name, a.Age, a.Athlete_ID)
	return err
}

func DeleteAthlete(aid string) error {
	sqlStatement := `DELETE FROM public.athlete WHERE Athlete_ID = $1`
	_, err := DBConn.Exec(sqlStatement, aid)
	return err
}
