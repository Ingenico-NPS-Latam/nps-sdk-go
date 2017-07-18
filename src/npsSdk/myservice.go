package npsSdk

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

func NewTaxesRequestStruct() *TaxesRequestStruct {
	p := new(TaxesRequestStruct)
	p.Xsi_Type = "tns.TaxesRequestStruct"

	return p
}

type TaxesRequestStruct struct {
	Xsi_Type string `xml:"xsi:type,attr"`

	TypeId string `xml:"TypeId,omitempty"`

	Amount string `xml:"Amount,omitempty"`

	Rate string `xml:"Rate,omitempty"`

	BaseAmount string `xml:"BaseAmount,omitempty"`
}

func NewArrayOf_TaxesRequestStruct() *ArrayOf_TaxesRequestStruct {
	p := new(ArrayOf_TaxesRequestStruct)
	p.Xsi_Type_SOAP = "SOAP-ENC:Array"
	p.Xsi_Type_SOAP_ARR = "tns:ArrayOf_TaxesRequestStruct"

	return p
}

type ArrayOf_TaxesRequestStruct struct {
	Xsi_Type_SOAP     string `xml:"xsi:type,attr"`
	Xsi_Type_SOAP_ARR string `xml:"SOAP-ENC:arrayType,attr"`

	Items []*TaxesRequestStruct `xml:"item,omitempty"`
}

func NewAmountAdditionalDetailsRequestStruct() *AmountAdditionalDetailsRequestStruct {
	p := new(AmountAdditionalDetailsRequestStruct)
	return p
}

type AmountAdditionalDetailsRequestStruct struct {
	XMLName xml.Name `xml:"psp_AmountAdditionalDetails"`
	Tip     string   `xml:"Tip,omitempty"`

	Discount string `xml:"Discount,omitempty"`

	Taxes *ArrayOf_TaxesRequestStruct
}

func NewAddressStruct() *AddressStruct {
	p := new(AddressStruct)
	return p
}

type AddressStruct struct {
	XMLName xml.Name //`xml:"AddressStruct"`

	Street string `xml:"Street,omitempty"`

	HouseNumber string `xml:"HouseNumber,omitempty"`

	AdditionalInfo string `xml:"AdditionalInfo,omitempty"`

	City string `xml:"City,omitempty"`

	StateProvince string `xml:"StateProvince,omitempty"`

	Country string `xml:"Country,omitempty"`

	ZipCode string `xml:"ZipCode,omitempty"`
}

func NewSellerDetailsStruct() *SellerDetailsStruct {
	p := new(SellerDetailsStruct)
	return p
}

type SellerDetailsStruct struct {
	XMLName xml.Name `xml:"SellerDetails"`

	IDNumber string `xml:"IDNumber,omitempty"`

	IDType string `xml:"IDType,omitempty"`

	Name string `xml:"Name,omitempty"`

	Invoice string `xml:"Invoice,omitempty"`

	PurchaseDescription string `xml:"PurchaseDescription,omitempty"`

	Address *AddressStruct

	MCC string `xml:"MCC,omitempty"`

	ChannelCode string `xml:"ChannelCode,omitempty"`

	GeoCode string `xml:"GeoCode,omitempty"`
}

func NewMerchantAdditionalDetailsStruct() *MerchantAdditionalDetailsStruct {
	p := new(MerchantAdditionalDetailsStruct)
	return p
}

type MerchantAdditionalDetailsStruct struct {
	XMLName xml.Name `xml:"psp_MerchantAdditionalDetails"`

	Type string `xml:"Type,omitempty"`

	SellerDetails *SellerDetailsStruct

	SdkInfo string `xml:"SdkInfo,omitempty"`

	ShoppingCartInfo string `xml:"ShoppingCartInfo,omitempty"`

	ShoppingCartPluginInfo string `xml:"ShoppingCartPluginInfo,omitempty"`
}

func NewCustomerAdditionalDetailsStruct() *CustomerAdditionalDetailsStruct {
	p := new(CustomerAdditionalDetailsStruct)
	return p
}

type CustomerAdditionalDetailsStruct struct {
	XMLName xml.Name `xml:"psp_CustomerAdditionalDetails"`

	EmailAddress string `xml:"EmailAddress,omitempty"`

	AlternativeEmailAddress string `xml:"AlternativeEmailAddress,omitempty"`

	IPAddress string `xml:"IPAddress,omitempty"`

	AccountID string `xml:"AccountID,omitempty"`

	AccountCreatedAt string `xml:"AccountCreatedAt,omitempty"`

	AccountPreviousActivity string `xml:"AccountPreviousActivity,omitempty"`

	AccountHasCredentials string `xml:"AccountHasCredentials,omitempty"`

	DeviceType string `xml:"DeviceType,omitempty"`

	DeviceFingerPrint string `xml:"DeviceFingerPrint,omitempty"`

	BrowserLanguage string `xml:"BrowserLanguage,omitempty"`

	HttpUserAgent string `xml:"HttpUserAgent,omitempty"`
}

func NewPersonStruct() *PersonStruct {
	p := new(PersonStruct)
	return p
}

type PersonStruct struct {
	XMLName xml.Name //`xml:"PersonStruct"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	PhoneNumber1 string `xml:"PhoneNumber1,omitempty"`

	PhoneNumber2 string `xml:"PhoneNumber2,omitempty"`

	Gender string `xml:"Gender,omitempty"`

	DateOfBirth string `xml:"DateOfBirth,omitempty"`

	Nationality string `xml:"Nationality,omitempty"`

	IDNumber string `xml:"IDNumber,omitempty"`

	IDType string `xml:"IDType,omitempty"`
}

func NewBillingDetailsStruct() *BillingDetailsStruct {
	p := new(BillingDetailsStruct)
	return p
}

type BillingDetailsStruct struct {
	XMLName xml.Name `xml:"psp_BillingDetails"`

	Invoice string `xml:"Invoice,omitempty"`

	InvoiceDate string `xml:"InvoiceDate,omitempty"`

	InvoiceAmount string `xml:"InvoiceAmount,omitempty"`

	InvoiceCurrency string `xml:"InvoiceCurrency,omitempty"`

	Person *PersonStruct

	Address *AddressStruct
}

func NewShippingDetailsStruct() *ShippingDetailsStruct {
	p := new(ShippingDetailsStruct)
	return p
}

type ShippingDetailsStruct struct {
	XMLName xml.Name `xml:"psp_ShippingDetails"`

	TrackingNumber string `xml:"TrackingNumber,omitempty"`

	Method string `xml:"Method,omitempty"`

	Carrier string `xml:"Carrier,omitempty"`

	DeliveryDate string `xml:"DeliveryDate,omitempty"`

	FreightAmount string `xml:"FreightAmount,omitempty"`

	GiftMessage string `xml:"GiftMessage,omitempty"`

	GiftWrapping string `xml:"GiftWrapping,omitempty"`

	PrimaryRecipient *PersonStruct

	SecondaryRecipient *PersonStruct

	Address *AddressStruct
}

func NewOrderItemStruct() *OrderItemStruct {
	p := new(OrderItemStruct)
	p.Xsi_Type = "tns.OrderItemStruct"
	return p
}

type OrderItemStruct struct {
	//XMLName xml.Name `xml:"OrderItemStruct"`
	Xsi_Type string `xml:"xsi:type,attr"`

	Quantity string `xml:"Quantity,omitempty"`

	UnitPrice string `xml:"UnitPrice,omitempty"`

	Description string `xml:"Description,omitempty"`

	Type string `xml:"Type,omitempty"`

	SkuCode string `xml:"SkuCode,omitempty"`

	ManufacturerPartNumber string `xml:"ManufacturerPartNumber,omitempty"`

	Risk string `xml:"Risk,omitempty"`
}

func NewArrayOf_OrderItemStruct() *ArrayOf_OrderItemStruct {
	p := new(ArrayOf_OrderItemStruct)
	p.Xsi_Type_SOAP = "SOAP-ENC:Array"
	p.Xsi_Type_SOAP_ARR = "tns:ArrayOf_OrderItemStruct"

	return p
}

type ArrayOf_OrderItemStruct struct {
	//	XMLName xml.Name `xml:"ArrayOf_OrderItemStruct"`
	Xsi_Type_SOAP     string `xml:"xsi:type,attr"`
	Xsi_Type_SOAP_ARR string `xml:"SOAP-ENC:arrayType,attr"`

	Items []*OrderItemStruct `xml:"item,omitempty"`
}

func NewOrderDetailsStruct() *OrderDetailsStruct {
	p := new(OrderDetailsStruct)
	return p
}

type OrderDetailsStruct struct {
	XMLName xml.Name `xml:"psp_OrderDetails"`

	OrderItems *ArrayOf_OrderItemStruct
}

func NewLegStruct() *LegStruct {
	p := new(LegStruct)
	return p
}

type LegStruct struct {
	XMLName xml.Name `xml:"LegStruct"`

	DepartureAirport string `xml:"DepartureAirport,omitempty"`

	DepartureDatetime string `xml:"DepartureDatetime,omitempty"`

	DepartureAirportTimezone string `xml:"DepartureAirportTimezone,omitempty"`

	ArrivalAirport string `xml:"ArrivalAirport,omitempty"`

	CarrierCode string `xml:"CarrierCode,omitempty"`

	FlightNumber string `xml:"FlightNumber,omitempty"`

	FareBasisCode string `xml:"FareBasisCode,omitempty"`

	FareClassCode string `xml:"FareClassCode,omitempty"`

	BaseFare string `xml:"BaseFare,omitempty"`

	BaseFareCurrency string `xml:"BaseFareCurrency,omitempty"`
}

func NewArrayOf_LegStruct() *ArrayOf_LegStruct {
	p := new(ArrayOf_LegStruct)
	p.Xsi_Type_SOAP = "SOAP-ENC:Array"
	p.Xsi_Type_SOAP_ARR = "tns:ArrayOf_LegStruct"

	return p
}

type ArrayOf_LegStruct struct {
	Xsi_Type_SOAP     string `xml:"xsi:type,attr"`
	Xsi_Type_SOAP_ARR string `xml:"SOAP-ENC:arrayType,attr"`

	Items []*LegStruct `xml:"item,omitempty"`
}

func NewPassengerStruct() *PassengerStruct {
	p := new(PassengerStruct)
	return p
}

