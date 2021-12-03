package ports

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/app"
	"github.com/ybalcin/ecommerce-study/internal/application/commands"
	"github.com/ybalcin/ecommerce-study/internal/application/queries"
	"log"
	"os"
)

type cLI struct {
	application *app.Application
	ctx         context.Context
}

var logger = log.New(os.Stdout, "", log.LstdFlags)

// NewCLI cli port
func NewCLI() *cLI {
	ctx := context.Background()

	applicationEntry := app.New(ctx)

	return &cLI{application: applicationEntry}
}

func (cli *cLI) CreateCampaign(command *commands.CreateCampaignCommand) {
	resp, err := cli.application.Commands.CreateCampaign.Handle(cli.ctx, command)
	if err != nil {
		logger.Println(err.Error())
		return
	}

	logger.Println(resp.String())
}

func (cli *cLI) CreateOrder(command *commands.CreateOrderCommand) {
	resp, err := cli.application.Commands.CreateOrder.Handle(cli.ctx, command)
	if err != nil {
		logger.Println(err.Error())
		return
	}

	logger.Println(resp.String())
}

func (cli *cLI) CreateProduct(command *commands.CreateProductCommand) {
	resp, err := cli.application.Commands.CreateProduct.Handle(cli.ctx, command)
	if err != nil {
		logger.Println(err.Error())
		return
	}

	logger.Println(resp.String())
}

func (cli *cLI) IncreaseTime(command *commands.IncreaseTimeCommand) {
	resp, err := cli.application.Commands.IncreaseTime.Handle(command)
	if err != nil {
		logger.Println(err.Error())
		return
	}

	logger.Println(resp.String())
}

func (cli *cLI) GetCampaignInfo(query *queries.GetCampaignInfoQuery) {
	resp, err := cli.application.Queries.GetCampaignInfo.Handle(cli.ctx, query)
	if err != nil {
		logger.Println(err.Error())
		return
	}

	logger.Println(resp.String())
}

func (cli *cLI) GetProductInfo(query *queries.GetProductInfoQuery) {
	resp, err := cli.application.Queries.GetProductInfo.Handle(cli.ctx, query)
	if err != nil {
		logger.Println(err.Error())
		return
	}

	logger.Println(resp.String())
}
