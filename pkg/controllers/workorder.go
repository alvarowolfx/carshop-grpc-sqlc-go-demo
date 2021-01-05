package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	carshop "com.aviebrantz.carshop/api"
	"com.aviebrantz.carshop/pkg/repository"
	"com.aviebrantz.carshop/pkg/validations"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WorkOrderController interface {
	carshop.WorkOrderServiceServer
}

type WorkOrderControllerDeps struct {
	DB repository.Querier
}

type workOrderController struct {
	deps    WorkOrderControllerDeps
	client  *http.Client
	encoder *jsonpb.Marshaler
}

// NewWorkOrderController
func NewWorkOrderController(deps WorkOrderControllerDeps) WorkOrderController {
	return &workOrderController{
		deps: deps,
	}
}

func (woc *workOrderController) RegisterWorkOrder(ctx context.Context, params *carshop.WorkOrderRequest) (*empty.Empty, error) {

	car, err := woc.deps.DB.GetCarByLicensePlate(ctx, params.LicensePlate)
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.FailedPrecondition, "Car not found")
	}

	if err != nil {
		return nil, err
	}

	workOrders, err := woc.deps.DB.GetRunningWorkOrders(ctx)

	for _, wo := range workOrders {
		if wo.LicensePlate == params.LicensePlate {
			return nil, status.Error(codes.FailedPrecondition, "There is already an open Work Order for this car")
		}
	}

	_, err = woc.deps.DB.CreateWorkOrder(ctx, repository.CreateWorkOrderParams{
		ChangeTires: params.ChangeTires,
		ChangeParts: params.ChangeParts,
		CarID:       int64(car.ID),
	})

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (woc *workOrderController) GetRunningWorkOrders(ctx context.Context, query *carshop.RunningWorkOrdersQuery) (*carshop.RunningWorkOrdersResponse, error) {
	workOrders, err := woc.deps.DB.GetRunningWorkOrders(ctx)

	if err == sql.ErrNoRows {
		res := &carshop.RunningWorkOrdersResponse{
			WorkOrder: []*carshop.WorkOrder{},
		}
		return res, nil
	}

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Failed to fetch work orders: %v", err))
	}

	res := &carshop.RunningWorkOrdersResponse{
		WorkOrder: []*carshop.WorkOrder{},
	}

	for _, wo := range workOrders {
		ps := carshop.WorkOrderStatus_value[string(wo.PreviousStatus)]
		cs := carshop.WorkOrderStatus_value[string(wo.CurrentStatus)]
		res.WorkOrder = append(res.WorkOrder, &carshop.WorkOrder{
			Id:             int64(wo.ID),
			LicensePlate:   wo.LicensePlate,
			ChangeTires:    wo.ChangeTires,
			ChangeParts:    wo.ChangeParts,
			Status:         carshop.WorkOrderStatus(cs),
			PreviousStatus: carshop.WorkOrderStatus(ps),
		})
	}

	return res, nil
}

func (woc *workOrderController) StartWorkOrderService(ctx context.Context, params *carshop.StartWorkOrderServiceRequest) (*empty.Empty, error) {
	services, err := woc.deps.DB.GetRunningServices(ctx, params.WorkOrderNumber)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Failed to check running services: %v", err))
	}

	if len(services) > 0 {
		return nil, status.Error(codes.FailedPrecondition, "There are services already being executed: %v")
	}

	serviceType := repository.ServiceType(params.Type.String())
	_, err = woc.deps.DB.RegisterWorkOrderServiceExecution(ctx, repository.RegisterWorkOrderServiceExecutionParams{
		WorkOrderID: params.WorkOrderNumber,
		Type:        serviceType,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Failed to register work order execution: %v", err))
	}

	err = woc.deps.DB.UpdateWorkOrderServiceStatus(ctx, repository.UpdateWorkOrderServiceStatusParams{
		ID:            int32(params.WorkOrderNumber),
		CurrentStatus: validations.FromServiceTypeToWorkStatus(serviceType),
	})

	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Failed to update work order: %v", err))
	}

	return &empty.Empty{}, nil
}

func (woc *workOrderController) FinishWorkOrderService(ctx context.Context, params *carshop.FinishWorkOrderServiceRequest) (*empty.Empty, error) {
	services, err := woc.deps.DB.GetRunningServices(ctx, params.WorkOrderNumber)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Failed to check running services: %v", err))
	}

	if len(services) <= 0 {
		return nil, status.Error(codes.FailedPrecondition, "There is no service being executed")
	}

	err = woc.deps.DB.EndWorkOrderServiceExecution(ctx, params.WorkOrderNumber)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Failed to register work order execution: %v", err))
	}

	err = woc.deps.DB.UpdateWorkOrderServiceStatus(ctx, repository.UpdateWorkOrderServiceStatusParams{
		ID:            int32(params.WorkOrderNumber),
		CurrentStatus: repository.WorkOrderStatusIDLE,
	})

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Failed to update work order: %v", err))
	}

	return &empty.Empty{}, nil
}

func (woc *workOrderController) EndWorkOrder(ctx context.Context, params *carshop.EndWorkOrderRequest) (*empty.Empty, error) {
	services, err := woc.deps.DB.GetRunningServices(ctx, params.WorkOrderNumber)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Failed to check running services: %v", err))
	}

	if len(services) > 0 {
		return nil, status.Error(codes.FailedPrecondition, "There are services being executed")
	}

	err = woc.deps.DB.UpdateWorkOrderServiceStatus(ctx, repository.UpdateWorkOrderServiceStatusParams{
		ID:            int32(params.WorkOrderNumber),
		CurrentStatus: repository.WorkOrderStatusDONE,
	})

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Failed to update work order: %v", err))
	}

	return &empty.Empty{}, nil
}
