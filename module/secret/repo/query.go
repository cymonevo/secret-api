package repo

const (
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
