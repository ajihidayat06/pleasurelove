package request

import "pleasurelove/internal/utils"

type ReqCategory struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
	Slug string `json:"slug"`
}

var ReqCategoryErrorMessage = map[string]string{
	"Name": "name required",
	"Code": "code required",
}

func (r *ReqCategory) ValidateRequestCreate() error {
	err := utils.ValidateCode(r.Code)
	if err != nil {
		return err
	}

	r.Slug = utils.GenerateSlug(r.Name)
	return nil
}

type ReqCategoryUpdate struct {
	ID   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
	Slug string `json:"slug"`
	AbstractRequest
}

var ReqCategoryUpdateErrorMessage = map[string]string{
	"ID":            "id required",
	"Name":          "name required",
	"Code":          "code required",
	"UpdateddAtStr": "updated_at required",
}

func (r *ReqCategoryUpdate) ValidateRequestUpdate() error {
	if err := r.ValidateUpdatedAt(); err != nil {
		return err
	}

	err := utils.ValidateUsername(r.Code)
	if err != nil {
		return err
	}

	r.Slug = utils.GenerateSlug(r.Name)
	return nil
}
