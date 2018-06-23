package models

import (
	"github.com/barretodaniel/hots-tier-list-api/db"
)

// Role represents the Hero's Role
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetRoles returns all the roles
func GetRoles() (*[]Role, error) {
	getRolesStmt, err := db.Get().Prepare(`SELECT * FROM roles`)

	if err != nil {
		return nil, err
	}

	defer getRolesStmt.Close()

	rows, err := getRolesStmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	roles := []Role{}

	for rows.Next() {
		var r Role
		err := rows.Scan(&r.ID, &r.Name)
		if err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}
	return &roles, nil
}
