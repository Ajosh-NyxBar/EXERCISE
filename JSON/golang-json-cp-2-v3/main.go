package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type LoanData struct {
	StartBalance int
	Data         []Loan
	Employees    []Employee
}

type Loan struct {
	Date        string
	EmployeeIDs []string
}

type Employee struct {
	ID       string
	Name     string
	Position string
}

type LoanRecord struct {
	MonthDate    string     `json:"month_date"`
	StartBalance int        `json:"start_balance"`
	EndBalance   int        `json:"end_balance"`
	Borrowers    []Borrower `json:"borrowers"`
}

type Borrower struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Total int    `json:"total_loan"`
}

func FindEmployee(id string, employees []Employee) (Employee, bool) {
	for _, employee := range employees {
		if employee.ID == id {
			return employee, true
		}
	}
	return Employee{}, false
}

func GetEndBalanceAndBorrowers(data LoanData) (int, []Borrower) {
	if len(data.Data) != 0 {
		var tempMap = map[string]int{}

		for _, v := range data.Data {
			for _, id := range v.EmployeeIDs {
				if data.StartBalance >= 50000 {
					tempMap[id] += 50000
					data.StartBalance -= 50000
				}
			}
		}

		var keys []string
		for k := range tempMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var borrowers []Borrower
		for _, k := range keys {
			if employee, ok := FindEmployee(k, data.Employees); ok {
				borrowers = append(borrowers, Borrower{
					ID:    employee.ID,
					Name:  employee.Name,
					Total: tempMap[k],
				})
			}
		}

		return data.StartBalance, borrowers
	} else {
		return 0, []Borrower{}
	}
}

func LoanReport(data LoanData) LoanRecord {
	endBalance, borrowers := GetEndBalanceAndBorrowers(data)
	monthDate := extractMonthDate(data.Data)
	return LoanRecord{
		MonthDate:    monthDate,
		StartBalance: data.StartBalance,
		EndBalance:   endBalance,
		Borrowers:    borrowers,
	}
}

func extractMonthDate(loans []Loan) string {
	if len(loans) > 0 {
		parts := strings.Split(loans[0].Date, "-")
		if len(parts) > 1 {
			return parts[1] + " " + parts[2]
		}
	}
	return ""
}

func RecordJSON(record LoanRecord, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(record)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}

func main() {
	records := LoanReport(LoanData{
		StartBalance: 500000,
		Data: []Loan{
			{"01-January-2021", []string{"1", "2"}},
			{"02-January-2021", []string{"1", "2", "3"}},
			{"03-January-2021", []string{"2", "3"}},
			{"04-January-2021", []string{"1", "3"}},
		},
		Employees: []Employee{
			{"1", "Employee A", "Manager"},
			{"2", "Employee B", "Staff"},
			{"3", "Employee C", "Staff"},
		},
	})

	err := RecordJSON(records, "loan-records.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(records)
}
