package cli
import(
	"fmt"
	"sort"
	"github.com/kjblanchard/baby_names/utils"
)

func ShowResults() {
	ratings := utils.GetRatedNames(RATINGS_FILE_NAME)
	var namesAndRatings []struct {
		Name   string
		Rating int
	}
	for _, item := range ratings {
		name := item.Name
		rank := item.Rank
		namesAndRatings = append(namesAndRatings, struct {
			Name   string
			Rating int
		}{Name: name, Rating: rank})
	}

	// Sort namesAndRatings by rating (descending order)
	sort.Slice(namesAndRatings, func(i, j int) bool {
		return namesAndRatings[i].Rating > namesAndRatings[j].Rating
	})

	// Present names and ratings
	fmt.Println("Names and Ratings (Ordered by Rating):")
	for _, item := range namesAndRatings {
		fmt.Printf("%s: %d\n", item.Name, item.Rating)
	}
}