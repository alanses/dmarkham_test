package signatures

type RandomJoke struct {
	Type  string           `json:"type"`
	Value *RandomJokeValue `json:"value"`
}

type RandomJokeValue struct {
	ID         int      `json:"id"`
	Joke       string   `json:"joke"`
	Categories []string `json:"categories"`
}