type PassengerStruct struct {
	XMLName xml.Name `xml:"Passenger"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	Type string `xml:"Type,omitempty"`

	DateOfBirth string `xml:"DateOfBirth,omitempty"`

	Nationality string `xml:"Nationality,omitempty"`

	IDNumber string `xml:"IDNumber,omitempty"`

	IDType string `xml:"IDType,omitempty"`

	IDCountry string `xml:"IDCountry,omitempty"`

	LoyaltyNumber string `xml:"LoyaltyNumber,omitempty"`

	LoyaltyTier string `xml:"LoyaltyTier,omitempty"`
}

func NewArrayOf_PassengerStruct() *ArrayOf_PassengerStruct {
	p := new(ArrayOf_PassengerStruct)
	p.Xsi_Type_SOAP = "SOAP-ENC:Array"
	p.Xsi_Type_SOAP_ARR = "tns:ArrayOf_PassengerStruct"

	return p
}

type ArrayOf_PassengerStruct struct {
	Xsi_Type_SOAP     string `xml:"xsi:type,attr"`
	Xsi_Type_SOAP_ARR string `xml:"SOAP-ENC:arrayType,attr"`

	Items []*PassengerStruct `xml:"item,omitempty"`
}

func NewCc() *Cc {
	p := new(Cc)
	return p
}

type Cc struct {
	XMLName xml.Name `xml:"cc"`

	CarrierPrefixCode string `xml:"CarrierPrefixCode,omitempty"`

	TravelAgentCode string `xml:"TravelAgentCode,omitempty"`

	TravelAgentName string `xml:"TravelAgentName,omitempty"`

	Date string `xml:"Date,omitempty"`

	Country string `xml:"Country,omitempty"`

	City string `xml:"City,omitempty"`

	Address string `xml:"Address,omitempty"`
}

func NewAirlineTicketStruct() *AirlineTicketStruct {
	p := new(AirlineTicketStruct)
	return p
}

type AirlineTicketStruct struct {
	XMLName xml.Name `xml:"AirlineTicketStruct"`

	TicketNumber string `xml:"TicketNumber,omitempty"`

	Eticket string `xml:"Eticket,omitempty"`

	Restricted string `xml:"Restricted,omitempty"`

	Issue *Cc

	TotalFareAmount string `xml:"TotalFareAmount,omitempty"`

	TotalTaxAmount string `xml:"TotalTaxAmount,omitempty"`

	TotalFeeAmount string `xml:"TotalFeeAmount,omitempty"`
}

func NewAirlineDetailsStruct() *AirlineDetailsStruct {
	p := new(AirlineDetailsStruct)
	return p
}

type AirlineDetailsStruct struct {
	XMLName xml.Name `xml:"AirlineDetailsStruct"`

	PNR string `xml:"PNR,omitempty"`

	Legs *ArrayOf_LegStruct

	Passengers *ArrayOf_PassengerStruct

	Ticket *AirlineTicketStruct
}

func NewVaultReference2pStruct() *VaultReference2pStruct {
	p := new(VaultReference2pStruct)
	return p
}

type VaultReference2pStruct struct {
	XMLName xml.Name `xml:"psp_VaultReference"`

	PaymentMethodToken string `xml:"PaymentMethodToken,omitempty"`

	PaymentMethodId string `xml:"PaymentMethodId,omitempty"`

	CustomerId string `xml:"CustomerId,omitempty"`
}

type ServiceStruct_PayOnLine_2p struct {
	XMLName xml.Name `xml:"PayOnLine_2p"`

	PSP_RequerimientoStruct_PayOnLine_2p *RequerimientoStruct_PayOnLine_2p
}

func NewRequerimientoStruct_PayOnLine_2p() *RequerimientoStruct_PayOnLine_2p {
	p := new(RequerimientoStruct_PayOnLine_2p)
	return p
}

type RequerimientoStruct_PayOnLine_2p struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_PayOnLine_2p"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TxSource string `xml:"psp_TxSource,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_Recurrent string `xml:"psp_Recurrent,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CardNumber string `xml:"psp_CardNumber,omitempty"`

	Psp_CardExpDate string `xml:"psp_CardExpDate,omitempty"`

	Psp_CardSecurityCode string `xml:"psp_CardSecurityCode,omitempty"`

	Psp_CardHolderName string `xml:"psp_CardHolderName,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_PurchaseDescription string `xml:"psp_PurchaseDescription,omitempty"`

	Psp_SoftDescriptor string `xml:"psp_SoftDescriptor,omitempty"`

	Psp_ForceProcessingMethod string `xml:"psp_ForceProcessingMethod,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_FirstPaymentDeferral string `xml:"psp_FirstPaymentDeferral,omitempty"`

	Psp_3dSecure_XID string `xml:"psp_3dSecure_XID,omitempty"`

	Psp_3dSecure_CAVV string `xml:"psp_3dSecure_CAVV,omitempty"`

	Psp_3dSecure_ECI string `xml:"psp_3dSecure_ECI,omitempty"`

	Psp_3dSecure_Secured string `xml:"psp_3dSecure_Secured,omitempty"`

	Psp_VisaArg_DA_DocType string `xml:"psp_VisaArg_DA_DocType,omitempty"`

	Psp_VisaArg_DA_DocNum string `xml:"psp_VisaArg_DA_DocNum,omitempty"`

	Psp_VisaArg_DA_DoorNum string `xml:"psp_VisaArg_DA_DoorNum,omitempty"`

	Psp_VisaArg_DA_BirthDay string `xml:"psp_VisaArg_DA_BirthDay,omitempty"`

	Psp_VisaArg_DA_Name string `xml:"psp_VisaArg_DA_Name,omitempty"`

	Psp_AmexArg_AVS_Address string `xml:"psp_AmexArg_AVS_Address,omitempty"`

	Psp_AmexArg_AVS_PostCode string `xml:"psp_AmexArg_AVS_PostCode,omitempty"`

	Psp_MasterArg_AVS_BirthDay string `xml:"psp_MasterArg_AVS_BirthDay,omitempty"`

	Psp_MasterArg_AVS_AdditionalsQtty string `xml:"psp_MasterArg_AVS_AdditionalsQtty,omitempty"`

	Psp_MasterArg_AVS_PostalCode string `xml:"psp_MasterArg_AVS_PostalCode,omitempty"`

	Psp_MasterArg_AVS_BranchEntity string `xml:"psp_MasterArg_AVS_BranchEntity,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsRequestStruct

	Psp_MerchantAdditionalDetails *MerchantAdditionalDetailsStruct

	Psp_CustomerAdditionalDetails *CustomerAdditionalDetailsStruct

	Psp_BillingDetails *BillingDetailsStruct

	Psp_ShippingDetails *ShippingDetailsStruct

	Psp_OrderDetails *OrderDetailsStruct

	Psp_AirlineDetails *AirlineDetailsStruct

	Psp_VaultReference *VaultReference2pStruct
}

type TaxesResponseStruct struct {
	XMLName xml.Name `xml:"TaxesResponseStruct"`

	TypeId string `xml:"TypeId,omitempty"`

	TypeDescription string `xml:"TypeDescription,omitempty"`

	Amount string `xml:"Amount,omitempty"`

	Rate string `xml:"Rate,omitempty"`

	BaseAmount string `xml:"BaseAmount,omitempty"`

	AppliedAmount string `xml:"AppliedAmount,omitempty"`

	Remarks string `xml:"Remarks,omitempty"`
}

type ArrayOf_TaxesResponseStruct struct {
	XMLName xml.Name `xml:"ArrayOf_TaxesResponseStruct"`

	Items []*TaxesResponseStruct `xml:"item,omitempty"`
}

type AmountAdditionalDetailsResponseStruct struct {
	XMLName xml.Name `xml:"AmountAdditionalDetailsResponseStruct"`

	Tip string `xml:"Tip,omitempty"`

	Discount string `xml:"Discount,omitempty"`

	Taxes *ArrayOf_TaxesResponseStruct
}

type FraudScreeningResultStruct struct {
	XMLName xml.Name `xml:"FraudScreeningResultStruct"`

	ResultCode string `xml:"ResultCode,omitempty"`

	ResultDescription string `xml:"ResultDescription,omitempty"`

	AdditionalInfo string `xml:"AdditionalInfo,omitempty"`
}

type VerificationServicesResultStruct struct {
	XMLName xml.Name `xml:"VerificationServicesResultStruct"`

	ResultCodeCardSecurityCode string `xml:"ResultCodeCardSecurityCode,omitempty"`

	ResultCodeBillingAddress string `xml:"ResultCodeBillingAddress,omitempty"`

	ResultCodeBillingAddressZipCode string `xml:"ResultCodeBillingAddressZipCode,omitempty"`

	ResultCodeBillingPersonIDType string `xml:"ResultCodeBillingPersonIDType,omitempty"`

	ResultCodeBillingPersonIDNumber string `xml:"ResultCodeBillingPersonIDNumber,omitempty"`

	ResultCodeBillingPersonDateOfBirth string `xml:"ResultCodeBillingPersonDateOfBirth,omitempty"`

	ResultCodeBillingPersonName string `xml:"ResultCodeBillingPersonName,omitempty"`

	ResultCodeBillingPersonPhoneNumber1 string `xml:"ResultCodeBillingPersonPhoneNumber1,omitempty"`

	ResultCodeCustomerEmailAddress string `xml:"ResultCodeCustomerEmailAddress,omitempty"`
}

type ServiceRespuestaStruct_PayOnLine_2p struct {
	XMLName xml.Name `xml:"PayOnLine_2pResponse"`

	PSP_RespuestaStruct_PayOnLine_2p *RespuestaStruct_PayOnLine_2p
}

type RespuestaStruct_PayOnLine_2p struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_PaymentAmount string `xml:"psp_PaymentAmount,omitempty"`

	Psp_Recurrent string `xml:"psp_Recurrent,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CardNumber string `xml:"psp_CardNumber,omitempty"`

	Psp_CardExpDate string `xml:"psp_CardExpDate,omitempty"`

	Psp_AuthorizationCode string `xml:"psp_AuthorizationCode,omitempty"`

	Psp_BatchNro string `xml:"psp_BatchNro,omitempty"`

	Psp_TicketNumber string `xml:"psp_TicketNumber,omitempty"`

	Psp_SequenceNumber string `xml:"psp_SequenceNumber,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_ClTrnId string `xml:"psp_ClTrnId,omitempty"`

	Psp_ClExternalMerchant string `xml:"psp_ClExternalMerchant,omitempty"`

	Psp_ClExternalTerminal string `xml:"psp_ClExternalTerminal,omitempty"`

	Psp_ClResponseCod string `xml:"psp_ClResponseCod,omitempty"`

	Psp_ClResponseMsg string `xml:"psp_ClResponseMsg,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_FirstPaymentDeferral string `xml:"psp_FirstPaymentDeferral,omitempty"`

	Psp_VisaArg_DA_Result string `xml:"psp_VisaArg_DA_Result,omitempty"`

	Psp_AmexArg_AVS_Result string `xml:"psp_AmexArg_AVS_Result,omitempty"`

	Psp_MasterArg_AVS_Result string `xml:"psp_MasterArg_AVS_Result,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsResponseStruct

	Psp_FraudScreeningResult *FraudScreeningResultStruct

	Psp_VerificationServicesResult *VerificationServicesResultStruct
}

type ServiceStruct_Authorize_2p struct {
	XMLName xml.Name `xml:"Authorize_2p"`

	PSP_RequerimientoStruct_Authorize_2p *RequerimientoStruct_Authorize_2p
}

func NewRequerimientoStruct_Authorize_2p() *RequerimientoStruct_Authorize_2p {
	p := new(RequerimientoStruct_Authorize_2p)
	return p
}

type RequerimientoStruct_Authorize_2p struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_Authorize_2p"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TxSource string `xml:"psp_TxSource,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_Recurrent string `xml:"psp_Recurrent,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CardNumber string `xml:"psp_CardNumber,omitempty"`

	Psp_CardExpDate string `xml:"psp_CardExpDate,omitempty"`

	Psp_CardSecurityCode string `xml:"psp_CardSecurityCode,omitempty"`

	Psp_CardHolderName string `xml:"psp_CardHolderName,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_PurchaseDescription string `xml:"psp_PurchaseDescription,omitempty"`

	Psp_SoftDescriptor string `xml:"psp_SoftDescriptor,omitempty"`

	Psp_ForceProcessingMethod string `xml:"psp_ForceProcessingMethod,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_3dSecure_XID string `xml:"psp_3dSecure_XID,omitempty"`

	Psp_3dSecure_CAVV string `xml:"psp_3dSecure_CAVV,omitempty"`

	Psp_3dSecure_ECI string `xml:"psp_3dSecure_ECI,omitempty"`

	Psp_3dSecure_Secured string `xml:"psp_3dSecure_Secured,omitempty"`

	Psp_VisaArg_DA_DocType string `xml:"psp_VisaArg_DA_DocType,omitempty"`

	Psp_VisaArg_DA_DocNum string `xml:"psp_VisaArg_DA_DocNum,omitempty"`

	Psp_VisaArg_DA_DoorNum string `xml:"psp_VisaArg_DA_DoorNum,omitempty"`

	Psp_VisaArg_DA_BirthDay string `xml:"psp_VisaArg_DA_BirthDay,omitempty"`

	Psp_VisaArg_DA_Name string `xml:"psp_VisaArg_DA_Name,omitempty"`

	Psp_AmexArg_AVS_Address string `xml:"psp_AmexArg_AVS_Address,omitempty"`

	Psp_AmexArg_AVS_PostCode string `xml:"psp_AmexArg_AVS_PostCode,omitempty"`

	Psp_MasterArg_AVS_BirthDay string `xml:"psp_MasterArg_AVS_BirthDay,omitempty"`

	Psp_MasterArg_AVS_AdditionalsQtty string `xml:"psp_MasterArg_AVS_AdditionalsQtty,omitempty"`

	Psp_MasterArg_AVS_PostalCode string `xml:"psp_MasterArg_AVS_PostalCode,omitempty"`

	Psp_MasterArg_AVS_BranchEntity string `xml:"psp_MasterArg_AVS_BranchEntity,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsRequestStruct

	Psp_MerchantAdditionalDetails *MerchantAdditionalDetailsStruct

	Psp_CustomerAdditionalDetails *CustomerAdditionalDetailsStruct

	Psp_BillingDetails *BillingDetailsStruct

	Psp_ShippingDetails *ShippingDetailsStruct

	Psp_OrderDetails *OrderDetailsStruct

	Psp_AirlineDetails *AirlineDetailsStruct

	Psp_VaultReference *VaultReference2pStruct
}

type ServiceRespuestaStruct_Authorize_2p struct {
	XMLName xml.Name `xml:"Authorize_2pResponse"`

	PSP_RespuestaStruct_Authorize_2p *RespuestaStruct_Authorize_2p
}

type RespuestaStruct_Authorize_2p struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_PaymentAmount string `xml:"psp_PaymentAmount,omitempty"`

	Psp_Recurrent string `xml:"psp_Recurrent,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CardNumber string `xml:"psp_CardNumber,omitempty"`

	Psp_CardExpDate string `xml:"psp_CardExpDate,omitempty"`

	Psp_AuthorizationCode string `xml:"psp_AuthorizationCode,omitempty"`

	Psp_BatchNro string `xml:"psp_BatchNro,omitempty"`

	Psp_TicketNumber string `xml:"psp_TicketNumber,omitempty"`

	Psp_SequenceNumber string `xml:"psp_SequenceNumber,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_ClTrnId string `xml:"psp_ClTrnId,omitempty"`

	Psp_ClExternalMerchant string `xml:"psp_ClExternalMerchant,omitempty"`

	Psp_ClExternalTerminal string `xml:"psp_ClExternalTerminal,omitempty"`

	Psp_ClResponseCod string `xml:"psp_ClResponseCod,omitempty"`

	Psp_ClResponseMsg string `xml:"psp_ClResponseMsg,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_VisaArg_DA_Result string `xml:"psp_VisaArg_DA_Result,omitempty"`

	Psp_AmexArg_AVS_Result string `xml:"psp_AmexArg_AVS_Result,omitempty"`

	Psp_MasterArg_AVS_Result string `xml:"psp_MasterArg_AVS_Result,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsResponseStruct

	Psp_FraudScreeningResult *FraudScreeningResultStruct

	Psp_VerificationServicesResult *VerificationServicesResultStruct
}

type ServiceStruct_BankPayment_2p struct {
	XMLName xml.Name `xml:"BankPayment_2p"`

	PSP_RequerimientoStruct_BankPayment_2p *RequerimientoStruct_BankPayment_2p
}

func NewRequerimientoStruct_BankPayment_2p() *RequerimientoStruct_BankPayment_2p {
	p := new(RequerimientoStruct_BankPayment_2p)
	return p
}

type RequerimientoStruct_BankPayment_2p struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_BankPayment_2p"`
	XMLName xml.Name `xml:"Requerimient"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TxSource string `xml:"psp_TxSource,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_ScreenDescription string `xml:"psp_ScreenDescription,omitempty"`

	Psp_TicketDescription string `xml:"psp_TicketDescription,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_Amount1 string `xml:"psp_Amount1,omitempty"`

	Psp_ExpDate1 string `xml:"psp_ExpDate1,omitempty"`

	Psp_Amount2 string `xml:"psp_Amount2,omitempty"`

	Psp_ExpDate2 string `xml:"psp_ExpDate2,omitempty"`

	Psp_Amount3 string `xml:"psp_Amount3,omitempty"`

	Psp_ExpDate3 string `xml:"psp_ExpDate3,omitempty"`

	Psp_MinAmount string `xml:"psp_MinAmount,omitempty"`

	Psp_ExpTime string `xml:"psp_ExpTime,omitempty"`

	Psp_ExpMark string `xml:"psp_ExpMark,omitempty"`

	Psp_CustomerBankId string `xml:"psp_CustomerBankId,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_ForceProcessingMethod string `xml:"psp_ForceProcessingMethod,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsRequestStruct

	Psp_MerchantAdditionalDetails *MerchantAdditionalDetailsStruct

	Psp_CustomerAdditionalDetails *CustomerAdditionalDetailsStruct

	Psp_BillingDetails *BillingDetailsStruct

	Psp_ShippingDetails *ShippingDetailsStruct

	Psp_OrderDetails *OrderDetailsStruct

	Psp_AirlineDetails *AirlineDetailsStruct
}

type ServiceRespuestaStruct_BankPayment_2p_BillingDetails struct {
	XMLName xml.Name `xml:"BankPayment_2p_BillingDetailsResponse"`

	PSP_RespuestaStruct_BankPayment_2p_BillingDetails *RespuestaStruct_BankPayment_2p_BillingDetails
}

