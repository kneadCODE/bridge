module github.com/kneadCODE/bridge/src/gatekeeper

go 1.19

require (
	github.com/kneadCODE/bridge/src/golib v0.0.0
	golang.org/x/exp v0.0.0-20221208152030-732eee02a75a
)

require (
	github.com/go-chi/chi/v5 v5.0.8 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
)

replace (
	github.com/kneadCODE/bridge/src/golib v0.0.0 => ../golib
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c => gopkg.in/yaml.v3 v3.0.1
)
