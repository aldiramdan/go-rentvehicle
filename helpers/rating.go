package helpers

func CalculateNewRating(totalRent int, currentRating, newRating float64) float64 {

	newRating = (currentRating*float64(totalRent) + newRating) / float64(totalRent+1)

	return newRating

}
