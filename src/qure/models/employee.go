package models

import (
	"database/sql"
	"fmt"
	"qure/db"
)

type Employee struct {
	Id		string		`json:"id"`
	Name 	string 		`json:"employee_name"`
	Salary 	string		`json:"employee_salary"`
	Age 	string 		`json:"employee_age"`
}

type Employees struct {
	Employees []Employee `json:"employee"`
}

var con *sql.DB

func GetEmployeeDB(id string) Employee {
	con := db.CreateCon()
	sqlStatement := "SELECT id, employee_name, employee_age, employee_salary FROM employee WHERE id = ?"

	emp := Employee{}

	err := con.QueryRow(sqlStatement, id).Scan(&emp.Id, &emp.Name, &emp.Salary, &emp.Age)

	if err != nil {
		fmt.Println(err)
	}

	return emp
}

func GetEmployeesDB() Employees {
	con := db.CreateCon()

	sqlStatement := "SELECT id, employee_name, employee_age, employee_salary FROM employee ORDER BY id"
	rows, err := con.Query(sqlStatement)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := Employees{}

	for rows.Next() {
		employee := Employee{}

		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)
		if err2 != nil {
			fmt.Println(err)
		}
		result.Employees = append(result.Employees, employee)
	}
	return result
}

func InsertEmployeeDB(emp Employee) string {
	con := db.CreateCon()

	sqlStatement := "INSERT INTO employee(employee_name, employee_salary, employee_age) VALUES(?, ?, ?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()
	result, err2 := stmt.Exec(emp.Name, emp.Salary, emp.Age)

	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())
	return emp.Name
}

func DeleteEmployeeDB(id string) bool {
	con := db.CreateCon()

	sqlStatement := "DELETE FROM employee WHERE id = ?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		fmt.Println(err)
	}

	result, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.RowsAffected())
	return true
}