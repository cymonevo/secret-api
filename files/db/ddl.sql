
CREATE TABLE scr_app (
    app_id UUID PRIMARY KEY,
    secret BYTEA DEFAULT NULL
);

CREATE TABLE scr_secret (
    id SERIAL PRIMARY KEY,
    app_id UUID NOT NULL,
    data BYTEA DEFAULT NULL,
    create_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    create_by INTEGER DEFAULT NULL,
    FOREIGN KEY (app_id) REFERENCES scr_app (app_id) ON DELETE CASCADE
);

CREATE INDEX ON scr_secret (app_id, create_time);