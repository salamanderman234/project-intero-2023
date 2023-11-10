package view

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	helper "github.com/salamanderman234/project-intro-2023/layanan-kelas/helpers"
)

type createViewServiceCallFunc func(createdForm domain.Form) (any, error)
type customBindFunc func() ([]error)

func basicCreateView(c echo.Context, fun createViewServiceCallFunc, form domain.Form, customBinds ...customBindFunc) (int, domain.BasicResponse) {
	respStatusCode := http.StatusCreated
	resp := domain.BasicResponse{
		Message: "success",
		Datas:   nil,
		Errors:  nil,
	}
	if len(customBinds) == 1 {
		if errs := customBinds[0](); len(errs) > 0 {
			resp.Message = "request error"
			resp.Errors = helper.GenerateBindingErrorDetail(errs)
			return http.StatusBadRequest, resp
		}
	}else if err := c.Bind(form); err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return respStatusCode, resp
	}
	created, err := fun(form)
	if err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return respStatusCode, resp
	}
	resp.Datas = map[string]any{
		"created_data": created,
	}
	return respStatusCode, resp
}

type searchViewServiceCallFunc func(q string, page uint, orderBy string, order string) (any, domain.Pagination, error)
type findViewServiceCallFunc func(id uint) (any, error)

func basicSearchView(c echo.Context, searchFunc searchViewServiceCallFunc, findFunc findViewServiceCallFunc) (int, domain.BasicResponse) {
	respStatusCode := http.StatusOK
	resp := domain.BasicResponse{
		Message: "ok",
		Datas: nil,	
		Errors: nil,
	}
	
	var page uint
	var query string
	var id uint
	var orderBy string
	var order string
	
	if errs := echo.QueryParamsBinder(c).
		Uint("page", &page).
		Uint("id", &id).
		String("q", &query).
		String("order_by", &orderBy).
		String("order", &order).
		BindErrors(); len(errs) > 0 {
		resp.Message = "request error"
		resp.Errors = helper.GenerateBindingErrorDetail(errs)
		return http.StatusBadRequest, resp
	}

	var datas any
	var err error
	var pagination domain.Pagination

	if id != 0 {
		datas, err = findFunc(id)
		resp.Datas = map[string]any{
			"data": datas,
		}
	} else {
		datas, pagination, err = searchFunc(query, page, orderBy, order)
		resp.Datas = map[string]any {
			"pagination" : pagination,
			"datas" : datas,
		}
	}

	if err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return respStatusCode, resp
	}
	return respStatusCode, resp
}

type updateViewServiceCallFunc func(id uint, updateForm domain.Form) (int, any, error)
func basicUpdateView(c echo.Context, fun updateViewServiceCallFunc, form domain.Form, customBinds ...customBindFunc) (int, domain.BasicResponse) {
	respStatusCode := http.StatusOK
	resp := domain.BasicResponse{
		Message: "success",
		Datas: nil,
		Errors: nil,
	}
	if len(customBinds) == 1 {
		if errs := customBinds[0](); len(errs) > 0 {
			resp.Message = "request error"
			resp.Errors = helper.GenerateBindingErrorDetail(errs)
			return http.StatusBadRequest, resp
		}
	}else if err := c.Bind(form); err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return respStatusCode, resp
	}
	var id uint
	if errs := echo.QueryParamsBinder(c).
		Uint("id", &id).
		BindErrors(); len(errs) > 0 {
		resp.Message = "request error"
		resp.Errors = helper.GenerateBindingErrorDetail(errs)
		return http.StatusBadRequest, resp
	}
	if id == 0 {
		respStatusCode = http.StatusBadRequest
		resp.Message = "request error"
		resp.Errors = []domain.ErrorDetail{
			{Field: "id",Type: "required", Message: "id field is required"},
		}
		return respStatusCode, resp
	}
	aff, data, err := fun(id, form)
	if err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return respStatusCode, resp
	}
	resp.Datas = map[string]any{
		"class_id" : id,
		"updated_data" : data,
		"rows_affected": aff,
	}
	return respStatusCode, resp
}

type deleteViewServiceCallFunc func(id uint) (error)
func basicDeleteView(c echo.Context, deleteFunc deleteViewServiceCallFunc) (int, domain.BasicResponse){
	respStatusCode := http.StatusOK
	resp := domain.BasicResponse{
		Message: "success",
		Datas: nil,
		Errors: nil,
	}
	var id uint
	if errs := echo.QueryParamsBinder(c).
		Uint("id", &id).
		BindErrors(); len(errs) > 0 {
		resp.Message = "request error"
		resp.Errors = helper.GenerateBindingErrorDetail(errs)
		return http.StatusBadRequest, resp
	}
	if id == 0 {
		respStatusCode = http.StatusBadRequest
		resp.Message = "request error"
		resp.Errors = []domain.ErrorDetail{
			{Field: "id",Type: "required", Message: "id field is required"},
		}
		return respStatusCode, resp
	}
	err := deleteFunc(id)
	if err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return respStatusCode, resp
	}
	resp.Datas = map[string]any{
		"deleted_id": id,
	}
	return respStatusCode, resp
}