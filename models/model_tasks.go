package Models

type Tarefa struct {
	ID        int    `json:NOT NULL AUTO_INCREMENT`
	Text      string `json:"text"`
	Feito     bool   `json:"feito"`
	Status    string `json:"status"`
	IDUsuario int    `json:NOT NULL AUTO_INCREMENT`
}

func (b *Tarefa) TableName() string {
	return "tarefa"
}
