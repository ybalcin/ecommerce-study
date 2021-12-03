package errors

import "errors"

type DomainError struct {
	*ErrorBase
}

var (
	productCodeIsEmptyError    = errors.New("product code is empty")
	productPriceIsInvalidError = errors.New("product price value is invalid")
	productStockIsInvalidError = errors.New("product stock value is invalid")

	campaignNameIsEmptyError                     = errors.New("campaign name is empty")
	campaignDurationIsInvalidError               = errors.New("campaign duration is invalid")
	campaignPriceManipulationLimitIsInvalidError = errors.New("campaign price manipulation limit is invalid")
	campaignApplyProductCodesNotEqualError       = errors.New("campaign's product code and product's code not equal")

	orderProductCodeIsEmptyError = errors.New("product code of order is empty")
	orderQuantityIsInvalidError  = errors.New("quantity of order is invalid")
)

func ThrowProductCodeShouldNotBeEmptyError() *DomainError {
	return &DomainError{
		ErrorBase: NewError(productCodeIsEmptyError),
	}
}

func ThrowProductPriceValueIsInvalidError() *DomainError {
	return &DomainError{
		ErrorBase: NewError(productPriceIsInvalidError),
	}
}

func ThrowProductStockValueIsInvalidError() *DomainError {
	return &DomainError{
		ErrorBase: NewError(productStockIsInvalidError),
	}
}

func ThrowCampaignNameShouldNotBeEmptyError() *DomainError {
	return &DomainError{
		ErrorBase: NewError(campaignNameIsEmptyError),
	}
}

func ThrowCampaignDurationIsInvalidError() *DomainError {
	return &DomainError{
		ErrorBase: NewError(campaignDurationIsInvalidError),
	}
}

func ThrowCampaignPriceManipulationLimitIsInvalidError() *DomainError {
	return &DomainError{
		ErrorBase: NewError(campaignPriceManipulationLimitIsInvalidError),
	}
}

func ThrowOrderProductCodeIsEmptyError() *DomainError {
	return &DomainError{
		ErrorBase: NewError(orderProductCodeIsEmptyError),
	}
}

func ThrowOrderQuantityIsInvalidError() *DomainError {
	return &DomainError{
		ErrorBase: NewError(orderQuantityIsInvalidError),
	}
}

func ThrowCampaignApplyProductCodesNotEqualError() *DomainError {
	return &DomainError{
		ErrorBase: NewError(campaignApplyProductCodesNotEqualError),
	}
}
