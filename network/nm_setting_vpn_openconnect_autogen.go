// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingVpnOpenconnectKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_CACERT:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_PROXY:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE:
		t = ktypeBoolean
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID:
		t = ktypeBoolean
	case NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_MTU:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING:
		t = ktypeString
	}
	return
}

// Check is key in current setting field
func isKeyInSettingVpnOpenconnect(key string) bool {
	switch key {
	case NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_CACERT:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_PROXY:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_MTU:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING:
		return true
	}
	return false
}

// Get key's default value
func getSettingVpnOpenconnectDefaultValue(key string) (value interface{}) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY:
		value = ""
	case NM_SETTING_VPN_OPENCONNECT_KEY_CACERT:
		value = ""
	case NM_SETTING_VPN_OPENCONNECT_KEY_PROXY:
		value = ""
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE:
		value = false
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER:
		value = ""
	case NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT:
		value = ""
	case NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY:
		value = ""
	case NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID:
		value = false
	case NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE:
		value = ""
	case NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT:
		value = ""
	case NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE:
		value = ""
	case NM_SETTING_VPN_OPENCONNECT_KEY_MTU:
		value = ""
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE:
		value = ""
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING:
		value = ""
	}
	return
}

// Get JSON value generally
func generalGetSettingVpnOpenconnectKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingVpnOpenconnectKeyJSON: invalide key", key)
	case NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY:
		value = getSettingVpnOpenconnectKeyGatewayJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_CACERT:
		value = getSettingVpnOpenconnectKeyCacertJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_PROXY:
		value = getSettingVpnOpenconnectKeyProxyJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE:
		value = getSettingVpnOpenconnectKeyCsdEnableJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER:
		value = getSettingVpnOpenconnectKeyCsdWrapperJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT:
		value = getSettingVpnOpenconnectKeyUsercertJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY:
		value = getSettingVpnOpenconnectKeyPrivkeyJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID:
		value = getSettingVpnOpenconnectKeyPemPassphraseFsidJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE:
		value = getSettingVpnOpenconnectKeyCookieJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT:
		value = getSettingVpnOpenconnectKeyGwcertJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE:
		value = getSettingVpnOpenconnectKeyAuthtypeJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_MTU:
		value = getSettingVpnOpenconnectKeyMtuJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE:
		value = getSettingVpnOpenconnectKeyStokenSourceJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING:
		value = getSettingVpnOpenconnectKeyStokenStringJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingVpnOpenconnectKeyJSON(data connectionData, key, valueJSON string) (err error) {
	switch key {
	default:
		logger.Error("generalSetSettingVpnOpenconnectKeyJSON: invalide key", key)
	case NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY:
		err = setSettingVpnOpenconnectKeyGatewayJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_CACERT:
		err = setSettingVpnOpenconnectKeyCacertJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_PROXY:
		err = setSettingVpnOpenconnectKeyProxyJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE:
		err = setSettingVpnOpenconnectKeyCsdEnableJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER:
		err = setSettingVpnOpenconnectKeyCsdWrapperJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT:
		err = setSettingVpnOpenconnectKeyUsercertJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY:
		err = setSettingVpnOpenconnectKeyPrivkeyJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID:
		err = setSettingVpnOpenconnectKeyPemPassphraseFsidJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE:
		err = setSettingVpnOpenconnectKeyCookieJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT:
		err = setSettingVpnOpenconnectKeyGwcertJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE:
		err = setSettingVpnOpenconnectKeyAuthtypeJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_MTU:
		err = setSettingVpnOpenconnectKeyMtuJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE:
		err = setSettingVpnOpenconnectKeyStokenSourceJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING:
		err = setSettingVpnOpenconnectKeyStokenStringJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingVpnOpenconnectKeyGatewayExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY)
}
func isSettingVpnOpenconnectKeyCacertExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CACERT)
}
func isSettingVpnOpenconnectKeyProxyExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PROXY)
}
func isSettingVpnOpenconnectKeyCsdEnableExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE)
}
func isSettingVpnOpenconnectKeyCsdWrapperExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER)
}
func isSettingVpnOpenconnectKeyUsercertExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT)
}
func isSettingVpnOpenconnectKeyPrivkeyExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY)
}
func isSettingVpnOpenconnectKeyPemPassphraseFsidExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID)
}
func isSettingVpnOpenconnectKeyCookieExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE)
}
func isSettingVpnOpenconnectKeyGwcertExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT)
}
func isSettingVpnOpenconnectKeyAuthtypeExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE)
}
func isSettingVpnOpenconnectKeyMtuExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_MTU)
}
func isSettingVpnOpenconnectKeyStokenSourceExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE)
}
func isSettingVpnOpenconnectKeyStokenStringExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING)
}

