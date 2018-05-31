package feed

import (
	"fmt"
	"../utils"
)

type Feed struct {
	Id				string `json:"id"`
	Name				string `json:"name"`
	Descrite			string `json:"descrite"`
	Power				uint32 `json:"power"`
	LastPower			uint32 `json:"power"`
	Owner				string `json:"owner"`
}

func NewFeed(name, describtion string, power uint32) *Feed {
	feed := &Feed{ utils.UniqueId(), name, describtion, power, power, utils.SYSTEM_USER_NAME};

	return feed;
}

func (feed *Feed) Consume(power uint32) error {
	last := feed.LastPower - power;
	if last < 0 {
		return fmt.Errorf("No more power");
	}

	return nil;
}

func (feed *Feed) TransferFeed(owner string) error {
	feed.Owner = owner;
	return nil;
}
