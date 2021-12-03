package ports

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/app"
	"github.com/ybalcin/ecommerce-study/internal/application/commands"
	"log"
	"os"
)

type CLI struct {
	application *app.Application
	ctx         context.Context
}

var logger = log.New(os.Stdout, "", log.LstdFlags)

// NewCLI cli port
func NewCLI() *CLI {
	ctx := context.Background()

	application := app.New(ctx)

	return &CLI{application: application}
}

func (cli *CLI) CreateCampaign(command *commands.CreateCampaignCommand) {
	resp, err := cli.application.Commands.CreateCampaign.Handle(cli.ctx, command)
	if err != nil {
		logger.Println(err.Error())
		return
	}

	logger.Println(resp.String())
}

func (cli *CLI) CreateOrder(command *commands.CreateOrderCommand) {
	resp, err := cli.application.Commands.CreateOrder.Handle(cli.ctx, command)
	if err != nil {
		logger.Println(err.Error())
		return
	}

	logger.Println(resp.String())
}

func (cli *CLI) CreateProduct(command *commands.CreateProductCommand) {
	resp, err := cli.application.Commands.CreateProduct.Handle(cli.ctx, command)
	if err != nil {
		logger.Println(err.Error())
		return
	}

	logger.Println(resp.String())
}

func (cli *CLI) IncreaseTime(command *commands.IncreaseTimeCommand) {
	resp, err := cli.application.Commands.IncreaseTime.Handle(command)
	if err != nil {
		logger.Println(err.Error())
		return
	}

	logger.Println(resp.String())
}
