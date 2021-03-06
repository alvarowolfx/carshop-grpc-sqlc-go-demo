// Code generated by sqlc. DO NOT EDIT.

package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createCarStmt, err = db.PrepareContext(ctx, createCar); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCar: %w", err)
	}
	if q.createOwnerStmt, err = db.PrepareContext(ctx, createOwner); err != nil {
		return nil, fmt.Errorf("error preparing query CreateOwner: %w", err)
	}
	if q.createWorkOrderStmt, err = db.PrepareContext(ctx, createWorkOrder); err != nil {
		return nil, fmt.Errorf("error preparing query CreateWorkOrder: %w", err)
	}
	if q.endWorkOrderStmt, err = db.PrepareContext(ctx, endWorkOrder); err != nil {
		return nil, fmt.Errorf("error preparing query EndWorkOrder: %w", err)
	}
	if q.endWorkOrderServiceExecutionStmt, err = db.PrepareContext(ctx, endWorkOrderServiceExecution); err != nil {
		return nil, fmt.Errorf("error preparing query EndWorkOrderServiceExecution: %w", err)
	}
	if q.getCarByLicensePlateStmt, err = db.PrepareContext(ctx, getCarByLicensePlate); err != nil {
		return nil, fmt.Errorf("error preparing query GetCarByLicensePlate: %w", err)
	}
	if q.getCarsByOwnerStmt, err = db.PrepareContext(ctx, getCarsByOwner); err != nil {
		return nil, fmt.Errorf("error preparing query GetCarsByOwner: %w", err)
	}
	if q.getOwnerByEmailStmt, err = db.PrepareContext(ctx, getOwnerByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetOwnerByEmail: %w", err)
	}
	if q.getOwnerByIDStmt, err = db.PrepareContext(ctx, getOwnerByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetOwnerByID: %w", err)
	}
	if q.getOwnerByNationalIDStmt, err = db.PrepareContext(ctx, getOwnerByNationalID); err != nil {
		return nil, fmt.Errorf("error preparing query GetOwnerByNationalID: %w", err)
	}
	if q.getRunningServicesStmt, err = db.PrepareContext(ctx, getRunningServices); err != nil {
		return nil, fmt.Errorf("error preparing query GetRunningServices: %w", err)
	}
	if q.getRunningWorkOrdersStmt, err = db.PrepareContext(ctx, getRunningWorkOrders); err != nil {
		return nil, fmt.Errorf("error preparing query GetRunningWorkOrders: %w", err)
	}
	if q.registerWorkOrderServiceExecutionStmt, err = db.PrepareContext(ctx, registerWorkOrderServiceExecution); err != nil {
		return nil, fmt.Errorf("error preparing query RegisterWorkOrderServiceExecution: %w", err)
	}
	if q.updateWorkOrderServiceStatusStmt, err = db.PrepareContext(ctx, updateWorkOrderServiceStatus); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateWorkOrderServiceStatus: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createCarStmt != nil {
		if cerr := q.createCarStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCarStmt: %w", cerr)
		}
	}
	if q.createOwnerStmt != nil {
		if cerr := q.createOwnerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createOwnerStmt: %w", cerr)
		}
	}
	if q.createWorkOrderStmt != nil {
		if cerr := q.createWorkOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createWorkOrderStmt: %w", cerr)
		}
	}
	if q.endWorkOrderStmt != nil {
		if cerr := q.endWorkOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing endWorkOrderStmt: %w", cerr)
		}
	}
	if q.endWorkOrderServiceExecutionStmt != nil {
		if cerr := q.endWorkOrderServiceExecutionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing endWorkOrderServiceExecutionStmt: %w", cerr)
		}
	}
	if q.getCarByLicensePlateStmt != nil {
		if cerr := q.getCarByLicensePlateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCarByLicensePlateStmt: %w", cerr)
		}
	}
	if q.getCarsByOwnerStmt != nil {
		if cerr := q.getCarsByOwnerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCarsByOwnerStmt: %w", cerr)
		}
	}
	if q.getOwnerByEmailStmt != nil {
		if cerr := q.getOwnerByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOwnerByEmailStmt: %w", cerr)
		}
	}
	if q.getOwnerByIDStmt != nil {
		if cerr := q.getOwnerByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOwnerByIDStmt: %w", cerr)
		}
	}
	if q.getOwnerByNationalIDStmt != nil {
		if cerr := q.getOwnerByNationalIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOwnerByNationalIDStmt: %w", cerr)
		}
	}
	if q.getRunningServicesStmt != nil {
		if cerr := q.getRunningServicesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRunningServicesStmt: %w", cerr)
		}
	}
	if q.getRunningWorkOrdersStmt != nil {
		if cerr := q.getRunningWorkOrdersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRunningWorkOrdersStmt: %w", cerr)
		}
	}
	if q.registerWorkOrderServiceExecutionStmt != nil {
		if cerr := q.registerWorkOrderServiceExecutionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing registerWorkOrderServiceExecutionStmt: %w", cerr)
		}
	}
	if q.updateWorkOrderServiceStatusStmt != nil {
		if cerr := q.updateWorkOrderServiceStatusStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateWorkOrderServiceStatusStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                    DBTX
	tx                                    *sql.Tx
	createCarStmt                         *sql.Stmt
	createOwnerStmt                       *sql.Stmt
	createWorkOrderStmt                   *sql.Stmt
	endWorkOrderStmt                      *sql.Stmt
	endWorkOrderServiceExecutionStmt      *sql.Stmt
	getCarByLicensePlateStmt              *sql.Stmt
	getCarsByOwnerStmt                    *sql.Stmt
	getOwnerByEmailStmt                   *sql.Stmt
	getOwnerByIDStmt                      *sql.Stmt
	getOwnerByNationalIDStmt              *sql.Stmt
	getRunningServicesStmt                *sql.Stmt
	getRunningWorkOrdersStmt              *sql.Stmt
	registerWorkOrderServiceExecutionStmt *sql.Stmt
	updateWorkOrderServiceStatusStmt      *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                    tx,
		tx:                                    tx,
		createCarStmt:                         q.createCarStmt,
		createOwnerStmt:                       q.createOwnerStmt,
		createWorkOrderStmt:                   q.createWorkOrderStmt,
		endWorkOrderStmt:                      q.endWorkOrderStmt,
		endWorkOrderServiceExecutionStmt:      q.endWorkOrderServiceExecutionStmt,
		getCarByLicensePlateStmt:              q.getCarByLicensePlateStmt,
		getCarsByOwnerStmt:                    q.getCarsByOwnerStmt,
		getOwnerByEmailStmt:                   q.getOwnerByEmailStmt,
		getOwnerByIDStmt:                      q.getOwnerByIDStmt,
		getOwnerByNationalIDStmt:              q.getOwnerByNationalIDStmt,
		getRunningServicesStmt:                q.getRunningServicesStmt,
		getRunningWorkOrdersStmt:              q.getRunningWorkOrdersStmt,
		registerWorkOrderServiceExecutionStmt: q.registerWorkOrderServiceExecutionStmt,
		updateWorkOrderServiceStatusStmt:      q.updateWorkOrderServiceStatusStmt,
	}
}
