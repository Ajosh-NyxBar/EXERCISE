package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Study struct {
	StudyName   string  `json:"study_name"`
	StudyCredit int     `json:"study_credit"`
	Grade       string  `json:"grade"`
}

type Report struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Semester int     `json:"semester"`
	Studies  []Study `json:"studies"`
}

func ReadJSON(filename string) (Report, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Report{}, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return Report{}, err
	}

	var report Report
	err = json.Unmarshal(bytes, &report)
	if err != nil {
		return Report{}, err
	}

	return report, nil
}

func GradePoint(report Report) float64 {
	gradeScale := map[string]float64{
		"A":  4.0,
		"AB": 3.5,
		"B":  3.0,
		"BC": 2.5,
		"C":  2.0,
		"CD": 1.5,
		"D":  1.0,
		"DE": 0.5,
		"E":  0.0,
	}

	var totalScore float64
	var totalCredits int

	for _, study := range report.Studies {
		scale, exists := gradeScale[study.Grade]
		if exists {
			totalScore += scale * float64(study.StudyCredit)
			totalCredits += study.StudyCredit
		}
	}

	if totalCredits == 0 {
		return 0.0
	}

	return totalScore / float64(totalCredits)
}

func main() {
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
