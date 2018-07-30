package npsSdk

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	CONSTANTS "github.com/Ingenico-NPS-Latam/nps-sdk-go/npsSdk/constants"
	"reflect"
	"sort"
	"strings"
)

var GetMerchDetNotAddServices = map[string]bool{
	CONSTANTS.QUERY_TXS:                          true,
	CONSTANTS.REFUND:                             true,
	CONSTANTS.SIMPLE_QUERY_TX:                    true,
	CONSTANTS.CAPTURE:                            true,
	CONSTANTS.CHANGE_SECRET_KEY:                  true,
	CONSTANTS.NOTIFY_FRAUD_SCREENING_REVIEW:      true,
	CONSTANTS.GET_IIN_DETAILS:                    true,
	CONSTANTS.QUERY_CARD_NUMBER:                  true,
	CONSTANTS.CREATE_PAYMENT_METHOD:              true,
	CONSTANTS.CREATE_PAYMENT_METHOD_FROM_PAYMENT: true,
	CONSTANTS.RETRIEVE_PAYMENT_METHOD:            true,
	CONSTANTS.UPDATE_PAYMENT_METHOD:              true,
	CONSTANTS.DELETE_PAYMENT_METHOD:              true,
	CONSTANTS.CREATE_CUSTOMER:                    true,
	CONSTANTS.RETRIEVE_CUSTOMER:                  true,
	CONSTANTS.UPDATE_CUSTOMER:                    true,
	CONSTANTS.DELETE_CUSTOMER:                    true,
	CONSTANTS.RECACHE_PAYMENT_METHOD_TOKEN:       true,
	CONSTANTS.CREATE_PAYMENT_METHOD_TOKEN:        true,
	CONSTANTS.RETRIEVE_PAYMENT_METHOD_TOKEN:      true,
	CONSTANTS.CREATE_CLIENT_SESSION:              true,
	CONSTANTS.GET_INSTALLMENTS_OPTIONS:           true,
	CONSTANTS.QUERY_CARD_DETAILS:                 true,
}

var key_config = map[string]int{
	"psp_Person.FirstName.max_length":                                 128,
	"psp_Person.LastName.max_length":                                  64,
	"psp_Person.MiddleName.max_length":                                64,
	"psp_Person.PhoneNumber1.max_length":                              32,
	"psp_Person.PhoneNumber2.max_length":                              32,
	"psp_Person.Gender.max_length":                                    1,
	"psp_Person.Nationality.max_length":                               3,
	"psp_Person.IDNumber.max_length":                                  40,
	"psp_Person.IDType.max_length":                                    5,
	"psp_Address.Street.max_length":                                   128,
	"psp_Address.HouseNumber.max_length":                              32,
	"psp_Address.AdditionalInfo.max_length":                           128,
	"psp_Address.City.max_length":                                     40,
	"psp_Address.StateProvince.max_length":                            40,
	"psp_Address.Country.max_length":                                  3,
	"psp_Address.ZipCode.max_length":                                  10,
	"psp_OrderItem.Description.max_length":                            127,
	"psp_OrderItem.Type.max_length":                                   20,
	"psp_OrderItem.SkuCode.max_length":                                48,
	"psp_OrderItem.ManufacturerPartNumber.max_length":                 30,
	"psp_OrderItem.Risk.max_length":                                   1,
	"psp_Leg.DepartureAirport.max_length":                             3,
	"psp_Leg.ArrivalAirport.max_length":                               3,
	"psp_Leg.CarrierCode.max_length":                                  2,
	"psp_Leg.FlightNumber.max_length":                                 5,
	"psp_Leg.FareBasisCode.max_length":                                15,
	"psp_Leg.FareClassCode.max_length":                                3,
	"psp_Leg.BaseFareCurrency.max_length":                             3,
	"psp_Passenger.FirstName.max_length":                              50,
	"psp_Passenger.LastName.max_length":                               30,
	"psp_Passenger.MiddleName.max_length":                             30,
	"psp_Passenger.Type.max_length":                                   1,
	"psp_Passenger.Nationality.max_length":                            3,
	"psp_Passenger.IDNumber.max_length":                               40,
	"psp_Passenger.IDType.max_length":                                 10,
	"psp_Passenger.IDCountry.max_length":                              3,
	"psp_Passenger.LoyaltyNumber.max_length":                          20,
	"psp_SellerDetails.IDNumber.max_length":                           40,
	"psp_SellerDetails.IDType.max_length":                             10,
	"psp_SellerDetails.Name.max_length":                               128,
	"psp_SellerDetails.Invoice.max_length":                            32,
	"psp_SellerDetails.PurchaseDescription.max_length":                32,
	"psp_SellerDetails.MCC.max_length":                                5,
	"psp_SellerDetails.ChannelCode.max_length":                        3,
	"psp_SellerDetails.GeoCode.max_length":                            5,
	"psp_TaxesRequest.TypeId.max_length":                              5,
	"psp_MerchantAdditionalDetails.Type.max_length":                   1,
	"psp_MerchantAdditionalDetails.SdkInfo.max_length":                48,
	"psp_MerchantAdditionalDetails.ShoppingCartInfo.max_length":       48,
	"psp_MerchantAdditionalDetails.ShoppingCartPluginInfo.max_length": 48,
	"psp_CustomerAdditionalDetails.IPAddress.max_length":              45,
	"psp_CustomerAdditionalDetails.AccountID.max_length":              128,
	"psp_CustomerAdditionalDetails.DeviceFingerPrint.max_length":      4000,
	"psp_CustomerAdditionalDetails.BrowserLanguage.max_length":        2,
	"psp_CustomerAdditionalDetails.HttpUserAgent.max_length":          255,
	"psp_BillingDetails.Invoice.max_length":                           32,
	"psp_BillingDetails.InvoiceCurrency.max_length":                   3,
	"psp_ShippingDetails.TrackingNumber.max_length":                   24,
	"psp_ShippingDetails.Method.max_length":                           3,
	"psp_ShippingDetails.Carrier.max_length":                          3,
	"psp_ShippingDetails.GiftMessage.max_length":                      200,
	"psp_AirlineDetails.TicketNumber.max_length":                      14,
	"psp_AirlineDetails.PNR.max_length":                               10,
	"psp_VaultReference.PaymentMethodToken.max_length":                64,
	"psp_VaultReference.PaymentMethodId.max_length":                   64,
	"psp_VaultReference.CustomerId.max_length":                        64,
}

