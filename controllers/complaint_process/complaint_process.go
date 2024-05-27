package complaint_process

import (
	"e-complaint-api/controllers/base"
	"e-complaint-api/controllers/complaint_process/request"
	"e-complaint-api/controllers/complaint_process/response"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ComplaintProcessController struct {
	complaintUseCase        entities.ComplaintUseCaseInterface
	complaintProcessUseCase entities.ComplaintProcessUseCaseInterface
}

func NewComplaintProcessController(complaintUseCase entities.ComplaintUseCaseInterface, complaintProcessUseCase entities.ComplaintProcessUseCaseInterface) *ComplaintProcessController {
	return &ComplaintProcessController{
		complaintUseCase:        complaintUseCase,
		complaintProcessUseCase: complaintProcessUseCase,
	}
}

func (cp *ComplaintProcessController) Create(c echo.Context) error {
	admin_id, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var complaintProcessRequest request.Create
	c.Bind(&complaintProcessRequest)

	complaintProcessRequest.AdminID = admin_id
	complaint_id := c.Param("complaint_id")
	complaintProcessRequest.ComplaintID = complaint_id

	err = cp.complaintUseCase.UpdateStatus(complaint_id, complaintProcessRequest.Status)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	complaintProcess, err := cp.complaintProcessUseCase.Create(complaintProcessRequest.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	complaintProcessResponse := response.CreateFromEntitiesToResponse(&complaintProcess)

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Create Complaint Process", complaintProcessResponse))
}

func (cp *ComplaintProcessController) GetByComplaintID(c echo.Context) error {
	complaint_id := c.Param("complaint_id")

	complaintProcesses, err := cp.complaintProcessUseCase.GetByComplaintID(complaint_id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var complaintProcessesResponse []*response.Get
	for _, complaintProcess := range complaintProcesses {
		complaintProcessesResponse = append(complaintProcessesResponse, response.GetFromEntitiesToResponse(&complaintProcess))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Complaint Process", complaintProcessesResponse))
}
