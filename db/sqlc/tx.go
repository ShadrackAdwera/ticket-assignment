package db

import (
	"context"
	"database/sql"
)

type TxStore interface {
	Querier
	NewTicketAssignment(ctx context.Context, args CreateAssignmentParams) (NewTicketAssignmentResult, error)
}

type TxSqlStore struct {
	*Queries
	db *sql.DB
}

func NewTxStore(db *sql.DB) TxStore {
	return &TxSqlStore{
		db:      db,
		Queries: New(db),
	}
}

// executes a function within a DB transancion
func (txStore *TxSqlStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := txStore.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})

	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		if rollbaccErr := tx.Rollback(); rollbaccErr != nil {
			return rollbaccErr
		}
		return err
	}

	return tx.Commit()
}

/*
THE TRANSACTION
- Assign ticket to agent
- Change Agent status from ACTIVE to INACTIVE
- Add a record to assignment table
*/

type NewTicketAssignmentResult struct {
	Ticket     Ticket     `json:"ticket"`
	Agent      Agent      `json:"agent"`
	Assignment Assignment `json:"assignment"`
}

func (txStore *TxSqlStore) NewTicketAssignment(ctx context.Context, args CreateAssignmentParams) (NewTicketAssignmentResult, error) {
	var ticketAssignmentResult NewTicketAssignmentResult

	err := txStore.execTx(ctx, func(q *Queries) error {
		var err error
		ticketAssignmentResult.Ticket, err = q.AssignTicketToAgent(context.Background(), AssignTicketToAgentParams{
			ID: args.TicketID,
			AssignedTo: sql.NullInt64{
				Int64: args.AgentID,
				Valid: true,
			},
			Status: "ASSIGNED",
		})

		if err != nil {
			return err
		}

		ticketAssignmentResult.Agent, err = q.UpdateAgent(context.Background(), UpdateAgentParams{
			ID:     args.AgentID,
			Status: "ACTIVE",
		})

		if err != nil {
			return err
		}
		ticketAssignmentResult.Assignment, err = q.CreateAssignment(context.Background(), CreateAssignmentParams{
			TicketID: args.TicketID,
			AgentID:  args.AgentID,
			Status:   "ASSIGNED",
		})

		if err != nil {
			return err
		}

		// lock ticket - lock agent

		return err
	})

	return ticketAssignmentResult, err
}
