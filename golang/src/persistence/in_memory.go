package persistence

import "notchman8600/authentication-provider/domain"

var AuthCodeList = make(map[string]domain.AuthCode)
var TokenCodeList = make(map[string]domain.TokenCode)
var SessionList = make(map[string]domain.Session)
