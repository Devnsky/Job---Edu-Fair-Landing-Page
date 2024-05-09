package pengunjung

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
	PengunjungService domain.PengunjungService
}

func NewApi(app *fiber.App, pengunjungService domain.PengunjungService) {
	a := api{
		PengunjungService: pengunjungService,
	}

	app.Post("/api/pengunjung/register", a.Register)
	app.Get("/pengunjung/export", a.ExportToExcel)
}

func (a api) Register(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	var PengunjungModel domain.PengunjungData
	if err := ctx.BodyParser(&PengunjungModel); err != nil {
		apiResponse := domain.ApiResponse{
			Code:    "01",
			Message: err.Error(),
		}
		util.ResponseInterceptor(c, &apiResponse)
		return ctx.Status(400).JSON(apiResponse)
	}
	apiResponse := a.PengunjungService.Register(c, PengunjungModel)
	util.ResponseInterceptor(c, &apiResponse)

	return ctx.Status(200).JSON(apiResponse)
}

func (a api) ExportToExcel(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	apiResponse := a.PengunjungService.GetAll(c)

	// Buat file Excel baru
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	data, ok := apiResponse.Data.([]domain.PengunjungData)
	if !ok {
		// Handle the error case where apiResponse.Data is not of type []domain.DataPeminjam
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid data type",
		})
	}

	headers := []string{"No", "Nama", "Nomor HP/WA", "Kategori Pengunjung", "Tahun Lulus", "Asal Sekolah/Instansi", "Jurusan"}
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
		row.AddCell().Value = d.Nama
		row.AddCell().Value = d.Nohp
		row.AddCell().Value = d.Kategori
		row.AddCell().Value = d.TahunLulus
		row.AddCell().Value = d.Asal
		row.AddCell().Value = d.Jurusan
		for _, cell := range row.Cells {
			cell.SetStyle(style)
		}
	}
	ctx.Response().Header.SetContentType("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Response().Header.Set("Content-Disposition", "attachment; filename=Pengunjung.xlsx")
	err = file.Write(ctx.Response().BodyWriter())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return nil
}
