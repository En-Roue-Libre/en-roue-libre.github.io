package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	events := GetEvents()

	for i, event := range events.PostedEvents {
		lateAmount := 0
		var compURLBuilder, specParamsBuilder, nameParamsBuilder strings.Builder
		compURLBuilder.WriteString("https://wowtbc.gg/wotlk/raid-comp/?c=")

		event.Details = GetEventDetails(event.ID)

		event.Tanks = []SignUp{}
		event.DPS = []SignUp{}
		event.Healers = []SignUp{}

		for _, signUp := range event.Details.SignUps {

			signUp.IconURL = iconURLMap[signUp.SpecName]
			signUp.ClassColor = specColorMap[signUp.SpecName]
			signUp.SpecID = specIDMap[signUp.SpecName]

			if signUp.ClassName == "Late" {
				lateAmount++
			}

			if signUp.ClassName != "Tentative" && signUp.ClassName != "Absence" && signUp.ClassName != "Bench" {
				specParamsBuilder.WriteString(signUp.SpecID)
				nameParamsBuilder.WriteString(signUp.Name + "!")
				if signUp.RoleName == "Tanks" {
					event.Tanks = append(event.Tanks, signUp)
				}
				if signUp.RoleName == "Healers" {
					event.Healers = append(event.Healers, signUp)
				}
				if signUp.RoleName == "Melee" || signUp.RoleName == "Ranged" {
					event.DPS = append(event.DPS, signUp)
				}
			}
		}

		event.SignUpsAmount = event.SignUpsAmount + lateAmount

		specParamsBuilder.WriteString(strings.Repeat("00", 30-event.SignUpsAmount))
		nameParamsBuilder.WriteString(strings.Repeat("!", 30-event.SignUpsAmount-1))

		compURLBuilder.WriteString(specParamsBuilder.String())
		compURLBuilder.WriteString("&n=")
		compURLBuilder.WriteString(nameParamsBuilder.String())

		event.CompoURL = compURLBuilder.String()

		events.PostedEvents[i] = event
	}

	file, err := os.OpenFile("src/_data/calendar.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(events)
}

func GetEvents() AllEvents {
	var urlBuilder strings.Builder
	var events AllEvents

	serverId := os.Getenv("RH_SERVERID")
	apiKey := os.Getenv("RH_APIKEY")
	method := "GET"

	urlBuilder.WriteString("https://raid-helper.dev/api/v2/servers/")
	urlBuilder.WriteString(serverId)
	urlBuilder.WriteString("/events")

	client := &http.Client{}
	req, err := http.NewRequest(method, urlBuilder.String(), nil)

	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", apiKey)

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &events); err != nil {
		panic(err)
	}
	return events
}

func GetEventDetails(eventID string) EventDetails {

	var urlBuilder strings.Builder
	var eventDetails EventDetails

	apiKey := os.Getenv("RH_APIKEY")
	method := "GET"

	urlBuilder.WriteString("https://raid-helper.dev/api/v2/events/")
	urlBuilder.WriteString(eventID)

	client := &http.Client{}
	req, err := http.NewRequest(method, urlBuilder.String(), nil)

	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", apiKey)

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &eventDetails); err != nil {
		panic(err)
	}
	return eventDetails
}

type AllEvents struct {
	ScheduledEvents []Event `json:"scheduledEvents"`
	PostedEvents    []Event `json:"postedEvents"`
}

type Event struct {
	Tanks         []SignUp     `json:"tanks"`
	Healers       []SignUp     `json:"healers"`
	DPS           []SignUp     `json:"dps"`
	CompoURL      string       `json:"compoUrl"`
	Details       EventDetails `json:"details"`
	SignUpsAmount int          `json:"signUpsAmount"`
	Description   string       `json:"description"`
	Title         string       `json:"title"`
	LeaderName    string       `json:"leaderName"`
	StartTime     int          `json:"startTime"`
	ID            string       `json:"id"`
}

