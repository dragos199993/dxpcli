package prompt

import (
	"fmt"

	"github.com/dragos199993/dxpcli/config"
	"github.com/dragos199993/dxpcli/utils"
	"github.com/manifoldco/promptui"
)

func PromptMarket(FinalMarket config.Market) config.Market {
	var marketList []string;
	var marketKeyValue = []config.Attribute {};
	for _, market := range config.AvailableMarkets {
		marketList = append(marketList, market.Market.Val)
		marketKeyValue = append(marketKeyValue, config.Attribute{market.Market.Key, market.Market.Val});
	}

	marketList = utils.UniqueNonEmptyElementsOf(marketList);
	prompt := promptui.Select{
		Label: "Select market:",
		Items: marketList,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		panic(err);
	}

	for _, v := range marketKeyValue {
    if v.Val == result {
			FinalMarket.Market = v.Key
    }
	}

	return FinalMarket;
}
