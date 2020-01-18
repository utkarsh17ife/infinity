package models

var Seasons = []string{"autmn", "monsoon", "winter", "summer"}

func validSeason(seasonToCheck string) bool {
	for _, item := range Seasons {
		if seasonToCheck == item {
			return true
		}
	}
	return false
}
