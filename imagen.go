package imagen

// Generic source definition
type Source struct {
	Name string   `yaml:"name"` // Package name, git repo etc
	Refs []string `yaml:"refs"` // Version, git revision etc
}

type Labels map[string]string

type ConfigFile struct {
	Configs []Config `json:"configs"`
}

type Config struct {
	Template string   `yaml:"template"`
	Bases    []Source `yaml:"bases"`
	Sources  []Source `yaml:"sources"`
	Labels   Labels   `yaml:"labels"`
}
