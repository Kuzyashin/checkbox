package handlers

import (
	_ "github.com/AKuzyashin/checkbox/docs"
	"github.com/AKuzyashin/checkbox/internal/models"
	"github.com/AKuzyashin/checkbox/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

type Handlers struct {
	msgChan  chan models.WorkerMsg
	services *services.Services
}

func NewHandlers(srv *services.Services, msgChan chan models.WorkerMsg) *Handlers {
	return &Handlers{services: srv, msgChan: msgChan}
}

func (h *Handlers) Register(app *gin.Engine) {
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.POST("/route", h.CreateRequest)
	app.GET("/route/:request_id", h.GetResult)
}

// CreateRequest godoc
// @Summary Create request for calculation.
// @Description create calculation request.
// @ID create-request
// @Tags routes
// @Accept application/json
// @Produce json
// @Success 200 {object} RouteCreatedJson
// @Failure 400 {object} ErrorJson
// @Failure 500 {object} ErrorJson
// @Router /route [post]
func (h *Handlers) CreateRequest(c *gin.Context) {
	var requestIn models.Route
	if err := c.ShouldBindJSON(&requestIn); err != nil {
		c.JSON(http.StatusBadRequest, ErrorJson{Error: err.Error()})
		return
	}
	err := h.services.Routes.CreateRequest(&requestIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorJson{Error: "Fatal error"})
	}
	h.msgChan <- requestIn.ToWorkerMsg()
	c.JSON(http.StatusCreated, RouteCreatedJson{RouteId: requestIn.ID})
}

// GetResult godoc
// @Summary Returns result of calculation.
// @Description get the result of calculation.
// @ID get-result-by-id
// @Tags routes
// @Accept */*
// @Produce json
// @Param request_id path uint true "Request ID"
// @Success 200 {object} models.Route
// @Failure 400 {object} ErrorJson "Bad request. Can not parse request_id"
// @Failure 500 {object} ErrorJson "Fatal error"
// @Failure 425 "Route calculation not completed yet. Try later"
// @Failure 404 "Route not found"
// @Router /route/{request_id} [get]
func (h *Handlers) GetResult(c *gin.Context) {
	var uri requestUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, ErrorJson{Error: "can not parse request_id"})
		return
	}
	result, err := h.services.Routes.GetResult(uri.RequestID)
	switch {
	case !result.IsValid():
		c.JSON(http.StatusNotFound, nil)
	case !result.HasResult():
		c.JSON(http.StatusTooEarly, nil)
	case err != nil:
		c.JSON(http.StatusInternalServerError, ErrorJson{Error: "Fatal error"})
	default:
		c.JSON(http.StatusOK, result)
	}

}
