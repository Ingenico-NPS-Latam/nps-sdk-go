package nps

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
	CONSTANTS "nps/constants"
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

	response := new(ServiceRespuestaStruct_PayOnLine_2p)

	servicioPayOnLine_2p := new(ServiceStruct_PayOnLine_2p)
	servicioPayOnLine_2p.PSP_RequerimientoStruct_PayOnLine_2p = request

	err := service.client.Call(CONSTANTS.PAY_ONLINE_2P, servicioPayOnLine_2p, response)
	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_PayOnLine_2p, nil
}

func (service *PaymentServicePlatformPortType) Authorize_2p(request *RequerimientoStruct_Authorize_2p) (*RespuestaStruct_Authorize_2p, error) {
	response := new(ServiceRespuestaStruct_Authorize_2p)

	servicioAuthorize_2p := new(ServiceStruct_Authorize_2p)
	servicioAuthorize_2p.PSP_RequerimientoStruct_Authorize_2p = request

	err := service.client.Call(CONSTANTS.AUTHORIZE_2P, servicioAuthorize_2p, response)
	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_Authorize_2p, nil
}

func (service *PaymentServicePlatformPortType) BankPayment_2p(request *RequerimientoStruct_BankPayment_2p) (*RespuestaStruct_BankPayment_2p, error) {
	response := new(ServiceRespuestaStruct_BankPayment_2p)

	servicioBankPayment_2p := new(ServiceStruct_BankPayment_2p)
	servicioBankPayment_2p.PSP_RequerimientoStruct_BankPayment_2p = request

	err := service.client.Call(CONSTANTS.BANK_PAYMENT_2P, servicioBankPayment_2p, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_BankPayment_2p, nil
}

func (service *PaymentServicePlatformPortType) SplitPayOnLine_2p(request *RequerimientoStruct_SplitPayOnLine_2p) (*RespuestaStruct_SplitPayOnLine_2p, error) {
	response := new(ServiceRespuestaStruct_SplitPayOnLine_2p)

	servicioSplitPayOnLine_2p := new(ServiceStruct_SplitPayOnLine_2p)
	servicioSplitPayOnLine_2p.PSP_RequerimientoStruct_SplitPayOnLine_2p = request

	err := service.client.Call(CONSTANTS.SPLIT_PAY_ONLINE_2P, servicioSplitPayOnLine_2p, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_SplitPayOnLine_2p, nil
}

func (service *PaymentServicePlatformPortType) SplitAuthorize_2p(request *RequerimientoStruct_SplitAuthorize_2p) (*RespuestaStruct_SplitAuthorize_2p, error) {
	response := new(ServiceRespuestaStruct_SplitAuthorize_2p)

	servicioSplitAuthorize_2p := new(ServiceStruct_SplitAuthorize_2p)
	servicioSplitAuthorize_2p.PSP_RequerimientoStruct_SplitAuthorize_2p = request

	err := service.client.Call(CONSTANTS.SPLIT_AUTHORIZE_2P, servicioSplitAuthorize_2p, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_SplitAuthorize_2p, nil
}

func (service *PaymentServicePlatformPortType) PayOnLine_3p(request *RequerimientoStruct_PayOnLine_3p) (*RespuestaStruct_PayOnLine_3p, error) {
	response := new(ServiceRespuestaStruct_PayOnLine_3p)

	servicioPayOnLine_3p := new(ServiceStruct_PayOnLine_3p)
	servicioPayOnLine_3p.PSP_RequerimientoStruct_PayOnLine_3p = request

	err := service.client.Call(CONSTANTS.PAY_ONLINE_3P, servicioPayOnLine_3p, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_PayOnLine_3p, nil
}

func (service *PaymentServicePlatformPortType) SplitPayOnLine_3p(request *RequerimientoStruct_SplitPayOnLine_3p) (*RespuestaStruct_SplitPayOnLine_3p, error) {
	response := new(ServiceRespuestaStruct_SplitPayOnLine_3p)

	servicioSplitPayOnLine_3p := new(ServiceStruct_SplitPayOnLine_3p)
	servicioSplitPayOnLine_3p.PSP_RequerimientoStruct_SplitPayOnLine_3p = request

	err := service.client.Call(CONSTANTS.SPLIT_PAY_ONLINE_3P, servicioSplitPayOnLine_3p, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_SplitPayOnLine_3p, nil
}

func (service *PaymentServicePlatformPortType) Authorize_3p(request *RequerimientoStruct_Authorize_3p) (*RespuestaStruct_Authorize_3p, error) {
	response := new(ServiceRespuestaStruct_Authorize_3p)

	servicioAuthorize_3p := new(ServiceStruct_Authorize_3p)
	servicioAuthorize_3p.PSP_RequerimientoStruct_Authorize_3p = request

	err := service.client.Call(CONSTANTS.AUTHORIZE_3P, servicioAuthorize_3p, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_Authorize_3p, nil
}

func (service *PaymentServicePlatformPortType) SplitAuthorize_3p(request *RequerimientoStruct_SplitAuthorize_3p) (*RespuestaStruct_SplitAuthorize_3p, error) {
	response := new(ServiceRespuestaStruct_SplitAuthorize_3p)

	serviceSplitAuthorize_3p := new(ServiceStruct_SplitAuthorize_3p)
	serviceSplitAuthorize_3p.PSP_RequerimientoStruct_SplitAuthorize_3p = request

	err := service.client.Call(CONSTANTS.SPLIT_AUTHORIZE_3P, serviceSplitAuthorize_3p, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_SplitAuthorize_3p, nil
}

func (service *PaymentServicePlatformPortType) BankPayment_3p(request *RequerimientoStruct_BankPayment_3p) (*RespuestaStruct_BankPayment_3p, error) {
	response := new(ServiceRespuestaStruct_BankPayment_3p)

	servicioBankPayment_3p := new(ServiceStruct_BankPayment_3p)
	servicioBankPayment_3p.PSP_RequerimientoStruct_BankPayment_3p = request

	err := service.client.Call(CONSTANTS.BANK_PAYMENT_3P, servicioBankPayment_3p, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_BankPayment_3p, nil
}

func (service *PaymentServicePlatformPortType) CashPayment_3p(request *RequerimientoStruct_CashPayment_3p) (*RespuestaStruct_CashPayment_3p, error) {
	response := new(ServiceRespuestaStruct_CashPayment_3p)

	servicioCashPayment_3p := new(ServiceStruct_CashPayment_3p)
	servicioCashPayment_3p.PSP_RequerimientoStruct_CashPayment_3p = request

	err := service.client.Call(CONSTANTS.CASH_PAYMENT_3P, servicioCashPayment_3p, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_CashPayment_3p, nil
}

func (service *PaymentServicePlatformPortType) Capture(request *RequerimientoStruct_Capture) (*RespuestaStruct_Capture, error) {
	response := new(ServiceRespuestaStruct_Capture)

	servicioCapture := new(ServiceStruct_Capture)
	servicioCapture.PSP_RequerimientoStruct_Capture = request

	err := service.client.Call(CONSTANTS.CAPTURE, servicioCapture, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_Capture, nil
}

func (service *PaymentServicePlatformPortType) Refund(request *RequerimientoStruct_Refund) (*RespuestaStruct_Refund, error) {
	response := new(ServiceRespuestaStruct_Refund)

	servicioRefund := new(ServiceStruct_Refund)
	servicioRefund.PSP_RequerimientoStruct_Refund = request

	err := service.client.Call(CONSTANTS.REFUND, servicioRefund, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_Refund, nil
}

func (service *PaymentServicePlatformPortType) GetInstallmentsOptions(request *RequerimientoStruct_GetInstallmentsOptions) (*RespuestaStruct_GetInstallmentsOptions, error) {
	response := new(ServiceRespuestaStruct_GetInstallmentsOptions)

	servicioGetInstallmentsOptions := new(ServiceStruct_GetInstallmentsOptions)
	servicioGetInstallmentsOptions.PSP_RequerimientoStruct_GetInstallmentsOptions = request

	err := service.client.Call(CONSTANTS.GET_INSTALLMENTS_OPTIONS, servicioGetInstallmentsOptions, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_GetInstallmentsOptions, nil
}

func (service *PaymentServicePlatformPortType) CreatePaymentMethodToken(request *RequerimientoStruct_CreatePaymentMethodToken) (*RespuestaStruct_CreatePaymentMethodToken, error) {
	response := new(ServiceRespuestaStruct_CreatePaymentMethodToken)

	if request.Psp_CardInputDetails != nil {
		request.Psp_CardInputDetails.XMLName = xml.Name{"", "psp_CardInputDetails"}
	}
	if request.Psp_Address != nil {
		request.Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	}
	if request.Psp_Person != nil {
		request.Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	}

	servicioCreatePaymentMethodToken := new(ServiceStruct_CreatePaymentMethodToken)
	servicioCreatePaymentMethodToken.PSP_RequerimientoStruct_CreatePaymentMethodToken = request

	err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD_TOKEN, servicioCreatePaymentMethodToken, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_CreatePaymentMethodToken, nil
}

func (service *PaymentServicePlatformPortType) RetrievePaymentMethodToken(request *RequerimientoStruct_RetrievePaymentMethodToken) (*RespuestaStruct_RetrievePaymentMethodToken, error) {
	response := new(ServiceRespuestaStruct_RetrievePaymentMethodToken)

	servicioRetrievePaymentMethodToken := new(ServiceStruct_RetrievePaymentMethodToken)
	servicioRetrievePaymentMethodToken.PSP_RequerimientoStruct_RetrievePaymentMethodToken = request

	err := service.client.Call(CONSTANTS.RETRIEVE_PAYMENT_METHOD_TOKEN, servicioRetrievePaymentMethodToken, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_RetrievePaymentMethodToken, nil
}

func (service *PaymentServicePlatformPortType) RecachePaymentMethodToken(request *RequerimientoStruct_RecachePaymentMethodToken) (*RespuestaStruct_RecachePaymentMethodToken, error) {
	response := new(ServiceRespuestaStruct_RecachePaymentMethodToken)

	if request.Psp_Address != nil {
		request.Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	}
	if request.Psp_Person != nil {
		request.Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	}

	servicioRecachePaymentMethodToken := new(ServiceStruct_RecachePaymentMethodToken)
	servicioRecachePaymentMethodToken.PSP_RequerimientoStruct_RecachePaymentMethodToken = request

	err := service.client.Call(CONSTANTS.RECACHE_PAYMENT_METHOD_TOKEN, servicioRecachePaymentMethodToken, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_RecachePaymentMethodToken, nil
}

func (service *PaymentServicePlatformPortType) CreateClientSession(request *RequerimientoStruct_CreateClientSession) (*RespuestaStruct_CreateClientSession, error) {
	response := new(ServiceRespuestaStruct_CreateClientSession)

	servicioCreateClientSession := new(ServiceStruct_CreateClientSession)
	servicioCreateClientSession.PSP_RequerimientoStruct_CreateClientSession = request

	err := service.client.Call(CONSTANTS.CREATE_CLIENT_SESSION, servicioCreateClientSession, response)
	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_CreateClientSession, nil
}

func (service *PaymentServicePlatformPortType) CreatePaymentMethod(request *RequerimientoStruct_CreatePaymentMethod) (*RespuestaStruct_CreatePaymentMethod, error) {
	response := new(ServiceRespuestaStruct_CreatePaymentMethod)

	servicioCreatePaymentMethod := new(ServiceStruct_CreatePaymentMethod)
	servicioCreatePaymentMethod.PSP_RequerimientoStruct_CreatePaymentMethod = request

	err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD, servicioCreatePaymentMethod, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_CreatePaymentMethod, nil
}

func (service *PaymentServicePlatformPortType) CreatePaymentMethodFromPayment(request *RequerimientoStruct_CreatePaymentMethodFromPayment) (*RespuestaStruct_CreatePaymentMethodFromPayment, error) {
	response := new(ServiceRespuestaStruct_CreatePaymentMethodFromPayment)

	servicioCreatePaymentMethodFromPayment := new(ServiceStruct_CreatePaymentMethodFromPayment)
	servicioCreatePaymentMethodFromPayment.PSP_RequerimientoStruct_CreatePaymentMethodFromPayment = request

	err := service.client.Call(CONSTANTS.CREATE_PAYMENT_METHOD_FROM_PAYMENT, servicioCreatePaymentMethodFromPayment, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_CreatePaymentMethodFromPayment, nil
}

func (service *PaymentServicePlatformPortType) RetrievePaymentMethod(request *RequerimientoStruct_RetrievePaymentMethod) (*RespuestaStruct_RetrievePaymentMethod, error) {
	response := new(ServiceRespuestaStruct_RetrievePaymentMethod)

	servicioRetrievePaymentMethod := new(ServiceStruct_RetrievePaymentMethod)
	servicioRetrievePaymentMethod.PSP_RequerimientoStruct_RetrievePaymentMethod = request

	err := service.client.Call(CONSTANTS.RETRIEVE_PAYMENT_METHOD, servicioRetrievePaymentMethod, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_RetrievePaymentMethod, nil
}

func (service *PaymentServicePlatformPortType) UpdatePaymentMethod(request *RequerimientoStruct_UpdatePaymentMethod) (*RespuestaStruct_UpdatePaymentMethod, error) {
	response := new(ServiceRespuestaStruct_UpdatePaymentMethod)

	servicioUpdatePaymentMethod := new(ServiceStruct_UpdatePaymentMethod)
	servicioUpdatePaymentMethod.PSP_RequerimientoStruct_UpdatePaymentMethod = request

	err := service.client.Call(CONSTANTS.UPDATE_PAYMENT_METHOD, servicioUpdatePaymentMethod, response)
	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_UpdatePaymentMethod, nil
}

func (service *PaymentServicePlatformPortType) DeletePaymentMethod(request *RequerimientoStruct_DeletePaymentMethod) (*RespuestaStruct_DeletePaymentMethod, error) {
	response := new(ServiceRespuestaStruct_DeletePaymentMethod)

	servicioDeletePaymentMethod := new(ServiceStruct_DeletePaymentMethod)
	servicioDeletePaymentMethod.PSP_RequerimientoStruct_DeletePaymentMethod = request

	err := service.client.Call(CONSTANTS.DELETE_PAYMENT_METHOD, servicioDeletePaymentMethod, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_DeletePaymentMethod, nil
}

func (service *PaymentServicePlatformPortType) CreateCustomer(request *RequerimientoStruct_CreateCustomer) (*RespuestaStruct_CreateCustomer, error) {
	response := new(ServiceRespuestaStruct_CreateCustomer)

	if request.Psp_Address != nil {
		request.Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	}
	if request.Psp_Person != nil {
		request.Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	}

	if request.Psp_PaymentMethod != nil && request.Psp_PaymentMethod.Person != nil {
		request.Psp_PaymentMethod.Person.XMLName = xml.Name{"", "Person"}
	}

	if request.Psp_PaymentMethod != nil && request.Psp_PaymentMethod.Address != nil {
		request.Psp_PaymentMethod.Address.XMLName = xml.Name{"", "Address"}
	}

	servicioCreateCustomer := new(ServiceStruct_CreateCustomer)
	servicioCreateCustomer.PSP_RequerimientoStruct_CreateCustomer = request

	err := service.client.Call(CONSTANTS.CREATE_CUSTOMER, servicioCreateCustomer, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_CreateCustomer, nil
}

func (service *PaymentServicePlatformPortType) RetrieveCustomer(request *RequerimientoStruct_RetrieveCustomer) (*RespuestaStruct_RetrieveCustomer, error) {
	response := new(ServiceRespuestaStruct_RetrieveCustomer)

	servicioRetrieveCustomer := new(ServiceStruct_RetrieveCustomer)
	servicioRetrieveCustomer.PSP_RequerimientoStruct_RetrieveCustomer = request

	err := service.client.Call(CONSTANTS.RETRIEVE_CUSTOMER, servicioRetrieveCustomer, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_RetrieveCustomer, nil
}

func (service *PaymentServicePlatformPortType) UpdateCustomer(request *RequerimientoStruct_UpdateCustomer) (*RespuestaStruct_UpdateCustomer, error) {
	response := new(ServiceRespuestaStruct_UpdateCustomer)

	if request.Psp_Address != nil {
		request.Psp_Address.XMLName = xml.Name{"", "psp_Address"}
	}
	if request.Psp_Person != nil {
		request.Psp_Person.XMLName = xml.Name{"", "psp_Person"}
	}
	if request.Psp_PaymentMethod != nil && request.Psp_PaymentMethod.Person != nil {
		request.Psp_PaymentMethod.Person.XMLName = xml.Name{"", "Person"}
	}

	if request.Psp_PaymentMethod != nil && request.Psp_PaymentMethod.Address != nil {
		request.Psp_PaymentMethod.Address.XMLName = xml.Name{"", "Address"}
	}

	servicioUpdateCustomer := new(ServiceStruct_UpdateCustomer)
	servicioUpdateCustomer.PSP_RequerimientoStruct_UpdateCustomer = request

	err := service.client.Call(CONSTANTS.UPDATE_CUSTOMER, servicioUpdateCustomer, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_UpdateCustomer, nil
}

func (service *PaymentServicePlatformPortType) DeleteCustomer(request *RequerimientoStruct_DeleteCustomer) (*RespuestaStruct_DeleteCustomer, error) {
	response := new(ServiceRespuestaStruct_DeleteCustomer)

	servicioDeleteCustomer := new(ServiceStruct_DeleteCustomer)
	servicioDeleteCustomer.PSP_RequerimientoStruct_DeleteCustomer = request

	err := service.client.Call(CONSTANTS.DELETE_CUSTOMER, servicioDeleteCustomer, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_DeleteCustomer, nil
}

func (service *PaymentServicePlatformPortType) SimpleQueryTx(request *RequerimientoStruct_SimpleQueryTx) (*RespuestaStruct_SimpleQueryTx, error) {
	response := new(ServiceRespuestaStruct_SimpleQueryTx)

	servicioSimpleQueryTx := new(ServiceStruct_SimpleQueryTx)
	servicioSimpleQueryTx.PSP_RequerimientoStruct_SimpleQueryTx = request

	err := service.client.Call(CONSTANTS.SIMPLE_QUERY_TX, servicioSimpleQueryTx, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_SimpleQueryTx, nil
}

func (service *PaymentServicePlatformPortType) QueryCardNumber(request *RequerimientoStruct_QueryCardNumber) (*RespuestaStruct_QueryCardNumber, error) {
	response := new(ServiceRespuestaStruct_QueryCardNumber)

	servicioQueryCardNumber := new(ServiceStruct_QueryCardNumber)
	servicioQueryCardNumber.PSP_RequerimientoStruct_QueryCardNumber = request

	err := service.client.Call(CONSTANTS.QUERY_CARD_NUMBER, servicioQueryCardNumber, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_QueryCardNumber, nil
}

func (service *PaymentServicePlatformPortType) QueryCardDetails(request *RequerimientoStruct_QueryCardDetails) (*RespuestaStruct_QueryCardDetails, error) {
	response := new(ServiceRespuestaStruct_QueryCardDetails)

	servicioQueryCardDetails := new(ServiceStruct_QueryCardDetails)
	servicioQueryCardDetails.PSP_RequerimientoStruct_QueryCardDetails = request

	err := service.client.Call(CONSTANTS.QUERY_CARD_DETAILS, servicioQueryCardDetails, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_QueryCardDetails, nil
}

func (service *PaymentServicePlatformPortType) QueryTxs(request *RequerimientoStruct_QueryTxs) (*RespuestaStruct_QueryTxs, error) {
	response := new(ServiceRespuestaStruct_QueryTxs)

	servicioQueryTxs := new(ServiceStruct_QueryTxs)
	servicioQueryTxs.PSP_RequerimientoStruct_QueryTxs = request

	err := service.client.Call(CONSTANTS.QUERY_TXS, servicioQueryTxs, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_QueryTxs, nil
}

func (service *PaymentServicePlatformPortType) GetIINDetails(request *RequerimientoStruct_GetIINDetails) (*RespuestaStruct_GetIINDetails, error) {
	response := new(ServiceRespuestaStruct_GetIINDetails)

	servicioGetIINDetails := new(ServiceStruct_GetIINDetails)
	servicioGetIINDetails.PSP_RequerimientoStruct_GetIINDetails = request

	err := service.client.Call(CONSTANTS.GET_IIN_DETAILS, servicioGetIINDetails, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_GetIINDetails, nil
}

func (service *PaymentServicePlatformPortType) ChangeSecretKey(request *RequerimientoStruct_ChangeSecretKey) (*RespuestaStruct_ChangeSecretKey, error) {
	response := new(ServiceRespuestaStruct_ChangeSecretKey)

	servicioChangeSecretKey := new(ServiceStruct_ChangeSecretKey)
	servicioChangeSecretKey.PSP_RequerimientoStruct_ChangeSecretKey = request

	err := service.client.Call(CONSTANTS.CHANGE_SECRET_KEY, servicioChangeSecretKey, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_ChangeSecretKey, nil
}

func (service *PaymentServicePlatformPortType) FraudScreening(request *RequerimientoStruct_FraudScreening) (*RespuestaStruct_FraudScreening, error) {
	response := new(ServiceRespuestaStruct_FraudScreening)

	servicioFraudScreening := new(ServiceStruct_FraudScreening)
	servicioFraudScreening.PSP_RequerimientoStruct_FraudScreening = request

	err := service.client.Call(CONSTANTS.FRAUD_SCREENING, servicioFraudScreening, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_FraudScreening, nil
}

func (service *PaymentServicePlatformPortType) NotifyFraudScreeningReview(request *RequerimientoStruct_NotifyFraudScreeningReview) (*RespuestaStruct_NotifyFraudScreeningReview, error) {
	response := new(ServiceRespuestaStruct_NotifyFraudScreeningReview)

	servicioNotifyFraudScreeningReview := new(ServiceStruct_NotifyFraudScreeningReview)
	servicioNotifyFraudScreeningReview.PSP_RequerimientoStruct_NotifyFraudScreeningReview = request

	err := service.client.Call(CONSTANTS.NOTIFY_FRAUD_SCREENING_REVIEW, servicioNotifyFraudScreeningReview, response)

	if err != nil {
		return nil, err
	}

	return response.PSP_RespuestaStruct_NotifyFraudScreeningReview, nil
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

	client := &http.Client{Transport: tr,
		Timeout: time.Duration(time.Duration(Configuration.execution_Timeout) * time.Second)}

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
