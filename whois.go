/*
 * Copyright (c) 2013 Vladimir Blaskov
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice, this
 *    list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
 * ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 * ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package whois

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
)

const (
	WHOIS_DOMAIN = ".whois-servers.net"
	WHOIS_PORT   = "43"
)

func Whois(query string, params map[string]string) (result string, err error) {
	var buffer []byte
	host, port := findHostPort(query, params)

	conn, err := net.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		return
	}

	fmt.Fprintf(conn, "%s\r\n", query)

	buffer, err = ioutil.ReadAll(conn)
	if err != nil {
		return
	}

	result = string(buffer[:])

	return
}

func findHostPort(query string, params map[string]string) (host, port string) {
	var ok bool

	if host, ok = params["host"]; !ok {
		fields := strings.Split(query, ".")
		tld := fields[len(fields)-1]

		host = fmt.Sprint(tld, WHOIS_DOMAIN)
	}

	if port, ok = params["port"]; !ok {
		port = WHOIS_PORT
	}

	return
}
