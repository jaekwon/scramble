package main

import (
	"log"
	"strings"
)

type EmailAddress struct {
	Name string
	Host string
}

func (addr *EmailAddress) String() string {
	return addr.Name+"@"+addr.Host
}

// Return this address but with everything after '#' stripped from Name.
func (addr *EmailAddress) StringNoHash() string {
	if hashIdx := strings.Index(addr.Name, "#"); hashIdx != -1 {
		return addr.Name[0:hashIdx]+"@"+addr.Host
	}
	return addr.Name+"@"+addr.Host
}

// Splits the Name into what precedes the first '#' and what follows
func (addr *EmailAddress) NameAndHash() (string, string) {
	if hashIdx := strings.Index(addr.Name, "#"); hashIdx != -1 {
		return addr.Name[:hashIdx], addr.Name[hashIdx+1:]
	}
	return addr.Name, ""
}

// "foo@bar.com" -> EmailAddress
func ParseEmailAddress(addr string) EmailAddress {
	parsed, ok := ParseEmailAddressSafe(addr)
	if !ok {
		log.Panicf("Invalid email address %s", addr)
	}
	return parsed
}

func ParseEmailAddressSafe(addr string) (EmailAddress, bool) {
	match := regexAddress.FindStringSubmatch(addr)
	if match == nil {
		return EmailAddress{}, false
	}
	return EmailAddress{match[1], match[2]}, true
}

// "foo@bar.com,baz@boo.com" -> []EmailAddress
func ParseEmailAddresses(addrList string) EmailAddresses {
	if addrList == "" {
		return nil
	}
	addrParts := strings.Split(addrList, ",")
	addrs := make([]EmailAddress, 0)
	for _, addrPart := range addrParts {
		addrs = append(addrs, ParseEmailAddress(addrPart))
	}
	return addrs
}

// "<foo@bar.com>,<baz@boo.com>" -> []EmailAddress
func ParseAngledEmailAddresses(addrList string, delim string) EmailAddresses {
	if addrList == "" {
		return nil
	}
	addrParts := strings.Split(addrList, delim)
	addrs := make([]EmailAddress, 0)
	for _, addrPart := range addrParts {
		if addrPart[0:1] != "<" || addrPart[len(addrPart)-1:] != ">" {
			log.Panicf("Invalid angled email address %s", addrPart)
		}
		addrPart = addrPart[1:len(addrPart)-1]
		addrs = append(addrs, ParseEmailAddress(addrPart))
	}
	return addrs
}

// "foo@bar.com,baz@boo.com" -> {<host>:[]EmailAddress}
func GroupAddrsByHost(addrList string) map[string]EmailAddresses {
	if addrList == "" {
		return nil
	}
	addrs := strings.Split(addrList, ",")
	hostAddrs := map[string]EmailAddresses{}
	for _, addr := range addrs {
		email := ParseEmailAddress(addr)
		hostAddrs[email.Host] = append(hostAddrs[email.Host], email)
	}
	return hostAddrs
}

// Like GroupAddrsByHost, but resolves the hostname to Mx host.
// Returns two maps where the second is all the hosts for which the Mx lookup failed.
func GroupAddrsByMxHost(addrList string) (map[string]EmailAddresses, map[string]EmailAddresses) {
	hostAddrs := GroupAddrsByHost(addrList)
	mxHostAddrs := map[string]EmailAddresses{}
	failedHostAddrs := map[string]EmailAddresses{}
	for host, addrs := range hostAddrs {
		var mxHost string
		// Skip lookup for self
		// This helps with localhost testing
		if host == GetConfig().SmtpMxHost {
			mxHostAddrs[host] = append(mxHostAddrs[host], addrs...)
			continue
		}
		// Lookup Mx record
		mxHost, err := smtpLookUp(host)
		if err != nil {
			failedHostAddrs[host] = addrs
		} else {
			mxHostAddrs[mxHost] = append(mxHostAddrs[mxHost], addrs...)
		}
	}
	return mxHostAddrs, failedHostAddrs
}

// This lets us add convenience methods to []EmailAddress
type EmailAddresses []EmailAddress

// -> "foo@bar.com,baz@boo.com"
func (addrs EmailAddresses) String() string {
	return strings.Join(addrs.Strings(), ",")
}

// -> "<foo@bar.com>,<baz@boo.com>"
func (addrs EmailAddresses) AngledString(delim string) string {
	if len(addrs) == 0 {
		return ""
	}
	return "<"+strings.Join(addrs.Strings(), ">"+delim+"<")+">"
}

// Like AngledString(), but drops the leftmost items such that
// the result is less than or equal to `limit` bytes.
// This is for storing email>ancestor_ids, for the References header.
func (addrs EmailAddresses) AngledStringCappedToBytes(delim string, limit int) string {
	if len(addrs) == 0 {
		return ""
	}
	res := []byte("<"+strings.Join(addrs.Strings(), ">"+delim+"<")+">")
	if len(res) <= limit {
		return string(res)
	}
	resPart := string(res[len(res)-limit:])
	addrStart := strings.Index(resPart, "<")
	if addrStart == -1 {
		return ""
	}
	return resPart[addrStart:]
}

// -> ["foo@bar.com","baz@boo.com"]
func (addrs EmailAddresses) Strings() []string {
	addrsList := []string{}
	for _, addr := range addrs {
		addrsList = append(addrsList, addr.String())
	}
	return addrsList
}

// Returns those addresses that have given host
func (addrs EmailAddresses) FilterByHost(host string) EmailAddresses {
	filtered := []EmailAddress{}
	for _, addr := range addrs {
		if addr.Host == host {
			filtered = append(filtered, addr)
		}
	}
	return filtered
}
