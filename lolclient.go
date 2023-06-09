package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	lolRequestUrl   = "https://euw1.api.riotgames.com"
	matchRequestUrl = "https://europe.api.riotgames.com"
)

type LolClient struct {
	apiKey     string
	httpClient *http.Client
}

func NewLolClient(apiKey string) *LolClient {
	return &LolClient{apiKey, &http.Client{}}
}

func (c *LolClient) GetLeagueBySummonerId(summonerId string) ([]League, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		strings.Join(
			[]string{lolRequestUrl, "/lol/league/v4/entries/by-summoner/", summonerId},
			"",
		),
		nil,
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-Riot-Token", c.apiKey)

	res, err := c.httpClient.Do(req)

	if err != nil {
		return []League{}, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return []League{}, err
	}

	var leagues []League
	err = json.Unmarshal(body, &leagues)

	if err != nil {
		return []League{}, err
	}

	return leagues, nil
}

func (c *LolClient) GetSummonerByName(summonerName string) (Summoner, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		strings.Join(
			[]string{lolRequestUrl, "/lol/summoner/v4/summoners/by-name/", summonerName},
			"",
		),
		nil,
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-Riot-Token", c.apiKey)

	res, err := c.httpClient.Do(req)

	if err != nil {
		return Summoner{}, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return Summoner{}, err
	}

	var summoner Summoner
	err = json.Unmarshal(body, &summoner)

	if err != nil {
		return Summoner{}, err
	}

	return summoner, nil
}

func (c *LolClient) GetLastMatchId(puuid string) (string, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		strings.Join(
			[]string{
				matchRequestUrl,
				"/lol/match/v5/matches/by-puuid/",
				puuid,
				"/ids?start=0&count=1",
			},
			"",
		),
		nil,
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-Riot-Token", c.apiKey)

	res, err := c.httpClient.Do(req)

	if err != nil {
		return "", err
	}

	var ids []string
	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &ids)

	return ids[0], err

}

func (c *LolClient) GetMatchById(id string) (Match, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		strings.Join([]string{matchRequestUrl, "/lol/match/v5/matches/", id}, ""),
		nil,
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-Riot-Token", c.apiKey)

	res, err := c.httpClient.Do(req)

	if err != nil {
		return Match{}, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return Match{}, err
	}

	var match Match
	err = json.Unmarshal(body, &match)

	if err != nil {
		return Match{}, err
	}

	return match, nil

}

//////////////////

//////////////////
//////////////////

//////////////////
//////////////////

//////////////////
//////////////////

// ////////////////
type Summoner struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

type League struct {
	LeagueID     string `json:"leagueId"`
	QueueType    string `json:"queueType"`
	Tier         string `json:"tier"`
	Rank         string `json:"rank"`
	SummonerID   string `json:"summonerId"`
	SummonerName string `json:"summonerName"`
	LeaguePoints int    `json:"leaguePoints"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	Veteran      bool   `json:"veteran"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshBlood"`
	HotStreak    bool   `json:"hotStreak"`
	MiniSeries   struct {
		Target   int    `json:"target"`
		Wins     int    `json:"wins"`
		Losses   int    `json:"losses"`
		Progress string `json:"progress"`
	} `json:"miniSeries"`
}

