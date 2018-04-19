package main

import (
	"fmt"
	"log"
	"github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk"
	CONSTANTS "github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk/constants"
	"reflect"
	"time"
)

func SendPayOnLine_2p(service *npsSdk.PaymentServicePlatformPortType) error {
	Person := npsSdk.NewPersonStruct()

        fName := reflect.ValueOf(Person).Elem().FieldByName("FirstName")
        lName := reflect.ValueOf(Person).Elem().FieldByName("LastName")
        phNumber1 := reflect.ValueOf(Person).Elem().FieldByName("PhoneNumber1")

        if fName.IsValid() { fName.SetString("Silvina") }
        if lName.IsValid() { lName.SetString("Falconi") }
        if phNumber1.IsValid() { phNumber1.SetString("52520960") }

	AmountAdditionalDetails := npsSdk.NewAmountAdditionalDetailsRequestStruct()
        tip := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Tip")
        discount := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Discount")

        if tip.IsValid() { tip.SetString("10") }
        if discount.IsValid() { discount.SetString("5") }

	Billing := npsSdk.NewBillingDetailsStruct()
        Invoice := reflect.ValueOf(Billing).Elem().FieldByName("Invoice")
        InvoiceAmount := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceAmount")
        InvoiceCurrency := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceCurrency")
        BillingPerson := reflect.ValueOf(Billing).Elem().FieldByName("Person")

        if Invoice.IsValid() { Invoice.SetString("100001234") }
        if InvoiceAmount.IsValid() { InvoiceAmount.SetString("990") }
        if InvoiceCurrency.IsValid() { InvoiceCurrency.SetString("032") }
        if BillingPerson.IsValid() { BillingPerson.Set(reflect.ValueOf(Person)) }

	SellerAddress := npsSdk.NewAddressStruct()
        City := reflect.ValueOf(SellerAddress).Elem().FieldByName("City")
        Country := reflect.ValueOf(SellerAddress).Elem().FieldByName("Country")
        Street := reflect.ValueOf(SellerAddress).Elem().FieldByName("Street")
        HouseNumber := reflect.ValueOf(SellerAddress).Elem().FieldByName("HouseNumber")
        
        if City.IsValid() { City.SetString("CABA") }
        if Country.IsValid() { Country.SetString("ARG") }
        if Street.IsValid() { Street.SetString("SellerStreet") }
        if HouseNumber.IsValid() { HouseNumber.SetString("1234") }

	SellerDetails := npsSdk.NewSellerDetailsStruct()
        SellerDetailsName := reflect.ValueOf(SellerDetails).Elem().FieldByName("Name")
        SellerDetailsAddress := reflect.ValueOf(SellerDetails).Elem().FieldByName("Address")

        if SellerDetailsName.IsValid() { SellerDetailsName.SetString("Seller Name") }
        if SellerDetailsAddress.IsValid() { SellerDetailsAddress.Set(reflect.ValueOf(SellerAddress)) }

	MerchantAdditionalDetails := npsSdk.NewMerchantAdditionalDetailsStruct()
        ShoppingCartInfo := reflect.ValueOf(MerchantAdditionalDetails).Elem().FieldByName("ShoppingCartInfo")
        MerchantAdditionalDetailsSellerDetails := reflect.ValueOf(MerchantAdditionalDetails).Elem().FieldByName("SellerDetails")
        
        if ShoppingCartInfo.IsValid() { ShoppingCartInfo.SetString("ShoppingCartInfo") }
        if MerchantAdditionalDetailsSellerDetails.IsValid() { MerchantAdditionalDetailsSellerDetails.Set(reflect.ValueOf(SellerDetails)) }

	CustomerAdditionalDetails := npsSdk.NewCustomerAdditionalDetailsStruct()
	EmailAddress := reflect.ValueOf(CustomerAdditionalDetails).Elem().FieldByName("EmailAddress")

        if EmailAddress.IsValid() { EmailAddress.SetString("mailAddr@mail.com.ar") }

      OrderItems := npsSdk.NewArrayOf_OrderItemStruct()

      Items := reflect.ValueOf(OrderItems).Elem().FieldByName("Items")
	
       if Items.Kind() == reflect.Slice {
         st := Items.Type()
         sliceType := st.Elem()
         if sliceType.Kind() == reflect.Ptr {
            sliceType = sliceType.Elem()
         }

         order1 := reflect.New(sliceType)
         order1.Elem().FieldByName("Description").SetString("producto 1") 
         order1.Elem().FieldByName("UnitPrice").SetString("10") 

         Items.Set(reflect.Append(Items, order1))

         order2 := reflect.New(sliceType)
         order2.Elem().FieldByName("Description").SetString("producto 2") 
         order2.Elem().FieldByName("UnitPrice").SetString("20") 

         Items.Set(reflect.Append(Items, order2))
         order3 := reflect.New(sliceType)
         order3.Elem().FieldByName("Description").SetString("producto 3") 
         order3.Elem().FieldByName("UnitPrice").SetString("30") 

         Items.Set(reflect.Append(Items, order3))

      }

      OrderDetails := npsSdk.NewOrderDetailsStruct()    
      OrderDetailsItems := reflect.ValueOf(OrderDetails).Elem().FieldByName("OrderItems")

      if OrderDetailsItems.IsValid() { OrderDetailsItems.Set(reflect.ValueOf(OrderItems)) }

	payOnline2p := npsSdk.NewRequerimientoStruct_PayOnLine_2p()
	Psp_Version := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_Version")
	Psp_MerchantId := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_MerchantId")
	Psp_TxSource := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_TxSource")
	Psp_MerchTxRef := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_MerchTxRef")
	Psp_MerchOrderId := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_MerchOrderId")
	Psp_Amount := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_Amount")
	Psp_NumPayments := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_NumPayments")
	Psp_Currency := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_Currency")
	Psp_Country := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_Country")
	Psp_Product := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_Product")
	Psp_CustomerMail := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_CustomerMail")
	Psp_CardNumber := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_CardNumber")
	Psp_CardExpDate := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_CardExpDate")
	Psp_CardSecurityCode := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_CardSecurityCode")
	Psp_SoftDescriptor := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_SoftDescriptor")
	Psp_PosDateTime := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_PosDateTime")

        Psp_UserId := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_UserId")
        if Psp_UserId.IsValid() { Psp_UserId.SetString("SFALCONI")} 

	if Psp_Version.IsValid() { Psp_Version.SetString("2.2")}
	if Psp_MerchantId.IsValid() { Psp_MerchantId.SetString("psp_test")}
	if Psp_TxSource.IsValid() { Psp_TxSource.SetString("WEB")}
        t := time.Now()
 	if Psp_MerchTxRef.IsValid() { Psp_MerchTxRef.SetString(t.Format("20060102150405")) }
	//if Psp_MerchTxRef.IsValid() { Psp_MerchTxRef.SetString("ORDER56666-3")}

	if Psp_MerchOrderId.IsValid() { Psp_MerchOrderId.SetString("ORDER56666")}
	if Psp_Amount.IsValid() { Psp_Amount.SetString("1000")}
	if Psp_NumPayments.IsValid() { Psp_NumPayments.SetString("1")}
	if Psp_Currency.IsValid() { Psp_Currency.SetString("032")}
	if Psp_Country.IsValid() { Psp_Country.SetString("ARG")}
	if Psp_Product.IsValid() { Psp_Product.SetString("14")}
	if Psp_CustomerMail.IsValid() { Psp_CustomerMail.SetString("yourmail@gmail")}
	if Psp_CardNumber.IsValid() { Psp_CardNumber.SetString("4507990000000010")}
	if Psp_CardExpDate.IsValid() { Psp_CardExpDate.SetString("1903")}
	if Psp_CardSecurityCode.IsValid() { Psp_CardSecurityCode.SetString("306")}
	if Psp_SoftDescriptor.IsValid() { Psp_SoftDescriptor.SetString("Sol Tropical E")}
	if Psp_PosDateTime.IsValid() { Psp_PosDateTime.SetString("2016-12-01 12:00:00")}

	Psp_OrderDetails := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_OrderDetails")
	Psp_CustomerAdditionalDetails := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_CustomerAdditionalDetails")
	Psp_AmountAdditionalDetails := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_AmountAdditionalDetails")
	Psp_BillingDetails := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_BillingDetails")
	Psp_MerchantAdditionalDetails := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_MerchantAdditionalDetails")

	if Psp_OrderDetails.IsValid() {  Psp_OrderDetails.Set(reflect.ValueOf(OrderDetails)) }
	if Psp_CustomerAdditionalDetails.IsValid() {  Psp_CustomerAdditionalDetails.Set(reflect.ValueOf(CustomerAdditionalDetails)) }
	if Psp_AmountAdditionalDetails.IsValid() {  Psp_AmountAdditionalDetails.Set(reflect.ValueOf(AmountAdditionalDetails)) }
	if Psp_BillingDetails.IsValid() {  Psp_BillingDetails.Set(reflect.ValueOf(Billing)) }
	if Psp_MerchantAdditionalDetails.IsValid() {  Psp_MerchantAdditionalDetails.Set(reflect.ValueOf(MerchantAdditionalDetails)) }

	resp, err := service.PayOnLine_2p(payOnline2p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendSplitPayOnLine_3p(service *npsSdk.PaymentServicePlatformPortType) error {

	splitPayOnline3p := npsSdk.NewRequerimientoStruct_SplitPayOnLine_3p()

	Psp_Version := reflect.ValueOf(splitPayOnline3p).Elem().FieldByName("Psp_Version")
        if Psp_Version.IsValid() { Psp_Version.SetString("2.2") }

	Psp_MerchantId := reflect.ValueOf(splitPayOnline3p).Elem().FieldByName("Psp_MerchantId")
	if Psp_MerchantId.IsValid() { Psp_MerchantId.SetString("psp_test") }

	Psp_TxSource :=  reflect.ValueOf(splitPayOnline3p).Elem().FieldByName("Psp_TxSource")
        if Psp_TxSource.IsValid() { Psp_TxSource.SetString("WEB") }

	Psp_MerchOrderId := reflect.ValueOf(splitPayOnline3p).Elem().FieldByName("Psp_MerchOrderId")
        if Psp_MerchOrderId.IsValid() { Psp_MerchOrderId.SetString("20170609134500") }

	Psp_ReturnURL := reflect.ValueOf(splitPayOnline3p).Elem().FieldByName("Psp_ReturnURL")
        if Psp_ReturnURL.IsValid() { Psp_ReturnURL.SetString("http://localhost/") }

	Psp_FrmLanguage := reflect.ValueOf(splitPayOnline3p).Elem().FieldByName("Psp_FrmLanguage")
        if Psp_FrmLanguage.IsValid() { Psp_FrmLanguage.SetString("es_AR") }

	Psp_Amount := reflect.ValueOf(splitPayOnline3p).Elem().FieldByName("Psp_Amount")
 	if Psp_Amount.IsValid() { Psp_Amount.SetString("15050") }

	Psp_Currency := reflect.ValueOf(splitPayOnline3p).Elem().FieldByName("Psp_Currency")
	if Psp_Currency.IsValid() { Psp_Currency.SetString("032") }

	Psp_Country := reflect.ValueOf(splitPayOnline3p).Elem().FieldByName("Psp_Country")
	if Psp_Country.IsValid() { Psp_Country.SetString("ARG") }

	Psp_Product := reflect.ValueOf(splitPayOnline3p).Elem().FieldByName("Psp_Product")
	if Psp_Product.IsValid() { Psp_Product.SetString("14") }

	Psp_PosDateTime := reflect.ValueOf(splitPayOnline3p).Elem().FieldByName("Psp_PosDateTime")
	if Psp_PosDateTime.IsValid() { Psp_PosDateTime.SetString("2016-12-01 12:00:00") }

	Transactions := reflect.ValueOf(splitPayOnline3p).Elem().FieldByName("Psp_Transactions")

	TransactionsArr := npsSdk.NewArrayOf_RequerimientoStruct_SplitPayOnLine_3p_Transactions()
	Items := reflect.ValueOf(TransactionsArr).Elem().FieldByName("Items")

       if Items.Kind() == reflect.Slice {
         st := Items.Type()
         sliceType := st.Elem()
         if sliceType.Kind() == reflect.Ptr {
            sliceType = sliceType.Elem()
         }

         trn := reflect.New(sliceType)
         trn.Elem().FieldByName("Psp_MerchantId").SetString("psp_test") 
         trn.Elem().FieldByName("Psp_MerchTxRef").SetString("ORDER66666-5") 
         trn.Elem().FieldByName("Psp_Product").SetString("14") 
         trn.Elem().FieldByName("Psp_Amount").SetString("10000") 
         trn.Elem().FieldByName("Psp_NumPayments").SetString("1") 

         Items.Set(reflect.Append(Items, trn))

         trn2 := reflect.New(sliceType)
         trn2.Elem().FieldByName("Psp_MerchantId").SetString("psp_test") 
         trn2.Elem().FieldByName("Psp_MerchTxRef").SetString("ORDER66666-6") 
         trn2.Elem().FieldByName("Psp_Product").SetString("14") 
         trn2.Elem().FieldByName("Psp_Amount").SetString("5050") 
         trn2.Elem().FieldByName("Psp_NumPayments").SetString("1") 

         Items.Set(reflect.Append(Items, trn2))

      }

      if Transactions.IsValid() { Transactions.Set(reflect.ValueOf(TransactionsArr)) }


	resp, err := service.SplitPayOnLine_3p(splitPayOnline3p)
	if err != nil {
		return err
	}

        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())

        ptrTransactions := reflect.ValueOf(resp).Elem().FieldByName("Psp_Transactions")
        respTransactions := ptrTransactions.Elem()
        if respTransactions.IsValid() {

	  RespItems := respTransactions.FieldByName("Items") //reflect.ValueOf(respTransactions).Elem().FieldByName("Items")
          if RespItems.IsValid() {

            fmt.Printf("RespItems = [%d]\n", RespItems.Len())
            for i := 0; i < RespItems.Len(); i++ {
  
		idxtrn := RespItems.Index(i).Elem()

		fmt.Printf("========== item[%d]= Kind[%s]========= \n", i, idxtrn.Kind().String())

		fmt.Printf("Psp_MerchantId         [%s]\n", idxtrn.FieldByName("Psp_MerchantId").String())

		fmt.Printf("Psp_TransactionId      [%s]\n", idxtrn.FieldByName("Psp_TransactionId").String())
		fmt.Printf("Psp_MerchTxRef         [%s]\n", idxtrn.FieldByName("Psp_MerchTxRef").String())
		fmt.Printf("Psp_MerchAdditionalRef [%s]\n", idxtrn.FieldByName("Psp_MerchAdditionalRef").String())
		fmt.Printf("Psp_Product            [%s]\n", idxtrn.FieldByName("Psp_Product").String())
		fmt.Printf("Psp_PromotionCode      [%s]\n", idxtrn.FieldByName("Psp_PromotionCode").String())
		fmt.Printf("Psp_Amount             [%s]\n", idxtrn.FieldByName("Psp_Amount").String())
		fmt.Printf("Psp_NumPayments        [%s]\n", idxtrn.FieldByName("Psp_NumPayments").String())
		fmt.Printf("Psp_Plan               [%s]\n", idxtrn.FieldByName("Psp_Plan").String())
		fmt.Printf("Psp_FirstPaymentDeferral [%s]\n", idxtrn.FieldByName("Psp_FirstPaymentDeferral").String())
		fmt.Printf("Psp_SoftDescriptor     [%s]\n", idxtrn.FieldByName("Psp_SoftDescriptor").String())
		fmt.Printf("Psp_CreatedAt          [%s]\n", idxtrn.FieldByName("Psp_CreatedAt").String())

            }
          }
}
	return nil
}

func SendAuthorize_2p(service *npsSdk.PaymentServicePlatformPortType) error {

	authorize2p := npsSdk.NewRequerimientoStruct_Authorize_2p()

	Psp_Version :=  reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_Version")
	if  Psp_Version.IsValid() { Psp_Version.SetString("2.2") }

	Psp_Merchant := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_MerchantId") 
	if  Psp_Merchant.IsValid() { Psp_Merchant.SetString("psp_test") }

	Psp_TxSource  := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_TxSource")
	if  Psp_TxSource.IsValid() { Psp_TxSource.SetString("WEB") }

	Psp_MerchTxRef := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_MerchTxRef") 
	if  Psp_MerchTxRef.IsValid() { Psp_MerchTxRef.SetString("20170630103600") }

	Psp_MerchOrderId := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_MerchOrderId") 
	if  Psp_MerchOrderId.IsValid() { Psp_MerchOrderId.SetString("20170630103600-1") }

	Psp_Amount := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_Amount") 
	if  Psp_Amount.IsValid() { Psp_Amount.SetString("15050") }

	Psp_NumPayments := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_NumPayments") 
	if  Psp_NumPayments.IsValid() { Psp_NumPayments.SetString("1") }

	Psp_Currency := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_Currency") 
	if  Psp_Currency.IsValid() {  Psp_Currency.SetString("032") }

	Psp_Country := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_Country") 
	if  Psp_Country.IsValid() { Psp_Country.SetString("ARG") }

	Psp_Product := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_Product") 
	if  Psp_Product.IsValid() { Psp_Product.SetString("14") }

	Psp_CardNumber := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_CardNumber") 
	if  Psp_CardNumber.IsValid() { Psp_CardNumber.SetString("4507990000000010") }

	Psp_CardExpDate := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_CardExpDate") 
	if  Psp_CardExpDate.IsValid() { Psp_CardExpDate.SetString("1812") }

	Psp_PosDateTime := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_PosDateTime") 
	if  Psp_PosDateTime.IsValid() { Psp_PosDateTime.SetString("2016-12-01 12:00:00") }

	CustomerAdditionalDetails := npsSdk.NewCustomerAdditionalDetailsStruct()
	EmailAddress := reflect.ValueOf(CustomerAdditionalDetails).Elem().FieldByName("EmailAddress")
 	if EmailAddress.IsValid() { EmailAddress.SetString("CustomerEmail@email.com.ar") }

	Person := npsSdk.NewPersonStruct()
	FirstName := reflect.ValueOf(Person).Elem().FieldByName("FirstName")
 	if FirstName.IsValid() { FirstName.SetString("Silvina") }

	LastName := reflect.ValueOf(Person).Elem().FieldByName("LastName")
 	if LastName.IsValid() { LastName.SetString("Falconi") }

	PhoneNumber1 := reflect.ValueOf(Person).Elem().FieldByName("PhoneNumber1")
 	if PhoneNumber1.IsValid() { PhoneNumber1.SetString("52520960") }

	Billing := npsSdk.NewBillingDetailsStruct()
	Invoice := reflect.ValueOf(Billing).Elem().FieldByName("Invoice")
 	if Invoice.IsValid() { Invoice.SetString("100001234") }

	InvoiceAmount := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceAmount")
 	if InvoiceAmount.IsValid() { InvoiceAmount.SetString("990") }

	InvoiceCurrency := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceCurrency")
 	if InvoiceCurrency.IsValid() { InvoiceCurrency.SetString("032") }

        billingPerson := reflect.ValueOf(Billing).Elem().FieldByName("Person")
 	if billingPerson.IsValid() { billingPerson.Set(reflect.ValueOf(Person)) }

	Taxes := npsSdk.NewArrayOf_TaxesRequestStruct()

        Items := reflect.ValueOf(Taxes).Elem().FieldByName("Items")
	if Items.Kind() == reflect.Slice {
          st := Items.Type()
          sliceType := st.Elem()
          if sliceType.Kind() == reflect.Ptr {
            sliceType = sliceType.Elem()
          }

         tax1 := reflect.New(sliceType)
         tax1.Elem().FieldByName("TypeId").SetString("500")
         tax1.Elem().FieldByName("Amount").SetString("50")  
         tax1.Elem().FieldByName("Rate").SetString("10")  
         Items.Set(reflect.Append(Items, tax1))
	}

	AmountAdditionalDetails := npsSdk.NewAmountAdditionalDetailsRequestStruct()

	amountAddTip := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Tip")
	if amountAddTip.IsValid() { amountAddTip.SetString("10") }

	amountAddDiscount := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Discount")
	if amountAddDiscount.IsValid() { amountAddDiscount.SetString("5") }

	amountAddTaxes := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Taxes")
	if amountAddTaxes.IsValid() { amountAddTaxes.Set(reflect.ValueOf(Taxes)) }

	SellerAddress := npsSdk.NewAddressStruct()

	sellerCity := reflect.ValueOf(SellerAddress).Elem().FieldByName("City")
 	if sellerCity.IsValid() { sellerCity.SetString("CABA") }

	sellerCountry := reflect.ValueOf(SellerAddress).Elem().FieldByName("Country")
 	if sellerCountry.IsValid() { sellerCountry.SetString("ARG") }

	sellerStreet := reflect.ValueOf(SellerAddress).Elem().FieldByName("Street")
 	if sellerStreet.IsValid() { sellerStreet.SetString("SellerStreet") }

	sellerHouseNumber := reflect.ValueOf(SellerAddress).Elem().FieldByName("HouseNumber")
 	if sellerHouseNumber.IsValid() { sellerHouseNumber.SetString("1234") }

	SellerDetails := npsSdk.NewSellerDetailsStruct()
	sellerDetailsName := reflect.ValueOf(SellerDetails).Elem().FieldByName("Name")
	if sellerDetailsName.IsValid() { sellerDetailsName.SetString("Seller Name") }

	sellerDetailsAddress := reflect.ValueOf(SellerDetails).Elem().FieldByName("Address")

	if sellerDetailsAddress.IsValid() { sellerDetailsAddress.Set(reflect.ValueOf(SellerAddress)) }

	MerchantAdditionalDetails := npsSdk.NewMerchantAdditionalDetailsStruct()

	shoppingCartInfo := reflect.ValueOf(MerchantAdditionalDetails).Elem().FieldByName("ShoppingCartInfo")
	if shoppingCartInfo.IsValid() { shoppingCartInfo.SetString("ShoppingCartInfo") }

	merchantAdditionalDetailsSellerDetails := reflect.ValueOf(MerchantAdditionalDetails).Elem().FieldByName("SellerDetails")
	if merchantAdditionalDetailsSellerDetails.IsValid() { merchantAdditionalDetailsSellerDetails.Set(reflect.ValueOf(SellerDetails)) }

        OrderItems := npsSdk.NewArrayOf_OrderItemStruct()
        orderItems := reflect.ValueOf(OrderItems).Elem().FieldByName("Items")
	if orderItems.Kind() == reflect.Slice {
          st := orderItems.Type()
          sliceType := st.Elem()
          if sliceType.Kind() == reflect.Ptr {
            sliceType = sliceType.Elem()
          }

          order1 := reflect.New(sliceType)
          order1.Elem().FieldByName("Description").SetString("producto 1")
          order1.Elem().FieldByName("UnitPrice").SetString("10")

          orderItems.Set(reflect.Append(orderItems, order1))

          order2 := reflect.New(sliceType)
          order2.Elem().FieldByName("Description").SetString("producto 2")
          order2.Elem().FieldByName("UnitPrice").SetString("20")

          orderItems.Set(reflect.Append(orderItems, order2))
	}

	OrderDetails := npsSdk.NewOrderDetailsStruct()
	orderDetailsItems := reflect.ValueOf(OrderDetails).Elem().FieldByName("OrderItems")
	if orderDetailsItems.IsValid() { orderDetailsItems.Set(reflect.ValueOf(OrderItems)) }

	ShippingDetails := npsSdk.NewShippingDetailsStruct()

	ShippingAddress := npsSdk.NewAddressStruct()

	shippingAddressStreet := reflect.ValueOf(ShippingAddress).Elem().FieldByName("Street")
 	if shippingAddressStreet.IsValid() { shippingAddressStreet.SetString("shipping street") }

	shippingAddressHouseNumber := reflect.ValueOf(ShippingAddress).Elem().FieldByName("HouseNumber")
 	if shippingAddressHouseNumber.IsValid() { shippingAddressHouseNumber.SetString("1234") }

	shippingAddressCity := reflect.ValueOf(ShippingAddress).Elem().FieldByName("City")
 	if shippingAddressCity.IsValid() { shippingAddressCity.SetString("CABA") }

	shippingAddressCountry := reflect.ValueOf(ShippingAddress).Elem().FieldByName("Country")
 	if shippingAddressCountry.IsValid() { shippingAddressCountry.SetString("ARG") }

	shippingAddressAddInfo := reflect.ValueOf(ShippingAddress).Elem().FieldByName("AdditionalInfo")
 	if shippingAddressAddInfo.IsValid() { shippingAddressAddInfo.SetString("AdditionalInfo of shipping") }


	PrimaryRecipient := npsSdk.NewPersonStruct()

	primRecFirstName := reflect.ValueOf(PrimaryRecipient).Elem().FieldByName("FirstName")
 	if primRecFirstName.IsValid() { primRecFirstName.SetString("Pepe") }

	primRecLastName := reflect.ValueOf(PrimaryRecipient).Elem().FieldByName("LastName")
 	if primRecLastName.IsValid() { primRecLastName.SetString("Juan") }

	shippingDetailsMethod := reflect.ValueOf(ShippingDetails).Elem().FieldByName("Method")
 	if shippingDetailsMethod.IsValid() { shippingDetailsMethod.SetString("10") }

	shippingDetailsAddress := reflect.ValueOf(ShippingDetails).Elem().FieldByName("Address")
 	if shippingDetailsAddress.IsValid() { shippingDetailsAddress.Set(reflect.ValueOf(ShippingAddress)) }

	shippingDetailsPrimaryRecipient := reflect.ValueOf(ShippingDetails).Elem().FieldByName("PrimaryRecipient")
 	if shippingDetailsPrimaryRecipient.IsValid() { shippingDetailsPrimaryRecipient.Set(reflect.ValueOf(PrimaryRecipient)) }

	psp_AmountAdditionalDetails := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_AmountAdditionalDetails")
 	if psp_AmountAdditionalDetails.IsValid() { psp_AmountAdditionalDetails.Set(reflect.ValueOf(AmountAdditionalDetails)) }

	psp_BillingDetails := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_BillingDetails")
 	if psp_BillingDetails.IsValid() { psp_BillingDetails.Set(reflect.ValueOf(Billing)) }

	psp_CustomerAdditionalDetails := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_CustomerAdditionalDetails")
 	if psp_CustomerAdditionalDetails.IsValid() { psp_CustomerAdditionalDetails.Set(reflect.ValueOf(CustomerAdditionalDetails)) }

	psp_MerchantAdditionalDetails := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_MerchantAdditionalDetails")
 	if psp_MerchantAdditionalDetails.IsValid() { psp_MerchantAdditionalDetails.Set(reflect.ValueOf(MerchantAdditionalDetails)) }

	psp_OrderDetails := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_OrderDetails")
 	if psp_OrderDetails.IsValid() { psp_OrderDetails.Set(reflect.ValueOf(OrderDetails)) }

	psp_ShippingDetails := reflect.ValueOf(authorize2p).Elem().FieldByName("Psp_ShippingDetails")
 	if psp_ShippingDetails.IsValid() { psp_ShippingDetails.Set(reflect.ValueOf(ShippingDetails)) }

	resp, err := service.Authorize_2p(authorize2p)
	if err != nil {
		return err
	}

        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())

	return nil
}

