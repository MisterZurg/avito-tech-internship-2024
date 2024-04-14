package repository

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/MisterZurg/avito-tech-internship-2024/api"
)

type BannersRepository struct {
	*sqlx.DB
}

// Получение всех баннеров c фильтрацией по фиче и/или тегу
// (GET /banner)
func (br *BannersRepository) GetBanner(ctx echo.Context, params api.GetBannerParams) error {
	//getBannerParams := new(api.GetBannerParams)
	//
	////// Loop through rows using only one struct
	//rows, err := br.DB.Queryx("SELECT banner_id, feature_id, array_to_json(regions), array_to_json(working_hours) FROM banners")
	//if err != nil {
	//	return ctx.JSON(http.StatusBadRequest, "Cannot select * from couriers in GetCouriers")
	//}

	return ctx.JSON(http.StatusBadRequest, "")
}

// Создание нового баннера
// (POST /banner)
func (br *BannersRepository) PostBanner(ctx echo.Context, params api.PostBannerParams) error {
	// banner := new(api.PostBannerJSONBody)
	//with create_banner AS (
	//	INSERT into banner (feature_id, is_active, "content") VALUES ($2, $3, $4) RETURNING "id", feature_id
	//),
	//create_banner_relation as (
	//	INSERT into banner_relation (banner_id, feature_id, tag_id)
	//SELECT cb.id as banner_id
	//, cb.feature_id as feature_id
	//, UNNEST($1::int[]) as tag_id
	//FROM create_banner AS cb
	//)
	//
	//SELECT "id" FROM create_banner;
	//row := br.DB.QueryRow(`INSERT INTO banners (feature_id, is_active, content) VALUES ($1, $2, $3) RETURNING banner_id`,
	//	banner.FeatureId,
	//	banner.IsActive,
	//	banner.Content,
	//)
	_ = br.DB.QueryRowx(`INSERT INTO banners (feature_id, is_active) VALUES ($1, $2) RETURNING banner_id`,
		1488,
		false,
	)

	return ctx.JSON(http.StatusBadRequest, "")
}

// Удаление баннера по идентификатору
// (DELETE /banner/{id})
func (br *BannersRepository) DeleteBannerId(ctx echo.Context, id int, params api.DeleteBannerIdParams) error {
	return ctx.JSON(http.StatusBadRequest, "")
}

// Обновление содержимого баннера
// (PATCH /banner/{id})
func (br *BannersRepository) PatchBannerId(ctx echo.Context, id int, params api.PatchBannerIdParams) error {
	return ctx.JSON(http.StatusBadRequest, "")
}

// Получение баннера для пользователя
// (GET /user_banner)
func (br *BannersRepository) GetUserBanner(ctx echo.Context, params api.GetUserBannerParams) error {
	return ctx.JSON(http.StatusBadRequest, "")
}