func AddSecureHashIfNoClientSession(request interface{}) error {

	var secureHashIdx int = -1
	var sessionIdx int = -1
	s := reflect.ValueOf(request).Elem()
	typeOfT := s.Type()

	m := make(map[string]string)
	keys := make([]string, 0)

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		if f.Kind().String() == "ptr" {

			if len(typeOfT.Field(i).Name) > 16 && typeOfT.Field(i).Name[:17] == "PSP_Requerimiento" {
				err := AddSecureHashIfNoClientSession(f.Interface())
				return err
			}

		} else {

			if f.Type().String() == "string" {
				if len(f.Interface().(string)) > 0 {
					if typeOfT.Field(i).Name == "Psp_ClientSession" {
						sessionIdx = i
						return nil
					}

					m[typeOfT.Field(i).Name] = f.Interface().(string)
					keys = append(keys, typeOfT.Field(i).Name)
				}
			}

			if typeOfT.Field(i).Name == "Psp_SecureHash" {
				secureHashIdx = i
			}
		}
	}

	if secureHashIdx < 0 && sessionIdx < 0 {
		return errors.New("Psp_ClientSession or Psp_SecureHash must be informed")
	}

	sort.Strings(keys)

	var buffer bytes.Buffer
	for i := 0; i < len(keys); i++ {
		buffer.WriteString(m[keys[i]])
	}

	buffer.WriteString(Configuration.GetSecretKey())

	data := []byte(buffer.String())

	hash := md5.Sum(data)

	var secureHash string = hex.EncodeToString(hash[:])
	s.Field(secureHashIdx).SetString(secureHash)
	return nil

}

func Sanitize(request interface{}, keyName string, level int) {

	s := reflect.ValueOf(request).Elem()

	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		if f.Kind().String() == "ptr" {

			var buffer bytes.Buffer
			buffer.WriteString("psp_")

			if typeOfT.Field(i).Name[:4] == "Psp_" {
				buffer.WriteString(typeOfT.Field(i).Name[4:])
			} else {
				buffer.WriteString(typeOfT.Field(i).Name)
			}

			if f.Interface() != reflect.Zero(f.Type()).Interface() {
				level = level + 1
				Sanitize(f.Interface(), buffer.String(), level)
			}

		} else {
			if f.Type().String() == "string" {

				if len(f.Interface().(string)) > 0 {
					var buffer bytes.Buffer
					buffer.WriteString(keyName)
					buffer.WriteString(".")
					buffer.WriteString(typeOfT.Field(i).Name)
					buffer.WriteString(".")
					buffer.WriteString("max_length")

					var max_length = key_config[buffer.String()]

					if (max_length > 0) && (len(f.Interface().(string)) > max_length) {

						var newValue string = f.String()
						newValue = newValue[:key_config[buffer.String()]]
						s.Field(i).SetString(newValue)
					}
				}
			}
		}
	}
}

