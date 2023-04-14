package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"math/rand"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	waProto "go.mau.fi/whatsmeow/binary/proto"
)

var lossPhrases = []string{
	"es mas malo que la infancia en siria!",
	"es peor que la morgana de biche!",
	"juega peor que el depor!",
	"esta una partida mas cerca del descenso, como el atleti!",
	"se merece un relojazo!",
	"deberia haber estudiao!",
	"tiene menos elo que pelo gibe!",
	"tira palla bobo!",
}

var winPhrases = []string{
	"ha chupado otro carrito!",
	"ha ganado de puto milagro!",
	"ha sobornado a negreira!",
	"parece VINI JR!",
	"va de relojazo en relojazo!",
	"se ha ganado unos kekos!",
	"ES IMPARABLE!",
	"ha despertado a pos!",
}

type ClientHandler struct {
	client      *whatsmeow.Client
	lolclient   *LolClient
	db          *gorm.DB
	playerCache map[string]map[string]string
	groupJID    types.JID
	adminJID    types.JID
}

func NewClientHandler(
	client *whatsmeow.Client,
	lolClient *LolClient,
	groupJID string,
	adminJID string,
) *ClientHandler {
	var accs []Account
	cache := map[string]map[string]string{}

	db, err := gorm.Open(sqlite.Open("botlevi.sqlite"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Account{})

	db.Find(&accs)
	for _, acc := range accs {
		matchId, _ := lolClient.GetLastMatchId(acc.Puuid)
		cache[acc.Puuid] = map[string]string{"lastMatchId": matchId}
	}

	chat, _ := types.ParseJID(groupJID)
	admin, _ := types.ParseJID(adminJID)

	return &ClientHandler{client, lolClient, db, cache, chat, admin}
}

func (c *ClientHandler) checkForNewMatches() {
	rand.Seed(time.Now().UnixNano())
	for range time.Tick(time.Second * 30) {
		for puuid, value := range c.playerCache {
			matchId, _ := c.lolclient.GetLastMatchId(puuid)
			if matchId != value["lastMatchId"] {
				value["lastMatchId"] = matchId
				match, err := c.lolclient.GetMatchById(matchId)

				if err != nil {
					continue
				}

				for _, v := range match.Info.Participants {
					if v.Puuid == puuid {
						if v.Win {
							c.client.Log.Infof("DR EARLY HA GANADO")
							c.SendMessage(
								fmt.Sprintf(
									"Bot: Ring Ring, VICTORIA! %s %s \n CAMPEON: %s \n DURACION: %d minutos \n STATS: %d/%d/%d \n DAÑO REALIZADO: %d \n HA PINGEADO UN TOTAL DE: %d \n ",
									v.SummonerName,
									winPhrases[rand.Intn(len(winPhrases)-0+1)+0],
									v.ChampionName,
									v.TimePlayed/60,
									v.Kills,
									v.Deaths,
									v.Assists,
									v.TotalDamageDealtToChampions,
									(v.HoldPings + v.PushPings + v.CommandPings + v.BasicPings + v.AssistMePings + v.OnMyWayPings + v.BaitPings + v.GetBackPings),
								),
							)
						} else {
							c.client.Log.Infof("DR EARLY HA PERDIDO")
							c.SendMessage(
								fmt.Sprintf("Bot: Ring Ring, DERROTA! %s %s \n CAMPEON: %s \n DURACION: %d minutos \n STATS: %d/%d/%d \n DAÑO REALIZADO: %d\n HA PINGEADO UN TOTAL DE: %d \n ",
									v.SummonerName,
									lossPhrases[rand.Intn(len(lossPhrases)-0+1)+0],
									v.ChampionName,
									v.TimePlayed/60,
									v.Kills,
									v.Deaths,
									v.Assists,
									v.TotalDamageDealtToChampions,
									(v.HoldPings + v.PushPings + v.CommandPings + v.BasicPings + v.AssistMePings + v.OnMyWayPings + v.BaitPings + v.GetBackPings),
								),
							)
						}

					}
				}
			}
		}

	}
}

func (c *ClientHandler) AddEventHandlers() {
	c.client.AddEventHandler(c.AddAccount)
	go c.checkForNewMatches()

}

func (c *ClientHandler) SendMessage(msg string) {
	c.client.SendMessage(
		context.Background(),
		c.groupJID,
		&waProto.Message{Conversation: proto.String(msg)},
	)
}

func (c *ClientHandler) retrievePlayerInfo(summonerName string) (Account, error) {
	summoner, err := c.lolclient.GetSummonerByName(summonerName)

	if err != nil {
		return Account{}, nil
	}

	return Account{
		Name:      summoner.Name,
		Accountid: summoner.AccountID,
		Id:        summoner.ID,
		Puuid:     summoner.Puuid,
	}, nil
}

func (c *ClientHandler) AddAccount(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		msg := strings.ToLower(v.Message.GetConversation())
		if strings.HasPrefix(msg, ".") {
			// param 0 should be the command
			// param 1 should be the value
			params := strings.Split(msg, " ")
			if len(params) == 0 {
				return
			}

			switch params[0] {
			case ".addaccount":
				if v.Info.Sender.String() != c.adminJID.String() {
					return
				}

				acc, err := c.retrievePlayerInfo(strings.Join(params[1:], ""))

				if err != nil {
					c.client.Log.Errorf("%s", err)
					return
				}

				c.db.Create(&acc)
				c.client.Log.Infof("Added account: %+v\n", acc)
				matchId, _ := c.lolclient.GetLastMatchId(acc.Puuid)
				c.playerCache[acc.Puuid] = map[string]string{"lastMatchId": matchId}
				var accs []string
				for k := range c.playerCache {
					accs = append(accs, k)
				}
				c.SendMessage(fmt.Sprintf("Tracking new account.. current accounts: %s", strings.Join(accs, "")))
			}
		}

	}
}
