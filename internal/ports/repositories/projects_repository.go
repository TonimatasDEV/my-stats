package repositories

type ProjectsRepository interface {
	Get() map[string]int
	StopTicker()
}
