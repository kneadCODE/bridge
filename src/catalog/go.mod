module github.com/kneadCODE/bridge/src/catalog

go 1.19

require (
	github.com/99designs/gqlgen v0.17.22
	github.com/go-chi/chi/v5 v5.0.8
	github.com/kneadCODE/bridge/src/golib v0.0.0
	github.com/vektah/gqlparser/v2 v2.5.1
	golang.org/x/exp v0.0.0-20221208152030-732eee02a75a
)

require (
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	github.com/urfave/cli/v2 v2.8.1 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	golang.org/x/mod v0.6.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.3.8 // indirect
	golang.org/x/tools v0.2.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/kneadCODE/bridge/src/golib v0.0.0 => ../golib
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c => gopkg.in/yaml.v3 v3.0.1
)
