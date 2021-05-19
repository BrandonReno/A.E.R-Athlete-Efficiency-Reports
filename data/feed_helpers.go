package data

func GetFeed() Workout_Feed {
	return workoutFeed // Returns the current Workout_Feed
}

func AddWorkout(w *Workout) {
	w.ID, w.User.ID = GetNextID() // Adjust the next ID and user ID of the passed workout to ensure they iterate
	workoutFeed = append(workoutFeed, w) // append the passed workout with adjusted IDs to the workout_feed
}

func UpdateWorkout(id int, w *Workout) error{
	ow, ind, err := FindWorkout(id) //Get the index of the URI ID and an error if it does not exist
	if err != nil{
		return err // if an error does exist return it back to the handlefunc
	}
	w.ID = ow.ID
	workoutFeed[ind] = w //Replace the workout_feed index with the new updated workout
	return nil 
}

func DeleteWorkout(id int) error{
	_, ind, err := FindWorkout(id) //Find the index for the workout in the feed
	if err != nil{ //if there is an error finding the workout return an error
		return err
	}
	
	workoutFeed = append(workoutFeed[:ind], workoutFeed[ind+1:]...) //remove the workout from the feed
	return nil //return nil error
}

func FindWorkout(id int) (*Workout,int, error){
	for ind, w := range workoutFeed{ //iterate through each workout in the workout feed
		if w.ID == id{ //if the workout id equals the URI id then return the index spot to be updated with nil error
			return w, ind, nil
		}
	}
	return nil, -1, ErrorWorkoutNotFound // if there is no matching id return -1 index with error stating that there is no matching id in the list
}

func GetNextID() (w_id int, a_id int){
	w_id = workoutFeed[len(workoutFeed) - 1].ID + 1 //Get the last workout in the list and return its id iterated by 1
	a_id = workoutFeed[len(workoutFeed) - 1].User.ID + 1 //Get the last workouts athlete in the list and return its id iterated by 1
	return
}