type EventDetails struct {
	Date             string           `json:"date"`
	SignUps          []SignUp         `json:"signUps"`
	Description      string           `json:"description"`
	ChannelType      string           `json:"channelType"`
	Title            string           `json:"title"`
	TemplateID       string           `json:"templateId"`
	LastUpdated      int              `json:"lastUpdated"`
	LeaderName       string           `json:"leaderName"`
	AdvancedSettings AdvancedSettings `json:"advancedSettings"`
	StartTime        int              `json:"startTime"`
	ID               string           `json:"id"`
}

type SignUp struct {
	Name       string `json:"name"`
	ClassName  string `json:"className"`
	Position   int    `json:"position"`
	SpecName   string `json:"specName,omitempty"`
	RoleName   string `json:"roleName,omitempty"`
	IconURL    string `json:"iconUrl,omitempty"`
	ClassColor string `json:"classColor,omitempty"`
	SpecID     string `json:"specId,omitempty"`
}

type AdvancedSettings struct {
	Limit     int    `json:"limit"`
	Image     string `json:"image"`
	Thumbnail string `json:"thumbnail"`
}

var iconURLMap = map[string]string{
	"Blood_Tank":    "https://wow.zamimg.com/images/wow/icons/medium/spell_deathknight_bloodpresence.jpg",
	"Frost_Tank":    "https://wow.zamimg.com/images/wow/icons/medium/spell_deathknight_frostpresence.jpg",
	"Unholy_Tank":   "https://wow.zamimg.com/images/wow/icons/medium/spell_deathknight_unholypresence.jpg",
	"Blood_DPS":     "https://wow.zamimg.com/images/wow/icons/medium/spell_deathknight_bloodpresence.jpg",
	"Frost_DPS":     "https://wow.zamimg.com/images/wow/icons/medium/spell_deathknight_frostpresence.jpg",
	"Unholy_DPS":    "https://wow.zamimg.com/images/wow/icons/medium/spell_deathknight_unholypresence.jpg",
	"Arms":          "https://wow.zamimg.com/images/wow/icons/medium/ability_warrior_savageblow.jpg",
	"Fury":          "https://wow.zamimg.com/images/wow/icons/medium/ability_warrior_innerrage.jpg",
	"Protection":    "https://wow.zamimg.com/images/wow/icons/medium/inv_shield_06.jpg",
	"Guardian":      "https://wow.zamimg.com/images/wow/icons/medium/ability_racial_bearform.jpg",
	"Balance":       "https://wow.zamimg.com/images/wow/icons/medium/spell_nature_starfall.jpg",
	"Feral":         "https://wow.zamimg.com/images/wow/icons/medium/ability_druid_catform.jpg",
	"Restoration":   "https://wow.zamimg.com/images/wow/icons/medium/spell_nature_healingtouch.jpg",
	"Protection1":   "https://wow.zamimg.com/images/wow/icons/medium/spell_holy_devotionaura.jpg",
	"Holy1":         "https://wow.zamimg.com/images/wow/icons/medium/spell_holy_holybolt.jpg",
	"Retribution":   "https://wow.zamimg.com/images/wow/icons/medium/spell_holy_auraoflight.jpg",
	"Assassination": "https://wow.zamimg.com/images/wow/icons/medium/ability_rogue_eviscerate.jpg",
	"Combat":        "https://wow.zamimg.com/images/wow/icons/medium/ability_backstab.jpg",
	"Subtlety":      "https://wow.zamimg.com/images/wow/icons/medium/ability_stealth.jpg",
	"Beastmastery":  "https://wow.zamimg.com/images/wow/icons/medium/ability_hunter_beasttaming.jpg",
	"Marksmanship":  "https://wow.zamimg.com/images/wow/icons/medium/ability_marksmanship.jpg",
	"Survival":      "https://wow.zamimg.com/images/wow/icons/medium/ability_hunter_swiftstrike.jpg",
	"Arcane":        "https://wow.zamimg.com/images/wow/icons/medium/spell_holy_magicalsentry.jpg",
	"Fire":          "https://wow.zamimg.com/images/wow/icons/medium/spell_fire_firebolt02.jpg",
	"Frost":         "https://wow.zamimg.com/images/wow/icons/medium/spell_frost_frostbolt02.jpg",
	"Affliction":    "https://wow.zamimg.com/images/wow/icons/medium/spell_shadow_deathcoil.jpg",
	"Demonology":    "https://wow.zamimg.com/images/wow/icons/medium/spell_shadow_metamorphosis.jpg",
	"Destruction":   "https://wow.zamimg.com/images/wow/icons/medium/spell_shadow_rainoffire.jpg",
	"Discipline":    "https://wow.zamimg.com/images/wow/icons/medium/spell_holy_wordfortitude.jpg",
	"Holy":          "https://wow.zamimg.com/images/wow/icons/medium/spell_holy_guardianspirit.jpg",
	"Shadow":        "https://wow.zamimg.com/images/wow/icons/medium/spell_shadow_shadowwordpain.jpg",
	"Smite":         "https://wow.zamimg.com/images/wow/icons/medium/spell_holy_holysmite.jpg",
	"Elemental":     "https://wow.zamimg.com/images/wow/icons/medium/spell_nature_lightning.jpg",
	"Enhancement":   "https://wow.zamimg.com/images/wow/icons/medium/spell_nature_lightningshield.jpg",
	"Restoration1":  "https://wow.zamimg.com/images/wow/icons/medium/spell_nature_magicimmunity.jpg",
}

