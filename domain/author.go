package domain

// Represent database structure
type Author struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`

	Note []Note // has many
}

// Represent response API
type AuthorResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}
