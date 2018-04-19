package npsSdk

import (
	"encoding/xml"
	CONSTANTS "npsSdk/constants"
	productionPack "github.com/Ingenico-NPS-Latam/nps-sdk-go/src/npsSdk/services/production"
	sandboxPack "github.com/Ingenico-NPS-Latam/nps-sdk-go/src/npsSdk/services/sandbox"
	stagingPack "github.com/Ingenico-NPS-Latam/nps-sdk-go/src/npsSdk/services/staging"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name


//func NewTaxesRequestStruct() *TaxesRequestStruct {
func NewTaxesRequestStruct() interface{} {
	//var p *TaxesRequestStruct
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p:=productionPack.NewTaxesRequestStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p:=stagingPack.NewTaxesRequestStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p:=sandboxPack.NewTaxesRequestStruct()
		return(p)
	}

        return nil

}

func NewArrayOf_TaxesRequestStruct() interface{} { //*ArrayOf_TaxesRequestStruct {
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewArrayOf_TaxesRequestStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewArrayOf_TaxesRequestStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewArrayOf_TaxesRequestStruct()
		return(p)
	}
	//	p.Xsi_Type_SOAP = "SOAP-ENC:Array"
	//	p.Xsi_Type_SOAP_ARR = "tns:ArrayOf_TaxesRequestStruct"

	return nil
}

func NewAmountAdditionalDetailsRequestStruct() interface{}{
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewAmountAdditionalDetailsRequestStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewAmountAdditionalDetailsRequestStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewAmountAdditionalDetailsRequestStruct()
		return(p)
	}

	return nil
}

func NewAddressStruct() interface{} { //*AddressStruct {
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewAddressStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewAddressStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewAddressStruct()
		return(p)
	}
	//p := new(AddressStruct)
	return nil
}

func NewSellerDetailsStruct() interface{} { //*SellerDetailsStruct {
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewSellerDetailsStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewSellerDetailsStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewSellerDetailsStruct()
		return(p)
	}
	return nil
}

func NewMerchantAdditionalDetailsStruct() interface{} { //*MerchantAdditionalDetailsStruct {
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewMerchantAdditionalDetailsStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewMerchantAdditionalDetailsStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewMerchantAdditionalDetailsStruct()
		return(p)
	}
	return nil
}

func NewCustomerAdditionalDetailsStruct() interface{} { //*CustomerAdditionalDetailsStruct {
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewCustomerAdditionalDetailsStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewCustomerAdditionalDetailsStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewCustomerAdditionalDetailsStruct()
		return(p)
	}
	return nil
}


func NewPersonStruct() interface{} { //*PersonStruct {

	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
               p:=productionPack.NewPersonStruct()
               return(p)
	case CONSTANTS.STAGING_ENV:
               p:=stagingPack.NewPersonStruct()
               return(p)
	case CONSTANTS.SANDBOX_ENV:
               p:=sandboxPack.NewPersonStruct()
               return(p)
	}
	return nil
}


func NewBillingDetailsStruct() interface{} { //*BillingDetailsStruct {
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p:=productionPack.NewBillingDetailsStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p:=stagingPack.NewBillingDetailsStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p:=sandboxPack.NewBillingDetailsStruct()
		return(p)
	}
	return nil
}

func NewShippingDetailsStruct() interface{} { //
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewShippingDetailsStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewShippingDetailsStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewShippingDetailsStruct()
		return(p)
	}
	return nil
}

func NewOrderItemStruct() interface{} { //*OrderItemStruct {
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewOrderItemStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewOrderItemStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewOrderItemStruct()
		return(p)
	}
	return nil
}

func NewArrayOf_OrderItemStruct() interface{} { //*ArrayOf_OrderItemStruct {
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewArrayOf_OrderItemStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewArrayOf_OrderItemStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewArrayOf_OrderItemStruct()
		return(p)
	}

	return nil
}

func NewOrderDetailsStruct() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewOrderDetailsStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewOrderDetailsStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewOrderDetailsStruct()
		return(p)
	}
	return nil
}

func NewLegStruct() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewLegStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewLegStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewLegStruct()
		return(p)
	}
	return nil
}

func NewArrayOf_LegStruct() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewArrayOf_LegStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewArrayOf_LegStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewArrayOf_LegStruct()
		return(p)
	}

	return nil
}

func NewPassengerStruct() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewPassengerStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewPassengerStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewPassengerStruct()
		return(p)
	}
	return nil
}

func NewArrayOf_PassengerStruct() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewArrayOf_PassengerStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewArrayOf_PassengerStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewArrayOf_PassengerStruct()
		return(p)
	}

	return nil
}

