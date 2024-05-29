package response

import (
	admin_response "e-complaint-api/controllers/admin/response"
	"e-complaint-api/entities"
)

type Update struct {
	ID          int                 `json:"id"`
	ComplaintID string              `json:"complaint_id"`
	Admin       *admin_response.Get `json:"admin"`
	Status      string              `json:"status"`
	Message     string              `json:"message"`
}

func UpdateFromEntitiesToResponse(data *entities.ComplaintProcess) *Update {
	return &Update{
		ID:          data.ID,
		ComplaintID: data.ComplaintID,
		Admin:       admin_response.GetFromEntitiesToResponse(&data.Admin),
		Status:      data.Status,
		Message:     data.Message,
	}
}