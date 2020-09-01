package Models

type Usuario struct {
	ID    int    `json:"NOT NULL AUTO_INCREMENT"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func (b *Usuario) TableName() string {
	return "usuario"
}
