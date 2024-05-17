package cli

import (
	"fmt"
	"github.com/kjblanchard/baby_names/utils"
)

const SECOND_RATINGS_FILE_NAME = "ratings2.json"

func CompareResults() {
	myRatings := utils.GetRatedNames(RATINGS_FILE_NAME)
	importedNames := utils.GetRatedNames(SECOND_RATINGS_FILE_NAME)
	for _, rating := range myRatings {
		for _, otherRating := range importedNames {
			if rating.Name != otherRating.Name {
				continue
			}
			fmt.Printf("Both have the name %s, my rating is %d and their rating is %d\n", rating.Name, rating.Rank, otherRating.Rank)
		}
	}
}
