package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	type Employee struct {
		Id			string		`json:"id"`
		Name 		string 		`json:"employee_name"`
		Salary		string		`json:"employee_salary"`
		Age			string		`json:"employee_age"`
	}

	type Employees struct {
		Employees 	[]Employee 	`json:"employee"`
	}

	db, err := sql.Open("mysql", "root:qwerty123@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	e.POST("/employee", func(c echo.Context) error {
		emp := new(Employee)
		if err := c.Bind(emp); err != nil {
			return err
		}

		sql := "INSERT INTO employee(employee_name, employee_age, employee_salary) VALUES(?, ?, ?)"
		stmt, err := db.Prepare(sql)

		if err != nil {
			fmt.Print(err.Error())
		}
		defer stmt.Close()
		result, err2 := stmt.Exec(emp.Name, emp.Salary, emp.Age)

		if err2 != nil {
			panic(err2)
		}
		fmt.Println(result.LastInsertId())

		return c.JSON(http.StatusCreated, emp.Name)
	})

	e.DELETE("/employee/:id", func(c echo.Context) error {
		requested_id := c.Param("id")
		sql := "DELETE FROM employee WHERE id = ?"
		stmt, err := db.Prepare(sql)
		if err != nil {
			fmt.Println(err)
		}
		result, err2 := stmt.Exec(requested_id)
		if err2 != nil {
			panic(err2)
		}
		fmt.Println(result.RowsAffected())
		return c.JSON(http.StatusOK, "Deleted")
	})

	e.GET("/employee/:id", func(c echo.Context) error {
		requested_id := c.Param("id")
		fmt.Println(requested_id)
		var name string
		var id string
		var salary string
		var age string

		err = db.QueryRow("SELECT id, employee_name, employee_age, employee_salary FROM employee WHERE id = ?", requested_id).Scan(&id, &name, &age, &salary)

		if err != nil {
			fmt.Println(err)
		}

		response := Employee{Id: id, Name: name, Salary: salary, Age: age}

		return c.JSON(http.StatusOK, response)
	})
}