package cli

import (
	"encoding/json"
	"fmt"
	"github.com/kjblanchard/baby_names/api"
	"github.com/kjblanchard/baby_names/models"
	"github.com/kjblanchard/baby_names/utils"
	"math/rand"
	"os"
	"strconv"
)

const RATINGS_FILE_NAME = "ratings.json"

func checkIfAlreadyRated(name string, names []models.BabyName) bool {
	for _, item := range names {
		if item.Name == name {
			return true
		}
	}
	return false
}

func startInteraction(gender string) {
	fmt.Printf("Lets rate some names! Currently we are going to rate %s names\n%s\n0 is terrible, 10 is great\n***\n", gender, "***")
}

func presentName(name models.BabyName) int {
	fmt.Println(name.Name)
	var currentInput string
	for {
		fmt.Print("Enter a number between 0 and 10: ")
		fmt.Scanln(&currentInput)
		rank, err := strconv.Atoi(currentInput)
		if err == nil && rank >= 0 && rank <= 10 {
			return rank
		}
		fmt.Println("Invalid input. Please enter a number between 0 and 10.")
	}
}

func addBabyName(name models.BabyName) {
	ratedNames := utils.GetRatedNames(RATINGS_FILE_NAME)
	ratedNames = append(ratedNames, name)
	data, err := json.Marshal(map[string][]models.BabyName{"names": ratedNames})
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(RATINGS_FILE_NAME, data, 0644)
	if err != nil {
		panic(err)
	}
}

func RateNames() {
	for {
		gender := rand.Intn(3)
		namesApiResponse := api.GetBabyNames(gender)
		ratings := utils.GetRatedNames(RATINGS_FILE_NAME)
		var namesNotRated []models.BabyName
		for _, name := range namesApiResponse {
			if !checkIfAlreadyRated(name.Name, ratings) {
				namesNotRated = append(namesNotRated, name)
			} else {

			}
		}
		genderStr, err := utils.GetGenderString(gender)
		if err != nil {
			panic("Gender is not valid!")
		}
		startInteraction(genderStr)
		for _, babyName := range namesNotRated {
			babyName.Rank = presentName(babyName)
			addBabyName(babyName)
		}
		var currentInput string
		fmt.Print("Would you like to continue rating names? (1 to continue, anything else quits): ")
		fmt.Scanln(&currentInput)
		if currentInput != "1" {
			break
		}
	}
}
