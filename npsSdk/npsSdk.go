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
	CONSTANTS "github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk/constants"
	"github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk/services/production"
	"github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk/services/sandbox"
	"github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk/services/staging"
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

func (service *PaymentServicePlatformPortType) PayOnLine_2p(request *RequerimientoStruct_PayOnLine_2p) (*RespuestaStruct_PayOnLine_2p, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_PayOnLine_2p)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_PayOnLine_2p)           
	  servicioPayOnLine_2p := new(prodservice.ServiceStruct_PayOnLine_2p)
	  servicioPayOnLine_2p.PSP_RequerimientoStruct_PayOnLine_2p = prodResquest
	  err := service.client.Call(CONSTANTS.PAY_ONLINE_2P, servicioPayOnLine_2p, prodResponse)
	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_PayOnLine_2p)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_PayOnLine_2p, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_PayOnLine_2p)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_PayOnLine_2p)

	  servicioPayOnLine_2p := new(stgservice.ServiceStruct_PayOnLine_2p)
	  servicioPayOnLine_2p.PSP_RequerimientoStruct_PayOnLine_2p = stgResquest

	  err := service.client.Call(CONSTANTS.PAY_ONLINE_2P, servicioPayOnLine_2p, stgResponse)
	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_PayOnLine_2p)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_PayOnLine_2p, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_PayOnLine_2p)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_PayOnLine_2p)

	  servicioPayOnLine_2p := new(sndservice.ServiceStruct_PayOnLine_2p)
	  servicioPayOnLine_2p.PSP_RequerimientoStruct_PayOnLine_2p = sndResquest

	  err := service.client.Call(CONSTANTS.PAY_ONLINE_2P, servicioPayOnLine_2p, sndResponse)
	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_PayOnLine_2p)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_PayOnLine_2p, response)

	  return response, nil
	}

	return nil, nil
}

func (service *PaymentServicePlatformPortType) Authorize_2p(request *RequerimientoStruct_Authorize_2p) (*RespuestaStruct_Authorize_2p, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_Authorize_2p)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_Authorize_2p)     

 	  servicioAuthorize_2p := new(prodservice.ServiceStruct_Authorize_2p)
	  servicioAuthorize_2p.PSP_RequerimientoStruct_Authorize_2p = prodResquest

	  err := service.client.Call(CONSTANTS.AUTHORIZE_2P, servicioAuthorize_2p, prodResponse)
	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_Authorize_2p)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_Authorize_2p, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_Authorize_2p)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_Authorize_2p)     

 	  servicioAuthorize_2p := new(stgservice.ServiceStruct_Authorize_2p)
	  servicioAuthorize_2p.PSP_RequerimientoStruct_Authorize_2p = stgResquest

	  err := service.client.Call(CONSTANTS.AUTHORIZE_2P, servicioAuthorize_2p, stgResponse)
	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_Authorize_2p)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_Authorize_2p, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_Authorize_2p)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_Authorize_2p)     

 	  servicioAuthorize_2p := new(sndservice.ServiceStruct_Authorize_2p)
	  servicioAuthorize_2p.PSP_RequerimientoStruct_Authorize_2p = sndResquest

	  err := service.client.Call(CONSTANTS.AUTHORIZE_2P, servicioAuthorize_2p, sndResponse)
	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_Authorize_2p)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_Authorize_2p, response)

	  return response, nil

       }

       return nil, nil
}

