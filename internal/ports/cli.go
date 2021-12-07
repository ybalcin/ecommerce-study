package ports

import (
	"context"
	"fmt"
	"github.com/ybalcin/ecommerce-study/internal/app"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/createcampaign"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/createorder"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/createproduct"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/increasetime"
	"github.com/ybalcin/ecommerce-study/internal/application/queries/getcampaigninfo"
	"github.com/ybalcin/ecommerce-study/internal/application/queries/getproductinfo"
	"github.com/ybalcin/ecommerce-study/internal/common"
	"strings"
)

type cli struct {
	application *app.Application
	ctx         context.Context
}

const (
	createProductEntry  string = "create_product"
	createCampaignEntry string = "create_campaign"
	createOrderEntry    string = "create_order"
	increaseTimeEntry   string = "increase_time"

	getCampaignInfoEntry string = "get_campaign_info"
	getProductInfoEntry  string = "get_product_info"
)

// NewCLI cli port
func NewCLI() *cli {
	ctx := context.Background()

	applicationEntry := app.New(ctx)

	return &cli{application: applicationEntry}
}

// Execute executes command
func (cli *cli) Execute(cmd string) {
	if cmd == "" {
		fmt.Println("cmd cannot be empty")
	}

	cmdTokens := strings.Split(cmd, " ")

	entry := common.ValueOfSlice(0, cmdTokens)
	if entry == "" {
		return
	}

	switch entry {
	case createCampaignEntry:
		cli.createCampaign(cmd)
	case createOrderEntry:
		cli.createOrder(cmd)
	case createProductEntry:
		cli.createProduct(cmd)
	case increaseTimeEntry:
		cli.increaseTime(cmd)
	case getCampaignInfoEntry:
		cli.getCampaignInfo(cmd)
	case getProductInfoEntry:
		cli.getProductInfo(cmd)
	}
}

func (cli *cli) createCampaign(cmd string) {
	command := createcampaign.Build(cmd)
	resp, err := cli.application.Commands.CreateCampaign.Handle(cli.ctx, command)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.String())
}

func (cli *cli) createOrder(cmd string) {
	command := createorder.Build(cmd)
	resp, err := cli.application.Commands.CreateOrder.Handle(cli.ctx, command)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.String())
}

func (cli *cli) createProduct(cmd string) {
	command := createproduct.Build(cmd)
	resp, err := cli.application.Commands.CreateProduct.Handle(cli.ctx, command)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.String())
}

func (cli *cli) increaseTime(cmd string) {
	command := increasetime.Build(cmd)
	resp, err := cli.application.Commands.IncreaseTime.Handle(command)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.String())
}

func (cli *cli) getCampaignInfo(cmd string) {
	query := getcampaigninfo.Build(cmd)
	resp, err := cli.application.Queries.GetCampaignInfo.Handle(cli.ctx, query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.String())
}

func (cli *cli) getProductInfo(cmd string) {
	query := getproductinfo.Build(cmd)
	resp, err := cli.application.Queries.GetProductInfo.Handle(cli.ctx, query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.String())
}
