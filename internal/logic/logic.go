package logic

import (
	"net/http"

	"github.com/PereRohit/util/log"
	respModel "github.com/PereRohit/util/model"

	"github.com/vatsal278/PdfConversion/internal/model"
	"github.com/vatsal278/PdfConversion/internal/repo/datasource"
)

//go:generate mockgen --build_flags=--mod=mod --destination=./../../pkg/mock/mock_logic.go --package=mock github.com/vatsal278/PdfConversion/internal/logic PdfConversionLogicIer

type PdfConversionLogicIer interface {
	Ping(*model.PingRequest) *respModel.Response
	HealthCheck() bool
}

type pdfConversionLogic struct {
	dummyDsSvc datasource.DataSource
}

func NewPdfConversionLogic(ds datasource.DataSource) PdfConversionLogicIer {
	return &pdfConversionLogic{
		dummyDsSvc: ds,
	}
}

func (l pdfConversionLogic) Ping(req *model.PingRequest) *respModel.Response {
	// add business logic here
	res, err := l.dummyDsSvc.Ping(&model.PingDs{
		Data: req.Data,
	})
	if err != nil {
		log.Error("datasource error", err)
		return &respModel.Response{
			Status:  http.StatusInternalServerError,
			Message: "",
			Data:    nil,
		}
	}
	return &respModel.Response{
		Status:  http.StatusOK,
		Message: "Pong",
		Data:    res,
	}
}

func (l pdfConversionLogic) HealthCheck() bool {
	// check all internal services are working fine
	return l.dummyDsSvc.HealthCheck()
}
