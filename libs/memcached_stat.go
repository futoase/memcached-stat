// reference from https://github.com/youtube/vitess/blob/master/go/memcache/memcache.go
// LICENSE is BSD-Style license.

package memcached_stat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Connection struct {
	conn     net.Conn
	buffered bufio.ReadWriter
}

type Result struct {
	Key   string
	Value []byte
	Flags uint16
	Cas   uint64
}

func Connect(address string) (conn *Connection, err error) {
	var network string
	if strings.Contains(address, "/") {
		network = "unix"
	} else {
		network = "tcp"
	}
	var nc net.Conn
	nc, err = net.Dial(network, address)
  if err != nil {
    return nil, err
  }
	return newConnection(nc), nil
}

func newConnection(nc net.Conn) *Connection {
	return &Connection{
		conn: nc,
		buffered: bufio.ReadWriter{
			Reader: bufio.NewReader(nc),
			Writer: bufio.NewWriter(nc),
		},
	}
}

func (mc *Connection) Close() {
	mc.conn.Close()
	mc.conn = nil
}

func (mc *Connection) Stats(argument string) (result []byte, err error) {
	defer handleError(&err)
	if argument == "" {
		mc.writestrings("stats\r\n")
	} else {
		mc.writestrings("stats ", argument, "\r\n")
	}
	mc.flush()
	for {
		l := mc.readline()
		if strings.HasPrefix(l, "END") {
			break
		}
		if strings.Contains(l, "ERROR") {
			return nil, NewMemcacheError(l)
		}
		result = append(result, l...)
		result = append(result, '\n')
	}
	return result, err
}

func (mc *Connection) writestrings(strs ...string) {
	for _, s := range strs {
		mc.writestring(s)
	}
}

func (mc *Connection) writestring(s string) {
	if _, err := mc.buffered.WriteString(s); err != nil {
		panic(NewMemcacheError("%s", err))
	}
}

func (mc *Connection) flush() {
	if err := mc.buffered.Flush(); err != nil {
		panic(NewMemcacheError("%s", err))
	}
}

func (mc *Connection) readline() string {
	mc.flush()
	l, isPrefix, err := mc.buffered.ReadLine()
	if isPrefix || err != nil {
		panic(NewMemcacheError("prefix: %v, %s", isPrefix, err))
	}
	return string(l)
}

func NewMemcacheError(format string, args ...interface{}) MemcacheError {
	return MemcacheError{fmt.Sprintf(format, args...)}
}

type MemcacheError struct {
	Message string
}

func (merr MemcacheError) Error() string {
	return merr.Message
}

func handleError(err *error) {
	if x := recover(); x != nil {
		*err = x.(MemcacheError)
	}
}
