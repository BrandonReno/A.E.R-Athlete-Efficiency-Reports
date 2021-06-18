package services

import(
	"github.com/BrandonReno/A.E.R/models"
)

func(db *DB) GetEfficiency(a *models.Athlete, ec chan <- models.Efficiency)  error{

	sqlStatement := `SELECT * FROM public.efficiency WHERE Athlete_ID = $1`
	row := db.DBConn.QueryRow(sqlStatement, a.Athlete_ID)

	var e models.Efficiency

	err := row.Scan(&e.Efficiency_ID, &e.Efficiency_Score, &e.Favorite_Sport, &e.Athlete_ID)

	if err != nil{
		return err
	}

	ec <- e

	return nil
}

func(db *DB) UpdateEfficiency(a *models.Athlete)error{
	ec := make(chan models.Efficiency, 1)
	wl := make(chan []models.Workout, 1)

	err := db.GetUserWorkouts(a.Athlete_ID, wl)
	
	if err != nil{
		return err
	}

	workouts := <- wl

	err = db.GetEfficiency(a, ec)

	if err != nil{
		return err
	}

	e := <- ec

	if err != nil{
		return err
	}

	e.CalculateEfficiency(workouts)
	e.CalculateSport(workouts)

	sqlStatement := `UPDATE public.efficiency 
					SET Efficiency_Score = $1,
						Favorite_Sport = $2
						WHERE Athlete_ID = $3`

	_, err = db.DBConn.Exec(sqlStatement, e.Efficiency_Score, e.Favorite_Sport, a.Athlete_ID)

	if err != nil{
		return err
	}
	return nil
}