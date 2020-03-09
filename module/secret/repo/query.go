package repo

const (
	insertAppQuery = `
		INSERT INTO scr_app (app_id, secret)
		VALUES (:app_id, :secret)`

	getAppQuery = `
		SELECT (app_id, secret)
		FROM scr_app
		WHERE app_id = $1`

	insertSecretQuery = `
		INSERT INTO scr_secret (app_id, data, create_by)
		VALUES (:app_id, :data, :create_by")`

	getAllSecretQuery = `
		SELECT (id, app_id, data, create_time, create_by)
		FROM scr_secret
		WHERE app_id = $1
		ORDER BY create_time DESC
		LIMIT $2`

	getLastSecretQuery = `
		SELECT (id, app_id, data, create_time, create_by)
		FROM scr_secret
		WHERE app_id = $1
		ORDER BY create_time DESC
		LIMIT 1`
)
