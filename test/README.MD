# Manual test

export HYPOLAS_HEALTHCHECK_HTTP_URL=https://httpbin.org/get
export HYPOLAS_HEALTHCHECK_HTTP_EXPECTED=gzip
export HYPOLAS_HEALTHCHECK_HTTP_JSON=headers__Accept-Encoding

go test -v .