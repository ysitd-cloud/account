package model

import "database/sql"

type Provider struct {
	ID           string `json:"id"`
	ClientID     string `json:"-"`
	ClientSecret string `json:"-"`
	Scopes       string `json:"scopes"`
	Name         string `json:"name"`
}

func ListProvider(db *sql.DB) ([]*Provider, error) {
	sql := "SELECT id, client_id, client_secret, scopes, name FROM connect"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var providers []*Provider
	for rows.Next() {
		var id, clientId, clientSecret, scopes, name string
		if err := rows.Scan(&id, &clientId, &clientSecret, &scopes, &name); err != nil {
			return nil, err
		}

		provider := &Provider{
			ID:           id,
			ClientID:     clientId,
			ClientSecret: clientSecret,
			Scopes:       scopes,
			Name:         name,
		}

		providers = append(providers, provider)
	}

	return providers, nil
}