func NewAirlineTicketStruct() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewAirlineTicketStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewAirlineTicketStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewAirlineTicketStruct()
		return(p)
	}
	return nil
}

func NewAirlineDetailsStruct() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewAirlineDetailsStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewAirlineDetailsStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewAirlineDetailsStruct()
		return(p)
	}
	return nil
}

func NewVaultReference2pStruct() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewVaultReference2pStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewVaultReference2pStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewVaultReference2pStruct()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_PayOnLine_2p() interface{} { //*RequerimientoStruct_PayOnLine_2p {
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_PayOnLine_2p()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_PayOnLine_2p()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_PayOnLine_2p()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_Authorize_2p()  interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_Authorize_2p()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_Authorize_2p()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_Authorize_2p()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_BankPayment_2p()  interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_BankPayment_2p()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_BankPayment_2p()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_BankPayment_2p()
		return(p)
	}
	return nil
}

func NewVaultReference3pStruct()   interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewVaultReference3pStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewVaultReference3pStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewVaultReference3pStruct()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_SplitPayOnLine_2p_Transactions() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_SplitPayOnLine_2p_Transactions()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_SplitPayOnLine_2p_Transactions()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_SplitPayOnLine_2p_Transactions()
		return(p)
	}
	return nil
}

func NewArrayOf_RequerimientoStruct_SplitPayOnLine_2p_Transactions() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewArrayOf_RequerimientoStruct_SplitPayOnLine_2p_Transactions()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewArrayOf_RequerimientoStruct_SplitPayOnLine_2p_Transactions()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewArrayOf_RequerimientoStruct_SplitPayOnLine_2p_Transactions()
		return(p)
	}

	return nil
}

func NewRequerimientoStruct_SplitPayOnLine_2p()  interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_SplitPayOnLine_2p()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_SplitPayOnLine_2p()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_SplitPayOnLine_2p()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_SplitAuthorize_2p_Transactions() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_SplitAuthorize_2p_Transactions()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_SplitAuthorize_2p_Transactions()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_SplitAuthorize_2p_Transactions()
		return(p)
	}
	return nil
}

func NewArrayOf_RequerimientoStruct_SplitAuthorize_2p_Transactions() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewArrayOf_RequerimientoStruct_SplitAuthorize_2p_Transactions()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewArrayOf_RequerimientoStruct_SplitAuthorize_2p_Transactions()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewArrayOf_RequerimientoStruct_SplitAuthorize_2p_Transactions()
		return(p)
	}

	return nil
}

func NewRequerimientoStruct_SplitAuthorize_2p() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_SplitAuthorize_2p()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_SplitAuthorize_2p()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_SplitAuthorize_2p()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_PayOnLine_3p() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_PayOnLine_3p()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_PayOnLine_3p()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_PayOnLine_3p()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_SplitPayOnLine_3p_Transactions() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_SplitPayOnLine_3p_Transactions()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_SplitPayOnLine_3p_Transactions()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_SplitPayOnLine_3p_Transactions()
		return(p)
	}

	return nil
}

func NewArrayOf_RequerimientoStruct_SplitPayOnLine_3p_Transactions() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewArrayOf_RequerimientoStruct_SplitPayOnLine_3p_Transactions()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewArrayOf_RequerimientoStruct_SplitPayOnLine_3p_Transactions()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewArrayOf_RequerimientoStruct_SplitPayOnLine_3p_Transactions()
		return(p)
	}

	return nil
}

func NewRequerimientoStruct_SplitPayOnLine_3p() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_SplitPayOnLine_3p()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_SplitPayOnLine_3p()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_SplitPayOnLine_3p()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_Authorize_3p() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_Authorize_3p()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_Authorize_3p()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_Authorize_3p()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_SplitAuthorize_3p_Transactions() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_SplitAuthorize_3p_Transactions()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_SplitAuthorize_3p_Transactions()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_SplitAuthorize_3p_Transactions()
		return(p)
	}
	return nil
}

func NewArrayOf_RequerimientoStruct_SplitAuthorize_3p_Transactions() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewArrayOf_RequerimientoStruct_SplitAuthorize_3p_Transactions()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewArrayOf_RequerimientoStruct_SplitAuthorize_3p_Transactions()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewArrayOf_RequerimientoStruct_SplitAuthorize_3p_Transactions()
		return(p)
	}

	return nil
}

