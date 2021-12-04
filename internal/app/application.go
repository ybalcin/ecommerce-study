package app

import (
	"context"
	"errors"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/createcampaign"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/createorder"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/createproduct"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/increasetime"
	"github.com/ybalcin/ecommerce-study/internal/application/queries/getcampaigninfo"
	"github.com/ybalcin/ecommerce-study/internal/application/queries/getproductinfo"
	"github.com/ybalcin/ecommerce-study/internal/infrastructure/adapters"
	"github.com/ybalcin/ecommerce-study/pkg/mgo"
)

type Application struct {
	Queries  Queries
	Commands Commands
	SysTime  *application.SystemTime
}

type Queries struct {
	GetCampaignInfo *getcampaigninfo.Handler
	GetProductInfo  *getproductinfo.Handler
}

type Commands struct {
	CreateCampaign *createcampaign.Handler
	CreateOrder    *createorder.Handler
	CreateProduct  *createproduct.Handler
	IncreaseTime   *increasetime.Handler
}

// New initializes new application
func New(ctx context.Context) *Application {
	sysTime := application.NewSystemTime()

	// TODO: json settings
	mgoStore := mgo.NewStore(ctx, "mongodb+srv://ecommerce-user:B9VeLojwHUidkeHP@cluster0.l1pmb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority", "ecommerce")
	if mgoStore == nil {
		panic(errors.New("mgo store nil"))
	}

	campaignMgoRepository, err := adapters.NewCampaignRepository(mgoStore)
	checkPanic(err)

	productMgoRepository, err := adapters.NewProductMongoRepository(mgoStore)
	checkPanic(err)

	orderMgoRepository, err := adapters.NewOrderRepository(mgoStore)
	checkPanic(err)

	// TODO: Remove collections for test purpose only
	err = campaignMgoRepository.DropCampaigns(ctx)
	checkPanic(err)

	err = orderMgoRepository.DropOrders(ctx)
	checkPanic(err)

	err = productMgoRepository.DropProducts(ctx)
	checkPanic(err)

	createCampaignCommandHandler := createcampaign.NewHandler(campaignMgoRepository, sysTime)

	createOrderCommandHandler := createorder.NewHandler(
		orderMgoRepository, productMgoRepository, campaignMgoRepository, sysTime)

	createProductCommandHandler := createproduct.NewHandler(productMgoRepository)

	increaseTimeCommandHandler := increasetime.NewHandler(sysTime)

	getCampaignInfoQueryHandler := getcampaigninfo.NewHandler(campaignMgoRepository, orderMgoRepository,
		productMgoRepository, sysTime)

	getProductInfoQueryHandler := getproductinfo.NewHandler(productMgoRepository, campaignMgoRepository,
		orderMgoRepository, sysTime)

	return &Application{
		Queries: Queries{
			GetCampaignInfo: getCampaignInfoQueryHandler,
			GetProductInfo:  getProductInfoQueryHandler,
		},
		Commands: Commands{
			CreateCampaign: createCampaignCommandHandler,
			CreateOrder:    createOrderCommandHandler,
			CreateProduct:  createProductCommandHandler,
			IncreaseTime:   increaseTimeCommandHandler,
		},
		SysTime: sysTime,
	}
}

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}
