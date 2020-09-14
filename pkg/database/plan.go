package database

import (
	"database/sql"
	"fluoride/internal/model"
	"fluoride/pkg/errors"
	"fmt"

	_ "github.com/jackc/pgx/stdlib" // For the pg driver
	"go.uber.org/zap"
)

// PlansExists checks whether the plans table exists
func (dbc *DBClient) PlansExists() (bool, error) {
	zap.S().Debugw("Querying database to check whether plans table exists")
	_, err := dbc.db.Queryx("SELECT 'core.plans'::regclass")

	if err != nil {
		zap.S().Debugw("Developer icon packs table does not exist")
		return false, err
	}

	return true, nil
}

// CreatePlan creates a new plan record
func (dbc *DBClient) CreatePlan(plan model.Plan) (string, error) {

	zap.S().Debugw("Make sure plans table exists before attempting to insert into it")
	if plansExists, _ := dbc.PlansExists(); !plansExists {
		zap.S().Errorf("Plans table does not exist, cannot create new plan")
		return "", fmt.Errorf("Plans table does not exist, cannot create new plan")
	}
	query := `
			INSERT INTO core.plans (name, description, introduction, price, billing_cycle)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING name
		`

	zap.S().Debugw("Inserting icon pack into developer's icon packs table")
	row := dbc.db.QueryRowx(query, plan.Name, plan.Description, plan.Intro, plan.Price, plan.BillingCycle)

	zap.S().Debugw("Scanning returned row")
	var planName string
	err := row.Scan(&planName)

	if err != nil {
		zap.S().Errorf("Failed to scan return value, error: %s", err)
		return "", err
	}

	zap.S().Debugw("Returning with the name of the new plan")
	return planName, nil

}

// GetAllPlans gets all the plans
func (dbc *DBClient) GetAllPlans() ([]model.Plan, error) {
	zap.S().Debugw("Make sure plans table exists before attempting to select from it")
	if plansExists, _ := dbc.PlansExists(); !plansExists {
		zap.S().Errorf("Plans table does not exist, cannot retrieve plans")
		return nil, fmt.Errorf("Plans table does not exist, cannot retrieve plans")
	}
	plans := []model.Plan{}
	zap.S().Debugw("Querying the database for all plans")
	rows, err := dbc.db.Queryx(`
		SELECT * FROM core.plans
	`)
	zap.S().Debugw("Scanning the result")
	for rows.Next() {
		var plan model.Plan
		err = rows.StructScan(&plan)
		plans = append(plans, plan)
	}
	if err == sql.ErrNoRows {
		zap.S().Errorf("No rows in the database!")
		return nil, err
	} else if err != nil {
		zap.S().Errorf(errors.ErrDatabase.Error())
		return nil, err
	}

	zap.S().Debugw("Returning with the list of all plans")
	return plans, nil
}
