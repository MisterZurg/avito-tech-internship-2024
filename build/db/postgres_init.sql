-- CREATE TABLE IF NOT EXISTS banners (
--     banner_id   SERIAL PRIMARY KEY,
--     feature_id  INTEGER,
--     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     is_active   BOOLEAN
-- );

CREATE TABLE IF NOT EXISTS banners (
    banner_id   SERIAL PRIMARY KEY,
    feature_id  INTEGER,
--     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_active   BOOLEAN
);