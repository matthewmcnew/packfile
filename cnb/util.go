package cnb

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"

	"github.com/sclevine/packfile"
)

func writeTOML(lt interface{}, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return toml.NewEncoder(f).Encode(lt)
}

func eachDir(dir string, fn func(name string) error) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		if err := fn(f.Name()); err != nil {
			return err
		}
	}
	return nil
}

func shellOverride(exec packfile.Exec, shell string) packfile.Exec {
	if exec.Shell == "" {
		exec.Shell = shell
	}
	return exec
}