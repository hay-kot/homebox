package faker

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Faker struct {
}

func NewFaker() *Faker {
	rand.Seed(time.Now().UnixNano())
	return &Faker{}
}

func (f *Faker) RandomString(length int) string {

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (f *Faker) RandomEmail() string {
	return f.RandomString(10) + "@email.com"
}

func (f *Faker) RandomBool() bool {
	return rand.Intn(2) == 1
}

func (f *Faker) RandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}