// Ensure field and key exists and not empty
func ensureFieldSettingVpnOpenconnectExists(data connectionData, errs fieldErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME))
	}
}
func ensureSettingVpnOpenconnectKeyGatewayNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnOpenconnectKeyGatewayExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyGateway(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyCacertNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnOpenconnectKeyCacertExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CACERT, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyCacert(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CACERT, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyProxyNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnOpenconnectKeyProxyExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PROXY, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyProxy(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PROXY, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyCsdEnableNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnOpenconnectKeyCsdEnableExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyCsdWrapperNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnOpenconnectKeyCsdWrapperExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyCsdWrapper(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyUsercertNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnOpenconnectKeyUsercertExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyUsercert(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyPrivkeyNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnOpenconnectKeyPrivkeyExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyPrivkey(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyPemPassphraseFsidNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnOpenconnectKeyPemPassphraseFsidExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyCookieNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnOpenconnectKeyCookieExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyCookie(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyGwcertNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnOpenconnectKeyGwcertExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyGwcert(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyAuthtypeNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnOpenconnectKeyAuthtypeExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyAuthtype(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyMtuNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnOpenconnectKeyMtuExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_MTU, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyMtu(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_MTU, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyStokenSourceNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnOpenconnectKeyStokenSourceExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyStokenSource(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyStokenStringNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnOpenconnectKeyStokenStringExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyStokenString(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING, NM_KEY_ERROR_EMPTY_VALUE)
	}
}

// Getter
func getSettingVpnOpenconnectKeyGateway(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY)
	value, ok := ivalue.(string)
	if !ok {
		logger.Errorf("getSettingVpnOpenconnectKeyGateway: value type is invalid, should be string instead of %#v", ivalue)
	}
	return
}
func getSettingVpnOpenconnectKeyCacert(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CACERT)
	value, ok := ivalue.(string)
	if !ok {
		logger.Errorf("getSettingVpnOpenconnectKeyCacert: value type is invalid, should be string instead of %#v", ivalue)
	}
	return
}
func getSettingVpnOpenconnectKeyProxy(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PROXY)
	value, ok := ivalue.(string)
	if !ok {
		logger.Errorf("getSettingVpnOpenconnectKeyProxy: value type is invalid, should be string instead of %#v", ivalue)
	}
	return
}
func getSettingVpnOpenconnectKeyCsdEnable(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE)
	value, ok := ivalue.(bool)
	if !ok {
		logger.Errorf("getSettingVpnOpenconnectKeyCsdEnable: value type is invalid, should be bool instead of %#v", ivalue)
	}
	return
}
func getSettingVpnOpenconnectKeyCsdWrapper(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER)
	value, ok := ivalue.(string)
	if !ok {
		logger.Errorf("getSettingVpnOpenconnectKeyCsdWrapper: value type is invalid, should be string instead of %#v", ivalue)
	}
	return
}
func getSettingVpnOpenconnectKeyUsercert(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT)
	value, ok := ivalue.(string)
	if !ok {
		logger.Errorf("getSettingVpnOpenconnectKeyUsercert: value type is invalid, should be string instead of %#v", ivalue)
	}
	return
}
func getSettingVpnOpenconnectKeyPrivkey(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY)
	value, ok := ivalue.(string)
	if !ok {
		logger.Errorf("getSettingVpnOpenconnectKeyPrivkey: value type is invalid, should be string instead of %#v", ivalue)
	}
	return
}
func getSettingVpnOpenconnectKeyPemPassphraseFsid(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID)
	value, ok := ivalue.(bool)
	if !ok {
		logger.Errorf("getSettingVpnOpenconnectKeyPemPassphraseFsid: value type is invalid, should be bool instead of %#v", ivalue)
	}
	return
}
func getSettingVpnOpenconnectKeyCookie(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE)
	value, ok := ivalue.(string)
	if !ok {
		logger.Errorf("getSettingVpnOpenconnectKeyCookie: value type is invalid, should be string instead of %#v", ivalue)
	}
	return
}
func getSettingVpnOpenconnectKeyGwcert(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT)
	value, ok := ivalue.(string)
	if !ok {
		logger.Errorf("getSettingVpnOpenconnectKeyGwcert: value type is invalid, should be string instead of %#v", ivalue)
	}
	return
}
func getSettingVpnOpenconnectKeyAuthtype(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE)
	value, ok := ivalue.(string)
	if !ok {
		logger.Errorf("getSettingVpnOpenconnectKeyAuthtype: value type is invalid, should be string instead of %#v", ivalue)
	}
	return
}
func getSettingVpnOpenconnectKeyMtu(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_MTU)
	value, ok := ivalue.(string)
	if !ok {
		logger.Errorf("getSettingVpnOpenconnectKeyMtu: value type is invalid, should be string instead of %#v", ivalue)
	}
	return
}
func getSettingVpnOpenconnectKeyStokenSource(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE)
	value, ok := ivalue.(string)
	if !ok {
		logger.Errorf("getSettingVpnOpenconnectKeyStokenSource: value type is invalid, should be string instead of %#v", ivalue)
	}
	return
}
func getSettingVpnOpenconnectKeyStokenString(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING)
	value, ok := ivalue.(string)
	if !ok {
		logger.Errorf("getSettingVpnOpenconnectKeyStokenString: value type is invalid, should be string instead of %#v", ivalue)
	}
	return
}

