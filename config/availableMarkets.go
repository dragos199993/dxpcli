package config

import (
	"github.com/dragos199993/dxpcli/constants"
)

var AvailableMarkets = []RMB{ 
	{
		Market: Attribute(constants.Subaru),
		Region: Attribute(constants.Latam),
		Language: Attribute(constants.English),
		Country: Attribute(constants.Australia),
	},
	{
		Market: Attribute(constants.BMW),
		Region: Attribute(constants.Latam),
		Language: Attribute(constants.SpanishChile),
		Country: Attribute(constants.Chile),
	},
	{
		Market: Attribute(constants.BMW),
		Region: Attribute(constants.Europe),
		Language: Attribute(constants.Finnish),
		Country: Attribute(constants.Finland),
	},
	{
		Market: Attribute(constants.Hino),
		Region: Attribute(constants.Latam),
		Language: Attribute(constants.SpanishColombia),
		Country: Attribute(constants.Colombia),
	},
}
