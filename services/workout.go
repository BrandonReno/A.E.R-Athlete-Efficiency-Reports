package services

import (
	"fmt"
	"github.com/BrandonReno/A.E.R/models"
)

func(d *DB) GetAllWorkouts(wc chan <- []models.Workout) error{
	var wl []models.Workout
	sqlStatement := `SELECT * FROM public.workout ORDER BY date DESC`

	rows, err := d.DBConn.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return err
	}

	var workout models.Workout
	for rows.Next(){
		err := rows.Scan(&workout.Workout_ID, &workout.Date, &workout.Description, &workout.Sport, &workout.Athlete_ID, &workout.Rating)
		if err != nil{
			return err
		}
		wl = append(wl, workout)
	}

	err = rows.Err()

	if err != nil{
		return err
	}

	wc <- wl

	return nil

}

func (d *DB) CreateWorkout(w *models.Workout) error{
	sqlStatement := `INSERT INTO public.workout(date, description, sport, athlete_id, rating)
					 VALUES($1, $2, $3, $4, $5);`
	_, err := d.DBConn.Exec(sqlStatement,w.Date, w.Description, w.Sport, w.Athlete_ID, w.Rating)
	if err != nil{
		return err
	}
	
	ac := make(chan models.Athlete, 1)

	err = d.GetAthlete(w.Athlete_ID, ac)

	athlete := <- ac

	if err != nil{
		return err
	}

	err = d.UpdateEfficiency(&athlete)
	if err != nil{
		return err
	}
	
	return nil
}

func (d *DB) GetUserWorkouts(id string, wc chan <- []models.Workout) error{
	var wl []models.Workout
	sqlStatement := `SELECT workout_id, date, description, sport, rating FROM public.workout WHERE Athlete_ID = $1`
	rows, err := d.DBConn.Query(sqlStatement, id)
	defer rows.Close()

	if err != nil{
		return err
	} 

	var workout models.Workout
	for rows.Next(){
		err := rows.Scan(&workout.Workout_ID, &workout.Date, &workout.Description, &workout.Sport, &workout.Rating)
		if err != nil{
			return err
		}
		wl = append(wl, workout)
	}

	err = rows.Err()

	if err != nil {
		return err
	}

	wc <- wl

	return nil
}

var cantFindWorkout error = fmt.Errorf("Can not find workout id in athletes workouts")

func (d *DB) GetSingleWorkout(aid string, wid int, wc chan <- models.Workout) (error){
	sqlStatement := `SELECT * FROM public.workout WHERE Workout_ID = $1 AND Athlete_ID = $2`

	row := d.DBConn.QueryRow(sqlStatement, wid, aid)

	if row.Err() != nil{
		return row.Err()
	}
	var w models.Workout
	err := row.Scan(&w.Workout_ID, &w.Date, &w.Description, &w.Sport, &w.Athlete_ID, &w.Rating)

	if err != nil{
		return err
	}

	wc <- w
	return nil
}

func (d *DB) UpdateWorkout(w *models.Workout) error{
	sqlStatement := `UPDATE public.workout 
					SET Date = $1,
						Description = $2,
						Sport = $3,
						Rating = $4
						WHERE Workout_ID = $5;`
	_, err := d.DBConn.Exec(sqlStatement, w.Date, w.Description, w.Sport, w.Rating, w.Workout_ID)
	return err
}

func (d *DB) DeleteWorkout(aid string, wid int) error{
	sqlStatement := `DELETE FROM public.workout WHERE Workout_ID = $1 AND Athlete_ID = $2`
	_, err := d.DBConn.Exec(sqlStatement, wid, aid)
	return err
}