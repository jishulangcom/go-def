package def

import "net/http"

const (
	HttpStatus100 = http.StatusContinue           // RFC 7231, 6.2.1
	HttpStatus101 = http.StatusSwitchingProtocols // RFC 7231, 6.2.2
	HttpStatus102 = http.StatusProcessing         // RFC 2518, 10.1
	HttpStatus103 = http.StatusEarlyHints         // RFC 8297

	HttpStatus200 = http.StatusOK                   // RFC 7231, 6.3.1
	HttpStatus201 = http.StatusCreated              // RFC 7231, 6.3.2
	HttpStatus202 = http.StatusAccepted             // RFC 7231, 6.3.3
	HttpStatus203 = http.StatusNonAuthoritativeInfo // RFC 7231, 6.3.4
	HttpStatus204 = http.StatusNoContent            // RFC 7231, 6.3.5
	HttpStatus205 = http.StatusResetContent         // RFC 7231, 6.3.6
	HttpStatus206 = http.StatusPartialContent       // RFC 7233, 4.1
	HttpStatus207 = http.StatusMultiStatus          // RFC 4918, 11.1
	HttpStatus208 = http.StatusAlreadyReported      // RFC 5842, 7.1
	HttpStatus226 = http.StatusIMUsed               // RFC 3229, 10.4.1

	HttpStatus300 = http.StatusMultipleChoices   // RFC 7231, 6.4.1
	HttpStatus301 = http.StatusMovedPermanently  // RFC 7231, 6.4.2
	HttpStatus302 = http.StatusFound             // RFC 7231, 6.4.3
	HttpStatus303 = http.StatusSeeOther          // RFC 7231, 6.4.4
	HttpStatus304 = http.StatusNotModified       // RFC 7232, 4.1
	HttpStatus305 = http.StatusUseProxy          // RFC 7231, 6.4.5
	HttpStatus306 = 306                          // RFC 7231, 6.4.6 (Unused)
	HttpStatus307 = http.StatusTemporaryRedirect // RFC 7231, 6.4.7
	HttpStatus308 = http.StatusPermanentRedirect // RFC 7538, 3

	HttpStatus400 = http.StatusBadRequest                   // RFC 7231, 6.5.1
	HttpStatus401 = http.StatusUnauthorized                 // RFC 7235, 3.1
	HttpStatus402 = http.StatusPaymentRequired              // RFC 7231, 6.5.2
	HttpStatus403 = http.StatusForbidden                    // RFC 7231, 6.5.3
	HttpStatus404 = http.StatusNotFound                     // RFC 7231, 6.5.4
	HttpStatus405 = http.StatusMethodNotAllowed             // RFC 7231, 6.5.5
	HttpStatus406 = http.StatusNotAcceptable                // RFC 7231, 6.5.6
	HttpStatus407 = http.StatusProxyAuthRequired            // RFC 7235, 3.2
	HttpStatus408 = http.StatusRequestTimeout               // RFC 7231, 6.5.7
	HttpStatus409 = http.StatusConflict                     // RFC 7231, 6.5.8
	HttpStatus410 = http.StatusGone                         // RFC 7231, 6.5.9
	HttpStatus411 = http.StatusLengthRequired               // RFC 7231, 6.5.10
	HttpStatus412 = http.StatusPreconditionFailed           // RFC 7232, 4.2
	HttpStatus413 = http.StatusRequestEntityTooLarge        // RFC 7231, 6.5.11
	HttpStatus414 = http.StatusRequestURITooLong            // RFC 7231, 6.5.12
	HttpStatus415 = http.StatusUnsupportedMediaType         // RFC 7231, 6.5.13
	HttpStatus416 = http.StatusRequestedRangeNotSatisfiable // RFC 7233, 4.4
	HttpStatus417 = http.StatusExpectationFailed            // RFC 7231, 6.5.14
	HttpStatus418 = http.StatusTeapot                       // RFC 7168, 2.3.3
	HttpStatus421 = http.StatusMisdirectedRequest           // RFC 7540, 9.1.2
	HttpStatus422 = http.StatusUnprocessableEntity          // RFC 4918, 11.2
	HttpStatus423 = http.StatusLocked                       // RFC 4918, 11.3
	HttpStatus424 = http.StatusFailedDependency             // RFC 4918, 11.4
	HttpStatus425 = http.StatusTooEarly                     // RFC 8470, 5.2.
	HttpStatus426 = http.StatusUpgradeRequired              // RFC 7231, 6.5.15
	HttpStatus428 = http.StatusPreconditionRequired         // RFC 6585, 3
	HttpStatus429 = http.StatusTooManyRequests              // RFC 6585, 4
	HttpStatus431 = http.StatusRequestHeaderFieldsTooLarge  // RFC 6585, 5
	HttpStatus451 = http.StatusUnavailableForLegalReasons   // RFC 7725, 3

	HttpStatus500 = http.StatusInternalServerError           // RFC 7231, 6.6.1
	HttpStatus501 = http.StatusNotImplemented                // RFC 7231, 6.6.2
	HttpStatus502 = http.StatusBadGateway                    // RFC 7231, 6.6.3
	HttpStatus503 = http.StatusServiceUnavailable            // RFC 7231, 6.6.4
	HttpStatus504 = http.StatusGatewayTimeout                // RFC 7231, 6.6.5
	HttpStatus505 = http.StatusHTTPVersionNotSupported       // RFC 7231, 6.6.6
	HttpStatus506 = http.StatusVariantAlsoNegotiates         // RFC 2295, 8.1
	HttpStatus507 = http.StatusInsufficientStorage           // RFC 4918, 11.5
	HttpStatus508 = http.StatusLoopDetected                  // RFC 5842, 7.2
	HttpStatus510 = http.StatusNotExtended                   // RFC 2774, 7
	HttpStatus511 = http.StatusNetworkAuthenticationRequired // RFC 6585, 6
)
