package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"qure/models"
)

func GetEmployee(c echo.Context) error {
	requested_id := c.Param("id")
	result := models.GetEmployeeDB(requested_id)
	return c.JSON(http.StatusOK, result)
}

func GetEmployees(c echo.Context) error {
	result := models.GetEmployeesDB()
	// println("foo")
	return c.JSON(http.StatusOK, result)
}

func PostEmployee(c echo.Context) error {
	emp := models.Employee{}
	if err := c.Bind(&emp); err != nil {
		return err
	}
	result := models.InsertEmployeeDB(emp)

	return c.JSON(http.StatusCreated, result)
}

func DeleteEmployee(c echo.Context) error {
	requested_id := c.Param("id")
	fmt.Println(requested_id)

	_ = models.DeleteEmployeeDB(requested_id)

	return c.JSON(http.StatusOK, "Deleted")
}