type Match struct {
	Metadata struct {
		DataVersion  string   `json:"dataVersion"`
		MatchID      string   `json:"matchId"`
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		GameCreation       int64  `json:"gameCreation"`
		GameDuration       int    `json:"gameDuration"`
		GameEndTimestamp   int64  `json:"gameEndTimestamp"`
		GameID             int64  `json:"gameId"`
		GameMode           string `json:"gameMode"`
		GameName           string `json:"gameName"`
		GameStartTimestamp int64  `json:"gameStartTimestamp"`
		GameType           string `json:"gameType"`
		GameVersion        string `json:"gameVersion"`
		MapID              int    `json:"mapId"`
		Participants       []struct {
			AllInPings    int `json:"allInPings"`
			AssistMePings int `json:"assistMePings"`
			Assists       int `json:"assists"`
			BaitPings     int `json:"baitPings"`
			BaronKills    int `json:"baronKills"`
			BasicPings    int `json:"basicPings"`
			BountyLevel   int `json:"bountyLevel"`
			Challenges    struct {
				One2AssistStreakCount                    int     `json:"12AssistStreakCount"`
				AbilityUses                              int     `json:"abilityUses"`
				AcesBefore15Minutes                      int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                 int     `json:"alliedJungleMonsterKills"`
				BaronTakedowns                           int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount           int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                               int     `json:"bountyGold"`
				BuffsStolen                              int     `json:"buffsStolen"`
				CompleteSupportQuestInTime               int     `json:"completeSupportQuestInTime"`
				ControlWardsPlaced                       int     `json:"controlWardsPlaced"`
				DamagePerMinute                          float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage              float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                     int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                      int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow               int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                               int     `json:"doubleAces"`
				DragonTakedowns                          int     `json:"dragonTakedowns"`
				EarlyLaningPhaseGoldExpAdvantage         int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                int     `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul         int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                    int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations             int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                  int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler         int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn   int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                        int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite            int     `json:"epicMonsterStolenWithoutSmite"`
				FlawlessAces                             int     `json:"flawlessAces"`
				FullTeamTakedown                         int     `json:"fullTeamTakedown"`
				GameLength                               float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                            float64 `json:"goldPerMinute"`
				HadOpenNexus                             int     `json:"hadOpenNexus"`
				ImmobilizeAndKillWithAlly                int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                         int     `json:"initialBuffCount"`
				InitialCrabCount                         int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                  int     `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster   int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall        int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                      float64 `json:"kda"`
				KillAfterHiddenWithAlly                  int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                        float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived    int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillsNearEnemyTurret                     int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner      int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack          int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                      int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster             int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                  int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage              int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                           int     `json:"legendaryCount"`
				LostAnInhibitor                          int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent             int     `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                           int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                 int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent              int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                        int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount               int     `json:"multiTurretRiftHeraldCount"`
				Multikills                               int     `json:"multikills"`
				MultikillsAfterAggressiveFlash           int     `json:"multikillsAfterAggressiveFlash"`
				MythicItemUsed                           int     `json:"mythicItemUsed"`
				OuterTurretExecutesBefore10Minutes       int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                         int     `json:"outnumberedKills"`
				OutnumberedNexusKill                     int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                  int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                              int     `json:"perfectGame"`
				PickKillWithAlly                         int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                int     `json:"playedChampSelectPosition"`
				PoroExplosions                           int     `json:"poroExplosions"`
				QuickCleanse                             int     `json:"quickCleanse"`
				QuickFirstTurret                         int     `json:"quickFirstTurret"`
				QuickSoloKills                           int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                      int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                        int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                         int     `json:"scuttleCrabKills"`
				SkillshotsDodged                         int     `json:"skillshotsDodged"`
				SkillshotsHit                            int     `json:"skillshotsHit"`
				SnowballsHit                             int     `json:"snowballsHit"`
				SoloBaronKills                           int     `json:"soloBaronKills"`
				SoloKills                                int     `json:"soloKills"`
				StealthWardsPlaced                       int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount               int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight          int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                    int     `json:"takedownOnFirstTurret"`
				Takedowns                                int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage      int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn         int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                   int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                        int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                 int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                           int     `json:"teamBaronKills"`
				TeamDamagePercentage                     float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                     int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                      int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                  int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                        int     `json:"turretPlatesTaken"`
				TurretTakedowns                          int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald               int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount             int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                            int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent         float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                     float64 `json:"visionScorePerMinute"`
				WardTakedowns                            int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                   int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                             int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			ChampExperience             int    `json:"champExperience"`
			ChampLevel                  int    `json:"champLevel"`
			ChampionID                  int    `json:"championId"`
			ChampionName                string `json:"championName"`
			ChampionTransform           int    `json:"championTransform"`
			CommandPings                int    `json:"commandPings"`
			ConsumablesPurchased        int    `json:"consumablesPurchased"`
			DamageDealtToBuildings      int    `json:"damageDealtToBuildings"`
			DamageDealtToObjectives     int    `json:"damageDealtToObjectives"`
			DamageDealtToTurrets        int    `json:"damageDealtToTurrets"`
			DamageSelfMitigated         int    `json:"damageSelfMitigated"`
			DangerPings                 int    `json:"dangerPings"`
			Deaths                      int    `json:"deaths"`
			DetectorWardsPlaced         int    `json:"detectorWardsPlaced"`
			DoubleKills                 int    `json:"doubleKills"`
			DragonKills                 int    `json:"dragonKills"`
			EligibleForProgression      bool   `json:"eligibleForProgression"`
			EnemyMissingPings           int    `json:"enemyMissingPings"`
			EnemyVisionPings            int    `json:"enemyVisionPings"`
			FirstBloodAssist            bool   `json:"firstBloodAssist"`
			FirstBloodKill              bool   `json:"firstBloodKill"`
			FirstTowerAssist            bool   `json:"firstTowerAssist"`
			FirstTowerKill              bool   `json:"firstTowerKill"`
			GameEndedInEarlySurrender   bool   `json:"gameEndedInEarlySurrender"`
			GameEndedInSurrender        bool   `json:"gameEndedInSurrender"`
			GetBackPings                int    `json:"getBackPings"`
			GoldEarned                  int    `json:"goldEarned"`
			GoldSpent                   int    `json:"goldSpent"`
			HoldPings                   int    `json:"holdPings"`
			IndividualPosition          string `json:"individualPosition"`
			InhibitorKills              int    `json:"inhibitorKills"`
			InhibitorTakedowns          int    `json:"inhibitorTakedowns"`
			InhibitorsLost              int    `json:"inhibitorsLost"`
			Item0                       int    `json:"item0"`
			Item1                       int    `json:"item1"`
			Item2                       int    `json:"item2"`
			Item3                       int    `json:"item3"`
			Item4                       int    `json:"item4"`
			Item5                       int    `json:"item5"`
			Item6                       int    `json:"item6"`
			ItemsPurchased              int    `json:"itemsPurchased"`
			KillingSprees               int    `json:"killingSprees"`
			Kills                       int    `json:"kills"`
			Lane                        string `json:"lane"`
			LargestCriticalStrike       int    `json:"largestCriticalStrike"`
			LargestKillingSpree         int    `json:"largestKillingSpree"`
			LargestMultiKill            int    `json:"largestMultiKill"`
			LongestTimeSpentLiving      int    `json:"longestTimeSpentLiving"`
			MagicDamageDealt            int    `json:"magicDamageDealt"`
			MagicDamageDealtToChampions int    `json:"magicDamageDealtToChampions"`
			MagicDamageTaken            int    `json:"magicDamageTaken"`
			NeedVisionPings             int    `json:"needVisionPings"`
			NeutralMinionsKilled        int    `json:"neutralMinionsKilled"`
			NexusKills                  int    `json:"nexusKills"`
			NexusLost                   int    `json:"nexusLost"`
			NexusTakedowns              int    `json:"nexusTakedowns"`
			ObjectivesStolen            int    `json:"objectivesStolen"`
			ObjectivesStolenAssists     int    `json:"objectivesStolenAssists"`
			OnMyWayPings                int    `json:"onMyWayPings"`
			ParticipantID               int    `json:"participantId"`
			PentaKills                  int    `json:"pentaKills"`
			Perks                       struct {
				StatPerks struct {
					Defense int `json:"defense"`
					Flex    int `json:"flex"`
					Offense int `json:"offense"`
				} `json:"statPerks"`
				Styles []struct {
					Description string `json:"description"`
					Selections  []struct {
						Perk int `json:"perk"`
						Var1 int `json:"var1"`
						Var2 int `json:"var2"`
						Var3 int `json:"var3"`
					} `json:"selections"`
					Style int `json:"style"`
				} `json:"styles"`
			} `json:"perks"`
			PhysicalDamageDealt            int    `json:"physicalDamageDealt"`
			PhysicalDamageDealtToChampions int    `json:"physicalDamageDealtToChampions"`
			PhysicalDamageTaken            int    `json:"physicalDamageTaken"`
			ProfileIcon                    int    `json:"profileIcon"`
			PushPings                      int    `json:"pushPings"`
			Puuid                          string `json:"puuid"`
			QuadraKills                    int    `json:"quadraKills"`
			RiotIDName                     string `json:"riotIdName"`
			RiotIDTagline                  string `json:"riotIdTagline"`
			Role                           string `json:"role"`
			SightWardsBoughtInGame         int    `json:"sightWardsBoughtInGame"`
			Spell1Casts                    int    `json:"spell1Casts"`
			Spell2Casts                    int    `json:"spell2Casts"`
			Spell3Casts                    int    `json:"spell3Casts"`
			Spell4Casts                    int    `json:"spell4Casts"`
			Summoner1Casts                 int    `json:"summoner1Casts"`
			Summoner1ID                    int    `json:"summoner1Id"`
			Summoner2Casts                 int    `json:"summoner2Casts"`
			Summoner2ID                    int    `json:"summoner2Id"`
			SummonerID                     string `json:"summonerId"`
			SummonerLevel                  int    `json:"summonerLevel"`
			SummonerName                   string `json:"summonerName"`
			TeamEarlySurrendered           bool   `json:"teamEarlySurrendered"`
			TeamID                         int    `json:"teamId"`
			TeamPosition                   string `json:"teamPosition"`
			TimeCCingOthers                int    `json:"timeCCingOthers"`
			TimePlayed                     int    `json:"timePlayed"`
			TotalDamageDealt               int    `json:"totalDamageDealt"`
			TotalDamageDealtToChampions    int    `json:"totalDamageDealtToChampions"`
			TotalDamageShieldedOnTeammates int    `json:"totalDamageShieldedOnTeammates"`
			TotalDamageTaken               int    `json:"totalDamageTaken"`
			TotalHeal                      int    `json:"totalHeal"`
			TotalHealsOnTeammates          int    `json:"totalHealsOnTeammates"`
			TotalMinionsKilled             int    `json:"totalMinionsKilled"`
			TotalTimeCCDealt               int    `json:"totalTimeCCDealt"`
			TotalTimeSpentDead             int    `json:"totalTimeSpentDead"`
			TotalUnitsHealed               int    `json:"totalUnitsHealed"`
			TripleKills                    int    `json:"tripleKills"`
			TrueDamageDealt                int    `json:"trueDamageDealt"`
			TrueDamageDealtToChampions     int    `json:"trueDamageDealtToChampions"`
			TrueDamageTaken                int    `json:"trueDamageTaken"`
			TurretKills                    int    `json:"turretKills"`
			TurretTakedowns                int    `json:"turretTakedowns"`
			TurretsLost                    int    `json:"turretsLost"`
			UnrealKills                    int    `json:"unrealKills"`
			VisionClearedPings             int    `json:"visionClearedPings"`
			VisionScore                    int    `json:"visionScore"`
			VisionWardsBoughtInGame        int    `json:"visionWardsBoughtInGame"`
			WardsKilled                    int    `json:"wardsKilled"`
			WardsPlaced                    int    `json:"wardsPlaced"`
			Win                            bool   `json:"win"`
			Challenges0                    struct {
				One2AssistStreakCount                  int     `json:"12AssistStreakCount"`
				AbilityUses                            int     `json:"abilityUses"`
				AcesBefore15Minutes                    int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills               int     `json:"alliedJungleMonsterKills"`
				BaronTakedowns                         int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount         int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                             int     `json:"bountyGold"`
				BuffsStolen                            int     `json:"buffsStolen"`
				CompleteSupportQuestInTime             int     `json:"completeSupportQuestInTime"`
				ControlWardsPlaced                     int     `json:"controlWardsPlaced"`
				DamagePerMinute                        float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage            float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                   int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                    int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow             int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                             int     `json:"doubleAces"`
				DragonTakedowns                        int     `json:"dragonTakedowns"`
				EarlyLaningPhaseGoldExpAdvantage       int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding              int     `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul       int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                  int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations           int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler       int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                      int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite          int     `json:"epicMonsterStolenWithoutSmite"`
				FlawlessAces                           int     `json:"flawlessAces"`
				FullTeamTakedown                       int     `json:"fullTeamTakedown"`
				GameLength                             float64 `json:"gameLength"`
				GoldPerMinute                          float64 `json:"goldPerMinute"`
				HadOpenNexus                           int     `json:"hadOpenNexus"`
				HighestWardKills                       int     `json:"highestWardKills"`
				ImmobilizeAndKillWithAlly              int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                       int     `json:"initialBuffCount"`
				InitialCrabCount                       int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                float64 `json:"jungleCsBefore10Minutes"`
				JunglerKillsEarlyJungle                int     `json:"junglerKillsEarlyJungle"`
				JunglerTakedownsNearDamagedEpicMonster int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall      int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                    float64 `json:"kda"`
				KillAfterHiddenWithAlly                int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                      float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived  int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillsNearEnemyTurret                   int     `json:"killsNearEnemyTurret"`
				KillsOnLanersEarlyJungleAsJungler      int     `json:"killsOnLanersEarlyJungleAsJungler"`
				KillsOnRecentlyHealedByAramPack        int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                    int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster           int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill              int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes              int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage            int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                         int     `json:"legendaryCount"`
				LostAnInhibitor                        int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent           float64 `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                         int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent               int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent            float64 `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                      int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount             int     `json:"multiTurretRiftHeraldCount"`
				Multikills                             int     `json:"multikills"`
				MultikillsAfterAggressiveFlash         int     `json:"multikillsAfterAggressiveFlash"`
				MythicItemUsed                         int     `json:"mythicItemUsed"`
				OuterTurretExecutesBefore10Minutes     int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                       int     `json:"outnumberedKills"`
				OutnumberedNexusKill                   int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                            int     `json:"perfectGame"`
				PickKillWithAlly                       int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition              int     `json:"playedChampSelectPosition"`
				PoroExplosions                         int     `json:"poroExplosions"`
				QuickCleanse                           int     `json:"quickCleanse"`
				QuickFirstTurret                       int     `json:"quickFirstTurret"`
				QuickSoloKills                         int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                    int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                      int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                       int     `json:"scuttleCrabKills"`
				SkillshotsDodged                       int     `json:"skillshotsDodged"`
				SkillshotsHit                          int     `json:"skillshotsHit"`
				SnowballsHit                           int     `json:"snowballsHit"`
				SoloBaronKills                         int     `json:"soloBaronKills"`
				SoloKills                              int     `json:"soloKills"`
				StealthWardsPlaced                     int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount             int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight        int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                  int     `json:"takedownOnFirstTurret"`
				Takedowns                              int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage    int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn       int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                 int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                      int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain               int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                         int     `json:"teamBaronKills"`
				TeamDamagePercentage                   float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                   int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                    int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount              int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                      int     `json:"turretPlatesTaken"`
				TurretTakedowns                        int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald             int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount           int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                          int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent       float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                   float64 `json:"visionScorePerMinute"`
				WardTakedowns                          int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                 int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                           int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			Challenges1 struct {
				One2AssistStreakCount                    int     `json:"12AssistStreakCount"`
				AbilityUses                              int     `json:"abilityUses"`
				AcesBefore15Minutes                      int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                 int     `json:"alliedJungleMonsterKills"`
				BaronTakedowns                           int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount           int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                               int     `json:"bountyGold"`
				BuffsStolen                              int     `json:"buffsStolen"`
				CompleteSupportQuestInTime               int     `json:"completeSupportQuestInTime"`
				ControlWardsPlaced                       int     `json:"controlWardsPlaced"`
				DamagePerMinute                          float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage              float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                     int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                      int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow               int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                               int     `json:"doubleAces"`
				DragonTakedowns                          int     `json:"dragonTakedowns"`
				EarlyLaningPhaseGoldExpAdvantage         int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                int     `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul         int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                    int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations             int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                  int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler         int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn   int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                        int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite            int     `json:"epicMonsterStolenWithoutSmite"`
				FlawlessAces                             int     `json:"flawlessAces"`
				FullTeamTakedown                         int     `json:"fullTeamTakedown"`
				GameLength                               float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                            float64 `json:"goldPerMinute"`
				HadOpenNexus                             int     `json:"hadOpenNexus"`
				ImmobilizeAndKillWithAlly                int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                         int     `json:"initialBuffCount"`
				InitialCrabCount                         int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                  int     `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster   int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall        int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                      float64 `json:"kda"`
				KillAfterHiddenWithAlly                  int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                        float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived    int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillingSprees                            int     `json:"killingSprees"`
				KillsNearEnemyTurret                     int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner      int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack          int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                      int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster             int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                  int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage              int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                           int     `json:"legendaryCount"`
				LostAnInhibitor                          int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent             int     `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                           int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                 int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent              int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                        int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount               int     `json:"multiTurretRiftHeraldCount"`
				Multikills                               int     `json:"multikills"`
				MultikillsAfterAggressiveFlash           int     `json:"multikillsAfterAggressiveFlash"`
				MythicItemUsed                           int     `json:"mythicItemUsed"`
				OuterTurretExecutesBefore10Minutes       int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                         int     `json:"outnumberedKills"`
				OutnumberedNexusKill                     int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                  int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                              int     `json:"perfectGame"`
				PickKillWithAlly                         int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                int     `json:"playedChampSelectPosition"`
				PoroExplosions                           int     `json:"poroExplosions"`
				QuickCleanse                             int     `json:"quickCleanse"`
				QuickFirstTurret                         int     `json:"quickFirstTurret"`
				QuickSoloKills                           int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                      int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                        int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                         int     `json:"scuttleCrabKills"`
				SkillshotsDodged                         int     `json:"skillshotsDodged"`
				SkillshotsHit                            int     `json:"skillshotsHit"`
				SnowballsHit                             int     `json:"snowballsHit"`
				SoloBaronKills                           int     `json:"soloBaronKills"`
				SoloKills                                int     `json:"soloKills"`
				StealthWardsPlaced                       int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount               int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight          int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                    int     `json:"takedownOnFirstTurret"`
				Takedowns                                int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage      int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn         int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                   int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                        int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                 int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                           int     `json:"teamBaronKills"`
				TeamDamagePercentage                     float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                     int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                      int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                  int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                        int     `json:"turretPlatesTaken"`
				TurretTakedowns                          int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald               int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount             int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                            int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent         float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                     float64 `json:"visionScorePerMinute"`
				WardTakedowns                            int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                   int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                             int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			Challenges2 struct {
				One2AssistStreakCount                    int     `json:"12AssistStreakCount"`
				AbilityUses                              int     `json:"abilityUses"`
				AcesBefore15Minutes                      int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                 int     `json:"alliedJungleMonsterKills"`
				BaronTakedowns                           int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount           int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                               int     `json:"bountyGold"`
				BuffsStolen                              int     `json:"buffsStolen"`
				CompleteSupportQuestInTime               int     `json:"completeSupportQuestInTime"`
				ControlWardsPlaced                       int     `json:"controlWardsPlaced"`
				DamagePerMinute                          float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage              float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                     int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                      int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow               int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                               int     `json:"doubleAces"`
				DragonTakedowns                          int     `json:"dragonTakedowns"`
				EarlyLaningPhaseGoldExpAdvantage         int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                int     `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul         int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                    int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations             int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                  int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler         int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn   int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                        int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite            int     `json:"epicMonsterStolenWithoutSmite"`
				FlawlessAces                             int     `json:"flawlessAces"`
				FullTeamTakedown                         int     `json:"fullTeamTakedown"`
				GameLength                               float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                            float64 `json:"goldPerMinute"`
				HadOpenNexus                             int     `json:"hadOpenNexus"`
				ImmobilizeAndKillWithAlly                int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                         int     `json:"initialBuffCount"`
				InitialCrabCount                         int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                  int     `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster   int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall        int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                      float64 `json:"kda"`
				KillAfterHiddenWithAlly                  int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                        float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived    int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillsNearEnemyTurret                     int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner      int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack          int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                      int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster             int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                  int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage              int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                           int     `json:"legendaryCount"`
				LostAnInhibitor                          int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent             int     `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                           int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                 int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent              int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                        int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount               int     `json:"multiTurretRiftHeraldCount"`
				Multikills                               int     `json:"multikills"`
				MultikillsAfterAggressiveFlash           int     `json:"multikillsAfterAggressiveFlash"`
				OuterTurretExecutesBefore10Minutes       int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                         int     `json:"outnumberedKills"`
				OutnumberedNexusKill                     int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                  int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                              int     `json:"perfectGame"`
				PickKillWithAlly                         int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                int     `json:"playedChampSelectPosition"`
				PoroExplosions                           int     `json:"poroExplosions"`
				QuickCleanse                             int     `json:"quickCleanse"`
				QuickFirstTurret                         int     `json:"quickFirstTurret"`
				QuickSoloKills                           int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                      int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                        int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                         int     `json:"scuttleCrabKills"`
				SkillshotsDodged                         int     `json:"skillshotsDodged"`
				SkillshotsHit                            int     `json:"skillshotsHit"`
				SnowballsHit                             int     `json:"snowballsHit"`
				SoloBaronKills                           int     `json:"soloBaronKills"`
				SoloKills                                int     `json:"soloKills"`
				StealthWardsPlaced                       int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount               int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight          int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                    int     `json:"takedownOnFirstTurret"`
				Takedowns                                int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage      int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn         int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                   int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                        int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                 int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                           int     `json:"teamBaronKills"`
				TeamDamagePercentage                     float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                     int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                      int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                  int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                        int     `json:"turretPlatesTaken"`
				TurretTakedowns                          int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald               int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount             int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                            int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent         float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                     float64 `json:"visionScorePerMinute"`
				WardTakedowns                            int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                   int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                             int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			Challenges3 struct {
				One2AssistStreakCount                    int     `json:"12AssistStreakCount"`
				AbilityUses                              int     `json:"abilityUses"`
				AcesBefore15Minutes                      int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                 int     `json:"alliedJungleMonsterKills"`
				BaronTakedowns                           int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount           int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                               int     `json:"bountyGold"`
				BuffsStolen                              int     `json:"buffsStolen"`
				CompleteSupportQuestInTime               int     `json:"completeSupportQuestInTime"`
				ControlWardsPlaced                       int     `json:"controlWardsPlaced"`
				DamagePerMinute                          float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage              float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                     int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                      int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow               int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                               int     `json:"doubleAces"`
				DragonTakedowns                          int     `json:"dragonTakedowns"`
				EarlyLaningPhaseGoldExpAdvantage         int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                float64 `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul         int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                    int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations             int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                  int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler         int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn   int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                        int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite            int     `json:"epicMonsterStolenWithoutSmite"`
				FlawlessAces                             int     `json:"flawlessAces"`
				FullTeamTakedown                         int     `json:"fullTeamTakedown"`
				GameLength                               float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                            float64 `json:"goldPerMinute"`
				HadOpenNexus                             int     `json:"hadOpenNexus"`
				ImmobilizeAndKillWithAlly                int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                         int     `json:"initialBuffCount"`
				InitialCrabCount                         int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                  int     `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster   int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall        int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                      int     `json:"kda"`
				KillAfterHiddenWithAlly                  int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                        float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived    int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillsNearEnemyTurret                     int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner      int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack          int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                      int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster             int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                  int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage              int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                           int     `json:"legendaryCount"`
				LostAnInhibitor                          int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent             int     `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                           int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                 int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent              int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                        int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount               int     `json:"multiTurretRiftHeraldCount"`
				Multikills                               int     `json:"multikills"`
				MultikillsAfterAggressiveFlash           int     `json:"multikillsAfterAggressiveFlash"`
				OuterTurretExecutesBefore10Minutes       int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                         int     `json:"outnumberedKills"`
				OutnumberedNexusKill                     int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                  int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                              int     `json:"perfectGame"`
				PickKillWithAlly                         int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                int     `json:"playedChampSelectPosition"`
				PoroExplosions                           int     `json:"poroExplosions"`
				QuickCleanse                             int     `json:"quickCleanse"`
				QuickFirstTurret                         int     `json:"quickFirstTurret"`
				QuickSoloKills                           int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                      int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                        int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                         int     `json:"scuttleCrabKills"`
				SkillshotsDodged                         int     `json:"skillshotsDodged"`
				SkillshotsHit                            int     `json:"skillshotsHit"`
				SnowballsHit                             int     `json:"snowballsHit"`
				SoloBaronKills                           int     `json:"soloBaronKills"`
				SoloKills                                int     `json:"soloKills"`
				StealthWardsPlaced                       int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount               int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight          int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                    int     `json:"takedownOnFirstTurret"`
				Takedowns                                int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage      int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn         int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                   int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                        int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                 int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                           int     `json:"teamBaronKills"`
				TeamDamagePercentage                     float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                     int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                      int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                  int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                        int     `json:"turretPlatesTaken"`
				TurretTakedowns                          int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald               int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount             int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                            int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent         float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                     float64 `json:"visionScorePerMinute"`
				WardTakedowns                            int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                   int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                             int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			Challenges4 struct {
				One2AssistStreakCount                     int     `json:"12AssistStreakCount"`
				AbilityUses                               int     `json:"abilityUses"`
				AcesBefore15Minutes                       int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                  int     `json:"alliedJungleMonsterKills"`
				BaronTakedowns                            int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount            int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                                int     `json:"bountyGold"`
				BuffsStolen                               int     `json:"buffsStolen"`
				CompleteSupportQuestInTime                int     `json:"completeSupportQuestInTime"`
				ControlWardTimeCoverageInRiverOrEnemyHalf float64 `json:"controlWardTimeCoverageInRiverOrEnemyHalf"`
				ControlWardsPlaced                        int     `json:"controlWardsPlaced"`
				DamagePerMinute                           float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage               float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                      int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                       int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow                int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                                int     `json:"doubleAces"`
				DragonTakedowns                           int     `json:"dragonTakedowns"`
				EarlyLaningPhaseGoldExpAdvantage          int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                 int     `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul          int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                     int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations              int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                   int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler          int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn    int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                         int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite             int     `json:"epicMonsterStolenWithoutSmite"`
				FirstTurretKilled                         float64 `json:"firstTurretKilled"`
				FlawlessAces                              int     `json:"flawlessAces"`
				FullTeamTakedown                          int     `json:"fullTeamTakedown"`
				GameLength                                float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner  int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                             float64 `json:"goldPerMinute"`
				HadOpenNexus                              int     `json:"hadOpenNexus"`
				HighestChampionDamage                     int     `json:"highestChampionDamage"`
				ImmobilizeAndKillWithAlly                 int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                          int     `json:"initialBuffCount"`
				InitialCrabCount                          int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                   int     `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster    int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall         int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                       int     `json:"kda"`
				KillAfterHiddenWithAlly                   int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                         float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived     int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillsNearEnemyTurret                      int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner       int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack           int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                       int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster              int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                 int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                   int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                 int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage               int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                            int     `json:"legendaryCount"`
				LostAnInhibitor                           int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent              int     `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                            int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                  int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent               int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                         int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount                int     `json:"multiTurretRiftHeraldCount"`
				Multikills                                int     `json:"multikills"`
				MultikillsAfterAggressiveFlash            int     `json:"multikillsAfterAggressiveFlash"`
				MythicItemUsed                            int     `json:"mythicItemUsed"`
				OuterTurretExecutesBefore10Minutes        int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                          int     `json:"outnumberedKills"`
				OutnumberedNexusKill                      int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                   int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                               int     `json:"perfectGame"`
				PickKillWithAlly                          int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                 int     `json:"playedChampSelectPosition"`
				PoroExplosions                            int     `json:"poroExplosions"`
				QuickCleanse                              int     `json:"quickCleanse"`
				QuickFirstTurret                          int     `json:"quickFirstTurret"`
				QuickSoloKills                            int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                       int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                         int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                          int     `json:"scuttleCrabKills"`
				SkillshotsDodged                          int     `json:"skillshotsDodged"`
				SkillshotsHit                             int     `json:"skillshotsHit"`
				SnowballsHit                              int     `json:"snowballsHit"`
				SoloBaronKills                            int     `json:"soloBaronKills"`
				SoloKills                                 int     `json:"soloKills"`
				SoloTurretsLategame                       int     `json:"soloTurretsLategame"`
				StealthWardsPlaced                        int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount                int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight           int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                     int     `json:"takedownOnFirstTurret"`
				Takedowns                                 int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage       int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn          int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                    int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                         int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                  int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                            int     `json:"teamBaronKills"`
				TeamDamagePercentage                      float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                      int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                       int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                 int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                   int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                         int     `json:"turretPlatesTaken"`
				TurretTakedowns                           int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald                int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount              int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                             int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent          float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                      float64 `json:"visionScorePerMinute"`
				WardTakedowns                             int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                    int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                              int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			Challenges5 struct {
				One2AssistStreakCount                     int     `json:"12AssistStreakCount"`
				AbilityUses                               int     `json:"abilityUses"`
				AcesBefore15Minutes                       int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                  int     `json:"alliedJungleMonsterKills"`
				BaronTakedowns                            int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount            int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                                int     `json:"bountyGold"`
				BuffsStolen                               int     `json:"buffsStolen"`
				CompleteSupportQuestInTime                int     `json:"completeSupportQuestInTime"`
				ControlWardTimeCoverageInRiverOrEnemyHalf float64 `json:"controlWardTimeCoverageInRiverOrEnemyHalf"`
				ControlWardsPlaced                        int     `json:"controlWardsPlaced"`
				DamagePerMinute                           float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage               float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                      int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                       int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow                int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                                int     `json:"doubleAces"`
				DragonTakedowns                           int     `json:"dragonTakedowns"`
				EarliestDragonTakedown                    float64 `json:"earliestDragonTakedown"`
				EarlyLaningPhaseGoldExpAdvantage          int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                 int     `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul          int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                     int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations              int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                   int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler          int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn    int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                         int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite             int     `json:"epicMonsterStolenWithoutSmite"`
				FirstTurretKilled                         float64 `json:"firstTurretKilled"`
				FlawlessAces                              int     `json:"flawlessAces"`
				FullTeamTakedown                          int     `json:"fullTeamTakedown"`
				GameLength                                float64 `json:"gameLength"`
				GoldPerMinute                             float64 `json:"goldPerMinute"`
				HadOpenNexus                              int     `json:"hadOpenNexus"`
				ImmobilizeAndKillWithAlly                 int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                          int     `json:"initialBuffCount"`
				InitialCrabCount                          int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                   float64 `json:"jungleCsBefore10Minutes"`
				JunglerKillsEarlyJungle                   int     `json:"junglerKillsEarlyJungle"`
				JunglerTakedownsNearDamagedEpicMonster    int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall         int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                       int     `json:"kda"`
				KillAfterHiddenWithAlly                   int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                         float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived     int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillsNearEnemyTurret                      int     `json:"killsNearEnemyTurret"`
				KillsOnLanersEarlyJungleAsJungler         int     `json:"killsOnLanersEarlyJungleAsJungler"`
				KillsOnRecentlyHealedByAramPack           int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                       int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster              int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                 int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                   int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                 int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage               int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                            int     `json:"legendaryCount"`
				LostAnInhibitor                           int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent              float64 `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                            int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                  int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent               float64 `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                         int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount                int     `json:"multiTurretRiftHeraldCount"`
				Multikills                                int     `json:"multikills"`
				MultikillsAfterAggressiveFlash            int     `json:"multikillsAfterAggressiveFlash"`
				MythicItemUsed                            int     `json:"mythicItemUsed"`
				OuterTurretExecutesBefore10Minutes        int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                          int     `json:"outnumberedKills"`
				OutnumberedNexusKill                      int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                   int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                               int     `json:"perfectGame"`
				PickKillWithAlly                          int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                 int     `json:"playedChampSelectPosition"`
				PoroExplosions                            int     `json:"poroExplosions"`
				QuickCleanse                              int     `json:"quickCleanse"`
				QuickFirstTurret                          int     `json:"quickFirstTurret"`
				QuickSoloKills                            int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                       int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                         int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                          int     `json:"scuttleCrabKills"`
				SkillshotsDodged                          int     `json:"skillshotsDodged"`
				SkillshotsHit                             int     `json:"skillshotsHit"`
				SnowballsHit                              int     `json:"snowballsHit"`
				SoloBaronKills                            int     `json:"soloBaronKills"`
				SoloKills                                 int     `json:"soloKills"`
				StealthWardsPlaced                        int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount                int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight           int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                     int     `json:"takedownOnFirstTurret"`
				Takedowns                                 int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage       int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn          int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                    int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                         int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                  int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                            int     `json:"teamBaronKills"`
				TeamDamagePercentage                      float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                      int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                       int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                 int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                   int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                         int     `json:"turretPlatesTaken"`
				TurretTakedowns                           int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald                int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount              int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                             int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent          float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                      float64 `json:"visionScorePerMinute"`
				WardTakedowns                             int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                    int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                              int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			Challenges6 struct {
				One2AssistStreakCount                     int     `json:"12AssistStreakCount"`
				AbilityUses                               int     `json:"abilityUses"`
				AcesBefore15Minutes                       int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                  int     `json:"alliedJungleMonsterKills"`
				BaronTakedowns                            int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount            int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                                int     `json:"bountyGold"`
				BuffsStolen                               int     `json:"buffsStolen"`
				CompleteSupportQuestInTime                int     `json:"completeSupportQuestInTime"`
				ControlWardTimeCoverageInRiverOrEnemyHalf float64 `json:"controlWardTimeCoverageInRiverOrEnemyHalf"`
				ControlWardsPlaced                        int     `json:"controlWardsPlaced"`
				DamagePerMinute                           float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage               float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                      int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                       int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow                int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                                int     `json:"doubleAces"`
				DragonTakedowns                           int     `json:"dragonTakedowns"`
				EarlyLaningPhaseGoldExpAdvantage          int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                 int     `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul          int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                     int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations              int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                   int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler          int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn    int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                         int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite             int     `json:"epicMonsterStolenWithoutSmite"`
				FirstTurretKilled                         float64 `json:"firstTurretKilled"`
				FlawlessAces                              int     `json:"flawlessAces"`
				FullTeamTakedown                          int     `json:"fullTeamTakedown"`
				GameLength                                float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner  int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                             float64 `json:"goldPerMinute"`
				HadOpenNexus                              int     `json:"hadOpenNexus"`
				HighestCrowdControlScore                  int     `json:"highestCrowdControlScore"`
				ImmobilizeAndKillWithAlly                 int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                          int     `json:"initialBuffCount"`
				InitialCrabCount                          int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                   int     `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster    int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall         int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                       float64 `json:"kda"`
				KillAfterHiddenWithAlly                   int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                         float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived     int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillingSprees                             int     `json:"killingSprees"`
				KillsNearEnemyTurret                      int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner       int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack           int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                       int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster              int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                 int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                   int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                 int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage               int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                            int     `json:"legendaryCount"`
				LostAnInhibitor                           int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent              int     `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                            int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                  int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent               int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                         int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount                int     `json:"multiTurretRiftHeraldCount"`
				Multikills                                int     `json:"multikills"`
				MultikillsAfterAggressiveFlash            int     `json:"multikillsAfterAggressiveFlash"`
				MythicItemUsed                            int     `json:"mythicItemUsed"`
				OuterTurretExecutesBefore10Minutes        int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                          int     `json:"outnumberedKills"`
				OutnumberedNexusKill                      int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                   int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                               int     `json:"perfectGame"`
				PickKillWithAlly                          int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                 int     `json:"playedChampSelectPosition"`
				PoroExplosions                            int     `json:"poroExplosions"`
				QuickCleanse                              int     `json:"quickCleanse"`
				QuickFirstTurret                          int     `json:"quickFirstTurret"`
				QuickSoloKills                            int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                       int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                         int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                          int     `json:"scuttleCrabKills"`
				SkillshotsDodged                          int     `json:"skillshotsDodged"`
				SkillshotsHit                             int     `json:"skillshotsHit"`
				SnowballsHit                              int     `json:"snowballsHit"`
				SoloBaronKills                            int     `json:"soloBaronKills"`
				SoloKills                                 int     `json:"soloKills"`
				SoloTurretsLategame                       int     `json:"soloTurretsLategame"`
				StealthWardsPlaced                        int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount                int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight           int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                     int     `json:"takedownOnFirstTurret"`
				Takedowns                                 int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage       int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn          int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                    int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                         int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                  int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                            int     `json:"teamBaronKills"`
				TeamDamagePercentage                      float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                      int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                       int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                 int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                   int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                         int     `json:"turretPlatesTaken"`
				TurretTakedowns                           int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald                int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount              int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                             int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent          float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                      float64 `json:"visionScorePerMinute"`
				WardTakedowns                             int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                    int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                              int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			Challenges7 struct {
				One2AssistStreakCount                     int     `json:"12AssistStreakCount"`
				AbilityUses                               int     `json:"abilityUses"`
				AcesBefore15Minutes                       int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                  int     `json:"alliedJungleMonsterKills"`
				BaronTakedowns                            int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount            int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                                int     `json:"bountyGold"`
				BuffsStolen                               int     `json:"buffsStolen"`
				CompleteSupportQuestInTime                int     `json:"completeSupportQuestInTime"`
				ControlWardTimeCoverageInRiverOrEnemyHalf float64 `json:"controlWardTimeCoverageInRiverOrEnemyHalf"`
				ControlWardsPlaced                        int     `json:"controlWardsPlaced"`
				DamagePerMinute                           float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage               float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                      int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                       int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow                int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                                int     `json:"doubleAces"`
				DragonTakedowns                           int     `json:"dragonTakedowns"`
				EarliestDragonTakedown                    float64 `json:"earliestDragonTakedown"`
				EarlyLaningPhaseGoldExpAdvantage          int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                 int     `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul          int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                     int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations              int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                   int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler          int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn    int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                         int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite             int     `json:"epicMonsterStolenWithoutSmite"`
				FirstTurretKilled                         float64 `json:"firstTurretKilled"`
				FlawlessAces                              int     `json:"flawlessAces"`
				FullTeamTakedown                          int     `json:"fullTeamTakedown"`
				GameLength                                float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner  int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                             float64 `json:"goldPerMinute"`
				HadOpenNexus                              int     `json:"hadOpenNexus"`
				ImmobilizeAndKillWithAlly                 int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                          int     `json:"initialBuffCount"`
				InitialCrabCount                          int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                   int     `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster    int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall         int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                       int     `json:"kda"`
				KillAfterHiddenWithAlly                   int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                         float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived     int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillingSprees                             int     `json:"killingSprees"`
				KillsNearEnemyTurret                      int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner       int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack           int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                       int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster              int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                 int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                   int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                 int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage               int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                            int     `json:"legendaryCount"`
				LostAnInhibitor                           int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent              int     `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                            int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                  int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent               int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                         int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount                int     `json:"multiTurretRiftHeraldCount"`
				Multikills                                int     `json:"multikills"`
				MultikillsAfterAggressiveFlash            int     `json:"multikillsAfterAggressiveFlash"`
				MythicItemUsed                            int     `json:"mythicItemUsed"`
				OuterTurretExecutesBefore10Minutes        int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                          int     `json:"outnumberedKills"`
				OutnumberedNexusKill                      int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                   int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                               int     `json:"perfectGame"`
				PickKillWithAlly                          int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                 int     `json:"playedChampSelectPosition"`
				PoroExplosions                            int     `json:"poroExplosions"`
				QuickCleanse                              int     `json:"quickCleanse"`
				QuickFirstTurret                          int     `json:"quickFirstTurret"`
				QuickSoloKills                            int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                       int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                         int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                          int     `json:"scuttleCrabKills"`
				SkillshotsDodged                          int     `json:"skillshotsDodged"`
				SkillshotsHit                             int     `json:"skillshotsHit"`
				SnowballsHit                              int     `json:"snowballsHit"`
				SoloBaronKills                            int     `json:"soloBaronKills"`
				SoloKills                                 int     `json:"soloKills"`
				StealthWardsPlaced                        int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount                int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight           int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                     int     `json:"takedownOnFirstTurret"`
				Takedowns                                 int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage       int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn          int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                    int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                         int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                  int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                            int     `json:"teamBaronKills"`
				TeamDamagePercentage                      float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                      int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                       int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                 int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                   int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                         int     `json:"turretPlatesTaken"`
				TurretTakedowns                           int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald                int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount              int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                             int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent          float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                      float64 `json:"visionScorePerMinute"`
				WardTakedowns                             int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                    int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                              int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			Challenges8 struct {
				One2AssistStreakCount                     int     `json:"12AssistStreakCount"`
				AbilityUses                               int     `json:"abilityUses"`
				AcesBefore15Minutes                       int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                  int     `json:"alliedJungleMonsterKills"`
				BaronTakedowns                            int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount            int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                                int     `json:"bountyGold"`
				BuffsStolen                               int     `json:"buffsStolen"`
				CompleteSupportQuestInTime                int     `json:"completeSupportQuestInTime"`
				ControlWardTimeCoverageInRiverOrEnemyHalf float64 `json:"controlWardTimeCoverageInRiverOrEnemyHalf"`
				ControlWardsPlaced                        int     `json:"controlWardsPlaced"`
				DamagePerMinute                           float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage               float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                      int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                       int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow                int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                                int     `json:"doubleAces"`
				DragonTakedowns                           int     `json:"dragonTakedowns"`
				EarliestDragonTakedown                    float64 `json:"earliestDragonTakedown"`
				EarlyLaningPhaseGoldExpAdvantage          int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                 float64 `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul          int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                     int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations              int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                   int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler          int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn    int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                         int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite             int     `json:"epicMonsterStolenWithoutSmite"`
				FirstTurretKilled                         float64 `json:"firstTurretKilled"`
				FlawlessAces                              int     `json:"flawlessAces"`
				FullTeamTakedown                          int     `json:"fullTeamTakedown"`
				GameLength                                float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner  int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                             float64 `json:"goldPerMinute"`
				HadOpenNexus                              int     `json:"hadOpenNexus"`
				HighestWardKills                          int     `json:"highestWardKills"`
				ImmobilizeAndKillWithAlly                 int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                          int     `json:"initialBuffCount"`
				InitialCrabCount                          int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                   int     `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster    int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall         int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                       int     `json:"kda"`
				KillAfterHiddenWithAlly                   int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                         float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived     int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillsNearEnemyTurret                      int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner       int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack           int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                       int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster              int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                 int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                   int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                 int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage               int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                            int     `json:"legendaryCount"`
				LostAnInhibitor                           int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent              int     `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                            int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                  int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent               int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                         int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount                int     `json:"multiTurretRiftHeraldCount"`
				Multikills                                int     `json:"multikills"`
				MultikillsAfterAggressiveFlash            int     `json:"multikillsAfterAggressiveFlash"`
				OuterTurretExecutesBefore10Minutes        int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                          int     `json:"outnumberedKills"`
				OutnumberedNexusKill                      int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                   int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                               int     `json:"perfectGame"`
				PickKillWithAlly                          int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                 int     `json:"playedChampSelectPosition"`
				PoroExplosions                            int     `json:"poroExplosions"`
				QuickCleanse                              int     `json:"quickCleanse"`
				QuickFirstTurret                          int     `json:"quickFirstTurret"`
				QuickSoloKills                            int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                       int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                         int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                          int     `json:"scuttleCrabKills"`
				SkillshotsDodged                          int     `json:"skillshotsDodged"`
				SkillshotsHit                             int     `json:"skillshotsHit"`
				SnowballsHit                              int     `json:"snowballsHit"`
				SoloBaronKills                            int     `json:"soloBaronKills"`
				SoloKills                                 int     `json:"soloKills"`
				StealthWardsPlaced                        int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount                int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight           int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                     int     `json:"takedownOnFirstTurret"`
				Takedowns                                 int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage       int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn          int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                    int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                         int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                  int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                            int     `json:"teamBaronKills"`
				TeamDamagePercentage                      float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                      int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                       int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                 int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                   int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                         int     `json:"turretPlatesTaken"`
				TurretTakedowns                           int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald                int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount              int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                             int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent          float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                      float64 `json:"visionScorePerMinute"`
				WardTakedowns                             int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                    int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                              int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
		} `json:"participants"`
		PlatformID string `json:"platformId"`
		QueueID    int    `json:"queueId"`
		Teams      []struct {
			Bans []struct {
				ChampionID int `json:"championId"`
				PickTurn   int `json:"pickTurn"`
			} `json:"bans"`
			Objectives struct {
				Baron struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"baron"`
				Champion struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"champion"`
				Dragon struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"dragon"`
				Inhibitor struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"inhibitor"`
				RiftHerald struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"riftHerald"`
				Tower struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"tower"`
			} `json:"objectives"`
			TeamID int  `json:"teamId"`
			Win    bool `json:"win"`
		} `json:"teams"`
		TournamentCode string `json:"tournamentCode"`
	} `json:"info"`
}
