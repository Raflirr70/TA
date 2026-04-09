package handler

import (
	"net/http"
	"strconv"
	models "tipes/internal/model"
	"tipes/internal/service"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	service service.EmployeeService
}

func NewEmployeeHandler(s service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: s}
}

func (h *EmployeeHandler) GetEmployees(c *gin.Context) {
	branchIDStr := c.Query("branch_id")
	branchID, err := strconv.Atoi(branchIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "branch_id tidak valid"})
		return
	}

	employees, err := h.service.GetEmployees(uint(branchID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func (h *EmployeeHandler) HireEmployee(c *gin.Context) {
	branchID, _ := strconv.Atoi(c.PostForm("branch_id"))
	roleID, _ := strconv.Atoi(c.PostForm("role_id"))

	emp := models.User{
		FirstName:   c.PostForm("firstname"),
		LastName:    c.PostForm("lastname"),
		Email:       c.PostForm("email"),
		Password:    c.PostForm("password"),
		NoTelephone: c.PostForm("no_telephone"),
		RoleID:      uint(roleID),
		BranchID:    uint(branchID),
	}

	if err := h.service.HireEmployee(emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "employee berhasil ditambahkan"})
}

func (h *EmployeeHandler) EditEmployee(c *gin.Context) {
	userID, _ := strconv.Atoi(c.PostForm("user_id"))
	// branchID, _ := strconv.Atoi(c.PostForm("branch_id"))

	emp := models.User{
		UserID:      uint(userID),
		FirstName:   c.PostForm("firstname"),
		LastName:    c.PostForm("lastname"),
		Email:       c.PostForm("email"),
		NoTelephone: c.PostForm("no_telephone"),
		// BranchID:    uint(branchID),
	}

	if err := h.service.EditEmployee(emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "employee berhasil diupdate"})
}

func (h *EmployeeHandler) FireEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee_id tidak valid"})
		return
	}

	if err := h.service.FireEmployee(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "employee berhasil di-nonaktifkan"})
}
