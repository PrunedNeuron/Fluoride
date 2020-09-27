package controller

import (
	"fmt"
	"net/http"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"github.com/PrunedNeuron/Fluoride/pkg/model"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// GetDevs renders all the devs in the database
// * GET /developers
// swagger:operation GET /developers Users GetDevs
//
// Gets the list of all the developers present in the database.
// ---
// produces:
//     - application/json
//
// responses:
//     "200":
//         description: OK
//         schema:
//             type: object
//             properties:
//                     status:
//                         type: string
//                         description: status message
//                         example: success
//                     developers:
//                         type: array
//                         items:
//                             "$ref": "#/definitions/User"
func GetDevs(w http.ResponseWriter, r *http.Request) {
	devs, err := devService.GetDevs()

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, &response{
		Status: "success",
		Devs:   devs,
	})
}

// GetDevCount renders the number of devs in the database
// * GET /developers/count
// swagger:operation GET /developers/count Users GetDevCount
//
// Get the total number of developers
//
// Fetches the total count of developers in the database
// ---
// produces:
//     - application/json
//
// responses:
//     "200":
//         description: OK
//         schema:
//             type: object
//             properties:
//                 status:
//                     type: string
//                     description: status message
//                 count:
//                     type: integer
//                     description: number of developers
func GetDevCount(w http.ResponseWriter, r *http.Request) {

	count, err := devService.GetDevCount()

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, &response{
		Status: "success",
		Count:  count,
	})
}

// GetDevByUsername renders the dev with the given username
// * GET /developers/{developer}
// swagger:operation GET /developers/{developer} Users GetDevByUsername
//
// Get a developer by username
//
// Fetches the developer with the provided username
// ---
// produces:
//     - application/json
//
// parameters:
//     - name: developer
//       in: path
//       description: developer username
//       required: true
//       type: string
//
// responses:
//     "200":
//         description: OK
//         schema:
//             type: object
//             properties:
//                     status:
//                         type: string
//                         description: status message
//                         example: success
//                     developers:
//                         type: array
//                         items:
//                             "$ref": "#/definitions/User"
//     "500":
//         description: bad request
//         schema:
//             type: object
//             properties:
//                     status:
//                         type: string
//                         description: status message
//                         example: failure
//                     message:
//                         type: string
//                         description: informational message
//                         example: server error
//                     error:
//                         type: string
//                         description: error message
//                         example: invalid dev
func GetDevByUsername(w http.ResponseWriter, r *http.Request) {

	// Get dev from url
	username := chi.URLParam(r, "developer")

	if username == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid dev")))
		return
	}

	dev, err := devService.GetDevByUsername(username)

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	var devs []model.User
	devs = append(devs, dev)

	render.JSON(w, r, &response{
		Status: "success",
		Devs:   devs,
	})
}

// * POST /developers
// ! TODO: Alias for POST /users with role = 'developer'
// swagger:operation POST /developers Users SaveDev
//
// Add a developer
//
// Saves a new developer to the database, rejecting the request on conflict
// ---
// consumes:
//     - application/json
//
// produces:
//     - application/json
//
// security:
//     - api_key: []
//
// parameters:
//     - name: request
//       in: body
//       description: information about the new developer
//       required: true
//       schema:
//           type: object
//           properties:
//              role:
//                  type: string
//                  description: role of the user (developer | admin)
//                  example: developer
//              name:
//                  type: string
//                  description: name of the developer
//                  example: Ayush Mishra
//              username:
//                  type: string
//                  description: username of the developer
//                  example: ayush
//              email:
//                  type: string
//                  description: email address of the developer
//                  example: am@ayushm.dev
//              url:
//                  type: string
//                  description: developer website URL
//                  example: https://ayushm.dev
//
// responses:
//     "200":
//         description: OK
//         schema:
//             type: object
//             properties:
//                 status:
//                     type: string
//                     description: status message
//                     example: success
//                 message:
//                     type: string
//                     description: informational message
//                     example: successfully added developer
//
//     "500":
//         description: server error
//         schema:
//             type: object
//             properties:
//                 status:
//                     type: string
//                     description: status message
//                     example: failure
//                 message:
//                     type: string
//                     description: informational message
//                     example: invalid request
//                 error:
//                     type: string
//                     description: error message
//                     example: user may already exist
