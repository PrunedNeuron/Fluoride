package swagger

// ! /developers

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

// * GET /developers/{developer}

// swagger:operation GET /developers/{developer} Users GetDevByUsername
//
// Gets the developer with the given username.
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

// * POST /developers

// swagger:operation POST /developers Users SaveDev
//
// Adds the given developer to the database.
// ---
// consumes:
//     - application/json
//
// produces:
//     - application/json
//
// parameters:
//     - name: request
//       in: body
//       description: information about the new developer
//       required: true
//       schema:
//           "$ref": "#/definitions/DevRequest"
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

// * GET /developers/count

// swagger:operation GET /developers/count Users GetDevCount
//
// Gets the number of developers in the database
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

// ! /packs

// * GET /packs

// swagger:operation GET /packs IconPacks GetPacks
//
// Gets the list of all the icon packs present in the database.
// ---
// produces:
//     - application/json
//
// responses:
//     "200":
//         description: OK
//         schema:
//             type: array
//             items:
//                 "$ref": "#/definitions/PackResponse"

// * POST /developers/{developer}/packs

// swagger:operation POST /developers/{developer}/packs IconPacks SavePack
//
// Adds the given icon pack to the database.
// ---
// consumes:
//     - application/json
// produces:
//     - application/json
//
// parameters:
//     - name: developer
//       in: path
//       description: username of the developer who owns the icon pack
//       required: true
//       type: string
//     - name: request
//       in: body
//       description: information about the new icon pack
//       required: true
//       schema:
//           "$ref": "#/definitions/PackRequest"
//
// responses:
//     "200":
//         description: OK
//     "500":
//         description: server error

// ! /icons

// * GET /icons

// swagger:operation GET /icons IconRequests GetIcons
//
// Gets the list of all the icon requests present in the database.
// ---
// produces:
//     - application/json
//
// responses:
//     "200":
//         description: OK
//         schema:
//             type: array
//             items:
//                 "$ref": "#/definitions/IconResponse"
