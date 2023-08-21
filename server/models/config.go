package models

type Config struct {
	AWS        []AWSConfig        `toml:"aws"`
	OCI        []OCIConfig        `toml:"oci"`
	Kubernetes []KubernetesConfig `toml:"k8s"`
}

type AWSConfig struct {
	Name    string `toml:"name"`
	Profile string `toml:"profile"`
	Source  string `toml:"source"`
	Path    string `toml:"path,omitempty"`
}

type OCIConfig struct {
	Name    string `toml:"name"`
	Profile string `toml:"profile"`
	Source  string `toml:"source"`
}

type KubernetesConfig struct {
	Name     string   `toml:"name"`
	Path     string   `toml:"path"`
	Contexts []string `toml:"contexts"`
}