type RespuestaStruct_BankPayment_2p_BillingDetails struct {
	XMLName xml.Name `xml:"Respuesta"`

	Invoice string `xml:"Invoice,omitempty"`

	InvoiceDate string `xml:"InvoiceDate,omitempty"`

	InvoiceAmount string `xml:"InvoiceAmount,omitempty"`

	InvoiceCurrency string `xml:"InvoiceCurrency,omitempty"`

	Person *PersonStruct

	Address *AddressStruct
}

type ServiceRespuestaStruct_BankPayment_2p struct {
	XMLName xml.Name `xml:"BankPayment_2pResponse"`

	PSP_RespuestaStruct_BankPayment_2p *RespuestaStruct_BankPayment_2p
}

type RespuestaStruct_BankPayment_2p struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_ScreenDescription string `xml:"psp_ScreenDescription,omitempty"`

	Psp_TicketDescription string `xml:"psp_TicketDescription,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_Amount1 string `xml:"psp_Amount1,omitempty"`

	Psp_ExpDate1 string `xml:"psp_ExpDate1,omitempty"`

	Psp_Amount2 string `xml:"psp_Amount2,omitempty"`

	Psp_ExpDate2 string `xml:"psp_ExpDate2,omitempty"`

	Psp_Amount3 string `xml:"psp_Amount3,omitempty"`

	Psp_ExpDate3 string `xml:"psp_ExpDate3,omitempty"`

	Psp_MinAmount string `xml:"psp_MinAmount,omitempty"`

	Psp_ExpTime string `xml:"psp_ExpTime,omitempty"`

	Psp_ExpMark string `xml:"psp_ExpMark,omitempty"`

	Psp_CustomerBankId string `xml:"psp_CustomerBankId,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsResponseStruct

	Psp_BillingDetails *RespuestaStruct_BankPayment_2p_BillingDetails
}

func NewVaultReference3pStruct() *VaultReference3pStruct {
	p := new(VaultReference3pStruct)
	return p
}

type VaultReference3pStruct struct {
	XMLName xml.Name `xml:"psp_VaultReference"`

	CustomerId string `xml:"CustomerId,omitempty"`
}

func NewRequerimientoStruct_SplitPayOnLine_2p_Transactions() *RequerimientoStruct_SplitPayOnLine_2p_Transactions {
	p := new(RequerimientoStruct_SplitPayOnLine_2p_Transactions)
	return p
}

type RequerimientoStruct_SplitPayOnLine_2p_Transactions struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_SplitPayOnLine_2p_Transactions"`
	//XMLName xml.Name `xml:"Requerimiento"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CardNumber string `xml:"psp_CardNumber,omitempty"`

	Psp_CardExpDate string `xml:"psp_CardExpDate,omitempty"`

	Psp_CardSecurityCode string `xml:"psp_CardSecurityCode,omitempty"`

	Psp_CardHolderName string `xml:"psp_CardHolderName,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_FirstPaymentDeferral string `xml:"psp_FirstPaymentDeferral,omitempty"`

	Psp_SoftDescriptor string `xml:"psp_SoftDescriptor,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsRequestStruct

	Psp_MerchantAdditionalDetails *MerchantAdditionalDetailsStruct

	Psp_BillingDetails *BillingDetailsStruct

	Psp_VaultReference *VaultReference2pStruct
}

func NewArrayOf_RequerimientoStruct_SplitPayOnLine_2p_Transactions() *ArrayOf_RequerimientoStruct_SplitPayOnLine_2p_Transactions {
	p := new(ArrayOf_RequerimientoStruct_SplitPayOnLine_2p_Transactions)
	p.Xsi_Type_SOAP = "SOAP-ENC:Array"
	p.Xsi_Type_SOAP_ARR = "tns:ArrayOf_RequerimientoStruct_SplitPayOnLine_2p_Transactions"

	return p
}

type ArrayOf_RequerimientoStruct_SplitPayOnLine_2p_Transactions struct {
	Xsi_Type_SOAP     string `xml:"xsi:type,attr"`
	Xsi_Type_SOAP_ARR string `xml:"SOAP-ENC:arrayType,attr"`

	Items []*RequerimientoStruct_SplitPayOnLine_2p_Transactions `xml:"item,omitempty"`
}

type ServiceStruct_SplitPayOnLine_2p struct {
	XMLName xml.Name `xml:"SplitPayOnLine_2p"`

	PSP_RequerimientoStruct_SplitPayOnLine_2p *RequerimientoStruct_SplitPayOnLine_2p
}

func NewRequerimientoStruct_SplitPayOnLine_2p() *RequerimientoStruct_SplitPayOnLine_2p {
	p := new(RequerimientoStruct_SplitPayOnLine_2p)
	return p
}

type RequerimientoStruct_SplitPayOnLine_2p struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_SplitPayOnLine_2p"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TxSource string `xml:"psp_TxSource,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Recurrent string `xml:"psp_Recurrent,omitempty"`

	Psp_UseMultipleProducts string `xml:"psp_UseMultipleProducts,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CardNumber string `xml:"psp_CardNumber,omitempty"`

	Psp_CardExpDate string `xml:"psp_CardExpDate,omitempty"`

	Psp_CardSecurityCode string `xml:"psp_CardSecurityCode,omitempty"`

	Psp_CardHolderName string `xml:"psp_CardHolderName,omitempty"`

	Psp_PurchaseDescription string `xml:"psp_PurchaseDescription,omitempty"`

	Psp_SoftDescriptor string `xml:"psp_SoftDescriptor,omitempty"`

	Psp_ForceProcessingMethod string `xml:"psp_ForceProcessingMethod,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_3dSecure_XID string `xml:"psp_3dSecure_XID,omitempty"`

	Psp_3dSecure_CAVV string `xml:"psp_3dSecure_CAVV,omitempty"`

	Psp_3dSecure_ECI string `xml:"psp_3dSecure_ECI,omitempty"`

	Psp_3dSecure_Secured string `xml:"psp_3dSecure_Secured,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`

	Psp_VaultReference *VaultReference2pStruct

	Psp_MerchantAdditionalDetails *MerchantAdditionalDetailsStruct

	Psp_CustomerAdditionalDetails *CustomerAdditionalDetailsStruct

	Psp_BillingDetails *BillingDetailsStruct

	Psp_ShippingDetails *ShippingDetailsStruct

	Psp_OrderDetails *OrderDetailsStruct

	Psp_AirlineDetails *AirlineDetailsStruct

	Psp_Transactions *ArrayOf_RequerimientoStruct_SplitPayOnLine_2p_Transactions `xml:"psp_Transactions,omitempty"`
}

type RespuestaStruct_SplitPayOnLine_2p_Transactions struct {
	//XMLName xml.Name `xml:"Respuesta"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_PaymentAmount string `xml:"psp_PaymentAmount,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CardNumber string `xml:"psp_CardNumber,omitempty"`

	Psp_CardExpDate string `xml:"psp_CardExpDate,omitempty"`

	Psp_AuthorizationCode string `xml:"psp_AuthorizationCode,omitempty"`

	Psp_BatchNro string `xml:"psp_BatchNro,omitempty"`

	Psp_TicketNumber string `xml:"psp_TicketNumber,omitempty"`

	Psp_SequenceNumber string `xml:"psp_SequenceNumber,omitempty"`

	Psp_ClTrnId string `xml:"psp_ClTrnId,omitempty"`

	Psp_ClExternalMerchant string `xml:"psp_ClExternalMerchant,omitempty"`

	Psp_ClExternalTerminal string `xml:"psp_ClExternalTerminal,omitempty"`

	Psp_ClResponseCod string `xml:"psp_ClResponseCod,omitempty"`

	Psp_ClResponseMsg string `xml:"psp_ClResponseMsg,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_FirstPaymentDeferral string `xml:"psp_FirstPaymentDeferral,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsResponseStruct

	Psp_FraudScreeningResult *FraudScreeningResultStruct

	Psp_VerificationServicesResult *VerificationServicesResultStruct
}

type ArrayOf_RespuestaStruct_SplitPayOnLine_2p_Transactions struct {
	XMLName xml.Name `xml:"ArrayOf_RespuestaStruct_SplitPayOnLine_2p_Transactions"`

	Items []*RespuestaStruct_SplitPayOnLine_2p_Transactions `xml:"item,omitempty"`
}

type ServiceRespuestaStruct_SplitPayOnLine_2p struct {
	XMLName xml.Name `xml:"SplitPayOnLine_2pResponse"`

	PSP_RespuestaStruct_SplitPayOnLine_2p *RespuestaStruct_SplitPayOnLine_2p
}

type RespuestaStruct_SplitPayOnLine_2p struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_UseMultipleProducts string `xml:"psp_UseMultipleProducts,omitempty"`

	Psp_Transactions *ArrayOf_RespuestaStruct_SplitPayOnLine_2p_Transactions //`xml:"psp_Transactions,omitempty"`
}

/*type ServiceStruct_SplitAuthorize_2p_Transactions struct {
	XMLName xml.Name `xml:"SplitAuthorize_2p_Transactions"`

	PSP_RequerimientoStruct_SplitAuthorize_2p_Transactions *RequerimientoStruct_SplitAuthorize_2p_Transactions
}*/

func NewRequerimientoStruct_SplitAuthorize_2p_Transactions() *RequerimientoStruct_SplitAuthorize_2p_Transactions {
	p := new(RequerimientoStruct_SplitAuthorize_2p_Transactions)
	return p
}

type RequerimientoStruct_SplitAuthorize_2p_Transactions struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_SplitAuthorize_2p_Transactions"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CardNumber string `xml:"psp_CardNumber,omitempty"`

	Psp_CardExpDate string `xml:"psp_CardExpDate,omitempty"`

	Psp_CardSecurityCode string `xml:"psp_CardSecurityCode,omitempty"`

	Psp_CardHolderName string `xml:"psp_CardHolderName,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_FirstPaymentDeferral string `xml:"psp_FirstPaymentDeferral,omitempty"`

	Psp_SoftDescriptor string `xml:"psp_SoftDescriptor,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsRequestStruct

	Psp_MerchantAdditionalDetails *MerchantAdditionalDetailsStruct

	Psp_BillingDetails *BillingDetailsStruct

	Psp_VaultReference *VaultReference2pStruct
}

func NewArrayOf_RequerimientoStruct_SplitAuthorize_2p_Transactions() *ArrayOf_RequerimientoStruct_SplitAuthorize_2p_Transactions {
	p := new(ArrayOf_RequerimientoStruct_SplitAuthorize_2p_Transactions)
	p.Xsi_Type_SOAP = "SOAP-ENC:Array"
	p.Xsi_Type_SOAP_ARR = "tns:ArrayOf_RequerimientoStruct_SplitAuthorize_2p_Transactions"

	return p
}

type ArrayOf_RequerimientoStruct_SplitAuthorize_2p_Transactions struct {
	Xsi_Type_SOAP     string `xml:"xsi:type,attr"`
	Xsi_Type_SOAP_ARR string `xml:"SOAP-ENC:arrayType,attr"`

	Items []*RequerimientoStruct_SplitAuthorize_2p_Transactions `xml:"item,omitempty"`
}

type ServiceStruct_SplitAuthorize_2p struct {
	XMLName xml.Name `xml:"SplitAuthorize_2p"`

	PSP_RequerimientoStruct_SplitAuthorize_2p *RequerimientoStruct_SplitAuthorize_2p
}

func NewRequerimientoStruct_SplitAuthorize_2p() *RequerimientoStruct_SplitAuthorize_2p {
	p := new(RequerimientoStruct_SplitAuthorize_2p)
	return p
}

type RequerimientoStruct_SplitAuthorize_2p struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_SplitAuthorize_2p"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TxSource string `xml:"psp_TxSource,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Recurrent string `xml:"psp_Recurrent,omitempty"`

	Psp_UseMultipleProducts string `xml:"psp_UseMultipleProducts,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CardNumber string `xml:"psp_CardNumber,omitempty"`

	Psp_CardExpDate string `xml:"psp_CardExpDate,omitempty"`

	Psp_CardSecurityCode string `xml:"psp_CardSecurityCode,omitempty"`

	Psp_CardHolderName string `xml:"psp_CardHolderName,omitempty"`

	Psp_PurchaseDescription string `xml:"psp_PurchaseDescription,omitempty"`

	Psp_SoftDescriptor string `xml:"psp_SoftDescriptor,omitempty"`

	Psp_ForceProcessingMethod string `xml:"psp_ForceProcessingMethod,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_3dSecure_XID string `xml:"psp_3dSecure_XID,omitempty"`

	Psp_3dSecure_CAVV string `xml:"psp_3dSecure_CAVV,omitempty"`

	Psp_3dSecure_ECI string `xml:"psp_3dSecure_ECI,omitempty"`

	Psp_3dSecure_Secured string `xml:"psp_3dSecure_Secured,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`

	Psp_MerchantAdditionalDetails *MerchantAdditionalDetailsStruct

	Psp_CustomerAdditionalDetails *CustomerAdditionalDetailsStruct

	Psp_BillingDetails *BillingDetailsStruct

	Psp_ShippingDetails *ShippingDetailsStruct

	Psp_OrderDetails *OrderDetailsStruct

	Psp_AirlineDetails *AirlineDetailsStruct

	Psp_Transactions *ArrayOf_RequerimientoStruct_SplitAuthorize_2p_Transactions `xml:"psp_Transactions,omitempty"`
}

type ServiceRespuestaStruct_SplitAuthorize_2p struct {
	XMLName xml.Name `xml:"SplitAuthorize_2pResponse"`

	PSP_RespuestaStruct_SplitAuthorize_2p *RespuestaStruct_SplitAuthorize_2p
}

type RespuestaStruct_SplitAuthorize_2p struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_UseMultipleProducts string `xml:"psp_UseMultipleProducts,omitempty"`

	Psp_Transactions *ArrayOf_RespuestaStruct_SplitPayOnLine_2p_Transactions //`xml:"psp_Transactions,omitempty"`
}

type ServiceStruct_PayOnLine_3p struct {
	XMLName xml.Name `xml:"PayOnLine_3p"`

	PSP_RequerimientoStruct_PayOnLine_3p *RequerimientoStruct_PayOnLine_3p
}

func NewRequerimientoStruct_PayOnLine_3p() *RequerimientoStruct_PayOnLine_3p {
	p := new(RequerimientoStruct_PayOnLine_3p)
	return p
}

