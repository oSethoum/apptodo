package resolvers

import (
	"app/backend/ent"
	"app/backend/graph/models"
	"context"
)


type TodoListenner struct {
	Context context.Context
	ID      int
	Event   models.Event
	Channel chan *ent.Todo
}

type TodosListenner struct {
	Context context.Context
	Channel chan *ent.TodoConnection
	Event   models.Event
	Query   *models.TodosQueryInput
}

