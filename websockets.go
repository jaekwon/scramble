package main

import (
	ec "github.com/jaekwon/go-websocket/easyconn"
)

var wsEmail2Conn = map[string][]*ec.Connection{}
var wsConn2Email = map[*ec.Connection]string{}
var wsConnOnClose = make(chan *ec.Connection)

func init() {
	go func() {
		for {
			select {
			case conn := <-wsConnOnClose:
				wsClose(conn)
			}
		}
	}()
}

func WSRegister(user *User, conn *ec.Connection) {
	wsEmail2Conn[user.EmailAddress] = append(wsEmail2Conn[user.EmailAddress], conn)
	wsConn2Email[conn] = user.EmailAddress
}

func WSSendMessage(user *User, message string) {
	conns := wsEmail2Conn[user.EmailAddress]
	bytes := []byte(message)
	for _, conn := range conns {
		conn.Send <- bytes
	}
}

// not threadsafe
func wsClose(conn *ec.Connection) {
	delete(wsConn2Email, conn)
	email := wsConn2Email[conn]
	if email == "" { panic("Expected email for conn, found none") }
	conns := wsEmail2Conn[email]
	if len(conns) == 1 {
		if conns[0] != conn { panic("Expected conn in list for "+email+", found something else") }
		delete(wsEmail2Conn, email)
	}
	filtered := []*ec.Connection{}
	for _, c := range wsEmail2Conn[email] {
		if conn == c { continue }
		filtered = append(filtered, c)
	}
	if len(filtered) != len(wsEmail2Conn[email])-1 {
		panic("Expected conn in list for "+email+", did not find exactly 1")
	}
	wsEmail2Conn[email] = filtered
}
