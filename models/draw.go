package models

import (
	"math/rand"
	"sort"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	UserName     string `json:"UserName" validate:"required"`
	GameUserName string `json:"GameUserName" validate:"required"`
	User_id      string `json:"User_id" validate:"required"`
	Email        string `json:"Email" validate:"required"`
}

type Teams struct {
	TeamName        string   `json:"TeamName"`
	Players         []Player `json:"Players"`
	Icon            string   `json:"Icon"`
	Points          int      `json:"Points"`
	GoalsScored     int      `json:"GoalsScored"`
	GoalsReceived   int      `json:"GoalsReceived"`
	GoalsDifference int      `json:"GoalsDifference"`
}

type BRTeams struct {
	TeamName    string   `json:"TeamName" validate:"required"`
	Players     []Player `json:"Players" validate:"required"`
	Wins        int      `json:"Wins"`
	Kills       int      `json:"Kills"`
	FirstBloods int      `json:"FirstBloods"`
}

type Draw struct {
	ID           primitive.ObjectID `bson:"_id" validate:"required"`
	Team1        Teams              `json:"Team1"`
	Team2        Teams              `json:"Team2"`
	Created_at   time.Time          `json:"Created_at" validate:"required"`
	Updated_at   time.Time          `json:"Updated_at" validate:"required"`
	TournamentId string             `json:"TournamentId" validate:"required"`
	DrawId       string             `json:"DrawId" validate:"required"`
	Stage        int                `json:"Stage" validate:"required"`
	Winner       string             `json:"Winner" validate:"eq=Team1|eq=Team2"`
	BRWinner     string             `json:"BRWinner"`
	Time         string             `json:"Time"`
	Date         string             `json:"Date"`
	Team1Score   int                `json:"Team1Score"`
	Team2Score   int                `json:"Team2Score"`
	Link         string             `json:"Link"`
	Group        string             `json:"Group"`
	BRTeams      []BRTeams          `json:"BRTeams"`
	IsPlayed     bool               `json:"IsPlayed"`
}

type AllDraws []Draw

// type DrawPairedDraw [][2]AllDraws
type AllTeams []Teams
type PairedDraw [][2]AllTeams

func (d *Draw) InitDraw() {

	d.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	d.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	d.ID = primitive.NewObjectID()
	d.DrawId = d.ID.Hex()

}

func (a AllTeams) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range a {
		newPosition := r.Intn(len(a) - 1)

		a[i], a[newPosition] = a[newPosition], a[i]
	}
}

func (tg TournamentGroups) ShuffleTournamentGroups() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Loop through the slice from the end to the beginning
	for i := len(tg) - 1; i > 0; i-- {
		// Pick a random index from 0 to i using the new random object
		j := r.Intn(i + 1)

		// Swap tg[i].Teams[1] with tg[j].Teams[1]
		tg[i].Teams[1], tg[j].Teams[1] = tg[j].Teams[1], tg[i].Teams[1]
	}
}

func (a AllTeams) GetTopTwo(name string) TournamentGroup {
	// Sort by Points first
	sort.SliceStable(a, func(i, j int) bool {
		return a[i].Points > a[j].Points
	})

	// Check if top 2 teams have the same points
	if len(a) > 2 && a[0].Points == a[1].Points {
		// Find the range where teams have the same points as the top team
		end := 2
		for end < len(a) && a[end].Points == a[0].Points {
			end++
		}

		// Sort that range by GoalsDifference
		sort.SliceStable(a[:end], func(i, j int) bool {
			return a[i].GoalsDifference > a[j].GoalsDifference
		})

		// If GoalsDifference is also the same, sort by GoalsScored
		if a[0].GoalsDifference == a[1].GoalsDifference {
			sort.SliceStable(a[:end], func(i, j int) bool {
				return a[i].GoalsScored > a[j].GoalsScored
			})

			// If GoalsScored is also the same, sort by alphabetical order
			if a[0].GoalsScored == a[1].GoalsScored {
				sort.SliceStable(a[:end], func(i, j int) bool {
					return strings.ToLower(a[i].TeamName) < strings.ToLower(a[j].TeamName)
				})
			}
		}
	}

	tg := TournamentGroup{
		Name:  name,
		Teams: a[:2],
	}

	// Return the top 2 teams
	return tg
}