func (service *PaymentServicePlatformPortType) BankPayment_2p(request *RequerimientoStruct_BankPayment_2p) (*RespuestaStruct_BankPayment_2p, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_BankPayment_2p)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_BankPayment_2p)     

	  servicioBankPayment_2p := new(prodservice.ServiceStruct_BankPayment_2p)
	  servicioBankPayment_2p.PSP_RequerimientoStruct_BankPayment_2p = prodResquest

	  err := service.client.Call(CONSTANTS.BANK_PAYMENT_2P, servicioBankPayment_2p, prodResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_BankPayment_2p)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_BankPayment_2p, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_BankPayment_2p)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_BankPayment_2p)     

	  servicioBankPayment_2p := new(stgservice.ServiceStruct_BankPayment_2p)
	  servicioBankPayment_2p.PSP_RequerimientoStruct_BankPayment_2p = stgResquest

	  err := service.client.Call(CONSTANTS.BANK_PAYMENT_2P, servicioBankPayment_2p, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_BankPayment_2p)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_BankPayment_2p, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_BankPayment_2p)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_BankPayment_2p)     

	  servicioBankPayment_2p := new(sndservice.ServiceStruct_BankPayment_2p)
	  servicioBankPayment_2p.PSP_RequerimientoStruct_BankPayment_2p = sndResquest

	  err := service.client.Call(CONSTANTS.BANK_PAYMENT_2P, servicioBankPayment_2p, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_BankPayment_2p)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_BankPayment_2p, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) SplitPayOnLine_2p(request *RequerimientoStruct_SplitPayOnLine_2p) (*RespuestaStruct_SplitPayOnLine_2p, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_SplitPayOnLine_2p)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_SplitPayOnLine_2p)

	  servicioSplitPayOnLine_2p := new(prodservice.ServiceStruct_SplitPayOnLine_2p)
	  servicioSplitPayOnLine_2p.PSP_RequerimientoStruct_SplitPayOnLine_2p = prodResquest

	  err := service.client.Call(CONSTANTS.SPLIT_PAY_ONLINE_2P, servicioSplitPayOnLine_2p, prodResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_SplitPayOnLine_2p)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_SplitPayOnLine_2p, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_SplitPayOnLine_2p)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_SplitPayOnLine_2p)

	  servicioSplitPayOnLine_2p := new(stgservice.ServiceStruct_SplitPayOnLine_2p)
	  servicioSplitPayOnLine_2p.PSP_RequerimientoStruct_SplitPayOnLine_2p = stgResquest

	  err := service.client.Call(CONSTANTS.SPLIT_PAY_ONLINE_2P, servicioSplitPayOnLine_2p, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_SplitPayOnLine_2p)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_SplitPayOnLine_2p, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_SplitPayOnLine_2p)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_SplitPayOnLine_2p)

	  servicioSplitPayOnLine_2p := new(sndservice.ServiceStruct_SplitPayOnLine_2p)
	  servicioSplitPayOnLine_2p.PSP_RequerimientoStruct_SplitPayOnLine_2p = sndResquest

	  err := service.client.Call(CONSTANTS.SPLIT_PAY_ONLINE_2P, servicioSplitPayOnLine_2p, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_SplitPayOnLine_2p)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_SplitPayOnLine_2p, response)

	  return response, nil


        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) SplitAuthorize_2p(request *RequerimientoStruct_SplitAuthorize_2p) (*RespuestaStruct_SplitAuthorize_2p, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_SplitAuthorize_2p)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_SplitAuthorize_2p)

	  servicioSplitAuthorize_2p := new(prodservice.ServiceStruct_SplitAuthorize_2p)
	  servicioSplitAuthorize_2p.PSP_RequerimientoStruct_SplitAuthorize_2p = prodResquest

	  err := service.client.Call(CONSTANTS.SPLIT_AUTHORIZE_2P, servicioSplitAuthorize_2p, prodResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_SplitAuthorize_2p)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_SplitAuthorize_2p, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_SplitAuthorize_2p)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_SplitAuthorize_2p)

	  servicioSplitAuthorize_2p := new(stgservice.ServiceStruct_SplitAuthorize_2p)
	  servicioSplitAuthorize_2p.PSP_RequerimientoStruct_SplitAuthorize_2p = stgResquest

	  err := service.client.Call(CONSTANTS.SPLIT_AUTHORIZE_2P, servicioSplitAuthorize_2p, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_SplitAuthorize_2p)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_SplitAuthorize_2p, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_SplitAuthorize_2p)
	  CopypStruct(request, sndResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_SplitAuthorize_2p)

	  servicioSplitAuthorize_2p := new(sndservice.ServiceStruct_SplitAuthorize_2p)
	  servicioSplitAuthorize_2p.PSP_RequerimientoStruct_SplitAuthorize_2p = sndResquest

	  err := service.client.Call(CONSTANTS.SPLIT_AUTHORIZE_2P, servicioSplitAuthorize_2p, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_SplitAuthorize_2p)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_SplitAuthorize_2p, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) PayOnLine_3p(request *RequerimientoStruct_PayOnLine_3p) (*RespuestaStruct_PayOnLine_3p, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_PayOnLine_3p)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_PayOnLine_3p)

	  servicioPayOnLine_3p := new(prodservice.ServiceStruct_PayOnLine_3p)
	  servicioPayOnLine_3p.PSP_RequerimientoStruct_PayOnLine_3p = prodResquest

	  err := service.client.Call(CONSTANTS.PAY_ONLINE_3P, servicioPayOnLine_3p, prodResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_PayOnLine_3p)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_PayOnLine_3p, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_PayOnLine_3p)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_PayOnLine_3p)

	  servicioPayOnLine_3p := new(stgservice.ServiceStruct_PayOnLine_3p)
	  servicioPayOnLine_3p.PSP_RequerimientoStruct_PayOnLine_3p = stgResquest

	  err := service.client.Call(CONSTANTS.PAY_ONLINE_3P, servicioPayOnLine_3p, stgResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_PayOnLine_3p)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_PayOnLine_3p, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_PayOnLine_3p)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_PayOnLine_3p)

	  servicioPayOnLine_3p := new(sndservice.ServiceStruct_PayOnLine_3p)
	  servicioPayOnLine_3p.PSP_RequerimientoStruct_PayOnLine_3p = sndResquest

	  err := service.client.Call(CONSTANTS.PAY_ONLINE_3P, servicioPayOnLine_3p, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_PayOnLine_3p)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_PayOnLine_3p, response)

	  return response, nil
        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) SplitPayOnLine_3p(request *RequerimientoStruct_SplitPayOnLine_3p) (*RespuestaStruct_SplitPayOnLine_3p, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_SplitPayOnLine_3p)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_SplitPayOnLine_3p)

	  servicioSplitPayOnLine_3p := new(prodservice.ServiceStruct_SplitPayOnLine_3p)
	  servicioSplitPayOnLine_3p.PSP_RequerimientoStruct_SplitPayOnLine_3p = prodResquest

	  err := service.client.Call(CONSTANTS.SPLIT_PAY_ONLINE_3P, servicioSplitPayOnLine_3p, prodResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_SplitPayOnLine_3p)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_SplitPayOnLine_3p, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_SplitPayOnLine_3p)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(prodservice.ServiceRespuestaStruct_SplitPayOnLine_3p)

	  servicioSplitPayOnLine_3p := new(stgservice.ServiceStruct_SplitPayOnLine_3p)
	  servicioSplitPayOnLine_3p.PSP_RequerimientoStruct_SplitPayOnLine_3p = stgResquest

	  err := service.client.Call(CONSTANTS.SPLIT_PAY_ONLINE_3P, servicioSplitPayOnLine_3p, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_SplitPayOnLine_3p)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_SplitPayOnLine_3p, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_SplitPayOnLine_3p)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_SplitPayOnLine_3p)

	  servicioSplitPayOnLine_3p := new(sndservice.ServiceStruct_SplitPayOnLine_3p)
	  servicioSplitPayOnLine_3p.PSP_RequerimientoStruct_SplitPayOnLine_3p = sndResquest

	  err := service.client.Call(CONSTANTS.SPLIT_PAY_ONLINE_3P, servicioSplitPayOnLine_3p, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_SplitPayOnLine_3p)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_SplitPayOnLine_3p, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) Authorize_3p(request *RequerimientoStruct_Authorize_3p) (*RespuestaStruct_Authorize_3p, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_Authorize_3p)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_Authorize_3p)

	  servicioAuthorize_3p := new(prodservice.ServiceStruct_Authorize_3p)
	  servicioAuthorize_3p.PSP_RequerimientoStruct_Authorize_3p = prodResquest

	  err := service.client.Call(CONSTANTS.AUTHORIZE_3P, servicioAuthorize_3p, prodResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_Authorize_3p)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_Authorize_3p, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_Authorize_3p)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_Authorize_3p)

	  servicioAuthorize_3p := new(stgservice.ServiceStruct_Authorize_3p)
	  servicioAuthorize_3p.PSP_RequerimientoStruct_Authorize_3p = stgResquest

	  err := service.client.Call(CONSTANTS.AUTHORIZE_3P, servicioAuthorize_3p, stgResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_Authorize_3p)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_Authorize_3p, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_Authorize_3p)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_Authorize_3p)

	  servicioAuthorize_3p := new(sndservice.ServiceStruct_Authorize_3p)
	  servicioAuthorize_3p.PSP_RequerimientoStruct_Authorize_3p = sndResquest

	  err := service.client.Call(CONSTANTS.AUTHORIZE_3P, servicioAuthorize_3p, sndResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_Authorize_3p)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_Authorize_3p, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) SplitAuthorize_3p(request *RequerimientoStruct_SplitAuthorize_3p) (*RespuestaStruct_SplitAuthorize_3p, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_SplitAuthorize_3p)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_SplitAuthorize_3p)

	  serviceSplitAuthorize_3p := new(prodservice.ServiceStruct_SplitAuthorize_3p)
	  serviceSplitAuthorize_3p.PSP_RequerimientoStruct_SplitAuthorize_3p = prodResquest

	  err := service.client.Call(CONSTANTS.SPLIT_AUTHORIZE_3P, serviceSplitAuthorize_3p, prodResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_SplitAuthorize_3p)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_SplitAuthorize_3p, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_SplitAuthorize_3p)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_SplitAuthorize_3p)

	  serviceSplitAuthorize_3p := new(stgservice.ServiceStruct_SplitAuthorize_3p)
	  serviceSplitAuthorize_3p.PSP_RequerimientoStruct_SplitAuthorize_3p = stgResquest

	  err := service.client.Call(CONSTANTS.SPLIT_AUTHORIZE_3P, serviceSplitAuthorize_3p, stgResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_SplitAuthorize_3p)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_SplitAuthorize_3p, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_SplitAuthorize_3p)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_SplitAuthorize_3p)

	  serviceSplitAuthorize_3p := new(sndservice.ServiceStruct_SplitAuthorize_3p)
	  serviceSplitAuthorize_3p.PSP_RequerimientoStruct_SplitAuthorize_3p = sndResquest

	  err := service.client.Call(CONSTANTS.SPLIT_AUTHORIZE_3P, serviceSplitAuthorize_3p, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_SplitAuthorize_3p)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_SplitAuthorize_3p, response)

	  return response, nil
        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) BankPayment_3p(request *RequerimientoStruct_BankPayment_3p) (*RespuestaStruct_BankPayment_3p, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_BankPayment_3p)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_BankPayment_3p)

	  servicioBankPayment_3p := new(prodservice.ServiceStruct_BankPayment_3p)
	  servicioBankPayment_3p.PSP_RequerimientoStruct_BankPayment_3p = prodResquest

	  err := service.client.Call(CONSTANTS.BANK_PAYMENT_3P, servicioBankPayment_3p, prodResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_BankPayment_3p)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_BankPayment_3p, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_BankPayment_3p)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_BankPayment_3p)

	  servicioBankPayment_3p := new(stgservice.ServiceStruct_BankPayment_3p)
	  servicioBankPayment_3p.PSP_RequerimientoStruct_BankPayment_3p = stgResquest

	  err := service.client.Call(CONSTANTS.BANK_PAYMENT_3P, servicioBankPayment_3p, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_BankPayment_3p)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_BankPayment_3p, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_BankPayment_3p)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_BankPayment_3p)

	  servicioBankPayment_3p := new(sndservice.ServiceStruct_BankPayment_3p)
	  servicioBankPayment_3p.PSP_RequerimientoStruct_BankPayment_3p = sndResquest

	  err := service.client.Call(CONSTANTS.BANK_PAYMENT_3P, servicioBankPayment_3p, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_BankPayment_3p)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_BankPayment_3p, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) CashPayment_3p(request *RequerimientoStruct_CashPayment_3p) (*RespuestaStruct_CashPayment_3p, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_CashPayment_3p)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_CashPayment_3p)

	  servicioCashPayment_3p := new(prodservice.ServiceStruct_CashPayment_3p)
	  servicioCashPayment_3p.PSP_RequerimientoStruct_CashPayment_3p = prodResquest

	  err := service.client.Call(CONSTANTS.CASH_PAYMENT_3P, servicioCashPayment_3p, prodResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_CashPayment_3p)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_CashPayment_3p, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_CashPayment_3p)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_CashPayment_3p)

	  servicioCashPayment_3p := new(stgservice.ServiceStruct_CashPayment_3p)
	  servicioCashPayment_3p.PSP_RequerimientoStruct_CashPayment_3p = stgResquest

	  err := service.client.Call(CONSTANTS.CASH_PAYMENT_3P, servicioCashPayment_3p, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_CashPayment_3p)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_CashPayment_3p, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_CashPayment_3p)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_CashPayment_3p)

	  servicioCashPayment_3p := new(sndservice.ServiceStruct_CashPayment_3p)
	  servicioCashPayment_3p.PSP_RequerimientoStruct_CashPayment_3p = sndResquest

	  err := service.client.Call(CONSTANTS.CASH_PAYMENT_3P, servicioCashPayment_3p, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_CashPayment_3p)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_CashPayment_3p, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) Capture(request *RequerimientoStruct_Capture) (*RespuestaStruct_Capture, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_Capture)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_Capture)

	  servicioCapture := new(prodservice.ServiceStruct_Capture)
	  servicioCapture.PSP_RequerimientoStruct_Capture = prodResquest

	  err := service.client.Call(CONSTANTS.CAPTURE, servicioCapture, prodResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_Capture)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_Capture, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_Capture)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_Capture)

	  servicioCapture := new(stgservice.ServiceStruct_Capture)
	  servicioCapture.PSP_RequerimientoStruct_Capture = stgResquest

	  err := service.client.Call(CONSTANTS.CAPTURE, servicioCapture, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_Capture)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_Capture, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_Capture)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_Capture)

	  servicioCapture := new(sndservice.ServiceStruct_Capture)
	  servicioCapture.PSP_RequerimientoStruct_Capture = sndResquest

	  err := service.client.Call(CONSTANTS.CAPTURE, servicioCapture, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_Capture)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_Capture, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) Refund(request *RequerimientoStruct_Refund) (*RespuestaStruct_Refund, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_Refund)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_Refund)

	  servicioRefund := new(prodservice.ServiceStruct_Refund)
	  servicioRefund.PSP_RequerimientoStruct_Refund = prodResquest

	  err := service.client.Call(CONSTANTS.REFUND, servicioRefund, prodResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_Refund)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_Refund, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_Refund)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_Refund)

	  servicioRefund := new(stgservice.ServiceStruct_Refund)
	  servicioRefund.PSP_RequerimientoStruct_Refund = stgResquest

	  err := service.client.Call(CONSTANTS.REFUND, servicioRefund, stgResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_Refund)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_Refund, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_Refund)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_Refund)

	  servicioRefund := new(sndservice.ServiceStruct_Refund)
	  servicioRefund.PSP_RequerimientoStruct_Refund = sndResquest

	  err := service.client.Call(CONSTANTS.REFUND, servicioRefund, sndResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_Refund)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_Refund, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) GetInstallmentsOptions(request *RequerimientoStruct_GetInstallmentsOptions) (*RespuestaStruct_GetInstallmentsOptions, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_GetInstallmentsOptions)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_GetInstallmentsOptions)

	  servicioGetInstallmentsOptions := new(prodservice.ServiceStruct_GetInstallmentsOptions)
	  servicioGetInstallmentsOptions.PSP_RequerimientoStruct_GetInstallmentsOptions = prodResquest

	  err := service.client.Call(CONSTANTS.GET_INSTALLMENTS_OPTIONS, servicioGetInstallmentsOptions, prodResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_GetInstallmentsOptions)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_GetInstallmentsOptions, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_GetInstallmentsOptions)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_GetInstallmentsOptions)

	  servicioGetInstallmentsOptions := new(stgservice.ServiceStruct_GetInstallmentsOptions)
	  servicioGetInstallmentsOptions.PSP_RequerimientoStruct_GetInstallmentsOptions = stgResquest

	  err := service.client.Call(CONSTANTS.GET_INSTALLMENTS_OPTIONS, servicioGetInstallmentsOptions, stgResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_GetInstallmentsOptions)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_GetInstallmentsOptions, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_GetInstallmentsOptions)
	  CopypStruct(request, sndResquest)

	  stgResponse := new(sndservice.ServiceRespuestaStruct_GetInstallmentsOptions)

	  servicioGetInstallmentsOptions := new(sndservice.ServiceStruct_GetInstallmentsOptions)
	  servicioGetInstallmentsOptions.PSP_RequerimientoStruct_GetInstallmentsOptions =sndResquest

	  err := service.client.Call(CONSTANTS.GET_INSTALLMENTS_OPTIONS, servicioGetInstallmentsOptions, stgResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_GetInstallmentsOptions)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_GetInstallmentsOptions, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) CreatePaymentMethodToken(request *RequerimientoStruct_CreatePaymentMethodToken) (*RespuestaStruct_CreatePaymentMethodToken, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_CreatePaymentMethodToken)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_CreatePaymentMethodToken)

	  if prodResquest.Psp_CardInputDetails != nil {
		prodResquest.Psp_CardInputDetails.XMLName = xml.Name{"", "psp_CardInputDetails"}
	  }
	  if prodResquest.Psp_Address != nil {
		prodResquest.Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if prodResquest.Psp_Person != nil {
		prodResquest.Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }

	  servicioCreatePaymentMethodToken := new(prodservice.ServiceStruct_CreatePaymentMethodToken)
	  servicioCreatePaymentMethodToken.PSP_RequerimientoStruct_CreatePaymentMethodToken = prodResquest

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD_TOKEN, servicioCreatePaymentMethodToken, prodResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_CreatePaymentMethodToken)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_CreatePaymentMethodToken, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_CreatePaymentMethodToken)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_CreatePaymentMethodToken)

	  if stgResquest.Psp_CardInputDetails != nil {
		stgResquest.Psp_CardInputDetails.XMLName = xml.Name{"", "psp_CardInputDetails"}
	  }
	  if stgResquest.Psp_Address != nil {
		stgResquest.Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if stgResquest.Psp_Person != nil {
		stgResquest.Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }

	  servicioCreatePaymentMethodToken := new(stgservice.ServiceStruct_CreatePaymentMethodToken)
	  servicioCreatePaymentMethodToken.PSP_RequerimientoStruct_CreatePaymentMethodToken = stgResquest

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD_TOKEN, servicioCreatePaymentMethodToken, stgResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_CreatePaymentMethodToken)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_CreatePaymentMethodToken, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_CreatePaymentMethodToken)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_CreatePaymentMethodToken)

	  if sndResquest.Psp_CardInputDetails != nil {
		sndResquest.Psp_CardInputDetails.XMLName = xml.Name{"", "psp_CardInputDetails"}
	  }
	  if sndResquest.Psp_Address != nil {
		sndResquest.Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if sndResquest.Psp_Person != nil {
		sndResquest.Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }

	  servicioCreatePaymentMethodToken := new(sndservice.ServiceStruct_CreatePaymentMethodToken)
	  servicioCreatePaymentMethodToken.PSP_RequerimientoStruct_CreatePaymentMethodToken = sndResquest

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD_TOKEN, servicioCreatePaymentMethodToken, sndResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_CreatePaymentMethodToken)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_CreatePaymentMethodToken, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) RetrievePaymentMethodToken(request *RequerimientoStruct_RetrievePaymentMethodToken) (*RespuestaStruct_RetrievePaymentMethodToken, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_RetrievePaymentMethodToken)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_RetrievePaymentMethodToken)

	  servicioRetrievePaymentMethodToken := new(prodservice.ServiceStruct_RetrievePaymentMethodToken)
	  servicioRetrievePaymentMethodToken.PSP_RequerimientoStruct_RetrievePaymentMethodToken = prodResquest

	  err := service.client.Call(CONSTANTS.RETRIEVE_PAYMENT_METHOD_TOKEN, servicioRetrievePaymentMethodToken, prodResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_RetrievePaymentMethodToken)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_RetrievePaymentMethodToken, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_RetrievePaymentMethodToken)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_RetrievePaymentMethodToken)

	  servicioRetrievePaymentMethodToken := new(stgservice.ServiceStruct_RetrievePaymentMethodToken)
	  servicioRetrievePaymentMethodToken.PSP_RequerimientoStruct_RetrievePaymentMethodToken = stgResquest

	  err := service.client.Call(CONSTANTS.RETRIEVE_PAYMENT_METHOD_TOKEN, servicioRetrievePaymentMethodToken, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_RetrievePaymentMethodToken)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_RetrievePaymentMethodToken, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_RetrievePaymentMethodToken)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_RetrievePaymentMethodToken)

	  servicioRetrievePaymentMethodToken := new(sndservice.ServiceStruct_RetrievePaymentMethodToken)
	  servicioRetrievePaymentMethodToken.PSP_RequerimientoStruct_RetrievePaymentMethodToken = sndResquest

	  err := service.client.Call(CONSTANTS.RETRIEVE_PAYMENT_METHOD_TOKEN, servicioRetrievePaymentMethodToken, sndResponse)

	  if err != nil {
		return nil, err
	  }
          
	  response := new(RespuestaStruct_RetrievePaymentMethodToken)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_RetrievePaymentMethodToken, response)
	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) RecachePaymentMethodToken(request *RequerimientoStruct_RecachePaymentMethodToken) (*RespuestaStruct_RecachePaymentMethodToken, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_RecachePaymentMethodToken)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_RecachePaymentMethodToken)

	  if prodResquest.Psp_Address != nil {
		prodResquest.Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if prodResquest.Psp_Person != nil {
		prodResquest.Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }

	  servicioRecachePaymentMethodToken := new(prodservice.ServiceStruct_RecachePaymentMethodToken)
	  servicioRecachePaymentMethodToken.PSP_RequerimientoStruct_RecachePaymentMethodToken = prodResquest

	  err := service.client.Call(CONSTANTS.RECACHE_PAYMENT_METHOD_TOKEN, servicioRecachePaymentMethodToken, prodResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_RecachePaymentMethodToken)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_RecachePaymentMethodToken, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_RecachePaymentMethodToken)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_RecachePaymentMethodToken)

	  if stgResquest.Psp_Address != nil {
		stgResquest.Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if stgResquest.Psp_Person != nil {
		stgResquest.Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }

	  servicioRecachePaymentMethodToken := new(stgservice.ServiceStruct_RecachePaymentMethodToken)
	  servicioRecachePaymentMethodToken.PSP_RequerimientoStruct_RecachePaymentMethodToken = stgResquest

	  err := service.client.Call(CONSTANTS.RECACHE_PAYMENT_METHOD_TOKEN, servicioRecachePaymentMethodToken, stgResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_RecachePaymentMethodToken)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_RecachePaymentMethodToken, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_RecachePaymentMethodToken)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_RecachePaymentMethodToken)

	  if sndResquest.Psp_Address != nil {
		sndResquest.Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if sndResquest.Psp_Person != nil {
		sndResquest.Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }

	  servicioRecachePaymentMethodToken := new(sndservice.ServiceStruct_RecachePaymentMethodToken)
	  servicioRecachePaymentMethodToken.PSP_RequerimientoStruct_RecachePaymentMethodToken = sndResquest

	  err := service.client.Call(CONSTANTS.RECACHE_PAYMENT_METHOD_TOKEN, servicioRecachePaymentMethodToken, sndResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_RecachePaymentMethodToken)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_RecachePaymentMethodToken, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) CreateClientSession(request *RequerimientoStruct_CreateClientSession) (*RespuestaStruct_CreateClientSession, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_CreateClientSession)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_CreateClientSession)

	  servicioCreateClientSession := new(prodservice.ServiceStruct_CreateClientSession)
	  servicioCreateClientSession.PSP_RequerimientoStruct_CreateClientSession = prodResquest

	  err := service.client.Call(CONSTANTS.CREATE_CLIENT_SESSION, servicioCreateClientSession, prodResponse)
	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_CreateClientSession)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_CreateClientSession, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_CreateClientSession)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_CreateClientSession)

	  servicioCreateClientSession := new(stgservice.ServiceStruct_CreateClientSession)
	  servicioCreateClientSession.PSP_RequerimientoStruct_CreateClientSession = stgResquest

	  err := service.client.Call(CONSTANTS.CREATE_CLIENT_SESSION, servicioCreateClientSession, stgResponse)
	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_CreateClientSession)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_CreateClientSession, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_CreateClientSession)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_CreateClientSession)

	  servicioCreateClientSession := new(sndservice.ServiceStruct_CreateClientSession)
	  servicioCreateClientSession.PSP_RequerimientoStruct_CreateClientSession = sndResquest

	  err := service.client.Call(CONSTANTS.CREATE_CLIENT_SESSION, servicioCreateClientSession, sndResponse)
	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_CreateClientSession)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_CreateClientSession, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) CreatePaymentMethod(request *RequerimientoStruct_CreatePaymentMethod) (*RespuestaStruct_CreatePaymentMethod, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_CreatePaymentMethod)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_CreatePaymentMethod)

	  servicioCreatePaymentMethod := new(prodservice.ServiceStruct_CreatePaymentMethod)
	  servicioCreatePaymentMethod.PSP_RequerimientoStruct_CreatePaymentMethod = prodResquest

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD, servicioCreatePaymentMethod, prodResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_CreatePaymentMethod)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_CreatePaymentMethod, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_CreatePaymentMethod)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_CreatePaymentMethod)

	  servicioCreatePaymentMethod := new(stgservice.ServiceStruct_CreatePaymentMethod)
	  servicioCreatePaymentMethod.PSP_RequerimientoStruct_CreatePaymentMethod = stgResquest

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD, servicioCreatePaymentMethod, stgResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_CreatePaymentMethod)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_CreatePaymentMethod, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_CreatePaymentMethod)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_CreatePaymentMethod)

	  servicioCreatePaymentMethod := new(sndservice.ServiceStruct_CreatePaymentMethod)
	  servicioCreatePaymentMethod.PSP_RequerimientoStruct_CreatePaymentMethod = sndResquest

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD, servicioCreatePaymentMethod, sndResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_CreatePaymentMethod)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_CreatePaymentMethod, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) CreatePaymentMethodFromPayment(request *RequerimientoStruct_CreatePaymentMethodFromPayment) (*RespuestaStruct_CreatePaymentMethodFromPayment, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_CreatePaymentMethodFromPayment)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_CreatePaymentMethodFromPayment)

	  servicioCreatePaymentMethodFromPayment := new(prodservice.ServiceStruct_CreatePaymentMethodFromPayment)
	  servicioCreatePaymentMethodFromPayment.PSP_RequerimientoStruct_CreatePaymentMethodFromPayment = prodResquest

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD_FROM_PAYMENT, servicioCreatePaymentMethodFromPayment, prodResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_CreatePaymentMethodFromPayment)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_CreatePaymentMethodFromPayment, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_CreatePaymentMethodFromPayment)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_CreatePaymentMethodFromPayment)

	  servicioCreatePaymentMethodFromPayment := new(stgservice.ServiceStruct_CreatePaymentMethodFromPayment)
	  servicioCreatePaymentMethodFromPayment.PSP_RequerimientoStruct_CreatePaymentMethodFromPayment = stgResquest

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD_FROM_PAYMENT, servicioCreatePaymentMethodFromPayment, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_CreatePaymentMethodFromPayment)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_CreatePaymentMethodFromPayment, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_CreatePaymentMethodFromPayment)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_CreatePaymentMethodFromPayment)

	  servicioCreatePaymentMethodFromPayment := new(sndservice.ServiceStruct_CreatePaymentMethodFromPayment)
	  servicioCreatePaymentMethodFromPayment.PSP_RequerimientoStruct_CreatePaymentMethodFromPayment = sndResquest

	  err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD_FROM_PAYMENT, servicioCreatePaymentMethodFromPayment, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_CreatePaymentMethodFromPayment)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_CreatePaymentMethodFromPayment, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) RetrievePaymentMethod(request *RequerimientoStruct_RetrievePaymentMethod) (*RespuestaStruct_RetrievePaymentMethod, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_RetrievePaymentMethod)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_RetrievePaymentMethod)

	  servicioRetrievePaymentMethod := new(prodservice.ServiceStruct_RetrievePaymentMethod)
	  servicioRetrievePaymentMethod.PSP_RequerimientoStruct_RetrievePaymentMethod = prodResquest

	  err := service.client.Call(CONSTANTS.RETRIEVE_PAYMENT_METHOD, servicioRetrievePaymentMethod, prodResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_RetrievePaymentMethod)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_RetrievePaymentMethod, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_RetrievePaymentMethod)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_RetrievePaymentMethod)

	  servicioRetrievePaymentMethod := new(stgservice.ServiceStruct_RetrievePaymentMethod)
	  servicioRetrievePaymentMethod.PSP_RequerimientoStruct_RetrievePaymentMethod = stgResquest

	  err := service.client.Call(CONSTANTS.RETRIEVE_PAYMENT_METHOD, servicioRetrievePaymentMethod, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_RetrievePaymentMethod)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_RetrievePaymentMethod, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_RetrievePaymentMethod)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_RetrievePaymentMethod)

	  servicioRetrievePaymentMethod := new(sndservice.ServiceStruct_RetrievePaymentMethod)
	  servicioRetrievePaymentMethod.PSP_RequerimientoStruct_RetrievePaymentMethod = sndResquest

	  err := service.client.Call(CONSTANTS.RETRIEVE_PAYMENT_METHOD, servicioRetrievePaymentMethod, sndResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_RetrievePaymentMethod)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_RetrievePaymentMethod, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) UpdatePaymentMethod(request *RequerimientoStruct_UpdatePaymentMethod) (*RespuestaStruct_UpdatePaymentMethod, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_UpdatePaymentMethod)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_UpdatePaymentMethod)

	  servicioUpdatePaymentMethod := new(prodservice.ServiceStruct_UpdatePaymentMethod)
	  servicioUpdatePaymentMethod.PSP_RequerimientoStruct_UpdatePaymentMethod = prodResquest

	  err := service.client.Call(CONSTANTS.UPDATE_PAYMENT_METHOD, servicioUpdatePaymentMethod, prodResponse)
	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_UpdatePaymentMethod)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_UpdatePaymentMethod, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_UpdatePaymentMethod)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_UpdatePaymentMethod)

	  servicioUpdatePaymentMethod := new(stgservice.ServiceStruct_UpdatePaymentMethod)
	  servicioUpdatePaymentMethod.PSP_RequerimientoStruct_UpdatePaymentMethod = stgResquest

	  err := service.client.Call(CONSTANTS.UPDATE_PAYMENT_METHOD, servicioUpdatePaymentMethod, stgResponse)
	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_UpdatePaymentMethod)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_UpdatePaymentMethod, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_UpdatePaymentMethod)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_UpdatePaymentMethod)

	  servicioUpdatePaymentMethod := new(sndservice.ServiceStruct_UpdatePaymentMethod)
	  servicioUpdatePaymentMethod.PSP_RequerimientoStruct_UpdatePaymentMethod = sndResquest

	  err := service.client.Call(CONSTANTS.UPDATE_PAYMENT_METHOD, servicioUpdatePaymentMethod, sndResponse)
	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_UpdatePaymentMethod)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_UpdatePaymentMethod, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) DeletePaymentMethod(request *RequerimientoStruct_DeletePaymentMethod) (*RespuestaStruct_DeletePaymentMethod, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_DeletePaymentMethod)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_DeletePaymentMethod)

	  servicioDeletePaymentMethod := new(prodservice.ServiceStruct_DeletePaymentMethod)
	  servicioDeletePaymentMethod.PSP_RequerimientoStruct_DeletePaymentMethod = prodResquest

	  err := service.client.Call(CONSTANTS.DELETE_PAYMENT_METHOD, servicioDeletePaymentMethod, prodResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_DeletePaymentMethod)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_DeletePaymentMethod, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_DeletePaymentMethod)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_DeletePaymentMethod)

	  servicioDeletePaymentMethod := new(stgservice.ServiceStruct_DeletePaymentMethod)
	  servicioDeletePaymentMethod.PSP_RequerimientoStruct_DeletePaymentMethod = stgResquest

	  err := service.client.Call(CONSTANTS.DELETE_PAYMENT_METHOD, servicioDeletePaymentMethod, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_DeletePaymentMethod)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_DeletePaymentMethod, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_DeletePaymentMethod)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_DeletePaymentMethod)

	  servicioDeletePaymentMethod := new(sndservice.ServiceStruct_DeletePaymentMethod)
	  servicioDeletePaymentMethod.PSP_RequerimientoStruct_DeletePaymentMethod = sndResquest

	  err := service.client.Call(CONSTANTS.DELETE_PAYMENT_METHOD, servicioDeletePaymentMethod, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_DeletePaymentMethod)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_DeletePaymentMethod, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) CreateCustomer(request *RequerimientoStruct_CreateCustomer) (*RespuestaStruct_CreateCustomer, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_CreateCustomer)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_CreateCustomer)

	  if prodResquest.Psp_Address != nil {
		prodResquest.Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if prodResquest.Psp_Person != nil {
		prodResquest.Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }
	  if prodResquest.Psp_PaymentMethod != nil && 
	     prodResquest.Psp_PaymentMethod.Person != nil {
		prodResquest.Psp_PaymentMethod.Person.XMLName = xml.Name{"", "Person"}
	  }
	  if prodResquest.Psp_PaymentMethod != nil && 
	     prodResquest.Psp_PaymentMethod.Address != nil {
		prodResquest.Psp_PaymentMethod.Address.XMLName = xml.Name{"", "Address"}
	  }

	  servicioCreateCustomer := new(prodservice.ServiceStruct_CreateCustomer)
	  servicioCreateCustomer.PSP_RequerimientoStruct_CreateCustomer = prodResquest

	  err := service.client.Call(CONSTANTS.CREATE_CUSTOMER, servicioCreateCustomer, prodResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_CreateCustomer)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_CreateCustomer, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_CreateCustomer)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_CreateCustomer)

	  if stgResquest.Psp_Address != nil {
		stgResquest.Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if stgResquest.Psp_Person != nil {
		stgResquest.Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }
	  if stgResquest.Psp_PaymentMethod != nil && 
	     stgResquest.Psp_PaymentMethod.Person != nil {
		stgResquest.Psp_PaymentMethod.Person.XMLName = xml.Name{"", "Person"}
	  }
	  if stgResquest.Psp_PaymentMethod != nil && 
	     stgResquest.Psp_PaymentMethod.Address != nil {
		stgResquest.Psp_PaymentMethod.Address.XMLName = xml.Name{"", "Address"}
	  }

	  servicioCreateCustomer := new(stgservice.ServiceStruct_CreateCustomer)
	  servicioCreateCustomer.PSP_RequerimientoStruct_CreateCustomer = stgResquest

	  err := service.client.Call(CONSTANTS.CREATE_CUSTOMER, servicioCreateCustomer, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_CreateCustomer)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_CreateCustomer, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_CreateCustomer)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_CreateCustomer)

	  if sndResquest.Psp_Address != nil {
		sndResquest.Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if sndResquest.Psp_Person != nil {
		sndResquest.Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }
	  if sndResquest.Psp_PaymentMethod != nil && 
	     sndResquest.Psp_PaymentMethod.Person != nil {
		sndResquest.Psp_PaymentMethod.Person.XMLName = xml.Name{"", "Person"}
	  }
	  if sndResquest.Psp_PaymentMethod != nil && 
	     sndResquest.Psp_PaymentMethod.Address != nil {
		sndResquest.Psp_PaymentMethod.Address.XMLName = xml.Name{"", "Address"}
	  }
	  servicioCreateCustomer := new(sndservice.ServiceStruct_CreateCustomer)
	  servicioCreateCustomer.PSP_RequerimientoStruct_CreateCustomer = sndResquest

	  err := service.client.Call(CONSTANTS.CREATE_CUSTOMER, servicioCreateCustomer, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_CreateCustomer)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_CreateCustomer, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) RetrieveCustomer(request *RequerimientoStruct_RetrieveCustomer) (*RespuestaStruct_RetrieveCustomer, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_RetrieveCustomer)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_RetrieveCustomer)

	  servicioRetrieveCustomer := new(prodservice.ServiceStruct_RetrieveCustomer)
	  servicioRetrieveCustomer.PSP_RequerimientoStruct_RetrieveCustomer = prodResquest

	  err := service.client.Call(CONSTANTS.RETRIEVE_CUSTOMER, servicioRetrieveCustomer, prodResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_RetrieveCustomer)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_RetrieveCustomer, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_RetrieveCustomer)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_RetrieveCustomer)

	  servicioRetrieveCustomer := new(stgservice.ServiceStruct_RetrieveCustomer)
	  servicioRetrieveCustomer.PSP_RequerimientoStruct_RetrieveCustomer = stgResquest

	  err := service.client.Call(CONSTANTS.RETRIEVE_CUSTOMER, servicioRetrieveCustomer, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_RetrieveCustomer)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_RetrieveCustomer, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_RetrieveCustomer)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_RetrieveCustomer)

	  servicioRetrieveCustomer := new(sndservice.ServiceStruct_RetrieveCustomer)
	  servicioRetrieveCustomer.PSP_RequerimientoStruct_RetrieveCustomer = sndResquest

	  err := service.client.Call(CONSTANTS.RETRIEVE_CUSTOMER, servicioRetrieveCustomer, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_RetrieveCustomer)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_RetrieveCustomer, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) UpdateCustomer(request *RequerimientoStruct_UpdateCustomer) (*RespuestaStruct_UpdateCustomer, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_UpdateCustomer)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_UpdateCustomer)

	  if prodResquest.Psp_Address != nil {
		prodResquest.Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if prodResquest.Psp_Person != nil {
		prodResquest.Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }
	  if prodResquest.Psp_PaymentMethod != nil && 
             prodResquest.Psp_PaymentMethod.Person != nil {
		prodResquest.Psp_PaymentMethod.Person.XMLName = xml.Name{"", "Person"}
	  }
	  if prodResquest.Psp_PaymentMethod != nil && 
             prodResquest.Psp_PaymentMethod.Address != nil {
		prodResquest.Psp_PaymentMethod.Address.XMLName = xml.Name{"", "Address"}
	  }

	  servicioUpdateCustomer := new(prodservice.ServiceStruct_UpdateCustomer)
	  servicioUpdateCustomer.PSP_RequerimientoStruct_UpdateCustomer = prodResquest

	  err := service.client.Call(CONSTANTS.UPDATE_CUSTOMER, servicioUpdateCustomer, prodResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_UpdateCustomer)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_UpdateCustomer, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_UpdateCustomer)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_UpdateCustomer)

	  if stgResquest.Psp_Address != nil {
		stgResquest.Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if stgResquest.Psp_Person != nil {
		stgResquest.Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }
	  if stgResquest.Psp_PaymentMethod != nil && 
             stgResquest.Psp_PaymentMethod.Person != nil {
		stgResquest.Psp_PaymentMethod.Person.XMLName = xml.Name{"", "Person"}
	  }
	  if stgResquest.Psp_PaymentMethod != nil && 
             stgResquest.Psp_PaymentMethod.Address != nil {
		stgResquest.Psp_PaymentMethod.Address.XMLName = xml.Name{"", "Address"}
	  }

	  servicioUpdateCustomer := new(stgservice.ServiceStruct_UpdateCustomer)
	  servicioUpdateCustomer.PSP_RequerimientoStruct_UpdateCustomer = stgResquest

	  err := service.client.Call(CONSTANTS.UPDATE_CUSTOMER, servicioUpdateCustomer, stgResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_UpdateCustomer)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_UpdateCustomer, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_UpdateCustomer)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_UpdateCustomer)

	  if sndResquest.Psp_Address != nil {
		sndResquest.Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	  }
	  if sndResquest.Psp_Person != nil {
		sndResquest.Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	  }
	  if sndResquest.Psp_PaymentMethod != nil && 
             sndResquest.Psp_PaymentMethod.Person != nil {
		sndResquest.Psp_PaymentMethod.Person.XMLName = xml.Name{"", "Person"}
	  }
	  if sndResquest.Psp_PaymentMethod != nil && 
             sndResquest.Psp_PaymentMethod.Address != nil {
		sndResquest.Psp_PaymentMethod.Address.XMLName = xml.Name{"", "Address"}
	  }

	  servicioUpdateCustomer := new(sndservice.ServiceStruct_UpdateCustomer)
	  servicioUpdateCustomer.PSP_RequerimientoStruct_UpdateCustomer = sndResquest

	  err := service.client.Call(CONSTANTS.UPDATE_CUSTOMER, servicioUpdateCustomer, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_UpdateCustomer)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_UpdateCustomer, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) DeleteCustomer(request *RequerimientoStruct_DeleteCustomer) (*RespuestaStruct_DeleteCustomer, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_DeleteCustomer)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_DeleteCustomer)

	  servicioDeleteCustomer := new(prodservice.ServiceStruct_DeleteCustomer)
	  servicioDeleteCustomer.PSP_RequerimientoStruct_DeleteCustomer = prodResquest

	  err := service.client.Call(CONSTANTS.DELETE_CUSTOMER, servicioDeleteCustomer, prodResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_DeleteCustomer)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_DeleteCustomer, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_DeleteCustomer)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_DeleteCustomer)

	  servicioDeleteCustomer := new(stgservice.ServiceStruct_DeleteCustomer)
	  servicioDeleteCustomer.PSP_RequerimientoStruct_DeleteCustomer = stgResquest

	  err := service.client.Call(CONSTANTS.DELETE_CUSTOMER, servicioDeleteCustomer, stgResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_DeleteCustomer)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_DeleteCustomer, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_DeleteCustomer)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_DeleteCustomer)

	  servicioDeleteCustomer := new(sndservice.ServiceStruct_DeleteCustomer)
	  servicioDeleteCustomer.PSP_RequerimientoStruct_DeleteCustomer = sndResquest

	  err := service.client.Call(CONSTANTS.DELETE_CUSTOMER, servicioDeleteCustomer, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_DeleteCustomer)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_DeleteCustomer, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) SimpleQueryTx(request *RequerimientoStruct_SimpleQueryTx) (*RespuestaStruct_SimpleQueryTx, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_SimpleQueryTx)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_SimpleQueryTx)

	  servicioSimpleQueryTx := new(prodservice.ServiceStruct_SimpleQueryTx)
	  servicioSimpleQueryTx.PSP_RequerimientoStruct_SimpleQueryTx = prodResquest

	  err := service.client.Call(CONSTANTS.SIMPLE_QUERY_TX, servicioSimpleQueryTx, prodResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_SimpleQueryTx)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_SimpleQueryTx, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_SimpleQueryTx)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_SimpleQueryTx)

	  servicioSimpleQueryTx := new(stgservice.ServiceStruct_SimpleQueryTx)
	  servicioSimpleQueryTx.PSP_RequerimientoStruct_SimpleQueryTx = stgResquest

	  err := service.client.Call(CONSTANTS.SIMPLE_QUERY_TX, servicioSimpleQueryTx, stgResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_SimpleQueryTx)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_SimpleQueryTx, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_SimpleQueryTx)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_SimpleQueryTx)

	  servicioSimpleQueryTx := new(sndservice.ServiceStruct_SimpleQueryTx)
	  servicioSimpleQueryTx.PSP_RequerimientoStruct_SimpleQueryTx = sndResquest

	  err := service.client.Call(CONSTANTS.SIMPLE_QUERY_TX, servicioSimpleQueryTx, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_SimpleQueryTx)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_SimpleQueryTx, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) QueryCardNumber(request *RequerimientoStruct_QueryCardNumber) (*RespuestaStruct_QueryCardNumber, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_QueryCardNumber)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_QueryCardNumber)

	  servicioQueryCardNumber := new(prodservice.ServiceStruct_QueryCardNumber)
	  servicioQueryCardNumber.PSP_RequerimientoStruct_QueryCardNumber = prodResquest

	  err := service.client.Call(CONSTANTS.QUERY_CARD_NUMBER, servicioQueryCardNumber, prodResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_QueryCardNumber)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_QueryCardNumber, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_QueryCardNumber)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_QueryCardNumber)

	  servicioQueryCardNumber := new(stgservice.ServiceStruct_QueryCardNumber)
	  servicioQueryCardNumber.PSP_RequerimientoStruct_QueryCardNumber = stgResquest

	  err := service.client.Call(CONSTANTS.QUERY_CARD_NUMBER, servicioQueryCardNumber, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_QueryCardNumber)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_QueryCardNumber, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_QueryCardNumber)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_QueryCardNumber)

	  servicioQueryCardNumber := new(sndservice.ServiceStruct_QueryCardNumber)
	  servicioQueryCardNumber.PSP_RequerimientoStruct_QueryCardNumber = sndResquest

	  err := service.client.Call(CONSTANTS.QUERY_CARD_NUMBER, servicioQueryCardNumber, sndResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_QueryCardNumber)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_QueryCardNumber, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) QueryCardDetails(request *RequerimientoStruct_QueryCardDetails) (*RespuestaStruct_QueryCardDetails, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_QueryCardDetails)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_QueryCardDetails)

	  servicioQueryCardDetails := new(prodservice.ServiceStruct_QueryCardDetails)
	  servicioQueryCardDetails.PSP_RequerimientoStruct_QueryCardDetails = prodResquest

	  err := service.client.Call(CONSTANTS.QUERY_CARD_DETAILS, servicioQueryCardDetails, prodResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_QueryCardDetails)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_QueryCardDetails, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_QueryCardDetails)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_QueryCardDetails)

	  servicioQueryCardDetails := new(stgservice.ServiceStruct_QueryCardDetails)
	  servicioQueryCardDetails.PSP_RequerimientoStruct_QueryCardDetails = stgResquest

	  err := service.client.Call(CONSTANTS.QUERY_CARD_DETAILS, servicioQueryCardDetails, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_QueryCardDetails)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_QueryCardDetails, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_QueryCardDetails)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_QueryCardDetails)

	  servicioQueryCardDetails := new(sndservice.ServiceStruct_QueryCardDetails)
	  servicioQueryCardDetails.PSP_RequerimientoStruct_QueryCardDetails = sndResquest

	  err := service.client.Call(CONSTANTS.QUERY_CARD_DETAILS, servicioQueryCardDetails, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_QueryCardDetails)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_QueryCardDetails, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) QueryTxs(request *RequerimientoStruct_QueryTxs) (*RespuestaStruct_QueryTxs, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_QueryTxs)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_QueryTxs)

	  servicioQueryTxs := new(prodservice.ServiceStruct_QueryTxs)
	  servicioQueryTxs.PSP_RequerimientoStruct_QueryTxs = prodResquest

	  err := service.client.Call(CONSTANTS.QUERY_TXS, servicioQueryTxs, prodResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_QueryTxs)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_QueryTxs, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_QueryTxs)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_QueryTxs)

	  servicioQueryTxs := new(stgservice.ServiceStruct_QueryTxs)
	  servicioQueryTxs.PSP_RequerimientoStruct_QueryTxs = stgResquest

	  err := service.client.Call(CONSTANTS.QUERY_TXS, servicioQueryTxs, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_QueryTxs)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_QueryTxs, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_QueryTxs)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_QueryTxs)

	  servicioQueryTxs := new(sndservice.ServiceStruct_QueryTxs)
	  servicioQueryTxs.PSP_RequerimientoStruct_QueryTxs = sndResquest

	  err := service.client.Call(CONSTANTS.QUERY_TXS, servicioQueryTxs, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_QueryTxs)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_QueryTxs, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) GetIINDetails(request *RequerimientoStruct_GetIINDetails) (*RespuestaStruct_GetIINDetails, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_GetIINDetails)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_GetIINDetails)

	  servicioGetIINDetails := new(prodservice.ServiceStruct_GetIINDetails)
	  servicioGetIINDetails.PSP_RequerimientoStruct_GetIINDetails = prodResquest

	  err := service.client.Call(CONSTANTS.GET_IIN_DETAILS, servicioGetIINDetails, prodResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_GetIINDetails)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_GetIINDetails, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_GetIINDetails)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_GetIINDetails)

	  servicioGetIINDetails := new(stgservice.ServiceStruct_GetIINDetails)
	  servicioGetIINDetails.PSP_RequerimientoStruct_GetIINDetails = stgResquest

	  err := service.client.Call(CONSTANTS.GET_IIN_DETAILS, servicioGetIINDetails, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_GetIINDetails)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_GetIINDetails, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_GetIINDetails)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_GetIINDetails)

	  servicioGetIINDetails := new(sndservice.ServiceStruct_GetIINDetails)
	  servicioGetIINDetails.PSP_RequerimientoStruct_GetIINDetails = sndResquest

	  err := service.client.Call(CONSTANTS.GET_IIN_DETAILS, servicioGetIINDetails, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_GetIINDetails)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_GetIINDetails, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) ChangeSecretKey(request *RequerimientoStruct_ChangeSecretKey) (*RespuestaStruct_ChangeSecretKey, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_ChangeSecretKey)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_ChangeSecretKey)

	  servicioChangeSecretKey := new(prodservice.ServiceStruct_ChangeSecretKey)
	  servicioChangeSecretKey.PSP_RequerimientoStruct_ChangeSecretKey = prodResquest

	  err := service.client.Call(CONSTANTS.CHANGE_SECRET_KEY, servicioChangeSecretKey, prodResponse)

	  if err != nil {
		return nil, err
	  }
	  response := new(RespuestaStruct_ChangeSecretKey)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_ChangeSecretKey, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_ChangeSecretKey)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_ChangeSecretKey)

	  servicioChangeSecretKey := new(stgservice.ServiceStruct_ChangeSecretKey)
	  servicioChangeSecretKey.PSP_RequerimientoStruct_ChangeSecretKey = stgResquest

	  err := service.client.Call(CONSTANTS.CHANGE_SECRET_KEY, servicioChangeSecretKey, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_ChangeSecretKey)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_ChangeSecretKey, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_ChangeSecretKey)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_ChangeSecretKey)

	  servicioChangeSecretKey := new(sndservice.ServiceStruct_ChangeSecretKey)
	  servicioChangeSecretKey.PSP_RequerimientoStruct_ChangeSecretKey = sndResquest

	  err := service.client.Call(CONSTANTS.CHANGE_SECRET_KEY, servicioChangeSecretKey, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_ChangeSecretKey)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_ChangeSecretKey, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) FraudScreening(request *RequerimientoStruct_FraudScreening) (*RespuestaStruct_FraudScreening, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_FraudScreening)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_FraudScreening)

	  servicioFraudScreening := new(prodservice.ServiceStruct_FraudScreening)
	  servicioFraudScreening.PSP_RequerimientoStruct_FraudScreening = prodResquest

	  err := service.client.Call(CONSTANTS.FRAUD_SCREENING, servicioFraudScreening, prodResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_FraudScreening)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_FraudScreening, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_FraudScreening)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_FraudScreening)

	  servicioFraudScreening := new(stgservice.ServiceStruct_FraudScreening)
	  servicioFraudScreening.PSP_RequerimientoStruct_FraudScreening = stgResquest

	  err := service.client.Call(CONSTANTS.FRAUD_SCREENING, servicioFraudScreening, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_FraudScreening)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_FraudScreening, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_FraudScreening)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_FraudScreening)

	  servicioFraudScreening := new(sndservice.ServiceStruct_FraudScreening)
	  servicioFraudScreening.PSP_RequerimientoStruct_FraudScreening = sndResquest

	  err := service.client.Call(CONSTANTS.FRAUD_SCREENING, servicioFraudScreening, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_FraudScreening)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_FraudScreening, response)

	  return response, nil

        }
        
        return nil, nil
}