type RequerimientoStruct_PayOnLine_3p struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_PayOnLine_3p"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TxSource string `xml:"psp_TxSource,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_ReturnURL string `xml:"psp_ReturnURL,omitempty"`

	Psp_FrmLanguage string `xml:"psp_FrmLanguage,omitempty"`

	Psp_FrmBackButtonURL string `xml:"psp_FrmBackButtonURL,omitempty"`

	Psp_FrmTimeout string `xml:"psp_FrmTimeout,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_PurchaseDescription string `xml:"psp_PurchaseDescription,omitempty"`

	Psp_SoftDescriptor string `xml:"psp_SoftDescriptor,omitempty"`

	Psp_ForceProcessingMethod string `xml:"psp_ForceProcessingMethod,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_PromotionCode string `xml:"psp_PromotionCode,omitempty"`

	Psp_FirstPaymentDeferral string `xml:"psp_FirstPaymentDeferral,omitempty"`

	Psp_3dSecureAction string `xml:"psp_3dSecureAction,omitempty"`

	Psp_VisaArg_DA_DocType string `xml:"psp_VisaArg_DA_DocType,omitempty"`

	Psp_VisaArg_DA_DocNum string `xml:"psp_VisaArg_DA_DocNum,omitempty"`

	Psp_VisaArg_DA_DoorNum string `xml:"psp_VisaArg_DA_DoorNum,omitempty"`

	Psp_VisaArg_DA_BirthDay string `xml:"psp_VisaArg_DA_BirthDay,omitempty"`

	Psp_VisaArg_DA_Name string `xml:"psp_VisaArg_DA_Name,omitempty"`

	Psp_AmexArg_AVS_Address string `xml:"psp_AmexArg_AVS_Address,omitempty"`

	Psp_AmexArg_AVS_PostCode string `xml:"psp_AmexArg_AVS_PostCode,omitempty"`

	Psp_MasterArg_AVS_BirthDay string `xml:"psp_MasterArg_AVS_BirthDay,omitempty"`

	Psp_MasterArg_AVS_AdditionalsQtty string `xml:"psp_MasterArg_AVS_AdditionalsQtty,omitempty"`

	Psp_MasterArg_AVS_PostalCode string `xml:"psp_MasterArg_AVS_PostalCode,omitempty"`

	Psp_MasterArg_AVS_BranchEntity string `xml:"psp_MasterArg_AVS_BranchEntity,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`

	Psp_VaultReference *VaultReference3pStruct

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsRequestStruct

	Psp_MerchantAdditionalDetails *MerchantAdditionalDetailsStruct

	Psp_CustomerAdditionalDetails *CustomerAdditionalDetailsStruct

	Psp_BillingDetails *BillingDetailsStruct

	Psp_ShippingDetails *ShippingDetailsStruct

	Psp_OrderDetails *OrderDetailsStruct

	Psp_AirlineDetails *AirlineDetailsStruct
}

type ServiceRespuestaStruct_PayOnLine_3p struct {
	XMLName xml.Name `xml:"PayOnLine_3pResponse"`

	PSP_RespuestaStruct_PayOnLine_3p *RespuestaStruct_PayOnLine_3p
}

type RespuestaStruct_PayOnLine_3p struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_Session3p string `xml:"psp_Session3p,omitempty"`

	Psp_FrontPSP_URL string `xml:"psp_FrontPSP_URL,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_PaymentAmount string `xml:"psp_PaymentAmount,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_FirstPaymentDeferral string `xml:"psp_FirstPaymentDeferral,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`
}

func NewRequerimientoStruct_SplitPayOnLine_3p_Transactions() *RequerimientoStruct_SplitPayOnLine_3p_Transactions {
	p := new(RequerimientoStruct_SplitPayOnLine_3p_Transactions)
	p.Xsi_Type = "tns.RequerimientoStruct_SplitPayOnLine_3p_Transactions"

	return p
}

type RequerimientoStruct_SplitPayOnLine_3p_Transactions struct {
	Xsi_Type string `xml:"xsi:type,attr"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_PromotionCode string `xml:"psp_PromotionCode,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_FirstPaymentDeferral string `xml:"psp_FirstPaymentDeferral,omitempty"`

	Psp_SoftDescriptor string `xml:"psp_SoftDescriptor,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsRequestStruct
}

func NewArrayOf_RequerimientoStruct_SplitPayOnLine_3p_Transactions() *ArrayOf_RequerimientoStruct_SplitPayOnLine_3p_Transactions {
	p := new(ArrayOf_RequerimientoStruct_SplitPayOnLine_3p_Transactions)
	p.Xsi_Type_SOAP = "SOAP-ENC:Array"
	p.Xsi_Type_SOAP_ARR = "tns:ArrayOf_RequerimientoStruct_SplitPayOnLine_3p_Transactions"

	return p
}

type ArrayOf_RequerimientoStruct_SplitPayOnLine_3p_Transactions struct {
	Xsi_Type_SOAP     string `xml:"xsi:type,attr"`
	Xsi_Type_SOAP_ARR string `xml:"SOAP-ENC:arrayType,attr"`

	Items []*RequerimientoStruct_SplitPayOnLine_3p_Transactions `xml:"item,omitempty"`
}

type ServiceStruct_SplitPayOnLine_3p struct {
	XMLName xml.Name `xml:"SplitPayOnLine_3p"`

	PSP_RequerimientoStruct_SplitPayOnLine_3p *RequerimientoStruct_SplitPayOnLine_3p
}

func NewRequerimientoStruct_SplitPayOnLine_3p() *RequerimientoStruct_SplitPayOnLine_3p {
	p := new(RequerimientoStruct_SplitPayOnLine_3p)
	return p
}

type RequerimientoStruct_SplitPayOnLine_3p struct {
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TxSource string `xml:"psp_TxSource,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_ReturnURL string `xml:"psp_ReturnURL,omitempty"`

	Psp_FrmLanguage string `xml:"psp_FrmLanguage,omitempty"`

	Psp_FrmBackButtonURL string `xml:"psp_FrmBackButtonURL,omitempty"`

	Psp_FrmTimeout string `xml:"psp_FrmTimeout,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_UseMultipleProducts string `xml:"psp_UseMultipleProducts,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_PurchaseDescription string `xml:"psp_PurchaseDescription,omitempty"`

	Psp_SoftDescriptor string `xml:"psp_SoftDescriptor,omitempty"`

	Psp_ForceProcessingMethod string `xml:"psp_ForceProcessingMethod,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_PromotionCode string `xml:"psp_PromotionCode,omitempty"`

	Psp_3dSecureAction string `xml:"psp_3dSecureAction,omitempty"`

	Psp_VisaArg_DA_DocType string `xml:"psp_VisaArg_DA_DocType,omitempty"`

	Psp_VisaArg_DA_DocNum string `xml:"psp_VisaArg_DA_DocNum,omitempty"`

	Psp_VisaArg_DA_DoorNum string `xml:"psp_VisaArg_DA_DoorNum,omitempty"`

	Psp_VisaArg_DA_BirthDay string `xml:"psp_VisaArg_DA_BirthDay,omitempty"`

	Psp_VisaArg_DA_Name string `xml:"psp_VisaArg_DA_Name,omitempty"`

	Psp_AmexArg_AVS_Address string `xml:"psp_AmexArg_AVS_Address,omitempty"`

	Psp_AmexArg_AVS_PostCode string `xml:"psp_AmexArg_AVS_PostCode,omitempty"`

	Psp_MasterArg_AVS_BirthDay string `xml:"psp_MasterArg_AVS_BirthDay,omitempty"`

	Psp_MasterArg_AVS_AdditionalsQtty string `xml:"psp_MasterArg_AVS_AdditionalsQtty,omitempty"`

	Psp_MasterArg_AVS_PostalCode string `xml:"psp_MasterArg_AVS_PostalCode,omitempty"`

	Psp_MasterArg_AVS_BranchEntity string `xml:"psp_MasterArg_AVS_BranchEntity,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`

	Psp_VaultReference *VaultReference3pStruct

	Psp_MerchantAdditionalDetails *MerchantAdditionalDetailsStruct

	Psp_CustomerAdditionalDetails *CustomerAdditionalDetailsStruct

	Psp_BillingDetails *BillingDetailsStruct

	Psp_ShippingDetails *ShippingDetailsStruct

	Psp_OrderDetails *OrderDetailsStruct

	Psp_AirlineDetails *AirlineDetailsStruct

	Psp_Transactions *ArrayOf_RequerimientoStruct_SplitPayOnLine_3p_Transactions `xml:"psp_Transactions,omitempty"`
}

type RespuestaStruct_SplitPayOnLine_3p_Transactions struct {
	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_PromotionCode string `xml:"psp_PromotionCode,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_FirstPaymentDeferral string `xml:"psp_FirstPaymentDeferral,omitempty"`

	Psp_SoftDescriptor string `xml:"psp_SoftDescriptor,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`
}

type ArrayOf_RespuestaStruct_SplitPayOnLine_3p_Transactions struct {
	Items []*RespuestaStruct_SplitPayOnLine_3p_Transactions `xml:"item,omitempty"`
}

type ServiceRespuestaStruct_SplitPayOnLine_3p struct {
	XMLName xml.Name `xml:"SplitPayOnLine_3pResponse"`

	PSP_RespuestaStruct_SplitPayOnLine_3p *RespuestaStruct_SplitPayOnLine_3p
}

type RespuestaStruct_SplitPayOnLine_3p struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_Session3p string `xml:"psp_Session3p,omitempty"`

	Psp_FrontPSP_URL string `xml:"psp_FrontPSP_URL,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_UseMultipleProducts string `xml:"psp_UseMultipleProducts,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_Transactions *ArrayOf_RespuestaStruct_SplitPayOnLine_3p_Transactions `xml:"psp_Transactions,omitempty"`
}

type ServiceStruct_Authorize_3p struct {
	XMLName xml.Name `xml:"Authorize_3p"`

	PSP_RequerimientoStruct_Authorize_3p *RequerimientoStruct_Authorize_3p
}

func NewRequerimientoStruct_Authorize_3p() *RequerimientoStruct_Authorize_3p {
	p := new(RequerimientoStruct_Authorize_3p)
	return p
}

type RequerimientoStruct_Authorize_3p struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_Authorize_3p"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TxSource string `xml:"psp_TxSource,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_ReturnURL string `xml:"psp_ReturnURL,omitempty"`

	Psp_FrmLanguage string `xml:"psp_FrmLanguage,omitempty"`

	Psp_FrmBackButtonURL string `xml:"psp_FrmBackButtonURL,omitempty"`

	Psp_FrmTimeout string `xml:"psp_FrmTimeout,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_PurchaseDescription string `xml:"psp_PurchaseDescription,omitempty"`

	Psp_SoftDescriptor string `xml:"psp_SoftDescriptor,omitempty"`

	Psp_ForceProcessingMethod string `xml:"psp_ForceProcessingMethod,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_PromotionCode string `xml:"psp_PromotionCode,omitempty"`

	Psp_3dSecureAction string `xml:"psp_3dSecureAction,omitempty"`

	Psp_VisaArg_DA_DocType string `xml:"psp_VisaArg_DA_DocType,omitempty"`

	Psp_VisaArg_DA_DocNum string `xml:"psp_VisaArg_DA_DocNum,omitempty"`

	Psp_VisaArg_DA_DoorNum string `xml:"psp_VisaArg_DA_DoorNum,omitempty"`

	Psp_VisaArg_DA_BirthDay string `xml:"psp_VisaArg_DA_BirthDay,omitempty"`

	Psp_VisaArg_DA_Name string `xml:"psp_VisaArg_DA_Name,omitempty"`

	Psp_AmexArg_AVS_Address string `xml:"psp_AmexArg_AVS_Address,omitempty"`

	Psp_AmexArg_AVS_PostCode string `xml:"psp_AmexArg_AVS_PostCode,omitempty"`

	Psp_MasterArg_AVS_BirthDay string `xml:"psp_MasterArg_AVS_BirthDay,omitempty"`

	Psp_MasterArg_AVS_AdditionalsQtty string `xml:"psp_MasterArg_AVS_AdditionalsQtty,omitempty"`

	Psp_MasterArg_AVS_PostalCode string `xml:"psp_MasterArg_AVS_PostalCode,omitempty"`

	Psp_MasterArg_AVS_BranchEntity string `xml:"psp_MasterArg_AVS_BranchEntity,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`

	Psp_VaultReference *VaultReference3pStruct

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsRequestStruct

	Psp_MerchantAdditionalDetails *MerchantAdditionalDetailsStruct

	Psp_CustomerAdditionalDetails *CustomerAdditionalDetailsStruct

	Psp_BillingDetails *BillingDetailsStruct

	Psp_ShippingDetails *ShippingDetailsStruct

	Psp_OrderDetails *OrderDetailsStruct

	Psp_AirlineDetails *AirlineDetailsStruct
}

type ServiceRespuestaStruct_Authorize_3p struct {
	XMLName xml.Name `xml:"Authorize_3pResponse"`

	PSP_RespuestaStruct_Authorize_3p *RespuestaStruct_Authorize_3p
}

type RespuestaStruct_Authorize_3p struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_Session3p string `xml:"psp_Session3p,omitempty"`

	Psp_FrontPSP_URL string `xml:"psp_FrontPSP_URL,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_PaymentAmount string `xml:"psp_PaymentAmount,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`
}

func NewRequerimientoStruct_SplitAuthorize_3p_Transactions() *RequerimientoStruct_SplitAuthorize_3p_Transactions {
	p := new(RequerimientoStruct_SplitAuthorize_3p_Transactions)
	return p
}

type RequerimientoStruct_SplitAuthorize_3p_Transactions struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_SplitAuthorize_3p_Transactions"`
	//XMLName xml.Name `xml:"Requerimiento"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_PromotionCode string `xml:"psp_PromotionCode,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_SoftDescriptor string `xml:"psp_SoftDescriptor,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsRequestStruct
}

func NewArrayOf_RequerimientoStruct_SplitAuthorize_3p_Transactions() *ArrayOf_RequerimientoStruct_SplitAuthorize_3p_Transactions {
	p := new(ArrayOf_RequerimientoStruct_SplitAuthorize_3p_Transactions)
	p.Xsi_Type_SOAP = "SOAP-ENC:Array"
	p.Xsi_Type_SOAP_ARR = "tns:ArrayOf_RequerimientoStruct_SplitAuthorize_3p_Transactions"

	return p
}

type ArrayOf_RequerimientoStruct_SplitAuthorize_3p_Transactions struct {
	Xsi_Type_SOAP     string `xml:"xsi:type,attr"`
	Xsi_Type_SOAP_ARR string `xml:"SOAP-ENC:arrayType,attr"`

	Items []*RequerimientoStruct_SplitAuthorize_3p_Transactions `xml:"item,omitempty"`
}

