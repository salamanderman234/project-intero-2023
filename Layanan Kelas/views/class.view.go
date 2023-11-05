package view

import (
	"math"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
)

type classView struct {
	classService domain.ClassService
}

func NewClassView(c domain.ClassService) domain.ClassView {
	return &classView{
		classService: c,
	}
}
func (cv *classView) Create(c echo.Context) error {
	requestContext := c.Request().Context()
	respStatusCode := http.StatusCreated
	resp := domain.BasicResponse{
		Message: "kelas berhasil dibuat",
		Datas: nil,
		Errors: nil,
	}
	var createForm domain.ClassCreateForm
	if err := c.Bind(&createForm); err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return c.JSON(respStatusCode, resp)
	}
	created, err := cv.classService.CreateClass(requestContext, createForm)
	if err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return c.JSON(respStatusCode, resp)
	}
	resp.Datas = map[string]any{
		"created_data": created,
	}
	return c.JSON(respStatusCode, resp)
}
func (cv *classView) Read(c echo.Context) error {
	requestContext := c.Request().Context()
	respStatusCode := http.StatusOK
	resp := domain.BasicResponse{
		Message: "ok",
		Datas: nil,	
		Errors: nil,
	}
	pageString := c.QueryParam("page")
	query := c.QueryParam("q")
	idString := c.QueryParam("id")
	
	if idString != "" {
		id,_ := strconv.Atoi(idString)
		data, err := cv.classService.GetClassInfo(requestContext, uint(id))
		if err != nil {
			respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
			return c.JSON(respStatusCode, resp)
		}
		resp.Datas = data
		return c.JSON(respStatusCode, resp)
	}
	pageInt,_ := strconv.Atoi(pageString)
	page := uint(math.Max(float64(pageInt), 1))
	datas, pagination, err := cv.classService.GetClassList(requestContext, query, page)
	if err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return c.JSON(respStatusCode, resp)
	}
	resp.Datas = map[string]any {
		"pagination" : pagination,
		"datas" : datas,
		"length" : len(datas),
	}
	return c.JSON(respStatusCode, resp)
}

func (cv *classView) Update(c echo.Context) error {
	requestContext := c.Request().Context()
	respStatusCode := http.StatusOK
	resp := domain.BasicResponse{
		Message: "kelas berhasil diupdate",
		Datas: nil,
		Errors: nil,
	}
	var updateForm domain.ClassUpdateForm
	if err := c.Bind(&updateForm); err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return c.JSON(respStatusCode, resp)
	}
	idString := c.QueryParam("id")
	id,_ := strconv.Atoi(idString)
	aff, updated, err := cv.classService.UpdateClassInfo(requestContext, uint(id), updateForm)
	if err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return c.JSON(respStatusCode, resp)
	}
	resp.Datas = map[string]any{
		"class_id" : id,
		"updated_data" : updated,
		"rows_affected": aff,
	}
	return c.JSON(respStatusCode, resp)
}
func (cv *classView) Delete(c echo.Context) error {
	requestContext := c.Request().Context()
	respStatusCode := http.StatusOK
	resp := domain.BasicResponse{
		Message: "kelas berhasil dihapus",
		Datas: nil,
		Errors: nil,
	}
	idString := c.QueryParam("id")
	id,_ := strconv.Atoi(idString)
	_, err := cv.classService.DeleteClass(requestContext, uint(id))
	if err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return c.JSON(respStatusCode, resp)
	}
	resp.Datas = map[string]any {
		"class_id" : id,
	}
	return c.JSON(respStatusCode, resp)
}