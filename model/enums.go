//go:generate enumer -type=Season -text -output=season.gen.go
//go:generate enumer -type=Diet -text -output=diet.gen.go
//go:generate enumer -type=Tag -text -output=tag.gen.go

package model

type Season int

const (
	Winter Season = iota
	Spring
	Summer
	Autumn
)

type Diet int

const (
	GlutenFree Diet = iota
	Vegetarian
	Vegan
)

type Tag int

const (
	HighProtein Tag = iota
	Fast
)
