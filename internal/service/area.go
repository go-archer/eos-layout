package service

import (
	"context"
	"eos-layout/internal/dto/request"
	"eos-layout/internal/dto/response"
	"eos-layout/internal/repository"
)

type AreaService interface {
	One(ctx context.Context, in *request.Area) (*response.Area, error)
	Find(ctx context.Context, in *request.Area) (*response.AreaList, error)
}

func NewAreaService(s *Service, areaRepo repository.AreaRepository) AreaService {
	return &areaService{s, areaRepo}
}

type areaService struct {
	*Service
	areaRepository repository.AreaRepository
}

func (s *areaService) One(ctx context.Context, in *request.Area) (*response.Area, error) {
	res, err := s.areaRepository.One(ctx, in.Level, in.ID)
	if err != nil {
		return nil, err
	}
	return &response.Area{
		ID:         res.AreaCode,
		ZipCode:    res.ZipCode,
		CityCode:   res.CityCode,
		Name:       res.Name,
		ShortName:  res.ShortName,
		MergerName: res.MergerName,
		Pinyin:     res.Pinyin,
		Lng:        res.Lng,
		Lat:        res.Lat,
	}, nil
}

func (s *areaService) Find(ctx context.Context, in *request.Area) (*response.AreaList, error) {
	res, err := s.areaRepository.Find(ctx, in.Level, in.ID, in.Key)
	if err != nil {
		return nil, err
	}
	items := make([]*response.Area, 0)
	for _, v := range res {
		items = append(items, &response.Area{
			ID:         v.AreaCode,
			ZipCode:    v.ZipCode,
			CityCode:   v.CityCode,
			Name:       v.Name,
			ShortName:  v.ShortName,
			MergerName: v.MergerName,
			Pinyin:     v.Pinyin,
			Lng:        v.Lng,
			Lat:        v.Lat,
		})
	}
	return &response.AreaList{Items: items}, nil
}
