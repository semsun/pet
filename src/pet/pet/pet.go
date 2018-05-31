package pet

import (
	"../utils"
)

type Pet struct {
	Id			string `json:"id"`
	Name			string `json:"pet_name"`
	Describtion		string `json:"pet_describtion"`
	GetTimestamp	int64  `json:"pet_get_timestmap"`
	Level			uint32 `json:"pet_level"`
	Quality			uint32 `json:"pet_quality"`
	Owner			string `json:"owner"`
}

func NewPet(name, describtion string, quality uint32) *Pet {
	timestamp := utils.CurTimestamp();
	pet := &Pet{utils.UniqueId(), name, describtion, timestamp, utils.INIT_LEVEL, quality, utils.SYSTEM_USER_NAME};

	return pet;
}

func (pet *Pet) TransferPet(owner string) error {
	pet.Owner = owner;
	return nil;
}

func (pet *Pet) feedPet(power int) {
	pet.Level = pet.Level + uint32(power) * pet.Quality;
}
