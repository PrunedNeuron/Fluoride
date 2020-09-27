package swagger

import "github.com/PrunedNeuron/Fluoride/pkg/model"

// DevRequest is the developer request body
// swagger:model
type DevRequest struct {
	// status of the request
	// example: success
	Role string `json:"role"`

	// name of the developer
	// example: John Doe
	Name string `json:"name"`

	// username of the developer
	// example: jdoe
	Username string `json:"username"`

	// email address of the developer
	// example: jdoe@gmail.com
	Email string `json:"email"`

	// website of the developer
	// example: https://ayushm.dev
	URL string `json:"url"`
}

// IconResponse is the developer response
// swagger:model
type IconResponse struct {
	// the status of the request
	// example: success
	Status string `json:"status"`

	// the response message
	// example: retrieved x icons
	Message string `json:"message"`

	// the list of icons
	Icons []model.Icon `json:"icons"`
}

// PackResponse is the icon pack response
// swagger:model
type PackResponse struct {
	// the status of the request
	// example: success
	Status string `json:"status"`

	// the list of packs
	Packs []model.Pack `json:"packs"`
}

// PackRequest is the icon pack request
// swagger:model
type PackRequest struct {
	// name of the icon pack
	// example: Amphetamine
	Name string `json:"name"`

	// username of the icon pack's developer
	// example: ayush
	DevUsername string `json:"developer_username"`

	// play store url of the icon pack
	// example: https://play.google.com/store/apps/details?id=com.trello
	URL string `json:"url"`

	// billing status of the icon pack
	// example: active
	BillingStatus string `json:"billing_status"`
}