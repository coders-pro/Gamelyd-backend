package models

import (
	"math/rand"
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
	Teams []Teams
	Fixtures
}

type Fixtures [][]Teams

type TournamentGroups []TournamentGroup

func (tg *TournamentGroup) PairTournamentGroups() TournamentGroup {
	var fixtures Fixtures
	for i := 0; i < len(tg.Teams); i++ {
		for j := i + 1; j < len(tg.Teams); j++ {
			fixtures = append(fixtures, []Teams{tg.Teams[i], tg.Teams[j]})
		}
	}

	tg.Fixtures = fixtures
	return *tg

}

func (a *TournamentGroup) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range a.Fixtures {
		newPosition := r.Intn(len(a.Fixtures) - 1)

		a.Fixtures[i], a.Fixtures[newPosition] = a.Fixtures[newPosition], a.Fixtures[i]
	}
}
