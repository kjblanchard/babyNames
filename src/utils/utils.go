package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kjblanchard/baby_names/models"
)

func GetGenderString(genderInt int) (string, error) {
	switch genderInt {
	case 0:
		return "Boy", nil
	case 1:
		return "Girl", nil
	case 2:
		return "Neutral", nil
	default:
		return "", fmt.Errorf("gender not recognized: %d", genderInt)
	}
}

func ChangeToWorkDir() {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exeDir := filepath.Dir(exePath)
	err = os.Chdir(exeDir)
	if err != nil {
		panic(err)
	}
}

func GetRatedNames(filename string) []models.BabyName {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var ratingsFile models.RatingsFile
	err = json.Unmarshal(file, &ratingsFile)
	if err != nil {
		panic(err)
	}
	return ratingsFile.BabyNames
}
