package npsSdk

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	CONSTANTS "github.com/Ingenico-NPS-Latam/nps-sdk-go/src/npsSdk/constants"
	"github.com/Ingenico-NPS-Latam/nps-sdk-go/src/npsSdk/services/production"
	"github.com/Ingenico-NPS-Latam/nps-sdk-go/src/npsSdk/services/sandbox"
	"github.com/Ingenico-NPS-Latam/nps-sdk-go/src/npsSdk/services/staging"
	"os"
	"syscall"
	"time"
)

type PaymentServicePlatformPortType struct {
	client *SOAPClient
}

func NewPaymentServicePlatformPortType(tls bool) *PaymentServicePlatformPortType {

	if len(Configuration.log_file) > 0 {
		f, err := os.OpenFile(Configuration.log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		defer f.Close()
		if err != nil {
			Configuration.npsLog.Fatalf("error opening file: %v", err)
		}

		syscall.Dup2(int(f.Fd()), 1)
		syscall.Dup2(int(f.Fd()), 2)
	}

	client := NewSOAPClient(Configuration.url, tls)

	return &PaymentServicePlatformPortType{
		client: client,
	}
}

func (service *PaymentServicePlatformPortType) SetHeader(header interface{}) {
	service.client.SetHeader(header)
}

//func (service *PaymentServicePlatformPortType) PayOnLine_2p(request *RequerimientoStruct_PayOnLine_2p) (*RespuestaStruct_PayOnLine_2p, error) {
func (service *PaymentServicePlatformPortType) PayOnLine_2p(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_PayOnLine_2p)
           
	  servicioPayOnLine_2p := new(prodservice.ServiceStruct_PayOnLine_2p)
	  servicioPayOnLine_2p.PSP_RequerimientoStruct_PayOnLine_2p = request.(*(prodservice.RequerimientoStruct_PayOnLine_2p))
	  err := service.client.Call(CONSTANTS.PAY_ONLINE_2P, servicioPayOnLine_2p, response)
	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_PayOnLine_2p, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_PayOnLine_2p)

	  servicioPayOnLine_2p := new(stgservice.ServiceStruct_PayOnLine_2p)
	  servicioPayOnLine_2p.PSP_RequerimientoStruct_PayOnLine_2p = request.(*(stgservice.RequerimientoStruct_PayOnLine_2p))
	  err := service.client.Call(CONSTANTS.PAY_ONLINE_2P, servicioPayOnLine_2p, response)
	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_PayOnLine_2p, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_PayOnLine_2p)

	  servicioPayOnLine_2p := new(sndservice.ServiceStruct_PayOnLine_2p)
	  servicioPayOnLine_2p.PSP_RequerimientoStruct_PayOnLine_2p = request.(*(sndservice.RequerimientoStruct_PayOnLine_2p))
	  err := service.client.Call(CONSTANTS.PAY_ONLINE_2P, servicioPayOnLine_2p, response)
	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_PayOnLine_2p, nil
	}

	return nil, nil
}

func (service *PaymentServicePlatformPortType) Authorize_2p(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_Authorize_2p)

 	  servicioAuthorize_2p := new(prodservice.ServiceStruct_Authorize_2p)
	  servicioAuthorize_2p.PSP_RequerimientoStruct_Authorize_2p = request.(*(prodservice.RequerimientoStruct_Authorize_2p))

	  err := service.client.Call(CONSTANTS.AUTHORIZE_2P, servicioAuthorize_2p, response)
	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_Authorize_2p, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_Authorize_2p)

 	  servicioAuthorize_2p := new(stgservice.ServiceStruct_Authorize_2p)
	  servicioAuthorize_2p.PSP_RequerimientoStruct_Authorize_2p = request.(*(stgservice.RequerimientoStruct_Authorize_2p))

	  err := service.client.Call(CONSTANTS.AUTHORIZE_2P, servicioAuthorize_2p, response)
	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_Authorize_2p, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_Authorize_2p)

 	  servicioAuthorize_2p := new(sndservice.ServiceStruct_Authorize_2p)
	  servicioAuthorize_2p.PSP_RequerimientoStruct_Authorize_2p = request.(*(sndservice.RequerimientoStruct_Authorize_2p))

	  err := service.client.Call(CONSTANTS.AUTHORIZE_2P, servicioAuthorize_2p, response)
	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_Authorize_2p, nil

       }

       return nil, nil
}

