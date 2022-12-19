package public

type ReqNotes struct {
	BookID int    `json:"book_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
