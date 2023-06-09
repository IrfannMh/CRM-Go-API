package customers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RequestHandler struct {
	ctrl ControllerCustomerInterface
}

type RequestHandlerInterface interface {
	Delete(c *gin.Context)
	GetAll(c *gin.Context)
	Create(c *gin.Context)
}

func NewRequestHandler(ctrl ControllerCustomerInterface) RequestHandlerInterface {
	return &RequestHandler{
		ctrl: ctrl,
	}
}

func DefaultRequestHandler(db *gorm.DB) RequestHandlerInterface {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}

type CreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
type CustomerResponse struct {
}

func (h RequestHandler) Create(c *gin.Context) {
	var req CreateRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := h.ctrl.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) GetAll(c *gin.Context) {
	email := c.Param("email")
	pageParam := c.Query("page")
	pageInt, _ := strconv.Atoi(pageParam)
	page := (pageInt - 1) * 6
	_ = page
	res, err := h.ctrl.GetAll(page, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) Delete(c *gin.Context) {
	customerID := c.Param("id")
	customer, err := h.ctrl.FindByID(customerID)

	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Customer Not Found"})
	}

	err = h.ctrl.Delete(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, MessageResponse{Message: "Delete Customer Success"})
}
