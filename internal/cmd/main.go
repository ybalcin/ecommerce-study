package main

import (
	"context"
	"fmt"
	"github.com/ybalcin/ecommerce-study/internal/app"
	"github.com/ybalcin/ecommerce-study/internal/application/commands"
)

func main() {

	ctx := context.Background()

	application := app.New(ctx)
	resp, err := application.Commands.CreateCampaign.Handle(ctx, &commands.CreateCampaignCommand{
		Name:                   "test1",
		ProductCode:            "asdasd",
		Duration:               1,
		PriceManipulationLimit: 1,
		TargetSalesCount:       1,
	})
	fmt.Println(resp.String())

	application.SysTime.Add(1)

	resp, err = application.Commands.CreateCampaign.Handle(ctx, &commands.CreateCampaignCommand{
		Name:                   "test1",
		ProductCode:            "asdasd",
		Duration:               1,
		PriceManipulationLimit: 1,
		TargetSalesCount:       1,
	})

	fmt.Println(resp)

	if err != nil {
		panic(err)
	}

}