func (service *PaymentServicePlatformPortType) BankPayment_2p(request interface{}) (interface{}, error) {
//func (service *PaymentServicePlatformPortType) BankPayment_2p(request *RequerimientoStruct_BankPayment_2p) (*RespuestaStruct_BankPayment_2p, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
  	  response := new(prodservice.ServiceRespuestaStruct_BankPayment_2p)

	  servicioBankPayment_2p := new(prodservice.ServiceStruct_BankPayment_2p)
	  servicioBankPayment_2p.PSP_RequerimientoStruct_BankPayment_2p = request.(*(prodservice.RequerimientoStruct_BankPayment_2p))

	  err := service.client.Call(CONSTANTS.BANK_PAYMENT_2P, servicioBankPayment_2p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_BankPayment_2p, nil

	case CONSTANTS.STAGING_ENV:
  	  response := new(stgservice.ServiceRespuestaStruct_BankPayment_2p)

	  servicioBankPayment_2p := new(stgservice.ServiceStruct_BankPayment_2p)
	  servicioBankPayment_2p.PSP_RequerimientoStruct_BankPayment_2p = request.(*(stgservice.RequerimientoStruct_BankPayment_2p))

	  err := service.client.Call(CONSTANTS.BANK_PAYMENT_2P, servicioBankPayment_2p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_BankPayment_2p, nil

	case CONSTANTS.SANDBOX_ENV:
  	  response := new(sndservice.ServiceRespuestaStruct_BankPayment_2p)

	  servicioBankPayment_2p := new(sndservice.ServiceStruct_BankPayment_2p)
	  servicioBankPayment_2p.PSP_RequerimientoStruct_BankPayment_2p = request.(*(sndservice.RequerimientoStruct_BankPayment_2p))

	  err := service.client.Call(CONSTANTS.BANK_PAYMENT_2P, servicioBankPayment_2p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_BankPayment_2p, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) SplitPayOnLine_2p(request *RequerimientoStruct_SplitPayOnLine_2p) (*RespuestaStruct_SplitPayOnLine_2p, error) {
func (service *PaymentServicePlatformPortType) SplitPayOnLine_2p(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_SplitPayOnLine_2p)

	  servicioSplitPayOnLine_2p := new(prodservice.ServiceStruct_SplitPayOnLine_2p)
	  servicioSplitPayOnLine_2p.PSP_RequerimientoStruct_SplitPayOnLine_2p = request.(*(prodservice.RequerimientoStruct_SplitPayOnLine_2p))

	  err := service.client.Call(CONSTANTS.SPLIT_PAY_ONLINE_2P, servicioSplitPayOnLine_2p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_SplitPayOnLine_2p, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_SplitPayOnLine_2p)

	  servicioSplitPayOnLine_2p := new(stgservice.ServiceStruct_SplitPayOnLine_2p)
	  servicioSplitPayOnLine_2p.PSP_RequerimientoStruct_SplitPayOnLine_2p = request.(*(stgservice.RequerimientoStruct_SplitPayOnLine_2p))

	  err := service.client.Call(CONSTANTS.SPLIT_PAY_ONLINE_2P, servicioSplitPayOnLine_2p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_SplitPayOnLine_2p, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_SplitPayOnLine_2p)

	  servicioSplitPayOnLine_2p := new(sndservice.ServiceStruct_SplitPayOnLine_2p)
	  servicioSplitPayOnLine_2p.PSP_RequerimientoStruct_SplitPayOnLine_2p = request.(*(sndservice.RequerimientoStruct_SplitPayOnLine_2p))

	  err := service.client.Call(CONSTANTS.SPLIT_PAY_ONLINE_2P, servicioSplitPayOnLine_2p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_SplitPayOnLine_2p, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) SplitAuthorize_2p(request *RequerimientoStruct_SplitAuthorize_2p) (*RespuestaStruct_SplitAuthorize_2p, error) {
func (service *PaymentServicePlatformPortType) SplitAuthorize_2p(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_SplitAuthorize_2p)

	  servicioSplitAuthorize_2p := new(prodservice.ServiceStruct_SplitAuthorize_2p)
	  servicioSplitAuthorize_2p.PSP_RequerimientoStruct_SplitAuthorize_2p = request.(*(prodservice.RequerimientoStruct_SplitAuthorize_2p))

	  err := service.client.Call(CONSTANTS.SPLIT_AUTHORIZE_2P, servicioSplitAuthorize_2p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_SplitAuthorize_2p, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_SplitAuthorize_2p)

	  servicioSplitAuthorize_2p := new(stgservice.ServiceStruct_SplitAuthorize_2p)
	  servicioSplitAuthorize_2p.PSP_RequerimientoStruct_SplitAuthorize_2p = request.(*(stgservice.RequerimientoStruct_SplitAuthorize_2p))

	  err := service.client.Call(CONSTANTS.SPLIT_AUTHORIZE_2P, servicioSplitAuthorize_2p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_SplitAuthorize_2p, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_SplitAuthorize_2p)

	  servicioSplitAuthorize_2p := new(sndservice.ServiceStruct_SplitAuthorize_2p)
	  servicioSplitAuthorize_2p.PSP_RequerimientoStruct_SplitAuthorize_2p = request.(*(sndservice.RequerimientoStruct_SplitAuthorize_2p))

	  err := service.client.Call(CONSTANTS.SPLIT_AUTHORIZE_2P, servicioSplitAuthorize_2p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_SplitAuthorize_2p, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) PayOnLine_3p(request *RequerimientoStruct_PayOnLine_3p) (*RespuestaStruct_PayOnLine_3p, error) {
func (service *PaymentServicePlatformPortType) PayOnLine_3p(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_PayOnLine_3p)

	  servicioPayOnLine_3p := new(prodservice.ServiceStruct_PayOnLine_3p)
	  servicioPayOnLine_3p.PSP_RequerimientoStruct_PayOnLine_3p = request.(*(prodservice.RequerimientoStruct_PayOnLine_3p))

	  err := service.client.Call(CONSTANTS.PAY_ONLINE_3P, servicioPayOnLine_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_PayOnLine_3p, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_PayOnLine_3p)

	  servicioPayOnLine_3p := new(stgservice.ServiceStruct_PayOnLine_3p)
	  servicioPayOnLine_3p.PSP_RequerimientoStruct_PayOnLine_3p = request.(*(stgservice.RequerimientoStruct_PayOnLine_3p))

	  err := service.client.Call(CONSTANTS.PAY_ONLINE_3P, servicioPayOnLine_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_PayOnLine_3p, nil
	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_PayOnLine_3p)

	  servicioPayOnLine_3p := new(sndservice.ServiceStruct_PayOnLine_3p)
	  servicioPayOnLine_3p.PSP_RequerimientoStruct_PayOnLine_3p = request.(*(sndservice.RequerimientoStruct_PayOnLine_3p))

	  err := service.client.Call(CONSTANTS.PAY_ONLINE_3P, servicioPayOnLine_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_PayOnLine_3p, nil
        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) SplitPayOnLine_3p(request *RequerimientoStruct_SplitPayOnLine_3p) (*RespuestaStruct_SplitPayOnLine_3p, error) {
func (service *PaymentServicePlatformPortType) SplitPayOnLine_3p(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_SplitPayOnLine_3p)

	  servicioSplitPayOnLine_3p := new(prodservice.ServiceStruct_SplitPayOnLine_3p)
	  servicioSplitPayOnLine_3p.PSP_RequerimientoStruct_SplitPayOnLine_3p = request.(*(prodservice.RequerimientoStruct_SplitPayOnLine_3p))

	  err := service.client.Call(CONSTANTS.SPLIT_PAY_ONLINE_3P, servicioSplitPayOnLine_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_SplitPayOnLine_3p, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_SplitPayOnLine_3p)

	  servicioSplitPayOnLine_3p := new(stgservice.ServiceStruct_SplitPayOnLine_3p)
	  servicioSplitPayOnLine_3p.PSP_RequerimientoStruct_SplitPayOnLine_3p = request.(*(stgservice.RequerimientoStruct_SplitPayOnLine_3p))

	  err := service.client.Call(CONSTANTS.SPLIT_PAY_ONLINE_3P, servicioSplitPayOnLine_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_SplitPayOnLine_3p, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_SplitPayOnLine_3p)

	  servicioSplitPayOnLine_3p := new(sndservice.ServiceStruct_SplitPayOnLine_3p)
	  servicioSplitPayOnLine_3p.PSP_RequerimientoStruct_SplitPayOnLine_3p = request.(*(sndservice.RequerimientoStruct_SplitPayOnLine_3p))

	  err := service.client.Call(CONSTANTS.SPLIT_PAY_ONLINE_3P, servicioSplitPayOnLine_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_SplitPayOnLine_3p, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) Authorize_3p(request *RequerimientoStruct_Authorize_3p) (*RespuestaStruct_Authorize_3p, error) {
func (service *PaymentServicePlatformPortType) Authorize_3p(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_Authorize_3p)

	  servicioAuthorize_3p := new(prodservice.ServiceStruct_Authorize_3p)
	  servicioAuthorize_3p.PSP_RequerimientoStruct_Authorize_3p = request.(*(prodservice.RequerimientoStruct_Authorize_3p))

	  err := service.client.Call(CONSTANTS.AUTHORIZE_3P, servicioAuthorize_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_Authorize_3p, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_Authorize_3p)

	  servicioAuthorize_3p := new(stgservice.ServiceStruct_Authorize_3p)
	  servicioAuthorize_3p.PSP_RequerimientoStruct_Authorize_3p = request.(*(stgservice.RequerimientoStruct_Authorize_3p))

	  err := service.client.Call(CONSTANTS.AUTHORIZE_3P, servicioAuthorize_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_Authorize_3p, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_Authorize_3p)

	  servicioAuthorize_3p := new(sndservice.ServiceStruct_Authorize_3p)
	  servicioAuthorize_3p.PSP_RequerimientoStruct_Authorize_3p = request.(*(sndservice.RequerimientoStruct_Authorize_3p))

	  err := service.client.Call(CONSTANTS.AUTHORIZE_3P, servicioAuthorize_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_Authorize_3p, nil


        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) SplitAuthorize_3p(request *RequerimientoStruct_SplitAuthorize_3p) (*RespuestaStruct_SplitAuthorize_3p, error) {
func (service *PaymentServicePlatformPortType) SplitAuthorize_3p(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_SplitAuthorize_3p)

	  serviceSplitAuthorize_3p := new(prodservice.ServiceStruct_SplitAuthorize_3p)
	  serviceSplitAuthorize_3p.PSP_RequerimientoStruct_SplitAuthorize_3p = request.(*(prodservice.RequerimientoStruct_SplitAuthorize_3p))

	  err := service.client.Call(CONSTANTS.SPLIT_AUTHORIZE_3P, serviceSplitAuthorize_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_SplitAuthorize_3p, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_SplitAuthorize_3p)

	  serviceSplitAuthorize_3p := new(stgservice.ServiceStruct_SplitAuthorize_3p)
	  serviceSplitAuthorize_3p.PSP_RequerimientoStruct_SplitAuthorize_3p = request.(*(stgservice.RequerimientoStruct_SplitAuthorize_3p))

	  err := service.client.Call(CONSTANTS.SPLIT_AUTHORIZE_3P, serviceSplitAuthorize_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_SplitAuthorize_3p, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_SplitAuthorize_3p)

	  serviceSplitAuthorize_3p := new(sndservice.ServiceStruct_SplitAuthorize_3p)
	  serviceSplitAuthorize_3p.PSP_RequerimientoStruct_SplitAuthorize_3p = request.(*(sndservice.RequerimientoStruct_SplitAuthorize_3p))

	  err := service.client.Call(CONSTANTS.SPLIT_AUTHORIZE_3P, serviceSplitAuthorize_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_SplitAuthorize_3p, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) BankPayment_3p(request *RequerimientoStruct_BankPayment_3p) (*RespuestaStruct_BankPayment_3p, error) {
func (service *PaymentServicePlatformPortType) BankPayment_3p(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_BankPayment_3p)

	  servicioBankPayment_3p := new(prodservice.ServiceStruct_BankPayment_3p)
	  servicioBankPayment_3p.PSP_RequerimientoStruct_BankPayment_3p = request.(*(prodservice.RequerimientoStruct_BankPayment_3p))

	  err := service.client.Call(CONSTANTS.BANK_PAYMENT_3P, servicioBankPayment_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_BankPayment_3p, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_BankPayment_3p)

	  servicioBankPayment_3p := new(stgservice.ServiceStruct_BankPayment_3p)
	  servicioBankPayment_3p.PSP_RequerimientoStruct_BankPayment_3p = request.(*(stgservice.RequerimientoStruct_BankPayment_3p))

	  err := service.client.Call(CONSTANTS.BANK_PAYMENT_3P, servicioBankPayment_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_BankPayment_3p, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_BankPayment_3p)

	  servicioBankPayment_3p := new(sndservice.ServiceStruct_BankPayment_3p)
	  servicioBankPayment_3p.PSP_RequerimientoStruct_BankPayment_3p = request.(*(sndservice.RequerimientoStruct_BankPayment_3p))

	  err := service.client.Call(CONSTANTS.BANK_PAYMENT_3P, servicioBankPayment_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_BankPayment_3p, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) CashPayment_3p(request *RequerimientoStruct_CashPayment_3p) (*RespuestaStruct_CashPayment_3p, error) {
func (service *PaymentServicePlatformPortType) CashPayment_3p(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_CashPayment_3p)

	  servicioCashPayment_3p := new(prodservice.ServiceStruct_CashPayment_3p)
	  servicioCashPayment_3p.PSP_RequerimientoStruct_CashPayment_3p = request.(*(prodservice.RequerimientoStruct_CashPayment_3p))

	  err := service.client.Call(CONSTANTS.CASH_PAYMENT_3P, servicioCashPayment_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CashPayment_3p, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_CashPayment_3p)

	  servicioCashPayment_3p := new(stgservice.ServiceStruct_CashPayment_3p)
	  servicioCashPayment_3p.PSP_RequerimientoStruct_CashPayment_3p = request.(*(stgservice.RequerimientoStruct_CashPayment_3p))

	  err := service.client.Call(CONSTANTS.CASH_PAYMENT_3P, servicioCashPayment_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CashPayment_3p, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_CashPayment_3p)

	  servicioCashPayment_3p := new(sndservice.ServiceStruct_CashPayment_3p)
	  servicioCashPayment_3p.PSP_RequerimientoStruct_CashPayment_3p = request.(*(sndservice.RequerimientoStruct_CashPayment_3p))

	  err := service.client.Call(CONSTANTS.CASH_PAYMENT_3P, servicioCashPayment_3p, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CashPayment_3p, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) Capture(request *RequerimientoStruct_Capture) (*RespuestaStruct_Capture, error) {
func (service *PaymentServicePlatformPortType) Capture(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_Capture)

	  servicioCapture := new(prodservice.ServiceStruct_Capture)
	  servicioCapture.PSP_RequerimientoStruct_Capture = request.(*(prodservice.RequerimientoStruct_Capture))

	  err := service.client.Call(CONSTANTS.CAPTURE, servicioCapture, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_Capture, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_Capture)

	  servicioCapture := new(stgservice.ServiceStruct_Capture)
	  servicioCapture.PSP_RequerimientoStruct_Capture = request.(*(stgservice.RequerimientoStruct_Capture))

	  err := service.client.Call(CONSTANTS.CAPTURE, servicioCapture, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_Capture, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_Capture)

	  servicioCapture := new(sndservice.ServiceStruct_Capture)
	  servicioCapture.PSP_RequerimientoStruct_Capture = request.(*(sndservice.RequerimientoStruct_Capture))

	  err := service.client.Call(CONSTANTS.CAPTURE, servicioCapture, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_Capture, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) Refund(request *RequerimientoStruct_Refund) (*RespuestaStruct_Refund, error) {
func (service *PaymentServicePlatformPortType) Refund(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_Refund)

	  servicioRefund := new(prodservice.ServiceStruct_Refund)
	  servicioRefund.PSP_RequerimientoStruct_Refund = request.(*(prodservice.RequerimientoStruct_Refund))

	  err := service.client.Call(CONSTANTS.REFUND, servicioRefund, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_Refund, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_Refund)

	  servicioRefund := new(stgservice.ServiceStruct_Refund)
	  servicioRefund.PSP_RequerimientoStruct_Refund = request.(*(stgservice.RequerimientoStruct_Refund))

	  err := service.client.Call(CONSTANTS.REFUND, servicioRefund, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_Refund, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_Refund)

	  servicioRefund := new(sndservice.ServiceStruct_Refund)
	  servicioRefund.PSP_RequerimientoStruct_Refund = request.(*(sndservice.RequerimientoStruct_Refund))

	  err := service.client.Call(CONSTANTS.REFUND, servicioRefund, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_Refund, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) GetInstallmentsOptions(request *RequerimientoStruct_GetInstallmentsOptions) (*RespuestaStruct_GetInstallmentsOptions, error) {
func (service *PaymentServicePlatformPortType) GetInstallmentsOptions(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_GetInstallmentsOptions)

	  servicioGetInstallmentsOptions := new(prodservice.ServiceStruct_GetInstallmentsOptions)
	  servicioGetInstallmentsOptions.PSP_RequerimientoStruct_GetInstallmentsOptions = request.(*(prodservice.RequerimientoStruct_GetInstallmentsOptions))

	  err := service.client.Call(CONSTANTS.GET_INSTALLMENTS_OPTIONS, servicioGetInstallmentsOptions, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_GetInstallmentsOptions, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_GetInstallmentsOptions)

	  servicioGetInstallmentsOptions := new(stgservice.ServiceStruct_GetInstallmentsOptions)
	  servicioGetInstallmentsOptions.PSP_RequerimientoStruct_GetInstallmentsOptions = request.(*(stgservice.RequerimientoStruct_GetInstallmentsOptions))

	  err := service.client.Call(CONSTANTS.GET_INSTALLMENTS_OPTIONS, servicioGetInstallmentsOptions, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_GetInstallmentsOptions, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_GetInstallmentsOptions)

	  servicioGetInstallmentsOptions := new(sndservice.ServiceStruct_GetInstallmentsOptions)
	  servicioGetInstallmentsOptions.PSP_RequerimientoStruct_GetInstallmentsOptions = request.(*(sndservice.RequerimientoStruct_GetInstallmentsOptions))

	  err := service.client.Call(CONSTANTS.GET_INSTALLMENTS_OPTIONS, servicioGetInstallmentsOptions, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_GetInstallmentsOptions, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) CreatePaymentMethodToken(request *RequerimientoStruct_CreatePaymentMethodToken) (*RespuestaStruct_CreatePaymentMethodToken, error) {
func (service *PaymentServicePlatformPortType) CreatePaymentMethodToken(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_CreatePaymentMethodToken)

	  if request.(*(prodservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_CardInputDetails != nil {
		request.(*(prodservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_CardInputDetails.XMLName = xml.Name{"", "psp_CardInputDetails"}
	  }
	  if request.(*(prodservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_Address != nil {
		request.(*(prodservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if request.(*(prodservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_Person != nil {
		request.(*(prodservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }

	  servicioCreatePaymentMethodToken := new(prodservice.ServiceStruct_CreatePaymentMethodToken)
	  servicioCreatePaymentMethodToken.PSP_RequerimientoStruct_CreatePaymentMethodToken = request.(*(prodservice.RequerimientoStruct_CreatePaymentMethodToken))

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD_TOKEN, servicioCreatePaymentMethodToken, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CreatePaymentMethodToken, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_CreatePaymentMethodToken)

	  if request.(*(stgservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_CardInputDetails != nil {
		request.(*(stgservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_CardInputDetails.XMLName = xml.Name{"", "psp_CardInputDetails"}
	  }
	  if request.(*(stgservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_Address != nil {
		request.(*(stgservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if request.(*(stgservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_Person != nil {
		request.(*(stgservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }

	  servicioCreatePaymentMethodToken := new(stgservice.ServiceStruct_CreatePaymentMethodToken)
	  servicioCreatePaymentMethodToken.PSP_RequerimientoStruct_CreatePaymentMethodToken = request.(*(stgservice.RequerimientoStruct_CreatePaymentMethodToken))

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD_TOKEN, servicioCreatePaymentMethodToken, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CreatePaymentMethodToken, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_CreatePaymentMethodToken)

	  if request.(*(sndservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_CardInputDetails != nil {
		request.(*(sndservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_CardInputDetails.XMLName = xml.Name{"", "psp_CardInputDetails"}
	  }
	  if request.(*(sndservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_Address != nil {
		request.(*(sndservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if request.(*(sndservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_Person != nil {
		request.(*(sndservice.RequerimientoStruct_CreatePaymentMethodToken)).Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }

	  servicioCreatePaymentMethodToken := new(sndservice.ServiceStruct_CreatePaymentMethodToken)
	  servicioCreatePaymentMethodToken.PSP_RequerimientoStruct_CreatePaymentMethodToken = request.(*(sndservice.RequerimientoStruct_CreatePaymentMethodToken))

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD_TOKEN, servicioCreatePaymentMethodToken, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CreatePaymentMethodToken, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) RetrievePaymentMethodToken(request *RequerimientoStruct_RetrievePaymentMethodToken) (*RespuestaStruct_RetrievePaymentMethodToken, error) {
func (service *PaymentServicePlatformPortType) RetrievePaymentMethodToken(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_RetrievePaymentMethodToken)

	  servicioRetrievePaymentMethodToken := new(prodservice.ServiceStruct_RetrievePaymentMethodToken)
	  servicioRetrievePaymentMethodToken.PSP_RequerimientoStruct_RetrievePaymentMethodToken = request.(*(prodservice.RequerimientoStruct_RetrievePaymentMethodToken))

	  err := service.client.Call(CONSTANTS.RETRIEVE_PAYMENT_METHOD_TOKEN, servicioRetrievePaymentMethodToken, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_RetrievePaymentMethodToken, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_RetrievePaymentMethodToken)

	  servicioRetrievePaymentMethodToken := new(stgservice.ServiceStruct_RetrievePaymentMethodToken)
	  servicioRetrievePaymentMethodToken.PSP_RequerimientoStruct_RetrievePaymentMethodToken = request.(*(stgservice.RequerimientoStruct_RetrievePaymentMethodToken))

	  err := service.client.Call(CONSTANTS.RETRIEVE_PAYMENT_METHOD_TOKEN, servicioRetrievePaymentMethodToken, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_RetrievePaymentMethodToken, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_RetrievePaymentMethodToken)

	  servicioRetrievePaymentMethodToken := new(sndservice.ServiceStruct_RetrievePaymentMethodToken)
	  servicioRetrievePaymentMethodToken.PSP_RequerimientoStruct_RetrievePaymentMethodToken = request.(*(sndservice.RequerimientoStruct_RetrievePaymentMethodToken))

	  err := service.client.Call(CONSTANTS.RETRIEVE_PAYMENT_METHOD_TOKEN, servicioRetrievePaymentMethodToken, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_RetrievePaymentMethodToken, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) RecachePaymentMethodToken(request *RequerimientoStruct_RecachePaymentMethodToken) (*RespuestaStruct_RecachePaymentMethodToken, error) {
func (service *PaymentServicePlatformPortType) RecachePaymentMethodToken(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_RecachePaymentMethodToken)

	  if request.(*(prodservice.RequerimientoStruct_RecachePaymentMethodToken)).Psp_Address != nil {
		request.(*(prodservice.RequerimientoStruct_RecachePaymentMethodToken)).Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if request.(*(prodservice.RequerimientoStruct_RecachePaymentMethodToken)).Psp_Person != nil {
		request.(*(prodservice.RequerimientoStruct_RecachePaymentMethodToken)).Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }

	  servicioRecachePaymentMethodToken := new(prodservice.ServiceStruct_RecachePaymentMethodToken)
	  servicioRecachePaymentMethodToken.PSP_RequerimientoStruct_RecachePaymentMethodToken = request.(*(prodservice.RequerimientoStruct_RecachePaymentMethodToken))

	  err := service.client.Call(CONSTANTS.RECACHE_PAYMENT_METHOD_TOKEN, servicioRecachePaymentMethodToken, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_RecachePaymentMethodToken, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_RecachePaymentMethodToken)

	  if request.(*(stgservice.RequerimientoStruct_RecachePaymentMethodToken)).Psp_Address != nil {
		request.(*(stgservice.RequerimientoStruct_RecachePaymentMethodToken)).Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if request.(*(stgservice.RequerimientoStruct_RecachePaymentMethodToken)).Psp_Person != nil {
		request.(*(stgservice.RequerimientoStruct_RecachePaymentMethodToken)).Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }

	  servicioRecachePaymentMethodToken := new(stgservice.ServiceStruct_RecachePaymentMethodToken)
	  servicioRecachePaymentMethodToken.PSP_RequerimientoStruct_RecachePaymentMethodToken = request.(*(stgservice.RequerimientoStruct_RecachePaymentMethodToken))

	  err := service.client.Call(CONSTANTS.RECACHE_PAYMENT_METHOD_TOKEN, servicioRecachePaymentMethodToken, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_RecachePaymentMethodToken, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_RecachePaymentMethodToken)

	  if request.(*(sndservice.RequerimientoStruct_RecachePaymentMethodToken)).Psp_Address != nil {
		request.(*(sndservice.RequerimientoStruct_RecachePaymentMethodToken)).Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if request.(*(sndservice.RequerimientoStruct_RecachePaymentMethodToken)).Psp_Person != nil {
		request.(*(sndservice.RequerimientoStruct_RecachePaymentMethodToken)).Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }

	  servicioRecachePaymentMethodToken := new(sndservice.ServiceStruct_RecachePaymentMethodToken)
	  servicioRecachePaymentMethodToken.PSP_RequerimientoStruct_RecachePaymentMethodToken = request.(*(sndservice.RequerimientoStruct_RecachePaymentMethodToken))

	  err := service.client.Call(CONSTANTS.RECACHE_PAYMENT_METHOD_TOKEN, servicioRecachePaymentMethodToken, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_RecachePaymentMethodToken, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) CreateClientSession(request *RequerimientoStruct_CreateClientSession) (*RespuestaStruct_CreateClientSession, error) {
func (service *PaymentServicePlatformPortType) CreateClientSession(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_CreateClientSession)

	  servicioCreateClientSession := new(prodservice.ServiceStruct_CreateClientSession)
	  servicioCreateClientSession.PSP_RequerimientoStruct_CreateClientSession = request.(*(prodservice.RequerimientoStruct_CreateClientSession))

	  err := service.client.Call(CONSTANTS.CREATE_CLIENT_SESSION, servicioCreateClientSession, response)
	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CreateClientSession, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_CreateClientSession)

	  servicioCreateClientSession := new(stgservice.ServiceStruct_CreateClientSession)
	  servicioCreateClientSession.PSP_RequerimientoStruct_CreateClientSession = request.(*(stgservice.RequerimientoStruct_CreateClientSession))

	  err := service.client.Call(CONSTANTS.CREATE_CLIENT_SESSION, servicioCreateClientSession, response)
	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CreateClientSession, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_CreateClientSession)

	  servicioCreateClientSession := new(sndservice.ServiceStruct_CreateClientSession)
	  servicioCreateClientSession.PSP_RequerimientoStruct_CreateClientSession = request.(*(sndservice.RequerimientoStruct_CreateClientSession))

	  err := service.client.Call(CONSTANTS.CREATE_CLIENT_SESSION, servicioCreateClientSession, response)
	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CreateClientSession, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) CreatePaymentMethod(request *RequerimientoStruct_CreatePaymentMethod) (*RespuestaStruct_CreatePaymentMethod, error) {
func (service *PaymentServicePlatformPortType) CreatePaymentMethod(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_CreatePaymentMethod)

	  servicioCreatePaymentMethod := new(prodservice.ServiceStruct_CreatePaymentMethod)
	  servicioCreatePaymentMethod.PSP_RequerimientoStruct_CreatePaymentMethod = request.(*(prodservice.RequerimientoStruct_CreatePaymentMethod))

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD, servicioCreatePaymentMethod, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CreatePaymentMethod, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_CreatePaymentMethod)

	  servicioCreatePaymentMethod := new(stgservice.ServiceStruct_CreatePaymentMethod)
	  servicioCreatePaymentMethod.PSP_RequerimientoStruct_CreatePaymentMethod = request.(*(stgservice.RequerimientoStruct_CreatePaymentMethod))

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD, servicioCreatePaymentMethod, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CreatePaymentMethod, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_CreatePaymentMethod)

	  servicioCreatePaymentMethod := new(sndservice.ServiceStruct_CreatePaymentMethod)
	  servicioCreatePaymentMethod.PSP_RequerimientoStruct_CreatePaymentMethod = request.(*(sndservice.RequerimientoStruct_CreatePaymentMethod))

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD, servicioCreatePaymentMethod, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CreatePaymentMethod, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) CreatePaymentMethodFromPayment(request *RequerimientoStruct_CreatePaymentMethodFromPayment) (*RespuestaStruct_CreatePaymentMethodFromPayment, error) {
func (service *PaymentServicePlatformPortType) CreatePaymentMethodFromPayment(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_CreatePaymentMethodFromPayment)

	  servicioCreatePaymentMethodFromPayment := new(prodservice.ServiceStruct_CreatePaymentMethodFromPayment)
	  servicioCreatePaymentMethodFromPayment.PSP_RequerimientoStruct_CreatePaymentMethodFromPayment = request.(*(prodservice.RequerimientoStruct_CreatePaymentMethodFromPayment))

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD_FROM_PAYMENT, servicioCreatePaymentMethodFromPayment, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CreatePaymentMethodFromPayment, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_CreatePaymentMethodFromPayment)

	  servicioCreatePaymentMethodFromPayment := new(stgservice.ServiceStruct_CreatePaymentMethodFromPayment)
	  servicioCreatePaymentMethodFromPayment.PSP_RequerimientoStruct_CreatePaymentMethodFromPayment = request.(*(stgservice.RequerimientoStruct_CreatePaymentMethodFromPayment))

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD_FROM_PAYMENT, servicioCreatePaymentMethodFromPayment, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CreatePaymentMethodFromPayment, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_CreatePaymentMethodFromPayment)

	  servicioCreatePaymentMethodFromPayment := new(sndservice.ServiceStruct_CreatePaymentMethodFromPayment)
	  servicioCreatePaymentMethodFromPayment.PSP_RequerimientoStruct_CreatePaymentMethodFromPayment = request.(*(sndservice.RequerimientoStruct_CreatePaymentMethodFromPayment))

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD_FROM_PAYMENT, servicioCreatePaymentMethodFromPayment, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CreatePaymentMethodFromPayment, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) RetrievePaymentMethod(request *RequerimientoStruct_RetrievePaymentMethod) (*RespuestaStruct_RetrievePaymentMethod, error) {
func (service *PaymentServicePlatformPortType) RetrievePaymentMethod(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_RetrievePaymentMethod)

	  servicioRetrievePaymentMethod := new(prodservice.ServiceStruct_RetrievePaymentMethod)
	  servicioRetrievePaymentMethod.PSP_RequerimientoStruct_RetrievePaymentMethod = request.(*(prodservice.RequerimientoStruct_RetrievePaymentMethod))

	  err := service.client.Call(CONSTANTS.RETRIEVE_PAYMENT_METHOD, servicioRetrievePaymentMethod, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_RetrievePaymentMethod, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_RetrievePaymentMethod)

	  servicioRetrievePaymentMethod := new(stgservice.ServiceStruct_RetrievePaymentMethod)
	  servicioRetrievePaymentMethod.PSP_RequerimientoStruct_RetrievePaymentMethod = request.(*(stgservice.RequerimientoStruct_RetrievePaymentMethod))

	  err := service.client.Call(CONSTANTS.RETRIEVE_PAYMENT_METHOD, servicioRetrievePaymentMethod, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_RetrievePaymentMethod, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_RetrievePaymentMethod)

	  servicioRetrievePaymentMethod := new(sndservice.ServiceStruct_RetrievePaymentMethod)
	  servicioRetrievePaymentMethod.PSP_RequerimientoStruct_RetrievePaymentMethod = request.(*(sndservice.RequerimientoStruct_RetrievePaymentMethod))

	  err := service.client.Call(CONSTANTS.RETRIEVE_PAYMENT_METHOD, servicioRetrievePaymentMethod, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_RetrievePaymentMethod, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) UpdatePaymentMethod(request *RequerimientoStruct_UpdatePaymentMethod) (*RespuestaStruct_UpdatePaymentMethod, error) {
func (service *PaymentServicePlatformPortType) UpdatePaymentMethod(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_UpdatePaymentMethod)

	  servicioUpdatePaymentMethod := new(prodservice.ServiceStruct_UpdatePaymentMethod)
	  servicioUpdatePaymentMethod.PSP_RequerimientoStruct_UpdatePaymentMethod = request.(*(prodservice.RequerimientoStruct_UpdatePaymentMethod))

	  err := service.client.Call(CONSTANTS.UPDATE_PAYMENT_METHOD, servicioUpdatePaymentMethod, response)
	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_UpdatePaymentMethod, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_UpdatePaymentMethod)

	  servicioUpdatePaymentMethod := new(stgservice.ServiceStruct_UpdatePaymentMethod)
	  servicioUpdatePaymentMethod.PSP_RequerimientoStruct_UpdatePaymentMethod = request.(*(stgservice.RequerimientoStruct_UpdatePaymentMethod))

	  err := service.client.Call(CONSTANTS.UPDATE_PAYMENT_METHOD, servicioUpdatePaymentMethod, response)
	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_UpdatePaymentMethod, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_UpdatePaymentMethod)

	  servicioUpdatePaymentMethod := new(sndservice.ServiceStruct_UpdatePaymentMethod)
	  servicioUpdatePaymentMethod.PSP_RequerimientoStruct_UpdatePaymentMethod = request.(*(sndservice.RequerimientoStruct_UpdatePaymentMethod))

	  err := service.client.Call(CONSTANTS.UPDATE_PAYMENT_METHOD, servicioUpdatePaymentMethod, response)
	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_UpdatePaymentMethod, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) DeletePaymentMethod(request *RequerimientoStruct_DeletePaymentMethod) (*RespuestaStruct_DeletePaymentMethod, error) {
func (service *PaymentServicePlatformPortType) DeletePaymentMethod(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_DeletePaymentMethod)

	  servicioDeletePaymentMethod := new(prodservice.ServiceStruct_DeletePaymentMethod)
	  servicioDeletePaymentMethod.PSP_RequerimientoStruct_DeletePaymentMethod = request.(*(prodservice.RequerimientoStruct_DeletePaymentMethod))

	  err := service.client.Call(CONSTANTS.DELETE_PAYMENT_METHOD, servicioDeletePaymentMethod, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_DeletePaymentMethod, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_DeletePaymentMethod)

	  servicioDeletePaymentMethod := new(stgservice.ServiceStruct_DeletePaymentMethod)
	  servicioDeletePaymentMethod.PSP_RequerimientoStruct_DeletePaymentMethod = request.(*(stgservice.RequerimientoStruct_DeletePaymentMethod))

	  err := service.client.Call(CONSTANTS.DELETE_PAYMENT_METHOD, servicioDeletePaymentMethod, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_DeletePaymentMethod, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_DeletePaymentMethod)

	  servicioDeletePaymentMethod := new(sndservice.ServiceStruct_DeletePaymentMethod)
	  servicioDeletePaymentMethod.PSP_RequerimientoStruct_DeletePaymentMethod = request.(*(sndservice.RequerimientoStruct_DeletePaymentMethod))

	  err := service.client.Call(CONSTANTS.DELETE_PAYMENT_METHOD, servicioDeletePaymentMethod, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_DeletePaymentMethod, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) CreateCustomer(request *RequerimientoStruct_CreateCustomer) (*RespuestaStruct_CreateCustomer, error) {
func (service *PaymentServicePlatformPortType) CreateCustomer(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_CreateCustomer)

	  if request.(*(prodservice.RequerimientoStruct_CreateCustomer)).Psp_Address != nil {
		request.(*(prodservice.RequerimientoStruct_CreateCustomer)).Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if request.(*(prodservice.RequerimientoStruct_CreateCustomer)).Psp_Person != nil {
		request.(*(prodservice.RequerimientoStruct_CreateCustomer)).Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }
	  if request.(*(prodservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod != nil && 
	     request.(*(prodservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod.Person != nil {
		request.(*(prodservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod.Person.XMLName = xml.Name{"", "Person"}
	  }
	  if request.(*(prodservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod != nil && 
	     request.(*(prodservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod.Address != nil {
		request.(*(prodservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod.Address.XMLName = xml.Name{"", "Address"}
	  }

	  servicioCreateCustomer := new(prodservice.ServiceStruct_CreateCustomer)
	  servicioCreateCustomer.PSP_RequerimientoStruct_CreateCustomer = request.(*(prodservice.RequerimientoStruct_CreateCustomer))

	  err := service.client.Call(CONSTANTS.CREATE_CUSTOMER, servicioCreateCustomer, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CreateCustomer, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_CreateCustomer)

	  if request.(*(stgservice.RequerimientoStruct_CreateCustomer)).Psp_Address != nil {
		request.(*(stgservice.RequerimientoStruct_CreateCustomer)).Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if request.(*(stgservice.RequerimientoStruct_CreateCustomer)).Psp_Person != nil {
		request.(*(stgservice.RequerimientoStruct_CreateCustomer)).Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }
	  if request.(*(stgservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod != nil && 
             request.(*(stgservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod.Person != nil {
		request.(*(stgservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod.Person.XMLName = xml.Name{"", "Person"}
	  }
	  if request.(*(stgservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod != nil && 
             request.(*(stgservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod.Address != nil {
		request.(*(stgservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod.Address.XMLName = xml.Name{"", "Address"}
	  }

	  servicioCreateCustomer := new(stgservice.ServiceStruct_CreateCustomer)
	  servicioCreateCustomer.PSP_RequerimientoStruct_CreateCustomer = request.(*(stgservice.RequerimientoStruct_CreateCustomer))

	  err := service.client.Call(CONSTANTS.CREATE_CUSTOMER, servicioCreateCustomer, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CreateCustomer, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_CreateCustomer)

	  if request.(*(sndservice.RequerimientoStruct_CreateCustomer)).Psp_Address != nil {
		request.(*(sndservice.RequerimientoStruct_CreateCustomer)).Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if request.(*(sndservice.RequerimientoStruct_CreateCustomer)).Psp_Person != nil {
		request.(*(sndservice.RequerimientoStruct_CreateCustomer)).Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }
	  if request.(*(sndservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod != nil && 
             request.(*(sndservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod.Person != nil {
		request.(*(sndservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod.Person.XMLName = xml.Name{"", "Person"}
	  }
	  if request.(*(sndservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod != nil && 
             request.(*(sndservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod.Address != nil {
		request.(*(sndservice.RequerimientoStruct_CreateCustomer)).Psp_PaymentMethod.Address.XMLName = xml.Name{"", "Address"}
	  }

	  servicioCreateCustomer := new(sndservice.ServiceStruct_CreateCustomer)
	  servicioCreateCustomer.PSP_RequerimientoStruct_CreateCustomer = request.(*(sndservice.RequerimientoStruct_CreateCustomer))

	  err := service.client.Call(CONSTANTS.CREATE_CUSTOMER, servicioCreateCustomer, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_CreateCustomer, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) RetrieveCustomer(request *RequerimientoStruct_RetrieveCustomer) (*RespuestaStruct_RetrieveCustomer, error) {
func (service *PaymentServicePlatformPortType) RetrieveCustomer(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_RetrieveCustomer)

	  servicioRetrieveCustomer := new(prodservice.ServiceStruct_RetrieveCustomer)
	  servicioRetrieveCustomer.PSP_RequerimientoStruct_RetrieveCustomer = request.(*(prodservice.RequerimientoStruct_RetrieveCustomer))

	  err := service.client.Call(CONSTANTS.RETRIEVE_CUSTOMER, servicioRetrieveCustomer, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_RetrieveCustomer, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_RetrieveCustomer)

	  servicioRetrieveCustomer := new(stgservice.ServiceStruct_RetrieveCustomer)
	  servicioRetrieveCustomer.PSP_RequerimientoStruct_RetrieveCustomer = request.(*(stgservice.RequerimientoStruct_RetrieveCustomer))

	  err := service.client.Call(CONSTANTS.RETRIEVE_CUSTOMER, servicioRetrieveCustomer, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_RetrieveCustomer, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_RetrieveCustomer)

	  servicioRetrieveCustomer := new(sndservice.ServiceStruct_RetrieveCustomer)
	  servicioRetrieveCustomer.PSP_RequerimientoStruct_RetrieveCustomer = request.(*(sndservice.RequerimientoStruct_RetrieveCustomer))

	  err := service.client.Call(CONSTANTS.RETRIEVE_CUSTOMER, servicioRetrieveCustomer, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_RetrieveCustomer, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) UpdateCustomer(request *RequerimientoStruct_UpdateCustomer) (*RespuestaStruct_UpdateCustomer, error) {
func (service *PaymentServicePlatformPortType) UpdateCustomer(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_UpdateCustomer)

	  if request.(*(prodservice.RequerimientoStruct_UpdateCustomer)).Psp_Address != nil {
		request.(*(prodservice.RequerimientoStruct_UpdateCustomer)).Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if request.(*(prodservice.RequerimientoStruct_UpdateCustomer)).Psp_Person != nil {
		request.(*(prodservice.RequerimientoStruct_UpdateCustomer)).Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }
	  if request.(*(prodservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod != nil && 
             request.(*(prodservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod.Person != nil {
		request.(*(prodservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod.Person.XMLName = xml.Name{"", "Person"}
	  }
	  if request.(*(prodservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod != nil && 
             request.(*(prodservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod.Address != nil {
		request.(*(prodservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod.Address.XMLName = xml.Name{"", "Address"}
	  }

	  servicioUpdateCustomer := new(prodservice.ServiceStruct_UpdateCustomer)
	  servicioUpdateCustomer.PSP_RequerimientoStruct_UpdateCustomer = request.(*(prodservice.RequerimientoStruct_UpdateCustomer))

	  err := service.client.Call(CONSTANTS.UPDATE_CUSTOMER, servicioUpdateCustomer, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_UpdateCustomer, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_UpdateCustomer)

	  if request.(*(stgservice.RequerimientoStruct_UpdateCustomer)).Psp_Address != nil {
		request.(*(stgservice.RequerimientoStruct_UpdateCustomer)).Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if request.(*(stgservice.RequerimientoStruct_UpdateCustomer)).Psp_Person != nil {
		request.(*(stgservice.RequerimientoStruct_UpdateCustomer)).Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }
	  if request.(*(stgservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod != nil && 
             request.(*(stgservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod.Person != nil {
		request.(*(stgservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod.Person.XMLName = xml.Name{"", "Person"}
	  }
	  if request.(*(stgservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod != nil && 
             request.(*(stgservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod.Address != nil {
		request.(*(stgservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod.Address.XMLName = xml.Name{"", "Address"}
	  }

	  servicioUpdateCustomer := new(stgservice.ServiceStruct_UpdateCustomer)
	  servicioUpdateCustomer.PSP_RequerimientoStruct_UpdateCustomer = request.(*(stgservice.RequerimientoStruct_UpdateCustomer))

	  err := service.client.Call(CONSTANTS.UPDATE_CUSTOMER, servicioUpdateCustomer, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_UpdateCustomer, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_UpdateCustomer)

	  if request.(*(sndservice.RequerimientoStruct_UpdateCustomer)).Psp_Address != nil {
		request.(*(sndservice.RequerimientoStruct_UpdateCustomer)).Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if request.(*(sndservice.RequerimientoStruct_UpdateCustomer)).Psp_Person != nil {
		request.(*(sndservice.RequerimientoStruct_UpdateCustomer)).Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }
	  if request.(*(sndservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod != nil && 
             request.(*(sndservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod.Person != nil {
		request.(*(sndservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod.Person.XMLName = xml.Name{"", "Person"}
	  }
	  if request.(*(sndservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod != nil && 
             request.(*(sndservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod.Address != nil {
		request.(*(prodservice.RequerimientoStruct_UpdateCustomer)).Psp_PaymentMethod.Address.XMLName = xml.Name{"", "Address"}
	  }

	  servicioUpdateCustomer := new(sndservice.ServiceStruct_UpdateCustomer)
	  servicioUpdateCustomer.PSP_RequerimientoStruct_UpdateCustomer = request.(*(sndservice.RequerimientoStruct_UpdateCustomer))

	  err := service.client.Call(CONSTANTS.UPDATE_CUSTOMER, servicioUpdateCustomer, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_UpdateCustomer, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) DeleteCustomer(request *RequerimientoStruct_DeleteCustomer) (*RespuestaStruct_DeleteCustomer, error) {
func (service *PaymentServicePlatformPortType) DeleteCustomer(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_DeleteCustomer)

	  servicioDeleteCustomer := new(prodservice.ServiceStruct_DeleteCustomer)
	  servicioDeleteCustomer.PSP_RequerimientoStruct_DeleteCustomer = request.(*(prodservice.RequerimientoStruct_DeleteCustomer))

	  err := service.client.Call(CONSTANTS.DELETE_CUSTOMER, servicioDeleteCustomer, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_DeleteCustomer, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_DeleteCustomer)

	  servicioDeleteCustomer := new(stgservice.ServiceStruct_DeleteCustomer)
	  servicioDeleteCustomer.PSP_RequerimientoStruct_DeleteCustomer = request.(*(stgservice.RequerimientoStruct_DeleteCustomer))

	  err := service.client.Call(CONSTANTS.DELETE_CUSTOMER, servicioDeleteCustomer, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_DeleteCustomer, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_DeleteCustomer)

	  servicioDeleteCustomer := new(sndservice.ServiceStruct_DeleteCustomer)
	  servicioDeleteCustomer.PSP_RequerimientoStruct_DeleteCustomer = request.(*(sndservice.RequerimientoStruct_DeleteCustomer))

	  err := service.client.Call(CONSTANTS.DELETE_CUSTOMER, servicioDeleteCustomer, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_DeleteCustomer, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) SimpleQueryTx(request *RequerimientoStruct_SimpleQueryTx) (*RespuestaStruct_SimpleQueryTx, error) {
func (service *PaymentServicePlatformPortType) SimpleQueryTx(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_SimpleQueryTx)

	  servicioSimpleQueryTx := new(prodservice.ServiceStruct_SimpleQueryTx)
	  servicioSimpleQueryTx.PSP_RequerimientoStruct_SimpleQueryTx = request.(*(prodservice.RequerimientoStruct_SimpleQueryTx))

	  err := service.client.Call(CONSTANTS.SIMPLE_QUERY_TX, servicioSimpleQueryTx, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_SimpleQueryTx, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_SimpleQueryTx)

	  servicioSimpleQueryTx := new(stgservice.ServiceStruct_SimpleQueryTx)
	  servicioSimpleQueryTx.PSP_RequerimientoStruct_SimpleQueryTx = request.(*(stgservice.RequerimientoStruct_SimpleQueryTx))

	  err := service.client.Call(CONSTANTS.SIMPLE_QUERY_TX, servicioSimpleQueryTx, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_SimpleQueryTx, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_SimpleQueryTx)

	  servicioSimpleQueryTx := new(sndservice.ServiceStruct_SimpleQueryTx)
	  servicioSimpleQueryTx.PSP_RequerimientoStruct_SimpleQueryTx = request.(*(sndservice.RequerimientoStruct_SimpleQueryTx))

	  err := service.client.Call(CONSTANTS.SIMPLE_QUERY_TX, servicioSimpleQueryTx, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_SimpleQueryTx, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) QueryCardNumber(request *RequerimientoStruct_QueryCardNumber) (*RespuestaStruct_QueryCardNumber, error) {
func (service *PaymentServicePlatformPortType) QueryCardNumber(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_QueryCardNumber)

	  servicioQueryCardNumber := new(prodservice.ServiceStruct_QueryCardNumber)
	  servicioQueryCardNumber.PSP_RequerimientoStruct_QueryCardNumber = request.(*(prodservice.RequerimientoStruct_QueryCardNumber))

	  err := service.client.Call(CONSTANTS.QUERY_CARD_NUMBER, servicioQueryCardNumber, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_QueryCardNumber, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_QueryCardNumber)

	  servicioQueryCardNumber := new(stgservice.ServiceStruct_QueryCardNumber)
	  servicioQueryCardNumber.PSP_RequerimientoStruct_QueryCardNumber = request.(*(stgservice.RequerimientoStruct_QueryCardNumber))

	  err := service.client.Call(CONSTANTS.QUERY_CARD_NUMBER, servicioQueryCardNumber, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_QueryCardNumber, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_QueryCardNumber)

	  servicioQueryCardNumber := new(sndservice.ServiceStruct_QueryCardNumber)
	  servicioQueryCardNumber.PSP_RequerimientoStruct_QueryCardNumber = request.(*(sndservice.RequerimientoStruct_QueryCardNumber))

	  err := service.client.Call(CONSTANTS.QUERY_CARD_NUMBER, servicioQueryCardNumber, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_QueryCardNumber, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) QueryCardDetails(request *RequerimientoStruct_QueryCardDetails) (*RespuestaStruct_QueryCardDetails, error) {
func (service *PaymentServicePlatformPortType) QueryCardDetails(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_QueryCardDetails)

	  servicioQueryCardDetails := new(prodservice.ServiceStruct_QueryCardDetails)
	  servicioQueryCardDetails.PSP_RequerimientoStruct_QueryCardDetails = request.(*(prodservice.RequerimientoStruct_QueryCardDetails))

	  err := service.client.Call(CONSTANTS.QUERY_CARD_DETAILS, servicioQueryCardDetails, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_QueryCardDetails, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_QueryCardDetails)

	  servicioQueryCardDetails := new(stgservice.ServiceStruct_QueryCardDetails)
	  servicioQueryCardDetails.PSP_RequerimientoStruct_QueryCardDetails = request.(*(stgservice.RequerimientoStruct_QueryCardDetails))

	  err := service.client.Call(CONSTANTS.QUERY_CARD_DETAILS, servicioQueryCardDetails, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_QueryCardDetails, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_QueryCardDetails)

	  servicioQueryCardDetails := new(sndservice.ServiceStruct_QueryCardDetails)
	  servicioQueryCardDetails.PSP_RequerimientoStruct_QueryCardDetails = request.(*(sndservice.RequerimientoStruct_QueryCardDetails))

	  err := service.client.Call(CONSTANTS.QUERY_CARD_DETAILS, servicioQueryCardDetails, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_QueryCardDetails, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) QueryTxs(request *RequerimientoStruct_QueryTxs) (*RespuestaStruct_QueryTxs, error) {
func (service *PaymentServicePlatformPortType) QueryTxs(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_QueryTxs)

	  servicioQueryTxs := new(prodservice.ServiceStruct_QueryTxs)
	  servicioQueryTxs.PSP_RequerimientoStruct_QueryTxs = request.(*(prodservice.RequerimientoStruct_QueryTxs))

	  err := service.client.Call(CONSTANTS.QUERY_TXS, servicioQueryTxs, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_QueryTxs, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_QueryTxs)

	  servicioQueryTxs := new(stgservice.ServiceStruct_QueryTxs)
	  servicioQueryTxs.PSP_RequerimientoStruct_QueryTxs = request.(*(stgservice.RequerimientoStruct_QueryTxs))

	  err := service.client.Call(CONSTANTS.QUERY_TXS, servicioQueryTxs, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_QueryTxs, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_QueryTxs)

	  servicioQueryTxs := new(sndservice.ServiceStruct_QueryTxs)
	  servicioQueryTxs.PSP_RequerimientoStruct_QueryTxs = request.(*(sndservice.RequerimientoStruct_QueryTxs))

	  err := service.client.Call(CONSTANTS.QUERY_TXS, servicioQueryTxs, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_QueryTxs, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) GetIINDetails(request *RequerimientoStruct_GetIINDetails) (*RespuestaStruct_GetIINDetails, error) {
func (service *PaymentServicePlatformPortType) GetIINDetails(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_GetIINDetails)

	  servicioGetIINDetails := new(prodservice.ServiceStruct_GetIINDetails)
	  servicioGetIINDetails.PSP_RequerimientoStruct_GetIINDetails = request.(*(prodservice.RequerimientoStruct_GetIINDetails))

	  err := service.client.Call(CONSTANTS.GET_IIN_DETAILS, servicioGetIINDetails, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_GetIINDetails, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_GetIINDetails)

	  servicioGetIINDetails := new(stgservice.ServiceStruct_GetIINDetails)
	  servicioGetIINDetails.PSP_RequerimientoStruct_GetIINDetails = request.(*(stgservice.RequerimientoStruct_GetIINDetails))

	  err := service.client.Call(CONSTANTS.GET_IIN_DETAILS, servicioGetIINDetails, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_GetIINDetails, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_GetIINDetails)

	  servicioGetIINDetails := new(sndservice.ServiceStruct_GetIINDetails)
	  servicioGetIINDetails.PSP_RequerimientoStruct_GetIINDetails = request.(*(sndservice.RequerimientoStruct_GetIINDetails))

	  err := service.client.Call(CONSTANTS.GET_IIN_DETAILS, servicioGetIINDetails, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_GetIINDetails, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) ChangeSecretKey(request *RequerimientoStruct_ChangeSecretKey) (*RespuestaStruct_ChangeSecretKey, error) {
func (service *PaymentServicePlatformPortType) ChangeSecretKey(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_ChangeSecretKey)

	  servicioChangeSecretKey := new(prodservice.ServiceStruct_ChangeSecretKey)
	  servicioChangeSecretKey.PSP_RequerimientoStruct_ChangeSecretKey = request.(*(prodservice.RequerimientoStruct_ChangeSecretKey))

	  err := service.client.Call(CONSTANTS.CHANGE_SECRET_KEY, servicioChangeSecretKey, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_ChangeSecretKey, nil
	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_ChangeSecretKey)

	  servicioChangeSecretKey := new(stgservice.ServiceStruct_ChangeSecretKey)
	  servicioChangeSecretKey.PSP_RequerimientoStruct_ChangeSecretKey = request.(*(stgservice.RequerimientoStruct_ChangeSecretKey))

	  err := service.client.Call(CONSTANTS.CHANGE_SECRET_KEY, servicioChangeSecretKey, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_ChangeSecretKey, nil
	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_ChangeSecretKey)

	  servicioChangeSecretKey := new(sndservice.ServiceStruct_ChangeSecretKey)
	  servicioChangeSecretKey.PSP_RequerimientoStruct_ChangeSecretKey = request.(*(sndservice.RequerimientoStruct_ChangeSecretKey))

	  err := service.client.Call(CONSTANTS.CHANGE_SECRET_KEY, servicioChangeSecretKey, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_ChangeSecretKey, nil
        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) FraudScreening(request *RequerimientoStruct_FraudScreening) (*RespuestaStruct_FraudScreening, error) {
func (service *PaymentServicePlatformPortType) FraudScreening(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_FraudScreening)

	  servicioFraudScreening := new(prodservice.ServiceStruct_FraudScreening)
	  servicioFraudScreening.PSP_RequerimientoStruct_FraudScreening = request.(*(prodservice.RequerimientoStruct_FraudScreening))

	  err := service.client.Call(CONSTANTS.FRAUD_SCREENING, servicioFraudScreening, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_FraudScreening, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_FraudScreening)

	  servicioFraudScreening := new(stgservice.ServiceStruct_FraudScreening)
	  servicioFraudScreening.PSP_RequerimientoStruct_FraudScreening = request.(*(stgservice.RequerimientoStruct_FraudScreening))

	  err := service.client.Call(CONSTANTS.FRAUD_SCREENING, servicioFraudScreening, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_FraudScreening, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_FraudScreening)

	  servicioFraudScreening := new(sndservice.ServiceStruct_FraudScreening)
	  servicioFraudScreening.PSP_RequerimientoStruct_FraudScreening = request.(*(sndservice.RequerimientoStruct_FraudScreening))

	  err := service.client.Call(CONSTANTS.FRAUD_SCREENING, servicioFraudScreening, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_FraudScreening, nil

        }
        
        return nil, nil
}

//func (service *PaymentServicePlatformPortType) NotifyFraudScreeningReview(request *RequerimientoStruct_NotifyFraudScreeningReview) (*RespuestaStruct_NotifyFraudScreeningReview, error) {
func (service *PaymentServicePlatformPortType) NotifyFraudScreeningReview(request interface{}) (interface{}, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  response := new(prodservice.ServiceRespuestaStruct_NotifyFraudScreeningReview)

	  servicioNotifyFraudScreeningReview := new(prodservice.ServiceStruct_NotifyFraudScreeningReview)
	  servicioNotifyFraudScreeningReview.PSP_RequerimientoStruct_NotifyFraudScreeningReview = request.(*(prodservice.RequerimientoStruct_NotifyFraudScreeningReview))

	  err := service.client.Call(CONSTANTS.NOTIFY_FRAUD_SCREENING_REVIEW, servicioNotifyFraudScreeningReview, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_NotifyFraudScreeningReview, nil

	case CONSTANTS.STAGING_ENV:
	  response := new(stgservice.ServiceRespuestaStruct_NotifyFraudScreeningReview)

	  servicioNotifyFraudScreeningReview := new(stgservice.ServiceStruct_NotifyFraudScreeningReview)
	  servicioNotifyFraudScreeningReview.PSP_RequerimientoStruct_NotifyFraudScreeningReview = request.(*(stgservice.RequerimientoStruct_NotifyFraudScreeningReview))

	  err := service.client.Call(CONSTANTS.NOTIFY_FRAUD_SCREENING_REVIEW, servicioNotifyFraudScreeningReview, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_NotifyFraudScreeningReview, nil

	case CONSTANTS.SANDBOX_ENV:
	  response := new(sndservice.ServiceRespuestaStruct_NotifyFraudScreeningReview)

	  servicioNotifyFraudScreeningReview := new(sndservice.ServiceStruct_NotifyFraudScreeningReview)
	  servicioNotifyFraudScreeningReview.PSP_RequerimientoStruct_NotifyFraudScreeningReview = request.(*(sndservice.RequerimientoStruct_NotifyFraudScreeningReview))

	  err := service.client.Call(CONSTANTS.NOTIFY_FRAUD_SCREENING_REVIEW, servicioNotifyFraudScreeningReview, response)

	  if err != nil {
		return nil, err
	  }

	  return response.PSP_RespuestaStruct_NotifyFraudScreeningReview, nil

        }
        
        return nil, nil
}

var timeout = time.Duration(time.Duration(Configuration.connection_Timeout))

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, time.Duration(time.Duration(Configuration.connection_Timeout)*time.Second))
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url    string
	tls    bool
	auth   *BasicAuth
	header interface{}
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {

		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool) *SOAPClient {
	return &SOAPClient{
		url: url,
		tls: tls,
	}
}

func (s *SOAPClient) SetHeader(header interface{}) {
	s.header = header
}

func Log(log_level string, format string, args ...interface{}) {
	if Configuration.debug {
		if Configuration.logLevel == CONSTANTS.DEBUG {
			Configuration.npsLog.Printf(log_level+" "+format, args...)
		} else {
			if Configuration.logLevel == CONSTANTS.INFO && log_level != CONSTANTS.DEBUG {
				Configuration.npsLog.Printf("%s %s", log_level, MaskData(format, args...))
			} else {
				if log_level == CONSTANTS.ERROR {
					Configuration.npsLog.Printf("%s %s", log_level, MaskData(format, args...))
				}
			}
		}
	}
}

func formatXML(data []byte) ([]byte, error) {
	b := &bytes.Buffer{}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	encoder := xml.NewEncoder(b)
	encoder.Indent(" ", "    ")
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			encoder.Flush()
			return b.Bytes(), nil
		}
		if err != nil {
			return nil, err
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			return nil, err
		}
	}
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {

	if Configuration.logLevel == CONSTANTS.DEBUG && Configuration.environment == CONSTANTS.PRODUCTION_ENV {
		Log(CONSTANTS.ERROR, "DEBUG level is not allowed in PRODUCTION ENVIRONMENT")
		err := errors.New("DEBUG level is not allowed in PRODUCTION ENVIRONMENT")
		return err
	}

	envelope := SOAPEnvelope{}

	if s.header != nil {
		envelope.Header = &SOAPHeader{Header: s.header}
	}

	if GetMerchDetNotAddServices[soapAction] != true {
		AddExtraInf(request)
	}

	if Configuration.sanitize {
		Sanitize(request, "", 0)
	}

	if err := AddSecureHashIfNoClientSession(request); err != nil {
		return err
	}

	envelope.Body.Content = request

	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	Log(CONSTANTS.DEBUG, " Buffer ")
	Log(CONSTANTS.INFO, "%s", buffer.String())
	Log(CONSTANTS.DEBUG, " url: [%s]", Configuration.url)

	req, err := http.NewRequest("POST", Configuration.url, buffer)

	if err != nil {
		return err
	}

	if soapAction != "" {
		var soapAux bytes.Buffer
		soapAux.WriteString(Configuration.url)
		soapAux.WriteString("/")
		soapAux.WriteString(soapAction)

		req.Header.Add("SOAPAction", soapAux.String())
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")

	//req.Header.Set("User-Agent", "gowsdl/0.1")
	//req.Close = true
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	if Configuration.proxy_url != "" {
		proxyUrl, err := url.Parse(Configuration.proxy_url)
		if err != nil {
			return err
		} else {

			if Configuration.proxy_username != "" {
				if Configuration.proxy_password != "" {
					req.SetBasicAuth(Configuration.proxy_username, Configuration.proxy_password)
				}
			}
			tr.Proxy = http.ProxyURL(proxyUrl)
		}
	}
	Log(CONSTANTS.DEBUG, "proxy config done")

	client := &http.Client{Transport: tr,
		Timeout: time.Duration(time.Duration(Configuration.execution_Timeout) * time.Second)}

	Log(CONSTANTS.DEBUG, "client set")
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Log(CONSTANTS.ERROR, err.Error())
		return err
	}

	if len(rawbody) == 0 {
		Log(CONSTANTS.ERROR, "empty response")
		return nil
	}

	/*	Log(CONSTANTS.DEBUG, "========= Response =========")
		Log(CONSTANTS.DEBUG, string(rawbody))
		Log(CONSTANTS.DEBUG, "============================")
	*/
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}

	err = xml.Unmarshal(rawbody, respEnvelope)

	Log(CONSTANTS.DEBUG, "========= XML Formatted =========")
	buf, _ := formatXML(rawbody)
	Log(CONSTANTS.DEBUG, string(buf))
	Log(CONSTANTS.DEBUG, "============================")

	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
