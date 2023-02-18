package main

import (
	"context"
	"strings"

	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/rs/zerolog/log"
)

func (a *app) SetupDemo() {
	csvText := `Import Ref,Location,Labels,Quantity,Name,Description,Insured,Serial Number,Model Number,Manufacturer,Notes,Purchase From,Purchased Price,Purchased Time,Lifetime Warranty,Warranty Expires,Warranty Details,Sold To,Sold Price,Sold Time,Sold Notes
,Garage,IOT;Home Assistant; Z-Wave,1,Zooz Universal Relay ZEN17,"Zooz 700 Series Z-Wave Universal Relay ZEN17 for Awnings, Garage Doors, Sprinklers, and More | 2 NO-C-NC Relays (20A, 10A) | Signal Repeater | Hub Required (Compatible with SmartThings and Hubitat)",,,ZEN17,Zooz,,Amazon,39.95,10/13/2021,,,,,,,
,Living Room,IOT;Home Assistant; Z-Wave,1,Zooz Motion Sensor,"Zooz Z-Wave Plus S2 Motion Sensor ZSE18 with Magnetic Mount, Works with Vera and SmartThings",,,ZSE18,Zooz,,Amazon,29.95,10/15/2021,,,,,,,
,Office,IOT;Home Assistant; Z-Wave,1,Zooz 110v Power Switch,"Zooz Z-Wave Plus Power Switch ZEN15 for 110V AC Units, Sump Pumps, Humidifiers, and More",,,ZEN15,Zooz,,Amazon,39.95,10/13/2021,,,,,,,
,Downstairs,IOT;Home Assistant; Z-Wave,1,Ecolink Z-Wave PIR Motion Sensor,"Ecolink Z-Wave PIR Motion Detector Pet Immune, White (PIRZWAVE2.5-ECO)",,,PIRZWAVE2.5-ECO,Ecolink,,Amazon,35.58,10/21/2020,,,,,,,
,Entry,IOT;Home Assistant; Z-Wave,1,Yale Security Touchscreen Deadbolt,"Yale Security YRD226-ZW2-619 YRD226ZW2619 Touchscreen Deadbolt, Satin Nickel",,,YRD226ZW2619,Yale,,Amazon,120.39,10/14/2020,,,,,,,
,Kitchen,IOT;Home Assistant; Z-Wave,1,Smart Rocker Light Dimmer,"UltraPro Z-Wave Smart Rocker Light Dimmer with QuickFit and SimpleWire, 3-Way Ready, Compatible with Alexa, Google Assistant, ZWave Hub Required, Repeater/Range Extender, White Paddle Only, 39351",,,‎39351,Honeywell,,Amazon,65.98,09/30/0202,,,,,,,
`

	registration := services.UserRegistration{
		Email:    "demo@example.com",
		Name:     "Demo",
		Password: "demo",
	}

	// First check if we've already setup a demo user and skip if so
	_, err := a.services.User.Login(context.Background(), registration.Email, registration.Password)
	if err == nil {
		return
	}

	_, err = a.services.User.RegisterUser(context.Background(), registration)
	if err != nil {
		log.Err(err).Msg("Failed to register demo user")
		log.Fatal().Msg("Failed to setup demo")
	}

	token, _ := a.services.User.Login(context.Background(), registration.Email, registration.Password)
	self, _ := a.services.User.GetSelf(context.Background(), token.Raw)

	_, err = a.services.Items.CsvImport(context.Background(), self.GroupID, strings.NewReader(csvText))
	if err != nil {
		log.Err(err).Msg("Failed to import CSV")
		log.Fatal().Msg("Failed to setup demo")
	}

	log.Info().Msg("Demo setup complete")
}
