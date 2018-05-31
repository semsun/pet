package player

type Player struct {
	Name		string `json:"name"`
	Pets		[]string `json:"pets"`
	Tasks		[]string `json:"tasks"`
	Feeds		[]string `json:"feeds"`
}

func NewPlayer(name string) *Player {
	new_player := &Player{};
	new_player.Name = name;
	
	return new_player;
}

func (t *Player) AddPet(petId string) {
	t.Pets = append(t.Pets, petId);
}

func (t *Player) AddFeed(feedId string) {
	t.Feeds = append(t.Feeds, feedId);
}

func (t *Player) AddTask(taskId string) {
	t.Tasks= append(t.Tasks, taskId);
}
