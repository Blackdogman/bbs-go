package services

var ApexSessionService = newApexSessionService()

func newApexSessionService() *apexSessionService {
	return &apexSessionService{}
}

type apexSessionService struct {
}