type ServiceStruct_SplitAuthorize_3p struct {
	XMLName xml.Name `xml:"SplitAuthorize_3p"`

	PSP_RequerimientoStruct_SplitAuthorize_3p *RequerimientoStruct_SplitAuthorize_3p
}

func NewRequerimientoStruct_SplitAuthorize_3p() *RequerimientoStruct_SplitAuthorize_3p {
	p := new(RequerimientoStruct_SplitAuthorize_3p)
	return p
}

type RequerimientoStruct_SplitAuthorize_3p struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_SplitAuthorize_3p"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TxSource string `xml:"psp_TxSource,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_ReturnURL string `xml:"psp_ReturnURL,omitempty"`

	Psp_FrmLanguage string `xml:"psp_FrmLanguage,omitempty"`

	Psp_FrmBackButtonURL string `xml:"psp_FrmBackButtonURL,omitempty"`

	Psp_FrmTimeout string `xml:"psp_FrmTimeout,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_UseMultipleProducts string `xml:"psp_UseMultipleProducts,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_PurchaseDescription string `xml:"psp_PurchaseDescription,omitempty"`

	Psp_SoftDescriptor string `xml:"psp_SoftDescriptor,omitempty"`

	Psp_ForceProcessingMethod string `xml:"psp_ForceProcessingMethod,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_PromotionCode string `xml:"psp_PromotionCode,omitempty"`

	Psp_3dSecureAction string `xml:"psp_3dSecureAction,omitempty"`

	Psp_VisaArg_DA_DocType string `xml:"psp_VisaArg_DA_DocType,omitempty"`

	Psp_VisaArg_DA_DocNum string `xml:"psp_VisaArg_DA_DocNum,omitempty"`

	Psp_VisaArg_DA_DoorNum string `xml:"psp_VisaArg_DA_DoorNum,omitempty"`

	Psp_VisaArg_DA_BirthDay string `xml:"psp_VisaArg_DA_BirthDay,omitempty"`

	Psp_VisaArg_DA_Name string `xml:"psp_VisaArg_DA_Name,omitempty"`

	Psp_AmexArg_AVS_Address string `xml:"psp_AmexArg_AVS_Address,omitempty"`

	Psp_AmexArg_AVS_PostCode string `xml:"psp_AmexArg_AVS_PostCode,omitempty"`

	Psp_MasterArg_AVS_BirthDay string `xml:"psp_MasterArg_AVS_BirthDay,omitempty"`

	Psp_MasterArg_AVS_AdditionalsQtty string `xml:"psp_MasterArg_AVS_AdditionalsQtty,omitempty"`

	Psp_MasterArg_AVS_PostalCode string `xml:"psp_MasterArg_AVS_PostalCode,omitempty"`

	Psp_MasterArg_AVS_BranchEntity string `xml:"psp_MasterArg_AVS_BranchEntity,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`

	Psp_VaultReference *VaultReference3pStruct

	Psp_MerchantAdditionalDetails *MerchantAdditionalDetailsStruct

	Psp_CustomerAdditionalDetails *CustomerAdditionalDetailsStruct

	Psp_BillingDetails *BillingDetailsStruct

	Psp_ShippingDetails *ShippingDetailsStruct

	Psp_OrderDetails *OrderDetailsStruct

	Psp_AirlineDetails *AirlineDetailsStruct

	Psp_Transactions *ArrayOf_RequerimientoStruct_SplitAuthorize_3p_Transactions `xml:"psp_Transactions,omitempty"`
}

type RespuestaStruct_SplitAuthorize_3p_Transactions struct {
	//	XMLName xml.Name `xml:"Respuesta"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_PromotionCode string `xml:"psp_PromotionCode,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_FirstPaymentDeferral string `xml:"psp_FirstPaymentDeferral,omitempty"`

	Psp_SoftDescriptor string `xml:"psp_SoftDescriptor,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`
}

type ArrayOf_RespuestaStruct_SplitAuthorize_3p_Transactions struct {
	XMLName xml.Name `xml:"ArrayOf_RespuestaStruct_SplitAuthorize_3p_Transactions"`

	Items []*RespuestaStruct_SplitAuthorize_3p_Transactions `xml:"item,omitempty"`
}

type ServiceRespuestaStruct_SplitAuthorize_3p struct {
	XMLName xml.Name `xml:"SplitAuthorize_3pResponse"`

	PSP_RespuestaStruct_SplitAuthorize_3p *RespuestaStruct_SplitAuthorize_3p
}

type RespuestaStruct_SplitAuthorize_3p struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_Session3p string `xml:"psp_Session3p,omitempty"`

	Psp_FrontPSP_URL string `xml:"psp_FrontPSP_URL,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_UseMultipleProducts string `xml:"psp_UseMultipleProducts,omitempty"`

	Psp_Transactions *ArrayOf_RespuestaStruct_SplitAuthorize_3p_Transactions //`xml:"psp_Transactions,omitempty"`
}

type ServiceStruct_BankPayment_3p struct {
	XMLName xml.Name `xml:"BankPayment_3p"`

	PSP_RequerimientoStruct_BankPayment_3p *RequerimientoStruct_BankPayment_3p
}

func NewRequerimientoStruct_BankPayment_3p() *RequerimientoStruct_BankPayment_3p {
	p := new(RequerimientoStruct_BankPayment_3p)
	return p
}

type RequerimientoStruct_BankPayment_3p struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_BankPayment_3p"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TxSource string `xml:"psp_TxSource,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_ReturnURL string `xml:"psp_ReturnURL,omitempty"`

	Psp_FrmLanguage string `xml:"psp_FrmLanguage,omitempty"`

	Psp_FrmBackButtonURL string `xml:"psp_FrmBackButtonURL,omitempty"`

	Psp_FrmTimeout string `xml:"psp_FrmTimeout,omitempty"`

	Psp_ScreenDescription string `xml:"psp_ScreenDescription,omitempty"`

	Psp_TicketDescription string `xml:"psp_TicketDescription,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_Amount1 string `xml:"psp_Amount1,omitempty"`

	Psp_ExpDate1 string `xml:"psp_ExpDate1,omitempty"`

	Psp_Amount2 string `xml:"psp_Amount2,omitempty"`

	Psp_ExpDate2 string `xml:"psp_ExpDate2,omitempty"`

	Psp_Amount3 string `xml:"psp_Amount3,omitempty"`

	Psp_ExpDate3 string `xml:"psp_ExpDate3,omitempty"`

	Psp_MinAmount string `xml:"psp_MinAmount,omitempty"`

	Psp_ExpTime string `xml:"psp_ExpTime,omitempty"`

	Psp_ExpMark string `xml:"psp_ExpMark,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_ForceProcessingMethod string `xml:"psp_ForceProcessingMethod,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsRequestStruct

	Psp_MerchantAdditionalDetails *MerchantAdditionalDetailsStruct

	Psp_CustomerAdditionalDetails *CustomerAdditionalDetailsStruct

	Psp_BillingDetails *BillingDetailsStruct

	Psp_ShippingDetails *ShippingDetailsStruct

	Psp_OrderDetails *OrderDetailsStruct

	Psp_AirlineDetails *AirlineDetailsStruct
}

type ServiceRespuestaStruct_BankPayment_3p struct {
	XMLName xml.Name `xml:"BankPayment_3pResponse"`

	PSP_RespuestaStruct_BankPayment_3p *RespuestaStruct_BankPayment_3p
}

type RespuestaStruct_BankPayment_3p struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_Session3p string `xml:"psp_Session3p,omitempty"`

	Psp_FrontPSP_URL string `xml:"psp_FrontPSP_URL,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_ScreenDescription string `xml:"psp_ScreenDescription,omitempty"`

	Psp_TicketDescription string `xml:"psp_TicketDescription,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_Amount1 string `xml:"psp_Amount1,omitempty"`

	Psp_ExpDate1 string `xml:"psp_ExpDate1,omitempty"`

	Psp_Amount2 string `xml:"psp_Amount2,omitempty"`

	Psp_ExpDate2 string `xml:"psp_ExpDate2,omitempty"`

	Psp_Amount3 string `xml:"psp_Amount3,omitempty"`

	Psp_ExpDate3 string `xml:"psp_ExpDate3,omitempty"`

	Psp_MinAmount string `xml:"psp_MinAmount,omitempty"`

	Psp_ExpTime string `xml:"psp_ExpTime,omitempty"`

	Psp_ExpMark string `xml:"psp_ExpMark,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`
}

type ServiceStruct_CashPayment_3p struct {
	XMLName xml.Name `xml:"CashPayment_3p"`

	PSP_RequerimientoStruct_CashPayment_3p *RequerimientoStruct_CashPayment_3p
}

func NewRequerimientoStruct_CashPayment_3p() *RequerimientoStruct_CashPayment_3p {
	p := new(RequerimientoStruct_CashPayment_3p)
	return p
}

type RequerimientoStruct_CashPayment_3p struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_CashPayment_3p"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TxSource string `xml:"psp_TxSource,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_ReturnURL string `xml:"psp_ReturnURL,omitempty"`

	Psp_FrmLanguage string `xml:"psp_FrmLanguage,omitempty"`

	Psp_FrmBackButtonURL string `xml:"psp_FrmBackButtonURL,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_FirstExpDate string `xml:"psp_FirstExpDate,omitempty"`

	Psp_DaysUntilSecondExpDate string `xml:"psp_DaysUntilSecondExpDate,omitempty"`

	Psp_SurchargeAmount string `xml:"psp_SurchargeAmount,omitempty"`

	Psp_DaysAvailableToPay string `xml:"psp_DaysAvailableToPay,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_CustomerDocNum string `xml:"psp_CustomerDocNum,omitempty"`

	Psp_PurchaseDescription string `xml:"psp_PurchaseDescription,omitempty"`

	Psp_ForceProcessingMethod string `xml:"psp_ForceProcessingMethod,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsRequestStruct

	Psp_MerchantAdditionalDetails *MerchantAdditionalDetailsStruct

	Psp_CustomerAdditionalDetails *CustomerAdditionalDetailsStruct

	Psp_BillingDetails *BillingDetailsStruct

	Psp_ShippingDetails *ShippingDetailsStruct

	Psp_OrderDetails *OrderDetailsStruct

	Psp_AirlineDetails *AirlineDetailsStruct
}

type ServiceRespuestaStruct_CashPayment_3p struct {
	XMLName xml.Name `xml:"CashPayment_3pResponse"`

	PSP_RespuestaStruct_CashPayment_3p *RespuestaStruct_CashPayment_3p
}

type RespuestaStruct_CashPayment_3p struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_Session3p string `xml:"psp_Session3p,omitempty"`

	Psp_FrontPSP_URL string `xml:"psp_FrontPSP_URL,omitempty"`

	Psp_BarCode string `xml:"psp_BarCode,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_FirstExpDate string `xml:"psp_FirstExpDate,omitempty"`

	Psp_DaysUntilSecondExpDate string `xml:"psp_DaysUntilSecondExpDate,omitempty"`

	Psp_SurchargeAmount string `xml:"psp_SurchargeAmount,omitempty"`

	Psp_DaysAvailableToPay string `xml:"psp_DaysAvailableToPay,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`
}

type ServiceStruct_Capture struct {
	XMLName xml.Name `xml:"Capture"`

	PSP_RequerimientoStruct_Capture *RequerimientoStruct_Capture
}

func NewRequerimientoStruct_Capture() *RequerimientoStruct_Capture {
	p := new(RequerimientoStruct_Capture)
	return p
}

type RequerimientoStruct_Capture struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_Capture"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TxSource string `xml:"psp_TxSource,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_TransactionId_Orig string `xml:"psp_TransactionId_Orig,omitempty"`

	Psp_AmountToCapture string `xml:"psp_AmountToCapture,omitempty"`

	Psp_UserId string `xml:"psp_UserId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsRequestStruct
}

type ServiceRespuestaStruct_Capture struct {
	XMLName xml.Name `xml:"CaptureResponse"`

	PSP_RespuestaStruct_Capture *RespuestaStruct_Capture
}

type RespuestaStruct_Capture struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_TransactionId_Orig string `xml:"psp_TransactionId_Orig,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_CapturedAmount string `xml:"psp_CapturedAmount,omitempty"`

	Psp_AuthorizedAmount_Orig string `xml:"psp_AuthorizedAmount_Orig,omitempty"`

	Psp_AuthorizationCode string `xml:"psp_AuthorizationCode,omitempty"`

	Psp_BatchNro string `xml:"psp_BatchNro,omitempty"`

	Psp_TicketNumber string `xml:"psp_TicketNumber,omitempty"`

	Psp_SequenceNumber string `xml:"psp_SequenceNumber,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_ClTrnId string `xml:"psp_ClTrnId,omitempty"`

	Psp_ClExternalMerchant string `xml:"psp_ClExternalMerchant,omitempty"`

	Psp_ClExternalTerminal string `xml:"psp_ClExternalTerminal,omitempty"`

	Psp_ClResponseCod string `xml:"psp_ClResponseCod,omitempty"`

	Psp_ClResponseMsg string `xml:"psp_ClResponseMsg,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsResponseStruct
}

type ServiceStruct_Refund struct {
	XMLName xml.Name `xml:"Refund"`

	PSP_RequerimientoStruct_Refund *RequerimientoStruct_Refund
}

func NewRequerimientoStruct_Refund() *RequerimientoStruct_Refund {
	p := new(RequerimientoStruct_Refund)
	return p
}

type RequerimientoStruct_Refund struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_Refund"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TxSource string `xml:"psp_TxSource,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_TransactionId_Orig string `xml:"psp_TransactionId_Orig,omitempty"`

	Psp_AmountToRefund string `xml:"psp_AmountToRefund,omitempty"`

	Psp_CardNumber string `xml:"psp_CardNumber,omitempty"`

	Psp_CardExpDate string `xml:"psp_CardExpDate,omitempty"`

	Psp_CardSecurityCode string `xml:"psp_CardSecurityCode,omitempty"`

	Psp_UserId string `xml:"psp_UserId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type ServiceRespuestaStruct_Refund struct {
	XMLName xml.Name `xml:"RefundResponse"`

	PSP_RespuestaStruct_Refund *RespuestaStruct_Refund
}

type RespuestaStruct_Refund struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_TransactionId_Orig string `xml:"psp_TransactionId_Orig,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_RefundedAmount string `xml:"psp_RefundedAmount,omitempty"`

	Psp_CapturedAmount_Orig string `xml:"psp_CapturedAmount_Orig,omitempty"`

	Psp_AuthorizationCode string `xml:"psp_AuthorizationCode,omitempty"`

	Psp_BatchNro string `xml:"psp_BatchNro,omitempty"`

	Psp_TicketNumber string `xml:"psp_TicketNumber,omitempty"`

	Psp_SequenceNumber string `xml:"psp_SequenceNumber,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_ClTrnId string `xml:"psp_ClTrnId,omitempty"`

	Psp_ClExternalMerchant string `xml:"psp_ClExternalMerchant,omitempty"`

	Psp_ClExternalTerminal string `xml:"psp_ClExternalTerminal,omitempty"`

	Psp_ClResponseCod string `xml:"psp_ClResponseCod,omitempty"`

	Psp_ClResponseMsg string `xml:"psp_ClResponseMsg,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`
}

