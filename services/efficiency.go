package services

import(
	"github.com/BrandonReno/A.E.R/models"
)

func(db *DB) GetEfficiency(a *models.Athlete) (models.Efficiency, error){

	sqlStatement := `SELECT * FROM public.efficiency WHERE Athlete_ID = $1`
	row := db.DBConn.QueryRow(sqlStatement, a.Athlete_ID)

	e := models.Efficiency{}

	err := row.Scan(&e.Athlete_ID, &e.Efficiency_ID, &e.Efficiency_Score, &e.Favorite_Sport)

	if err != nil{
		return models.Efficiency{}, err
	}

	return e, nil
}

func(db *DB) UpdateEfficiency(a *models.Athlete)error{
	workouts, err := db.GetUserWorkouts(a.Athlete_ID)

	if err != nil{
		return err
	}
	e, err := db.GetEfficiency(a)

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