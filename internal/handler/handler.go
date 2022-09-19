package handler

import (
	"fmt"
	"net/http"

	"github.com/PereRohit/util/request"
	"github.com/PereRohit/util/response"

	"github.com/vatsal278/PdfConversion/internal/logic"
	"github.com/vatsal278/PdfConversion/internal/model"
	"github.com/vatsal278/PdfConversion/internal/repo/datasource"
)

const PdfConversionName = "pdfConversion"

//go:generate mockgen --build_flags=--mod=mod --destination=./../../pkg/mock/mock_handler.go --package=mock github.com/vatsal278/PdfConversion/internal/handler PdfConversionHandler

type PdfConversionHandler interface {
	HealthChecker
	Ping(w http.ResponseWriter, r *http.Request)
}

type pdfConversion struct {
	logic logic.PdfConversionLogicIer
}

func NewPdfConversion(ds datasource.DataSource) PdfConversionHandler {
	svc := &pdfConversion{
		logic: logic.NewPdfConversionLogic(ds),
	}
	AddHealthChecker(svc)
	return svc
}

func (svc pdfConversion) HealthCheck() (svcName string, msg string, stat bool) {
	set := false
	defer func() {
		svcName = PdfConversionName
		if !set {
			msg = ""
			stat = true
		}
	}()
	stat = svc.logic.HealthCheck()
	set = true
	return
}

func (svc pdfConversion) Ping(w http.ResponseWriter, r *http.Request) {
	req := &model.PingRequest{}

	suggestedCode, err := request.FromJson(r, req)
	if err != nil {
		response.ToJson(w, suggestedCode, fmt.Sprintf("FAILED: %s", err.Error()), nil)
		return
	}
	// call logic
	resp := svc.logic.Ping(req)
	response.ToJson(w, resp.Status, resp.Message, resp.Data)
	return
}