type ServiceStruct_GetInstallmentsOptions struct {
	XMLName xml.Name `xml:"GetInstallmentsOptions"`

	PSP_RequerimientoStruct_GetInstallmentsOptions *RequerimientoStruct_GetInstallmentsOptions
}

func NewRequerimientoStruct_GetInstallmentsOptions() *RequerimientoStruct_GetInstallmentsOptions {
	p := new(RequerimientoStruct_GetInstallmentsOptions)
	return p
}

type RequerimientoStruct_GetInstallmentsOptions struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_GetInstallmentsOptions"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_PaymentMethodToken string `xml:"psp_PaymentMethodToken,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`

	Psp_ClientSession string `xml:"psp_ClientSession,omitempty"`
}

func NewInstallmentsOptionsStruct() *InstallmentsOptionsStruct {
	p := new(InstallmentsOptionsStruct)
	return p
}

type InstallmentsOptionsStruct struct {
	XMLName xml.Name `xml:"InstallmentsOptionsStruct"`

	NumPayments string `xml:"NumPayments,omitempty"`

	InstallmentAmount string `xml:"InstallmentAmount,omitempty"`

	InterestRate string `xml:"InterestRate,omitempty"`
}

func NewArrayOf_InstallmentsOptionsStruct() *ArrayOf_InstallmentsOptionsStruct {
	p := new(ArrayOf_InstallmentsOptionsStruct)
	p.Xsi_Type_SOAP = "SOAP-ENC:Array"
	p.Xsi_Type_SOAP_ARR = "tns:ArrayOf_InstallmentsOptionsStruct"

	return p
}

type ArrayOf_InstallmentsOptionsStruct struct {
	Xsi_Type_SOAP     string `xml:"xsi:type,attr"`
	Xsi_Type_SOAP_ARR string `xml:"SOAP-ENC:arrayType,attr"`

	Items []*InstallmentsOptionsStruct `xml:"item,omitempty"`
}

type ServiceRespuestaStruct_GetInstallmentsOptions struct {
	XMLName xml.Name `xml:"GetInstallmentsOptionsResponse"`

	PSP_RespuestaStruct_GetInstallmentsOptions *RespuestaStruct_GetInstallmentsOptions
}

type RespuestaStruct_GetInstallmentsOptions struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_InstallmentsOptions *ArrayOf_InstallmentsOptionsStruct

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`
}

func NewCardInputDetailsStruct() *CardInputDetailsStruct {
	p := new(CardInputDetailsStruct)
	return p
}

type CardInputDetailsStruct struct {
	XMLName xml.Name

	Number string `xml:"Number,omitempty"`

	ExpirationDate string `xml:"ExpirationDate,omitempty"`

	SecurityCode string `xml:"SecurityCode,omitempty"`

	HolderName string `xml:"HolderName,omitempty"`
}

type ServiceStruct_CreatePaymentMethodToken struct {
	XMLName xml.Name `xml:"CreatePaymentMethodToken"`

	PSP_RequerimientoStruct_CreatePaymentMethodToken *RequerimientoStruct_CreatePaymentMethodToken
}

func NewRequerimientoStruct_CreatePaymentMethodToken() *RequerimientoStruct_CreatePaymentMethodToken {
	p := new(RequerimientoStruct_CreatePaymentMethodToken)
	return p
}

type RequerimientoStruct_CreatePaymentMethodToken struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_CreatePaymentMethodToken"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CardInputDetails *CardInputDetailsStruct

	Psp_Person *PersonStruct

	Psp_Address *AddressStruct

	Psp_ClientSession string `xml:"psp_ClientSession,omitempty"`
}

type CardOutputDetailsStruct struct {
	XMLName xml.Name `xml:"CardOutputDetailsStruct"`

	Number string `xml:"Number,omitempty"`

	ExpirationDate string `xml:"ExpirationDate,omitempty"`

	ExpirationYear string `xml:"ExpirationYear,omitempty"`

	ExpirationMonth string `xml:"ExpirationMonth,omitempty"`

	HolderName string `xml:"HolderName,omitempty"`

	IIN string `xml:"IIN,omitempty"`

	Last4 string `xml:"Last4,omitempty"`

	MaskedNumber string `xml:"MaskedNumber,omitempty"`

	MaskedNumberAlternative string `xml:"MaskedNumberAlternative,omitempty"`
}

type ServiceRespuestaStruct_CreatePaymentMethodToken struct {
	XMLName xml.Name `xml:"CreatePaymentMethodTokenResponse"`

	PSP_RespuestaStruct_CreatePaymentMethodToken *RespuestaStruct_CreatePaymentMethodToken
}

type RespuestaStruct_CreatePaymentMethodToken struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_PaymentMethodToken string `xml:"psp_PaymentMethodToken,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CardOutputDetails *CardOutputDetailsStruct

	Psp_Person *PersonStruct

	Psp_Address *AddressStruct

	Psp_AlreadyUsed string `xml:"psp_AlreadyUsed,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`

	Psp_UpdatedAt string `xml:"psp_UpdatedAt,omitempty"`
}

type ServiceStruct_RetrievePaymentMethodToken struct {
	XMLName xml.Name `xml:"RetrievePaymentMethodToken"`

	PSP_RequerimientoStruct_RetrievePaymentMethodToken *RequerimientoStruct_RetrievePaymentMethodToken
}

func NewRequerimientoStruct_RetrievePaymentMethodToken() *RequerimientoStruct_RetrievePaymentMethodToken {
	p := new(RequerimientoStruct_RetrievePaymentMethodToken)
	return p
}

type RequerimientoStruct_RetrievePaymentMethodToken struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_RetrievePaymentMethodToken"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_PaymentMethodToken string `xml:"psp_PaymentMethodToken,omitempty"`

	Psp_ClientSession string `xml:"psp_ClientSession,omitempty"`
}

type ServiceRespuestaStruct_RetrievePaymentMethodToken struct {
	XMLName xml.Name `xml:"RetrievePaymentMethodTokenResponse"`

	PSP_RespuestaStruct_RetrievePaymentMethodToken *RespuestaStruct_RetrievePaymentMethodToken
}

type RespuestaStruct_RetrievePaymentMethodToken struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_PaymentMethodToken string `xml:"psp_PaymentMethodToken,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CardOutputDetails *CardOutputDetailsStruct

	Psp_Person *PersonStruct

	Psp_Address *AddressStruct

	Psp_AlreadyUsed string `xml:"psp_AlreadyUsed,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`

	Psp_UpdatedAt string `xml:"psp_UpdatedAt,omitempty"`
}

type ServiceStruct_RecachePaymentMethodToken struct {
	XMLName xml.Name `xml:"RecachePaymentMethodToken"`

	PSP_RequerimientoStruct_RecachePaymentMethodToken *RequerimientoStruct_RecachePaymentMethodToken
}

func NewRequerimientoStruct_RecachePaymentMethodToken() *RequerimientoStruct_RecachePaymentMethodToken {
	p := new(RequerimientoStruct_RecachePaymentMethodToken)
	return p

}

type RequerimientoStruct_RecachePaymentMethodToken struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_RecachePaymentMethodToken"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_PaymentMethodId string `xml:"psp_PaymentMethodId,omitempty"`

	Psp_CardSecurityCode string `xml:"psp_CardSecurityCode,omitempty"`

	Psp_Person *PersonStruct

	Psp_Address *AddressStruct

	Psp_ClientSession string `xml:"psp_ClientSession,omitempty"`
}

type ServiceRespuestaStruct_RecachePaymentMethodToken struct {
	XMLName xml.Name `xml:"RecachePaymentMethodTokenResponse"`

	PSP_RespuestaStruct_RecachePaymentMethodToken *RespuestaStruct_RecachePaymentMethodToken
}

type RespuestaStruct_RecachePaymentMethodToken struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_PaymentMethodToken string `xml:"psp_PaymentMethodToken,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CardOutputDetails *CardOutputDetailsStruct

	Psp_Person *PersonStruct

	Psp_Address *AddressStruct

	Psp_AlreadyUsed string `xml:"psp_AlreadyUsed,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`

	Psp_UpdatedAt string `xml:"psp_UpdatedAt,omitempty"`
}

type ServiceStruct_CreateClientSession struct {
	XMLName xml.Name `xml:"CreateClientSession"`

	PSP_RequerimientoStruct_CreateClientSession *RequerimientoStruct_CreateClientSession
}

func NewRequerimientoStruct_CreateClientSession() *RequerimientoStruct_CreateClientSession {
	p := new(RequerimientoStruct_CreateClientSession)
	return p
}

type RequerimientoStruct_CreateClientSession struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_CreateClientSession"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type ServiceRespuestaStruct_CreateClientSession struct {
	XMLName xml.Name `xml:"CreateClientSessionResponse"`

	PSP_RespuestaStruct_CreateClientSession *RespuestaStruct_CreateClientSession
}

type RespuestaStruct_CreateClientSession struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_ClientSession string `xml:"psp_ClientSession,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`
}

func NewPaymentMethodInputDetailsStruct() *PaymentMethodInputDetailsStruct {
	p := new(PaymentMethodInputDetailsStruct)
	return p
}

type PaymentMethodInputDetailsStruct struct {
	XMLName xml.Name `xml:"psp_PaymentMethod"`

	PaymentMethodToken string `xml:"PaymentMethodToken,omitempty"`

	CardInputDetails *CardInputDetailsStruct

	Person *PersonStruct

	Address *AddressStruct
}

type ServiceStruct_CreatePaymentMethod struct {
	XMLName xml.Name `xml:"CreatePaymentMethod"`

	PSP_RequerimientoStruct_CreatePaymentMethod *RequerimientoStruct_CreatePaymentMethod
}

func NewRequerimientoStruct_CreatePaymentMethod() *RequerimientoStruct_CreatePaymentMethod {
	p := new(RequerimientoStruct_CreatePaymentMethod)
	return p
}

type RequerimientoStruct_CreatePaymentMethod struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_CreatePaymentMethod"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_PaymentMethod *PaymentMethodInputDetailsStruct

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_SetAsCustomerDefault string `xml:"psp_SetAsCustomerDefault,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type PaymentMethodOutputDetailsStruct struct {
	XMLName xml.Name `xml:"PaymentMethodOutputDetailsStruct"`

	PaymentMethodId string `xml:"PaymentMethodId,omitempty"`

	Product string `xml:"Product,omitempty"`

	CardOutputDetails *CardOutputDetailsStruct

	FingerPrint string `xml:"FingerPrint,omitempty"`

	Person *PersonStruct

	Address *AddressStruct

	CreatedAt string `xml:"CreatedAt,omitempty"`

	UpdatedAt string `xml:"UpdatedAt,omitempty"`
}

type ServiceRespuestaStruct_CreatePaymentMethod struct {
	XMLName xml.Name `xml:"CreatePaymentMethodResponse"`

	PSP_RespuestaStruct_CreatePaymentMethod *RespuestaStruct_CreatePaymentMethod
}

type RespuestaStruct_CreatePaymentMethod struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_PaymentMethod *PaymentMethodOutputDetailsStruct

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`
}

type ServiceStruct_CreatePaymentMethodFromPayment struct {
	XMLName xml.Name `xml:"CreatePaymentMethodFromPayment"`

	PSP_RequerimientoStruct_CreatePaymentMethodFromPayment *RequerimientoStruct_CreatePaymentMethodFromPayment
}

func NewRequerimientoStruct_CreatePaymentMethodFromPayment() *RequerimientoStruct_CreatePaymentMethodFromPayment {
	p := new(RequerimientoStruct_CreatePaymentMethodFromPayment)
	return p
}

type RequerimientoStruct_CreatePaymentMethodFromPayment struct {
	XMLName xml.Name `xml:"RequerimientoStruct_CreatePaymentMethodFromPayment"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_SetAsCustomerDefault string `xml:"psp_SetAsCustomerDefault,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type ServiceRespuestaStruct_CreatePaymentMethodFromPayment struct {
	XMLName xml.Name `xml:"CreatePaymentMethodFromPaymentResponse"`

	PSP_RespuestaStruct_CreatePaymentMethodFromPayment *RespuestaStruct_CreatePaymentMethodFromPayment
}

type RespuestaStruct_CreatePaymentMethodFromPayment struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_PaymentMethod *PaymentMethodOutputDetailsStruct

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`
}

type ServiceStruct_RetrievePaymentMethod struct {
	XMLName xml.Name `xml:"RetrievePaymentMethod"`

	PSP_RequerimientoStruct_RetrievePaymentMethod *RequerimientoStruct_RetrievePaymentMethod
}

func NewRequerimientoStruct_RetrievePaymentMethod() *RequerimientoStruct_RetrievePaymentMethod {
	p := new(RequerimientoStruct_RetrievePaymentMethod)
	return p
}

type RequerimientoStruct_RetrievePaymentMethod struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_RetrievePaymentMethod"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_PaymentMethodId string `xml:"psp_PaymentMethodId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type ServiceRespuestaStruct_RetrievePaymentMethod struct {
	XMLName xml.Name `xml:"RetrievePaymentMethodResponse"`

	PSP_RespuestaStruct_RetrievePaymentMethod *RespuestaStruct_RetrievePaymentMethod
}

