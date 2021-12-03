package main

import (
	"github.com/ybalcin/ecommerce-study/internal/application/commands"
	"github.com/ybalcin/ecommerce-study/internal/application/queries"
	"github.com/ybalcin/ecommerce-study/internal/ports"
)

func main() {

	cli := ports.NewCLI()

	//createCampaign := &commands.CreateCampaignCommand{
	//	Name:                   "deneme",
	//	ProductCode:            "P1",
	//	Duration:               5,
	//	PriceManipulationLimit: 20,
	//	TargetSalesCount:       100,
	//}
	////
	//cli.CreateCampaign(createCampaign)
	//
	//createCampaign = &commands.CreateCampaignCommand{
	//	Name:                   "deneme",
	//	ProductCode:            "P1",
	//	Duration:               5,
	//	PriceManipulationLimit: 25,
	//	TargetSalesCount:       1,
	//}
	//
	//cli.CreateCampaign(createCampaign)
	//
	//createProduct := &commands.CreateProductCommand{
	//	ProductCode: "P1",
	//	Price:       100,
	//	Stock:       10,
	//}
	//
	//cli.CreateProduct(createProduct)

	cli.IncreaseTime(&commands.IncreaseTimeCommand{Hours: 1})

	cli.GetProductInfo(&queries.GetProductInfoQuery{Code: "P1"})

	cli.GetCampaignInfo(&queries.GetCampaignInfoQuery{Name: "deneme"})

	cli.CreateOrder(&commands.CreateOrderCommand{
		ProductCode: "P1",
		Quantity:    10,
	})

	cli.GetCampaignInfo(&queries.GetCampaignInfoQuery{Name: "deneme"})

	//cli.GetProductInfo(&queries.GetProductInfoQuery{Code: "P1"})
	//
	//cli.IncreaseTime(&commands.IncreaseTimeCommand{Hours: 1})
	//
	//cli.GetProductInfo(&queries.GetProductInfoQuery{Code: "P1"})
	//
	//cli.IncreaseTime(&commands.IncreaseTimeCommand{Hours: 2})
	//
	//cli.GetProductInfo(&queries.GetProductInfoQuery{Code: "P1"})
}
