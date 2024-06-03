package eng

import (
	"encoding/json"
	"log"
)

type Occupancy struct {
	Adult  int `json:"adult"`
	Child  int `json:"child"`
	Infant int `json:"infant"`
}

type Base struct {
	Price        int       `json:"price"`
	MinOccupancy Occupancy `json:"min_occupancy"`
	MaxOccupancy Occupancy `json:"max_occupancy"`
}

type Extra struct {
	Adult  int `json:"adult"`
	Child  int `json:"child"`
	Infant int `json:"infant"`
}

type ServiceRates struct {
	Breakfast       Extra `json:"Breakfast"`
	Lunch           Extra `json:"Lunch"`
	Dinner          Extra `json:"Dinner"`
	AirportTransfer Extra `json:"Airport Transfer"`
	SpaServices     Extra `json:"Spa Services"`
}

type RoomRate struct {
	RoomID   string       `json:"roomId"`
	Type     string       `json:"type"`
	View     string       `json:"view"`
	Base     Base         `json:"base"`
	Extra    Extra        `json:"extra"`
	Services ServiceRates `json:"services"`
}

type Pax struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	AgeCategory string `json:"age_category"`
}

type Service struct {
	Title string `json:"title"`
	Pax   []Pax  `json:"pax"`
}

type RoomBooking struct {
	RoomID   string    `json:"roomId"`
	Pax      []Pax     `json:"pax"`
	Services []Service `json:"services"`
}

type Itinerary struct {
	Date string      `json:"date"`
	Room RoomBooking `json:"room"`
}

type Booking struct {
	Date         string `json:"date"`
	CheckInDate  string `json:"checkInDate"`
	CheckOutDate string `json:"checkOutDate"`
	Stay         struct {
		Nights int `json:"nights"`
		Days   int `json:"days"`
	} `json:"stay"`
	Itinerary []Itinerary `json:"itinerary"`
}

type Data struct {
	TotalPrice int        `json:"total_price"`
	RoomRates  []RoomRate `json:"room_rates"`
	Booking    Booking    `json:"booking"`
}

func CalculatePrice(jsonData string) string {
	var data Data

	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatal(err)
	}

	totalPrice := 0

	for _, itinerary := range data.Booking.Itinerary {
		roomID := itinerary.Room.RoomID
		var roomRate RoomRate

		for _, rr := range data.RoomRates {
			if rr.RoomID == roomID {
				roomRate = rr
				break
			}
		}

		// Base price
		totalPrice += roomRate.Base.Price

		// Extra charges for additional occupants
		numAdults := 0
		numChildren := 0
		numInfants := 0

		for _, pax := range itinerary.Room.Pax {
			switch pax.AgeCategory {
			case "adult":
				numAdults++
			case "child":
				numChildren++
			case "infant":
				numInfants++
			}
		}

		if numAdults > roomRate.Base.MinOccupancy.Adult {
			totalPrice += (numAdults - roomRate.Base.MinOccupancy.Adult) * roomRate.Extra.Adult
		}
		if numChildren > roomRate.Base.MinOccupancy.Child {
			totalPrice += (numChildren - roomRate.Base.MinOccupancy.Child) * roomRate.Extra.Child
		}
		if numInfants > roomRate.Base.MinOccupancy.Infant {
			totalPrice += (numInfants - roomRate.Base.MinOccupancy.Infant) * roomRate.Extra.Infant
		}

		// Charges for services
		for _, service := range itinerary.Room.Services {
			var serviceRate Extra
			switch service.Title {
			case "Breakfast":
				serviceRate = roomRate.Services.Breakfast
			case "Lunch":
				serviceRate = roomRate.Services.Lunch
			case "Dinner":
				serviceRate = roomRate.Services.Dinner
			case "Airport Transfer":
				serviceRate = roomRate.Services.AirportTransfer
			case "Spa":
				serviceRate = roomRate.Services.SpaServices
			}

			for _, pax := range service.Pax {
				switch pax.AgeCategory {
				case "adult":
					totalPrice += serviceRate.Adult
				case "child":
					totalPrice += serviceRate.Child
				case "infant":
					totalPrice += serviceRate.Infant
				}
			}
		}
	}

	data.TotalPrice = totalPrice

	json, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	return string(json)
}
