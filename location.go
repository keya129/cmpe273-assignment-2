package main


type Location struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	State   	string    `json:"state"`
	Zip      	string    		`json:"zip"`
	Coordinate	Coordinate		`json:"cordinate"`
}
type Coordinate struct{
	Lat float64				`json:"lat"`
	Lng float64				`json:"lng"`
}

type MyLocations []MyLocations
