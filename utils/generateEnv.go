package utils

import (
	"fmt"
	"os"

	"github.com/TwinProduction/go-color"
	"github.com/dragos199993/dxpcli/config"
)

func GenerateEnv(FinalMarket config.Market) {
	f, err := os.Create(".env.local.override")
	if err != nil {
		fmt.Println(err)
		panic(err);
	}
	l, err := f.WriteString(fmt.Sprintf(`
REACT_APP_BRAND_CODE=%s
REACT_APP_COUNTRY=%s
REACT_APP_REGION=%s
REACT_DEFAULT_LANGUAGE=%s
`, FinalMarket.Market, FinalMarket.Country, FinalMarket.Region, FinalMarket.Language));

	if err != nil {
		fmt.Println(err)
		f.Close()
		panic(err);
	}

	fmt.Println(l, color.Ize(color.Green,`
.env.local.override generated successfully!

Restart application to have the latest market applied.`))
	err = f.Close()
	if err != nil {
			fmt.Println(err)
			panic(err);
	}
}
