// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingVpnPptpKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_VPN_PPTP_KEY_GATEWAY:
		t = ktypeString
	case NM_SETTING_VPN_PPTP_KEY_USER:
		t = ktypeString
	case NM_SETTING_VPN_PPTP_KEY_PASSWORD_FLAGS:
		t = ktypeUint32
	case NM_SETTING_VPN_PPTP_KEY_PASSWORD:
		t = ktypeString
	case NM_SETTING_VPN_PPTP_KEY_DOMAIN:
		t = ktypeString
	}
	return
}

// Check is key in current setting field
func isKeyInSettingVpnPptp(key string) bool {
	switch key {
	case NM_SETTING_VPN_PPTP_KEY_GATEWAY:
		return true
	case NM_SETTING_VPN_PPTP_KEY_USER:
		return true
	case NM_SETTING_VPN_PPTP_KEY_PASSWORD_FLAGS:
		return true
	case NM_SETTING_VPN_PPTP_KEY_PASSWORD:
		return true
	case NM_SETTING_VPN_PPTP_KEY_DOMAIN:
		return true
	}
	return false
}

// Get key's default value
func getSettingVpnPptpDefaultValue(key string) (value interface{}) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_VPN_PPTP_KEY_GATEWAY:
		value = ""
	case NM_SETTING_VPN_PPTP_KEY_USER:
		value = ""
	case NM_SETTING_VPN_PPTP_KEY_PASSWORD_FLAGS:
		value = 0
	case NM_SETTING_VPN_PPTP_KEY_PASSWORD:
		value = ""
	case NM_SETTING_VPN_PPTP_KEY_DOMAIN:
		value = ""
	}
	return
}

// Get JSON value generally
func generalGetSettingVpnPptpKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingVpnPptpKeyJSON: invalide key", key)
	case NM_SETTING_VPN_PPTP_KEY_GATEWAY:
		value = getSettingVpnPptpKeyGatewayJSON(data)
	case NM_SETTING_VPN_PPTP_KEY_USER:
		value = getSettingVpnPptpKeyUserJSON(data)
	case NM_SETTING_VPN_PPTP_KEY_PASSWORD_FLAGS:
		value = getSettingVpnPptpKeyPasswordFlagsJSON(data)
	case NM_SETTING_VPN_PPTP_KEY_PASSWORD:
		value = getSettingVpnPptpKeyPasswordJSON(data)
	case NM_SETTING_VPN_PPTP_KEY_DOMAIN:
		value = getSettingVpnPptpKeyDomainJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingVpnPptpKeyJSON(data connectionData, key, valueJSON string) (err error) {
	switch key {
	default:
		logger.Error("generalSetSettingVpnPptpKeyJSON: invalide key", key)
	case NM_SETTING_VPN_PPTP_KEY_GATEWAY:
		err = setSettingVpnPptpKeyGatewayJSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_USER:
		err = setSettingVpnPptpKeyUserJSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_PASSWORD_FLAGS:
		err = setSettingVpnPptpKeyPasswordFlagsJSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_PASSWORD:
		err = setSettingVpnPptpKeyPasswordJSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_DOMAIN:
		err = setSettingVpnPptpKeyDomainJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingVpnPptpKeyGatewayExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_GATEWAY)
}
func isSettingVpnPptpKeyUserExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_USER)
}
func isSettingVpnPptpKeyPasswordFlagsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_PASSWORD_FLAGS)
}
func isSettingVpnPptpKeyPasswordExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_PASSWORD)
}
func isSettingVpnPptpKeyDomainExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_DOMAIN)
}

