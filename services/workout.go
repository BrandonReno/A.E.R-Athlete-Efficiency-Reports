package services

import (
	"fmt"
	"github.com/BrandonReno/A.E.R/models"
)

func CreateWorkout(w *models.Workout) error{
	sqlStatement := `INSERT INTO public.workout(date, description, sport, athlete_id, rating)
					 VALUES($1, $2, $3, $4, $5);`
	_, err := DBConn.Exec(sqlStatement,w.Date, w.Description, w.Sport, w.Athlete_ID, w.Rating)
	return err
}

func GetUserWorkouts(id string) ([]models.Workout, error){
	var wl []models.Workout
	sqlStatement := `SELECT workout_id, date, description, sport, rating FROM public.workout WHERE Athlete_ID = $1`
	rows, err := DBConn.Query(sqlStatement, id)
	defer rows.Close()

	if err != nil{
		return nil, err
	} 

	var workout models.Workout
	for rows.Next(){
		err := rows.Scan(&workout.Workout_ID, &workout.Date, &workout.Description, &workout.Sport, &workout.Rating)
		if err != nil{
			return nil, err
		}
		wl = append(wl, workout)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return wl, nil
}

var cantFindWorkout error = fmt.Errorf("Can not find workout id in athletes workouts")

func GetSingleWorkout(aid string, wid int) (models.Workout, error){
	wl, err := GetUserWorkouts(aid)

	if err != nil{
		return models.Workout{}, err
	}

	for _, w := range wl{
		if w.Workout_ID == wid{
			return w, nil
		}
	}

	return models.Workout{}, cantFindWorkout
}

func UpdateWorkout(w *models.Workout) error{
	sqlStatement := `UPDATE public.workout 
					SET Date = $1,
						Description = $2,
						Sport = $3,
						Rating = $4
						WHERE Workout_ID = $5;`
	_, err := DBConn.Exec(sqlStatement, w.Date, w.Description, w.Sport, w.Rating, w.Workout_ID)
	return err
}

func DeleteWorkout(aid string, wid int) error{
	sqlStatement := `DELETE FROM public.workout WHERE Workout_ID = $1 AND Athlete_ID = $2`
	_, err := DBConn.Exec(sqlStatement, wid, aid)
	return err
}