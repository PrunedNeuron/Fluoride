package controller

import (
	"fluoride/internal/model"
	"fluoride/pkg/errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

// CreatePlan creates a new plan
func CreatePlan(w http.ResponseWriter, r *http.Request) {

	plan := new(model.Plan)

	err := render.DecodeJSON(r.Body, &plan)

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(err))
		return
	}

	if plan.Name == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Missing name value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Missing name value")))
		return
	}

	if plan.Description == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Missing description value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Missing description value")))
		return
	}

	if plan.Intro == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Missing introduction value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Missing introduction value")))
		return
	}

	price, _ := strconv.ParseFloat(plan.Price, 32)

	if plan.Price == "" || price <= 0.0 {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Missing or invalid price value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Missing or invalid price value")))
		return
	}

	if plan.BillingCycle <= 0 {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Missing billing period value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Missing billing period value")))
		return
	}

	planName, err := planService.CreatePlan(plan)

	render.JSON(w, r, &response{
		Status:  "success",
		Message: "Created plan with name " + planName,
	})
}

// GetPlans gets all the plans
func GetPlans(w http.ResponseWriter, r *http.Request) {
	plans, err := planService.GetPlans()

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, &response{
		Status: "success",
		Plans:  plans,
	})
}
