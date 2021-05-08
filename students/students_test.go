package students

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestStudentList(t *testing.T) {
	t.Run("success fetch", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		rows := sqlmock.NewRows([]string{"id", "name", "grade"}).
			AddRow(1, "student1", 1).
			AddRow(2, "student2", 1)

		mock.ExpectQuery("select id, name, grade from students").WillReturnRows(rows).RowsWillBeClosed()

		app := New(db)
		gin.SetMode(gin.TestMode)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, err = http.NewRequest("GET", "/v1/students", nil)
		assert.Nil(t, err)
		app.StudentList(c)
		var mockResponse response
		err = json.NewDecoder(rec.Body).Decode(&mockResponse)
		assert.Nil(t, err)
		responseData := mockResponse.Data.(map[string]interface{})
		responseListData, isExist := responseData["list"]
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, 200, mockResponse.Status)
		assert.True(t, isExist)
		assert.Len(t, responseListData, 2)
	})
	t.Run("error query", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		err = errors.New("unexpected error query")
		mock.ExpectQuery("select id, name, grade from students").WillReturnError(err).RowsWillBeClosed()

		app := New(db)
		gin.SetMode(gin.TestMode)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, err = http.NewRequest("GET", "/v1/students", nil)
		assert.Nil(t, err)
		app.StudentList(c)
		var mockResponse response
		err = json.NewDecoder(rec.Body).Decode(&mockResponse)
		assert.Nil(t, err)
		responseData := mockResponse.Data.(map[string]interface{})
		_, isExist := responseData["list"]
		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, 500, mockResponse.Status)
		assert.False(t, isExist)
	})
	t.Run("error scan", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		err = errors.New("scan error")
		rows := sqlmock.NewRows([]string{"id", "name", "grade"}).
			AddRow(1, "student1", false).
			RowError(1, err)

		mock.ExpectQuery("select id, name, grade from students").WillReturnRows(rows).RowsWillBeClosed()

		app := New(db)
		gin.SetMode(gin.TestMode)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, err = http.NewRequest("GET", "/v1/student", nil)
		assert.Nil(t, err)
		app.StudentList(c)
		var mockResponse response
		err = json.NewDecoder(rec.Body).Decode(&mockResponse)
		assert.Nil(t, err)
		responseData := mockResponse.Data.(map[string]interface{})
		_, isExist := responseData["list"]
		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, 500, mockResponse.Status)
		assert.False(t, isExist)
	})
}

func TestStudentAdd(t *testing.T) {
	t.Run("success add", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		var data studentData
		data.Name = "student 3"
		data.Grade = 2
		mock.ExpectExec("insert into students (name, grade) values(?, ?)").WithArgs(data.Name, data.Grade).WillReturnResult(sqlmock.NewResult(1, 1))

		app := New(db)
		gin.SetMode(gin.TestMode)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		var param = url.Values{}
		param.Set("name", data.Name)
		param.Set("grade", strconv.Itoa(data.Grade))
		var payload = bytes.NewBufferString(param.Encode())
		c.Request, err = http.NewRequest("POST", "/v1/student", payload)
		assert.Nil(t, err)
		c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		app.StudentAdd(c)
		var mockResponse response
		err = json.NewDecoder(rec.Body).Decode(&mockResponse)
		assert.Nil(t, err)
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, 200, mockResponse.Status)
	})
	t.Run("error insert", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		var data studentData
		data.Name = "student 3"
		data.Grade = 2
		mock.ExpectExec("insert into students (name, grade) values(?, ?)").WithArgs(data.Name, data.Grade).WillReturnError(errors.New("unexpected error when insert"))

		app := New(db)
		gin.SetMode(gin.TestMode)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		var param = url.Values{}
		param.Set("name", data.Name)
		param.Set("grade", strconv.Itoa(data.Grade))
		var payload = bytes.NewBufferString(param.Encode())
		c.Request, err = http.NewRequest("POST", "/v1/student", payload)
		assert.Nil(t, err)
		c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		app.StudentAdd(c)
		var mockResponse response
		err = json.NewDecoder(rec.Body).Decode(&mockResponse)
		assert.Nil(t, err)
		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, 500, mockResponse.Status)
	})
}