var specColorMap = map[string]string{
	"Blood_Tank":    "dk-color",
	"Frost_Tank":    "dk-color",
	"Unholy_Tank":   "dk-color",
	"Blood_DPS":     "dk-color",
	"Frost_DPS":     "dk-color",
	"Unholy_DPS":    "dk-color",
	"Arms":          "warrior-color",
	"Fury":          "warrior-color",
	"Protection":    "warrior-color",
	"Guardian":      "druid-color",
	"Balance":       "druid-color",
	"Feral":         "druid-color",
	"Restoration":   "druid-color",
	"Protection1":   "paladin-color",
	"Holy1":         "paladin-color",
	"Retribution":   "paladin-color",
	"Assassination": "rogue-color",
	"Combat":        "rogue-color",
	"Subtlety":      "rogue-color",
	"Beastmastery":  "hunter-color",
	"Marksmanship":  "hunter-color",
	"Survival":      "hunter-color",
	"Arcane":        "mage-color",
	"Fire":          "mage-color",
	"Frost":         "mage-color",
	"Affliction":    "warlock-color",
	"Demonology":    "warlock-color",
	"Destruction":   "warlock-color",
	"Discipline":    "priest-color",
	"Holy":          "priest-color",
	"Shadow":        "priest-color",
	"Smite":         "priest-color",
	"Elemental":     "shaman-color",
	"Enhancement":   "shaman-color",
	"Restoration1":  "shaman-color",
}

var specIDMap = map[string]string{
	"Blood_Tank":    "28",
	"Frost_Tank":    "29",
	"Unholy_Tank":   "30",
	"Blood_DPS":     "28",
	"Frost_DPS":     "29",
	"Unholy_DPS":    "30",
	"Arms":          "25",
	"Fury":          "26",
	"Protection":    "27",
	"Guardian":      "02",
	"Balance":       "01",
	"Feral":         "02",
	"Restoration":   "03",
	"Protection1":   "11",
	"Holy1":         "10",
	"Retribution":   "12",
	"Assassination": "16",
	"Combat":        "17",
	"Subtlety":      "18",
	"Beastmastery":  "04",
	"Marksmanship":  "05",
	"Survival":      "06",
	"Arcane":        "07",
	"Fire":          "08",
	"Frost":         "09",
	"Affliction":    "22",
	"Demonology":    "23",
	"Destruction":   "24",
	"Discipline":    "13",
	"Holy":          "14",
	"Shadow":        "15",
	"Smite":         "13",
	"Elemental":     "19",
	"Enhancement":   "20",
	"Restoration1":  "21",
}
