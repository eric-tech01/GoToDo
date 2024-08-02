package gateway

import (
	"context"
	"netix-todo-app/backend/global"
	"netix-todo-app/backend/lib/utils"
	todoModel "netix-todo-app/backend/model/todo"
	todoSrv "netix-todo-app/backend/service/todo"

	log "github.com/eric-tech01/simple-log"
)

// App struct
type TodoController struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewTodoController() *TodoController {
	return &TodoController{}
}

func (a *TodoController) Startup(ctx context.Context) {
	a.ctx = ctx
}
func (a *TodoController) Save(t *todoModel.Todo) *todoModel.Todo {
	if t.Id == 0 {
		t.Id = utils.GetTimeStampMillSecond()
	}
	if err := t.Save(); err != nil {
		log.Error(err)
		return nil
	}
	return t
}
func (a *TodoController) Load() todoModel.TodoList {
	return todoSrv.FetchAll()

}
func (a *TodoController) Delete(id int64) string {
	t := todoModel.Todo{Id: id}
	if err := t.Delete(); err != nil {
		log.Error(err)
		return "删除失败"
	}
	return ""

}

func (a *TodoController) Finish(id int64) string {
	t := todoModel.Todo{Id: id}
	if err := t.Get(); err != nil {
		log.Error(err)
		return global.SYSTEM_EXCEPTION
	}
	t.Completed = true
	if err := t.Save(); err != nil {
		log.Error(err)
		return global.SYSTEM_EXCEPTION
	}
	return global.SUCCESS
}
