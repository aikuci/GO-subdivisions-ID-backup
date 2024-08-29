package mapper

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"
)

type DistrictMapper struct{}

func NewDistrictMapper() *DistrictMapper {
	return &DistrictMapper{}
}

func (m *DistrictMapper) ModelToResponse(district *entity.District) *model.DistrictResponse {
	var province *model.ProvinceResponse
	if district.Province.ID > 0 {
		provinceMapper := NewProvinceMapper()
		province = provinceMapper.ModelToResponse(&district.Province)
	}
	var city *model.CityResponse
	if district.City.ID > 0 {
		cityMapper := NewCityMapper()
		city = cityMapper.ModelToResponse(&district.City)
	}

	return &model.DistrictResponse{
		BaseCollectionResponse: appmodel.BaseCollectionResponse[int]{ID: district.ID},
		IDCity:                 district.CityID,
		IDProvince:             district.ProvinceID,
		Code:                   district.Code,
		Name:                   district.Name,
		PostalCodes:            district.PostalCodes,

		Province: province,
		City:     city,
	}
}
