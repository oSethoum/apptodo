package resolvers

import (
   "app/backend/graph/models"
   "app/backend/ent"
)


func NotifyTodosListenners(r *mutationResolver, event models.Event) {
	r.TodosListennersMutext.Lock()
	for key := range r.TodosListenners {
		if event == r.TodosListenners[key].Event {
			query := r.TodosListenners[key].Query
			if query == nil {
				query = &models.TodosQueryInput{}
			}

			entities, err := r.Client.Todo.Query().Paginate(r.TodosListenners[key].Context, query.After, query.First, query.Before, query.Last, ent.WithTodoFilter(query.Where.Filter), ent.WithTodoOrder(query.OrderBy))

			if err == nil {
				r.TodosListenners[key].Channel <- entities
			}

		}
	}
	r.TodosListennersMutext.Unlock()
}

func NotifyTodoListenners(r *mutationResolver, event models.Event, entity *ent.Todo) {
	r.TodoListennersMutext.Lock()
	for key := range r.TodoListenners {
		if event == r.TodoListenners[key].Event && r.TodoListenners[key].ID == entity.ID {
			r.TodoListenners[key].Channel <- entity
		}
	}
	r.TodoListennersMutext.Unlock()
}


