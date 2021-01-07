package domain

import "com.aviebrantz.carshop/pkg/common/repository"

func FromServiceTypeToWorkStatus(t repository.ServiceType) repository.WorkOrderStatus {
	switch t {
	case repository.ServiceTypeCHANGEPARTS:
		return repository.WorkOrderStatusCHANGINGPARTS
	case repository.ServiceTypeCHANGETIRES:
		return repository.WorkOrderStatusCHANGINGTIRES
	case repository.ServiceTypeDIAGNOSTIC:
		return repository.WorkOrderStatusDIAGNOSTICS
	case repository.ServiceTypeWASH:
		return repository.WorkOrderStatusWASHING
	default:
		return repository.WorkOrderStatusIDLE
	}
}
