package libs

// Options global options
type Options struct {
	RootFolder    string
	ConfigFile    string
	LogFile       string
	SSHPrivateKey string
	PassPhrase    string
	Credentials   string

	Port string
	User string

	Commands    []string
	Concurrency int
	Retry       int
	Verbose     bool
	Debug       bool

	InputFile string
	Inputs    []string
}
