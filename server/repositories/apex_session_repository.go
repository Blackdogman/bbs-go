package repositories

var ApexSessionRepository = newApexSessionRepository()

func newApexSessionRepository() *apexSessionRepository {
	return &apexSessionRepository{}
}

type apexSessionRepository struct {
}
