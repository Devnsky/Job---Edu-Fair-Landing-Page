package pengunjung

import (
	"context"
	"eduFair/domain"
)

type service struct {
	PengunjungRepository domain.PengunjungRepository
}

func NewService(pengunjungRepositori domain.PengunjungRepository) domain.PengunjungService {
	return &service{
		PengunjungRepository: pengunjungRepositori,
	}
}

func (s service) Register(ctx context.Context, data domain.PengunjungData) domain.ApiResponse {
	exists, err := s.PengunjungRepository.NameExist(ctx, data.Nama)
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
		}
	}
	if exists {
		return domain.ApiResponse{
			Code:    "400",
			Message: "Name already exists",
		}
	}

	pengunjungModel := domain.Pengunjung{
		Nama:       data.Nama,
		Nohp:       data.Nohp,
		TahunLulus: data.TahunLulus,
		Asal:       data.Asal,
		Jurusan:    data.Jurusan,
		Kategori:   data.Kategori,
	}

	err = s.PengunjungRepository.Register(ctx, &pengunjungModel)
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
		}
	}

	return domain.ApiResponse{
		Code:    "200",
		Message: "Success",
		Data:    pengunjungModel,
	}
}

func (s service) GetAll(ctx context.Context) domain.ApiResponse {
	alat, err := s.PengunjungRepository.GetAll(ctx)
	if err != nil {
		return domain.ApiResponse{
			Code:    "505",
			Message: "SYSTEM MALFUNCTION(03)",
		}
	}
	var pengunjungModel []domain.PengunjungData
	for _, v := range alat {
		pengunjungModel = append(pengunjungModel, domain.PengunjungData{
			Nama:       v.Nama,
			Nohp:       v.Nohp,
			TahunLulus: v.TahunLulus,
			Asal:       v.Asal,
			Jurusan:    v.Jurusan,
			Kategori:   v.Kategori,
		})
	}
	return domain.ApiResponse{
		Code:    "00",
		Message: "APPROVED",
		Data:    pengunjungModel,
	}
}
