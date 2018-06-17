package services

import (
	"database/sql"
)

// HeroService provides a service to retrieve heroes
type HeroService struct {
	db *sql.DB
}

// Hero represents the Hero Model
type Hero struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Portrait   string `json:"portrait"`
	Tier       int    `json:"tier"`
	AttackType string `json:"attackType"`
	Role       string `json:"role"`
}

// GetHeroService instantiates and returns a new HeroService
func GetHeroService(db *sql.DB) *HeroService {
	return &HeroService{db: db}
}

// GetHeroes gets all the heroes available with its role
func (hs *HeroService) GetHeroes() (*[]Hero, error) {
	getHeroesStmt, err := hs.db.Prepare(`SELECT 
		h.id, h.name, h.portrait, h.tier, h.attack_type, r.name
		FROM heroes h 
		INNER JOIN roles r on h.role_id = r.id
		ORDER BY h.tier, r.id`)
	if err != nil {
		return nil, err
	}

	defer getHeroesStmt.Close()

	rows, err := getHeroesStmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	heroes := []Hero{}

	for rows.Next() {
		var h Hero
		err := rows.Scan(&h.ID, &h.Name, &h.Portrait, &h.Tier, &h.AttackType, &h.Role)
		if err != nil {
			return nil, err
		}
		heroes = append(heroes, h)
	}
	return &heroes, nil
}
