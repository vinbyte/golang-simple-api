package students

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Student struct {
	db *sql.DB
}

type studentData struct {
	ID    int    `json:"id" form:"-"`
	Name  string `json:"name" form:"name"`
	Grade int    `json:"grade" form:"grade"`
}

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type studentResponse struct {
	List []studentData `json:"list"`
}

func New(db *sql.DB) Student {
	data := Student{
		db: db,
	}
	return data
}

// StudentList untuk mengambil daftar siswa di tabel students
func (s *Student) StudentList(c *gin.Context) {
	//prepare response
	var res response
	res.Status = 200
	res.Message = "success"
	var resData studentResponse
	//make data property empty JSON object ( {} )
	res.Data = new(struct{})
	//make list property empty JSON array
	resData.List = make([]studentData, 0)

	//get from db
	query := "select id, name, grade from students"
	rows, err := s.db.Query(query)
	if err != nil {
		res.Status = 500
		res.Message = err.Error()
		c.JSON(500, res)
		return
	}
	defer rows.Close()
	for rows.Next() {
		temp := studentData{}
		err := rows.Scan(&temp.ID, &temp.Name, &temp.Grade)
		if err != nil {
			res.Status = 500
			res.Message = err.Error()
			c.JSON(500, res)
		}
		resData.List = append(resData.List, temp)
	}

	//send the result
	res.Data = resData
	c.JSON(200, res)
}

func (s *Student) StudentAdd(c *gin.Context) {
	//prepare response
	var res response
	res.Status = 200
	res.Message = "success"
	//make data property empty JSON object ( {} )
	res.Data = new(struct{})

	//bind parameter into struct
	var data studentData
	c.Bind(&data)

	//insert the data
	query := `insert into students (name, grade) values(?, ?)`
	_, err := s.db.Exec(query, data.Name, data.Grade)
	if err != nil {
		res.Status = 500
		res.Message = err.Error()
		c.JSON(500, res)
		return
	}

	//send response
	c.JSON(200, res)
}
