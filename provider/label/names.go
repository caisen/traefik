package label

// Traefik labels
const (
	Prefix                                         = "traefik."
	SuffixBackend                                  = "backend"
	SuffixDomain                                   = "domain"
	SuffixEnable                                   = "enable"
	SuffixPort                                     = "port"
	SuffixPortIndex                                = "portIndex"
	SuffixProtocol                                 = "protocol"
	SuffixTags                                     = "tags"
	SuffixWeight                                   = "weight"
	SuffixBackendID                                = "backend.id"
	SuffixBackendCircuitBreaker                    = "backend.circuitbreaker"
	SuffixBackendCircuitBreakerExpression          = "backend.circuitbreaker.expression"
	SuffixBackendHealthCheckPath                   = "backend.healthcheck.path"
	SuffixBackendHealthCheckPort                   = "backend.healthcheck.port"
	SuffixBackendHealthCheckInterval               = "backend.healthcheck.interval"
	SuffixBackendLoadBalancer                      = "backend.loadbalancer"
	SuffixBackendLoadBalancerMethod                = SuffixBackendLoadBalancer + ".method"
	SuffixBackendLoadBalancerSticky                = SuffixBackendLoadBalancer + ".sticky"
	SuffixBackendLoadBalancerStickiness            = SuffixBackendLoadBalancer + ".stickiness"
	SuffixBackendLoadBalancerStickinessCookieName  = SuffixBackendLoadBalancer + ".stickiness.cookieName"
	SuffixBackendMaxConnAmount                     = "backend.maxconn.amount"
	SuffixBackendMaxConnExtractorFunc              = "backend.maxconn.extractorfunc"
	SuffixFrontend                                 = "frontend"
	SuffixFrontendAuthBasic                        = "frontend.auth.basic"
	SuffixFrontendBackend                          = "frontend.backend"
	SuffixFrontendEntryPoints                      = "frontend.entryPoints"
	SuffixFrontendHeaders                          = "frontend.headers."
	SuffixFrontendRequestHeaders                   = SuffixFrontendHeaders + "customRequestHeaders"
	SuffixFrontendResponseHeaders                  = SuffixFrontendHeaders + "customResponseHeaders"
	SuffixFrontendHeadersAllowedHosts              = SuffixFrontendHeaders + "allowedHosts"
	SuffixFrontendHeadersHostsProxyHeaders         = SuffixFrontendHeaders + "hostsProxyHeaders"
	SuffixFrontendHeadersSSLRedirect               = SuffixFrontendHeaders + "SSLRedirect"
	SuffixFrontendHeadersSSLTemporaryRedirect      = SuffixFrontendHeaders + "SSLTemporaryRedirect"
	SuffixFrontendHeadersSSLHost                   = SuffixFrontendHeaders + "SSLHost"
	SuffixFrontendHeadersSSLProxyHeaders           = SuffixFrontendHeaders + "SSLProxyHeaders"
	SuffixFrontendHeadersSTSSeconds                = SuffixFrontendHeaders + "STSSeconds"
	SuffixFrontendHeadersSTSIncludeSubdomains      = SuffixFrontendHeaders + "STSIncludeSubdomains"
	SuffixFrontendHeadersSTSPreload                = SuffixFrontendHeaders + "STSPreload"
	SuffixFrontendHeadersForceSTSHeader            = SuffixFrontendHeaders + "forceSTSHeader"
	SuffixFrontendHeadersFrameDeny                 = SuffixFrontendHeaders + "frameDeny"
	SuffixFrontendHeadersCustomFrameOptionsValue   = SuffixFrontendHeaders + "customFrameOptionsValue"
	SuffixFrontendHeadersContentTypeNosniff        = SuffixFrontendHeaders + "contentTypeNosniff"
	SuffixFrontendHeadersBrowserXSSFilter          = SuffixFrontendHeaders + "browserXSSFilter"
	SuffixFrontendHeadersContentSecurityPolicy     = SuffixFrontendHeaders + "contentSecurityPolicy"
	SuffixFrontendHeadersPublicKey                 = SuffixFrontendHeaders + "publicKey"
	SuffixFrontendHeadersReferrerPolicy            = SuffixFrontendHeaders + "referrerPolicy"
	SuffixFrontendHeadersIsDevelopment             = SuffixFrontendHeaders + "isDevelopment"
	SuffixFrontendPassHostHeader                   = "frontend.passHostHeader"
	SuffixFrontendPassTLSCert                      = "frontend.passTLSCert"
	SuffixFrontendPriority                         = "frontend.priority"
	SuffixFrontendRateLimitExtractorFunc           = "frontend.rateLimit.extractorFunc"
	SuffixFrontendRedirectEntryPoint               = "frontend.redirect.entryPoint"
	SuffixFrontendRedirectRegex                    = "frontend.redirect.regex"
	SuffixFrontendRedirectReplacement              = "frontend.redirect.replacement"
	SuffixFrontendRule                             = "frontend.rule"
	SuffixFrontendRuleType                         = "frontend.rule.type"
	SuffixFrontendWhitelistSourceRange             = "frontend.whitelistSourceRange"
	TraefikDomain                                  = Prefix + SuffixDomain
	TraefikEnable                                  = Prefix + SuffixEnable
	TraefikPort                                    = Prefix + SuffixPort
	TraefikPortIndex                               = Prefix + SuffixPortIndex
	TraefikProtocol                                = Prefix + SuffixProtocol
	TraefikTags                                    = Prefix + SuffixTags
	TraefikWeight                                  = Prefix + SuffixWeight
	TraefikBackend                                 = Prefix + SuffixBackend
	TraefikBackendID                               = Prefix + SuffixBackendID
	TraefikBackendCircuitBreaker                   = Prefix + SuffixBackendCircuitBreaker
	TraefikBackendCircuitBreakerExpression         = Prefix + SuffixBackendCircuitBreakerExpression
	TraefikBackendHealthCheckPath                  = Prefix + SuffixBackendHealthCheckPath
	TraefikBackendHealthCheckPort                  = Prefix + SuffixBackendHealthCheckPort
	TraefikBackendHealthCheckInterval              = Prefix + SuffixBackendHealthCheckInterval
	TraefikBackendLoadBalancer                     = Prefix + SuffixBackendLoadBalancer
	TraefikBackendLoadBalancerMethod               = Prefix + SuffixBackendLoadBalancerMethod
	TraefikBackendLoadBalancerSticky               = Prefix + SuffixBackendLoadBalancerSticky
	TraefikBackendLoadBalancerStickiness           = Prefix + SuffixBackendLoadBalancerStickiness
	TraefikBackendLoadBalancerStickinessCookieName = Prefix + SuffixBackendLoadBalancerStickinessCookieName
	TraefikBackendMaxConnAmount                    = Prefix + SuffixBackendMaxConnAmount
	TraefikBackendMaxConnExtractorFunc             = Prefix + SuffixBackendMaxConnExtractorFunc
	TraefikFrontend                                = Prefix + SuffixFrontend
	TraefikFrontendAuthBasic                       = Prefix + SuffixFrontendAuthBasic
	TraefikFrontendEntryPoints                     = Prefix + SuffixFrontendEntryPoints
	TraefikFrontendPassHostHeader                  = Prefix + SuffixFrontendPassHostHeader
	TraefikFrontendPassTLSCert                     = Prefix + SuffixFrontendPassTLSCert
	TraefikFrontendPriority                        = Prefix + SuffixFrontendPriority
	TraefikFrontendRateLimitExtractorFunc          = Prefix + SuffixFrontendRateLimitExtractorFunc
	TraefikFrontendRedirectEntryPoint              = Prefix + SuffixFrontendRedirectEntryPoint
	TraefikFrontendRedirectRegex                   = Prefix + SuffixFrontendRedirectRegex
	TraefikFrontendRedirectReplacement             = Prefix + SuffixFrontendRedirectReplacement
	TraefikFrontendRule                            = Prefix + SuffixFrontendRule
	TraefikFrontendRuleType                        = Prefix + SuffixFrontendRuleType // k8s only
	TraefikFrontendWhitelistSourceRange            = Prefix + SuffixFrontendWhitelistSourceRange
	TraefikFrontendHeaders                         = Prefix + SuffixFrontendHeaders
	TraefikFrontendRequestHeaders                  = Prefix + SuffixFrontendRequestHeaders
	TraefikFrontendResponseHeaders                 = Prefix + SuffixFrontendResponseHeaders
	TraefikFrontendAllowedHosts                    = Prefix + SuffixFrontendHeadersAllowedHosts
	TraefikFrontendHostsProxyHeaders               = Prefix + SuffixFrontendHeadersHostsProxyHeaders
	TraefikFrontendSSLRedirect                     = Prefix + SuffixFrontendHeadersSSLRedirect
	TraefikFrontendSSLTemporaryRedirect            = Prefix + SuffixFrontendHeadersSSLTemporaryRedirect
	TraefikFrontendSSLHost                         = Prefix + SuffixFrontendHeadersSSLHost
	TraefikFrontendSSLProxyHeaders                 = Prefix + SuffixFrontendHeadersSSLProxyHeaders
	TraefikFrontendSTSSeconds                      = Prefix + SuffixFrontendHeadersSTSSeconds
	TraefikFrontendSTSIncludeSubdomains            = Prefix + SuffixFrontendHeadersSTSIncludeSubdomains
	TraefikFrontendSTSPreload                      = Prefix + SuffixFrontendHeadersSTSPreload
	TraefikFrontendForceSTSHeader                  = Prefix + SuffixFrontendHeadersForceSTSHeader
	TraefikFrontendFrameDeny                       = Prefix + SuffixFrontendHeadersFrameDeny
	TraefikFrontendCustomFrameOptionsValue         = Prefix + SuffixFrontendHeadersCustomFrameOptionsValue
	TraefikFrontendContentTypeNosniff              = Prefix + SuffixFrontendHeadersContentTypeNosniff
	TraefikFrontendBrowserXSSFilter                = Prefix + SuffixFrontendHeadersBrowserXSSFilter
	TraefikFrontendContentSecurityPolicy           = Prefix + SuffixFrontendHeadersContentSecurityPolicy
	TraefikFrontendPublicKey                       = Prefix + SuffixFrontendHeadersPublicKey
	TraefikFrontendReferrerPolicy                  = Prefix + SuffixFrontendHeadersReferrerPolicy
	TraefikFrontendIsDevelopment                   = Prefix + SuffixFrontendHeadersIsDevelopment
	BaseFrontendErrorPage                          = "frontend.errors."
	SuffixErrorPageBackend                         = "backend"
	SuffixErrorPageQuery                           = "query"
	SuffixErrorPageStatus                          = "status"
	BaseFrontendRateLimit                          = "frontend.rateLimit.rateSet."
	SuffixRateLimitPeriod                          = "period"
	SuffixRateLimitAverage                         = "average"
	SuffixRateLimitBurst                           = "burst"
)
