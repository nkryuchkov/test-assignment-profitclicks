package service

type operationFunc func([]int64) int64

func (s *Service) isAllowedOperation(operation string) bool {
	return s.mapOperationToFunc(operation) != nil
}

func (s *Service) mapOperationToFunc(operation string) operationFunc {
	switch operation {
	case "sum":
		return s.sumOperation
	case "count":
		return s.countOperation
	default:
		return nil
	}
}

func (s *Service) sumOperation(list []int64) int64 {
	var sum int64
	for _, v := range list {
		sum += v
	}
	return sum
}

func (s *Service) countOperation(list []int64) int64 {
	return int64(len(list))
}
