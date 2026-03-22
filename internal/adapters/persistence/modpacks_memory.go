package persistence

import (
	"my-stats/internal/domain"
	"my-stats/internal/ports/repositories"
	"sync"
	"time"
)

type MemoryModpacksRepository struct {
	data       map[string]domain.Modpack
	ticker     *time.Ticker
	updateFunc func(data map[string]domain.Modpack)
	mu         sync.RWMutex
	done       chan struct{}
}

func NewMemoryModpacksRepository(updateFunc func(data map[string]domain.Modpack)) repositories.ModpacksRepository {
	repo := &MemoryModpacksRepository{
		data:       make(map[string]domain.Modpack),
		updateFunc: updateFunc,
		done:       make(chan struct{}),
	}

	updateFunc(repo.data)
	repo.ticker = time.NewTicker(time.Minute)

	go repo.startTicker()

	return repo
}

func (m *MemoryModpacksRepository) startTicker() {
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

func (m *MemoryModpacksRepository) Get() map[string]domain.Modpack {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make(map[string]domain.Modpack, len(m.data))
	for k, v := range m.data {
		result[k] = v
	}

	return result
}

func (m *MemoryModpacksRepository) StopTicker() {
	close(m.done)
	m.ticker.Stop()
}
