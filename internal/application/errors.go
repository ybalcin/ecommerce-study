package application

import (
	"errors"
	"fmt"
)

const (
	campaignRepositoryIsNilError string = "campaign repository nil"
	productRepositoryIsNilError  string = "product repository nil"
	orderRepositoryNilError      string = "order repository nil"

	createCampaignCommandNilError string = "create campaign command nil"
	createOrderCommandNilError    string = "create order command nil"
	createProductCommandNilError  string = "create product command nil"
	increaseTimeCommandNilError   string = "increase time command nil"

	createCampaignCommandHandlerNilError string = "create campaign command handler nil"
	createOrderCommandHandlerNilError    string = "create order command handler nil"
	createProductCommandHandlerNilError  string = "create product command handler nil"
	increaseTimeCommandHandlerNilError   string = "increase time command handler nil"

	getCampaignInfoQueryNilError string = "get campaign info query nil"
	getProductInfoQueryNilError  string = "get product info query nil"

	getCampaignInfoQueryHandlerNilError string = "get campaign info query handler nil"
	getProductInfoQueryHandlerNilError  string = "get product info query handler nil"

	systemTimeNilError string = "system time nil"

	campaignNotFoundError  string = "%s campaign not found"
	orderCouldNotCreated   string = "%s %d order could not created"
	productCouldNotCreated string = "%s product could not created"
	productNotFoundError   string = "%s product not found"
	productOutOfStockError string = "%s product out of stock"
)

func ThrowCampaignRepositoryCannotBeNilError() error {
	return errors.New(campaignRepositoryIsNilError)
}

func ThrowProductRepositoryCannotBeNil() error {
	return errors.New(productRepositoryIsNilError)
}

func ThrowOrderRepositoryCannotBeNilError() error {
	return errors.New(orderRepositoryNilError)
}

func ThrowCreateCampaignCommandHandlerCannotBeNilError() error {
	return errors.New(createCampaignCommandHandlerNilError)
}

func ThrowCreateOrderCommandHandlerCannotBeNilError() error {
	return errors.New(createOrderCommandHandlerNilError)
}

func ThrowCreateProductCommandHandlerCannotBeNilError() error {
	return errors.New(createProductCommandHandlerNilError)
}

func ThrowIncreaseHourCommandHandlerNilError() error {
	return errors.New(increaseTimeCommandHandlerNilError)
}

func ThrowGetCampaignInfoQueryHandlerCannotBeNilError() error {
	return errors.New(getCampaignInfoQueryHandlerNilError)
}

func ThrowGetProductInfoQueryHandlerCannotNilError() error {
	return errors.New(getProductInfoQueryHandlerNilError)
}

func ThrowSystemTimeCannotBeNilError() error {
	return errors.New(systemTimeNilError)
}

func ThrowCreateCampaignCommandCannotNilError() error {
	return errors.New(createCampaignCommandNilError)
}

func ThrowCreateOrderCommandCannotBeNilError() error {
	return errors.New(createOrderCommandNilError)
}

func ThrowIncreaseTimeCommandCannotBeNilError() error {
	return errors.New(increaseTimeCommandNilError)
}

func ThrowCreateProductCommandCannotBeNilError() error {
	return errors.New(createProductCommandNilError)
}

func ThrowGetCampaignInfoQueryNilError() error {
	return errors.New(getCampaignInfoQueryNilError)
}

func ThrowGetProductInfoQueryNilError() error {
	return errors.New(getProductInfoQueryNilError)
}

func ThrowCampaignNotFoundError(name string) error {
	return fmt.Errorf(campaignNotFoundError, name)
}

func ThrowOrderCouldNotCreated(pCode string, quantity int) error {
	return fmt.Errorf(orderCouldNotCreated, pCode, quantity)
}

func ThrowProductCouldNotCreated(pCode string) error {
	return fmt.Errorf(productCouldNotCreated, pCode)
}

func ThrowProductNotFoundError(pCode string) error {
	return fmt.Errorf(productNotFoundError, pCode)
}

func ThrowProductOutOfStockError(pCode string) error {
	return fmt.Errorf(productOutOfStockError, pCode)
}
