/*
Copyright © 2021 NAME HERE dragos199993@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/dragos199993/dxpcli/config"
	"github.com/dragos199993/dxpcli/prompt"
	"github.com/dragos199993/dxpcli/utils"
	"github.com/spf13/cobra"
)

var FinalMarket = config.Market {
	Market: "",
	Region: "",
	Language: "",
	Country: "",
}

var marketCmd = &cobra.Command{
	Use:   "market",
	Short: "Output an available DXP Market",
	Long: `Used to output an available DXP Market`,
	Run: func(cmd *cobra.Command, args []string) {
		FinalMarket = prompt.PromptMarket(FinalMarket);
		FinalMarket = prompt.PromptRegion(FinalMarket);
		FinalMarket = prompt.PromptCountry(FinalMarket);
		FinalMarket = prompt.PromptLanguage(FinalMarket);
		utils.GenerateEnv(FinalMarket);
	},
}

func init() {
	rootCmd.AddCommand(marketCmd)
}

type Attribute struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

