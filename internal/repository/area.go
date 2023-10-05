package repository

import (
	"context"
	"eos-layout/internal/model"
	"eos-layout/internal/status"
)

type AreaRepository interface {
	One(ctx context.Context, level, id int64) (*model.Area, error)
	Find(ctx context.Context, level, id int64, key string) ([]*model.Area, error)
}

func NewAreaRepository(r *Repository) AreaRepository {
	return &areaRepository{r}
}

type areaRepository struct {
	*Repository
}

func (r areaRepository) Find(ctx context.Context, level int64, id int64, key string) ([]*model.Area, error) {
	db := r.DB(ctx)
	tx := db.Where("level=?", level)
	if id != 0 {
		tx = tx.Where("parent_code=?", id)
	}
	if len(key) != 0 {
		tx = tx.Where("(name LIKE ? or pinyin LIKE ?)", key+"%", key+"%")
	}
	areas := make([]*model.Area, 0)
	err := tx.Find(&areas).Error
	if err != nil {
		return nil, err
	}
	return areas, nil
}

func (r areaRepository) One(ctx context.Context, level int64, id int64) (*model.Area, error) {
	area := &model.Area{}
	db := r.DB(ctx)
	err := db.Where("level=? and area_code=?", level, id).First(area).Error
	if err != nil && !status.IsRecordNotFound(err) {
		return nil, err
	}
	return area, nil
}
