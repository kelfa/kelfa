package objects

import (
	"net"
	"time"
)

type DataPoint struct {
	ID          string        // Unique ID of the request - es: x-edge-request-id
	DateTime    time.Time     // DateTime of the request - date and time
	ElapsedTime time.Duration // The amount of time the server used to read the request, prepare the response and send it - time-taken

	// Client
	ClientIP        net.IP // Requester IP - c-ip
	ClientUserAgent string // Requester user agent - cs(User-Agent)

	// Request
	RequestDomain   string // The requested domain - es: x-host-header
	RequestURI      string // Requested URI - cs-uri-stem
	RequestUriQuery string // Query string portion of the URI - cs-uri-query
	RequestReferer  string // Referrer, if present - cs(Referer)
	RequestMethod   string // DELETE, GET, HEAD, OPTIONS, PATCH, POST, or PUT - cs-method
	RequestCookie   string // Cookie - cs(Cookie)
	RequestSize     int    // Bytes sent by the client to the server - cs-bytes

	// Communication
	CommunicationProtocol    string // The request protocol like HTTP, HTTPS - cs-protocol
	CommunicationTLSProtocol string // SSL/TLS protocol used - ssl-protocol
	CoomunicationTLSCipher   string // SSL/TLS cipher used - ssl-cipher
	// x-forwarded-for

	// Response
	ResponseSize int // Bytes sent by the server to the client - sc-bytes
	ResponseCode int // HTTP response code - sc-status

	// CDN
	CDNLocation string // Location of the CDN that handled the request - es: x-edge-location
	CDNHost     string // Host of the CDN that handled the request - es: cs(Host)
	CDNCache    string // Hit, Miss, Error, ... - es: x-edge-result-type
}
