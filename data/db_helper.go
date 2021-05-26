package data

import(
	"math/rand"
	"fmt"
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

func AddAthlete(a *Athlete) error{
	id := GenerateID(10, charset)
	current_date := time.Now()
	sqlStatement := `INSERT INTO public.athlete VALUES($1, $2, $3, $4, $5);`
	_, err := DBConn.Exec(sqlStatement, id, a.First_Name, a.Last_Name, a.Age, current_date)
	return err
}

func GetAthlete(id string) (Athlete, error){
	var athlete Athlete
	sqlStatement := `SELECT * FROM public.athlete WHERE Athlete_ID = $1;`
	row := DBConn.QueryRow(sqlStatement, id)
	err := row.Scan(&athlete.Athlete_ID, &athlete.First_Name, &athlete.Last_Name, &athlete.Age, &athlete.Joined)

	if err != nil{
		return Athlete{},err
	}

	return athlete, nil
}

func UpdateAthlete(a *Athlete) error{
	sqlStatement := `UPDATE public.athlete 
					SET First_Name = $1,
						Last_Name = $2,
						Age = $3
						WHERE Athlete_ID = $4;`
	_, err := DBConn.Exec(sqlStatement, a.First_Name, a.Last_Name, a.Age, a.Athlete_ID)
	return err
}








func CreateWorkout(w *Workout) error{
	sqlStatement := `INSERT INTO public.workout(date, description, sport, athlete_id, rating)
					 VALUES($1, $2, $3, $4, $5);`
	_, err := DBConn.Exec(sqlStatement,w.Date, w.Description, w.Sport, w.Athlete_ID, w.Rating)
	return err
}

func GetUserWorkouts(id string) ([]Workout, error){
	var wl []Workout
	sqlStatement := `SELECT date, description, sport, rating FROM public.workout WHERE Athlete_ID = $1`
	rows, err := DBConn.Query(sqlStatement, id)
	defer rows.Close()

	if err != nil{
		return nil, err
	} 

	var workout Workout
	for rows.Next(){
		err := rows.Scan(&workout.Date, &workout.Description, &workout.Sport, &workout.Rating)
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

func GetSingleWorkout(aid string, wid int) (Workout, error){
	wl, err := GetUserWorkouts(aid)

	if err != nil{
		return Workout{}, err
	}

	for _, w := range wl{
		if w.Workout_ID == wid{
			return w, nil
		}
	}

	return Workout{}, cantFindWorkout
}

func UpdateWorkout(w *Workout) error{
	sqlStatement := `UPDATE public.workout 
					SET Date = $1,
						Description = $2,
						Sport = $3,
						Rating = $4
						WHERE Workout_ID = $5;`
	_, err := DBConn.Exec(sqlStatement, w.Date, w.Description, w.Sport, w.Rating, w.Workout_ID)
	return err
}