// Setter
func setSettingVpnOpenconnectKeyGateway(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY, value)
}
func setSettingVpnOpenconnectKeyCacert(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CACERT, value)
}
func setSettingVpnOpenconnectKeyProxy(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PROXY, value)
}
func setSettingVpnOpenconnectKeyCsdEnable(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE, value)
}
func setSettingVpnOpenconnectKeyCsdWrapper(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER, value)
}
func setSettingVpnOpenconnectKeyUsercert(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT, value)
}
func setSettingVpnOpenconnectKeyPrivkey(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY, value)
}
func setSettingVpnOpenconnectKeyPemPassphraseFsid(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID, value)
}
func setSettingVpnOpenconnectKeyCookie(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE, value)
}
func setSettingVpnOpenconnectKeyGwcert(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT, value)
}
func setSettingVpnOpenconnectKeyAuthtype(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE, value)
}
func setSettingVpnOpenconnectKeyMtu(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_MTU, value)
}
func setSettingVpnOpenconnectKeyStokenSource(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE, value)
}
func setSettingVpnOpenconnectKeyStokenString(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING, value)
}

// JSON Getter
func getSettingVpnOpenconnectKeyGatewayJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY))
	return
}
func getSettingVpnOpenconnectKeyCacertJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CACERT, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_CACERT))
	return
}
func getSettingVpnOpenconnectKeyProxyJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PROXY, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_PROXY))
	return
}
func getSettingVpnOpenconnectKeyCsdEnableJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE))
	return
}
func getSettingVpnOpenconnectKeyCsdWrapperJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER))
	return
}
func getSettingVpnOpenconnectKeyUsercertJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT))
	return
}
func getSettingVpnOpenconnectKeyPrivkeyJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY))
	return
}
func getSettingVpnOpenconnectKeyPemPassphraseFsidJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID))
	return
}
func getSettingVpnOpenconnectKeyCookieJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE))
	return
}
func getSettingVpnOpenconnectKeyGwcertJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT))
	return
}
func getSettingVpnOpenconnectKeyAuthtypeJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE))
	return
}
func getSettingVpnOpenconnectKeyMtuJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_MTU, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_MTU))
	return
}
func getSettingVpnOpenconnectKeyStokenSourceJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE))
	return
}
func getSettingVpnOpenconnectKeyStokenStringJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING))
	return
}

// JSON Setter
func setSettingVpnOpenconnectKeyGatewayJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY))
}
func setSettingVpnOpenconnectKeyCacertJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CACERT, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_CACERT))
}
func setSettingVpnOpenconnectKeyProxyJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PROXY, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_PROXY))
}
func setSettingVpnOpenconnectKeyCsdEnableJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE))
}
func setSettingVpnOpenconnectKeyCsdWrapperJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER))
}
func setSettingVpnOpenconnectKeyUsercertJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT))
}
func setSettingVpnOpenconnectKeyPrivkeyJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY))
}
func setSettingVpnOpenconnectKeyPemPassphraseFsidJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID))
}
func setSettingVpnOpenconnectKeyCookieJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE))
}
func setSettingVpnOpenconnectKeyGwcertJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT))
}
func setSettingVpnOpenconnectKeyAuthtypeJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE))
}
func setSettingVpnOpenconnectKeyMtuJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_MTU, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_MTU))
}
func setSettingVpnOpenconnectKeyStokenSourceJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE))
}
func setSettingVpnOpenconnectKeyStokenStringJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING))
}

// Logic JSON Setter

// Remover
func removeSettingVpnOpenconnectKeyGateway(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY)
}
func removeSettingVpnOpenconnectKeyCacert(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CACERT)
}
func removeSettingVpnOpenconnectKeyProxy(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PROXY)
}
func removeSettingVpnOpenconnectKeyCsdEnable(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE)
}
func removeSettingVpnOpenconnectKeyCsdWrapper(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER)
}
func removeSettingVpnOpenconnectKeyUsercert(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT)
}
func removeSettingVpnOpenconnectKeyPrivkey(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY)
}
func removeSettingVpnOpenconnectKeyPemPassphraseFsid(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID)
}
func removeSettingVpnOpenconnectKeyCookie(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE)
}
func removeSettingVpnOpenconnectKeyGwcert(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT)
}
func removeSettingVpnOpenconnectKeyAuthtype(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE)
}
func removeSettingVpnOpenconnectKeyMtu(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_MTU)
}
func removeSettingVpnOpenconnectKeyStokenSource(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE)
}
func removeSettingVpnOpenconnectKeyStokenString(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING)
}
