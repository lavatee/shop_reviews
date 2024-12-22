package reviews

type Review struct {
	Id         int    `db:"id"`
	Text       string `db:"text"`
	UserId     int    `db:"user_id"`
	ProductId  int    `db:"product_id"`
	Evaluation int    `db:"evaluation"`
}
