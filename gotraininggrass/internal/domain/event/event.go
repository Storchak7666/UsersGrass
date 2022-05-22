package event

type Event struct {
	Id   int64
	Name string
}

type User struct {
	Id      int64  `db:"id"`
	Name    string `db:"Name"`
	Age     int64  `db:"Age"`
	City    string `db:"City"`
	Country string `db:"Country"`
}
