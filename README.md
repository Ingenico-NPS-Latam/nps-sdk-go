#  Go SDK

##  Avalilability
Supports GO 1.7 and above

##  How to Install

* [Download binary release](https://github.com/Ingenico-NPS-Latam/nps-sdk-go/releases)
* Download and build locally: `go get github.com/Ingenico-NPS-Latam/nps-sdk-go/src/npsSdk`

##  Configuration

It's a basic configuration of the SDK

```go
package main

import (
	"fmt"
	"log"
	"npsSdk"
	CONSTANTS "npsSdk/constants"
	"reflect"
	"time"
)

func main() {

  err := npsSdk.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	  "secret_key":  "_YOUR_SECRET_KEY_",
	})

  if err != nil {
	log.Fatalf("error configuring sdk: %v", err)
  }
}

```

Here is an simple example request:

```go
package main

import (
	"fmt"
	"log"
	"npsSdk"
	CONSTANTS "npsSdk/constants"
	"reflect"
	"time"
)

func SendPayOnLine_2p(service *npsSdk.PaymentServicePlatformPortType) error {
	Person := npsSdk.NewPersonStruct()

	fName := reflect.ValueOf(Person).Elem().FieldByName("FirstName")
	if fName.IsValid() { fName.SetString("Silvina") }

	lName := reflect.ValueOf(Person).Elem().FieldByName("LastName")
	if lName.IsValid() { lName.SetString("Falconi") }

	phNumber1 := reflect.ValueOf(Person).Elem().FieldByName("PhoneNumber1")
	if phNumber1.IsValid() { phNumber1.SetString("52520960") }

	AmountAdditionalDetails := npsSdk.NewAmountAdditionalDetailsRequestStruct()
	tip := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Tip")
	if tip.IsValid() { tip.SetString("10") }

	discount := reflect.ValueOf(AmountAdditionalDetails).Elem().FieldByName("Discount")
	if discount.IsValid() { discount.SetString("5") }

	Billing := npsSdk.NewBillingDetailsStruct()
	Invoice := reflect.ValueOf(Billing).Elem().FieldByName("Invoice")
	if Invoice.IsValid() { Invoice.SetString("100001234") }

	InvoiceAmount := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceAmount")
	if InvoiceAmount.IsValid() { InvoiceAmount.SetString("990") }

	InvoiceCurrency := reflect.ValueOf(Billing).Elem().FieldByName("InvoiceCurrency")
	if InvoiceCurrency.IsValid() { InvoiceCurrency.SetString("032") }

	BillingPerson := reflect.ValueOf(Billing).Elem().FieldByName("Person")
	if BillingPerson.IsValid() { BillingPerson.Set(reflect.ValueOf(Person)) }

	SellerAddress := npsSdk.NewAddressStruct()
	City := reflect.ValueOf(SellerAddress).Elem().FieldByName("City")
	if City.IsValid() { City.SetString("CABA") }

	Country := reflect.ValueOf(SellerAddress).Elem().FieldByName("Country")
	if Country.IsValid() { Country.SetString("ARG") }

	Street := reflect.ValueOf(SellerAddress).Elem().FieldByName("Street")
	if Street.IsValid() { Street.SetString("SellerStreet") }

	HouseNumber := reflect.ValueOf(SellerAddress).Elem().FieldByName("HouseNumber")
	if HouseNumber.IsValid() { HouseNumber.SetString("1234") }

	SellerDetails := npsSdk.NewSellerDetailsStruct()
	SellerDetailsName := reflect.ValueOf(SellerDetails).Elem().FieldByName("Name")
	if SellerDetailsName.IsValid() { SellerDetailsName.SetString("Seller Name") }

	SellerDetailsAddress := reflect.ValueOf(SellerDetails).Elem().FieldByName("Address")
	if SellerDetailsAddress.IsValid() { SellerDetailsAddress.Set(reflect.ValueOf(SellerAddress)) }

	MerchantAdditionalDetails := npsSdk.NewMerchantAdditionalDetailsStruct()
	ShoppingCartInfo := reflect.ValueOf(MerchantAdditionalDetails).Elem().FieldByName("ShoppingCartInfo")
	if ShoppingCartInfo.IsValid() { ShoppingCartInfo.SetString("ShoppingCartInfo") }

	MerchantAdditionalDetailsSellerDetails := reflect.ValueOf(MerchantAdditionalDetails).Elem().FieldByName("SellerDetails")
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
	}

	OrderDetails := npsSdk.NewOrderDetailsStruct()    
	OrderDetailsItems := reflect.ValueOf(OrderDetails).Elem().FieldByName("OrderItems")
	if OrderDetailsItems.IsValid() { OrderDetailsItems.Set(reflect.ValueOf(OrderItems)) }

	payOnline2p := npsSdk.NewRequerimientoStruct_PayOnLine_2p()
	Psp_Version := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_Version")
	if Psp_Version.IsValid() { Psp_Version.SetString("2.2")}

	Psp_MerchantId := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_MerchantId")
	if Psp_MerchantId.IsValid() { Psp_MerchantId.SetString("psp_test")}

	Psp_TxSource := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_TxSource")
	if Psp_TxSource.IsValid() { Psp_TxSource.SetString("WEB")}

	Psp_MerchTxRef := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_MerchTxRef")
        t := time.Now()
 	if Psp_MerchTxRef.IsValid() { Psp_MerchTxRef.SetString(t.Format("20060102150405")) }

	Psp_MerchOrderId := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_MerchOrderId")
	if Psp_MerchOrderId.IsValid() { Psp_MerchOrderId.SetString("ORDER56666")}

	Psp_Amount := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_Amount")
	if Psp_Amount.IsValid() { Psp_Amount.SetString("1000")}

	Psp_NumPayments := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_NumPayments")
	if Psp_NumPayments.IsValid() { Psp_NumPayments.SetString("1")}

	Psp_Currency := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_Currency")
	if Psp_Currency.IsValid() { Psp_Currency.SetString("032")}

	Psp_Country := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_Country")
	if Psp_Country.IsValid() { Psp_Country.SetString("ARG")}

	Psp_Product := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_Product")
	if Psp_Product.IsValid() { Psp_Product.SetString("14")}

	Psp_CustomerMail := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_CustomerMail")
	if Psp_CustomerMail.IsValid() { Psp_CustomerMail.SetString("yourmail@gmail")}

	Psp_CardNumber := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_CardNumber")
	if Psp_CardNumber.IsValid() { Psp_CardNumber.SetString("4507990000000010")}

	Psp_CardExpDate := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_CardExpDate")
	if Psp_CardExpDate.IsValid() { Psp_CardExpDate.SetString("1903")}

	Psp_CardSecurityCode := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_CardSecurityCode")
	if Psp_CardSecurityCode.IsValid() { Psp_CardSecurityCode.SetString("306")}

	Psp_SoftDescriptor := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_SoftDescriptor")
	if Psp_SoftDescriptor.IsValid() { Psp_SoftDescriptor.SetString("Sol Tropical E")}

	Psp_PosDateTime := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_PosDateTime")
	if Psp_PosDateTime.IsValid() { Psp_PosDateTime.SetString("2016-12-01 12:00:00")}

        Psp_UserId := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_UserId")
        if Psp_UserId.IsValid() { Psp_UserId.SetString("SFALCONI")} 

	Psp_OrderDetails := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_OrderDetails")
	if Psp_OrderDetails.IsValid() {  Psp_OrderDetails.Set(reflect.ValueOf(OrderDetails)) }

	Psp_CustomerAdditionalDetails := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_CustomerAdditionalDetails")
	if Psp_CustomerAdditionalDetails.IsValid() {  Psp_CustomerAdditionalDetails.Set(reflect.ValueOf(CustomerAdditionalDetails)) }

	Psp_AmountAdditionalDetails := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_AmountAdditionalDetails")
	if Psp_AmountAdditionalDetails.IsValid() {  Psp_AmountAdditionalDetails.Set(reflect.ValueOf(AmountAdditionalDetails)) }

	Psp_BillingDetails := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_BillingDetails")
	if Psp_BillingDetails.IsValid() {  Psp_BillingDetails.Set(reflect.ValueOf(Billing)) }

	Psp_MerchantAdditionalDetails := reflect.ValueOf(payOnline2p).Elem().FieldByName("Psp_MerchantAdditionalDetails")
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

func main() {

  err := npsSdk.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	  "secret_key":  "_YOUR_SECRET_KEY_",
	})

  if err != nil {
	log.Fatalf("error configuring sdk: %v", err)
  }

  service := npsSdk.NewPaymentServicePlatformPortType(true)

  err = SendPayOnLine_2p(service)

  if err != nil {
    fmt.Printf("\n")
    fmt.Printf("ERROR [%s]\n", err)
  }

}

```

##  Environments

```go
package main

import (
	"fmt"
	"log"
	"npsSdk"
	CONSTANTS "npsSdk/constants"
	"reflect"
	"time"
)

func main() {

  err := npsSdk.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	})
  err := npsSdk.Configure(map[string]interface{}{
	  "environment": CONSTANTS.PRODUCTION_ENV,
	})
  err := npsSdk.Configure(map[string]interface{}{
	  "environment": CONSTANTS.STAGING_ENV,
	})
}

```

##  Error handling

All exceptions than can occur will be detailed inside of the response provided by NPS .

```go 

  err = SendPayOnLine_2p(service)

  if err != nil {
    fmt.Printf("\n")
    fmt.Printf("ERROR [%s]\n", err)
  }
```

##  Advanced configurations

### Logging
Nps SDK allows you to log what’s happening with you request inside of our SDK, it logs by default to stout.
The SDK uses the custom logger that you use for your project.

An example using your project logger.

```go
package main

import (
	"fmt"
	"log"
	"npsSdk"
	CONSTANTS "npsSdk/constants"
	"reflect"
	"time"
)

func main() {

  var appLog = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
  f, err := os.OpenFile("appLog.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  defer f.Close()
  if err != nil {
    log.Fatalf("error opening file: %v", err)
  }
  appLog.SetOutput(f)
  appLog.Println("MAIN begin")
 
  err := npsSdk.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	  "secret_key":  "_YOUR_SECRET_KEY_",
	  "debug":       true,
	  "log_level":   CONSTANTS.DEBUG,
	  "npsLog":      appLog,      
	})
}
```

### LogLevel

The "INFO" level will write concise information of the request and will mask sensitive data of the request. 
The "DEBUG" level will write information about the request to let developers debug it in a more detailed way.

```go
package main

import (
	"fmt"
	"log"
	"npsSdk"
	CONSTANTS "npsSdk/constants"
	"reflect"
	"time"
)

func main() {

  err := npsSdk.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	  "secret_key":  "_YOUR_SECRET_KEY_",
	  "debug":       true,
	  "log_level":   CONSTANTS.INFO,
	})
}
```

### Sanitize

Sanitize allows the SDK to truncate to a fixed size some fields that could make request fail, like extremely long name.

```go
package main

import (
	"fmt"
	"log"
	"npsSdk"
	CONSTANTS "npsSdk/constants"
	"reflect"
	"time"
)

func main() {

  err := npsSdk.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	  "secret_key":  "_YOUR_SECRET_KEY_",
	  "sanitize":    true,
	})
}
```

## Timeout

you can change the timeout of the request.

ExecutionTimeout(Default=60 seconds): you can change the execution timeout of the request.

ConnectionTimeout(Default=60 seconds): you can change the connection timeout of the request.

```go
package main

import (
	"fmt"
	"log"
	"npsSdk"
	CONSTANTS "npsSdk/constants"
	"reflect"
	"time"
)

func main() {

  err := npsSdk.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	  "secret_key":  "_YOUR_SECRET_KEY_",
	  "conn_timeout": 65,
	  "timeout": 65,
	})
}

```

### Proxy

Proxy configuration

```go
package main

import (
	"fmt"
	"log"
	"npsSdk"
	CONSTANTS "npsSdk/constants"
	"reflect"
	"time"
)

func main() {

  err := npsSdk.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	  "secret_key":  "_YOUR_SECRET_KEY_",
	  "proxy_url": "http://yourproxy",
	  "proxy_username": "proxyUsername",
	  "proxy_password": "proxyPassword",
	})
}

```
