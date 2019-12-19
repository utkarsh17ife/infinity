package common_models

// Season ....
type Season int

// Season ....
const (
	Spring Season = iota
	Autmn
	Winter
	Summer
	Monsoon
)

func (s Season) String() string {
	return [...]string{"Spring", "Autmn", "Winter", "Summer", "Monsoon"}[s]
}