type RespuestaStruct_RetrievePaymentMethod struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_PaymentMethod *PaymentMethodOutputDetailsStruct

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`
}

func NewCardInputUpdateDetailsStruct() *CardInputUpdateDetailsStruct {
	p := new(CardInputUpdateDetailsStruct)
	return p
}

type CardInputUpdateDetailsStruct struct {
	XMLName xml.Name `xml:"CardInputUpdateDetailsStruct"`

	ExpirationDate string `xml:"ExpirationDate,omitempty"`

	HolderName string `xml:"HolderName,omitempty"`
}

type ServiceStruct_UpdatePaymentMethod struct {
	XMLName xml.Name `xml:"UpdatePaymentMethod"`

	PSP_RequerimientoStruct_UpdatePaymentMethod *RequerimientoStruct_UpdatePaymentMethod
}

func NewRequerimientoStruct_UpdatePaymentMethod() *RequerimientoStruct_UpdatePaymentMethod {
	p := new(RequerimientoStruct_UpdatePaymentMethod)
	return p
}

type RequerimientoStruct_UpdatePaymentMethod struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_UpdatePaymentMethod"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_PaymentMethodId string `xml:"psp_PaymentMethodId,omitempty"`

	Psp_CardInputDetails *CardInputUpdateDetailsStruct

	Psp_Person *PersonStruct

	Psp_Address *AddressStruct

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type ServiceRespuestaStruct_UpdatePaymentMethod struct {
	XMLName xml.Name `xml:"UpdatePaymentMethodResponse"`

	PSP_RespuestaStruct_UpdatePaymentMethod *RespuestaStruct_UpdatePaymentMethod
}

type RespuestaStruct_UpdatePaymentMethod struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_PaymentMethod *PaymentMethodOutputDetailsStruct

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`
}

type ServiceStruct_DeletePaymentMethod struct {
	XMLName xml.Name `xml:"DeletePaymentMethod"`

	PSP_RequerimientoStruct_DeletePaymentMethod *RequerimientoStruct_DeletePaymentMethod
}

func NewRequerimientoStruct_DeletePaymentMethod() *RequerimientoStruct_DeletePaymentMethod {
	p := new(RequerimientoStruct_DeletePaymentMethod)
	return p
}

type RequerimientoStruct_DeletePaymentMethod struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_DeletePaymentMethod"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_PaymentMethodId string `xml:"psp_PaymentMethodId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type ServiceRespuestaStruct_DeletePaymentMethod struct {
	XMLName xml.Name `xml:"DeletePaymentMethodResponse"`

	PSP_RespuestaStruct_DeletePaymentMethod *RespuestaStruct_DeletePaymentMethod
}

type RespuestaStruct_DeletePaymentMethod struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_PaymentMethod *PaymentMethodOutputDetailsStruct

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`
}

type ServiceStruct_CreateCustomer struct {
	XMLName xml.Name `xml:"CreateCustomer"`

	PSP_RequerimientoStruct_CreateCustomer *RequerimientoStruct_CreateCustomer
}

func NewRequerimientoStruct_CreateCustomer() *RequerimientoStruct_CreateCustomer {
	p := new(RequerimientoStruct_CreateCustomer)
	return p
}

type RequerimientoStruct_CreateCustomer struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_CreateCustomer"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_EmailAddress string `xml:"psp_EmailAddress,omitempty"`

	Psp_AlternativeEmailAddress string `xml:"psp_AlternativeEmailAddress,omitempty"`

	Psp_AccountID string `xml:"psp_AccountID,omitempty"`

	Psp_AccountCreatedAt string `xml:"psp_AccountCreatedAt,omitempty"`

	Psp_Person *PersonStruct

	Psp_Address *AddressStruct

	Psp_PaymentMethod *PaymentMethodInputDetailsStruct

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type PaymentMethodsOutputDetailsStruct struct {
	XMLName xml.Name `xml:"PaymentMethodsOutputDetailsStruct"`

	PaymentMethodId string `xml:"PaymentMethodId,omitempty"`

	Product string `xml:"Product,omitempty"`

	CardOutputDetails *CardOutputDetailsStruct

	FingerPrint string `xml:"FingerPrint,omitempty"`

	Person *PersonStruct

	Address *AddressStruct

	CreatedAt string `xml:"CreatedAt,omitempty"`

	UpdatedAt string `xml:"UpdatedAt,omitempty"`
}

type ArrayOf_PaymentMethodsOutputDetailsStruct struct {
	XMLName xml.Name `xml:"ArrayOf_PaymentMethodsOutputDetailsStruct"`

	Items []*PaymentMethodsOutputDetailsStruct `xml:"item,omitempty"`
}

type ServiceRespuestaStruct_CreateCustomer struct {
	XMLName xml.Name `xml:"CreateCustomerResponse"`

	PSP_RespuestaStruct_CreateCustomer *RespuestaStruct_CreateCustomer
}

type RespuestaStruct_CreateCustomer struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_EmailAddress string `xml:"psp_EmailAddress,omitempty"`

	Psp_AlternativeEmailAddress string `xml:"psp_AlternativeEmailAddress,omitempty"`

	Psp_AccountID string `xml:"psp_AccountID,omitempty"`

	Psp_AccountCreatedAt string `xml:"psp_AccountCreatedAt,omitempty"`

	Psp_Person *PersonStruct

	Psp_Address *AddressStruct

	Psp_DefaultPaymentMethodId string `xml:"psp_DefaultPaymentMethodId,omitempty"`

	Psp_PaymentMethods *ArrayOf_PaymentMethodsOutputDetailsStruct

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`
}

type ServiceStruct_RetrieveCustomer struct {
	XMLName xml.Name `xml:"RetrieveCustomer"`

	PSP_RequerimientoStruct_RetrieveCustomer *RequerimientoStruct_RetrieveCustomer
}

func NewRequerimientoStruct_RetrieveCustomer() *RequerimientoStruct_RetrieveCustomer {
	p := new(RequerimientoStruct_RetrieveCustomer)
	return p
}

type RequerimientoStruct_RetrieveCustomer struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_RetrieveCustomer"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type ServiceRespuestaStruct_RetrieveCustomer struct {
	XMLName xml.Name `xml:"RetrieveCustomerResponse"`

	PSP_RespuestaStruct_RetrieveCustomer *RespuestaStruct_RetrieveCustomer
}

type RespuestaStruct_RetrieveCustomer struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_EmailAddress string `xml:"psp_EmailAddress,omitempty"`

	Psp_AlternativeEmailAddress string `xml:"psp_AlternativeEmailAddress,omitempty"`

	Psp_AccountID string `xml:"psp_AccountID,omitempty"`

	Psp_AccountCreatedAt string `xml:"psp_AccountCreatedAt,omitempty"`

	Psp_Person *PersonStruct

	Psp_Address *AddressStruct

	Psp_DefaultPaymentMethodId string `xml:"psp_DefaultPaymentMethodId,omitempty"`

	Psp_PaymentMethods *ArrayOf_PaymentMethodsOutputDetailsStruct

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`
}

type ServiceStruct_UpdateCustomer struct {
	XMLName xml.Name `xml:"UpdateCustomer"`

	PSP_RequerimientoStruct_UpdateCustomer *RequerimientoStruct_UpdateCustomer
}

func NewRequerimientoStruct_UpdateCustomer() *RequerimientoStruct_UpdateCustomer {
	p := new(RequerimientoStruct_UpdateCustomer)
	return p
}

type RequerimientoStruct_UpdateCustomer struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_UpdateCustomer"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_EmailAddress string `xml:"psp_EmailAddress,omitempty"`

	Psp_AlternativeEmailAddress string `xml:"psp_AlternativeEmailAddress,omitempty"`

	Psp_AccountID string `xml:"psp_AccountID,omitempty"`

	Psp_AccountCreatedAt string `xml:"psp_AccountCreatedAt,omitempty"`

	Psp_Person *PersonStruct `xml:"psp_Person,omitempty"`

	Psp_Address *AddressStruct `xml:"psp_Address,omitempty"`

	Psp_DefaultPaymentMethodId string `xml:"psp_DefaultPaymentMethodId,omitempty"`

	Psp_PaymentMethod *PaymentMethodInputDetailsStruct

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type ServiceRespuestaStruct_UpdateCustomer struct {
	XMLName xml.Name `xml:"UpdateCustomerResponse"`

	PSP_RespuestaStruct_UpdateCustomer *RespuestaStruct_UpdateCustomer
}

type RespuestaStruct_UpdateCustomer struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_EmailAddress string `xml:"psp_EmailAddress,omitempty"`

	Psp_AlternativeEmailAddress string `xml:"psp_AlternativeEmailAddress,omitempty"`

	Psp_AccountID string `xml:"psp_AccountID,omitempty"`

	Psp_AccountCreatedAt string `xml:"psp_AccountCreatedAt,omitempty"`

	Psp_Person *PersonStruct `xml:"psp_Person,omitempty"`

	Psp_Address *AddressStruct `xml:"psp_Address,omitempty"`

	Psp_DefaultPaymentMethodId string `xml:"psp_DefaultPaymentMethodId,omitempty"`

	Psp_PaymentMethods *ArrayOf_PaymentMethodsOutputDetailsStruct

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`
}

type ServiceStruct_DeleteCustomer struct {
	XMLName xml.Name `xml:"DeleteCustomer"`

	PSP_RequerimientoStruct_DeleteCustomer *RequerimientoStruct_DeleteCustomer
}

func NewRequerimientoStruct_DeleteCustomer() *RequerimientoStruct_DeleteCustomer {
	p := new(RequerimientoStruct_DeleteCustomer)
	return p
}

type RequerimientoStruct_DeleteCustomer struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_DeleteCustomer"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type ServiceRespuestaStruct_DeleteCustomer struct {
	XMLName xml.Name `xml:"DeleteCustomerResponse"`

	PSP_RespuestaStruct_DeleteCustomer *RespuestaStruct_DeleteCustomer
}

type RespuestaStruct_DeleteCustomer struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_EmailAddress string `xml:"psp_EmailAddress,omitempty"`

	Psp_AlternativeEmailAddress string `xml:"psp_AlternativeEmailAddress,omitempty"`

	Psp_AccountID string `xml:"psp_AccountID,omitempty"`

	Psp_AccountCreatedAt string `xml:"psp_AccountCreatedAt,omitempty"`

	Psp_Person *PersonStruct

	Psp_Address *AddressStruct

	Psp_DefaultPaymentMethodId string `xml:"psp_DefaultPaymentMethodId,omitempty"`

	Psp_PaymentMethods *ArrayOf_PaymentMethodsOutputDetailsStruct

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`
}

type ServiceStruct_SimpleQueryTx struct {
	XMLName xml.Name `xml:"SimpleQueryTx"`

	PSP_RequerimientoStruct_SimpleQueryTx *RequerimientoStruct_SimpleQueryTx
}

func NewRequerimientoStruct_SimpleQueryTx() *RequerimientoStruct_SimpleQueryTx {
	p := new(RequerimientoStruct_SimpleQueryTx)
	return p
}

type RequerimientoStruct_SimpleQueryTx struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_SimpleQueryTx"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_QueryCriteria string `xml:"psp_QueryCriteria,omitempty"`

	Psp_QueryCriteriaId string `xml:"psp_QueryCriteriaId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type RespuestaStruct_SimpleQueryTx_Transactions struct {
	XMLName xml.Name `xml:"psp_Transaction"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TxSource string `xml:"psp_TxSource,omitempty"`

	Psp_Operation string `xml:"psp_Operation,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_PaymentAmount string `xml:"psp_PaymentAmount,omitempty"`

	Psp_FirstPaymentDeferral string `xml:"psp_FirstPaymentDeferral,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_AuthorizationCode string `xml:"psp_AuthorizationCode,omitempty"`

	Psp_BatchNro string `xml:"psp_BatchNro,omitempty"`

	Psp_SequenceNumber string `xml:"psp_SequenceNumber,omitempty"`

	Psp_TicketNumber string `xml:"psp_TicketNumber,omitempty"`

	Psp_CardNumber_FSD string `xml:"psp_CardNumber_FSD,omitempty"`

	Psp_CardNumber_LFD string `xml:"psp_CardNumber_LFD,omitempty"`

	Psp_CardMask string `xml:"psp_CardMask,omitempty"`

	Psp_CardHolderName string `xml:"psp_CardHolderName,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_CustomerIpAddress string `xml:"psp_CustomerIpAddress,omitempty"`

	Psp_CustomerHttpUserAgent string `xml:"psp_CustomerHttpUserAgent,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_ClTrnId string `xml:"psp_ClTrnId,omitempty"`

	Psp_ClExternalMerchant string `xml:"psp_ClExternalMerchant,omitempty"`

	Psp_ClExternalTerminal string `xml:"psp_ClExternalTerminal,omitempty"`

	Psp_ClResponseCod string `xml:"psp_ClResponseCod,omitempty"`

	Psp_ClResponseMsg string `xml:"psp_ClResponseMsg,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_PurchaseDescription string `xml:"psp_PurchaseDescription,omitempty"`

	Psp_SoftDescriptor string `xml:"psp_SoftDescriptor,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_VisaArg_VBV_Secured string `xml:"psp_VisaArg_VBV_Secured,omitempty"`

	Psp_3dSecure_XID string `xml:"psp_3dSecure_XID,omitempty"`

	Psp_3dSecure_CAVV string `xml:"psp_3dSecure_CAVV,omitempty"`

	Psp_3dSecure_Eci string `xml:"psp_3dSecure_Eci,omitempty"`

	Psp_3dSecure_EciMsg string `xml:"psp_3dSecure_EciMsg,omitempty"`

	Psp_3dSecure_Secured string `xml:"psp_3dSecure_Secured,omitempty"`

	Psp_VisaArg_DA_Result string `xml:"psp_VisaArg_DA_Result,omitempty"`

	Psp_AmexArg_AVS_Result string `xml:"psp_AmexArg_AVS_Result,omitempty"`

	Psp_MasterArg_AVS_Result string `xml:"psp_MasterArg_AVS_Result,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsResponseStruct

	Psp_FraudScreeningResult *FraudScreeningResultStruct

	Psp_VerificationServicesResult *VerificationServicesResultStruct

	Psp_MerchantAdditionalDetails *MerchantAdditionalDetailsStruct

	Psp_CustomerAdditionalDetails *CustomerAdditionalDetailsStruct

	Psp_BillingDetails *BillingDetailsStruct

	Psp_ShippingDetails *ShippingDetailsStruct

	Psp_OrderDetails *OrderDetailsStruct

	Psp_AirlineDetails *AirlineDetailsStruct
}

type ServiceRespuestaStruct_SimpleQueryTx struct {
	XMLName xml.Name `xml:"SimpleQueryTxResponse"`

	PSP_RespuestaStruct_SimpleQueryTx *RespuestaStruct_SimpleQueryTx
}

