package service

type TaskID int

type Task struct {
	ID          TaskID
	Description string
}

type TaskRepository interface {
	Create(Task) (Task, error)
	Delete(TaskID) (Task, error)
	Find(TaskID) (Task, error)
	Update(Task) (Task, error)
}

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Create(description string) (Task, error) {
	return s.repo.Create(Task{Description: description})
}

func (s *TaskService) Delete(id TaskID) (Task, error) {
	return s.repo.Delete(id)
}

func (s *TaskService) Find(id TaskID) (Task, error) {
	return s.repo.Find(id)
}

func (s *TaskService) Update(task Task) (Task, error) {
	return s.repo.Update(task)
}
