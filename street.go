package my1562api

// StreetsList list of streets
type StreetsList []Street

// Street street from the 1562 database
type Street struct {
	ID   int
	Name string
}
