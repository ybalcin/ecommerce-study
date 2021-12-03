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
)

type cLI struct {
	application *app.Application
	ctx         context.Context
}

// NewCLI cli port
func NewCLI() *cLI {
	ctx := context.Background()

	applicationEntry := app.New(ctx)

	return &cLI{application: applicationEntry}
}

func (cli *cLI) CreateCampaign(command *createcampaign.Command) {
	resp, err := cli.application.Commands.CreateCampaign.Handle(cli.ctx, command)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.String())
}

func (cli *cLI) CreateOrder(command *createorder.Command) {
	resp, err := cli.application.Commands.CreateOrder.Handle(cli.ctx, command)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.String())
}

func (cli *cLI) CreateProduct(command *createproduct.Command) {
	resp, err := cli.application.Commands.CreateProduct.Handle(cli.ctx, command)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.String())
}

func (cli *cLI) IncreaseTime(command *increasetime.Command) {
	resp, err := cli.application.Commands.IncreaseTime.Handle(command)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.String())
}

func (cli *cLI) GetCampaignInfo(query *getcampaigninfo.Query) {
	resp, err := cli.application.Queries.GetCampaignInfo.Handle(cli.ctx, query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.String())
}

func (cli *cLI) GetProductInfo(query *getproductinfo.Query) {
	resp, err := cli.application.Queries.GetProductInfo.Handle(cli.ctx, query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.String())
}
