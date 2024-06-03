package eng

func GetRates() []Rate {
	return []Rate{
		{
			RuleID:      "1",
			Label:       "Luxury Room Base Price Adjustment",
			Description: "Increase base price for luxury rooms for below 9000 to 10000",
			Elgibilities: []Eligibility{
				{
					ID:           "1",
					Label:        "Applicable when the amount is less than 9000",
					Param:        "room_rates.#(type==Luxury).base.price",
					OperatorType: "LTE",
					Value:        Action{Operator: "LTE", Value: "9000"},
				},
			},
			Rewards: []Reward{
				{
					ID:         "1",
					Label:      "Increase base price for the room",
					Param:      "room_rates.#(type==Luxury).base.price",
					ActionType: "ADD",
					Value:      Action{Operator: "SUB", Value: "2000"},
				},
			},
		},
		{
			RuleID:      "2",
			Label:       "Service Charge for Extra Breakfast",
			Description: "Add extra charge for breakfast when more than 2 pax opt for it",
			Elgibilities: []Eligibility{
				{
					ID:           "2",
					Label:        "Applicable when breakfast pax is more than 2",
					Param:        "booking.itinerary.#.room.services.#(title=Breakfast).pax.#",
					OperatorType: "GTE",
					Value:        Action{Operator: "flat", Value: "3"},
				},
			},
			Rewards: []Reward{
				{
					ID:         "2",
					Label:      "Add extra breakfast charge per pax",
					Param:      "room_rates.#.services.Breakfast.adult",
					ActionType: "OVERRIDE",
					Value:      Action{Operator: "OVERRIDE", Value: "800"},
				},
				{
					ID:         "2",
					Label:      "Add extra breakfast charge per pax",
					Param:      "room_rates.#.services.Breakfast.child",
					ActionType: "OVERRIDE",
					Value:      Action{Operator: "OVERRIDE", Value: "400"},
				},
				{
					ID:         "2",
					Label:      "Add extra breakfast charge per pax",
					Param:      "room_rates.#.services.Breakfast.infant",
					ActionType: "OVERRIDE",
					Value:      Action{Operator: "OVERRIDE", Value: "200"},
				},
			},
		},
		{
			RuleID:      "3",
			Label:       "Discount for Premium Members",
			Description: "Provide a discount on base price for premium members",
			Elgibilities: []Eligibility{
				{
					ID:           "3",
					Label:        "Applicable for premium membership levels",
					Param:        "membership.type",
					OperatorType: "EQ",
					Value:        Action{Operator: "EQ", Value: "Premium"},
				},
			},
			Rewards: []Reward{
				{
					ID:         "3",
					Label:      "Discount on base price",
					Param:      "room_rates.#.base.price",
					ActionType: "OVERRIDE",
					Value:      Action{Operator: "OVERRIDE", Value: "500"},
				},
			},
		},
		// {
		// 	RuleID:      "5",
		// 	Label:       "Long Stay Discount",
		// 	Description: "Provide a discount for stays longer than 5 nights",
		// 	Elgibilities: []Eligibility{
		// 		{
		// 			ID:           "5",
		// 			Label:        "Applicable when stay is longer than 5 nights",
		// 			Param:        "booking.stay.nights",
		// 			OperatorType: "GTE",
		// 			Value:        Action{Operator: "flat", Value: "6"},
		// 		},
		// 	},
		// 	Rewards: []Reward{
		// 		{
		// 			ID:         "5",
		// 			Label:      "Discount on total price",
		// 			Param:      "room_rates",
		// 			ActionType: "SUB",
		// 			Value:      Action{Operator: "SUB", Value: "200"},
		// 		},
		// 	},
		// },
	}
}