// Ensure field and key exists and not empty
func ensureFieldSettingVpnPptpExists(data connectionData, errs fieldErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_VF_VPN_PPTP_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_VF_VPN_PPTP_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_VF_VPN_PPTP_SETTING_NAME))
	}
}
func ensureSettingVpnPptpKeyGatewayNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyGatewayExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_GATEWAY, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnPptpKeyGateway(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_GATEWAY, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnPptpKeyUserNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyUserExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_USER, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnPptpKeyUser(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_USER, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnPptpKeyPasswordFlagsNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyPasswordFlagsExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_PASSWORD_FLAGS, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnPptpKeyPasswordNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyPasswordExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_PASSWORD, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnPptpKeyPassword(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_PASSWORD, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnPptpKeyDomainNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyDomainExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_DOMAIN, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnPptpKeyDomain(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_DOMAIN, NM_KEY_ERROR_EMPTY_VALUE)
	}
}

// Getter
func getSettingVpnPptpKeyGateway(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_GATEWAY)
	value, ok := ivalue.(string)
	if !ok {
		logger.Errorf("getSettingVpnPptpKeyGateway: value type is invalid, should be string instead of %#v", ivalue)
	}
	return
}
func getSettingVpnPptpKeyUser(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_USER)
	value, ok := ivalue.(string)
	if !ok {
		logger.Errorf("getSettingVpnPptpKeyUser: value type is invalid, should be string instead of %#v", ivalue)
	}
	return
}
func getSettingVpnPptpKeyPasswordFlags(data connectionData) (value uint32) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_PASSWORD_FLAGS)
	value, ok := ivalue.(uint32)
	if !ok {
		logger.Errorf("getSettingVpnPptpKeyPasswordFlags: value type is invalid, should be uint32 instead of %#v", ivalue)
	}
	return
}
func getSettingVpnPptpKeyPassword(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_PASSWORD)
	value, ok := ivalue.(string)
	if !ok {
		logger.Errorf("getSettingVpnPptpKeyPassword: value type is invalid, should be string instead of %#v", ivalue)
	}
	return
}
func getSettingVpnPptpKeyDomain(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_DOMAIN)
	value, ok := ivalue.(string)
	if !ok {
		logger.Errorf("getSettingVpnPptpKeyDomain: value type is invalid, should be string instead of %#v", ivalue)
	}
	return
}

// Setter
func setSettingVpnPptpKeyGateway(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_GATEWAY, value)
}
func setSettingVpnPptpKeyUser(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_USER, value)
}
func setSettingVpnPptpKeyPasswordFlags(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_PASSWORD_FLAGS, value)
}
func setSettingVpnPptpKeyPassword(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_PASSWORD, value)
}
func setSettingVpnPptpKeyDomain(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_DOMAIN, value)
}

// JSON Getter
func getSettingVpnPptpKeyGatewayJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_GATEWAY, getSettingVpnPptpKeyType(NM_SETTING_VPN_PPTP_KEY_GATEWAY))
	return
}
func getSettingVpnPptpKeyUserJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_USER, getSettingVpnPptpKeyType(NM_SETTING_VPN_PPTP_KEY_USER))
	return
}
func getSettingVpnPptpKeyPasswordFlagsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_PASSWORD_FLAGS, getSettingVpnPptpKeyType(NM_SETTING_VPN_PPTP_KEY_PASSWORD_FLAGS))
	return
}
func getSettingVpnPptpKeyPasswordJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_PASSWORD, getSettingVpnPptpKeyType(NM_SETTING_VPN_PPTP_KEY_PASSWORD))
	return
}
func getSettingVpnPptpKeyDomainJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_DOMAIN, getSettingVpnPptpKeyType(NM_SETTING_VPN_PPTP_KEY_DOMAIN))
	return
}

// JSON Setter
func setSettingVpnPptpKeyGatewayJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_GATEWAY, valueJSON, getSettingVpnPptpKeyType(NM_SETTING_VPN_PPTP_KEY_GATEWAY))
}
func setSettingVpnPptpKeyUserJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_USER, valueJSON, getSettingVpnPptpKeyType(NM_SETTING_VPN_PPTP_KEY_USER))
}
func setSettingVpnPptpKeyPasswordFlagsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_PASSWORD_FLAGS, valueJSON, getSettingVpnPptpKeyType(NM_SETTING_VPN_PPTP_KEY_PASSWORD_FLAGS))
}
func setSettingVpnPptpKeyPasswordJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_PASSWORD, valueJSON, getSettingVpnPptpKeyType(NM_SETTING_VPN_PPTP_KEY_PASSWORD))
}
func setSettingVpnPptpKeyDomainJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_DOMAIN, valueJSON, getSettingVpnPptpKeyType(NM_SETTING_VPN_PPTP_KEY_DOMAIN))
}

// Logic JSON Setter

// Remover
func removeSettingVpnPptpKeyGateway(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_GATEWAY)
}
func removeSettingVpnPptpKeyUser(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_USER)
}
func removeSettingVpnPptpKeyPasswordFlags(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_PASSWORD_FLAGS)
}
func removeSettingVpnPptpKeyPassword(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_PASSWORD)
}
func removeSettingVpnPptpKeyDomain(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_DOMAIN)
}
