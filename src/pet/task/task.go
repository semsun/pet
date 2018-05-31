package task

import (
	//"fmt"
	"../utils"
)

type Task struct {
	Id		string `json:"id"`
	Name		string `json:"task_name"`
	Describtion	string `json:"task_descrition"`
	State		int    `json:"task_state"`
	AwardType	uint16 `json:"task_award_type"`
	AwardObjId	string `json:"task_award_obj"`
	Owner		string `json:"owner"`
}

const (
	CREATE_TASK int = 1000
	COMPLETE_TASK int = 9000
)

const (
	AWARD_TYPE_PET uint16 = 1000
	AWARD_TYPE_FOOD uint16 = 2000
)

func NewTask(name, describtion, awardObjId string, awardType uint16) *Task {
	task := &Task{utils.UniqueId(), name, describtion, CREATE_TASK, awardType, awardObjId, utils.SYSTEM_USER_NAME};

	return task;
}

func (task *Task) TakeTask(owner string) error {
	task.Owner = owner;
	return nil;
}

func (task *Task) GetAward() (uint16, string) {
	return task.AwardType, task.AwardObjId;
}