func AddExtraInf(request interface{}) {

	s := reflect.ValueOf(request).Elem()

	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		if f.Kind().String() == "ptr" {

			if len(typeOfT.Field(i).Name) > 16 && typeOfT.Field(i).Name[:17] == "PSP_Requerimiento" {
				AddExtraInf(f.Interface())
				return
			}

			if typeOfT.Field(i).Name == "Psp_MerchantAdditionalDetails" {
				if f.Interface() == reflect.Zero(f.Type()).Interface() {
					merStruct := reflect.New(f.Type().Elem())
					s.Field(i).Set(merStruct)
				}

				AddExtraInf(f.Interface())
				return
			}

		} else {
			if f.Type().String() == "string" {

				if typeOfT.Field(i).Name == "SdkInfo" {
					var buffer bytes.Buffer
					buffer.WriteString(CONSTANTS.SDK_NAME)
					buffer.WriteString(" SDK ")
					buffer.WriteString("Version: ")
					buffer.WriteString(CONSTANTS.SDK_VERSION)

					var sdkInfo string = buffer.String()
					s.Field(i).SetString(sdkInfo)
					return
				}
			}
		}
	}

}

func getStringInBetween(str string, start string, end string) string {
	s := strings.Index(str, start)
	if s == -1 {
		return ""
	}
	s += len(start)
	e := strings.Index(str, end)
	return str[s:e]
}

func mask_card_number(data string) string {
	var masked_card string = ""
	if len(data) > 10 {
		masked_card = data[:6] + strings.Repeat("*", len(data)-10) + data[len(data)-4:]
	} else {
		masked_card = strings.Repeat("*", len(data))
	}

	return masked_card
}

func mask_exp_date(data string) string {

	return "****"
}

func mask_cvc(data string) string {

	return strings.Repeat("*", len(data))
}

func MaskData(format string, args ...interface{}) string {

	var data string = fmt.Sprintf(format, args...)

	var cardNumber = getStringInBetween(data, "<psp_CardNumber>", "</psp_CardNumber>")
	var expDate = getStringInBetween(data, "<psp_CardExpDate>", "</psp_CardExpDate>")
	var cvc = getStringInBetween(data, "<psp_CardSecurityCode>", "</psp_CardSecurityCode>")
	var token_cNumber = getStringInBetween(data, "<Number>", "</Number>")
	var token_expDate = getStringInBetween(data, "<ExpirationDate>", "</ExpirationDate>")
	var token_cvc = getStringInBetween(data, "<SecurityCode>", "</SecurityCode>")

	rCardNumber := strings.NewReplacer("<psp_CardNumber>"+cardNumber+"</psp_CardNumber>", "<psp_CardNumber>"+mask_card_number(cardNumber)+"</psp_CardNumber>")
	data = rCardNumber.Replace(data)

	rExpDate := strings.NewReplacer("<psp_CardExpDate>"+expDate+"</psp_CardExpDate>", "<psp_CardExpDate>"+mask_exp_date(expDate)+"</psp_CardExpDate>")
	data = rExpDate.Replace(data)

	rCvc := strings.NewReplacer("<psp_CardSecurityCode>"+cvc+"</psp_CardSecurityCode>", "<psp_CardSecurityCode>"+mask_cvc(cvc)+"</psp_CardSecurityCode>")
	data = rCvc.Replace(data)

	rTokenCNumber := strings.NewReplacer("<Number>"+token_cNumber+"</Number>", "<Number>"+mask_card_number(token_cNumber)+"</Number>")
	data = rTokenCNumber.Replace(data)

	rTokenExpDate := strings.NewReplacer("<ExpirationDate>"+token_expDate+"</ExpirationDate>", "<ExpirationDate>"+mask_exp_date(token_expDate)+"</ExpirationDate>")
	data = rTokenExpDate.Replace(data)

	rTokenCvc := strings.NewReplacer("<SecurityCode>"+token_cvc+"</SecurityCode>", "<SecurityCode>"+mask_cvc(token_cvc)+"</SecurityCode>")
	data = rTokenCvc.Replace(data)

	return data

}
