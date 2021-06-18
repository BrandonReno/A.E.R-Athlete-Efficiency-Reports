package services

import (
	"math/rand"
	"time"
	"github.com/BrandonReno/A.E.R/models"
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

func GetDate() string{
	current := time.Now()
	date := time.Date(current.Year(), current.Month(), current.Day(),0,0,0,0, current.Location())
	return date.String()
}

func (d *DB) AddAthlete(a *models.Athlete) error{
	id := GenerateID(10, charset)
	current_date := GetDate()
	sqlStatement := `INSERT INTO public.athlete VALUES($1, $2, $3, $4, $5);`
	_, err := d.DBConn.Exec(sqlStatement, id, a.First_Name, a.Last_Name, a.Age, current_date)
	return err
}

func (d *DB) GetAthlete(id string, ac chan <- models.Athlete) error{
	var athlete models.Athlete
	sqlStatement := `SELECT * FROM public.athlete WHERE Athlete_ID = $1;`
	row := d.DBConn.QueryRow(sqlStatement, id)
	err := row.Scan(&athlete.Athlete_ID, &athlete.First_Name, &athlete.Last_Name, &athlete.Age, &athlete.Joined)

	if err != nil{
		return err
	}

	ac <- athlete

	return  nil
}

func (d *DB) GetAllAthletes(ac chan []models.Athlete) error{
	sqlStatement := `SELECT * FROM public.athlete;`
	rows, err := d.DBConn.Query(sqlStatement)
	
	if err != nil{
		return err
	}

	defer rows.Close()

	var al []models.Athlete
	var athlete models.Athlete
	for rows.Next(){
		err := rows.Scan(&athlete.Athlete_ID, &athlete.First_Name, &athlete.Last_Name, &athlete.Age, &athlete.Joined)
		if err != nil{
			return err
		}
		al = append(al, athlete)
	}
	err = rows.Err()
	if err != nil{
		return err
	}

	ac <- al
	
	return nil
}

func (d *DB) UpdateAthlete(a *models.Athlete) error{
	sqlStatement := `UPDATE public.athlete 
					SET First_Name = $1,
						Last_Name = $2,
						Age = $3
						WHERE Athlete_ID = $4;`
	_, err := d.DBConn.Exec(sqlStatement, a.First_Name, a.Last_Name, a.Age, a.Athlete_ID)
	return err
}

func (d *DB) DeleteAthlete(aid string) error{
	sqlStatement := `DELETE FROM public.athlete WHERE Athlete_ID = $1`
	_, err := d.DBConn.Exec(sqlStatement, aid)
	return err
}
