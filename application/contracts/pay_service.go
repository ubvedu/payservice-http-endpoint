package contracts

import "payervice-http-endpoint/application/contracts/dto"

type PayService interface {
	Charge(r dto.ChargeRequest) dto.ChargeResult
}
