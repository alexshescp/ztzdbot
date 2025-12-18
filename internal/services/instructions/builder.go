package instructions

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetUniversalProtocol() string {
	return "1) Изоляция\n2) Сохранение артефактов\n3) Диагностика\n4) Минимизация ущерба\n5) Контакт инженера"
}
