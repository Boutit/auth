package store

const (
	createAccountSQL = `
		INSERT into account (
			username,
			name,
			city_lat
		) VALUES (
			:username,
			:name,
			:citylat
		)
	`
)