// Pair method to pair up elements in the RegisterTournamentSlice

func (a AllTeams) Pair() PairedDraw {
	// Group the slice into pairs
	pairs := make(PairedDraw, 0)

	for i := 0; i < len(a); i += 2 {
		if i+1 < len(a) {
			// Create a pair of slices, each containing one element
			pair := [2]AllTeams{
				a[i : i+1],   // First element as a slice
				a[i+1 : i+2], // Second element as a slice
			}
			pairs = append(pairs, pair)
		} else {
			// Handle the case where there's an odd number of elements
			pair := [2]AllTeams{
				a[i : i+1], // Last element as a slice
				nil,        // No second element
			}
			pairs = append(pairs, pair)
		}
	}

	return pairs
}

func (p PairedDraw) Generate1V1DrawsFromPairs(tournament Tournament) []Draw {
	allDraws := []Draw{}
	// loop draws paired draws and create new draws
	for _, pairedDraw := range p {
		newDraw := Draw{}

		newDraw.Team1.Players = pairedDraw[0][0].Players
		newDraw.Team1.TeamName = pairedDraw[0][0].TeamName
		newDraw.Team1.Icon = pairedDraw[0][0].Icon

		if len(pairedDraw[1]) == 1 {
			newDraw.Team2.Players = pairedDraw[1][0].Players
			newDraw.Team2.TeamName = pairedDraw[1][0].TeamName
			newDraw.Team2.Icon = pairedDraw[1][0].Icon
		} else {
			// add default team for automatic qualification
			newDraw.Team2.Players = nil
			newDraw.Team2.TeamName = "Automatic Qualification"
			newDraw.Team2.Icon = ""
		}

		newDraw.ID = primitive.NewObjectID()
		newDraw.Stage = tournament.Stage + 1
		newDraw.DrawId = newDraw.ID.Hex()
		newDraw.TournamentId = tournament.TournamentId

		allDraws = append(allDraws, newDraw)
	}

	return allDraws
}

// extract all winning teams from draw matches
func (a AllDraws) ExtractTeam() AllTeams {
	teams := AllTeams{}

	for _, oneTeam := range a {
		team := Teams{}
		// if winner is team 2
		if oneTeam.Winner == "Team2" {
			team = oneTeam.Team2
		} else {
			// if winner is team one or there is an automatic qualification
			team = oneTeam.Team1
		}
		teams = append(teams, team)
	}
	return teams
}

func (ts *AllTeams) UpdateTeamScore(name string, gf, ga, points int) {
	for i := range *ts {
		if (*ts)[i].TeamName == name {
			(*ts)[i].Points += points
			(*ts)[i].GoalsScored += gf
			(*ts)[i].GoalsReceived += ga
			(*ts)[i].GoalsDifference = (*ts)[i].GoalsScored - (*ts)[i].GoalsReceived
			break
		}
	}

	sort.SliceStable(*ts, func(i, j int) bool {
		if (*ts)[i].Points != (*ts)[j].Points {
			return (*ts)[i].Points > (*ts)[j].Points
		}
		if (*ts)[i].GoalsDifference != (*ts)[j].GoalsDifference {
			return (*ts)[i].GoalsDifference > (*ts)[j].GoalsDifference
		}
		if (*ts)[i].GoalsScored != (*ts)[j].GoalsScored {
			return (*ts)[i].GoalsScored > (*ts)[j].GoalsScored
		}
		return (*ts)[i].TeamName < (*ts)[j].TeamName
	})
}
