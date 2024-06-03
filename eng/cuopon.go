package eng

func GetPromotions() []Promotion {
	return []Promotion{
		{
			ID:          "1",
			Label:       "10% discount for longer stay",
			Description: "You get a 10% off from total amount",
			Elgibilities: []Eligibility{
				{
					ID:           "1",
					Label:        "Applicable when the number of stay is longer than 5 nights",
					Param:        "booking.stay.nights",
					OperatorType: "GTE",
					Value:        Action{Operator: "flat", Value: "5"},
				},
			},
			Rewards: []Reward{
				{
					ID:         "1",
					Label:      "Discount the total amount",
					Param:      "total_price",
					ActionType: "SUB",
					Value:      Action{Operator: "MULPER", Value: "0.10"},
				},
			},
		},
		{
			ID:          "1",
			Label:       "Season Discount",
			Description: "You get a free night from us",
			Elgibilities: []Eligibility{
				{
					ID:           "1",
					Label:        "Applicable when the number of stay is longer than 5 nights",
					Param:        "booking.stay.nights",
					OperatorType: "GTE",
					Value:        Action{Operator: "flat", Value: "5"},
				},
			},
			Rewards: []Reward{
				{
					ID:         "1",
					Label:      "Discount the total amount",
					Param:      "room_rates.0.base.price",
					ActionType: "SUB",
					Value:      Action{Operator: "OVERRIDE", Value: "0"},
				},
			},
		},
	}
}
