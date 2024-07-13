package main

import (
	"fmt"
	"sync"
)

type Storage struct {
	m     sync.Mutex
	Tasks map[string]Task //storing as map like in DB
	Users map[string]User
}

func NewStorage() *Storage {
	return &Storage{
		Tasks: make(map[string]Task),
		Users: make(map[string]User),
	}
}

func (s *Storage) GetTasks() []Task {
	tasks := make([]Task, 0, len(s.Tasks)) //turning map into slice to return an array. why can't we also return a map?

	for _, task := range s.Tasks {
		tasks = append(tasks, task)
	}

	return tasks
}

func (s *Storage) CreateTask(task Task) (string, bool) {
	s.m.Lock()

	defer s.m.Unlock()

	id, err := gonanoid.New()

	if err != nil {
		return "", false
	}

	task.ID = id
	s.Tasks[task.ID] = task

	fmt.Print(task)
	return id, true
}

func (s *Storage) UpdadeTask(id string, task Task) bool {
	s.m.Lock()

	defer s.m.Unlock()

	t, ok := s.Tasks[id]

	if !ok {
		return false
	}

	task.ID = t.ID
	s.Tasks[task.ID] = task

	return true
}

func (s *Storage) DeleteTask(id string) bool {
	s.m.Lock()

	defer s.m.Unlock()

	_, ok := s.Tasks[id]

	if !ok {
		return false
	}

	delete(s.Tasks, id)
	return true
}

func (s *Storage) GetUser(username string) (User, bool) {
	fmt.Print(username)
	s.m.Lock()

	defer s.m.Unlock()

	user, ok := s.Users[username]

	fmt.Print(user)
	return user, ok
}

func (s *Storage) CreateUser(u User) bool {
	s.m.Lock()

	defer s.m.Unlock()

	_, ok := s.Users[u.Username]

	if ok {
		return false
	}

	s.Users[u.Username] = u
	return true
}
