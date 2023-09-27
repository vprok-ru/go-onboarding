package main

import "main/internal/chain-of-responsibility"

func main() {
	stage := chain_of_responsibility.NewCarClient(nil)
	stage = chain_of_responsibility.NewCarDealership(stage)
	stage = chain_of_responsibility.NewCarFactory(stage)

	stageInfo, err := stage.CheckStatus(chain_of_responsibility.UnknownStatus)
	if err != nil {
		println(err.Error())
	} else {
		println(stageInfo)
	}

	stageInfo, err = stage.CheckStatus(chain_of_responsibility.StatusFactory)
	if err != nil {
		println(err.Error())
	} else {
		println(stageInfo)
	}

	stageInfo, err = stage.CheckStatus(chain_of_responsibility.StatusClient)
	if err != nil {
		println(err.Error())
	} else {
		println(stageInfo)
	}
}
