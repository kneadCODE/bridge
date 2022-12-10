package app

type Env string

const (
	// EnvDev is the dev env
	EnvDev = Env("dev")
	// EnvStaging is the staging env
	EnvStaging = Env("staging")
	// EnvProd is the prod env
	EnvProd = Env("prod")
)

// String returns the string representation of Env
func (e Env) String() string {
	return string(e)
}
