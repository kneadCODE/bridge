module github.com/kneadCODE/bridge/src/golib

go 1.19

require (
	github.com/99designs/gqlgen v0.17.22
	github.com/go-chi/chi/v5 v5.0.8
	github.com/stretchr/testify v1.7.1
	github.com/vektah/gqlparser/v2 v2.5.1
	golang.org/x/exp v0.0.0-20221215174704-0915cd710c24
)

require (
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 => golang.org/x/crypto v0.4.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 => golang.org/x/crypto v0.4.0
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859 => golang.org/x/net v0.4.0
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 => golang.org/x/net v0.4.0
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f => golang.org/x/net v0.4.0
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b => golang.org/x/net v0.4.0
	golang.org/x/net v0.3.0 => golang.org/x/net v0.4.0
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a => golang.org/x/sys v0.3.0
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 => golang.org/x/sys v0.3.0
	golang.org/x/sys v0.0.0-20210423082822-04245dca01da => golang.org/x/sys v0.3.0
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 => golang.org/x/sys v0.3.0
	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 => golang.org/x/sys v0.3.0
	golang.org/x/text v0.3.0 => golang/golang.org/x/text v0.5.0
	golang.org/x/text v0.3.3 => golang/golang.org/x/text v0.5.0
	golang.org/x/text v0.3.6 => golang/golang.org/x/text v0.5.0
	gopkg.in/yaml.v2 v2.2.2 => gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v2 v2.2.4 => gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v2 v2.2.8 => gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c => gopkg.in/yaml.v3 v3.0.1
)
