package Controllers

import (
	"errors"

	Config "github.com/renancrc99/1sti.challenge.backend/config"
	"github.com/renancrc99/1sti.challenge.backend/graph/model"

	Models "github.com/renancrc99/1sti.challenge.backend/models"
)

func CriarUsuario(input *model.UsuarioInput) error {
	usuario := new(Models.Usuario)
	if input.Nome == "" {
		return errors.New("Campo nome é obrigatório")
	}
	if input.Email == "" {
		return errors.New("Campo email é obrigatório")
	}
	usuario.ID = input.ID
	usuario.Nome = input.Nome
	usuario.Email = input.Email

	if err := Config.DB.Create(usuario).Error; err != nil {
		return err
	}
	return nil
}

func CriarTarefaPorUsuario(tarefaInput model.TarefaInput) error {
	tarefa := new(Models.Tarefa)
	tarefa.IDUsuario = tarefaInput.UsuarioInput.ID
	tarefa.ID = tarefaInput.ID
	tarefa.Text = tarefaInput.Text
	tarefa.Feito = tarefaInput.Feito
	tarefa.Status = tarefaInput.Status
	if err := Config.DB.Create(tarefa).Error; err != nil {
		return err
	}
	return nil
}

func EditarTarefaPorUsuario(tarefaInput model.TarefaInput) error {
	var tarefa Models.Tarefa
	Config.DB.Where("id = ?", tarefaInput.ID).Find(&tarefa)

	tarefa.Text = tarefaInput.Text
	tarefa.Feito = tarefaInput.Feito
	tarefa.Status = tarefaInput.Status
	if err := Config.DB.Save(tarefa).Error; err != nil {
		return err
	}
	return nil
}

func MudarStatusTarefa(tarefaInput model.TarefaInput) error {
	var tarefa Models.Tarefa
	Config.DB.Where("id = ?", tarefaInput.ID).Find(&tarefa)
	tarefa.Status = tarefaInput.Status
	if err := Config.DB.Save(&tarefa).Error; err != nil {
		return err
	}
	return nil
}

func ListarTarefasUsuario(input model.UsuarioInput) []Models.Tarefa {
	var tarefa []Models.Tarefa
	Config.DB.Where("id_usuario = ?", input.ID).Find(&tarefa)
	return tarefa

}

func DeletarUsuario(id int) error {
	var usuario Models.Usuario
	Config.DB.Where("id = ?", id).Find(&usuario)
	if err := Config.DB.Delete(usuario).Error; err != nil {
		return err
	}
	return nil
}

func DeletarTarefa(id int) error {
	var tarefa Models.Tarefa
	Config.DB.Where("id = ?", id).Find(&tarefa)
	if err := Config.DB.Delete(tarefa).Error; err != nil {
		return err
	}
	return nil
}
