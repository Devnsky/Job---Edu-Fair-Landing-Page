package feedback

import (
	"context"
	"eduFair/domain"
	"eduFair/internal/util"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tealeg/xlsx"
)

type api struct {
	FeedbackService domain.FeedbackService
}

func NewApi(app *fiber.App, feedbackService domain.FeedbackService) {
	a := api{
		FeedbackService: feedbackService,
	}

	app.Post("/api/feedback/register", a.Register)
	app.Get("/feedback/export", a.Export)
}

func (a api) Register(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	var feedbacks domain.FeedbackData
	if err := ctx.BodyParser(&feedbacks); err != nil {
		apiResponse := domain.ApiResponse{
			Code:    "01",
			Message: err.Error(),
		}
		util.ResponseInterceptor(c, &apiResponse)
		return ctx.Status(400).JSON(apiResponse)
	}
	apiResponse := a.FeedbackService.Register(c, feedbacks)
	util.ResponseInterceptor(c, &apiResponse)

	return ctx.Status(200).JSON(apiResponse)
}

func (a api) Export(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	apiResponse := a.FeedbackService.GetAll(c)

	// Buat file Excel baru
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	data, ok := apiResponse.Data.([]domain.FeedbackData)
	if !ok {
		// Handle the error case where apiResponse.Data is not of type []domain.DataPeminjam
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid data type",
		})
	}

	headers := []string{"No", "Feedbacks"}
	headerRow := sheet.AddRow()

	for _, h := range headers {
		cell := headerRow.AddCell()
		cell.Value = h
		style := xlsx.NewStyle()
		style.Font = *xlsx.NewFont(12, "Calibri")
		style.Font.Bold = true
		style.Fill = *xlsx.NewFill("solid", "FFFF00", "FFFFFFFF")
		style.Border = *xlsx.NewBorder("thin", "thin", "thin", "thick")
		cell.SetStyle(style)
	}

	for i, d := range data {
		row := sheet.AddRow()
		style := xlsx.NewStyle()
		style.Font = *xlsx.NewFont(11, "Calibri")
		style.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")
		row.AddCell().Value = strconv.Itoa(i + 1)
		feed := row.AddCell()
		feed.Value = d.Feedback
		for _, cell := range row.Cells {
			cell.SetStyle(style)
			cell.GetStyle().Alignment.WrapText = true
			cell.GetStyle().Alignment.Vertical = "center"
		}
	}
	ctx.Response().Header.SetContentType("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Response().Header.Set("Content-Disposition", "attachment; filename=Feedback.xlsx")
	err = file.Write(ctx.Response().BodyWriter())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return nil
}
