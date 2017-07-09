package models

import (
	"github.com/jinzhu/gorm"

	//"fmt"
)

type Tournament struct {
	gorm.Model

	Matches []Match `gorm:"ForeignKey:TournamentID"`
	Teams []Team `gorm:"ForeignKey:TournamentID"`
}

func (tour *Tournament) CreateTeams(db *gorm.DB, request TeamRequest) []Team {
	teamRaw := request.Teams
	var teams []Team
	for _, raw := range teamRaw {
		team := Team {
			TournamentID: tour.ID,
			Member1ID: uint(raw.Member1_id),
			Member2ID: uint(raw.Member2_id),
		}
		db.Create(&team)
		teams = append(teams, team)
	}
	return teams
}

func (tour *Tournament) CreateMatches(db *gorm.DB, teams []Team) []Match {
	nTeam := len(teams)
	// Add a dummy teams if number of teams is odd
	if nTeam % 2 == 1 {
		nTeam += 1
	}

	candidates := make([]int, nTeam)
	for i := range candidates {
		candidates[i] = i
	}

	var matches []Match
	// For each round
	for round := 1; round < nTeam; round++ {
		// Team i in upper row
		for i := 0; i < nTeam / 2; i++ {
			// Team j in lower row
			j := nTeam - i - 1

			team1 := candidates[i]
			team2 := candidates[j]
			// Change Home/Away
			if round % 2 == 0 {
				team1, team2 = team2, team1
			}

			// Next if i or j is dummy team
			if len(teams) % 2 == 1 && (team1 == nTeam - 1 || team2 == nTeam - 1) {
				continue
			}

			match := Match {
				TournamentID: tour.ID,
				Team1ID: teams[team1].ID,
				Team2ID: teams[team2].ID,
			}

			db.Create(&match)
			matches = append(matches, match)
		}

		// Rotate candidate, follow the rule
		// https://stackoverflow.com/questions/6648512/scheduling-algorithm-for-a-round-robin-tournament
		tail := append(candidates[nTeam-1:], candidates[1:nTeam-1]...)
		candidates = append(candidates[:1], tail...)
	}

	return matches
}
