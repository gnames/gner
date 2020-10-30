package txt

type Volume struct {
	ID    string `json:"volumeId"`
	Pages []Page
}

type Page struct {
	ID   string `json:"pageId"`
	Text []rune `json:"-"`
}