func NewRequerimientoStruct_SplitAuthorize_3p() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_SplitAuthorize_3p()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_SplitAuthorize_3p()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_SplitAuthorize_3p()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_BankPayment_3p() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_BankPayment_3p()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_BankPayment_3p()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_BankPayment_3p()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_CashPayment_3p() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_CashPayment_3p()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_CashPayment_3p()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_CashPayment_3p()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_Capture() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_Capture()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_Capture()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_Capture()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_Refund() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_Refund()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_Refund()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_Refund()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_GetInstallmentsOptions() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_GetInstallmentsOptions()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_GetInstallmentsOptions()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_GetInstallmentsOptions()
		return(p)
	}
	return nil
}

func NewInstallmentsOptionsStruct() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewInstallmentsOptionsStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewInstallmentsOptionsStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewInstallmentsOptionsStruct()
		return(p)
	}
	return nil
}

func NewArrayOf_InstallmentsOptionsStruct() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewArrayOf_InstallmentsOptionsStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewArrayOf_InstallmentsOptionsStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewArrayOf_InstallmentsOptionsStruct()
		return(p)
	}

	return nil
}

func NewCardInputDetailsStruct() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewCardInputDetailsStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewCardInputDetailsStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewCardInputDetailsStruct()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_CreatePaymentMethodToken() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_CreatePaymentMethodToken()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_CreatePaymentMethodToken()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_CreatePaymentMethodToken()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_RetrievePaymentMethodToken() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_RetrievePaymentMethodToken()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_RetrievePaymentMethodToken()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_RetrievePaymentMethodToken()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_RecachePaymentMethodToken() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_RecachePaymentMethodToken()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_RecachePaymentMethodToken()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_RecachePaymentMethodToken()
		return(p)
	}
	return nil

}

func NewRequerimientoStruct_CreateClientSession() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_CreateClientSession()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_CreateClientSession()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_CreateClientSession()
		return(p)
	}
	return nil
}

func NewPaymentMethodInputDetailsStruct() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewPaymentMethodInputDetailsStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewPaymentMethodInputDetailsStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewPaymentMethodInputDetailsStruct()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_CreatePaymentMethod() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_CreatePaymentMethod()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_CreatePaymentMethod()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_CreatePaymentMethod()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_CreatePaymentMethodFromPayment() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_CreatePaymentMethodFromPayment()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_CreatePaymentMethodFromPayment()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_CreatePaymentMethodFromPayment()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_RetrievePaymentMethod() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_RetrievePaymentMethod()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_RetrievePaymentMethod()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_RetrievePaymentMethod()
		return(p)
	}
	return nil
}

func NewCardInputUpdateDetailsStruct() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewCardInputUpdateDetailsStruct()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewCardInputUpdateDetailsStruct()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewCardInputUpdateDetailsStruct()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_UpdatePaymentMethod() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_UpdatePaymentMethod()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_UpdatePaymentMethod()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_UpdatePaymentMethod()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_DeletePaymentMethod() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_DeletePaymentMethod()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_DeletePaymentMethod()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_DeletePaymentMethod()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_CreateCustomer() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_CreateCustomer()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_CreateCustomer()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_CreateCustomer()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_RetrieveCustomer() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_RetrieveCustomer()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_RetrieveCustomer()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_RetrieveCustomer()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_UpdateCustomer() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_UpdateCustomer()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_UpdateCustomer()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_UpdateCustomer()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_DeleteCustomer() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_DeleteCustomer()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_DeleteCustomer()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_DeleteCustomer()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_SimpleQueryTx() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_SimpleQueryTx()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_SimpleQueryTx()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_SimpleQueryTx()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_QueryCardNumber() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_QueryCardNumber()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_QueryCardNumber()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_QueryCardNumber()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_QueryCardDetails() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_QueryCardDetails()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_QueryCardDetails()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_QueryCardDetails()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_QueryTxs() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_QueryTxs()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_QueryTxs()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_QueryTxs()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_GetIINDetails() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_GetIINDetails()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_GetIINDetails()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_GetIINDetails()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_ChangeSecretKey() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_ChangeSecretKey()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_ChangeSecretKey()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_ChangeSecretKey()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_FraudScreening() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_FraudScreening()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_FraudScreening()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_FraudScreening()
		return(p)
	}
	return nil
}

func NewRequerimientoStruct_NotifyFraudScreeningReview() interface{} { 
	switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
		p := productionPack.NewRequerimientoStruct_NotifyFraudScreeningReview()
		return(p)
	case CONSTANTS.STAGING_ENV:
		p := stagingPack.NewRequerimientoStruct_NotifyFraudScreeningReview()
		return(p)
	case CONSTANTS.SANDBOX_ENV:
		p := sandboxPack.NewRequerimientoStruct_NotifyFraudScreeningReview()
		return(p)
	}
	return nil
}

