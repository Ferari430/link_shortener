package link

import (
	"math/rand"
	"my_project/internal/stat"

	"gorm.io/gorm"
)

// то что будет в БД
type Link struct {
	gorm.Model
	Url   string      `json:"url" validate:"required,url"`
	Hash  string      `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}


func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: string(GenerateHash(10)),
	}
}

var letterRunes = []rune("abcdfghjkl")

func GenerateHash(n int) []rune {
	result := make([]rune, n)
	for i := range n {
		result[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return result

}
