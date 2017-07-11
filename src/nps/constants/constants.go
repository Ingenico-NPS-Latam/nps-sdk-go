package constants

const (
    SDK_VERSION = "1.0"
    SDK_NAME = "GOLANG SDK "
    PAY_ONLINE_2P = "PayOnLine_2p"
    AUTHORIZE_2P = "Authorize_2p"
    QUERY_TXS = "QueryTxs"
    SIMPLE_QUERY_TX = "SimpleQueryTx"
    REFUND = "Refund"
    CAPTURE = "Capture"
    AUTHORIZE_3P = "Authorize_3p"
    BANK_PAYMENT_3P = "BankPayment_3p"
    CASH_PAYMENT_3P = "CashPayment_3p"
    CHANGE_SECRET_KEY = "ChangeSecretKey"
    FRAUD_SCREENING = "FraudScreening"
    NOTIFY_FRAUD_SCREENING_REVIEW = "NotifyFraudScreeningReview"
    PAY_ONLINE_3P = "PayOnLine_3p"
    SPLIT_AUTHORIZE_3P = "SplitAuthorize_3p"
    SPLIT_PAY_ONLINE_3P = "SplitPayOnLine_3p"
    BANK_PAYMENT_2P = "BankPayment_2p"
    SPLIT_PAY_ONLINE_2P = "SplitPayOnLine_2p"
    SPLIT_AUTHORIZE_2P = "SplitAuthorize_2p"
    QUERY_CARD_NUMBER = "QueryCardNumber"
    QUERY_CARD_DETAILS = "QueryCardDetails"
    GET_IIN_DETAILS = "GetIINDetails"
    WSDL_FOLDER = "/NpsSDK/wsdl/"
    STAGING_WSDL_FILE = "staging.wsdl"
    PRODUCTION_WSDL_FILE = "production.wsdl"
    SANDBOX_WSDL_FILE = "sandbox.wsdl"
    DEVELOPMENT_WSDL_FILE = "development.wsdl"
    CREATE_PAYMENT_METHOD = "CreatePaymentMethod"
    CREATE_PAYMENT_METHOD_FROM_PAYMENT = "CreatePaymentMethodFromPayment"
    RETRIEVE_PAYMENT_METHOD = "RetrievePaymentMethod"
    UPDATE_PAYMENT_METHOD = "UpdatePaymentMethod"
    DELETE_PAYMENT_METHOD = "DeletePaymentMethod"
    CREATE_CUSTOMER = "CreateCustomer"
    RETRIEVE_CUSTOMER = "RetrieveCustomer"
    UPDATE_CUSTOMER = "UpdateCustomer"
    DELETE_CUSTOMER = "DeleteCustomer"
    RECACHE_PAYMENT_METHOD_TOKEN =  "RecachePaymentMethodToken"
    CREATE_PAYMENT_METHOD_TOKEN = "CreatePaymentMethodToken"
    RETRIEVE_PAYMENT_METHOD_TOKEN =  "RetrievePaymentMethodToken"
    CREATE_CLIENT_SESSION = "CreateClientSession"
    GET_INSTALLMENTS_OPTIONS = "GetInstallmentsOptions"
    DEBUG = "DEBUG"
    INFO = "INFO"
    ERROR = "ERROR"

    PRODUCTION_ENV = 0
    STAGING_ENV = 1
    SANDBOX_ENV = 2
    DEVELOPMENT_ENV = 3
    LOCAL_ENV = 4

    PRODUCTION_URL = "https://sandbox.nps.com.ar:443/ws.php"
    STAGING_URL = "https://sandbox.nps.com.ar:443/ws.php"
    SANDBOX_URL = "https://sandbox.nps.com.ar:443/ws.php"
    DEVELOPMENT_URL = "https://sandbox.nps.com.ar:443/ws.php"
    LOCAL_URL = "https://psp.nps.com.ar/ws.php"
    DEF_TIMEOUT = 60
)