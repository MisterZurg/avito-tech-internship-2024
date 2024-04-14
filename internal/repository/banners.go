package repository

import (
	"database/sql"
	"github.com/lib/pq"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/MisterZurg/avito-tech-internship-2024/api"
)

type BannersRepository struct {
	*sqlx.DB
}

type Banner struct {
	ID        int       `db:"id"`
	TagIDs    []int64   `db:"tag_ids"`
	FeatureID int       `db:"feature_id"`
	Content   []byte    `db:"content"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Получение всех баннеров c фильтрацией по фиче и/или тегу
// (GET /banner)
func (br *BannersRepository) GetBanner(ctx echo.Context, params api.GetBannerParams) error {
	bannerResponse := make([]Banner, 0)
	// if !(params.FeatureId != nil || params.TagId != nil) {
	rows, err := br.DB.Queryx(`
			SELECT
				b.id,
				b.feature_id,
				b.content,
				b.is_active,
				b.created_at,
				b.updated_at,
				(SELECT array_agg(tag_id) FROM banners_relation AS br WHERE br.banner_id = b.id) AS tag_ids
			FROM banners as b
			ORDER BY b.created_at DESC
			LIMIT $1 OFFSET $2
			`,
		params.Limit,
		params.Offset,
	)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Cannot select * from orders in GetOrders")
	}

	for rows.Next() {
		//var tagIds []int64
		banner := new(Banner)
		err := rows.Scan(
			&banner.ID,
			&banner.FeatureID,
			&banner.Content,
			&banner.IsActive,
			&banner.CreatedAt,
			&banner.UpdatedAt,
			(*pq.Int64Array)(&banner.TagIDs),
		)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, "cannot fill banner in GetBanner")
		}
		//banner.TagIDs = tagIds
		bannerResponse = append(bannerResponse, *banner)
	}
	return ctx.JSON(http.StatusOK, bannerResponse)
}

// Создание нового баннера
// (POST /banner)
func (br *BannersRepository) PostBanner(ctx echo.Context, params api.PostBannerParams) error {
	banner := new(api.PostBannerJSONBody)

	if err := ctx.Bind(&banner); err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid format for PostBanner")
	}

	_ = br.DB.QueryRowx(`
		WITH create_banner AS (
			INSERT INTO banners (feature_id, is_active, content) VALUES ($1, $2, $3) RETURNING id, feature_id
		),
		create_banner_relation AS (
			INSERT INTO banners_relation (banner_id, feature_id, tag_id)
		SELECT cb.id AS banner_id
		, cb.feature_id AS feature_id
		, UNNEST($4::INT[]) AS tag_id
		FROM create_banner AS cb
		)
		SELECT "id" FROM create_banner;
	`,
		banner.FeatureId,
		banner.IsActive,
		banner.Content,
		banner.TagIds,
	)

	return ctx.JSON(http.StatusCreated, "Created")
}

// Удаление баннера по идентификатору
// (DELETE /banner/{id})
func (br *BannersRepository) DeleteBannerId(ctx echo.Context, id int, params api.DeleteBannerIdParams) error {
	_ = br.DB.QueryRowx(
		`DELETE from banners WHERE id = $1`,
		id,
	)
	return ctx.JSON(http.StatusBadRequest, "")
}

// Обновление содержимого баннера
// (PATCH /banner/{id})
func (br *BannersRepository) PatchBannerId(ctx echo.Context, id int, params api.PatchBannerIdParams) error {
	row := br.DB.QueryRowx(`
		SELECT
			b.id,
			b.feature_id,
			b.content,
			b.is_active,
			b.created_at,
			b.updated_at,
			(SELECT ARRAY_AGG(tag_id) FROM banners_relation AS br WHERE br.banner_id = b.id) as tag_ids
		FROM banners as b
		WHERE b.id = $1;
	`,
		id,
	)
	banner := new(Banner)
	err := row.Scan(
		&banner.ID,
		&banner.FeatureID,
		&banner.Content,
		&banner.IsActive,
		&banner.CreatedAt,
		&banner.UpdatedAt,
		(*pq.Int64Array)(&banner.TagIDs),
	)
	if err == sql.ErrNoRows {
		return ctx.JSON(http.StatusNotFound, "")
	}

	updBanner := new(api.PostBannerJSONBody)

	if err := ctx.Bind(&updBanner); err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid format for PostBanner")
	}

	_ = br.DB.QueryRowx(`
		UPDATE banners SET 
			feature_id=$2,
			is_active=$3, 
			content=$4
		WHERE id=$1
		`,
		id,
		updBanner.FeatureId,
		updBanner.IsActive,
		updBanner.Content,
	)

	return ctx.JSON(http.StatusBadRequest, "")

}

// Получение баннера для пользователя
// (GET /user_banner)
func (br *BannersRepository) GetUserBanner(ctx echo.Context, params api.GetUserBannerParams) error {
	row := br.DB.QueryRowx(`
	WITH find_banner AS (
		SELECT banner_id, tag_id, feature_id FROM banners_relation WHERE tag_id=$1 AND feature_id=$2
		)
	SELECT
		b.id,
		b.feature_id,
		b.content,
		b.is_active,
		b.created_at,
		b.updated_at,
		(SELECT ARRAY_AGG(tag_id) FROM banners_relation AS br WHERE br.banner_id = b.id) AS tag_ids
	FROM banners AS b JOIN find_banner AS fb ON (b.id = fb.banner_id);
	`,
		params.TagId,
		params.FeatureId,
	)

	userBanner := new(Banner)
	err := row.Scan(
		&userBanner.ID,
		&userBanner.FeatureID,
		&userBanner.Content,
		&userBanner.IsActive,
		&userBanner.CreatedAt,
		&userBanner.UpdatedAt,
		(*pq.Int64Array)(&userBanner.TagIDs),
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "cannot fill Banner GetUserBanner")
	}
	return ctx.JSON(http.StatusOK, userBanner.Content)
}
