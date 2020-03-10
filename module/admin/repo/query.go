package repo

const (
	insertAppQuery = `
		INSERT INTO scr_app (app_id, secret)
		VALUES (:app_id, :secret)`

	getAppQuery = `
		SELECT (app_id, secret)
		FROM scr_app
		WHERE app_id = $1`
)
