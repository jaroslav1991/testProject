package randomizer

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

type RandomAnnouncements struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	IdPhoto     string `json:"id_photo"`
}

var randomName = []string{"notebook", "car", "house", "table", "TV", "Bike"}
var randomPrice = []int{12500, 24560, 19990, 75322, 25910, 33445, 94999, 89150, 15000, 22190}

func Init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomBite(n int) *RandomAnnouncements {
	var randomAnnouncement RandomAnnouncements
	num := make([]int, n)
	c := make([]string, n)

	for item := range c {
		c[item] = randomName[rand.Intn(len(randomName))]
		randomAnnouncement.Name = c[item]
	}

	for j := range num {
		num[j] = randomPrice[rand.Intn(len(randomPrice))]
		randomAnnouncement.Price = num[j]
	}

	return &RandomAnnouncements{
		Name:        randomAnnouncement.Name,
		Description: "test description",
		Price:       randomAnnouncement.Price,
		IdPhoto:     "qwerty123",
	}
}

func SenderRandomString() []byte {
	Init()
	var randomString = RandomBite(1)

	bodyRes, err := json.Marshal(randomString)
	if err != nil {
		return nil
	}
	log.Println(string(bodyRes))
	return bodyRes
}
