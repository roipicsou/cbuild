package persfile

type Config struct {
	Compilateur string   `yaml:"compilateur"`
	Executable  string   `yaml:"executable"`
	ListDir     []string `yaml:"listDir"`
}

func NewConfig(exec string, name string) *Config {
	return &Config{
		Compilateur: exec,
		Executable:  name,
		ListDir: []string{
			".",
		},
	}
}
