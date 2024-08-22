package models

import (
	"math"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tournament struct {
	ID                 primitive.ObjectID `bson:"_id" validate:"required"`
	Name               *string            `json:"Name" validate:"required,min=2,max=100"`
	GameName           *string            `json:"GameName" validate:"required,min=2,max=100"`
	Icon               *string            `json:"icon" validate:"required,min=1,max=100"`
	TournamentType     *string            `json:"TournamentType" validate:"required,eq=PUBLIC|eq=PRIVATE"`
	Payment            *string            `json:"Payment" validate:"required,eq=FREE|eq=PAID|eq=SPONSORED"`
	TournamentMode     *string            `json:"TournamentMode" validate:"required,eq=MULTIPLAYER|eq=BATTLEROYALE|eq=GROUPS"`
	TournamentSize     *int               `json:"TournamentSize" validate:"required"`
	Team               *string            `json:"Team" validate:"required,eq=SINGLE|eq=DUO|eq=SQUAD"`
	Shuffle            *string            `json:"Shuffle" validate:"required,eq=MANUAL|eq=AUTOMATIC"`
	Created_at         time.Time          `json:"Created_at" validate:"required"`
	Updated_at         time.Time          `json:"Updated_at"`
	TournamentId       string             `json:"TournamentId" validate:"required"`
	User_id            string             `json:"User_id" validate:"required"`
	Active             bool               `json:"Active"`
	IsDeleted          bool               `json:"IsDeleted"`
	IsSuspended        bool               `json:"IsSuspended"`
	Start              bool               `json:"Start"`
	Link               string             `json:"Link"`
	Date               string             `json:"Date" validate:"required"`
	IsPaid             bool               `json:"IsPaid"`
	IsDrawn            bool               `json:"IsDrawn"`
	RefNumber          string             `json:"RefNumber"`
	PaymentChannel     string             `json:"PaymentChannel"`
	Amount             int                `json:"Amount"`
	RegistrationAmount int                `json:"RegistrationAmount"`
	Note               string             `json:"Note" validate:"required"`
	AcceptedInvites    []string           `json:"AcceptedInvites"`
	Platform           string             `json:"Platform"`
	Winner             Teams              `json:"Winner"`
	Stage              int                `json:"Stage"`
	Groups             TournamentGroups   `json:"Groups"`
	PointSystem        PointSystem        `json:"PointSystem"`
}

type PointSystem struct {
	Win  int `json:"Win"`
	Draw int `json:"Draw"`
}

type TournamentGroup struct {
	Name  string `json:"GroupName"`
	Teams AllTeams
}

type TournamentGroups []TournamentGroup

func (tg *TournamentGroup) PairTournamentGroups(tournamentId string) AllDraws {
	var drawSlice AllDraws
	var draw Draw

	for i := 0; i < len(tg.Teams); i++ {
		for j := i + 1; j < len(tg.Teams); j++ {
			draw.Team1 = tg.Teams[i]
			draw.Team2 = tg.Teams[j]
			draw.InitDraw()
			draw.TournamentId = tournamentId
			draw.Group = tg.Name
			drawSlice = append(drawSlice, draw)
		}
	}

	return drawSlice

}

func GroupTeams(tournamentParticipants AllRegTeams) TournamentGroups {
	var groups TournamentGroups

	maxNumberPerGroup := 4
	minNumberPerGroup := 3
	groupNumber := int(math.Ceil(float64(len(tournamentParticipants)) / float64(maxNumberPerGroup)))
	allTeams := tournamentParticipants.ExtractTeam()
	allTeams.Shuffle()

	for i := 0; i < groupNumber; i++ {

		currentGroup := maxNumberPerGroup * (i + 1)

		if i == groupNumber-1 {

			group := TournamentGroup{
				Name:  "Group" + " " + strconv.Itoa(i+1),
				Teams: allTeams[currentGroup-maxNumberPerGroup:],
			}

			count := 1
			for {
				if len(group.Teams) >= minNumberPerGroup {
					break
				} else {
					teamGroups := groups[len(groups)-count].Teams

					if len(teamGroups) > minNumberPerGroup {
						group.Teams = append(group.Teams, teamGroups[len(teamGroups)-1])
						groups[len(groups)-count].Teams = teamGroups[:len(teamGroups)-1]
					} else {
						count++
					}

				}

				if count > len(groups) {
					break
				}
			}

			groups = append(groups, group)

		} else {
			group := TournamentGroup{
				Name:  "Group" + " " + strconv.Itoa(i+1),
				Teams: allTeams[currentGroup-maxNumberPerGroup : currentGroup],
			}

			groups = append(groups, group)
		}

	}

	return groups
}
