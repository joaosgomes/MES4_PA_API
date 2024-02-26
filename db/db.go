package db

import (
	"time"

	"api/models"

	"github.com/google/uuid"
)

func DbEvents() []*models.Event {
	return []*models.Event{

		{
			Id:          uuid.New(),
			Name:        "Oktoberfest",
			Description: "The 189th Oktoberfest will be held at Munich's Theresienwiese from September 21 to October 6, 2024.",
			Location:    "Munich",
			Date:        time.Date(2024, time.October, 6, 0, 0, 0, 0, time.UTC),
			Src_Image:   "https://images.unsplash.com/photo-1605493717198-ebe4197f606d?auto=format&fit=crop",
			CreateDate:  time.Now(),
			VersionDate: time.Now(),
			Capacity:    0,
			Temperature: float64(1),
			IsActive:    true,
		},
		{
			Id:          uuid.New(),
			Name:        "Academy Awards (Oscar)",
			Description: "An Academy Award of Merit, an Oscar statuette recognizes achievements that have an extraordinary influence upon the advancement of the motion picture arts and sciences. This award is generally reserved for achievements that have changed the course of filmmaking since their introduction.",
			Location:    "Hollywood",
			Date:        time.Date(2024, time.March, 10, 23, 0, 0, 0, time.UTC),
			Src_Image:   "https://images.unsplash.com/photo-1528669608060-7b7282176925?auto=format&fit=crop",
			CreateDate:  time.Now(),
			VersionDate: time.Now(),
			Capacity:    0,
			Temperature: float64(1),
			IsActive:    true,
		},
		{
			Id:          uuid.New(),
			Name:        "Super Bowl",
			Description: "The Super Bowl is the biggest and most important American football game of the year. It is the National Football League (NFL) yearly championship game. The game is played between the winning teams from the NFL's two conferences, the American Football Conference (AFC) and the National Football Conference (NFC).",
			Location:    "USA",
			Date:        time.Now().Add(time.Hour),
			Src_Image:   "https://images.unsplash.com/photo-1642189523635-912d71d9cf12?auto=format&fit=crop",
			CreateDate:  time.Now(),
			VersionDate: time.Now(),
			Capacity:    0,
			Temperature: float64(1),
			IsActive:    true,
		},
		{
			Id:          uuid.New(),
			Name:        "Tomorrowland",
			Description: "Set to welcome more than 400 of the world's finest electronic music artists across 16 mesmerizing stages during two weekends of magic",
			Location:    "Belgium",
			Date:        time.Now().Add(time.Hour),
			Src_Image:   "https://images.unsplash.com/photo-1520095972714-909e91b038e5?auto=format&fit=crop",
			CreateDate:  time.Now(),
			VersionDate: time.Now(),
			Capacity:    0,
			Temperature: float64(1),
			IsActive:    true,
		},
		{
			Id:          uuid.New(),
			Name:        "Euro 2024",
			Description: "Euro 2024 is the 17th UEFA European Football Championship, a quadrennial international men's football championship of Europe organized by UEFA. It brings together national teams from across Europe to compete for the prestigious title and showcases top-level football talent and exciting matches.",
			Location:    "Germany/Europe",
			Date:        time.Now().Add(time.Hour + 1),
			Src_Image:   "https://images.unsplash.com/photo-1434648957308-5e6a859697e8?auto=format&fit=crop",
			CreateDate:  time.Now(),
			VersionDate: time.Now(),
			Capacity:    0,
			Temperature: float64(1),
			IsActive:    true,
		},
		{
			Id:          uuid.New(),
			Name:        "Running of the Bulls",
			Description: "Running of the Bulls",
			Location:    "Pamplona/Spain",
			Date:        time.Now().Add(time.Hour),
			Src_Image:   "https://images.unsplash.com/photo-1636600489499-9f453f1af3e7?auto=format&fit=crop",
			CreateDate:  time.Now(),
			VersionDate: time.Now(),
			Capacity:    0,
			Temperature: float64(1),
			IsActive:    true,
		},
		{
			Id:          uuid.New(),
			Name:        "Day of the Dead",
			Description: "Mexicans honor their ancestors on Day of the Dead, but they're also reminding themselves that death is just a part of life. Hanging out with skeletons reminds people that one day they will be skeletons—but not for a very long time!",
			Location:    "Mexico",
			Date:        time.Now().Add(time.Hour),
			Src_Image:   "https://images.unsplash.com/photo-1633053663400-655b31fb88ac?auto=format&fit=crop",
			CreateDate:  time.Now(),
			VersionDate: time.Now(),
			Capacity:    0,
			Temperature: float64(1),
			IsActive:    true,
		},
		{
			Id:          uuid.New(),
			Name:        "Yi Peng Lantern Festival",
			Description: "It symbolizes the beauty of the full moon. As the exact date varies each year, the festival will be held on November 15-16,2024 During Yi Peng, locals gather at temples to listen to the monks' prayers and offer lanterns as a form of devotion.",
			Location:    "Mexico",
			Date:        time.Now().Add(time.Hour),
			Src_Image:   "https://images.unsplash.com/photo-1510673398445-94f476ef2cbc?auto=format&fit=crop",
			CreateDate:  time.Now(),
			VersionDate: time.Now(),
			Capacity:    0,
			Temperature: float64(1),
			IsActive:    true,
		},
	}
}

func DbImages() []*models.Image {
	return []*models.Image{

		{
			Id:          uuid.New(),
			Alt:         "",
			Description: "",
			Src_Image:   "https://images.unsplash.com/photo-1605493717198-ebe4197f606d?auto=format&fit=crop",
			Size:        1,
		},
		{
			Id:          uuid.New(),
			Alt:         "",
			Description: "",
			Src_Image:   "https://images.unsplash.com/photo-1528669608060-7b7282176925?auto=format&fit=crop",
			Size:        1,
		},
		{
			Id:          uuid.New(),
			Alt:         "",
			Description: "",
			Src_Image:   "https://images.unsplash.com/photo-1642189523635-912d71d9cf12?auto=format&fit=crop",
			Size:        1,
		},
		{
			Id:          uuid.New(),
			Alt:         "",
			Description: "",
			Src_Image:   "https://images.unsplash.com/photo-1520095972714-909e91b038e5?auto=format&fit=crop",
			Size:        1,
		},
		{
			Id:          uuid.New(),
			Alt:         "",
			Description: "",
			Src_Image:   "https://images.unsplash.com/photo-1434648957308-5e6a859697e8?auto=format&fit=crop",
			Size:        1,
		},
		{
			Id:          uuid.New(),
			Alt:         "",
			Description: "",
			Src_Image:   "https://images.unsplash.com/photo-1636600489499-9f453f1af3e7?auto=format&fit=crop",
			Size:        1,
		},
		{
			Id:          uuid.New(),
			Alt:         "",
			Description: "",
			Src_Image:   "https://images.unsplash.com/photo-1633053663400-655b31fb88ac?auto=format&fit=crop",
			Size:        1,
		},
		{
			Id:          uuid.New(),
			Alt:         "",
			Description: "",
			Src_Image:   "https://images.unsplash.com/photo-1510673398445-94f476ef2cbc?auto=format&fit=crop",
			Size:        1,
		},
	}
}
