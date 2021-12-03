package app

import (
	"context"
	"errors"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/application/commands"
	"github.com/ybalcin/ecommerce-study/internal/application/queries"
	"github.com/ybalcin/ecommerce-study/internal/infrastructure/adapters"
	"github.com/ybalcin/ecommerce-study/pkg/mgo"
)

type Application struct {
	Queries  Queries
	Commands Commands
	SysTime  *application.SystemTime
}

type Queries struct {
	GetCampaignInfo *queries.GetCampaignInfoQueryHandler
	GetProductInfo  *queries.GetProductInfoQueryHandler
}

type Commands struct {
	CreateCampaign *commands.CreateCampaignCommandHandler
	CreateOrder    *commands.CreateOrderCommandHandler
	CreateProduct  *commands.CreateProductCommandHandler
	IncreaseTime   *commands.IncreaseTimeCommandHandler
}

// New initializes new application
func New(ctx context.Context) *Application {
	sysTime := application.NewSystemTime()

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

	createCampaignCommandHandler := commands.NewCreateCampaignCommandHandler(campaignMgoRepository, sysTime)

	createOrderCommandHandler := commands.NewCreateOrderCommandHandler(
		orderMgoRepository, productMgoRepository, campaignMgoRepository, sysTime)

	createProductCommandHandler := commands.NewCreateProductCommandHandler(productMgoRepository)

	increaseTimeCommandHandler := commands.NewIncreaseTimeCommandHandler(sysTime)

	getCampaignInfoQueryHandler := queries.NewGetCampaignInfoQueryHandler(campaignMgoRepository, orderMgoRepository,
		productMgoRepository, sysTime)

	getProductInfoQueryHandler := queries.NewGetProductInfoQueryHandler(productMgoRepository, campaignMgoRepository,
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
