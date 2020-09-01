package testes

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	Config "github.com/renancrc99/1sti.challenge.backend/config"
	Controllers "github.com/renancrc99/1sti.challenge.backend/controllers"
	"github.com/renancrc99/1sti.challenge.backend/graph/model"
)

func TestAdicionarUsuarioComSucesso(t *testing.T) {
	var err error
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	usuario := model.UsuarioInput{ID: 456423148, Nome: "Carlos Antônio", Email: "carlosant@gmail.com.br"}
	err = Controllers.CriarUsuario(&usuario)
	if err != nil {
		t.Errorf("Erro ao tentar incluir usuário: %s", err)
	} else {
		t.Log("Teste com sucesso.")
	}
	//Excluir usuario apos o teste
	Controllers.DeletarUsuario(456423148)
}

func TestAdicionarUsuarioSemNome(t *testing.T) {
	var err error
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	usuario := model.UsuarioInput{ID: 1213457, Email: "Andrecarv@hotmail.com.br"}
	err = Controllers.CriarUsuario(&usuario)
	if err != nil {
		t.Log("Teste com sucesso.")
	} else {
		t.Errorf("Incluiu usuário sem o nome.")
	}
}

func TestCriarEditarTarefaPorUsuario(t *testing.T) {
	var err error
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	usuario := model.UsuarioInput{ID: 45612345}
	tarefas := model.TarefaInput{ID: 45645132154, Feito: true, Text: "Fazer o café", Status: "Concluído", UsuarioInput: &usuario}
	err = Controllers.CriarTarefaPorUsuario(tarefas)
	if err != nil {
		t.Errorf("Erro ao tentar incluir tarefa: %s", err)
	} else {
		t.Log("Teste com sucesso.")
	}
	tarefas.Text = "Novo teste"
	err = Controllers.EditarTarefaPorUsuario(tarefas)
	if err != nil {
		t.Errorf("Erro ao tentar editar tarefa: %s", err)
	} else {
		t.Log("Teste com sucesso.")
	}
	//Excluir tarefa apos o teste
	Controllers.DeletarTarefa(45645132154)
}

func TestMudarStatusTarefa(t *testing.T) {
	var err error
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	usuario := model.UsuarioInput{ID: 123456789}
	tarefas := model.TarefaInput{ID: 123456, Feito: true, Text: "Pintar a casa", Status: "Em progresso", UsuarioInput: &usuario}
	tarefas.Status = "Concluído"
	err = Controllers.MudarStatusTarefa(tarefas)
	if err != nil {
		t.Errorf("Erro ao tentar alterar o status da tarefa: %s", err)
	} else {
		t.Log("Teste com sucesso.")
	}
	//Excluir tarefa apos o teste
	Controllers.DeletarTarefa(123456)

}

func TestListarTarefasUsuario(t *testing.T) {
	var err error
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	usuario := model.UsuarioInput{ID: 123456789, Nome: "Jose Figueiredo", Email: "josefigueira@outlook.com"}
	Controllers.CriarUsuario(&usuario)
	tarefa := model.TarefaInput{ID: 12345678, Feito: true, Text: "Lavar o carro", Status: "Pendente", UsuarioInput: &usuario}
	Controllers.CriarTarefaPorUsuario(tarefa)
	tarefa2 := model.TarefaInput{ID: 123456789, Feito: false, Text: "Comprar suco", Status: "Pendente", UsuarioInput: &usuario}
	Controllers.CriarTarefaPorUsuario(tarefa2)
	arrtarefas := Controllers.ListarTarefasUsuario(usuario)

	if err != nil || arrtarefas[0].ID != 12345678 || arrtarefas[0].Text != "Lavar o carro" {
		t.Errorf("Erro ao listar as tarefas: %s", err)
	} else {
		t.Log("Teste com sucesso.")
	}
	Controllers.DeletarUsuario(123456789)
	Controllers.DeletarTarefa(123456789)
	Controllers.DeletarTarefa(12345678)

}
