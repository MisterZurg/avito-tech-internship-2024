CREATE TABLE IF NOT EXISTS banners (
    id SERIAL PRIMARY KEY,
    feature_id INT NOT NULL,
    content jsonb NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_active BOOL NOT NULL
);

CREATE TABLE IF NOT EXISTS banners_relation (
    banner_id INT NOT NULL REFERENCES banners ON DELETE CASCADE,
    feature_id INT NOT NULL,
    tag_id INT NOT NULL,
    UNIQUE(feature_id, tag_id)
);

CREATE INDEX IF NOT EXISTS banners_id_relation ON banners_relation (banner_id);