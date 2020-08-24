package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	Controllers "github.com/renancrc99/1sti.challenge.backend/controllers"
	"github.com/renancrc99/1sti.challenge.backend/graph/generated"
	"github.com/renancrc99/1sti.challenge.backend/graph/model"
)

func (r *mutationResolver) CriarUsuario(ctx context.Context, input *model.UsuarioInput) (*model.Usuario, error) {
	err := Controllers.CriarUsuario(input)
	return nil, err
}

func (r *mutationResolver) CriarTarefaPorUsuario(ctx context.Context, input model.TarefaInput) (*model.Tarefa, error) {
	err := Controllers.CriarTarefaPorUsuario(input)
	return nil, err
}

func (r *mutationResolver) EditarTarefaPorUsuario(ctx context.Context, input model.TarefaInput) (*model.Tarefa, error) {
	err := Controllers.EditarTarefaPorUsuario(input)
	return nil, err
}

func (r *mutationResolver) MudarStatusTarefa(ctx context.Context, input model.TarefaInput) (*model.Tarefa, error) {
	err := Controllers.MudarStatusTarefa(input)
	return nil, err
}

func (r *queryResolver) ListarTarefasUsuario(ctx context.Context, input model.UsuarioInput) ([]*model.Tarefa, error) {
	tarefas := Controllers.ListarTarefasUsuario(input)
	var retorno []*model.Tarefa
	for i := 0; i < len(tarefas); i++ {
		var nova = new(model.Tarefa)
		nova.ID = tarefas[i].ID
		nova.Feito = tarefas[i].Feito
		nova.Text = tarefas[i].Text
		nova.Status = tarefas[i].Status
		retorno = append(retorno, nova)

	}
	return retorno, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
