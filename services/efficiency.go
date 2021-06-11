package services

import(
	"github.com/BrandonReno/A.E.R/models"
)

func(db *DB) GetEfficiency(a *models.Athlete) (models.Efficiency, error){
	athlete_id := a.Athlete_ID
	sqlStatement := `SELECT * FROM public.efficiency WHERE Athlete_ID = $1`
	row := db.DBConn.QueryRow(sqlStatement, athlete_id)

	e := models.Efficiency{}

	err := row.Scan(&e.Athlete_ID, &e.Efficiency_ID, &e.Efficiency_Score, &e.Favorite_Sport)

	if err != nil{
		return models.Efficiency{}, err
	}
	//Calculations will go here but for now just return
	return e, nil
}