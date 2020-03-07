package schema

type Todo struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}