func SendCreateClientSession(service *npsSdk.PaymentServicePlatformPortType) error {

	createClientSession := npsSdk.NewRequerimientoStruct_CreateClientSession()

	psp_Version := reflect.ValueOf(createClientSession).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(createClientSession).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PosDateTime := reflect.ValueOf(createClientSession).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2017-06-19 12:00:00") }

	resp, err := service.CreateClientSession(createClientSession)
	if err != nil {
		return err
	}

        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        Psp_ClientSession := reflect.ValueOf(resp).Elem().FieldByName("Psp_ClientSession")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	fmt.Printf("ClientSession = [%s]\n", Psp_ClientSession.String())

	return nil
}

func SendCreatePaymentMethod(service *npsSdk.PaymentServicePlatformPortType) error {

	createPaymentMethod := npsSdk.NewRequerimientoStruct_CreatePaymentMethod()

	psp_Version := reflect.ValueOf(createPaymentMethod).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(createPaymentMethod).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PosDateTime := reflect.ValueOf(createPaymentMethod).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2017-06-19 12:00:00") }

	PaymentMethod := npsSdk.NewPaymentMethodInputDetailsStruct()

	CardInputDetails := npsSdk.NewCardInputDetailsStruct()

        holderName := reflect.ValueOf(CardInputDetails).Elem().FieldByName("HolderName")
	if holderName.IsValid() { holderName.SetString("tester") }

        expirationDate := reflect.ValueOf(CardInputDetails).Elem().FieldByName("ExpirationDate")
	if expirationDate.IsValid() { expirationDate.SetString("1812") }

        cardNumber := reflect.ValueOf(CardInputDetails).Elem().FieldByName("Number")
	if cardNumber.IsValid() { cardNumber.SetString("4507990000000010") }

        securityCode := reflect.ValueOf(CardInputDetails).Elem().FieldByName("SecurityCode")
	if securityCode.IsValid() { securityCode.SetString("123") }

	cardInputDetails := reflect.ValueOf(PaymentMethod).Elem().FieldByName("CardInputDetails")
	if cardInputDetails.IsValid() { cardInputDetails.Set(reflect.ValueOf(CardInputDetails)) }

	psp_PaymentMethod := reflect.ValueOf(createPaymentMethod).Elem().FieldByName("Psp_PaymentMethod")
	if psp_PaymentMethod.IsValid() { psp_PaymentMethod.Set(reflect.ValueOf(PaymentMethod)) }

	resp, err := service.CreatePaymentMethod(createPaymentMethod)
	if err != nil {
		return err
	}

        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendCreatePaymentMethodToken(service *npsSdk.PaymentServicePlatformPortType) error {

	createPaymentMethodToken := npsSdk.NewRequerimientoStruct_CreatePaymentMethodToken()
	psp_Version := reflect.ValueOf(createPaymentMethodToken).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(createPaymentMethodToken).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PosDateTime := reflect.ValueOf(createPaymentMethodToken).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2017-06-19 12:00:00") }

	psp_ClientSession := reflect.ValueOf(createPaymentMethodToken).Elem().FieldByName("Psp_ClientSession")
	if psp_ClientSession.IsValid() { psp_ClientSession.SetString("w9mXsBXOlK5mx340F9WsXVoCTUsila2lqTBeqTwezfkbTt0li4NLXrFeFsbfXO9J") }

	CardInputDetails := npsSdk.NewCardInputDetailsStruct()

        holderName := reflect.ValueOf(CardInputDetails).Elem().FieldByName("HolderName")
	if holderName.IsValid() { holderName.SetString("tester") }

        expirationDate := reflect.ValueOf(CardInputDetails).Elem().FieldByName("ExpirationDate")
	if expirationDate.IsValid() { expirationDate.SetString("1812") }

        cardNumber := reflect.ValueOf(CardInputDetails).Elem().FieldByName("Number")
	if cardNumber.IsValid() { cardNumber.SetString("4507990000000010") }

        securityCode := reflect.ValueOf(CardInputDetails).Elem().FieldByName("SecurityCode")
	if securityCode.IsValid() { securityCode.SetString("123") }

	cardInputDetails := reflect.ValueOf(createPaymentMethodToken).Elem().FieldByName("Psp_CardInputDetails")
	if cardInputDetails.IsValid() { cardInputDetails.Set(reflect.ValueOf(CardInputDetails)) }

	person := npsSdk.NewPersonStruct()
        firstName := reflect.ValueOf(person).Elem().FieldByName("FirstName")
	if firstName.IsValid() { firstName.SetString("Silvina") }

        lastName := reflect.ValueOf(person).Elem().FieldByName("LastName")
	if lastName.IsValid() { lastName.SetString("Falconi") }

        phoneNumber1 := reflect.ValueOf(person).Elem().FieldByName("PhoneNumber1")
	if phoneNumber1.IsValid() { phoneNumber1.SetString("52520960") }

	psp_Person := reflect.ValueOf(createPaymentMethodToken).Elem().FieldByName("Psp_Person")
	if psp_Person.IsValid() { psp_Person.Set(reflect.ValueOf(person)) }

	address := npsSdk.NewAddressStruct()
        street := reflect.ValueOf(address).Elem().FieldByName("Street")
	if street.IsValid() { street.SetString("pelegrini") }

        city := reflect.ValueOf(address).Elem().FieldByName("City")
	if city.IsValid() { city.SetString("CABA") }

        country := reflect.ValueOf(address).Elem().FieldByName("Country")
	if country.IsValid() { country.SetString("ARG") }

        houseNumber := reflect.ValueOf(address).Elem().FieldByName("HouseNumber")
	if houseNumber.IsValid() { houseNumber.SetString("1111") }

	psp_Address := reflect.ValueOf(createPaymentMethodToken).Elem().FieldByName("Psp_Address")
	if psp_Address.IsValid() { psp_Address.Set(reflect.ValueOf(address)) }


	resp, err := service.CreatePaymentMethodToken(createPaymentMethodToken)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendBankPayment_2p(service *npsSdk.PaymentServicePlatformPortType) error {

	Taxes := npsSdk.NewArrayOf_TaxesRequestStruct()

        Items := reflect.ValueOf(Taxes).Elem().FieldByName("Items")
	if Items.Kind() == reflect.Slice {
          st := Items.Type()
          sliceType := st.Elem()
          if sliceType.Kind() == reflect.Ptr {
            sliceType = sliceType.Elem()
          }

         tax1 := reflect.New(sliceType)
         tax1.Elem().FieldByName("TypeId").SetString("500")
         tax1.Elem().FieldByName("Amount").SetString("50")  
         tax1.Elem().FieldByName("Rate").SetString("10")  
         Items.Set(reflect.Append(Items, tax1))
	}

	AmountAdditionalDetails := npsSdk.NewAmountAdditionalDetailsRequestStruct()

	amountAddTip := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Tip")
	if amountAddTip.IsValid() { amountAddTip.SetString("10") }

	amountAddDiscount := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Discount")
	if amountAddDiscount.IsValid() { amountAddDiscount.SetString("5") }

	amountAddTaxes := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Taxes")
	if amountAddTaxes.IsValid() { amountAddTaxes.Set(reflect.ValueOf(Taxes)) }

	Person := npsSdk.NewPersonStruct()
        firstName := reflect.ValueOf(Person).Elem().FieldByName("FirstName")
	if firstName.IsValid() { firstName.SetString("Silvina") }

        lastName := reflect.ValueOf(Person).Elem().FieldByName("LastName")
	if lastName.IsValid() { lastName.SetString("Falconi") }

        phoneNumber1 := reflect.ValueOf(Person).Elem().FieldByName("PhoneNumber1")
	if phoneNumber1.IsValid() { phoneNumber1.SetString("52520960") }

        idNumber := reflect.ValueOf(Person).Elem().FieldByName("IDNumber")
	if idNumber.IsValid() { idNumber.SetString("11111111") }

        idType := reflect.ValueOf(Person).Elem().FieldByName("IDType")
	if idType.IsValid() { idType.SetString("100") }


	Billing := npsSdk.NewBillingDetailsStruct()
        invoice := reflect.ValueOf(Billing).Elem().FieldByName("Invoice")
	if invoice.IsValid() { invoice.SetString("100001234") }

        invoiceAmount := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceAmount")
	if invoiceAmount.IsValid() { invoiceAmount.SetString("990") }

        invoiceCurrency := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceCurrency")
	if invoiceCurrency.IsValid() { invoiceCurrency.SetString("032") }

        person := reflect.ValueOf(Billing).Elem().FieldByName("Person")
	if person.IsValid() { person.Set(reflect.ValueOf(Person)) }

	CustomerAdditionalDetails := npsSdk.NewCustomerAdditionalDetailsStruct()

        emailAddress := reflect.ValueOf(CustomerAdditionalDetails).Elem().FieldByName("EmailAddress")
	if emailAddress.IsValid() { emailAddress.SetString("CustomerEmail@email.com.ar") }

	BankPayment_2p := npsSdk.NewRequerimientoStruct_BankPayment_2p()
	psp_Version := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_Product := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_Product")
	if psp_Product.IsValid() { psp_Product.SetString("320") }

	psp_MerchantId := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_MerchTxRef := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_MerchTxRef")
        t := time.Now()
 	if psp_MerchTxRef.IsValid() { psp_MerchTxRef.SetString(t.Format("20060102150405")) }

	psp_MerchOrderId := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_MerchOrderId")
	if psp_MerchOrderId.IsValid() { psp_MerchOrderId.SetString("ORD-10001") }

	psp_ScreenDescription := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_ScreenDescription")
	if psp_ScreenDescription.IsValid() { psp_ScreenDescription.SetString("CR DESC") }

	psp_TicketDescription := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_TicketDescription")
	if psp_TicketDescription.IsValid() { psp_TicketDescription.SetString("ticket desc") }

	psp_CustomerBankId := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_CustomerBankId")
	if psp_CustomerBankId.IsValid() { psp_CustomerBankId.SetString("HSBC") }

	psp_Amount1 := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_Amount1")
	if psp_Amount1.IsValid() { psp_Amount1.SetString("10000") }

	psp_Amount2 := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_Amount2")
	if psp_Amount2.IsValid() { psp_Amount2.SetString("20000") }

	psp_Country := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_Country")
	if psp_Country.IsValid() { psp_Country.SetString("ARG") }

	psp_Currency := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_Currency")
	if psp_Currency.IsValid() { psp_Currency.SetString("032") }

	psp_PosDateTime := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2017-06-19 12:00:00") }

	psp_TxSource := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_TxSource")
	if psp_TxSource.IsValid() { psp_TxSource.SetString("WEB") }

	psp_AmountAdditionalDetails := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_AmountAdditionalDetails")
	if psp_AmountAdditionalDetails.IsValid() { psp_AmountAdditionalDetails.Set(reflect.ValueOf(AmountAdditionalDetails)) }

	psp_BillingDetails := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_BillingDetails")
	if psp_BillingDetails.IsValid() { psp_BillingDetails.Set(reflect.ValueOf(Billing)) }

	psp_CustomerAdditionalDetails := reflect.ValueOf(BankPayment_2p).Elem().FieldByName("Psp_CustomerAdditionalDetails")
	if psp_CustomerAdditionalDetails.IsValid() { psp_CustomerAdditionalDetails.Set(reflect.ValueOf(CustomerAdditionalDetails)) }

	resp, err := service.BankPayment_2p(BankPayment_2p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")

        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendSplitPayOnLine_2p(service *npsSdk.PaymentServicePlatformPortType) error {

	VaultReference := npsSdk.NewVaultReference2pStruct()
	paymentMethodToken := reflect.ValueOf(VaultReference).Elem().FieldByName("PaymentMethodToken")
	if paymentMethodToken.IsValid() { paymentMethodToken.SetString("pm-token_lCY303k0vHvS5W06sPwzgoSHNt0VRrkG") }

	TransactionsArr := npsSdk.NewArrayOf_RequerimientoStruct_SplitPayOnLine_2p_Transactions()
	Items := reflect.ValueOf(TransactionsArr).Elem().FieldByName("Items")

       if Items.Kind() == reflect.Slice {
         st := Items.Type()
         sliceType := st.Elem()
         if sliceType.Kind() == reflect.Ptr {
            sliceType = sliceType.Elem()
         }

         trn := reflect.New(sliceType)
         trn.Elem().FieldByName("Psp_MerchantId").SetString("psp_test") 
         trn.Elem().FieldByName("Psp_MerchTxRef").SetString("ORDER66666-5") 
         trn.Elem().FieldByName("Psp_Product").SetString("55") 
         trn.Elem().FieldByName("Psp_Amount").SetString("10000") 
         trn.Elem().FieldByName("Psp_NumPayments").SetString("1") 

         Items.Set(reflect.Append(Items, trn))

         trn2 := reflect.New(sliceType)
         trn2.Elem().FieldByName("Psp_MerchantId").SetString("psp_test") 
         trn2.Elem().FieldByName("Psp_MerchTxRef").SetString("ORDER66666-6") 
         trn2.Elem().FieldByName("Psp_Product").SetString("23") 
         trn2.Elem().FieldByName("Psp_Amount").SetString("5050") 
         trn2.Elem().FieldByName("Psp_NumPayments").SetString("1") 

         Items.Set(reflect.Append(Items, trn2))

      }

      SplitPayOnLine_2p := npsSdk.NewRequerimientoStruct_SplitPayOnLine_2p()

      Transactions := reflect.ValueOf(SplitPayOnLine_2p).Elem().FieldByName("Psp_Transactions")
      if Transactions.IsValid() { Transactions.Set(reflect.ValueOf(TransactionsArr)) }

	psp_Version := reflect.ValueOf(SplitPayOnLine_2p).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_Product := reflect.ValueOf(SplitPayOnLine_2p).Elem().FieldByName("Psp_Product")
	if psp_Product.IsValid() { psp_Product.SetString("320") }

	psp_MerchantId := reflect.ValueOf(SplitPayOnLine_2p).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_Amount := reflect.ValueOf(SplitPayOnLine_2p).Elem().FieldByName("Psp_Amount")
	if psp_Amount.IsValid() { psp_Amount.SetString("1000") }

	psp_CardHolderName := reflect.ValueOf(SplitPayOnLine_2p).Elem().FieldByName("Psp_CardHolderName")
	if psp_CardHolderName.IsValid() { psp_CardHolderName.SetString("holder") }

	psp_CardNumber := reflect.ValueOf(SplitPayOnLine_2p).Elem().FieldByName("Psp_CardNumber")
	if psp_CardNumber.IsValid() { psp_CardNumber.SetString("4507990000000010") }

	psp_CardExpDate := reflect.ValueOf(SplitPayOnLine_2p).Elem().FieldByName("Psp_CardExpDate")
	if psp_CardExpDate.IsValid() { psp_CardExpDate.SetString("1903") }

	psp_CardSecurityCode := reflect.ValueOf(SplitPayOnLine_2p).Elem().FieldByName("Psp_CardSecurityCode")
	if psp_CardSecurityCode.IsValid() { psp_CardSecurityCode.SetString("306") }

	psp_TxSource := reflect.ValueOf(SplitPayOnLine_2p).Elem().FieldByName("Psp_TxSource")
	if psp_TxSource.IsValid() { psp_TxSource.SetString("WEB") }

	psp_MerchOrderId := reflect.ValueOf(SplitPayOnLine_2p).Elem().FieldByName("Psp_MerchOrderId")
	if psp_MerchOrderId.IsValid() { psp_MerchOrderId.SetString("ORD-20170629123000") }

	psp_Currency := reflect.ValueOf(SplitPayOnLine_2p).Elem().FieldByName("Psp_Currency")
	if psp_Currency.IsValid() { psp_Currency.SetString("032") }

	psp_Country := reflect.ValueOf(SplitPayOnLine_2p).Elem().FieldByName("Psp_Country")
	if psp_Country.IsValid() { psp_Country.SetString("ARG") }

	psp_PosDateTime := reflect.ValueOf(SplitPayOnLine_2p).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2016-12-01 12:00:00") }


	resp, err := service.SplitPayOnLine_2p(SplitPayOnLine_2p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendSplitAuthorize_2p(service *npsSdk.PaymentServicePlatformPortType) error {

	CustomerAdditionalDetails := npsSdk.NewCustomerAdditionalDetailsStruct()
	emailAddress := reflect.ValueOf(CustomerAdditionalDetails).Elem().FieldByName("EmailAddress")
	if emailAddress.IsValid() { emailAddress.SetString("CustomerEmail@email.com.ar") }

	Person := npsSdk.NewPersonStruct()
        firstName := reflect.ValueOf(Person).Elem().FieldByName("FirstName")
	if firstName.IsValid() { firstName.SetString("Silvina") }

        lastName := reflect.ValueOf(Person).Elem().FieldByName("LastName")
	if lastName.IsValid() { lastName.SetString("Falconi") }

        phoneNumber1 := reflect.ValueOf(Person).Elem().FieldByName("PhoneNumber1")
	if phoneNumber1.IsValid() { phoneNumber1.SetString("52520960") }


	Billing := npsSdk.NewBillingDetailsStruct()
        invoice := reflect.ValueOf(Billing).Elem().FieldByName("Invoice")
	if invoice.IsValid() { invoice.SetString("100001234") }

        invoiceAmount := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceAmount")
	if invoiceAmount.IsValid() { invoiceAmount.SetString("990") }

        invoiceCurrency := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceCurrency")
	if invoiceCurrency.IsValid() { invoiceCurrency.SetString("032") }

        person := reflect.ValueOf(Billing).Elem().FieldByName("Person")
	if person.IsValid() { person.Set(reflect.ValueOf(Person)) }


	SellerAddress := npsSdk.NewAddressStruct()
        sellerAddressCity := reflect.ValueOf(SellerAddress).Elem().FieldByName("City")
	if sellerAddressCity.IsValid() { sellerAddressCity.SetString("CABA") }

        sellerAddressCountry := reflect.ValueOf(SellerAddress).Elem().FieldByName("Country")
	if sellerAddressCountry.IsValid() { sellerAddressCountry.SetString("ARG") }

        sellerAddressStreet := reflect.ValueOf(SellerAddress).Elem().FieldByName("Street")
	if sellerAddressStreet.IsValid() { sellerAddressStreet.SetString("SellerStreet") }

        sellerAddressHouseNumber := reflect.ValueOf(SellerAddress).Elem().FieldByName("HouseNumber")
	if sellerAddressHouseNumber.IsValid() { sellerAddressHouseNumber.SetString("1234") }
        
	SellerDetails := npsSdk.NewSellerDetailsStruct()
        sellerDetailsName := reflect.ValueOf(SellerDetails).Elem().FieldByName("Name")
	if sellerDetailsName.IsValid() { sellerDetailsName.SetString("Seller Name") }

        sellerDetailsAddress := reflect.ValueOf(SellerDetails).Elem().FieldByName("Address")
	if sellerDetailsAddress.IsValid() { sellerDetailsAddress.Set(reflect.ValueOf(SellerAddress)) }


	MerchantAdditionalDetails := npsSdk.NewMerchantAdditionalDetailsStruct()
        shoppingCartInfo := reflect.ValueOf(MerchantAdditionalDetails).Elem().FieldByName("ShoppingCartInfo")
	if shoppingCartInfo.IsValid() { shoppingCartInfo.SetString("ShoppingCartInfo") }

        sellerDetails := reflect.ValueOf(MerchantAdditionalDetails).Elem().FieldByName("SellerDetails")
	if sellerDetails.IsValid() { sellerDetails.Set(reflect.ValueOf(SellerDetails)) }

      OrderItems := npsSdk.NewArrayOf_OrderItemStruct()

      Items := reflect.ValueOf(OrderItems).Elem().FieldByName("Items")
	
       if Items.Kind() == reflect.Slice {
         st := Items.Type()
         sliceType := st.Elem()
         if sliceType.Kind() == reflect.Ptr {
            sliceType = sliceType.Elem()
         }

         order1 := reflect.New(sliceType)
         order1.Elem().FieldByName("Description").SetString("producto 1") 
         order1.Elem().FieldByName("UnitPrice").SetString("10") 

         Items.Set(reflect.Append(Items, order1))

         order2 := reflect.New(sliceType)
         order2.Elem().FieldByName("Description").SetString("producto 2") 
         order2.Elem().FieldByName("UnitPrice").SetString("20") 

         Items.Set(reflect.Append(Items, order2))
         order3 := reflect.New(sliceType)
         order3.Elem().FieldByName("Description").SetString("producto 3") 
         order3.Elem().FieldByName("UnitPrice").SetString("30") 

         Items.Set(reflect.Append(Items, order3))

      }

      OrderDetails := npsSdk.NewOrderDetailsStruct()    
      OrderDetailsItems := reflect.ValueOf(OrderDetails).Elem().FieldByName("OrderItems")

      if OrderDetailsItems.IsValid() { OrderDetailsItems.Set(reflect.ValueOf(OrderItems)) }

	ShippingDetails := npsSdk.NewShippingDetailsStruct()

	ShippingAddress := npsSdk.NewAddressStruct()
        shippingAddressStreet := reflect.ValueOf(ShippingAddress).Elem().FieldByName("Street")
	if shippingAddressStreet.IsValid() { shippingAddressStreet.SetString("shipping street") }

        shippingAddressHouseNumber := reflect.ValueOf(ShippingAddress).Elem().FieldByName("HouseNumber")
	if shippingAddressHouseNumber.IsValid() { shippingAddressHouseNumber.SetString("1234") }

        shippingAddressHouseCity := reflect.ValueOf(ShippingAddress).Elem().FieldByName("City")
	if shippingAddressHouseCity.IsValid() { shippingAddressHouseCity.SetString("CABA") }

        shippingAddressHouseCountry := reflect.ValueOf(ShippingAddress).Elem().FieldByName("Country")
	if shippingAddressHouseCountry.IsValid() { shippingAddressHouseCountry.SetString("ARG") }

        additionalInfo := reflect.ValueOf(ShippingAddress).Elem().FieldByName("AdditionalInfo")
	if additionalInfo.IsValid() { additionalInfo.SetString("AdditionalInfo of shipping") }


	PrimaryRecipient := npsSdk.NewPersonStruct()
        primaryRecipientFirstName := reflect.ValueOf(PrimaryRecipient).Elem().FieldByName("FirstName")
	if primaryRecipientFirstName.IsValid() { primaryRecipientFirstName.SetString("Pepe") }

        primaryRecipientFirstNameLastName := reflect.ValueOf(PrimaryRecipient).Elem().FieldByName("LastName")
	if primaryRecipientFirstNameLastName.IsValid() { primaryRecipientFirstNameLastName.SetString("Juan") }

        shippingDetailsMethod := reflect.ValueOf(ShippingDetails).Elem().FieldByName("Method")
	if shippingDetailsMethod.IsValid() { shippingDetailsMethod.SetString("10") }

        shippingDetailsAddress := reflect.ValueOf(ShippingDetails).Elem().FieldByName("Address")
	if shippingDetailsAddress.IsValid() { shippingDetailsAddress.Set(reflect.ValueOf(ShippingAddress)) }

        shippingDetailsPrimaryRecipient := reflect.ValueOf(ShippingDetails).Elem().FieldByName("PrimaryRecipient")
	if shippingDetailsPrimaryRecipient.IsValid() { shippingDetailsPrimaryRecipient.Set(reflect.ValueOf(PrimaryRecipient)) }

	VaultReference := npsSdk.NewVaultReference2pStruct()
        paymentMethodToken := reflect.ValueOf(VaultReference).Elem().FieldByName("PaymentMethodToken")
	if paymentMethodToken.IsValid() { paymentMethodToken.SetString("pm-token_lCY303k0vHvS5W06sPwzgoSHNt0VRrkG") }

	TransactionsArr := npsSdk.NewArrayOf_RequerimientoStruct_SplitAuthorize_2p_Transactions()
	TransacItems := reflect.ValueOf(TransactionsArr).Elem().FieldByName("Items")

       if TransacItems.Kind() == reflect.Slice {
         st := TransacItems.Type()
         sliceType := st.Elem()
         if sliceType.Kind() == reflect.Ptr {
            sliceType = sliceType.Elem()
         }

         trn := reflect.New(sliceType)
         trn.Elem().FieldByName("Psp_MerchantId").SetString("psp_test") 
         trn.Elem().FieldByName("Psp_MerchTxRef").SetString("ORDER66666-5") 
         trn.Elem().FieldByName("Psp_Product").SetString("14") 
         trn.Elem().FieldByName("Psp_Amount").SetString("10000") 
         trn.Elem().FieldByName("Psp_NumPayments").SetString("1") 
         trn.Elem().FieldByName("Psp_VaultReference").Set(reflect.ValueOf(VaultReference)) 

         TransacItems.Set(reflect.Append(TransacItems, trn))

         trn2 := reflect.New(sliceType)
         trn2.Elem().FieldByName("Psp_MerchantId").SetString("psp_test") 
         trn2.Elem().FieldByName("Psp_MerchTxRef").SetString("ORDER66666-6") 
         trn2.Elem().FieldByName("Psp_Product").SetString("14") 
         trn2.Elem().FieldByName("Psp_Amount").SetString("5050") 
         trn2.Elem().FieldByName("Psp_NumPayments").SetString("1") 

         TransacItems.Set(reflect.Append(TransacItems, trn2))

      }

      SplitAuthorize2p := npsSdk.NewRequerimientoStruct_SplitAuthorize_2p()
      Transactions := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_Transactions")
      if Transactions.IsValid() { Transactions.Set(reflect.ValueOf(TransactionsArr)) }

      psp_Version := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_Version")
      if psp_Version.IsValid() { psp_Version.SetString("2.2") }

      psp_MerchantId := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_MerchantId")
      if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

      psp_TxSource := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_TxSource")
      if psp_TxSource.IsValid() { psp_TxSource.SetString("WEB") }

      psp_MerchOrderId := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_MerchOrderId")
      if psp_MerchOrderId.IsValid() { psp_MerchOrderId.SetString("20170609150900-1") }

      psp_Amount := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_Amount")
      if psp_Amount.IsValid() { psp_Amount.SetString("15050") }

      psp_Currency := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_Currency")
      if psp_Currency.IsValid() { psp_Currency.SetString("032") }

      psp_Country := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_Country")
      if psp_Country.IsValid() { psp_Country.SetString("ARG") }

      psp_Product := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_Product")
      if psp_Product.IsValid() { psp_Product.SetString("14") }

      psp_CardNumber := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_CardNumber")
      if psp_CardNumber.IsValid() { psp_CardNumber.SetString("4507990000000010") }

      psp_CardExpDate := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_CardExpDate")
      if psp_CardExpDate.IsValid() { psp_CardExpDate.SetString("1903") }

      psp_PosDateTime := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_PosDateTime")
      if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2016-12-01 12:00:00") }

      psp_BillingDetails := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_BillingDetails")
      if psp_BillingDetails.IsValid() { psp_BillingDetails.Set(reflect.ValueOf(Billing)) }

      psp_CustomerAdditionalDetails := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_CustomerAdditionalDetails")
      if psp_CustomerAdditionalDetails.IsValid() { psp_CustomerAdditionalDetails.Set(reflect.ValueOf(CustomerAdditionalDetails)) }

      psp_MerchantAdditionalDetails := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_MerchantAdditionalDetails")
      if psp_MerchantAdditionalDetails.IsValid() { psp_MerchantAdditionalDetails.Set(reflect.ValueOf(MerchantAdditionalDetails)) }

      psp_OrderDetails := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_OrderDetails")
      if psp_OrderDetails.IsValid() { psp_OrderDetails.Set(reflect.ValueOf(OrderDetails)) }

      psp_ShippingDetails := reflect.ValueOf(SplitAuthorize2p).Elem().FieldByName("Psp_ShippingDetails")
      if psp_ShippingDetails.IsValid() { psp_ShippingDetails.Set(reflect.ValueOf(ShippingDetails)) }

	resp, err := service.SplitAuthorize_2p(SplitAuthorize2p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendPayOnLine_3p(service *npsSdk.PaymentServicePlatformPortType) error {

	Person := npsSdk.NewPersonStruct()
        firstName := reflect.ValueOf(Person).Elem().FieldByName("FirstName")
        lastName := reflect.ValueOf(Person).Elem().FieldByName("LastName")
        phoneNumber1 := reflect.ValueOf(Person).Elem().FieldByName("PhoneNumber1")

        if firstName.IsValid() { firstName.SetString("Silvina") }
        if lastName.IsValid() { lastName.SetString("Falconi") }
        if phoneNumber1.IsValid() { phoneNumber1.SetString("52520960") }

	AmountAdditionalDetails := npsSdk.NewAmountAdditionalDetailsRequestStruct()
        tip := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Tip")
        discount := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Discount")

        if tip.IsValid() { tip.SetString("10") }
        if discount.IsValid() { discount.SetString("5") }

	Billing := npsSdk.NewBillingDetailsStruct()
        invoice := reflect.ValueOf(Billing).Elem().FieldByName("Invoice")
        invoiceAmount := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceAmount")
        invoiceCurrency := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceCurrency")
        billingPerson := reflect.ValueOf(Billing).Elem().FieldByName("Person")

        if invoice.IsValid() { invoice.SetString("100001234") }
        if invoiceAmount.IsValid() { invoiceAmount.SetString("990") }
        if invoiceCurrency.IsValid() { invoiceCurrency.SetString("032") }
        if billingPerson.IsValid() { billingPerson.Set(reflect.ValueOf(Person)) }

	SellerAddress := npsSdk.NewAddressStruct()
        city := reflect.ValueOf(SellerAddress).Elem().FieldByName("City")
        country := reflect.ValueOf(SellerAddress).Elem().FieldByName("Country")
        street := reflect.ValueOf(SellerAddress).Elem().FieldByName("Street")
        houseNumber := reflect.ValueOf(SellerAddress).Elem().FieldByName("HouseNumber")
        
        if city.IsValid() { city.SetString("CABA") }
        if country.IsValid() { country.SetString("ARG") }
        if street.IsValid() { street.SetString("SellerStreet") }
        if houseNumber.IsValid() { houseNumber.SetString("1234") }

	SellerDetails := npsSdk.NewSellerDetailsStruct()
        sellerDetailsName := reflect.ValueOf(SellerDetails).Elem().FieldByName("Name")
        sellerDetailsAddress := reflect.ValueOf(SellerDetails).Elem().FieldByName("Address")

        if sellerDetailsName.IsValid() { sellerDetailsName.SetString("Seller Name") }
        if sellerDetailsAddress.IsValid() { sellerDetailsAddress.Set(reflect.ValueOf(SellerAddress)) }

	MerchantAdditionalDetails := npsSdk.NewMerchantAdditionalDetailsStruct()
        shoppingCartInfo := reflect.ValueOf(MerchantAdditionalDetails).Elem().FieldByName("ShoppingCartInfo")
        merchantAdditionalDetailsSellerDetails := reflect.ValueOf(MerchantAdditionalDetails).Elem().FieldByName("SellerDetails")
        
        if shoppingCartInfo.IsValid() { shoppingCartInfo.SetString("ShoppingCartInfo") }
        if merchantAdditionalDetailsSellerDetails.IsValid() { merchantAdditionalDetailsSellerDetails.Set(reflect.ValueOf(SellerDetails)) }

	CustomerAdditionalDetails := npsSdk.NewCustomerAdditionalDetailsStruct()
	emailAddress := reflect.ValueOf(CustomerAdditionalDetails).Elem().FieldByName("EmailAddress")
        if emailAddress.IsValid() { emailAddress.SetString("mailAddr@mail.com.ar") }

      OrderItems := npsSdk.NewArrayOf_OrderItemStruct()

      Items := reflect.ValueOf(OrderItems).Elem().FieldByName("Items")
	
       if Items.Kind() == reflect.Slice {
         st := Items.Type()
         sliceType := st.Elem()
         if sliceType.Kind() == reflect.Ptr {
            sliceType = sliceType.Elem()
         }

         order1 := reflect.New(sliceType)
         order1.Elem().FieldByName("Description").SetString("producto 1") 
         order1.Elem().FieldByName("UnitPrice").SetString("10") 

         Items.Set(reflect.Append(Items, order1))

         order2 := reflect.New(sliceType)
         order2.Elem().FieldByName("Description").SetString("producto 2") 
         order2.Elem().FieldByName("UnitPrice").SetString("20") 

         Items.Set(reflect.Append(Items, order2))
         order3 := reflect.New(sliceType)
         order3.Elem().FieldByName("Description").SetString("producto 3") 
         order3.Elem().FieldByName("UnitPrice").SetString("30") 

         Items.Set(reflect.Append(Items, order3))

      }

      OrderDetails := npsSdk.NewOrderDetailsStruct()    
      OrderDetailsItems := reflect.ValueOf(OrderDetails).Elem().FieldByName("OrderItems")

      if OrderDetailsItems.IsValid() { OrderDetailsItems.Set(reflect.ValueOf(OrderItems)) }

	PayOnline3p := npsSdk.NewRequerimientoStruct_PayOnLine_3p()

        psp_Version := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_Version")
        if psp_Version.IsValid() { psp_Version.SetString("2.2") }

        psp_MerchantId := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_MerchantId")
        if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

        psp_TxSource := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_TxSource")
        if psp_TxSource.IsValid() { psp_TxSource.SetString("WEB") }

        psp_MerchTxRef := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_MerchTxRef")
        if psp_MerchTxRef.IsValid() { psp_MerchTxRef.SetString("ORDER56666-3") }

        psp_MerchOrderId := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_MerchOrderId")
        if psp_MerchOrderId.IsValid() { psp_MerchOrderId.SetString("ORDER56666") }

        psp_Amount := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_Amount")
        if psp_Amount.IsValid() { psp_Amount.SetString("1000") }

        psp_NumPayments := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_NumPayments")
        if psp_NumPayments.IsValid() { psp_NumPayments.SetString("1") }

        psp_Currency := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_Currency")
        if psp_Currency.IsValid() { psp_Currency.SetString("032") }

        psp_Country := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_Country")
        if psp_Country.IsValid() { psp_Country.SetString("ARG") }

        psp_Product := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_Product")
        if psp_Product.IsValid() { psp_Product.SetString("14") }

        psp_CustomerMail := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_CustomerMail")
        if psp_CustomerMail.IsValid() { psp_CustomerMail.SetString("yourmail@gmail") }

        psp_SoftDescriptor := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_SoftDescriptor")
        if psp_SoftDescriptor.IsValid() { psp_SoftDescriptor.SetString("Sol Tropical E") }

        psp_PosDateTime := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_PosDateTime")
        if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2016-12-01 12:00:00") }

        psp_ReturnURL := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_ReturnURL")
        if psp_ReturnURL.IsValid() { psp_ReturnURL.SetString("http://psp-client.localhost/simple_query_tx.php") }

        psp_FrmLanguage := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_FrmLanguage")
        if psp_FrmLanguage.IsValid() { psp_FrmLanguage.SetString("en_US") }

        psp_OrderDetails := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_OrderDetails")
        if psp_OrderDetails.IsValid() { psp_OrderDetails.Set(reflect.ValueOf(OrderDetails)) }

        psp_CustomerAdditionalDetails := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_CustomerAdditionalDetails")
        if psp_CustomerAdditionalDetails.IsValid() { psp_CustomerAdditionalDetails.Set(reflect.ValueOf(CustomerAdditionalDetails)) }

        psp_AmountAdditionalDetails := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_AmountAdditionalDetails")
        if psp_AmountAdditionalDetails.IsValid() { psp_AmountAdditionalDetails.Set(reflect.ValueOf(AmountAdditionalDetails)) }

        psp_BillingDetails := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_BillingDetails")
        if psp_BillingDetails.IsValid() { psp_BillingDetails.Set(reflect.ValueOf(Billing)) }

        psp_MerchantAdditionalDetails := reflect.ValueOf(PayOnline3p).Elem().FieldByName("Psp_MerchantAdditionalDetails")
        if psp_MerchantAdditionalDetails.IsValid() { psp_MerchantAdditionalDetails.Set(reflect.ValueOf(MerchantAdditionalDetails)) }

	resp, err := service.PayOnLine_3p(PayOnline3p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendAuthorize_3p(service *npsSdk.PaymentServicePlatformPortType) error {

	CustomerAdditionalDetails := npsSdk.NewCustomerAdditionalDetailsStruct()
	emailAddress := reflect.ValueOf(CustomerAdditionalDetails).Elem().FieldByName("EmailAddress")
        if emailAddress.IsValid() { emailAddress.SetString("CustomerEmail@email.com.ar") }

	Person := npsSdk.NewPersonStruct()
        firstName := reflect.ValueOf(Person).Elem().FieldByName("FirstName")
        lastName := reflect.ValueOf(Person).Elem().FieldByName("LastName")
        phoneNumber1 := reflect.ValueOf(Person).Elem().FieldByName("PhoneNumber1")

        if firstName.IsValid() { firstName.SetString("Silvina") }
        if lastName.IsValid() { lastName.SetString("Falconi") }
        if phoneNumber1.IsValid() { phoneNumber1.SetString("52520960") }

	Billing := npsSdk.NewBillingDetailsStruct()
        invoice := reflect.ValueOf(Billing).Elem().FieldByName("Invoice")
        invoiceAmount := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceAmount")
        invoiceCurrency := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceCurrency")
        billingPerson := reflect.ValueOf(Billing).Elem().FieldByName("Person")

        if invoice.IsValid() { invoice.SetString("100001234") }
        if invoiceAmount.IsValid() { invoiceAmount.SetString("990") }
        if invoiceCurrency.IsValid() { invoiceCurrency.SetString("032") }
        if billingPerson.IsValid() { billingPerson.Set(reflect.ValueOf(Person)) }

	SellerAddress := npsSdk.NewAddressStruct()
        city := reflect.ValueOf(SellerAddress).Elem().FieldByName("City")
        country := reflect.ValueOf(SellerAddress).Elem().FieldByName("Country")
        street := reflect.ValueOf(SellerAddress).Elem().FieldByName("Street")
        houseNumber := reflect.ValueOf(SellerAddress).Elem().FieldByName("HouseNumber")
        
        if city.IsValid() { city.SetString("CABA") }
        if country.IsValid() { country.SetString("ARG") }
        if street.IsValid() { street.SetString("SellerStreet") }
        if houseNumber.IsValid() { houseNumber.SetString("1234") }

	SellerDetails := npsSdk.NewSellerDetailsStruct()
        sellerDetailsName := reflect.ValueOf(SellerDetails).Elem().FieldByName("Name")
        sellerDetailsAddress := reflect.ValueOf(SellerDetails).Elem().FieldByName("Address")

        if sellerDetailsName.IsValid() { sellerDetailsName.SetString("Seller Name") }
        if sellerDetailsAddress.IsValid() { sellerDetailsAddress.Set(reflect.ValueOf(SellerAddress)) }

      OrderItems := npsSdk.NewArrayOf_OrderItemStruct()

      Items := reflect.ValueOf(OrderItems).Elem().FieldByName("Items")
	
       if Items.Kind() == reflect.Slice {
         st := Items.Type()
         sliceType := st.Elem()
         if sliceType.Kind() == reflect.Ptr {
            sliceType = sliceType.Elem()
         }

         order1 := reflect.New(sliceType)
         order1.Elem().FieldByName("Description").SetString("producto 1") 
         order1.Elem().FieldByName("UnitPrice").SetString("10") 

         Items.Set(reflect.Append(Items, order1))

         order2 := reflect.New(sliceType)
         order2.Elem().FieldByName("Description").SetString("producto 2") 
         order2.Elem().FieldByName("UnitPrice").SetString("20") 

         Items.Set(reflect.Append(Items, order2))
         order3 := reflect.New(sliceType)
         order3.Elem().FieldByName("Description").SetString("producto 3") 
         order3.Elem().FieldByName("UnitPrice").SetString("30") 

         Items.Set(reflect.Append(Items, order3))

      }

      OrderDetails := npsSdk.NewOrderDetailsStruct()    
      OrderDetailsItems := reflect.ValueOf(OrderDetails).Elem().FieldByName("OrderItems")

      if OrderDetailsItems.IsValid() { OrderDetailsItems.Set(reflect.ValueOf(OrderItems)) }

	ShippingDetails := npsSdk.NewShippingDetailsStruct()

	ShippingAddress := npsSdk.NewAddressStruct()

	shippingAddressStreet := reflect.ValueOf(ShippingAddress).Elem().FieldByName("Street")
 	if shippingAddressStreet.IsValid() { shippingAddressStreet.SetString("shipping street") }

	shippingAddressHouseNumber := reflect.ValueOf(ShippingAddress).Elem().FieldByName("HouseNumber")
 	if shippingAddressHouseNumber.IsValid() { shippingAddressHouseNumber.SetString("1234") }

	shippingAddressCity := reflect.ValueOf(ShippingAddress).Elem().FieldByName("City")
 	if shippingAddressCity.IsValid() { shippingAddressCity.SetString("CABA") }

	shippingAddressCountry := reflect.ValueOf(ShippingAddress).Elem().FieldByName("Country")
 	if shippingAddressCountry.IsValid() { shippingAddressCountry.SetString("ARG") }

	shippingAddressAddInfo := reflect.ValueOf(ShippingAddress).Elem().FieldByName("AdditionalInfo")
 	if shippingAddressAddInfo.IsValid() { shippingAddressAddInfo.SetString("AdditionalInfo of shipping") }

	PrimaryRecipient := npsSdk.NewPersonStruct()

	primRecFirstName := reflect.ValueOf(PrimaryRecipient).Elem().FieldByName("FirstName")
 	if primRecFirstName.IsValid() { primRecFirstName.SetString("Pepe") }

	primRecLastName := reflect.ValueOf(PrimaryRecipient).Elem().FieldByName("LastName")
 	if primRecLastName.IsValid() { primRecLastName.SetString("Juan") }

	shippingDetailsMethod := reflect.ValueOf(ShippingDetails).Elem().FieldByName("Method")
 	if shippingDetailsMethod.IsValid() { shippingDetailsMethod.SetString("10") }

	shippingDetailsAddress := reflect.ValueOf(ShippingDetails).Elem().FieldByName("Address")
 	if shippingDetailsAddress.IsValid() { shippingDetailsAddress.Set(reflect.ValueOf(ShippingAddress)) }

	shippingDetailsPrimaryRecipient := reflect.ValueOf(ShippingDetails).Elem().FieldByName("PrimaryRecipient")
 	if shippingDetailsPrimaryRecipient.IsValid() { shippingDetailsPrimaryRecipient.Set(reflect.ValueOf(PrimaryRecipient)) }

	Authorize3p := npsSdk.NewRequerimientoStruct_Authorize_3p()

        psp_Version := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_Version")
        if psp_Version.IsValid() { psp_Version.SetString("2.2") }

        psp_MerchantId := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_MerchantId")
        if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

        psp_TxSource := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_TxSource")
        if psp_TxSource.IsValid() { psp_TxSource.SetString("WEB") }

        psp_MerchTxRef := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_MerchTxRef")
        if psp_MerchTxRef.IsValid() { psp_MerchTxRef.SetString("ORDER56666-3") }

        psp_MerchOrderId := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_MerchOrderId")
        if psp_MerchOrderId.IsValid() { psp_MerchOrderId.SetString("ORDER56666") }

        psp_Amount := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_Amount")
        if psp_Amount.IsValid() { psp_Amount.SetString("15050") }

        psp_NumPayments := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_NumPayments")
        if psp_NumPayments.IsValid() { psp_NumPayments.SetString("1") }

        psp_Currency := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_Currency")
        if psp_Currency.IsValid() { psp_Currency.SetString("032") }

        psp_Country := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_Country")
        if psp_Country.IsValid() { psp_Country.SetString("ARG") }

        psp_Product := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_Product")
        if psp_Product.IsValid() { psp_Product.SetString("14") }

        psp_CustomerMail := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_CustomerMail")
        if psp_CustomerMail.IsValid() { psp_CustomerMail.SetString("yourmail@gmail") }

        psp_SoftDescriptor := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_SoftDescriptor")
        if psp_SoftDescriptor.IsValid() { psp_SoftDescriptor.SetString("Sol Tropical E") }

        psp_PosDateTime := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_PosDateTime")
        if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2016-12-01 12:00:00") }

        psp_ReturnURL := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_ReturnURL")
        if psp_ReturnURL.IsValid() { psp_ReturnURL.SetString("http://psp-client.localhost/simple_query_tx.php") }

        psp_FrmLanguage := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_FrmLanguage")
        if psp_FrmLanguage.IsValid() { psp_FrmLanguage.SetString("en_US") }

        psp_OrderDetails := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_OrderDetails")
        if psp_OrderDetails.IsValid() { psp_OrderDetails.Set(reflect.ValueOf(OrderDetails)) }

        psp_ShippingDetails := reflect.ValueOf(Authorize3p).Elem().FieldByName("Psp_ShippingDetails")
        if psp_ShippingDetails.IsValid() { psp_ShippingDetails.Set(reflect.ValueOf(ShippingDetails)) }

	resp, err := service.Authorize_3p(Authorize3p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendSplitAuthorize_3p(service *npsSdk.PaymentServicePlatformPortType) error {
      OrderItems := npsSdk.NewArrayOf_OrderItemStruct()

      Items := reflect.ValueOf(OrderItems).Elem().FieldByName("Items")
	
       if Items.Kind() == reflect.Slice {
         st := Items.Type()
         sliceType := st.Elem()
         if sliceType.Kind() == reflect.Ptr {
            sliceType = sliceType.Elem()
         }

         order1 := reflect.New(sliceType)
         order1.Elem().FieldByName("Description").SetString("producto 1") 
         order1.Elem().FieldByName("UnitPrice").SetString("10") 

         Items.Set(reflect.Append(Items, order1))

         order2 := reflect.New(sliceType)
         order2.Elem().FieldByName("Description").SetString("producto 2") 
         order2.Elem().FieldByName("UnitPrice").SetString("20") 

         Items.Set(reflect.Append(Items, order2))
         order3 := reflect.New(sliceType)
         order3.Elem().FieldByName("Description").SetString("producto 3") 
         order3.Elem().FieldByName("UnitPrice").SetString("30") 

         Items.Set(reflect.Append(Items, order3))

      }

      OrderDetails := npsSdk.NewOrderDetailsStruct()    
      OrderDetailsItems := reflect.ValueOf(OrderDetails).Elem().FieldByName("OrderItems")

      if OrderDetailsItems.IsValid() { OrderDetailsItems.Set(reflect.ValueOf(OrderItems)) }

	ShippingDetails := npsSdk.NewShippingDetailsStruct()

	ShippingAddress := npsSdk.NewAddressStruct()

	shippingAddressStreet := reflect.ValueOf(ShippingAddress).Elem().FieldByName("Street")
 	if shippingAddressStreet.IsValid() { shippingAddressStreet.SetString("shipping street") }

	shippingAddressHouseNumber := reflect.ValueOf(ShippingAddress).Elem().FieldByName("HouseNumber")
 	if shippingAddressHouseNumber.IsValid() { shippingAddressHouseNumber.SetString("1234") }

	shippingAddressCity := reflect.ValueOf(ShippingAddress).Elem().FieldByName("City")
 	if shippingAddressCity.IsValid() { shippingAddressCity.SetString("CABA") }

	shippingAddressCountry := reflect.ValueOf(ShippingAddress).Elem().FieldByName("Country")
 	if shippingAddressCountry.IsValid() { shippingAddressCountry.SetString("ARG") }

	shippingAddressAddInfo := reflect.ValueOf(ShippingAddress).Elem().FieldByName("AdditionalInfo")
 	if shippingAddressAddInfo.IsValid() { shippingAddressAddInfo.SetString("AdditionalInfo of shipping") }

	PrimaryRecipient := npsSdk.NewPersonStruct()

	primRecFirstName := reflect.ValueOf(PrimaryRecipient).Elem().FieldByName("FirstName")
 	if primRecFirstName.IsValid() { primRecFirstName.SetString("Pepe") }

	primRecLastName := reflect.ValueOf(PrimaryRecipient).Elem().FieldByName("LastName")
 	if primRecLastName.IsValid() { primRecLastName.SetString("Juan") }

	shippingDetailsMethod := reflect.ValueOf(ShippingDetails).Elem().FieldByName("Method")
 	if shippingDetailsMethod.IsValid() { shippingDetailsMethod.SetString("10") }

	shippingDetailsAddress := reflect.ValueOf(ShippingDetails).Elem().FieldByName("Address")
 	if shippingDetailsAddress.IsValid() { shippingDetailsAddress.Set(reflect.ValueOf(ShippingAddress)) }

	shippingDetailsPrimaryRecipient := reflect.ValueOf(ShippingDetails).Elem().FieldByName("PrimaryRecipient")
 	if shippingDetailsPrimaryRecipient.IsValid() { shippingDetailsPrimaryRecipient.Set(reflect.ValueOf(PrimaryRecipient)) }

	TransactionsArr := npsSdk.NewArrayOf_RequerimientoStruct_SplitAuthorize_3p_Transactions()
	trnItems := reflect.ValueOf(TransactionsArr).Elem().FieldByName("Items")

       if trnItems.Kind() == reflect.Slice {
         st := trnItems.Type()
         sliceType := st.Elem()
         if sliceType.Kind() == reflect.Ptr {
            sliceType = sliceType.Elem()
         }

         trn := reflect.New(sliceType)
         trn.Elem().FieldByName("Psp_MerchantId").SetString("psp_test") 
         trn.Elem().FieldByName("Psp_MerchTxRef").SetString("ORDER66666-5") 
         trn.Elem().FieldByName("Psp_Product").SetString("14") 
         trn.Elem().FieldByName("Psp_Amount").SetString("10000") 
         trn.Elem().FieldByName("Psp_NumPayments").SetString("1") 

         trnItems.Set(reflect.Append(trnItems, trn))

         trn2 := reflect.New(sliceType)
         trn2.Elem().FieldByName("Psp_MerchantId").SetString("psp_test") 
         trn2.Elem().FieldByName("Psp_MerchTxRef").SetString("ORDER66666-6") 
         trn2.Elem().FieldByName("Psp_Product").SetString("14") 
         trn2.Elem().FieldByName("Psp_Amount").SetString("5050") 
         trn2.Elem().FieldByName("Psp_NumPayments").SetString("1") 

         trnItems.Set(reflect.Append(trnItems, trn2))

      }


	SplitAuthorize3p := npsSdk.NewRequerimientoStruct_SplitAuthorize_3p()
	Transactions := reflect.ValueOf(SplitAuthorize3p).Elem().FieldByName("Psp_Transactions")
        if Transactions.IsValid() { Transactions.Set(reflect.ValueOf(TransactionsArr)) }

        psp_Version := reflect.ValueOf(SplitAuthorize3p).Elem().FieldByName("Psp_Version")
        if psp_Version.IsValid() { psp_Version.SetString("2.2") }

        psp_MerchantId := reflect.ValueOf(SplitAuthorize3p).Elem().FieldByName("Psp_MerchantId")
        if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

        psp_TxSource := reflect.ValueOf(SplitAuthorize3p).Elem().FieldByName("Psp_TxSource")
        if psp_TxSource.IsValid() { psp_TxSource.SetString("WEB") }

        psp_MerchOrderId := reflect.ValueOf(SplitAuthorize3p).Elem().FieldByName("Psp_MerchOrderId")
        if psp_MerchOrderId.IsValid() { psp_MerchOrderId.SetString("ORDER56666") }

        psp_Amount := reflect.ValueOf(SplitAuthorize3p).Elem().FieldByName("Psp_Amount")
        if psp_Amount.IsValid() { psp_Amount.SetString("15050") }

        psp_Currency := reflect.ValueOf(SplitAuthorize3p).Elem().FieldByName("Psp_Currency")
        if psp_Currency.IsValid() { psp_Currency.SetString("032") }

        psp_Country := reflect.ValueOf(SplitAuthorize3p).Elem().FieldByName("Psp_Country")
        if psp_Country.IsValid() { psp_Country.SetString("ARG") }

        psp_Product := reflect.ValueOf(SplitAuthorize3p).Elem().FieldByName("Psp_Product")
        if psp_Product.IsValid() { psp_Product.SetString("14") }

        psp_PosDateTime := reflect.ValueOf(SplitAuthorize3p).Elem().FieldByName("Psp_PosDateTime")
        if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2016-12-01 12:00:00") }

        psp_ReturnURL := reflect.ValueOf(SplitAuthorize3p).Elem().FieldByName("Psp_ReturnURL")
        if psp_ReturnURL.IsValid() { psp_ReturnURL.SetString("http://psp-client.localhost/simple_query_tx.php") }

        psp_FrmLanguage := reflect.ValueOf(SplitAuthorize3p).Elem().FieldByName("Psp_FrmLanguage")
        if psp_FrmLanguage.IsValid() { psp_FrmLanguage.SetString("en_US") }

        psp_OrderDetails := reflect.ValueOf(SplitAuthorize3p).Elem().FieldByName("Psp_OrderDetails")
        if psp_OrderDetails.IsValid() { psp_OrderDetails.Set(reflect.ValueOf(OrderDetails)) }

        psp_ShippingDetails := reflect.ValueOf(SplitAuthorize3p).Elem().FieldByName("Psp_ShippingDetails")
        if psp_ShippingDetails.IsValid() { psp_ShippingDetails.Set(reflect.ValueOf(ShippingDetails)) }

	resp, err := service.SplitAuthorize_3p(SplitAuthorize3p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendBankPayment_3p(service *npsSdk.PaymentServicePlatformPortType) error {
	Taxes := npsSdk.NewArrayOf_TaxesRequestStruct()

        Items := reflect.ValueOf(Taxes).Elem().FieldByName("Items")
	if Items.Kind() == reflect.Slice {
          st := Items.Type()
          sliceType := st.Elem()
          if sliceType.Kind() == reflect.Ptr {
            sliceType = sliceType.Elem()
          }

         tax1 := reflect.New(sliceType)
         tax1.Elem().FieldByName("TypeId").SetString("500")
         tax1.Elem().FieldByName("Amount").SetString("50")  
         tax1.Elem().FieldByName("Rate").SetString("10")  
         Items.Set(reflect.Append(Items, tax1))
	}

	AmountAdditionalDetails := npsSdk.NewAmountAdditionalDetailsRequestStruct()

	amountAddTip := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Tip")
	if amountAddTip.IsValid() { amountAddTip.SetString("10") }

	amountAddDiscount := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Discount")
	if amountAddDiscount.IsValid() { amountAddDiscount.SetString("5") }

	amountAddTaxes := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Taxes")
	if amountAddTaxes.IsValid() { amountAddTaxes.Set(reflect.ValueOf(Taxes)) }

	Person := npsSdk.NewPersonStruct()
        firstName := reflect.ValueOf(Person).Elem().FieldByName("FirstName")
	if firstName.IsValid() { firstName.SetString("Silvina") }

        lastName := reflect.ValueOf(Person).Elem().FieldByName("LastName")
	if lastName.IsValid() { lastName.SetString("Falconi") }

        phoneNumber1 := reflect.ValueOf(Person).Elem().FieldByName("PhoneNumber1")
	if phoneNumber1.IsValid() { phoneNumber1.SetString("52520960") }

        idNumber := reflect.ValueOf(Person).Elem().FieldByName("IDNumber")
	if idNumber.IsValid() { idNumber.SetString("11111111") }

        idType := reflect.ValueOf(Person).Elem().FieldByName("IDType")
	if idType.IsValid() { idType.SetString("100") }


	Billing := npsSdk.NewBillingDetailsStruct()
        invoice := reflect.ValueOf(Billing).Elem().FieldByName("Invoice")
	if invoice.IsValid() { invoice.SetString("100001234") }

        invoiceAmount := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceAmount")
	if invoiceAmount.IsValid() { invoiceAmount.SetString("990") }

        invoiceCurrency := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceCurrency")
	if invoiceCurrency.IsValid() { invoiceCurrency.SetString("032") }

        person := reflect.ValueOf(Billing).Elem().FieldByName("Person")
	if person.IsValid() { person.Set(reflect.ValueOf(Person)) }


	CustomerAdditionalDetails := npsSdk.NewCustomerAdditionalDetailsStruct()
        emailAddress := reflect.ValueOf(Billing).Elem().FieldByName("EmailAddress")
	if emailAddress.IsValid() { emailAddress.SetString("CustomerEmail@email.com.ar") }

	BankPayment_3p := npsSdk.NewRequerimientoStruct_BankPayment_3p()
	psp_Version := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_Product := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_Product")
	if psp_Product.IsValid() { psp_Product.SetString("320") }

	psp_MerchantId := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_TxSource := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_TxSource")
	if psp_TxSource.IsValid() { psp_TxSource.SetString("WEB") }

	psp_MerchTxRef := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_MerchTxRef")
        t := time.Now()
 	if psp_MerchTxRef.IsValid() { psp_MerchTxRef.SetString(t.Format("20060102150405")) }

	psp_MerchOrderId := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_MerchOrderId")
	if psp_MerchOrderId.IsValid() { psp_MerchOrderId.SetString("ORD-10001") }

	psp_ScreenDescription := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_ScreenDescription")
	if psp_ScreenDescription.IsValid() { psp_ScreenDescription.SetString("SCR DESC") }

	psp_TicketDescription := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_TicketDescription")
	if psp_TicketDescription.IsValid() { psp_TicketDescription.SetString("ticket desc") }

	psp_ReturnURL := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_ReturnURL")
	if psp_ReturnURL.IsValid() { psp_ReturnURL.SetString("http://localhost/") }

	psp_FrmLanguage := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_FrmLanguage")
	if psp_FrmLanguage.IsValid() { psp_FrmLanguage.SetString("es_AR") }

	psp_Amount1 := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_Amount1")
	if psp_Amount1.IsValid() { psp_Amount1.SetString("10000") }

	psp_ExpDate1 := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_ExpDate1")
	if psp_ExpDate1.IsValid() { psp_ExpDate1.SetString("2017-09-01") }

	psp_Amount2 := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_Amount2")
	if psp_Amount2.IsValid() { psp_Amount2.SetString("20000") }

	psp_ExpDate2 := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_ExpDate2")
	if psp_ExpDate2.IsValid() { psp_ExpDate2.SetString("2017-09-01") }

	psp_Country := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_Country")
	if psp_Country.IsValid() { psp_Country.SetString("ARG") }

	psp_Currency := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_Currency")
	if psp_Currency.IsValid() { psp_Currency.SetString("032") }

	psp_PosDateTime := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2016-12-01 12:00:00") }

        psp_AmountAdditionalDetails := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_AmountAdditionalDetails")
        if psp_AmountAdditionalDetails.IsValid() { psp_AmountAdditionalDetails.Set(reflect.ValueOf(AmountAdditionalDetails)) }

        psp_BillingDetails := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_BillingDetails")
        if psp_BillingDetails.IsValid() { psp_BillingDetails.Set(reflect.ValueOf(Billing)) }

        psp_CustomerAdditionalDetails := reflect.ValueOf(BankPayment_3p).Elem().FieldByName("Psp_CustomerAdditionalDetails")
        if psp_CustomerAdditionalDetails.IsValid() { psp_CustomerAdditionalDetails.Set(reflect.ValueOf(CustomerAdditionalDetails)) }


	resp, err := service.BankPayment_3p(BankPayment_3p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendCashPayment_3p(service *npsSdk.PaymentServicePlatformPortType) error {
	Taxes := npsSdk.NewArrayOf_TaxesRequestStruct()

        Items := reflect.ValueOf(Taxes).Elem().FieldByName("Items")
	if Items.Kind() == reflect.Slice {
          st := Items.Type()
          sliceType := st.Elem()
          if sliceType.Kind() == reflect.Ptr {
            sliceType = sliceType.Elem()
          }

         tax1 := reflect.New(sliceType)
         tax1.Elem().FieldByName("TypeId").SetString("500")
         tax1.Elem().FieldByName("Amount").SetString("50")  
         tax1.Elem().FieldByName("Rate").SetString("10")  
         Items.Set(reflect.Append(Items, tax1))
	}

	AmountAdditionalDetails := npsSdk.NewAmountAdditionalDetailsRequestStruct()

	amountAddTip := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Tip")
	if amountAddTip.IsValid() { amountAddTip.SetString("10") }

	amountAddDiscount := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Discount")
	if amountAddDiscount.IsValid() { amountAddDiscount.SetString("5") }

	amountAddTaxes := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Taxes")
	if amountAddTaxes.IsValid() { amountAddTaxes.Set(reflect.ValueOf(Taxes)) }

	Person := npsSdk.NewPersonStruct()
        firstName := reflect.ValueOf(Person).Elem().FieldByName("FirstName")
	if firstName.IsValid() { firstName.SetString("Silvina") }

        lastName := reflect.ValueOf(Person).Elem().FieldByName("LastName")
	if lastName.IsValid() { lastName.SetString("Falconi") }

        phoneNumber1 := reflect.ValueOf(Person).Elem().FieldByName("PhoneNumber1")
	if phoneNumber1.IsValid() { phoneNumber1.SetString("52520960") }

        idNumber := reflect.ValueOf(Person).Elem().FieldByName("IDNumber")
	if idNumber.IsValid() { idNumber.SetString("11111111") }

        idType := reflect.ValueOf(Person).Elem().FieldByName("IDType")
	if idType.IsValid() { idType.SetString("100") }


	Billing := npsSdk.NewBillingDetailsStruct()
        invoice := reflect.ValueOf(Billing).Elem().FieldByName("Invoice")
	if invoice.IsValid() { invoice.SetString("100001234") }

        invoiceAmount := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceAmount")
	if invoiceAmount.IsValid() { invoiceAmount.SetString("990") }

        invoiceCurrency := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceCurrency")
	if invoiceCurrency.IsValid() { invoiceCurrency.SetString("032") }

        person := reflect.ValueOf(Billing).Elem().FieldByName("Person")
	if person.IsValid() { person.Set(reflect.ValueOf(Person)) }


	CustomerAdditionalDetails := npsSdk.NewCustomerAdditionalDetailsStruct()
        emailAddress := reflect.ValueOf(Billing).Elem().FieldByName("EmailAddress")
	if emailAddress.IsValid() { emailAddress.SetString("CustomerEmail@email.com.ar") }

      OrderItems := npsSdk.NewArrayOf_OrderItemStruct()

      orderItems := reflect.ValueOf(OrderItems).Elem().FieldByName("Items")
	
       if orderItems.Kind() == reflect.Slice {
         st := orderItems.Type()
         sliceType := st.Elem()
         if sliceType.Kind() == reflect.Ptr {
            sliceType = sliceType.Elem()
         }

         order1 := reflect.New(sliceType)
         order1.Elem().FieldByName("Description").SetString("producto 1") 
         order1.Elem().FieldByName("UnitPrice").SetString("10") 

         orderItems.Set(reflect.Append(orderItems, order1))

         order2 := reflect.New(sliceType)
         order2.Elem().FieldByName("Description").SetString("producto 2") 
         order2.Elem().FieldByName("UnitPrice").SetString("20") 

         orderItems.Set(reflect.Append(orderItems, order2))
      }

      OrderDetails := npsSdk.NewOrderDetailsStruct()    
      OrderDetailsItems := reflect.ValueOf(OrderDetails).Elem().FieldByName("OrderItems")

      if OrderDetailsItems.IsValid() { OrderDetailsItems.Set(reflect.ValueOf(OrderItems)) }

	CashPayment_3p := npsSdk.NewRequerimientoStruct_CashPayment_3p()

	psp_Version := reflect.ValueOf(CashPayment_3p).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_Product := reflect.ValueOf(CashPayment_3p).Elem().FieldByName("Psp_Product")
	if psp_Product.IsValid() { psp_Product.SetString("320") }

	psp_MerchantId := reflect.ValueOf(CashPayment_3p).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_TxSource := reflect.ValueOf(CashPayment_3p).Elem().FieldByName("Psp_TxSource")
	if psp_TxSource.IsValid() { psp_TxSource.SetString("WEB") }

	psp_MerchTxRef := reflect.ValueOf(CashPayment_3p).Elem().FieldByName("Psp_MerchTxRef")
        t := time.Now()
 	if psp_MerchTxRef.IsValid() { psp_MerchTxRef.SetString(t.Format("20060102150405")) }

	psp_MerchOrderId := reflect.ValueOf(CashPayment_3p).Elem().FieldByName("Psp_MerchOrderId")
	if psp_MerchOrderId.IsValid() { psp_MerchOrderId.SetString("ORD-10001") }

	psp_ReturnURL := reflect.ValueOf(CashPayment_3p).Elem().FieldByName("Psp_ReturnURL")
	if psp_ReturnURL.IsValid() { psp_ReturnURL.SetString("http://localhost/") }

	psp_FrmLanguage := reflect.ValueOf(CashPayment_3p).Elem().FieldByName("Psp_FrmLanguage")
	if psp_FrmLanguage.IsValid() { psp_FrmLanguage.SetString("es_AR") }

	psp_Amount := reflect.ValueOf(CashPayment_3p).Elem().FieldByName("Psp_Amount")
	if psp_Amount.IsValid() { psp_Amount.SetString("11000") }

	psp_Country := reflect.ValueOf(CashPayment_3p).Elem().FieldByName("Psp_Country")
	if psp_Country.IsValid() { psp_Country.SetString("ARG") }

	psp_Currency := reflect.ValueOf(CashPayment_3p).Elem().FieldByName("Psp_Currency")
	if psp_Currency.IsValid() { psp_Currency.SetString("032") }

	psp_PosDateTime := reflect.ValueOf(CashPayment_3p).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2016-12-01 12:00:00") }

        psp_AmountAdditionalDetails := reflect.ValueOf(CashPayment_3p).Elem().FieldByName("Psp_AmountAdditionalDetails")
        if psp_AmountAdditionalDetails.IsValid() { psp_AmountAdditionalDetails.Set(reflect.ValueOf(AmountAdditionalDetails)) }

        psp_BillingDetails := reflect.ValueOf(CashPayment_3p).Elem().FieldByName("Psp_BillingDetails")
        if psp_BillingDetails.IsValid() { psp_BillingDetails.Set(reflect.ValueOf(Billing)) }

        psp_CustomerAdditionalDetails := reflect.ValueOf(CashPayment_3p).Elem().FieldByName("Psp_CustomerAdditionalDetails")
        if psp_CustomerAdditionalDetails.IsValid() { psp_CustomerAdditionalDetails.Set(reflect.ValueOf(CustomerAdditionalDetails)) }


	resp, err := service.CashPayment_3p(CashPayment_3p)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendCapture(service *npsSdk.PaymentServicePlatformPortType) error {
	Taxes := npsSdk.NewArrayOf_TaxesRequestStruct()

        Items := reflect.ValueOf(Taxes).Elem().FieldByName("Items")
	if Items.Kind() == reflect.Slice {
          st := Items.Type()
          sliceType := st.Elem()
          if sliceType.Kind() == reflect.Ptr {
            sliceType = sliceType.Elem()
          }

         tax1 := reflect.New(sliceType)
         tax1.Elem().FieldByName("TypeId").SetString("500")
         tax1.Elem().FieldByName("Amount").SetString("50")  
         tax1.Elem().FieldByName("Rate").SetString("10")  
         Items.Set(reflect.Append(Items, tax1))
	}

	AmountAdditionalDetails := npsSdk.NewAmountAdditionalDetailsRequestStruct()

	amountAddTip := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Tip")
	if amountAddTip.IsValid() { amountAddTip.SetString("10") }

	amountAddDiscount := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Discount")
	if amountAddDiscount.IsValid() { amountAddDiscount.SetString("5") }

	amountAddTaxes := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Taxes")
	if amountAddTaxes.IsValid() { amountAddTaxes.Set(reflect.ValueOf(Taxes)) }

	Capture := npsSdk.NewRequerimientoStruct_Capture()
	psp_Version := reflect.ValueOf(Capture).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_Product := reflect.ValueOf(Capture).Elem().FieldByName("Psp_Product")
	if psp_Product.IsValid() { psp_Product.SetString("320") }

	psp_MerchantId := reflect.ValueOf(Capture).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_TxSource := reflect.ValueOf(Capture).Elem().FieldByName("Psp_TxSource")
	if psp_TxSource.IsValid() { psp_TxSource.SetString("WEB") }

	psp_MerchTxRef := reflect.ValueOf(Capture).Elem().FieldByName("Psp_MerchTxRef")
        t := time.Now()
 	if psp_MerchTxRef.IsValid() { psp_MerchTxRef.SetString(t.Format("20060102150405")) }

	psp_MerchOrderId := reflect.ValueOf(Capture).Elem().FieldByName("Psp_MerchOrderId")
	if psp_MerchOrderId.IsValid() { psp_MerchOrderId.SetString("ORD-10001") }

	psp_PosDateTime := reflect.ValueOf(Capture).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2016-12-01 12:00:00") }

	psp_TransactionId_Orig := reflect.ValueOf(Capture).Elem().FieldByName("Psp_TransactionId_Orig")
	if psp_TransactionId_Orig.IsValid() { psp_TransactionId_Orig.SetString("1868712") }

	psp_AmountToCapture := reflect.ValueOf(Capture).Elem().FieldByName("Psp_AmountToCapture")
	if psp_AmountToCapture.IsValid() { psp_AmountToCapture.SetString("15000") }

        psp_AmountAdditionalDetails := reflect.ValueOf(Capture).Elem().FieldByName("Psp_AmountAdditionalDetails")
        if psp_AmountAdditionalDetails.IsValid() { psp_AmountAdditionalDetails.Set(reflect.ValueOf(AmountAdditionalDetails)) }

	resp, err := service.Capture(Capture)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendRefund(service *npsSdk.PaymentServicePlatformPortType) error {

	Refund := npsSdk.NewRequerimientoStruct_Refund()
	psp_Version := reflect.ValueOf(Refund).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_Product := reflect.ValueOf(Refund).Elem().FieldByName("Psp_Product")
	if psp_Product.IsValid() { psp_Product.SetString("320") }

	psp_MerchantId := reflect.ValueOf(Refund).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_TxSource := reflect.ValueOf(Refund).Elem().FieldByName("Psp_TxSource")
	if psp_TxSource.IsValid() { psp_TxSource.SetString("WEB") }

	psp_MerchTxRef := reflect.ValueOf(Refund).Elem().FieldByName("Psp_MerchTxRef")
        t := time.Now()
 	if psp_MerchTxRef.IsValid() { psp_MerchTxRef.SetString(t.Format("20060102150405")) }

	psp_MerchOrderId := reflect.ValueOf(Refund).Elem().FieldByName("Psp_MerchOrderId")
	if psp_MerchOrderId.IsValid() { psp_MerchOrderId.SetString("ORD-10001") }

	psp_PosDateTime := reflect.ValueOf(Refund).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2016-12-01 12:00:00") }

	psp_TransactionId_Orig := reflect.ValueOf(Refund).Elem().FieldByName("Psp_TransactionId_Orig")
	if psp_TransactionId_Orig.IsValid() { psp_TransactionId_Orig.SetString("1868717") }

	psp_AmountToRefund := reflect.ValueOf(Refund).Elem().FieldByName("Psp_AmountToRefund")
	if psp_AmountToRefund.IsValid() { psp_AmountToRefund.SetString("10000") }

	resp, err := service.Refund(Refund)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendRetrievePaymentMethodToken(service *npsSdk.PaymentServicePlatformPortType) error {

	RetrievePaymentMethodToken := npsSdk.NewRequerimientoStruct_RetrievePaymentMethodToken()
	psp_Version := reflect.ValueOf(RetrievePaymentMethodToken).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(RetrievePaymentMethodToken).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PaymentMethodToken := reflect.ValueOf(RetrievePaymentMethodToken).Elem().FieldByName("Psp_PaymentMethodToken")
	if psp_PaymentMethodToken.IsValid() { psp_PaymentMethodToken.SetString("FG2WE0mVhoOm4U2a1R8qsBe307mtiVy0") }

	psp_ClientSession := reflect.ValueOf(RetrievePaymentMethodToken).Elem().FieldByName("Psp_ClientSession")
	if psp_ClientSession.IsValid() { psp_ClientSession.SetString("w9mXsBXOlK5mx340F9WsXVoCTUsila2lqTBeqTwezfkbTt0li4NLXrFeFsbfXO9J") }

	resp, err := service.RetrievePaymentMethodToken(RetrievePaymentMethodToken)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendRecachePaymentMethodToken(service *npsSdk.PaymentServicePlatformPortType) error {

	RecachePaymentMethodToken := npsSdk.NewRequerimientoStruct_RecachePaymentMethodToken()
	psp_Version := reflect.ValueOf(RecachePaymentMethodToken).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(RecachePaymentMethodToken).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PaymentMethodId := reflect.ValueOf(RecachePaymentMethodToken).Elem().FieldByName("Psp_PaymentMethodId")
	if psp_PaymentMethodId.IsValid() { psp_PaymentMethodId.SetString("ZSUjUAEcb5uQ3qtNSS7N7gyDMPtODkjt") }

	psp_ClientSession := reflect.ValueOf(RecachePaymentMethodToken).Elem().FieldByName("Psp_ClientSession")
	if psp_ClientSession.IsValid() { psp_ClientSession.SetString("w9mXsBXOlK5mx340F9WsXVoCTUsila2lqTBeqTwezfkbTt0li4NLXrFeFsbfXO9J") }

	psp_CardSecurityCode := reflect.ValueOf(RecachePaymentMethodToken).Elem().FieldByName("Psp_CardSecurityCode")
	if psp_CardSecurityCode.IsValid() { psp_CardSecurityCode.SetString("123") }

	resp, err := service.RecachePaymentMethodToken(RecachePaymentMethodToken)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendCreatePaymentMethodFromPayment(service *npsSdk.PaymentServicePlatformPortType) error {

	createPaymentMethodFromPayment := npsSdk.NewRequerimientoStruct_CreatePaymentMethodFromPayment()
	psp_Version := reflect.ValueOf(createPaymentMethodFromPayment).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(createPaymentMethodFromPayment).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PosDateTime := reflect.ValueOf(createPaymentMethodFromPayment).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2017-06-19 12:00:00") }

	psp_TransactionId := reflect.ValueOf(createPaymentMethodFromPayment).Elem().FieldByName("Psp_TransactionId")
	if psp_TransactionId.IsValid() { psp_TransactionId.SetString("1868712") }

	resp, err := service.CreatePaymentMethodFromPayment(createPaymentMethodFromPayment)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendRetrievePaymentMethod(service *npsSdk.PaymentServicePlatformPortType) error {

	RetrievePaymentMethod := npsSdk.NewRequerimientoStruct_RetrievePaymentMethod()
	psp_Version := reflect.ValueOf(RetrievePaymentMethod).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(RetrievePaymentMethod).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PaymentMethodId := reflect.ValueOf(RetrievePaymentMethod).Elem().FieldByName("Psp_PaymentMethodId")
	if psp_PaymentMethodId.IsValid() { psp_PaymentMethodId.SetString("ZSUjUAEcb5uQ3qtNSS7N7gyDMPtODkjt") }

	psp_PosDateTime := reflect.ValueOf(RetrievePaymentMethod).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2016-12-01 12:00:00") }

	resp, err := service.RetrievePaymentMethod(RetrievePaymentMethod)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendUpdatePaymentMethod(service *npsSdk.PaymentServicePlatformPortType) error {

	UpdatePaymentMethod := npsSdk.NewRequerimientoStruct_UpdatePaymentMethod()
	psp_Version := reflect.ValueOf(UpdatePaymentMethod).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(UpdatePaymentMethod).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PosDateTime := reflect.ValueOf(UpdatePaymentMethod).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2016-12-01 12:00:00") }

	cardInputDetails := npsSdk.NewCardInputUpdateDetailsStruct()

        holderName := reflect.ValueOf(cardInputDetails).Elem().FieldByName("HolderName")
	if holderName.IsValid() { holderName.SetString("tester") }

        expirationDate := reflect.ValueOf(cardInputDetails).Elem().FieldByName("ExpirationDate")
	if expirationDate.IsValid() { expirationDate.SetString("1812") }

        cardNumber := reflect.ValueOf(cardInputDetails).Elem().FieldByName("Number")
	if cardNumber.IsValid() { cardNumber.SetString("4507990000000010") }

        securityCode := reflect.ValueOf(cardInputDetails).Elem().FieldByName("SecurityCode")
	if securityCode.IsValid() { securityCode.SetString("123") }

	psp_CardInputDetails := reflect.ValueOf(UpdatePaymentMethod).Elem().FieldByName("Psp_CardInputDetails")
	if psp_CardInputDetails.IsValid() { psp_CardInputDetails.Set(reflect.ValueOf(cardInputDetails)) }

	psp_PaymentMethodId := reflect.ValueOf(UpdatePaymentMethod).Elem().FieldByName("Psp_PaymentMethodId")
	if psp_PaymentMethodId.IsValid() { psp_PaymentMethodId.SetString("ZSUjUAEcb5uQ3qtNSS7N7gyDMPtODkjt") }

	resp, err := service.UpdatePaymentMethod(UpdatePaymentMethod)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendDeletePaymentMethod(service *npsSdk.PaymentServicePlatformPortType) error {

	DeletePaymentMethod := npsSdk.NewRequerimientoStruct_DeletePaymentMethod()

	psp_Version := reflect.ValueOf(DeletePaymentMethod).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(DeletePaymentMethod).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PosDateTime := reflect.ValueOf(DeletePaymentMethod).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2016-12-01 12:00:00") }

        psp_PaymentMethodId := reflect.ValueOf(DeletePaymentMethod).Elem().FieldByName("Psp_PaymentMethodId")
	if psp_PaymentMethodId.IsValid() { psp_PaymentMethodId.SetString("ZSUjUAEcb5uQ3qtNSS7N7gyDMPtODkjt") }


	resp, err := service.DeletePaymentMethod(DeletePaymentMethod)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendCreateCustomer(service *npsSdk.PaymentServicePlatformPortType) error {

	Person := npsSdk.NewPersonStruct()
        firstName := reflect.ValueOf(Person).Elem().FieldByName("FirstName")
        if firstName.IsValid() { firstName.SetString("Silvina") }

        lastName := reflect.ValueOf(Person).Elem().FieldByName("LastName")
        if lastName.IsValid() { lastName.SetString("Falconi") }

        phoneNumber1 := reflect.ValueOf(Person).Elem().FieldByName("PhoneNumber1")
        if phoneNumber1.IsValid() { phoneNumber1.SetString("52520960") }

	Address := npsSdk.NewAddressStruct()
        city := reflect.ValueOf(Address).Elem().FieldByName("City")
        if city.IsValid() { city.SetString("CABA") }

        country := reflect.ValueOf(Address).Elem().FieldByName("Country")
        if country.IsValid() { country.SetString("ARG") }

        street := reflect.ValueOf(Address).Elem().FieldByName("Street")
        if street.IsValid() { street.SetString("SellerStreet") }

        houseNumber := reflect.ValueOf(Address).Elem().FieldByName("HouseNumber")
        if houseNumber.IsValid() { houseNumber.SetString("1234") }

	PmPerson := npsSdk.NewPersonStruct()
        pmPersonFirstName := reflect.ValueOf(PmPerson).Elem().FieldByName("FirstName")
        if pmPersonFirstName.IsValid() { pmPersonFirstName.SetString("Silvina") }

        pmPersonLastName := reflect.ValueOf(PmPerson).Elem().FieldByName("LastName")
        if pmPersonLastName.IsValid() { pmPersonLastName.SetString("Falconi") }

        pmPersonPhoneNumber1 := reflect.ValueOf(PmPerson).Elem().FieldByName("PhoneNumber1")
        if pmPersonPhoneNumber1.IsValid() { pmPersonPhoneNumber1.SetString("52520960") }

	PmAddress := npsSdk.NewAddressStruct()
        PmAddressCity := reflect.ValueOf(PmAddress).Elem().FieldByName("City")
        if PmAddressCity.IsValid() { PmAddressCity.SetString("CABA") }

        PmAddressCountry := reflect.ValueOf(PmAddress).Elem().FieldByName("Country")
        if PmAddressCountry.IsValid() { PmAddressCountry.SetString("ARG") }

        PmAddressStreet := reflect.ValueOf(PmAddress).Elem().FieldByName("Street")
        if PmAddressStreet.IsValid() { PmAddressStreet.SetString("SellerStreet") }

        PmAddressHouseNumber := reflect.ValueOf(PmAddress).Elem().FieldByName("HouseNumber")
        if PmAddressHouseNumber.IsValid() { PmAddressHouseNumber.SetString("1234") }

	CardInputDetails := npsSdk.NewCardInputDetailsStruct()

        holderName := reflect.ValueOf(CardInputDetails).Elem().FieldByName("HolderName")
	if holderName.IsValid() { holderName.SetString("tester") }

        expirationDate := reflect.ValueOf(CardInputDetails).Elem().FieldByName("ExpirationDate")
	if expirationDate.IsValid() { expirationDate.SetString("1812") }

        cardNumber := reflect.ValueOf(CardInputDetails).Elem().FieldByName("Number")
	if cardNumber.IsValid() { cardNumber.SetString("4507990000000010") }

        securityCode := reflect.ValueOf(CardInputDetails).Elem().FieldByName("SecurityCode")
	if securityCode.IsValid() { securityCode.SetString("123") }

	pm := npsSdk.NewPaymentMethodInputDetailsStruct()
	address := reflect.ValueOf(pm).Elem().FieldByName("Address")
	if address.IsValid() { address.Set(reflect.ValueOf(PmAddress)) }

	person := reflect.ValueOf(pm).Elem().FieldByName("Person")
	if person.IsValid() { person.Set(reflect.ValueOf(PmPerson)) }

	cardInputDetails := reflect.ValueOf(pm).Elem().FieldByName("CardInputDetails")
	if cardInputDetails.IsValid() { cardInputDetails.Set(reflect.ValueOf(CardInputDetails)) }

	CreateCustomer := npsSdk.NewRequerimientoStruct_CreateCustomer()
	psp_Version := reflect.ValueOf(CreateCustomer).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(CreateCustomer).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_EmailAddress := reflect.ValueOf(CreateCustomer).Elem().FieldByName("Psp_EmailAddress")
	if psp_EmailAddress.IsValid() { psp_EmailAddress.SetString("pspEmail@emai.com") }

	psp_PosDateTime := reflect.ValueOf(CreateCustomer).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2017-06-19 12:00:00") }

	psp_Person := reflect.ValueOf(CreateCustomer).Elem().FieldByName("Psp_Person")
	if psp_Person.IsValid() { psp_Person.Set(reflect.ValueOf(Person)) }

	psp_Address := reflect.ValueOf(CreateCustomer).Elem().FieldByName("Psp_Address")
	if psp_Address.IsValid() { psp_Address.Set(reflect.ValueOf(Address)) }

	psp_PaymentMethod := reflect.ValueOf(CreateCustomer).Elem().FieldByName("Psp_PaymentMethod")
	if psp_PaymentMethod.IsValid() { psp_PaymentMethod.Set(reflect.ValueOf(pm)) }

	resp, err := service.CreateCustomer(CreateCustomer)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendRetrieveCustomer(service *npsSdk.PaymentServicePlatformPortType) error {

	RetrieveCustomer := npsSdk.NewRequerimientoStruct_RetrieveCustomer()
	psp_Version := reflect.ValueOf(RetrieveCustomer).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(RetrieveCustomer).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PosDateTime := reflect.ValueOf(RetrieveCustomer).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2017-06-19 12:00:00") }

	psp_CustomerId := reflect.ValueOf(RetrieveCustomer).Elem().FieldByName("Psp_CustomerId")
	if psp_MerchantId.IsValid() { psp_CustomerId.SetString("8H9T839wLt83DldqV0qxCMJC3hwvL8yF") }

	resp, err := service.RetrieveCustomer(RetrieveCustomer)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendUpdateCustomer(service *npsSdk.PaymentServicePlatformPortType) error {

	Person := npsSdk.NewPersonStruct()
        firstName := reflect.ValueOf(Person).Elem().FieldByName("FirstName")
        if firstName.IsValid() { firstName.SetString("Silvina") }

        lastName := reflect.ValueOf(Person).Elem().FieldByName("LastName")
        if lastName.IsValid() { lastName.SetString("Falconi") }

        phoneNumber1 := reflect.ValueOf(Person).Elem().FieldByName("PhoneNumber1")
        if phoneNumber1.IsValid() { phoneNumber1.SetString("123456") }

	Address := npsSdk.NewAddressStruct()
        city := reflect.ValueOf(Address).Elem().FieldByName("City")
        if city.IsValid() { city.SetString("CABA") }

        country := reflect.ValueOf(Address).Elem().FieldByName("Country")
        if country.IsValid() { country.SetString("ARG") }

        street := reflect.ValueOf(Address).Elem().FieldByName("Street")
        if street.IsValid() { street.SetString("SellerStreet") }

        houseNumber := reflect.ValueOf(Address).Elem().FieldByName("HouseNumber")
        if houseNumber.IsValid() { houseNumber.SetString("1234") }

	PmPerson := npsSdk.NewPersonStruct()
        pmPersonFirstName := reflect.ValueOf(PmPerson).Elem().FieldByName("FirstName")
        if pmPersonFirstName.IsValid() { pmPersonFirstName.SetString("Silvina") }

        pmPersonLastName := reflect.ValueOf(PmPerson).Elem().FieldByName("LastName")
        if pmPersonLastName.IsValid() { pmPersonLastName.SetString("Falconi") }

        pmPersonPhoneNumber1 := reflect.ValueOf(PmPerson).Elem().FieldByName("PhoneNumber1")
        if pmPersonPhoneNumber1.IsValid() { pmPersonPhoneNumber1.SetString("52520960") }

	PmAddress := npsSdk.NewAddressStruct()
        PmAddressCity := reflect.ValueOf(PmAddress).Elem().FieldByName("City")
        if PmAddressCity.IsValid() { PmAddressCity.SetString("CABA") }

        PmAddressCountry := reflect.ValueOf(PmAddress).Elem().FieldByName("Country")
        if PmAddressCountry.IsValid() { PmAddressCountry.SetString("ARG") }

        PmAddressStreet := reflect.ValueOf(PmAddress).Elem().FieldByName("Street")
        if PmAddressStreet.IsValid() { PmAddressStreet.SetString("SellerStreet") }

        PmAddressHouseNumber := reflect.ValueOf(PmAddress).Elem().FieldByName("HouseNumber")
        if PmAddressHouseNumber.IsValid() { PmAddressHouseNumber.SetString("1234") }

	pm := npsSdk.NewPaymentMethodInputDetailsStruct()
	paymentMethodToken := reflect.ValueOf(pm).Elem().FieldByName("PaymentMethodToken")
	if paymentMethodToken.IsValid() { paymentMethodToken.SetString("pm-token_lCY303k0vHvS5W06sPwzgoSHNt0VRrkG") }

	address := reflect.ValueOf(pm).Elem().FieldByName("Address")
	if address.IsValid() { address.Set(reflect.ValueOf(PmAddress)) }

	person := reflect.ValueOf(pm).Elem().FieldByName("Person")
	if person.IsValid() { person.Set(reflect.ValueOf(PmPerson)) }

	UpdateCustomer := npsSdk.NewRequerimientoStruct_UpdateCustomer()
	psp_Version := reflect.ValueOf(UpdateCustomer).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(UpdateCustomer).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PosDateTime := reflect.ValueOf(UpdateCustomer).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2017-06-19 12:00:00") }

	psp_CustomerId := reflect.ValueOf(UpdateCustomer).Elem().FieldByName("Psp_CustomerId")
	if psp_CustomerId.IsValid() { psp_CustomerId.SetString("8H9T839wLt83DldqV0qxCMJC3hwvL8yF") }

	psp_EmailAddress := reflect.ValueOf(UpdateCustomer).Elem().FieldByName("Psp_EmailAddress")
	if psp_EmailAddress.IsValid() { psp_EmailAddress.SetString("CustomerEmail@email.com.ar") }

	psp_Person := reflect.ValueOf(UpdateCustomer).Elem().FieldByName("Psp_Person")
	if psp_Person.IsValid() { psp_Person.Set(reflect.ValueOf(Person)) }

	psp_Address := reflect.ValueOf(UpdateCustomer).Elem().FieldByName("Psp_Address")
	if psp_Address.IsValid() { psp_Address.Set(reflect.ValueOf(Address)) }

	psp_PaymentMethod := reflect.ValueOf(UpdateCustomer).Elem().FieldByName("Psp_PaymentMethod")
	if psp_PaymentMethod.IsValid() { psp_PaymentMethod.Set(reflect.ValueOf(pm)) }

	resp, err := service.UpdateCustomer(UpdateCustomer)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendDeleteCustomer(service *npsSdk.PaymentServicePlatformPortType) error {

	DeleteCustomer := npsSdk.NewRequerimientoStruct_DeleteCustomer()
	psp_Version := reflect.ValueOf(DeleteCustomer).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(DeleteCustomer).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PosDateTime := reflect.ValueOf(DeleteCustomer).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2017-06-19 12:00:00") }

	psp_CustomerId := reflect.ValueOf(DeleteCustomer).Elem().FieldByName("Psp_CustomerId")
	if psp_CustomerId.IsValid() { psp_CustomerId.SetString("8H9T839wLt83DldqV0qxCMJC3hwvL8yF") }

	resp, err := service.DeleteCustomer(DeleteCustomer)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendQueryCardNumber(service *npsSdk.PaymentServicePlatformPortType) error {

	QueryCardNumber := npsSdk.NewRequerimientoStruct_QueryCardNumber()
	psp_Version := reflect.ValueOf(QueryCardNumber).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(QueryCardNumber).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PosDateTime := reflect.ValueOf(QueryCardNumber).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2017-06-19 12:00:00") }

	psp_QueryCriteria := reflect.ValueOf(QueryCardNumber).Elem().FieldByName("Psp_QueryCriteria")
	if psp_QueryCriteria.IsValid() { psp_QueryCriteria.SetString("T") }

	psp_QueryCriteriaId := reflect.ValueOf(QueryCardNumber).Elem().FieldByName("Psp_QueryCriteriaId")
	if psp_QueryCriteriaId.IsValid() { psp_QueryCriteriaId.SetString("76577") }

	resp, err := service.QueryCardNumber(QueryCardNumber)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendQueryCardDetails(service *npsSdk.PaymentServicePlatformPortType) error {

	QueryCardDetails := npsSdk.NewRequerimientoStruct_QueryCardDetails()
	psp_Version := reflect.ValueOf(QueryCardDetails).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(QueryCardDetails).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PosDateTime := reflect.ValueOf(QueryCardDetails).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2017-06-19 12:00:00") }

	psp_QueryCriteria := reflect.ValueOf(QueryCardDetails).Elem().FieldByName("Psp_QueryCriteria")
	if psp_QueryCriteria.IsValid() { psp_QueryCriteria.SetString("T") }

	psp_QueryCriteriaId := reflect.ValueOf(QueryCardDetails).Elem().FieldByName("Psp_QueryCriteriaId")
	if psp_QueryCriteriaId.IsValid() { psp_QueryCriteriaId.SetString("76577") }

	resp, err := service.QueryCardDetails(QueryCardDetails)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendSimpleQueryTx(service *npsSdk.PaymentServicePlatformPortType) error {

	SimpleQueryTx := npsSdk.NewRequerimientoStruct_SimpleQueryTx()
	psp_Version := reflect.ValueOf(SimpleQueryTx).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(SimpleQueryTx).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PosDateTime := reflect.ValueOf(SimpleQueryTx).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2017-06-19 12:00:00") }

	psp_QueryCriteria := reflect.ValueOf(SimpleQueryTx).Elem().FieldByName("Psp_QueryCriteria")
	if psp_QueryCriteria.IsValid() { psp_QueryCriteria.SetString("T") }

	psp_QueryCriteriaId := reflect.ValueOf(SimpleQueryTx).Elem().FieldByName("Psp_QueryCriteriaId")
	if psp_QueryCriteriaId.IsValid() { psp_QueryCriteriaId.SetString("160117") }

	resp, err := service.SimpleQueryTx(SimpleQueryTx)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())

        ptrTransactions := reflect.ValueOf(resp).Elem().FieldByName("Psp_Transaction")
        respTransactions := ptrTransactions.Elem()
        if respTransactions.IsValid() {

	  psp_AuthorizationCode := respTransactions.FieldByName("Psp_AuthorizationCode")
          if psp_AuthorizationCode.IsValid() {
		fmt.Printf("Psp_AuthorizationCode = [%s]\n", psp_AuthorizationCode.String())
          }
        }

	return nil
}

func SendQueryTxs(service *npsSdk.PaymentServicePlatformPortType) error {

	QueryTxs := npsSdk.NewRequerimientoStruct_QueryTxs()
	psp_Version := reflect.ValueOf(QueryTxs).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(QueryTxs).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PosDateTime := reflect.ValueOf(QueryTxs).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2017-06-19 12:00:00") }

	psp_QueryCriteria := reflect.ValueOf(QueryTxs).Elem().FieldByName("Psp_QueryCriteria")
	if psp_QueryCriteria.IsValid() { psp_QueryCriteria.SetString("T") }

	psp_QueryCriteriaId := reflect.ValueOf(QueryTxs).Elem().FieldByName("Psp_QueryCriteriaId")
	if psp_QueryCriteriaId.IsValid() { psp_QueryCriteriaId.SetString("160117") }

	resp, err := service.QueryTxs(QueryTxs)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendGetInstallmentsOptions(service *npsSdk.PaymentServicePlatformPortType) error {

	GetInstallmentsOptions := npsSdk.NewRequerimientoStruct_GetInstallmentsOptions()
	psp_Version := reflect.ValueOf(GetInstallmentsOptions).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(GetInstallmentsOptions).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PosDateTime := reflect.ValueOf(GetInstallmentsOptions).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2017-06-19 12:00:00") }

	psp_Amount := reflect.ValueOf(GetInstallmentsOptions).Elem().FieldByName("Psp_Amount")
	if psp_Amount.IsValid() { psp_Amount.SetString("1000") }

	psp_Product := reflect.ValueOf(GetInstallmentsOptions).Elem().FieldByName("Psp_Product")
	if psp_Product.IsValid() { psp_Product.SetString("14") }

	psp_Currency := reflect.ValueOf(GetInstallmentsOptions).Elem().FieldByName("Psp_Currency")
	if psp_Currency.IsValid() { psp_Currency.SetString("032") }

	psp_Country := reflect.ValueOf(GetInstallmentsOptions).Elem().FieldByName("Psp_Country")
	if psp_Country.IsValid() { psp_Country.SetString("ARG") }

	psp_ClientSession := reflect.ValueOf(GetInstallmentsOptions).Elem().FieldByName("Psp_ClientSession")
	if psp_ClientSession.IsValid() { psp_ClientSession.SetString("w9mXsBXOlK5mx340F9WsXVoCTUsila2lqTBeqTwezfkbTt0li4NLXrFeFsbfXO9J") }

	resp, err := service.GetInstallmentsOptions(GetInstallmentsOptions)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
	return nil
}

func SendGetIINDetails(service *npsSdk.PaymentServicePlatformPortType) error {

	GetIINDetails := npsSdk.NewRequerimientoStruct_GetIINDetails()
	psp_Version := reflect.ValueOf(GetIINDetails).Elem().FieldByName("Psp_Version")
	if psp_Version.IsValid() { psp_Version.SetString("2.2") }

	psp_MerchantId := reflect.ValueOf(GetIINDetails).Elem().FieldByName("Psp_MerchantId")
	if psp_MerchantId.IsValid() { psp_MerchantId.SetString("psp_test") }

	psp_PosDateTime := reflect.ValueOf(GetIINDetails).Elem().FieldByName("Psp_PosDateTime")
	if psp_PosDateTime.IsValid() { psp_PosDateTime.SetString("2017-06-19 12:00:00") }

	psp_IIN := reflect.ValueOf(GetIINDetails).Elem().FieldByName("Psp_IIN")
	if psp_IIN.IsValid() { psp_IIN.SetString("450799") }

	psp_ClientSession := reflect.ValueOf(GetIINDetails).Elem().FieldByName("Psp_ClientSession")
	if psp_ClientSession.IsValid() { psp_ClientSession.SetString("w9mXsBXOlK5mx340F9WsXVoCTUsila2lqTBeqTwezfkbTt0li4NLXrFeFsbfXO9J") }

	resp, err := service.GetIINDetails(GetIINDetails)
	if err != nil {
		return err
	}

	fmt.Printf("\n")
        Psp_ResponseCod := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseCod")
        Psp_ResponseMsg := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseMsg")
        Psp_ResponseExtended := reflect.ValueOf(resp).Elem().FieldByName("Psp_ResponseExtended")
        fmt.Printf("Response = [%s] [%s]\n", Psp_ResponseCod.String(), Psp_ResponseMsg.String())
        fmt.Printf("Extended = [%s]\n", Psp_ResponseExtended.String())
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

	err := npsSdk.Configure(map[string]interface{}{
		//"environment": CONSTANTS.LOCAL_ENV,
		//"environment": CONSTANTS.STAGING_ENV,
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

	service := npsSdk.NewPaymentServicePlatformPortType(true)

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
	//err = SendUpdateCustomer(service)
	//err = SendDeleteCustomer(service)

	//err = SendQueryCardNumber(service)
	//err = SendQueryCardDetails(service)

	//err = SendSimpleQueryTx(service)
	//err = SendQueryTxs(service)

	err = SendGetInstallmentsOptions(service)
	//err = SendGetIINDetails(service)

	if err != nil {
		fmt.Printf("\n")
		fmt.Printf("ERROR [%s]\n", err)
	}

}
