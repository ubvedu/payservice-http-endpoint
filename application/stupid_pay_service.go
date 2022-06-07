package application

import (
	"fmt"
	"http-endpoint/application/contracts/dto"
	"net/http"
	"strconv"
)

type StupidPayService struct {
	count int
}

func (s *StupidPayService) Charge(request dto.ChargeRequest) dto.ChargeResult {
	s.count++
	result := dto.ChargeResult{StatusCode: http.StatusOK, Status: "Ok", Uuid: strconv.Itoa(s.count)}
	fmt.Printf("%+v\n%+v\n\n", request, result)
	return result
}

func NewStupidPayService() *StupidPayService {
	return &StupidPayService{}
}
