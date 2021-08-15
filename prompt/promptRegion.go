package prompt

import (
	"fmt"

	"github.com/dragos199993/dxpcli/config"
	"github.com/dragos199993/dxpcli/utils"
	"github.com/manifoldco/promptui"
)

func PromptRegion(FinalMarket config.Market) config.Market {
	var regionList []string
	var regionKeyValue = []config.Attribute {};
	for _, market := range config.AvailableMarkets {
		if market.Market.Key == FinalMarket.Market {
			regionList = append(regionList, market.Region.Val)
			regionKeyValue = append(regionKeyValue, config.Attribute{market.Region.Key, market.Region.Val});
		}
	}

	regionList = utils.UniqueNonEmptyElementsOf(regionList);
	prompt := promptui.Select{
		Label: "Select region:",
		Items: regionList,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		panic(err);
	}

	for _, v := range regionKeyValue {
    if v.Val == result {
			FinalMarket.Region = v.Key
    }
	}

	return FinalMarket;
}