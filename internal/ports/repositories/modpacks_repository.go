package repositories

import "my-stats/internal/domain"

type ModpacksRepository interface {
	Get() map[string]domain.Modpack
	StopTicker()
}