type RespuestaStruct_SimpleQueryTx struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_QueryCriteria string `xml:"psp_QueryCriteria,omitempty"`

	Psp_QueryCriteriaId string `xml:"psp_QueryCriteriaId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_Transaction *RespuestaStruct_SimpleQueryTx_Transactions
}

type ServiceStruct_QueryCardNumber struct {
	XMLName xml.Name `xml:"QueryCardNumber"`

	PSP_RequerimientoStruct_QueryCardNumber *RequerimientoStruct_QueryCardNumber
}

func NewRequerimientoStruct_QueryCardNumber() *RequerimientoStruct_QueryCardNumber {
	p := new(RequerimientoStruct_QueryCardNumber)
	return p
}

type RequerimientoStruct_QueryCardNumber struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_QueryCardNumber"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_QueryCriteria string `xml:"psp_QueryCriteria,omitempty"`

	Psp_QueryCriteriaId string `xml:"psp_QueryCriteriaId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type ServiceRespuestaStruct_QueryCardNumber struct {
	XMLName xml.Name `xml:"QueryCardNumberResponse"`

	PSP_RespuestaStruct_QueryCardNumber *RespuestaStruct_QueryCardNumber
}

type RespuestaStruct_QueryCardNumber struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_QueryCriteria string `xml:"psp_QueryCriteria,omitempty"`

	Psp_QueryCriteriaId string `xml:"psp_QueryCriteriaId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_CardNumber string `xml:"psp_CardNumber,omitempty"`
}

type ServiceStruct_QueryCardDetails struct {
	XMLName xml.Name `xml:"QueryCardDetails"`

	PSP_RequerimientoStruct_QueryCardDetails *RequerimientoStruct_QueryCardDetails
}

func NewRequerimientoStruct_QueryCardDetails() *RequerimientoStruct_QueryCardDetails {
	p := new(RequerimientoStruct_QueryCardDetails)
	return p
}

type RequerimientoStruct_QueryCardDetails struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_QueryCardDetails"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_QueryCriteria string `xml:"psp_QueryCriteria,omitempty"`

	Psp_QueryCriteriaId string `xml:"psp_QueryCriteriaId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type ServiceRespuestaStruct_QueryCardDetails struct {
	XMLName xml.Name `xml:"QueryCardDetailsResponse"`

	PSP_RespuestaStruct_QueryCardDetails *RespuestaStruct_QueryCardDetails
}

type RespuestaStruct_QueryCardDetails struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_QueryCriteria string `xml:"psp_QueryCriteria,omitempty"`

	Psp_QueryCriteriaId string `xml:"psp_QueryCriteriaId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_CardOutputDetails *CardOutputDetailsStruct
}

type ServiceStruct_QueryTxs struct {
	XMLName xml.Name `xml:"QueryTxs"`

	PSP_RequerimientoStruct_QueryTxs *RequerimientoStruct_QueryTxs
}

func NewRequerimientoStruct_QueryTxs() *RequerimientoStruct_QueryTxs {
	p := new(RequerimientoStruct_QueryTxs)
	return p
}

type RequerimientoStruct_QueryTxs struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_QueryTxs"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_QueryCriteria string `xml:"psp_QueryCriteria,omitempty"`

	Psp_QueryCriteriaId string `xml:"psp_QueryCriteriaId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type RespuestaStruct_QueryTxs_Transactions struct {
	//XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_TransactionId string `xml:"psp_TransactionId,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TxSource string `xml:"psp_TxSource,omitempty"`

	Psp_Operation string `xml:"psp_Operation,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_PaymentAmount string `xml:"psp_PaymentAmount,omitempty"`

	Psp_FirstPaymentDeferral string `xml:"psp_FirstPaymentDeferral,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_AuthorizationCode string `xml:"psp_AuthorizationCode,omitempty"`

	Psp_BatchNro string `xml:"psp_BatchNro,omitempty"`

	Psp_SequenceNumber string `xml:"psp_SequenceNumber,omitempty"`

	Psp_TicketNumber string `xml:"psp_TicketNumber,omitempty"`

	Psp_CardNumber_FSD string `xml:"psp_CardNumber_FSD,omitempty"`

	Psp_CardNumber_LFD string `xml:"psp_CardNumber_LFD,omitempty"`

	Psp_CardMask string `xml:"psp_CardMask,omitempty"`

	Psp_CardHolderName string `xml:"psp_CardHolderName,omitempty"`

	Psp_CustomerMail string `xml:"psp_CustomerMail,omitempty"`

	Psp_CustomerId string `xml:"psp_CustomerId,omitempty"`

	Psp_CustomerHttpUserAgent string `xml:"psp_CustomerHttpUserAgent,omitempty"`

	Psp_CustomerIpAddress string `xml:"psp_CustomerIpAddress,omitempty"`

	Psp_MerchantMail string `xml:"psp_MerchantMail,omitempty"`

	Psp_ClTrnId string `xml:"psp_ClTrnId,omitempty"`

	Psp_ClExternalMerchant string `xml:"psp_ClExternalMerchant,omitempty"`

	Psp_ClExternalTerminal string `xml:"psp_ClExternalTerminal,omitempty"`

	Psp_ClResponseCod string `xml:"psp_ClResponseCod,omitempty"`

	Psp_ClResponseMsg string `xml:"psp_ClResponseMsg,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_PurchaseDescription string `xml:"psp_PurchaseDescription,omitempty"`

	Psp_SoftDescriptor string `xml:"psp_SoftDescriptor,omitempty"`

	Psp_Plan string `xml:"psp_Plan,omitempty"`

	Psp_VisaArg_VBV_Secured string `xml:"psp_VisaArg_VBV_Secured,omitempty"`

	Psp_3dSecure_XID string `xml:"psp_3dSecure_XID,omitempty"`

	Psp_3dSecure_CAVV string `xml:"psp_3dSecure_CAVV,omitempty"`

	Psp_3dSecure_Eci string `xml:"psp_3dSecure_Eci,omitempty"`

	Psp_3dSecure_EciMsg string `xml:"psp_3dSecure_EciMsg,omitempty"`

	Psp_3dSecure_Secured string `xml:"psp_3dSecure_Secured,omitempty"`

	Psp_VisaArg_DA_Result string `xml:"psp_VisaArg_DA_Result,omitempty"`

	Psp_AmexArg_AVS_Result string `xml:"psp_AmexArg_AVS_Result,omitempty"`

	Psp_MasterArg_AVS_Result string `xml:"psp_MasterArg_AVS_Result,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`

	Psp_AmountAdditionalDetails *AmountAdditionalDetailsResponseStruct

	Psp_FraudScreeningResult *FraudScreeningResultStruct

	Psp_VerificationServicesResult *VerificationServicesResultStruct

	Psp_MerchantAdditionalDetails *MerchantAdditionalDetailsStruct

	Psp_CustomerAdditionalDetails *CustomerAdditionalDetailsStruct

	Psp_BillingDetails *BillingDetailsStruct

	Psp_ShippingDetails *ShippingDetailsStruct

	Psp_OrderDetails *OrderDetailsStruct

	Psp_AirlineDetails *AirlineDetailsStruct
}

type ArrayOf_RespuestaStruct_QueryTxs_Transactions struct {
	//XMLName xml.Name `xml:"ArrayOf_RespuestaStruct_QueryTxs_Transactions"`

	Items []*RespuestaStruct_QueryTxs_Transactions `xml:"item,omitempty"`
}

type ServiceRespuestaStruct_QueryTxs struct {
	XMLName xml.Name `xml:"QueryTxsResponse"`

	PSP_RespuestaStruct_QueryTxs *RespuestaStruct_QueryTxs
}

type RespuestaStruct_QueryTxs struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_QueryCriteria string `xml:"psp_QueryCriteria,omitempty"`

	Psp_QueryCriteriaId string `xml:"psp_QueryCriteriaId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_Transactions *ArrayOf_RespuestaStruct_QueryTxs_Transactions `xml:"psp_Transactions,omitempty"`
}

type ServiceStruct_GetIINDetails struct {
	XMLName xml.Name `xml:"GetIINDetails"`

	PSP_RequerimientoStruct_GetIINDetails *RequerimientoStruct_GetIINDetails
}

func NewRequerimientoStruct_GetIINDetails() *RequerimientoStruct_GetIINDetails {
	p := new(RequerimientoStruct_GetIINDetails)
	return p
}

type RequerimientoStruct_GetIINDetails struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_GetIINDetails"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_IIN string `xml:"psp_IIN,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`

	Psp_ClientSession string `xml:"psp_ClientSession,omitempty"`
}

type ServiceRespuestaStruct_GetIINDetails struct {
	XMLName xml.Name `xml:"GetIINDetailsResponse"`

	PSP_RespuestaStruct_GetIINDetails *RespuestaStruct_GetIINDetails
}

type RespuestaStruct_GetIINDetails struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_FinancialInstitution string `xml:"psp_FinancialInstitution,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`
}

type ServiceStruct_ChangeSecretKey struct {
	XMLName xml.Name `xml:"ChangeSecretKey"`

	PSP_RequerimientoStruct_ChangeSecretKey *RequerimientoStruct_ChangeSecretKey
}

func NewRequerimientoStruct_ChangeSecretKey() *RequerimientoStruct_ChangeSecretKey {
	p := new(RequerimientoStruct_ChangeSecretKey)
	return p
}

type RequerimientoStruct_ChangeSecretKey struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_ChangeSecretKey"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_NewSecretKey string `xml:"psp_NewSecretKey,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type ServiceRespuestaStruct_ChangeSecretKey struct {
	XMLName xml.Name `xml:"ChangeSecretKeyResponse"`

	PSP_RespuestaStruct_ChangeSecretKey *RespuestaStruct_ChangeSecretKey
}

type RespuestaStruct_ChangeSecretKey struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`
}

type ServiceStruct_FraudScreening struct {
	XMLName xml.Name `xml:"FraudScreening"`

	PSP_RequerimientoStruct_FraudScreening *RequerimientoStruct_FraudScreening
}

func NewRequerimientoStruct_FraudScreening() *RequerimientoStruct_FraudScreening {
	p := new(RequerimientoStruct_FraudScreening)
	return p
}

type RequerimientoStruct_FraudScreening struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_FraudScreening"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_TxSource string `xml:"psp_TxSource,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_FirstPaymentDeferral string `xml:"psp_FirstPaymentDeferral,omitempty"`

	Psp_Recurrent string `xml:"psp_Recurrent,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CardNumber string `xml:"psp_CardNumber,omitempty"`

	Psp_CardExpDate string `xml:"psp_CardExpDate,omitempty"`

	Psp_CardHolderName string `xml:"psp_CardHolderName,omitempty"`

	Psp_PurchaseDescription string `xml:"psp_PurchaseDescription,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`

	Psp_MerchantAdditionalDetails *MerchantAdditionalDetailsStruct

	Psp_CustomerAdditionalDetails *CustomerAdditionalDetailsStruct

	Psp_BillingDetails *BillingDetailsStruct

	Psp_ShippingDetails *ShippingDetailsStruct

	Psp_OrderDetails *OrderDetailsStruct

	Psp_AirlineDetails *AirlineDetailsStruct
}

type ServiceRespuestaStruct_FraudScreening struct {
	XMLName xml.Name `xml:"FraudScreeningResponse"`

	PSP_RespuestaStruct_FraudScreening *RespuestaStruct_FraudScreening
}

type RespuestaStruct_FraudScreening struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_Result *FraudScreeningResultStruct

	Psp_OrderId string `xml:"psp_OrderId,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_MerchTxRef string `xml:"psp_MerchTxRef,omitempty"`

	Psp_MerchOrderId string `xml:"psp_MerchOrderId,omitempty"`

	Psp_MerchAdditionalRef string `xml:"psp_MerchAdditionalRef,omitempty"`

	Psp_Amount string `xml:"psp_Amount,omitempty"`

	Psp_NumPayments string `xml:"psp_NumPayments,omitempty"`

	Psp_FirstPaymentDeferral string `xml:"psp_FirstPaymentDeferral,omitempty"`

	Psp_Recurrent string `xml:"psp_Recurrent,omitempty"`

	Psp_Currency string `xml:"psp_Currency,omitempty"`

	Psp_Country string `xml:"psp_Country,omitempty"`

	Psp_Product string `xml:"psp_Product,omitempty"`

	Psp_CardNumber string `xml:"psp_CardNumber,omitempty"`

	Psp_CardExpDate string `xml:"psp_CardExpDate,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_CreatedAt string `xml:"psp_CreatedAt,omitempty"`
}

type ServiceStruct_NotifyFraudScreeningReview struct {
	XMLName xml.Name `xml:"NotifyFraudScreeningReview"`

	PSP_RequerimientoStruct_NotifyFraudScreeningReview *RequerimientoStruct_NotifyFraudScreeningReview
}

func NewRequerimientoStruct_NotifyFraudScreeningReview() *RequerimientoStruct_NotifyFraudScreeningReview {
	p := new(RequerimientoStruct_NotifyFraudScreeningReview)
	return p
}

type RequerimientoStruct_NotifyFraudScreeningReview struct {
	//XMLName xml.Name `xml:"RequerimientoStruct_NotifyFraudScreeningReview"`
	XMLName xml.Name `xml:"Requerimiento"`

	Psp_Version string `xml:"psp_Version,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_Criteria string `xml:"psp_Criteria,omitempty"`

	Psp_CriteriaId string `xml:"psp_CriteriaId,omitempty"`

	Psp_ReviewResult string `xml:"psp_ReviewResult,omitempty"`

	Psp_UserId string `xml:"psp_UserId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`

	Psp_SecureHash string `xml:"psp_SecureHash,omitempty"`
}

type ServiceRespuestaStruct_NotifyFraudScreeningReview struct {
	XMLName xml.Name `xml:"NotifyFraudScreeningReviewResponse"`

	PSP_RespuestaStruct_NotifyFraudScreeningReview *RespuestaStruct_NotifyFraudScreeningReview
}

type RespuestaStruct_NotifyFraudScreeningReview struct {
	XMLName xml.Name `xml:"Respuesta"`

	Psp_ResponseCod string `xml:"psp_ResponseCod,omitempty"`

	Psp_ResponseMsg string `xml:"psp_ResponseMsg,omitempty"`

	Psp_ResponseExtended string `xml:"psp_ResponseExtended,omitempty"`

	Psp_MerchantId string `xml:"psp_MerchantId,omitempty"`

	Psp_Criteria string `xml:"psp_Criteria,omitempty"`

	Psp_CriteriaId string `xml:"psp_CriteriaId,omitempty"`

	Psp_PosDateTime string `xml:"psp_PosDateTime,omitempty"`
}
