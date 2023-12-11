package repositories

import (
	"database/sql"
	"sync"

	service "github.com/flat35hd99/play-sqlite-go/service"
	_ "github.com/glebarez/go-sqlite"
)

type TaskRepositoryImpl struct {
	lock sync.Mutex
	db   *sql.DB
}

func NewTaskRepository() service.TaskRepository {
	db, err := sql.Open("sqlite", "tasks.db")
	if err != nil {
		panic(err)
	}
	migration(db)

	return &TaskRepositoryImpl{db: db}
}

func migration(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, description TEXT)")
	if err != nil {
		panic(err)
	}
}

func (r *TaskRepositoryImpl) Create(uncommitedTask service.Task) (service.Task, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	var commitedTask service.Task
	err := r.db.QueryRow("INSERT INTO tasks (description) VALUES (?) RETURNING id, description;", uncommitedTask.Description).Scan(&commitedTask.ID, &commitedTask.Description)
	if err != nil {
		return service.Task{}, err
	}

	return commitedTask, nil
}

func (r *TaskRepositoryImpl) Delete(id service.TaskID) (service.Task, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	_, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return service.Task{}, err
	}

	return service.Task{}, nil
}

func (r *TaskRepositoryImpl) Find(id service.TaskID) (service.Task, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	var task service.Task
	err := r.db.QueryRow("SELECT id, description FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Description)
	if err != nil {
		return service.Task{}, err
	}

	return task, nil
}

func (r *TaskRepositoryImpl) Update(task service.Task) (service.Task, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	_, err := r.db.Exec("UPDATE tasks SET description = ? WHERE id = ?", task.Description, task.ID)
	if err != nil {
		return service.Task{}, err
	}

	return task, nil
}
