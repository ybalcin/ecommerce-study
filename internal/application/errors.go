package application

import "errors"

const (
	campaignRepositoryIsNilError string = "campaign repository nil"
	productRepositoryIsNilError  string = "product repository nil"
	orderRepositoryNilError      string = "order repository nil"

	createCampaignCommandNilError string = "create campaign command nil"
	createOrderCommandNilError    string = "create order command nil"
	createProductCommandNilError  string = "create product command nil"
	increaseHourCommandNilError   string = "increase hour command nil"

	createCampaignCommandHandlerNilError string = "create campaign command handler nil"
	createOrderCommandHandlerNilError    string = "create order command handler nil"
	createProductCommandHandlerNilError  string = "create product command handler nil"
	increaseHourCommandHandlerNilError   string = "increase hour command handler nil"

	getCampaignInfoQueryNilError string = "get campaign info query nil"
	getProductInfoQueryNilError  string = "get product info query nil"

	getCampaignInfoQueryHandlerNilError string = "get campaign info query handler nil"
	getProductInfoQueryHandlerNilError  string = "get product info query handler nil"

	systemTimeNilError string = "system time nil"
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
	return errors.New(increaseHourCommandHandlerNilError)
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

func ThrowIncreaseHourCommandCannotBeNilError() error {
	return errors.New(increaseHourCommandNilError)
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
