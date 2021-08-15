package prompt

import (
	"fmt"

	"github.com/dragos199993/dxpcli/config"
	"github.com/dragos199993/dxpcli/utils"
	"github.com/manifoldco/promptui"
)


func PromptLanguage(FinalMarket config.Market) config.Market {
	var languageList []string
	var languageKeyValue = []config.Attribute {};
	for _, market := range config.AvailableMarkets {
		if market.Market.Key == FinalMarket.Market && 
				market.Region.Key == FinalMarket.Region  && 
				market.Country.Key == FinalMarket.Country {
			languageList = append(languageList, market.Language.Val)
			languageKeyValue = append(languageKeyValue, config.Attribute{market.Language.Key, market.Language.Val});
		}
	}

	languageList = utils.UniqueNonEmptyElementsOf(languageList);
	prompt := promptui.Select{
		Label: "Select language:",
		Items: languageList,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		panic(err);
	}

	for _, v := range languageKeyValue {
    if v.Val == result {
			FinalMarket.Language = v.Key
    }
	}

	return FinalMarket;
}