func (service *PaymentServicePlatformPortType) NotifyFraudScreeningReview(request *RequerimientoStruct_NotifyFraudScreeningReview) (*RespuestaStruct_NotifyFraudScreeningReview, error) {
        switch Configuration.environment {
	case CONSTANTS.PRODUCTION_ENV:
	  prodResquest := new(prodservice.RequerimientoStruct_NotifyFraudScreeningReview)
	  CopypStruct(request, prodResquest)

	  prodResponse := new(prodservice.ServiceRespuestaStruct_NotifyFraudScreeningReview)

	  servicioNotifyFraudScreeningReview := new(prodservice.ServiceStruct_NotifyFraudScreeningReview)
	  servicioNotifyFraudScreeningReview.PSP_RequerimientoStruct_NotifyFraudScreeningReview = prodResquest

	  err := service.client.Call(CONSTANTS.NOTIFY_FRAUD_SCREENING_REVIEW, servicioNotifyFraudScreeningReview, prodResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_NotifyFraudScreeningReview)
	  CopypStruct(prodResponse.PSP_RespuestaStruct_NotifyFraudScreeningReview, response)

	  return response, nil

	case CONSTANTS.STAGING_ENV:
	  stgResquest := new(stgservice.RequerimientoStruct_NotifyFraudScreeningReview)
	  CopypStruct(request, stgResquest)

	  stgResponse := new(stgservice.ServiceRespuestaStruct_NotifyFraudScreeningReview)

	  servicioNotifyFraudScreeningReview := new(stgservice.ServiceStruct_NotifyFraudScreeningReview)
	  servicioNotifyFraudScreeningReview.PSP_RequerimientoStruct_NotifyFraudScreeningReview = stgResquest

	  err := service.client.Call(CONSTANTS.NOTIFY_FRAUD_SCREENING_REVIEW, servicioNotifyFraudScreeningReview, stgResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_NotifyFraudScreeningReview)
	  CopypStruct(stgResponse.PSP_RespuestaStruct_NotifyFraudScreeningReview, response)

	  return response, nil

	case CONSTANTS.SANDBOX_ENV:
	  sndResquest := new(sndservice.RequerimientoStruct_NotifyFraudScreeningReview)
	  CopypStruct(request, sndResquest)

	  sndResponse := new(sndservice.ServiceRespuestaStruct_NotifyFraudScreeningReview)

	  servicioNotifyFraudScreeningReview := new(sndservice.ServiceStruct_NotifyFraudScreeningReview)
	  servicioNotifyFraudScreeningReview.PSP_RequerimientoStruct_NotifyFraudScreeningReview = sndResquest

	  err := service.client.Call(CONSTANTS.NOTIFY_FRAUD_SCREENING_REVIEW, servicioNotifyFraudScreeningReview, sndResponse)

	  if err != nil {
		return nil, err
	  }

	  response := new(RespuestaStruct_NotifyFraudScreeningReview)
	  CopypStruct(sndResponse.PSP_RespuestaStruct_NotifyFraudScreeningReview, response)

	  return response, nil

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
