package engine

type Game struct {
	Id      string         `json:"id"`
	Players []Player       `json:"players"`
	Turn    int            `json:"turn"`
	State   string         `json:"state"`
	Data    map[string]any `json:"data"`
}

type Player struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Coins     int      `json:"coins"`
	Influence []string `json:"influence"`
	Alive     bool     `json:"alive"`
}

type Action struct {
	GameId   string         `json:"game_id"`
	PlayerId string         `json:"player_id"`
	Type     string         `json:"type"`
	Payload  map[string]any `json:"payload"`
}
