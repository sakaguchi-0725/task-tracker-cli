package persistence

import (
	"github.com/sakaguchi-0725/task-tracker/internal/domain/model"
	"github.com/sakaguchi-0725/task-tracker/internal/domain/repository"
	"github.com/sakaguchi-0725/task-tracker/internal/infra/dao"
	"github.com/sakaguchi-0725/task-tracker/internal/infra/dto"
)

type taskPersistence struct {
	dao dao.JsonDAO[dto.Task]
}

func (t *taskPersistence) Create(task model.Task) error {
	taskDTO := dto.NewTaskDTO(task)

	if err := t.dao.Create(&taskDTO); err != nil {
		return err
	}

	return nil
}

func (t *taskPersistence) DeleteByID(id model.TaskID) error {
	err := t.dao.Where("ID", id.String()).Delete()
	if err != nil {
		return err
	}

	return nil
}

func (t *taskPersistence) FindAll() ([]model.Task, error) {
	var tasks []dto.Task
	if err := t.dao.Find(&tasks); err != nil {
		return []model.Task{}, err
	}

	res := make([]model.Task, len(tasks))
	for i, task := range tasks {
		res[i] = model.RecreateTask(
			model.TaskID(task.ID),
			task.Title,
			task.Description,
			model.Status(task.Status),
		)
	}

	return res, nil
}

func (t *taskPersistence) FindByID(id model.TaskID) (model.Task, error) {
	var task dto.Task
	if err := t.dao.Where("ID", id.String()).First(&task); err != nil {
		return model.Task{}, err
	}

	return model.RecreateTask(
		id,
		task.Title,
		task.Description,
		model.Status(task.Status),
	), nil
}

func (t *taskPersistence) Update(task model.Task) error {
	taskDTO := dto.NewTaskDTO(task)
	if err := t.dao.Update(&taskDTO); err != nil {
		return err
	}

	return nil
}

func NewTaskPersistence(filePath string) repository.TaskRepository {
	dao := dao.NewJsonDAO[dto.Task](filePath)

	return &taskPersistence{dao}
}
