package public

import (
	"fmt"

	"github.com/tonyhhyip/vodka"
)

type securityMiddleware struct {
	// STSSeconds is the max-age of the Strict-Transport-Security header.
	// Default is 0, which would NOT include the header.
	STSSeconds int64
	// If STSIncludeSubdomains is set to true, the `includeSubdomains` will
	// be appended to the Strict-Transport-Security header. Default is false.
	STSIncludeSubdomains bool
	// If FrameDeny is set to true, adds the X-Frame-Options header with
	// the value of `DENY`. Default is false.
	FrameDeny bool
	// If ContentTypeNosniff is true, adds the X-Content-Type-Options header
	// with the value `nosniff`. Default is false.
	ContentTypeNosniff bool
	// If BrowserXssFilter is true, adds the X-XSS-Protection header with
	// the value `1; mode=block`. Default is false.
	BrowserXssFilter bool
	// ContentSecurityPolicy allows the Content-Security-Policy header value
	// to be set with a custom value. Default is "".
	ContentSecurityPolicy string
	// HTTP header "Referrer-Policy" governs which referrer information, sent in the Referrer header, should be included with requests made.
	ReferrerPolicy string
	// When true, the whole secury policy applied by the middleware is disable
	// completely.
	IsDevelopment bool
	// Prevent Internet Explorer from executing downloads in your site’s context
	IENoOpen bool
}

func (sm *securityMiddleware) Wrap(next vodka.Handler) vodka.Handler {
	return vodka.HandlerFunc(func(c *vodka.Context) {
		sm.handle(c)
		next.Handle(c)
	})
}

func (sm *securityMiddleware) handle(c *vodka.Context) {
	header := c.Response.Header()

	// Frame Options header.
	if sm.FrameDeny {
		header.Set("X-Frame-Options", "DENY")
	}

	// Content Type Options header.
	if sm.ContentTypeNosniff {
		header.Set("X-Content-Type-Options", "nosniff")
	}

	// XSS Protection header.
	if sm.BrowserXssFilter {
		header.Set("X-Xss-Protection", "1; mode=block")
	}

	// Content Security Policy header.
	if len(sm.ContentSecurityPolicy) > 0 {
		header.Set("Content-Security-Policy", sm.ContentSecurityPolicy)
	}

	// Strict Transport Security header.
	if sm.STSSeconds != 0 {
		header.Set("Strict-Transport-Security", fmt.Sprintf("max-age=%d", sm.STSSeconds))
	}

	// X-Download-Options header.
	if sm.IENoOpen {
		header.Set("X-Download-Options", "noopen")
	}
}
