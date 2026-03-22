package persistence

import (
	"my-stats/internal/ports/repositories"
	"sync"
	"time"
)

type MemoryProjectsRepository struct {
	data       map[string]int
	ticker     *time.Ticker
	updateFunc func(data map[string]int)
	mu         sync.RWMutex
	done       chan struct{}
}

func NewMemoryProjectsRepository(updateFunc func(data map[string]int)) repositories.ProjectsRepository {
	repo := &MemoryProjectsRepository{
		data:       make(map[string]int),
		updateFunc: updateFunc,
		done:       make(chan struct{}),
	}

	repo.ticker = time.NewTicker(1 * time.Minute)

	go repo.startTicker()

	return repo
}

func (m *MemoryProjectsRepository) startTicker() {
	for {
		select {
		case <-m.ticker.C:
			m.mu.Lock()
			m.updateFunc(m.data)
			m.mu.Unlock()
		case <-m.done:
			return
		}
	}
}

func (m *MemoryProjectsRepository) Get() map[string]int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make(map[string]int, len(m.data))
	for k, v := range m.data {
		result[k] = v
	}

	return result
}

func (m *MemoryProjectsRepository) StopTicker() {
	close(m.done)
	m.ticker.Stop()
}
