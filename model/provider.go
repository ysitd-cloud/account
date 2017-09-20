package model

import (
	"database/sql"
	"strings"
)

type Provider struct {
	ID           string   `json:"id"`
	ClientID     string   `json:"-"`
	ClientSecret string   `json:"-"`
	Scopes       []string `json:"scopes"`
	Name         string   `json:"name"`
	RedirectURL  string   `json:"redirect_url"`
	Icon         string   `json:"icon"`
}

func ListProvider(db *sql.DB) ([]*Provider, error) {
	query := "SELECT id, client_id, client_secret, scopes, name, redirect_url, icon FROM connect"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var providers []*Provider
	for rows.Next() {
		var id, clientId, clientSecret, scopes, name, redirectURL, icon string
		if err := rows.Scan(&id, &clientId, &clientSecret, &scopes, &name, &redirectURL, &icon); err != nil {
			return nil, err
		}

		provider := &Provider{
			ID:           id,
			ClientID:     clientId,
			ClientSecret: clientSecret,
			Scopes:       processScope(scopes),
			Name:         name,
			RedirectURL:  redirectURL,
			Icon:         icon,
		}

		providers = append(providers, provider)
	}

	return providers, nil
}

func GetProviderByID(db *sql.DB, id string) (*Provider, error) {
	query := "SELECT client_id, client_secret, scopes, name, redirect_url, icon FROM connect WHERE id = $1"
	row := db.QueryRow(query, id)
	var clientId, clientSecret, scope, name, redirectURL, icon string
	if err := row.Scan(&clientId, &clientSecret, &scope, &name, &redirectURL, &icon); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	provider := &Provider{
		ID:           id,
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       processScope(scope),
		Name:         name,
		RedirectURL:  redirectURL,
		Icon:         icon,
	}

	return provider, nil
}

func processScope(scope string) []string {
	return strings.Split(scope, "|")
}
