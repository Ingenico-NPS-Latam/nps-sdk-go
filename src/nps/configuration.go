package nps

import (
	"errors"
	"log"
	CONSTANTS "nps/constants"
	"os"
)

type MyConfiguration struct {
	environment        int
	secret_key         string
	connection_Timeout int
	execution_Timeout  int
	debug              bool
	verifPeer          bool
	logLevel           string
	sanitize           bool
	url                string
	log_file           string
	npsLog             *log.Logger
	proxy_url          string
	proxy_username     string
	proxy_password     string
}

var Configuration = NewMyConfiguration()

func NewMyConfiguration() *MyConfiguration {
	c := new(MyConfiguration)
	c.environment = CONSTANTS.SANDBOX_ENV
	c.url = CONSTANTS.SANDBOX_URL
	c.secret_key = ""
	c.debug = false
	c.execution_Timeout = 60
	c.verifPeer = true
	c.logLevel = CONSTANTS.INFO
	c.sanitize = true
	c.log_file = ""
	c.npsLog = nil
	c.proxy_url = ""
	c.proxy_username = ""
	c.proxy_password = ""

	return c
}

//Configure Sdk
//environment = CONSTANTS.PRODUCTION_ENV|CONSTANTS.STAGING_ENV|CONSTANTS.SANDBOX_ENV|CONSTANTS.DEVELOPMENT_ENV
//secret_key = string
//debug=bool (default false)
//conn_timeout (default 10)
//timeout=int (default 60)
//cert_verify_peer= bool (default true)
//log_level= [CONSTANTS.DEBUG|CONSTANTS.INFO|CONSTANTS.ERROR] default(CONSTANTS.INFO)
//sanitize= bool (default true)
//log_file= string
//npsLog= log.logger (default nil)
//proxy_url = string
//proxy_username = string
//proxy_password = string
func Configure(m map[string]interface{}) error {

	for key, value := range m {
		if key == "environment" {
			Configuration.environment = value.(int)
			switch Configuration.environment {
			case CONSTANTS.PRODUCTION_ENV:
				Configuration.url = CONSTANTS.PRODUCTION_URL
			case CONSTANTS.STAGING_ENV:
				Configuration.url = CONSTANTS.STAGING_URL
			case CONSTANTS.SANDBOX_ENV:
				Configuration.url = CONSTANTS.SANDBOX_URL
			case CONSTANTS.DEVELOPMENT_ENV:
				Configuration.url = CONSTANTS.DEVELOPMENT_URL
			case CONSTANTS.LOCAL_ENV:
				Configuration.url = CONSTANTS.LOCAL_URL
			default:
				errMsg := "environment must be [CONSTANTS.PRODUCTION_ENV|CONSTANTS.STAGING_ENV|CONSTANTS.SANDBOX_ENV|CONSTANTS.DEVELOPMENT_ENV]"
				log.Println(errMsg)
				err := errors.New(errMsg)
				return err
			}
		}
		if key == "secret_key" {
			Configuration.secret_key = value.(string)
		}
		if key == "debug" {
			Configuration.debug = value.(bool)
		}
		if key == "conn_timeout" {
			Configuration.connection_Timeout = value.(int)
		}
		if key == "timeout" {
			Configuration.execution_Timeout = value.(int)
		}
		if key == "cert_verify_peer" {
			Configuration.verifPeer = value.(bool)
		}
		if key == "log_level" {
			Configuration.logLevel = value.(string)
		}
		if key == "sanitize" {
			Configuration.sanitize = value.(bool)
		}
		if key == "log_file" {
			Configuration.log_file = value.(string)
		}
		if key == "npsLog" {
			Configuration.npsLog = value.(*log.Logger)
		}
		if key == "proxy_url" {
			Configuration.proxy_url = value.(string)
		}
		if key == "proxy_username" {
			Configuration.proxy_username = value.(string)
		}
		if key == "proxy_password" {
			Configuration.proxy_password = value.(string)
		}
	}

	if Configuration.npsLog == nil {
		Configuration.npsLog = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	}
	return nil
}

func (c *MyConfiguration) GetSecretKey() string {
	return c.secret_key
}
