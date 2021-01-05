package validations

import (
	"com.aviebrantz.carshop/pkg/repository"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CheckUniqueConstraintError(err error) error {
	if pgerr, ok := err.(*pq.Error); ok {
		if pgerr.Code.Name() == "unique_violation" {
			return status.Errorf(codes.AlreadyExists, "unique constraint validation - %s", pgerr.Constraint)
		}
	}
	return err
}

func FromErrorResponsesToGrpcError(msg string, errors []*ErrorResponse) error {
	st := status.New(codes.InvalidArgument, msg)
	br := &errdetails.BadRequest{}
	for _, err := range errors {
		fv := &errdetails.BadRequest_FieldViolation{
			Field:       err.FailedField,
			Description: err.Error,
		}
		br.FieldViolations = append(br.FieldViolations, fv)
	}
	st, err := st.WithDetails(br)
	if err != nil {
		return status.Error(codes.Internal, "Failed to build error msg")
	}
	return st.Err()
}

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
