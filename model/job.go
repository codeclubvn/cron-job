package model

// JobConfig represents a job configuration at a time
type JobConfig struct {
	Name          string       `yaml:"name"`
	Spec          string       `yaml:"spec"`
	HandlerConfig *HandlerHttp `yaml:"handler"`
}

// HandlerHttp represents a http handler configuration
type HandlerHttp struct {
	Method         string            `yaml:"method"`
	URL            string            `yaml:"url"`
	Headers        map[string]string `yaml:"headers"`
	Body           string            `yaml:"body"`
	TimeoutSeconds int               `yaml:"timeout_seconds"`
}
