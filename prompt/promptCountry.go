package prompt

import (
	"fmt"

	"github.com/dragos199993/dxpcli/config"
	"github.com/dragos199993/dxpcli/utils"
	"github.com/manifoldco/promptui"
)

func PromptCountry(FinalMarket config.Market) config.Market {
	var countryList []string
	var countryKeyValue = []config.Attribute {};
	for _, market := range config.AvailableMarkets {
		if market.Market.Key == FinalMarket.Market && market.Region.Key == FinalMarket.Region {
			countryList = append(countryList, market.Country.Val)
			countryKeyValue = append(countryKeyValue, config.Attribute{market.Country.Key, market.Country.Val});
		}
	}

	countryList = utils.UniqueNonEmptyElementsOf(countryList);
	prompt := promptui.Select{
		Label: "Select country:",
		Items: countryList,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		panic(err);
	}

	for _, v := range countryKeyValue {
    if v.Val == result {
			FinalMarket.Country = v.Key
    }
	}

	return FinalMarket;
}
