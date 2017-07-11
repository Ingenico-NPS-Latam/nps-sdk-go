package main

import (
	"fmt"
	"log"
	"nps"
	CONSTANTS "nps/constants"
)

func SendPayOnLine_2p(service *nps.PaymentServicePlatformPortType) error {

	Person := nps.NewPersonStruct()
	Person.FirstName = "Silvina"
	Person.LastName = "Falconi"
	Person.PhoneNumber1 = "52520960"

	AmountAdditionalDetails := nps.NewAmountAdditionalDetailsRequestStruct()
	AmountAdditionalDetails.Tip = "10"
	AmountAdditionalDetails.Discount = "5"

	Billing := nps.NewBillingDetailsStruct()
	Billing.Invoice = "100001234"
	//Billing.InvoiceDate = "2017-06-29 12:00:00"
	Billing.InvoiceAmount = "990"
	Billing.InvoiceCurrency = "032"
	Billing.Person = Person

	SellerAddress := nps.NewAddressStruct()
	SellerAddress.City = "CABA"
	SellerAddress.Country = "ARG"
	SellerAddress.Street = "SellerStreet"
	SellerAddress.HouseNumber = "1234"

	SellerDetails := nps.NewSellerDetailsStruct()
	SellerDetails.Name = "Seller Name"
	SellerDetails.Address = SellerAddress

	MerchantAdditionalDetails := nps.NewMerchantAdditionalDetailsStruct()
	MerchantAdditionalDetails.ShoppingCartInfo = "ShoppingCartInfo"
	MerchantAdditionalDetails.SellerDetails = SellerDetails

	CustomerAdditionalDetails := nps.NewCustomerAdditionalDetailsStruct()
	CustomerAdditionalDetails.EmailAddress = "mailAddr@mail.com.ar"

	order1 := nps.NewOrderItemStruct()
	order1.Description = "producto 1"
	order1.UnitPrice = "10"

	order2 := nps.NewOrderItemStruct()
	order2.Description = "producto 2"
	order2.UnitPrice = "20"

	OrderDetails := nps.NewOrderDetailsStruct()
	OrderDetails.OrderItems = nps.NewArrayOf_OrderItemStruct()
	OrderDetails.OrderItems.Items = make([]*nps.OrderItemStruct, 0)
	OrderDetails.OrderItems.Items = append(OrderDetails.OrderItems.Items, order1)
	OrderDetails.OrderItems.Items = append(OrderDetails.OrderItems.Items, order2)

	payOnline2p := nps.NewRequerimientoStruct_PayOnLine_2p()

	payOnline2p.Psp_Version = "2.2"
	payOnline2p.Psp_MerchantId = "psp_test"
	payOnline2p.Psp_TxSource = "WEB"
	payOnline2p.Psp_MerchTxRef = "ORDER56666-3"
	payOnline2p.Psp_MerchOrderId = "ORDER56666"
	payOnline2p.Psp_Amount = "1000"
	payOnline2p.Psp_NumPayments = "1"
	payOnline2p.Psp_Currency = "032"
	payOnline2p.Psp_Country = "ARG"
	payOnline2p.Psp_Product = "14"
	payOnline2p.Psp_CustomerMail = "yourmail@gmail"
	payOnline2p.Psp_CardNumber = "4507990000000010"
	payOnline2p.Psp_CardExpDate = "1903"
	payOnline2p.Psp_CardSecurityCode = "306"
	payOnline2p.Psp_SoftDescriptor = "Sol Tropical E"
	payOnline2p.Psp_PosDateTime = "2016-12-01 12:00:00"

	payOnline2p.Psp_OrderDetails = OrderDetails
	payOnline2p.Psp_CustomerAdditionalDetails = CustomerAdditionalDetails
	payOnline2p.Psp_AmountAdditionalDetails = AmountAdditionalDetails
	payOnline2p.Psp_BillingDetails = Billing
	payOnline2p.Psp_MerchantAdditionalDetails = MerchantAdditionalDetails

	resp, err := service.PayOnLine_2p(payOnline2p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendSplitPayOnLine_3p(service *nps.PaymentServicePlatformPortType) error {

	splitPayOnline3p := nps.NewRequerimientoStruct_SplitPayOnLine_3p()

	splitPayOnline3p.Psp_Version = "2.2"
	splitPayOnline3p.Psp_MerchantId = "psp_test"
	splitPayOnline3p.Psp_TxSource = "WEB"
	splitPayOnline3p.Psp_MerchOrderId = "20170609134500"
	splitPayOnline3p.Psp_ReturnURL = "http://localhost/"
	splitPayOnline3p.Psp_FrmLanguage = "es_AR"
	splitPayOnline3p.Psp_Amount = "15050"
	splitPayOnline3p.Psp_Currency = "032"
	splitPayOnline3p.Psp_Country = "ARG"
	splitPayOnline3p.Psp_Product = "14"
	splitPayOnline3p.Psp_PosDateTime = "2016-12-01 12:00:00"

	trn := nps.NewRequerimientoStruct_SplitPayOnLine_3p_Transactions()
	trn.Psp_MerchantId = "psp_test"
	trn.Psp_MerchTxRef = "ORDER66666-5"
	trn.Psp_Product = "14"
	trn.Psp_Amount = "10000"
	trn.Psp_NumPayments = "1"

	trn2 := nps.NewRequerimientoStruct_SplitPayOnLine_3p_Transactions()
	trn2.Psp_MerchantId = "psp_test"
	trn2.Psp_MerchTxRef = "ORDER66666-6"
	trn2.Psp_Product = "14"
	trn2.Psp_Amount = "5050"
	trn2.Psp_NumPayments = "1"

	Transactions := nps.NewArrayOf_RequerimientoStruct_SplitPayOnLine_3p_Transactions()
	Transactions.Items = make([]*nps.RequerimientoStruct_SplitPayOnLine_3p_Transactions, 0)

	Transactions.Items = append(Transactions.Items, trn)
	Transactions.Items = append(Transactions.Items, trn2)

	splitPayOnline3p.Psp_Transactions = Transactions

	resp, err := service.SplitPayOnLine_3p(splitPayOnline3p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)

	for i := 0; i < len(resp.Psp_Transactions.Items); i++ {
		f := resp.Psp_Transactions.Items[i]
		fmt.Printf("========== item[%d]========== \n", i)
		fmt.Printf("Psp_MerchantId         [%s]\n", f.Psp_MerchantId)
		fmt.Printf("Psp_TransactionId      [%s]\n", f.Psp_TransactionId)
		fmt.Printf("Psp_MerchTxRef         [%s]\n", f.Psp_MerchTxRef)
		fmt.Printf("Psp_MerchAdditionalRef [%s]\n", f.Psp_MerchAdditionalRef)
		fmt.Printf("Psp_Product            [%s]\n", f.Psp_Product)
		fmt.Printf("Psp_PromotionCode      [%s]\n", f.Psp_PromotionCode)
		fmt.Printf("Psp_Amount             [%s]\n", f.Psp_Amount)
		fmt.Printf("Psp_NumPayments        [%s]\n", f.Psp_NumPayments)
		fmt.Printf("Psp_Plan               [%s]\n", f.Psp_Plan)
		fmt.Printf("Psp_FirstPaymentDeferral [%s]\n", f.Psp_FirstPaymentDeferral)
		fmt.Printf("Psp_SoftDescriptor     [%s]\n", f.Psp_SoftDescriptor)
		fmt.Printf("Psp_CreatedAt          [%s]\n", f.Psp_CreatedAt)

	}

	return nil
}

func SendAuthorize_2p(service *nps.PaymentServicePlatformPortType) error {

	authorize2p := nps.NewRequerimientoStruct_Authorize_2p()

	authorize2p.Psp_Version = "2.2"
	authorize2p.Psp_MerchantId = "psp_test"
	authorize2p.Psp_TxSource = "WEB"
	authorize2p.Psp_MerchTxRef = "20170630103600"
	authorize2p.Psp_MerchOrderId = "20170630103600-1"
	authorize2p.Psp_Amount = "15050"
	authorize2p.Psp_NumPayments = "1"
	authorize2p.Psp_Currency = "032"
	authorize2p.Psp_Country = "ARG"
	authorize2p.Psp_Product = "14"
	authorize2p.Psp_CardNumber = "4507990000000010"
	authorize2p.Psp_CardExpDate = "1812"
	authorize2p.Psp_PosDateTime = "2016-12-01 12:00:00"

	CustomerAdditionalDetails := nps.NewCustomerAdditionalDetailsStruct()
	CustomerAdditionalDetails.EmailAddress = "CustomerEmail@email.com.ar"

	Person := nps.NewPersonStruct()
	Person.FirstName = "Silvina"
	Person.LastName = "Falconi"
	Person.PhoneNumber1 = "52520960"

	Billing := nps.NewBillingDetailsStruct()
	Billing.Invoice = "100001234"
	//Billing.InvoiceDate = "20170601"
	Billing.InvoiceAmount = "990"
	Billing.InvoiceCurrency = "032"
	Billing.Person = Person

	Taxes := nps.NewArrayOf_TaxesRequestStruct()
	Taxes.Items = make([]*nps.TaxesRequestStruct, 0)

	tax1 := nps.NewTaxesRequestStruct()
	tax1.TypeId = "500"
	tax1.Amount = "50"
	tax1.Rate = "10"

	Taxes.Items = append(Taxes.Items, tax1)

	AmountAdditionalDetails := nps.NewAmountAdditionalDetailsRequestStruct()
	AmountAdditionalDetails.Tip = "10"
	AmountAdditionalDetails.Discount = "5"
	AmountAdditionalDetails.Taxes = Taxes

	SellerAddress := nps.NewAddressStruct()
	SellerAddress.City = "CABA"
	SellerAddress.Country = "ARG"
	SellerAddress.Street = "SellerStreet"
	SellerAddress.HouseNumber = "1234"

	SellerDetails := nps.NewSellerDetailsStruct()
	SellerDetails.Name = "Seller Name"
	SellerDetails.Address = SellerAddress

	MerchantAdditionalDetails := nps.NewMerchantAdditionalDetailsStruct()
	MerchantAdditionalDetails.ShoppingCartInfo = "ShoppingCartInfo"
	MerchantAdditionalDetails.SellerDetails = SellerDetails

	order1 := nps.NewOrderItemStruct()
	order1.Description = "producto 1"
	order1.UnitPrice = "10"

	order2 := nps.NewOrderItemStruct()
	order2.Description = "producto 2"
	order2.UnitPrice = "20"

	OrderDetails := nps.NewOrderDetailsStruct()
	OrderDetails.OrderItems = nps.NewArrayOf_OrderItemStruct()
	OrderDetails.OrderItems.Items = make([]*nps.OrderItemStruct, 0)
	OrderDetails.OrderItems.Items = append(OrderDetails.OrderItems.Items, order1)
	OrderDetails.OrderItems.Items = append(OrderDetails.OrderItems.Items, order2)

	ShippingDetails := nps.NewShippingDetailsStruct()

	ShippingAddress := nps.NewAddressStruct()
	ShippingAddress.Street = "shipping street"
	ShippingAddress.HouseNumber = "1234"
	ShippingAddress.City = "CABA"
	ShippingAddress.Country = "ARG"
	ShippingAddress.AdditionalInfo = "AdditionalInfo of shipping"

	PrimaryRecipient := nps.NewPersonStruct()
	PrimaryRecipient.FirstName = "Pepe"
	PrimaryRecipient.LastName = "Juan"

	ShippingDetails.Method = "10"
	ShippingDetails.Address = ShippingAddress
	ShippingDetails.PrimaryRecipient = PrimaryRecipient

	authorize2p.Psp_AmountAdditionalDetails = AmountAdditionalDetails
	authorize2p.Psp_BillingDetails = Billing
	authorize2p.Psp_CustomerAdditionalDetails = CustomerAdditionalDetails
	authorize2p.Psp_MerchantAdditionalDetails = MerchantAdditionalDetails
	authorize2p.Psp_OrderDetails = OrderDetails
	authorize2p.Psp_ShippingDetails = ShippingDetails

	resp, err := service.Authorize_2p(authorize2p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendCreateClientSession(service *nps.PaymentServicePlatformPortType) error {

	createClientSession := nps.NewRequerimientoStruct_CreateClientSession()

	createClientSession.Psp_Version = "2.2"
	createClientSession.Psp_MerchantId = "psp_test"
	createClientSession.Psp_PosDateTime = "2017-06-19 12:00:00"

	resp, err := service.CreateClientSession(createClientSession)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	fmt.Printf("ClientSession = [%s]\n", resp.Psp_ClientSession)
	return nil
}

func SendCreatePaymentMethod(service *nps.PaymentServicePlatformPortType) error {

	createPaymentMethod := nps.NewRequerimientoStruct_CreatePaymentMethod()

	createPaymentMethod.Psp_Version = "2.2"
	createPaymentMethod.Psp_MerchantId = "psp_test"
	createPaymentMethod.Psp_PosDateTime = "2017-06-19 12:00:00"

	paymentMethod := nps.NewPaymentMethodInputDetailsStruct()
	cardInputDetails := nps.NewCardInputDetailsStruct()
	cardInputDetails.HolderName = "tester"
	cardInputDetails.ExpirationDate = "1812"
	cardInputDetails.Number = "4507990000000010"
	cardInputDetails.SecurityCode = "123"

	paymentMethod.CardInputDetails = cardInputDetails

	createPaymentMethod.Psp_PaymentMethod = paymentMethod

	resp, err := service.CreatePaymentMethod(createPaymentMethod)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func CreatePaymentMethodToken(service *nps.PaymentServicePlatformPortType) error {

	createPaymentMethodToken := nps.NewRequerimientoStruct_CreatePaymentMethodToken()

	createPaymentMethodToken.Psp_Version = "2.2"
	createPaymentMethodToken.Psp_MerchantId = "psp_test"
	createPaymentMethodToken.Psp_Product = "14"
	createPaymentMethodToken.Psp_ClientSession = "ku4sGsxTKlbYQ9PJxwS4FGduR2YZdyMtQn8xcurxFRmCUuUXj6BP9wtYLAfcGJew"

	cardInputDetails := nps.NewCardInputDetailsStruct()
	cardInputDetails.ExpirationDate = "1812"
	cardInputDetails.HolderName = "Silvina Falconi"
	cardInputDetails.Number = "4507990000000010"
	cardInputDetails.SecurityCode = "123"
	createPaymentMethodToken.Psp_CardInputDetails = cardInputDetails

	person := nps.NewPersonStruct()
	person.FirstName = "Silvina"
	person.LastName = "Falconi"
	person.PhoneNumber1 = "52520960"
	createPaymentMethodToken.Psp_Person = person

	address := nps.NewAddressStruct()
	address.Street = "pelegrini"
	address.City = "CABA"
	address.Country = "ARG"
	address.HouseNumber = "1111"
	createPaymentMethodToken.Psp_Address = address

	resp, err := service.CreatePaymentMethodToken(createPaymentMethodToken)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendBankPayment_2p(service *nps.PaymentServicePlatformPortType) error {

	Taxes := nps.NewArrayOf_TaxesRequestStruct()
	Taxes.Items = make([]*nps.TaxesRequestStruct, 0)

	tax1 := nps.NewTaxesRequestStruct()
	tax1.TypeId = "500"
	tax1.Amount = "50"
	tax1.Rate = "10"

	Taxes.Items = append(Taxes.Items, tax1)

	AmountAdditionalDetails := nps.NewAmountAdditionalDetailsRequestStruct()
	AmountAdditionalDetails.Tip = "10"
	AmountAdditionalDetails.Discount = "5"
	AmountAdditionalDetails.Taxes = Taxes

	Person := nps.NewPersonStruct()
	Person.FirstName = "Silvina"
	Person.LastName = "Falconi"
	Person.PhoneNumber1 = "52520960"
	Person.IDNumber = "11111111"
	Person.IDType = "100"

	Billing := nps.NewBillingDetailsStruct()
	Billing.Invoice = "100001234"
	//Billing.InvoiceDate = "20170601"
	Billing.InvoiceAmount = "990"
	Billing.InvoiceCurrency = "032"
	Billing.Person = Person

	CustomerAdditionalDetails := nps.NewCustomerAdditionalDetailsStruct()
	CustomerAdditionalDetails.EmailAddress = "CustomerEmail@email.com.ar"

	BankPayment_2p := nps.NewRequerimientoStruct_BankPayment_2p()
	BankPayment_2p.Psp_Version = "2.2"
	BankPayment_2p.Psp_Product = "320"
	BankPayment_2p.Psp_MerchantId = "psp_test"
	BankPayment_2p.Psp_TxSource = "TxSource"
	BankPayment_2p.Psp_MerchTxRef = "11111"
	BankPayment_2p.Psp_MerchOrderId = "ORD-10001"
	BankPayment_2p.Psp_ScreenDescription = "SCR DESC"
	BankPayment_2p.Psp_TicketDescription = "ticket desc"
	BankPayment_2p.Psp_CustomerBankId = "HSBC"

	BankPayment_2p.Psp_Amount1 = "10000"
	BankPayment_2p.Psp_Amount2 = "20000"
	BankPayment_2p.Psp_Country = "ARG"
	BankPayment_2p.Psp_Currency = "032"
	BankPayment_2p.Psp_PosDateTime = "2016-12-01 12:00:00"
	BankPayment_2p.Psp_TxSource = "WEB"

	BankPayment_2p.Psp_AmountAdditionalDetails = AmountAdditionalDetails
	BankPayment_2p.Psp_BillingDetails = Billing
	BankPayment_2p.Psp_CustomerAdditionalDetails = CustomerAdditionalDetails

	resp, err := service.BankPayment_2p(BankPayment_2p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendSplitPayOnLine_2p(service *nps.PaymentServicePlatformPortType) error {

	VaultReference := nps.NewVaultReference2pStruct()
	//VaultReference.CustomerId = "5555"
	VaultReference.PaymentMethodToken = "pm-token_lCY303k0vHvS5W06sPwzgoSHNt0VRrkG"

	trn := nps.NewRequerimientoStruct_SplitPayOnLine_2p_Transactions()
	trn.Psp_MerchantId = "psp_test"
	trn.Psp_MerchTxRef = "ORDER66666-5"
	trn.Psp_Product = "14"
	trn.Psp_Amount = "10000"
	trn.Psp_NumPayments = "1"
	trn.Psp_VaultReference = VaultReference

	trn2 := nps.NewRequerimientoStruct_SplitPayOnLine_2p_Transactions()
	trn2.Psp_MerchantId = "psp_test"
	trn2.Psp_MerchTxRef = "ORDER66666-6"
	trn2.Psp_Product = "14"
	trn2.Psp_Amount = "5050"
	trn2.Psp_NumPayments = "1"
	trn2.Psp_VaultReference = VaultReference

	Transactions := nps.NewArrayOf_RequerimientoStruct_SplitPayOnLine_2p_Transactions()
	Transactions.Items = make([]*nps.RequerimientoStruct_SplitPayOnLine_2p_Transactions, 0)

	Transactions.Items = append(Transactions.Items, trn)
	Transactions.Items = append(Transactions.Items, trn2)

	SplitPayOnLine_2p := nps.NewRequerimientoStruct_SplitPayOnLine_2p()
	SplitPayOnLine_2p.Psp_Version = "2.2"
	SplitPayOnLine_2p.Psp_Product = "14"
	SplitPayOnLine_2p.Psp_MerchantId = "psp_test"
	SplitPayOnLine_2p.Psp_Amount = "1000"
	SplitPayOnLine_2p.Psp_CardHolderName = "holder"
	SplitPayOnLine_2p.Psp_CardNumber = "4507990000000010"
	SplitPayOnLine_2p.Psp_CardExpDate = "1903"
	SplitPayOnLine_2p.Psp_CardSecurityCode = "306"
	SplitPayOnLine_2p.Psp_TxSource = "WEB"
	SplitPayOnLine_2p.Psp_MerchOrderId = "ORD-20170629123000"
	SplitPayOnLine_2p.Psp_Currency = "032"
	SplitPayOnLine_2p.Psp_Country = "ARG"
	SplitPayOnLine_2p.Psp_PosDateTime = "2016-12-01 12:00:00"

	//SplitPayOnLine_2p.Psp_VaultReference = VaultReference
	SplitPayOnLine_2p.Psp_Transactions = Transactions

	resp, err := service.SplitPayOnLine_2p(SplitPayOnLine_2p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendSplitAuthorize_2p(service *nps.PaymentServicePlatformPortType) error {

	CustomerAdditionalDetails := nps.NewCustomerAdditionalDetailsStruct()
	CustomerAdditionalDetails.EmailAddress = "CustomerEmail@email.com.ar"

	Person := nps.NewPersonStruct()
	Person.FirstName = "Silvina"
	Person.LastName = "Falconi"
	Person.PhoneNumber1 = "52520960"

	Billing := nps.NewBillingDetailsStruct()
	Billing.Invoice = "100001234"
	//Billing.InvoiceDate = "20170601"
	Billing.InvoiceAmount = "990"
	Billing.InvoiceCurrency = "032"
	Billing.Person = Person

	SellerAddress := nps.NewAddressStruct()
	SellerAddress.City = "CABA"
	SellerAddress.Country = "ARG"
	SellerAddress.Street = "SellerStreet"
	SellerAddress.HouseNumber = "1234"

	SellerDetails := nps.NewSellerDetailsStruct()
	SellerDetails.Name = "Seller Name"
	SellerDetails.Address = SellerAddress

	MerchantAdditionalDetails := nps.NewMerchantAdditionalDetailsStruct()
	MerchantAdditionalDetails.ShoppingCartInfo = "ShoppingCartInfo"
	MerchantAdditionalDetails.SellerDetails = SellerDetails

	order1 := nps.NewOrderItemStruct()
	order1.Description = "producto 1"
	order1.UnitPrice = "10"

	order2 := nps.NewOrderItemStruct()
	order2.Description = "producto 2"
	order2.UnitPrice = "20"

	OrderDetails := nps.NewOrderDetailsStruct()
	OrderDetails.OrderItems = nps.NewArrayOf_OrderItemStruct()
	OrderDetails.OrderItems.Items = make([]*nps.OrderItemStruct, 0)
	OrderDetails.OrderItems.Items = append(OrderDetails.OrderItems.Items, order1)
	OrderDetails.OrderItems.Items = append(OrderDetails.OrderItems.Items, order2)

	ShippingDetails := nps.NewShippingDetailsStruct()

	ShippingAddress := nps.NewAddressStruct()
	ShippingAddress.Street = "shipping street"
	ShippingAddress.HouseNumber = "1234"
	ShippingAddress.City = "CABA"
	ShippingAddress.Country = "ARG"
	ShippingAddress.AdditionalInfo = "AdditionalInfo of shipping"

	PrimaryRecipient := nps.NewPersonStruct()
	PrimaryRecipient.FirstName = "Pepe"
	PrimaryRecipient.LastName = "Juan"

	ShippingDetails.Method = "10"
	ShippingDetails.Address = ShippingAddress
	ShippingDetails.PrimaryRecipient = PrimaryRecipient

	VaultReference := nps.NewVaultReference2pStruct()
	VaultReference.PaymentMethodToken = "pm-token_lCY303k0vHvS5W06sPwzgoSHNt0VRrkG"

	trn := nps.NewRequerimientoStruct_SplitAuthorize_2p_Transactions()
	trn.Psp_MerchantId = "psp_test"
	trn.Psp_MerchTxRef = "ORDER66666-5"
	trn.Psp_Product = "14"
	trn.Psp_Amount = "10000"
	trn.Psp_NumPayments = "1"
	trn.Psp_VaultReference = VaultReference

	trn2 := nps.NewRequerimientoStruct_SplitAuthorize_2p_Transactions()
	trn2.Psp_MerchantId = "psp_test"
	trn2.Psp_MerchTxRef = "ORDER66666-6"
	trn2.Psp_Product = "14"
	trn2.Psp_Amount = "5050"
	trn2.Psp_NumPayments = "1"
	trn2.Psp_VaultReference = VaultReference

	Transactions := nps.NewArrayOf_RequerimientoStruct_SplitAuthorize_2p_Transactions()
	Transactions.Items = make([]*nps.RequerimientoStruct_SplitAuthorize_2p_Transactions, 0)

	Transactions.Items = append(Transactions.Items, trn)
	Transactions.Items = append(Transactions.Items, trn2)

	SplitAuthorize2p := nps.NewRequerimientoStruct_SplitAuthorize_2p()

	SplitAuthorize2p.Psp_Version = "2.2"
	SplitAuthorize2p.Psp_MerchantId = "psp_test"
	SplitAuthorize2p.Psp_TxSource = "WEB"
	SplitAuthorize2p.Psp_MerchOrderId = "20170609150900-1"
	SplitAuthorize2p.Psp_Amount = "15050"
	SplitAuthorize2p.Psp_Currency = "032"
	SplitAuthorize2p.Psp_Country = "ARG"
	SplitAuthorize2p.Psp_Product = "14"
	SplitAuthorize2p.Psp_CardNumber = "4507990000000010"
	SplitAuthorize2p.Psp_CardExpDate = "1903"
	SplitAuthorize2p.Psp_PosDateTime = "2016-12-01 12:00:00"

	SplitAuthorize2p.Psp_BillingDetails = Billing
	SplitAuthorize2p.Psp_CustomerAdditionalDetails = CustomerAdditionalDetails
	SplitAuthorize2p.Psp_MerchantAdditionalDetails = MerchantAdditionalDetails
	SplitAuthorize2p.Psp_OrderDetails = OrderDetails
	SplitAuthorize2p.Psp_ShippingDetails = ShippingDetails
	SplitAuthorize2p.Psp_Transactions = Transactions

	resp, err := service.SplitAuthorize_2p(SplitAuthorize2p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendPayOnLine_3p(service *nps.PaymentServicePlatformPortType) error {

	Person := nps.NewPersonStruct()
	Person.FirstName = "Silvina"
	Person.LastName = "Falconi"
	Person.PhoneNumber1 = "52520960"

	AmountAdditionalDetails := nps.NewAmountAdditionalDetailsRequestStruct()
	AmountAdditionalDetails.Tip = "10"
	AmountAdditionalDetails.Discount = "5"

	Billing := nps.NewBillingDetailsStruct()
	Billing.Invoice = "100001234"
	//Billing.InvoiceDate = "2017-06-29 12:00:00"
	Billing.InvoiceAmount = "990"
	Billing.InvoiceCurrency = "032"
	Billing.Person = Person

	SellerAddress := nps.NewAddressStruct()
	SellerAddress.City = "CABA"
	SellerAddress.Country = "ARG"
	SellerAddress.Street = "SellerStreet"
	SellerAddress.HouseNumber = "1234"

	SellerDetails := nps.NewSellerDetailsStruct()
	SellerDetails.Name = "Seller Name"
	SellerDetails.Address = SellerAddress

	MerchantAdditionalDetails := nps.NewMerchantAdditionalDetailsStruct()
	MerchantAdditionalDetails.ShoppingCartInfo = "ShoppingCartInfo"
	MerchantAdditionalDetails.SellerDetails = SellerDetails

	CustomerAdditionalDetails := nps.NewCustomerAdditionalDetailsStruct()
	CustomerAdditionalDetails.EmailAddress = "mailAddr@mail.com.ar"

	order1 := nps.NewOrderItemStruct()
	order1.Description = "producto 1"
	order1.UnitPrice = "10"

	order2 := nps.NewOrderItemStruct()
	order2.Description = "producto 2"
	order2.UnitPrice = "20"

	OrderDetails := nps.NewOrderDetailsStruct()
	OrderDetails.OrderItems = nps.NewArrayOf_OrderItemStruct()
	OrderDetails.OrderItems.Items = make([]*nps.OrderItemStruct, 0)
	OrderDetails.OrderItems.Items = append(OrderDetails.OrderItems.Items, order1)
	OrderDetails.OrderItems.Items = append(OrderDetails.OrderItems.Items, order2)

	PayOnline3p := nps.NewRequerimientoStruct_PayOnLine_3p()

	PayOnline3p.Psp_Version = "2.2"
	PayOnline3p.Psp_MerchantId = "psp_test"
	PayOnline3p.Psp_TxSource = "WEB"
	PayOnline3p.Psp_MerchTxRef = "ORDER56666-3"
	PayOnline3p.Psp_MerchOrderId = "ORDER56666"
	PayOnline3p.Psp_Amount = "1000"
	PayOnline3p.Psp_NumPayments = "1"
	PayOnline3p.Psp_Currency = "032"
	PayOnline3p.Psp_Country = "ARG"
	PayOnline3p.Psp_Product = "14"
	PayOnline3p.Psp_CustomerMail = "yourmail@gmail"
	PayOnline3p.Psp_SoftDescriptor = "Sol Tropical E"
	PayOnline3p.Psp_PosDateTime = "2016-12-01 12:00:00"
	PayOnline3p.Psp_ReturnURL = "'http://psp-client.localhost/simple_query_tx.php"
	PayOnline3p.Psp_FrmLanguage = "en_US"

	PayOnline3p.Psp_OrderDetails = OrderDetails
	PayOnline3p.Psp_CustomerAdditionalDetails = CustomerAdditionalDetails
	PayOnline3p.Psp_AmountAdditionalDetails = AmountAdditionalDetails
	PayOnline3p.Psp_BillingDetails = Billing
	PayOnline3p.Psp_MerchantAdditionalDetails = MerchantAdditionalDetails

	resp, err := service.PayOnLine_3p(PayOnline3p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendAuthorize_3p(service *nps.PaymentServicePlatformPortType) error {

	CustomerAdditionalDetails := nps.NewCustomerAdditionalDetailsStruct()
	CustomerAdditionalDetails.EmailAddress = "CustomerEmail@email.com.ar"

	Person := nps.NewPersonStruct()
	Person.FirstName = "Silvina"
	Person.LastName = "Falconi"
	Person.PhoneNumber1 = "52520960"

	Billing := nps.NewBillingDetailsStruct()
	Billing.Invoice = "100001234"
	//Billing.InvoiceDate = "20170601"
	Billing.InvoiceAmount = "990"
	Billing.InvoiceCurrency = "032"
	Billing.Person = Person

	SellerAddress := nps.NewAddressStruct()
	SellerAddress.City = "CABA"
	SellerAddress.Country = "ARG"
	SellerAddress.Street = "SellerStreet"
	SellerAddress.HouseNumber = "1234"

	SellerDetails := nps.NewSellerDetailsStruct()
	SellerDetails.Name = "Seller Name"
	SellerDetails.Address = SellerAddress

	MerchantAdditionalDetails := nps.NewMerchantAdditionalDetailsStruct()
	MerchantAdditionalDetails.ShoppingCartInfo = "ShoppingCartInfo"
	MerchantAdditionalDetails.SellerDetails = SellerDetails

	order1 := nps.NewOrderItemStruct()
	order1.Description = "producto 1"
	order1.UnitPrice = "10"

	order2 := nps.NewOrderItemStruct()
	order2.Description = "producto 2"
	order2.UnitPrice = "20"

	OrderDetails := nps.NewOrderDetailsStruct()
	OrderDetails.OrderItems = nps.NewArrayOf_OrderItemStruct()
	OrderDetails.OrderItems.Items = make([]*nps.OrderItemStruct, 0)
	OrderDetails.OrderItems.Items = append(OrderDetails.OrderItems.Items, order1)
	OrderDetails.OrderItems.Items = append(OrderDetails.OrderItems.Items, order2)

	ShippingDetails := nps.NewShippingDetailsStruct()

	ShippingAddress := nps.NewAddressStruct()
	ShippingAddress.Street = "shipping street"
	ShippingAddress.HouseNumber = "1234"
	ShippingAddress.City = "CABA"
	ShippingAddress.Country = "ARG"
	ShippingAddress.AdditionalInfo = "AdditionalInfo of shipping"

	PrimaryRecipient := nps.NewPersonStruct()
	PrimaryRecipient.FirstName = "Pepe"
	PrimaryRecipient.LastName = "Juan"

	ShippingDetails.Method = "10"
	ShippingDetails.Address = ShippingAddress
	ShippingDetails.PrimaryRecipient = PrimaryRecipient

	Authorize3p := nps.NewRequerimientoStruct_Authorize_3p()

	Authorize3p.Psp_Version = "2.2"
	Authorize3p.Psp_MerchantId = "psp_test"
	Authorize3p.Psp_TxSource = "WEB"
	Authorize3p.Psp_MerchTxRef = "20170609150900"
	Authorize3p.Psp_MerchOrderId = "20170609150900-1"
	Authorize3p.Psp_Amount = "15050"
	Authorize3p.Psp_NumPayments = "1"
	Authorize3p.Psp_Currency = "032"
	Authorize3p.Psp_Country = "ARG"
	Authorize3p.Psp_Product = "14"
	Authorize3p.Psp_PosDateTime = "2016-12-01 12:00:00"

	Authorize3p.Psp_ReturnURL = "http://psp-client.localhost/simple_query_tx.php"
	Authorize3p.Psp_FrmLanguage = "en_US"

	Authorize3p.Psp_OrderDetails = OrderDetails
	Authorize3p.Psp_ShippingDetails = ShippingDetails

	resp, err := service.Authorize_3p(Authorize3p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendSplitAuthorize_3p(service *nps.PaymentServicePlatformPortType) error {

	order1 := nps.NewOrderItemStruct()
	order1.Description = "producto 1"
	order1.UnitPrice = "10"

	order2 := nps.NewOrderItemStruct()
	order2.Description = "producto 2"
	order2.UnitPrice = "20"

	OrderDetails := nps.NewOrderDetailsStruct()
	OrderDetails.OrderItems = nps.NewArrayOf_OrderItemStruct()
	OrderDetails.OrderItems.Items = make([]*nps.OrderItemStruct, 0)
	OrderDetails.OrderItems.Items = append(OrderDetails.OrderItems.Items, order1)
	OrderDetails.OrderItems.Items = append(OrderDetails.OrderItems.Items, order2)

	ShippingDetails := nps.NewShippingDetailsStruct()

	ShippingAddress := nps.NewAddressStruct()
	ShippingAddress.Street = "shipping street"
	ShippingAddress.HouseNumber = "1234"
	ShippingAddress.City = "CABA"
	ShippingAddress.Country = "ARG"
	ShippingAddress.AdditionalInfo = "AdditionalInfo of shipping"

	PrimaryRecipient := nps.NewPersonStruct()
	PrimaryRecipient.FirstName = "Pepe"
	PrimaryRecipient.LastName = "Juan"

	ShippingDetails.Method = "10"
	ShippingDetails.Address = ShippingAddress
	ShippingDetails.PrimaryRecipient = PrimaryRecipient

	trn := nps.NewRequerimientoStruct_SplitAuthorize_3p_Transactions()
	trn.Psp_MerchantId = "psp_test"
	trn.Psp_MerchTxRef = "ORDER66666-5"
	trn.Psp_Product = "14"
	trn.Psp_Amount = "10000"
	trn.Psp_NumPayments = "1"

	trn2 := nps.NewRequerimientoStruct_SplitAuthorize_3p_Transactions()
	trn2.Psp_MerchantId = "psp_test"
	trn2.Psp_MerchTxRef = "ORDER66666-6"
	trn2.Psp_Product = "14"
	trn2.Psp_Amount = "5050"
	trn2.Psp_NumPayments = "1"

	Transactions := nps.NewArrayOf_RequerimientoStruct_SplitAuthorize_3p_Transactions()
	Transactions.Items = make([]*nps.RequerimientoStruct_SplitAuthorize_3p_Transactions, 0)

	Transactions.Items = append(Transactions.Items, trn)
	Transactions.Items = append(Transactions.Items, trn2)

	SplitAuthorize3p := nps.NewRequerimientoStruct_SplitAuthorize_3p()

	SplitAuthorize3p.Psp_Version = "2.2"
	SplitAuthorize3p.Psp_MerchantId = "psp_test"
	SplitAuthorize3p.Psp_TxSource = "WEB"
	SplitAuthorize3p.Psp_MerchOrderId = "20170609150900-1"
	SplitAuthorize3p.Psp_Amount = "15050"
	SplitAuthorize3p.Psp_Currency = "032"
	SplitAuthorize3p.Psp_Country = "ARG"
	SplitAuthorize3p.Psp_Product = "14"
	SplitAuthorize3p.Psp_PosDateTime = "2016-12-01 12:00:00"

	SplitAuthorize3p.Psp_ReturnURL = "http://psp-client.localhost/simple_query_tx.php"
	SplitAuthorize3p.Psp_FrmLanguage = "en_US"

	SplitAuthorize3p.Psp_OrderDetails = OrderDetails
	SplitAuthorize3p.Psp_ShippingDetails = ShippingDetails
	SplitAuthorize3p.Psp_Transactions = Transactions

	resp, err := service.SplitAuthorize_3p(SplitAuthorize3p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendBankPayment_3p(service *nps.PaymentServicePlatformPortType) error {

	Taxes := nps.NewArrayOf_TaxesRequestStruct()
	Taxes.Items = make([]*nps.TaxesRequestStruct, 0)

	tax1 := nps.NewTaxesRequestStruct()
	tax1.TypeId = "500"
	tax1.Amount = "50"
	tax1.Rate = "10"

	Taxes.Items = append(Taxes.Items, tax1)

	AmountAdditionalDetails := nps.NewAmountAdditionalDetailsRequestStruct()
	AmountAdditionalDetails.Tip = "10"
	AmountAdditionalDetails.Discount = "5"
	AmountAdditionalDetails.Taxes = Taxes

	Person := nps.NewPersonStruct()
	Person.FirstName = "Silvina"
	Person.LastName = "Falconi"
	Person.PhoneNumber1 = "52520960"
	Person.IDNumber = "11111111"
	Person.IDType = "100"

	Billing := nps.NewBillingDetailsStruct()
	Billing.Invoice = "100001234"
	//Billing.InvoiceDate = "20170601"
	Billing.InvoiceAmount = "990"
	Billing.InvoiceCurrency = "032"
	Billing.Person = Person

	CustomerAdditionalDetails := nps.NewCustomerAdditionalDetailsStruct()
	CustomerAdditionalDetails.EmailAddress = "CustomerEmail@email.com.ar"

	BankPayment_3p := nps.NewRequerimientoStruct_BankPayment_3p()
	BankPayment_3p.Psp_Version = "2.2"
	BankPayment_3p.Psp_Product = "320"
	BankPayment_3p.Psp_MerchantId = "psp_test"
	BankPayment_3p.Psp_TxSource = "TxSource"
	BankPayment_3p.Psp_MerchTxRef = "11111"
	BankPayment_3p.Psp_MerchOrderId = "ORD-10001"
	BankPayment_3p.Psp_ScreenDescription = "SCR DESC"
	BankPayment_3p.Psp_TicketDescription = "ticket desc"
	BankPayment_3p.Psp_ReturnURL = "http://localhost/"
	BankPayment_3p.Psp_FrmLanguage = "es_AR"

	BankPayment_3p.Psp_Amount1 = "10000"
	BankPayment_3p.Psp_ExpDate1 = "2017-09-01"
	BankPayment_3p.Psp_Amount2 = "20000"
	BankPayment_3p.Psp_ExpDate2 = "2017-09-10"
	BankPayment_3p.Psp_Country = "ARG"
	BankPayment_3p.Psp_Currency = "032"
	BankPayment_3p.Psp_PosDateTime = "2016-12-01 12:00:00"
	BankPayment_3p.Psp_TxSource = "WEB"

	BankPayment_3p.Psp_AmountAdditionalDetails = AmountAdditionalDetails
	BankPayment_3p.Psp_BillingDetails = Billing
	BankPayment_3p.Psp_CustomerAdditionalDetails = CustomerAdditionalDetails

	resp, err := service.BankPayment_3p(BankPayment_3p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendCashPayment_3p(service *nps.PaymentServicePlatformPortType) error {

	Taxes := nps.NewArrayOf_TaxesRequestStruct()
	Taxes.Items = make([]*nps.TaxesRequestStruct, 0)

	tax1 := nps.NewTaxesRequestStruct()
	tax1.TypeId = "500"
	tax1.Amount = "50"
	tax1.Rate = "10"

	Taxes.Items = append(Taxes.Items, tax1)

	AmountAdditionalDetails := nps.NewAmountAdditionalDetailsRequestStruct()
	AmountAdditionalDetails.Tip = "10"
	AmountAdditionalDetails.Discount = "5"
	AmountAdditionalDetails.Taxes = Taxes

	Person := nps.NewPersonStruct()
	Person.FirstName = "Silvina"
	Person.LastName = "Falconi"
	Person.PhoneNumber1 = "52520960"
	Person.IDNumber = "11111111"
	Person.IDType = "100"

	Billing := nps.NewBillingDetailsStruct()
	Billing.Invoice = "100001234"
	//Billing.InvoiceDate = "20170601"
	Billing.InvoiceAmount = "990"
	Billing.InvoiceCurrency = "032"
	Billing.Person = Person

	CustomerAdditionalDetails := nps.NewCustomerAdditionalDetailsStruct()
	CustomerAdditionalDetails.EmailAddress = "CustomerEmail@email.com.ar"

	order1 := nps.NewOrderItemStruct()
	order1.Description = "producto 1"
	order1.UnitPrice = "10"

	order2 := nps.NewOrderItemStruct()
	order2.Description = "producto 2"
	order2.UnitPrice = "20"

	OrderDetails := nps.NewOrderDetailsStruct()
	OrderDetails.OrderItems = nps.NewArrayOf_OrderItemStruct()
	OrderDetails.OrderItems.Items = make([]*nps.OrderItemStruct, 0)
	OrderDetails.OrderItems.Items = append(OrderDetails.OrderItems.Items, order1)
	OrderDetails.OrderItems.Items = append(OrderDetails.OrderItems.Items, order2)

	CashPayment_3p := nps.NewRequerimientoStruct_CashPayment_3p()
	CashPayment_3p.Psp_Version = "2.2"
	CashPayment_3p.Psp_Product = "320"
	CashPayment_3p.Psp_MerchantId = "psp_test"
	CashPayment_3p.Psp_TxSource = "TxSource"
	CashPayment_3p.Psp_MerchTxRef = "11111"
	CashPayment_3p.Psp_MerchOrderId = "ORD-10001"
	CashPayment_3p.Psp_ReturnURL = "http://localhost/"
	CashPayment_3p.Psp_FrmLanguage = "es_AR"
	CashPayment_3p.Psp_Amount = "11000"
	CashPayment_3p.Psp_Currency = "032"
	CashPayment_3p.Psp_Country = "ARG"
	CashPayment_3p.Psp_PosDateTime = "2016-12-01 12:00:00"
	CashPayment_3p.Psp_TxSource = "WEB"

	CashPayment_3p.Psp_AmountAdditionalDetails = AmountAdditionalDetails
	CashPayment_3p.Psp_BillingDetails = Billing
	CashPayment_3p.Psp_OrderDetails = OrderDetails

	resp, err := service.CashPayment_3p(CashPayment_3p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendCapture(service *nps.PaymentServicePlatformPortType) error {

	Taxes := nps.NewArrayOf_TaxesRequestStruct()
	Taxes.Items = make([]*nps.TaxesRequestStruct, 0)

	tax1 := nps.NewTaxesRequestStruct()
	tax1.TypeId = "500"
	tax1.Amount = "50"
	tax1.Rate = "10"

	Taxes.Items = append(Taxes.Items, tax1)

	AmountAdditionalDetails := nps.NewAmountAdditionalDetailsRequestStruct()
	AmountAdditionalDetails.Tip = "10"
	AmountAdditionalDetails.Discount = "5"
	AmountAdditionalDetails.Taxes = Taxes

	Capture := nps.NewRequerimientoStruct_Capture()
	Capture.Psp_Version = "2.2"
	Capture.Psp_MerchantId = "psp_test"
	Capture.Psp_TxSource = "TxSource"
	Capture.Psp_MerchTxRef = "20170630103800"
	Capture.Psp_PosDateTime = "2016-12-01 12:00:00"
	Capture.Psp_TxSource = "WEB"
	Capture.Psp_TransactionId_Orig = "1868712"
	Capture.Psp_AmountToCapture = "15000"

	Capture.Psp_AmountAdditionalDetails = AmountAdditionalDetails

	resp, err := service.Capture(Capture)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendRefund(service *nps.PaymentServicePlatformPortType) error {

	Refund := nps.NewRequerimientoStruct_Refund()
	Refund.Psp_Version = "2.2"
	Refund.Psp_MerchantId = "psp_test"
	Refund.Psp_TxSource = "TxSource"
	Refund.Psp_MerchTxRef = "20170630104000"
	Refund.Psp_PosDateTime = "2016-12-01 12:00:00"
	Refund.Psp_TxSource = "WEB"
	Refund.Psp_TransactionId_Orig = "1868717"
	Refund.Psp_AmountToRefund = "10000"

	resp, err := service.Refund(Refund)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendRetrievePaymentMethodToken(service *nps.PaymentServicePlatformPortType) error {

	RetrievePaymentMethodToken := nps.NewRequerimientoStruct_RetrievePaymentMethodToken()
	RetrievePaymentMethodToken.Psp_Version = "2.2"
	RetrievePaymentMethodToken.Psp_MerchantId = "psp_test"
	RetrievePaymentMethodToken.Psp_PaymentMethodToken = "FG2WE0mVhoOm4U2a1R8qsBe307mtiVy0"
	RetrievePaymentMethodToken.Psp_ClientSession = "iYgdiXyl56vszeEpCRGmS1JiNYg1xSnYXzQuiFWP4Q2nTwbPiWwZruUzXmqrXYR9"

	resp, err := service.RetrievePaymentMethodToken(RetrievePaymentMethodToken)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendCreatePaymentMethodToken(service *nps.PaymentServicePlatformPortType) error {

	cardInputDetails := nps.NewCardInputDetailsStruct()
	cardInputDetails.HolderName = "tester"
	cardInputDetails.ExpirationDate = "1812"
	cardInputDetails.Number = "4507990000000010"
	cardInputDetails.SecurityCode = "123"

	CreatePaymentMethodToken := nps.NewRequerimientoStruct_CreatePaymentMethodToken()

	CreatePaymentMethodToken.Psp_Version = "2.2"
	CreatePaymentMethodToken.Psp_MerchantId = "psp_test"
	CreatePaymentMethodToken.Psp_Product = "14"
	CreatePaymentMethodToken.Psp_CardInputDetails = cardInputDetails
	CreatePaymentMethodToken.Psp_ClientSession = "iYgdiXyl56vszeEpCRGmS1JiNYg1xSnYXzQuiFWP4Q2nTwbPiWwZruUzXmqrXYR9"

	resp, err := service.CreatePaymentMethodToken(CreatePaymentMethodToken)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendRecachePaymentMethodToken(service *nps.PaymentServicePlatformPortType) error {

	RecachePaymentMethodToken := nps.NewRequerimientoStruct_RecachePaymentMethodToken()
	RecachePaymentMethodToken.Psp_Version = "2.2"
	RecachePaymentMethodToken.Psp_MerchantId = "psp_test"
	RecachePaymentMethodToken.Psp_PaymentMethodId = "ZSUjUAEcb5uQ3qtNSS7N7gyDMPtODkjt"
	RecachePaymentMethodToken.Psp_ClientSession = "iYgdiXyl56vszeEpCRGmS1JiNYg1xSnYXzQuiFWP4Q2nTwbPiWwZruUzXmqrXYR9"
	RecachePaymentMethodToken.Psp_CardSecurityCode = "123"

	resp, err := service.RecachePaymentMethodToken(RecachePaymentMethodToken)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendCreatePaymentMethodFromPayment(service *nps.PaymentServicePlatformPortType) error {

	createPaymentMethodFromPayment := nps.NewRequerimientoStruct_CreatePaymentMethodFromPayment()

	createPaymentMethodFromPayment.Psp_Version = "2.2"
	createPaymentMethodFromPayment.Psp_MerchantId = "psp_test"
	createPaymentMethodFromPayment.Psp_PosDateTime = "2017-06-19 12:00:00"
	createPaymentMethodFromPayment.Psp_TransactionId = "1868712"

	resp, err := service.CreatePaymentMethodFromPayment(createPaymentMethodFromPayment)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendRetrievePaymentMethod(service *nps.PaymentServicePlatformPortType) error {

	RetrievePaymentMethod := nps.NewRequerimientoStruct_RetrievePaymentMethod()
	RetrievePaymentMethod.Psp_Version = "2.2"
	RetrievePaymentMethod.Psp_MerchantId = "psp_test"
	RetrievePaymentMethod.Psp_PaymentMethodId = "ZSUjUAEcb5uQ3qtNSS7N7gyDMPtODkjt"
	RetrievePaymentMethod.Psp_PosDateTime = "2016-12-01 12:00:00"

	resp, err := service.RetrievePaymentMethod(RetrievePaymentMethod)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendUpdatePaymentMethod(service *nps.PaymentServicePlatformPortType) error {

	UpdatePaymentMethod := nps.NewRequerimientoStruct_UpdatePaymentMethod()

	UpdatePaymentMethod.Psp_Version = "2.2"
	UpdatePaymentMethod.Psp_MerchantId = "psp_test"
	UpdatePaymentMethod.Psp_PosDateTime = "2017-06-19 12:00:00"

	paymentMethod := nps.NewPaymentMethodInputDetailsStruct()
	cardInputDetails := nps.NewCardInputDetailsStruct()
	cardInputDetails.HolderName = "tester"
	cardInputDetails.ExpirationDate = "1812"
	cardInputDetails.Number = "4507990000000010"
	cardInputDetails.SecurityCode = "123"

	paymentMethod.CardInputDetails = cardInputDetails

	UpdatePaymentMethod.Psp_PaymentMethodId = "ZSUjUAEcb5uQ3qtNSS7N7gyDMPtODkjt"

	resp, err := service.UpdatePaymentMethod(UpdatePaymentMethod)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendDeletePaymentMethod(service *nps.PaymentServicePlatformPortType) error {

	DeletePaymentMethod := nps.NewRequerimientoStruct_DeletePaymentMethod()

	DeletePaymentMethod.Psp_Version = "2.2"
	DeletePaymentMethod.Psp_MerchantId = "psp_test"
	DeletePaymentMethod.Psp_PosDateTime = "2017-06-19 12:00:00"
	DeletePaymentMethod.Psp_PaymentMethodId = "ZSUjUAEcb5uQ3qtNSS7N7gyDMPtODkjt"

	resp, err := service.DeletePaymentMethod(DeletePaymentMethod)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendCreateCustomer(service *nps.PaymentServicePlatformPortType) error {

	Person := nps.NewPersonStruct()
	Person.FirstName = "Silvina"
	Person.LastName = "Falconi"
	Person.PhoneNumber1 = "52520960"

	Address := nps.NewAddressStruct()
	Address.City = "CABA"
	Address.Country = "ARG"
	Address.Street = "SellerStreet"
	Address.HouseNumber = "1234"

	PmPerson := nps.NewPersonStruct()
	PmPerson.FirstName = "Silvina"
	PmPerson.LastName = "Falconi"
	PmPerson.PhoneNumber1 = "52520960"

	PmAddress := nps.NewAddressStruct()
	PmAddress.City = "CABA"
	PmAddress.Country = "ARG"
	PmAddress.Street = "SellerStreet"
	PmAddress.HouseNumber = "1234"

	pm := nps.NewPaymentMethodInputDetailsStruct()
	pm.Address = PmAddress
	pm.Person = PmPerson

	CreateCustomer := nps.NewRequerimientoStruct_CreateCustomer()

	CreateCustomer.Psp_Version = "2.2"
	CreateCustomer.Psp_MerchantId = "psp_test"
	CreateCustomer.Psp_PosDateTime = "2017-06-19 12:00:00"
	CreateCustomer.Psp_EmailAddress = "CustomerEmail@email.com.ar"

	CreateCustomer.Psp_Person = Person
	CreateCustomer.Psp_Address = Address
	CreateCustomer.Psp_PaymentMethod = pm

	resp, err := service.CreateCustomer(CreateCustomer)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendRetrieveCustomer(service *nps.PaymentServicePlatformPortType) error {

	RetrieveCustomer := nps.NewRequerimientoStruct_RetrieveCustomer()

	RetrieveCustomer.Psp_Version = "2.2"
	RetrieveCustomer.Psp_MerchantId = "psp_test"
	RetrieveCustomer.Psp_PosDateTime = "2017-06-19 12:00:00"
	RetrieveCustomer.Psp_CustomerId = "xbqrLYhuPP8PnM5uQDw1GunU4V7rfT9Y"

	resp, err := service.RetrieveCustomer(RetrieveCustomer)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendUpdateCustomer(service *nps.PaymentServicePlatformPortType) error {

	Person := nps.NewPersonStruct()
	Person.FirstName = "Silvina"
	Person.LastName = "Falconi"
	Person.PhoneNumber1 = "52520960"

	Address := nps.NewAddressStruct()
	Address.City = "CABA"
	Address.Country = "ARG"
	Address.Street = "SellerStreet"
	Address.HouseNumber = "1234"

	PmPerson := nps.NewPersonStruct()
	PmPerson.FirstName = "Silvina"
	PmPerson.LastName = "Falconi"
	PmPerson.PhoneNumber1 = "52520960"

	PmAddress := nps.NewAddressStruct()
	PmAddress.City = "CABA"
	PmAddress.Country = "ARG"
	PmAddress.Street = "SellerStreet"
	PmAddress.HouseNumber = "1234"

	pm := nps.NewPaymentMethodInputDetailsStruct()
	pm.Address = PmAddress
	pm.Person = PmPerson

	UpdateCustomer := nps.NewRequerimientoStruct_UpdateCustomer()

	UpdateCustomer.Psp_Version = "2.2"
	UpdateCustomer.Psp_MerchantId = "psp_test"
	UpdateCustomer.Psp_PosDateTime = "2017-06-19 12:00:00"
	UpdateCustomer.Psp_CustomerId = "xbqrLYhuPP8PnM5uQDw1GunU4V7rfT9Y"
	UpdateCustomer.Psp_EmailAddress = "CustomerEmail@email.com.ar"

	UpdateCustomer.Psp_Person = Person
	UpdateCustomer.Psp_Address = Address
	UpdateCustomer.Psp_PaymentMethod = pm

	resp, err := service.UpdateCustomer(UpdateCustomer)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendDeleteCustomer(service *nps.PaymentServicePlatformPortType) error {

	DeleteCustomer := nps.NewRequerimientoStruct_DeleteCustomer()

	DeleteCustomer.Psp_Version = "2.2"
	DeleteCustomer.Psp_MerchantId = "psp_test"
	DeleteCustomer.Psp_PosDateTime = "2017-06-19 12:00:00"
	DeleteCustomer.Psp_CustomerId = "xbqrLYhuPP8PnM5uQDw1GunU4V7rfT9Y"

	resp, err := service.DeleteCustomer(DeleteCustomer)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendQueryCardNumber(service *nps.PaymentServicePlatformPortType) error {

	QueryCardNumber := nps.NewRequerimientoStruct_QueryCardNumber()

	QueryCardNumber.Psp_Version = "2.2"
	QueryCardNumber.Psp_MerchantId = "psp_test"
	QueryCardNumber.Psp_PosDateTime = "2017-06-19 12:00:00"
	QueryCardNumber.Psp_QueryCriteria = "T"
	QueryCardNumber.Psp_QueryCriteriaId = "76577"

	resp, err := service.QueryCardNumber(QueryCardNumber)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendQueryCardDetails(service *nps.PaymentServicePlatformPortType) error {

	QueryCardDetails := nps.NewRequerimientoStruct_QueryCardDetails()

	QueryCardDetails.Psp_Version = "2.2"
	QueryCardDetails.Psp_MerchantId = "psp_test"
	QueryCardDetails.Psp_PosDateTime = "2017-06-19 12:00:00"
	QueryCardDetails.Psp_QueryCriteria = "T"
	QueryCardDetails.Psp_QueryCriteriaId = "76577"

	resp, err := service.QueryCardDetails(QueryCardDetails)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	return nil
}

func SendSimpleQueryTx(service *nps.PaymentServicePlatformPortType) error {

	SimpleQueryTx := nps.NewRequerimientoStruct_SimpleQueryTx()

	SimpleQueryTx.Psp_Version = "2.2"
	SimpleQueryTx.Psp_MerchantId = "psp_test"
	SimpleQueryTx.Psp_PosDateTime = "2017-06-19 12:00:00"
	SimpleQueryTx.Psp_QueryCriteria = "T"
	SimpleQueryTx.Psp_QueryCriteriaId = "1868712"

	resp, err := service.SimpleQueryTx(SimpleQueryTx)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)
	if resp.Psp_Transaction != nil && len(resp.Psp_Transaction.Psp_AuthorizationCode) > 0 {
		fmt.Printf("Psp_AuthorizationCode = [%s]\n", resp.Psp_Transaction.Psp_AuthorizationCode)
	}

	return nil
}

func SendQueryTxs(service *nps.PaymentServicePlatformPortType) error {

	QueryTxs := nps.NewRequerimientoStruct_QueryTxs()

	QueryTxs.Psp_Version = "2.2"
	QueryTxs.Psp_MerchantId = "psp_test"
	QueryTxs.Psp_PosDateTime = "2017-06-19 12:00:00"
	QueryTxs.Psp_QueryCriteria = "T"
	QueryTxs.Psp_QueryCriteriaId = "1868712"

	resp, err := service.QueryTxs(QueryTxs)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)

	return nil
}

func SendGetInstallmentsOptions(service *nps.PaymentServicePlatformPortType) error {

	GetInstallmentsOptions := nps.NewRequerimientoStruct_GetInstallmentsOptions()

	GetInstallmentsOptions.Psp_Version = "2.2"
	GetInstallmentsOptions.Psp_MerchantId = "psp_test"
	GetInstallmentsOptions.Psp_PosDateTime = "2017-06-19 12:00:00"

	GetInstallmentsOptions.Psp_Amount = "1000"
	GetInstallmentsOptions.Psp_Product = "14"
	GetInstallmentsOptions.Psp_Currency = "032"
	GetInstallmentsOptions.Psp_Country = "ARG"

	GetInstallmentsOptions.Psp_ClientSession = "iYgdiXyl56vszeEpCRGmS1JiNYg1xSnYXzQuiFWP4Q2nTwbPiWwZruUzXmqrXYR9"

	resp, err := service.GetInstallmentsOptions(GetInstallmentsOptions)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)

	return nil
}

func SendGetIINDetails(service *nps.PaymentServicePlatformPortType) error {

	GetIINDetails := nps.NewRequerimientoStruct_GetIINDetails()

	GetIINDetails.Psp_Version = "2.2"
	GetIINDetails.Psp_MerchantId = "psp_test"
	GetIINDetails.Psp_PosDateTime = "2017-06-19 12:00:00"
	GetIINDetails.Psp_IIN = "450799"

	GetIINDetails.Psp_ClientSession = "iYgdiXyl56vszeEpCRGmS1JiNYg1xSnYXzQuiFWP4Q2nTwbPiWwZruUzXmqrXYR9"

	resp, err := service.GetIINDetails(GetIINDetails)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	fmt.Printf("Response = [%s] [%s]\n", resp.Psp_ResponseCod, resp.Psp_ResponseMsg)
	fmt.Printf("Extended = [%s]\n", resp.Psp_ResponseExtended)

	return nil
}
func main() {

	/*******************************************************************
	  Configuring my appLog and use it by sdk:
	  var appLog = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
		f, err := os.OpenFile("appLog.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		defer f.Close()
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		appLog.SetOutput(f)
		appLog.Println("MAIN begin")
	  **********************************************************************/

	err := nps.Configure(map[string]interface{}{
		"environment": CONSTANTS.SANDBOX_ENV,
		"secret_key":  "IeShlZMDk8mp8VA6vy41mLnVggnj1yqHcJyNqIYaRINZnXdiTfhF0Ule9WNAUCR6",
		"debug":       true,
		"log_level":   CONSTANTS.DEBUG,
		//"conn_timeout": 10,
		//"timeout":      60,
		//"npsLog":      appLog      // use appLog at configuration
	})
	if err != nil {
		log.Fatalf("error configuring sdk: %v", err)
	}

	service := nps.NewPaymentServicePlatformPortType(true)

	//err = SendPayOnLine_2p(service)
	//err = SendAuthorize_2p(service)
	//err = SendBankPayment_2p(service)
	//err = SendSplitPayOnLine_2p(service)
	//err = SendSplitAuthorize_2p(service)

	//err = SendPayOnLine_3p(service)
	//err = SendSplitPayOnLine_3p(service)
	//err = SendAuthorize_3p(service)
	//err = SendSplitAuthorize_3p(service)
	//err = SendBankPayment_3p(service)
	//err = SendCashPayment_3p(service)

	//err = SendCapture(service)
	//err = SendRefund(service)

	//err = SendCreatePaymentMethod(service)
	//err = SendCreatePaymentMethodToken(service)
	//err = SendRetrievePaymentMethodToken(service)
	//err = SendRecachePaymentMethodToken(service)
	//err = SendCreateClientSession(service)

	//err = SendCreatePaymentMethodFromPayment(service)
	//err = SendRetrievePaymentMethod(service)
	//err = SendUpdatePaymentMethod(service)
	//err = SendDeletePaymentMethod(service)

	//err = SendCreateCustomer(service)
	//err = SendRetrieveCustomer(service)
	err = SendUpdateCustomer(service)
	//err = SendDeleteCustomer(service)

	//err = SendQueryCardNumber(service)
	//err = SendQueryCardDetails(service)

	//err = SendSimpleQueryTx(service)
	//err = SendQueryTxs(service)

	//err = SendGetInstallmentsOptions(service)
	//err = SendGetIINDetails(service)

	if err != nil {
		fmt.Printf("\n")
		fmt.Printf("ERROR [%s]\n", err)
	}

}
