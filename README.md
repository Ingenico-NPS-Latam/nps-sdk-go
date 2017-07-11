# Go SDK

## Avalilability
Supports GO 1.7 and above

## How to Install

* [Download binary release](https://github.com/Ingenico-NPS-Latam/nps-sdk-go/releases)
* Download and build locally: `go get github.com/Ingenico-NPS-Latam/nps-sdk-go/...`

## Configuration

It's a basic configuration of the SDK

```go
package main

import (
	"fmt"
	"log"
	"nps"
	CONSTANTS "nps/constants"
)

func main() {

  err := nps.Configure(map[string]interface{}{
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
	"nps"
	CONSTANTS "nps/constants"
)

func SendPayOnLine_2p(service *nps.PaymentServicePlatformPortType) error {

  Person := nps.NewPersonStruct()
  Person.FirstName = "First Name"
  Person.LastName = "Last Name"
  Person.PhoneNumber1 = "52520960"
  
  AmountAdditionalDetails := nps.NewAmountAdditionalDetailsRequestStruct()
  AmountAdditionalDetails.Tip = "10"
  AmountAdditionalDetails.Discount = "5"
  
  Billing := nps.NewBillingDetailsStruct()
  Billing.Invoice = "100001234"
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

func main() {

  err := nps.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	  "secret_key":  "_YOUR_SECRET_KEY_",
	})

  if err != nil {
	log.Fatalf("error configuring sdk: %v", err)
  }

  service := nps.NewPaymentServicePlatformPortType(true)

  err = SendPayOnLine_2p(service)

  if err != nil {
    fmt.Printf("\n")
    fmt.Printf("ERROR [%s]\n", err)
  }

}


## Environments

```go
package main

import (
	"fmt"
	"log"
	"nps"
	CONSTANTS "nps/constants"
)

func main() {

  err := nps.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	})
  err := nps.Configure(map[string]interface{}{
	  "environment": CONSTANTS.PRODUCTION_ENV,
	})
  err := nps.Configure(map[string]interface{}{
	  "environment": CONSTANTS.STAGING_ENV,
	})

## Error handling

All exceptions than can occur will be detailed inside of the response provided by NPS .

```go 

  err = SendPayOnLine_2p(service)

  if err != nil {
    fmt.Printf("\n")
    fmt.Printf("ERROR [%s]\n", err)
  }
```

## Advanced configurations

Nps SDK allows you to log what’s happening with you request inside of our SDK, it logs by default to stout.
The SDK uses the custom logger that you use for your project.

An example using your project logger.

```go
package main

import (
	"fmt"
	"log"
	"nps"
	CONSTANTS "nps/constants"
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
 
  err := nps.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	  "secret_key":  "_YOUR_SECRET_KEY_",
	  "debug":       true,
	  "log_level":   CONSTANTS.DEBUG,
	  "npsLog":      appLog,      
	})
```

The "INFO" level will write concise information of the request and will mask sensitive data of the request. 
The "DEBUG" level will write information about the request to let developers debug it in a more detailed way.

```go
package main

import (
	"fmt"
	"log"
	"nps"
	CONSTANTS "nps/constants"
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

  err := nps.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	  "secret_key":  "_YOUR_SECRET_KEY_",
	  "debug":       true,
	  "log_level":   CONSTANTS.INFO,
	  "npsLog":      appLog,      
	})
```

Sanitize allows the SDK to truncate to a fixed size some fields that could make request fail, like extremely long name.

```go
package main

import (
	"fmt"
	"log"
	"nps"
	CONSTANTS "nps/constants"
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

  err := nps.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	  "secret_key":  "_YOUR_SECRET_KEY_",
	  "sanitize":    true,
	})
```
you can change the timeout of the request.

ExecutionTimeout(Default=60 seconds): you can change the execution timeout of the request.

ConnectionTimeout(Default=60 seconds): you can change the connection timeout of the request.

```go
package main

import (
	"fmt"
	"log"
	"nps"
	CONSTANTS "nps/constants"
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

  err := nps.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	  "secret_key":  "_YOUR_SECRET_KEY_",
	  "conn_timeout": 65,
	  "timeout": 65,
	})

```

Proxy configuration

```go
package main

import (
	"fmt"
	"log"
	"nps"
	CONSTANTS "nps/constants"
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

  err := nps.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	  "secret_key":  "_YOUR_SECRET_KEY_",
	  "proxy_url": "http://yourproxy",
	  "timeout": 65,
	})

Configuration::secretKey(“your key here”);
Configuration::proxyUrl("http://yourproxy");
Configuration::proxy_username("proxyUsername");
Configuration::proxy_password("proxyPassword");

```
