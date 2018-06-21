#  Go SDK
 

##  Avalilability
Supports GO 1.7 and above

##  How to Install

* Download and build locally: `go get github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk`

##  Configuration

It's a basic configuration of the SDK

```go
package main

import (
	"fmt"
	"log"
	"github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk"
	CONSTANTS "github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk/constants"
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
	"github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk"
	CONSTANTS "github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk/constants"
)

func SendPayOnLine_2p(service *npsSdk.PaymentServicePlatformPortType) error {

  Person := npsSdk.NewPersonStruct()
  Person.FirstName = "John"
  Person.LastName = "Doe"
  Person.PhoneNumber1 = "+1 011 11111111"
  
  AmountAdditionalDetails := npsSdk.NewAmountAdditionalDetailsRequestStruct()
  AmountAdditionalDetails.Tip = "10"
  AmountAdditionalDetails.Discount = "5"
  
  Billing := npsSdk.NewBillingDetailsStruct()
  Billing.Invoice = "54877555"
  Billing.InvoiceAmount = "15050"
  Billing.InvoiceCurrency = "032"
  Billing.Person = Person
  
  SellerAddress := npsSdk.NewAddressStruct()
  SellerAddress.City = "Miami"
  SellerAddress.Country = "USA"
  SellerAddress.Street = "SellerStreet"
  SellerAddress.HouseNumber = "1245"

  SellerDetails := npsSdk.NewSellerDetailsStruct()
  SellerDetails.Name = "John Doe"
  SellerDetails.Address = SellerAddress
  
  MerchantAdditionalDetails := npsSdk.NewMerchantAdditionalDetailsStruct()
  MerchantAdditionalDetails.ShoppingCartInfo = "ShoppingCartInfo"
  MerchantAdditionalDetails.SellerDetails = SellerDetails
  CustomerAdditionalDetails := npsSdk.NewCustomerAdditionalDetailsStruct()
  CustomerAdditionalDetails.EmailAddress = "jdoe@email.com"
  
  order1 := npsSdk.NewOrderItemStruct()
  order1.Description = "Blue pencil"
  order1.UnitPrice = "10050"

  OrderDetails := npsSdk.NewOrderDetailsStruct()
  OrderDetails.OrderItems = npsSdk.NewArrayOf_OrderItemStruct()
  OrderDetails.OrderItems.Items = make([]*npsSdk.OrderItemStruct, 0)
  OrderDetails.OrderItems.Items = append(OrderDetails.OrderItems.Items, order1)

  payOnline2p := npsSdk.NewRequerimientoStruct_PayOnLine_2p()

  payOnline2p.Psp_Version = "2.2"
  payOnline2p.Psp_MerchantId = "psp_test"
  payOnline2p.Psp_TxSource = "WEB"
  payOnline2p.Psp_MerchTxRef = "ORDERX1466Xz-3"
  payOnline2p.Psp_MerchOrderId = "ORDERX1466Xz"
  payOnline2p.Psp_Amount = "15050"
  payOnline2p.Psp_NumPayments = "1"
  payOnline2p.Psp_Currency = "032"
  payOnline2p.Psp_Country = "ARG"
  payOnline2p.Psp_Product = "14"
  payOnline2p.Psp_CustomerMail = "jdoe@email.com"
  payOnline2p.Psp_CardNumber = "4507990000000010"
  payOnline2p.Psp_CardExpDate = "1912"
  payOnline2p.Psp_CardSecurityCode = "306"
  payOnline2p.Psp_SoftDescriptor = "Compra Art 15"
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
	"github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk"
	CONSTANTS "github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk/constants"
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
	"os"
	"github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk"
	CONSTANTS "github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk/constants"

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
 
  err = npsSdk.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	  "secret_key":  "_YOUR_SECRET_KEY_",
	  "debug":       true,
	  "log_level":   CONSTANTS.DEBUG,
	  "npsLog":      appLog,      
	})
}
```

### Loglevel

The "INFO" level will write concise information of the request and will mask sensitive data of the request. 
The "DEBUG" level will write information about the request to let developers debug it in a more detailed way.

```go
package main

import (
	"fmt"
	"log"
	"os"
	"github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk"
	CONSTANTS "github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk/constants"
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
	"github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk"
	CONSTANTS "github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk/constants"
)

func main() {

  err := npsSdk.Configure(map[string]interface{}{
	  "environment": CONSTANTS.SANDBOX_ENV,
	  "secret_key":  "_YOUR_SECRET_KEY_",
	  "sanitize":    true,
	})
}
```

### Timeout

You can change the timeout of the request.

ExecutionTimeout(Default=10 seconds): you can change the execution timeout of the request.

ConnectionTimeout(Default=60 seconds): you can change the connection timeout of the request.

```go
package main

import (
	"fmt"
	"log"
	"github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk"
	CONSTANTS "github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk/constants"
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

### Proxy configuration

```go
package main

import (
	"fmt"
	"log"
	"github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk"
	CONSTANTS "github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk/constants"
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
