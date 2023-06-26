package buf

import (
	"gopkg.in/yaml.v3"
	"os"
)

type (
	// Dep holds the buf dependency from buf.lock file.
	Dep struct {
		Remote     string `yaml:"remote"`
		Owner      string `yaml:"owner"`
		Repository string `yaml:"repository"`
		Commit     string `yaml:"commit"`
	}
	// Lock holds the buf.lock file.
	Lock struct {
		Version string `yaml:"version"`
		Deps    Deps   `yaml:"deps"`
	}

	// Deps holds the list of buf dependency from buf.lock file.
	Deps []Dep
)

// Dependencies returns the list of repository dependencies name.
func (d Deps) Dependencies() (dependencies []string) {
	for _, dep := range d {
		dependencies = append(dependencies, dep.Repository)
	}
	return
}

// ParseBufLock parse dependencies from the buf.lock file.
func ParseBufLock(path string) (Lock, error) {
	f, err := os.Open(path)
	if err != nil {
		return Lock{}, err
	}
	defer f.Close()

	var lock Lock
	return lock, yaml.NewDecoder(f).Decode(&lock)
}
