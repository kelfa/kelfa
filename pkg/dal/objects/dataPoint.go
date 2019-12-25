package objects

import (
	"net"
	"time"
)

type DataPoint struct {
	ID              string        // Unique ID of the request - es: x-edge-request-id
	DateTime        time.Time     // DateTime of the request - date and time
	ElapsedTime     time.Duration // The amount of time the server used to read the request, prepare the response and send it - time-taken
	TimeToFirstByte time.Duration // The number of seconds between receiving the request and writing the first byte of the response, as measured on the server - time-to-first-byte  [from 2019-12-12]

	// Client
	ClientIP        net.IP // Requester IP - c-ip
	ClientUserAgent string // Requester user agent - cs(User-Agent)
	ClientPort      int    // Requester port number - c-port [from 2019-12-12]

	// Request
	RequestDomain   string // The requested domain - es: x-host-header
	RequestURI      string // Requested URI - cs-uri-stem
	RequestURIQuery string // Query string portion of the URI - cs-uri-query
	RequestReferer  string // Referrer, if present - cs(Referer)
	RequestMethod   string // DELETE, GET, HEAD, OPTIONS, PATCH, POST, or PUT - cs-method
	RequestCookie   string // Cookie - cs(Cookie)
	RequestSize     int    // Bytes sent by the client to the server - cs-bytes

	// Communication
	CommunicationProtocol    string // The request protocol like HTTP, HTTPS - cs-protocol
	CommunicationTLSProtocol string // SSL/TLS protocol used - ssl-protocol
	CommunicationTLSCipher   string // SSL/TLS cipher used - ssl-cipher
	// x-forwarded-for

	// Response
	ResponseSize              int    // Bytes sent by the server to the client - sc-bytes
	ResponseCode              int    // HTTP response code - sc-status
	ResponseContentType       string // The value of the HTTP Content-Type header of the response - sc-content-type [from 2019-12-12]
	ResponseContentLen        int    // The value of the HTTP Content-Length header of the response - sc-content-len [from 2019-12-12]
	ResponseContentRangeBegin int    // When the response contains the HTTP Content-Range header, this field contains the range start value - sc-range-start [from 2019-12-12]
	ResponseContentRangeEnd   int    // When the response contains the HTTP Content-Range header, this field contains the range end value - sc-range-end [from 2019-12-12]

	// CDN
	CDNLocation      string // Location of the CDN that handled the request - es: x-edge-location
	CDNHost          string // Host of the CDN that handled the request - es: cs(Host)
	CDNCache         string // Hit, Miss, Error, ... - es: x-edge-result-type
	CDNCacheDetailed string // When the result type is an error, this field contains the specific type of error - x-edge-detailed-result-type [from 2019-12-12]
}
