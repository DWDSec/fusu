package fusu

const (
	version  = "0.0.1"
	build    = "beta-dev"
	codename = "DWDSEC"
	intro    = "A distributed container competition platform developed by goalng."
)

func Version() string {
	return version
}

func VersionStatement() []string {
	return []string{
		"Fusu " + Version() + " (" + codename + ") " + build,
		intro,
	}
}
