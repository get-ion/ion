Basic Built'n Handlers
------------

| Middleware | Example |
| -----------|-------------|
| [Basic Authentication](basicauth) | https://github.com/get-ion/ion/tree/master/_examples#authentication |
| [Localization and Internationalization](i18n) | https://github.com/get-ion/ion/tree/master/_examples#miscellaneous |
| [Request Logger](logger) | https://github.com/get-ion/ion/tree/master/_examples#miscellaneous |
| [Profiling (pprof)](pprof) | https://github.com/get-ion/ion/tree/master/_examples#miscellaneous |
| [Recovery](recover) | https://github.com/get-ion/ion/tree/master/_examples#miscellaneous |


Third-Party Handlers
------------

ion has its own middleware form of `func(ctx context.Context)` but it's also compatible with all `net/http` middleware forms. See [here](https://github.com/get-ion/ion/tree/master/_examples/convert-handlers).

Here's a small list of third-party handlers:

> handlers that are ported to work with _ion_'s handlers contains an _Example_ link at the _Description_ column.

| Middleware | Author | Description |
| -----------|--------|-------------|
| [jwt](https://github.com/get-ion/middleware/tree/master/jwt) | [Auth0](https://github.com/auth0) | Middleware checks for a JWT on the `Authorization` header on incoming requests and decodes it. [Example](https://github.com/get-ion/middleware/tree/master/jwt/_example) |
| [cors](https://github.com/get-ion/middleware/tree/master/cors) | [rs](https://github.com/rs) | HTTP Access Control. [Example](https://github.com/get-ion/middleware/tree/master/cors/_example) |
| [secure](https://github.com/get-ion/middleware/tree/master/secure) | [Cory Jacobsen](https://github.com/unrolled) | Middleware that implements a few quick security wins. [Example](https://github.com/get-ion/middleware/tree/master/secure/_example/main.go) |
| [tollbooth](https://github.com/get-ion/middleware/tree/master/tollboothic) | [Didip Kerabat](https://github.com/didip) | Generic middleware to rate-limit HTTP requests. [Example](https://github.com/get-ion/middleware/tree/master/tollbooth/_examples/limit-handler)|
| [cloudwatch](https://github.com/get-ion/middleware/tree/master/cloudwatch) | [Colin Steele](https://github.com/cvillecsteele) | AWS cloudwatch metrics middleware. [Example](https://github.com/get-ion/middleware/tree/master/cloudwatch/_example) |
| [goth](https://github.com/markbates/goth) | [Mark Bates](https://github.com/markbates) | OAuth, OAuth2 authentication. [Example](https://github.com/get-ion/ion/tree/master/_examples/authentication/oauth2) |
| [binding](https://github.com/mholt/binding) | [Matt Holt](https://github.com/mholt) | Data binding from HTTP requests into structs |
| [csp](https://github.com/awakenetworks/csp) | [Awake Networks](https://github.com/awakenetworks) | [Content Security Policy](https://www.w3.org/TR/CSP2/) (CSP) support |
| [delay](https://github.com/jeffbmartinez/delay) | [Jeff Martinez](https://github.com/jeffbmartinez) | Add delays/latency to endpoints. Useful when testing effects of high latency |
| [New Relic Go Agent](https://github.com/yadvendar/negroni-newrelic-go-agent) | [Yadvendar Champawat](https://github.com/yadvendar) | Official [New Relic Go Agent](https://github.com/newrelic/go-agent)  |
| [gorelic](https://github.com/jingweno/negroni-gorelic) | [Jingwen Owen Ou](https://github.com/jingweno) | New Relic agent for Go runtime |
| [onthefly](https://github.com/xyproto/onthefly) | [Alexander Rødseth](https://github.com/xyproto) | Generate TinySVG, HTML and CSS on the fly |
| [permissions2](https://github.com/xyproto/permissions2) | [Alexander Rødseth](https://github.com/xyproto) | Cookies, users and permissions |
| [prometheus](https://github.com/zbindenren/negroni-prometheus) | [Rene Zbinden](https://github.com/zbindenren) | Easily create metrics endpoint for the [prometheus](http://prometheus.io) instrumentation tool |
| [RestGate](https://github.com/pjebs/restgate) | [Prasanga Siripala](https://github.com/pjebs) | Secure authentication for REST API endpoints |
| [stats](https://github.com/thoas/stats) | [Florent Messa](https://github.com/thoas) | Store information about your web application (response time, etc.) |
| [VanGoH](https://github.com/auroratechnologies/vangoh) | [Taylor Wrobel](https://github.com/twrobel3) | Configurable [AWS-Style](http://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html) HMAC authentication middleware |
| [xrequestid](https://github.com/pilu/xrequestid) | [Andrea Franz](https://github.com/pilu) | Middleware that assigns a random X-Request-Id header to each request |
| [digits](https://github.com/bamarni/digits) | [Bilal Amarni](https://github.com/bamarni) | Middleware that handles [Twitter Digits](https://get.digits.com/) authentication |


Feel free to put up your own middleware in this list!