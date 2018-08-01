package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	StudentCode string `json:"studentcode"`
	NameThai    string `json:"namethai"`
	NameEng     string `json:"nameeng"`
	Year        string `json:"year"`
	Semester    string `json:"semester"`
	CampusNo    string `json:"campusno"`
}

func FetchAllStudent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			students []Student
			student  Student
			result   gin.H
		)
		rows, err := db.Query("select * from XXX_STUDENT where STD_CODE like '591%'")

		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"error":  err.Error(),
			}
			c.JSON(http.StatusBadRequest, result)
		} else {
			for rows.Next() {
				rows.Scan(&student.StudentCode, &student.NameThai, &student.NameEng, &student.Year, &student.Semester, &student.CampusNo)
				students = append(students, student)
			}
			result = gin.H{
				"result": students,
				"error":  nil,
			}
			c.JSON(http.StatusOK, result)
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}
}

func Fetchstudent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			student Student
			result  gin.H
		)
		id := c.Param("id")
		rows := db.QueryRow("select * from XXX_STUDENT where std_code = ?", id)
		err := rows.Scan(&student.StudentCode, &student.NameThai, &student.NameEng, &student.Year, &student.Semester, &student.CampusNo)
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"error":  err.Error(),
			}
			c.JSON(http.StatusBadRequest, result)
		} else {

			result = gin.H{
				"result": student,
				"error":  nil,
			}
			c.JSON(http.StatusOK, result)
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}
}

func Createstudent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			student Student
			result  gin.H
		)
		student.StudentCode = c.PostForm("studentcode")
		student.NameThai = c.PostForm("namethai")
		student.NameEng = c.PostForm("nameeng")
		student.Year = c.PostForm("year")
		student.Semester = c.PostForm("semester")
		student.CampusNo = c.PostForm("campusno")
		_, err := db.Exec(`insert into XXX_STUDENT (STD_CODE,NAME_THAI,NAME_ENG,ENROLL_YEAR,ENROLL_SEMESTER,CAMPUS_NO) values (?,?,?,?,?,?)`, student.StudentCode, student.NameThai, student.NameEng, student.Year, student.Semester, student.CampusNo)
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"error":  err.Error(),
			}
			c.JSON(http.StatusBadRequest, result)
		} else {
			result = gin.H{
				"result": true,
				"error":  nil,
			}
			c.JSON(http.StatusOK, result)
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}
}

func Updatestudent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			student Student
			result  gin.H
		)

		student.StudentCode = c.PostForm("studentcode")
		student.NameThai = c.PostForm("namethai")
		student.NameEng = c.PostForm("nameeng")
		student.Year = c.PostForm("year")
		student.Semester = c.PostForm("semester")
		student.CampusNo = c.PostForm("campusno")
		_, err := db.Exec("update XXX_STUDENT set NAME_THAI = ? ,NAME_ENG = ? ,ENROLL_YEAR = ? ,ENROLL_SEMESTER = ? ,CAMPUS_NO = ? WHERE STD_CODE = ?", student.NameThai, student.NameEng, student.Year, student.Semester, student.CampusNo, student.StudentCode)
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"error":  err.Error(),
			}
			c.JSON(http.StatusBadRequest, result)
		} else {
			result = gin.H{
				"result": true,
				"error":  nil,
			}
			c.JSON(http.StatusOK, result)
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}
}

func Deletestudent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			result gin.H
			id     string
		)
		id = c.Param("id")
		_, err := db.Exec("delete from XXX_STUDENT where STD_CODE = ?", id)
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"error":  err.Error(),
			}
			c.JSON(http.StatusBadRequest, result)
		} else {
			result = gin.H{
				"result": true,
				"error":  nil,
			}
			c.JSON(http.StatusOK, result)
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}
}
