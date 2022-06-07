package contracts

import "http-endpoint/application/contracts/dto"

type PayService interface {
	Charge(r dto.ChargeRequest) dto.ChargeResult
}
