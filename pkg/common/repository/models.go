// Code generated by sqlc. DO NOT EDIT.

package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type CarSize string

const (
	CarSizeSmall  CarSize = "small"
	CarSizeMedium CarSize = "medium"
	CarSizeLarge  CarSize = "large"
)

func (e *CarSize) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CarSize(s)
	case string:
		*e = CarSize(s)
	default:
		return fmt.Errorf("unsupported scan type for CarSize: %T", src)
	}
	return nil
}

type ServiceType string

const (
	ServiceTypeDIAGNOSTIC  ServiceType = "DIAGNOSTIC"
	ServiceTypeCHANGEPARTS ServiceType = "CHANGE_PARTS"
	ServiceTypeCHANGETIRES ServiceType = "CHANGE_TIRES"
	ServiceTypeWASH        ServiceType = "WASH"
)

func (e *ServiceType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ServiceType(s)
	case string:
		*e = ServiceType(s)
	default:
		return fmt.Errorf("unsupported scan type for ServiceType: %T", src)
	}
	return nil
}

type WorkOrderStatus string

const (
	WorkOrderStatusCREATED       WorkOrderStatus = "CREATED"
	WorkOrderStatusDIAGNOSTICS   WorkOrderStatus = "DIAGNOSTICS"
	WorkOrderStatusCHANGINGPARTS WorkOrderStatus = "CHANGING_PARTS"
	WorkOrderStatusCHANGINGTIRES WorkOrderStatus = "CHANGING_TIRES"
	WorkOrderStatusWASHING       WorkOrderStatus = "WASHING"
	WorkOrderStatusIDLE          WorkOrderStatus = "IDLE"
	WorkOrderStatusFINISHED      WorkOrderStatus = "FINISHED"
	WorkOrderStatusDONE          WorkOrderStatus = "DONE"
)

func (e *WorkOrderStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = WorkOrderStatus(s)
	case string:
		*e = WorkOrderStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for WorkOrderStatus: %T", src)
	}
	return nil
}

type Car struct {
	ID           int32     `json:"id"`
	LicensePlate string    `json:"license_plate"`
	Size         CarSize   `json:"size"`
	NumWheels    int16     `json:"num_wheels"`
	Color        string    `json:"color"`
	OwnerID      int64     `json:"owner_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Owner struct {
	ID         int32     `json:"id"`
	Email      string    `json:"email"`
	NationalID string    `json:"national_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type WorkOrder struct {
	ID             int32           `json:"id"`
	ChangeTires    bool            `json:"change_tires"`
	ChangeParts    bool            `json:"change_parts"`
	CurrentStatus  WorkOrderStatus `json:"current_status"`
	PreviousStatus WorkOrderStatus `json:"previous_status"`
	CarID          int64           `json:"car_id"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

type WorkOrderServiceExecution struct {
	ID          int32        `json:"id"`
	Type        ServiceType  `json:"type"`
	WorkOrderID int64        `json:"work_order_id"`
	CreatedAt   time.Time    `json:"created_at"`
	FinishedAt  sql.NullTime `json:"finished_at"`
}
