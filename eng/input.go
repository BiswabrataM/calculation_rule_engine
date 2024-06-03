package eng

func GetInputJSON() string {
	return `{
        "total_price": 0,
        "room_rates": [
            {
                "roomId": "12",
                "type": "Luxury",
                "view": "Sea View",
                "base": {
                    "price": 7000,
                    "min_occupancy": {
                        "adult": 1,
                        "child": 0,
                        "infant": 0
                    },
                    "max_occupancy": {
                        "adult": 3,
                        "child": 2,
                        "infant": 2
                    }
                },
                "extra": {
                    "adult": 1000,
                    "child": 500,
                    "infant": 100
                },
                "services": {
                    "Breakfast": {
                        "adult": 500,
                        "child": 100,
                        "infant": 0
                    },
                    "Lunch": {
                        "adult": 1000,
                        "child": 500,
                        "infant": 100
                    },
                    "Dinner": {
                        "adult": 1000,
                        "child": 500,
                        "infant": 100
                    },
                    "Airport Transfer": {
                        "adult": 1000,
                        "child": 500,
                        "infant": 100
                    },
                    "Spa Services": {
                        "adult": 1000,
                        "child": 500,
                        "infant": 100
                    }
                }
            },
            {
                "roomId": "13",
                "type": "Normal",
                "view": "No view",
                "base": {
                    "price": 2020,
                    "min_occupancy": {
                        "adult": 1,
                        "child": 0,
                        "infant": 0
                    },
                    "max_occupancy": {
                        "adult": 3,
                        "child": 2,
                        "infant": 2
                    }
                },
                "extra": {
                    "adult": 500,
                    "child": 100,
                    "infant": 50
                },
                "services": {
                    "Breakfast": {
                        "adult": 500,
                        "child": 100,
                        "infant": 0
                    },
                    "Lunch": {
                        "adult": 1000,
                        "child": 500,
                        "infant": 100
                    },
                    "Dinner": {
                        "adult": 1000,
                        "child": 500,
                        "infant": 100
                    },
                    "Airport Transfer": {
                        "adult": 1000,
                        "child": 500,
                        "infant": 100
                    },
                    "Spa Services": {
                        "adult": 1000,
                        "child": 500,
                        "infant": 100
                    }
                }
            }
        ],
        "booking": {
            "date": "2024-05-27T10:30:00Z",
            "checkInDate": "2024-06-01",
            "checkOutDate": "2024-06-05",
            "stay": {
                "nights": 6,
                "days": 4
            },
            "itinerary": [
                {
                    "date": "2024-06-01",
                    "room": {
                        "roomId": "12",
                        "pax": [
                            {
                                "id": 1,
                                "name": "John Doe 1",
                                "age_category": "adult"
                            },
                            {
                                "id": 2,
                                "name": "John Doe 2",
                                "age_category": "adult"
                            }
                        ],
                        "services": [
                            {
                                "title": "Breakfast",
                                "pax": [
                                    {
                                        "id": 1,
                                        "name": "John Doe 1",
                                        "age_category": "adult"
                                    },
                                    {
                                        "id": 2,
                                        "name": "John Doe 2",
                                        "age_category": "adult"
                                    },
                                    {
                                        "id": 1,
                                        "name": "John Doe 1",
                                        "age_category": "adult"
                                    }
                                ]
                            }
                        ]
                    }
                },
                {
                    "date": "2024-06-02",
                    "room": {
                        "roomId": "13",
                        "pax": [
                            {
                                "id": 1,
                                "name": "John Doe 1",
                                "age_category": "adult"
                            },
                            {
                                "id": 2,
                                "name": "John Doe 2",
                                "age_category": "adult"
                            },
                            {
                                "id": 3,
                                "name": "Baby Doe",
                                "age_category": "infant"
                            },
                            {
                                "id": 4,
                                "name": "Kid Doe",
                                "age_category": "child"
                            }
                        ],
                        "services": [
                            {
                                "title": "Breakfast",
                                "pax": [
                                    {
                                        "id": 1,
                                        "name": "John Doe 1",
                                        "age_category": "adult"
                                    },
                                    {
                                        "id": 1,
                                        "name": "John Doe 1",
                                        "age_category": "adult"
                                    },
                                    {
                                        "id": 1,
                                        "name": "John Doe 1",
                                        "age_category": "adult"
                                    }
                                ]
                            }
                        ]
                    }
                },
                {
                    "date": "2024-06-03",
                    "room": {
                        "roomId": "13",
                        "pax": [
                            {
                                "id": 1,
                                "name": "John Doe 1",
                                "age_category": "adult"
                            },
                            {
                                "id": 2,
                                "name": "John Doe 2",
                                "age_category": "adult"
                            },
                            {
                                "id": 3,
                                "name": "Baby Doe",
                                "age_category": "infant"
                            },
                            {
                                "id": 4,
                                "name": "Kid Doe",
                                "age_category": "child"
                            }
                        ],
                        "services": [
                            {
                                "title": "Lunch",
                                "pax": [
                                    {
                                        "id": 1,
                                        "name": "John Doe 1",
                                        "age_category": "adult"
                                    }
                                ]
                            }
                        ]
                    }
                },
                {
                    "date": "2024-06-04",
                    "room": {
                        "roomId": "13",
                        "pax": [
                            {
                                "id": 1,
                                "name": "John Doe 1",
                                "age_category": "adult"
                            },
                            {
                                "id": 2,
                                "name": "John Doe 2",
                                "age_category": "adult"
                            },
                            {
                                "id": 3,
                                "name": "Baby Doe",
                                "age_category": "infant"
                            },
                            {
                                "id": 4,
                                "name": "Kid Doe",
                                "age_category": "child"
                            }
                        ],
                        "services": [
                            {
                                "title": "Spa",
                                "pax": [
                                    {
                                        "id": 1,
                                        "name": "John Doe 1",
                                        "age_category": "adult"
                                    },
                                    {
                                        "id": 2,
                                        "name": "John Doe 2",
                                        "age_category": "adult"
                                    },
                                    {
                                        "id": 3,
                                        "name": "Baby Doe",
                                        "age_category": "infant"
                                    },
                                    {
                                        "id": 4,
                                        "name": "Kid Doe",
                                        "age_category": "child"
                                    }
                                ]
                            }
                        ]
                    }
                }
            ]
        },
        "user": {
            "id": "123",
            "name": "So and so"
        },
        "membership": {
            "id": "12",
            "code": "M/RP/001",
            "type": "Premium"
        }
    }`
}
