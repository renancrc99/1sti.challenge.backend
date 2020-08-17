package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/renancrc99/1sti.challenge.backend/graph/generated"
	"github.com/renancrc99/1sti.challenge.backend/graph/model"
)

func (r *mutationResolver) CriarUsuario(ctx context.Context, input *model.UsuarioInput) (*model.Usuario, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CriarTarefaPorUsuario(ctx context.Context, input model.TarefaInput) (*model.Tarefa, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditarTarefaPorUsuario(ctx context.Context, input model.TarefaInput) (*model.Tarefa, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MudarStatusTarefa(ctx context.Context, input model.TarefaInput) (*model.Tarefa, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListarTarefasUsuario(ctx context.Context, input model.UsuarioInput) ([]*model.Tarefa, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
