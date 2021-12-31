package utils

import (
	"regexp"
)

func VerifyAccount(account string) bool {
	if VerifyIsPhoneNumber(account) {
		return true
	} else if VerifyIsMail(account) {
		return true
	} else {
		return false
	}
}

func VerifyIsPasswd(passwd string) bool {
	if len(passwd) < 8 {
		return false
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, passwd); !b || err != nil {
		return false
	}
	if b, err := regexp.MatchString(a_z, passwd); !b || err != nil {
		return false
	}
	if b, err := regexp.MatchString(A_Z, passwd); !b || err != nil {
		return false
	}
	if b, err := regexp.MatchString(symbol, passwd); !b || err != nil {
		return false
	}
	return true
}

func VerifyIsIpaddress(ip string) bool {
	var ok bool
	ok, _ = regexp.MatchString("^((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}$", ip)
	return ok
}

func VerifyIsPhoneNumber(phone string) bool {
	var ok bool
	ok, _ = regexp.MatchString("^400[0-9]{7}|^1[3456789]\\d{9}$|^0[0-9]{2,3}-[0-9]{8}", phone)
	return ok
}

func VerifyIsMail(mail string) bool {
	var ok bool
	ok, _ = regexp.MatchString("[A-Za-z0-9\u4e00-\u9fa5\\-_.]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$", mail)
	return ok
}

func VerifyIsTaxID(taxID string) bool {
	var ok bool
	ok, _ = regexp.MatchString("^[A-Z0-9]{15}$|^[A-Z0-9]{18}$|^[A-Z0-9]{20}$", taxID)
	return ok
}

func VerifyIsCompany(name string) bool {
	var ok bool
	ok, _ = regexp.MatchString("^[0-9\u4e00-\u9fa5\\(\\)（）a-zA-Z&]{2,50}$", name)
	return ok
}

func VerifyIsBankNo(no string) bool {
	var ok bool
	ok, _ = regexp.MatchString("^([1-9]{1})(\\d{14}|\\d{18})$", no)
	return ok
}

func VerifyIsZipCode(code string) bool {
	var ok bool
	ok, _ = regexp.MatchString("^[1-9][0-9]{5}$", code)
	return ok
}
