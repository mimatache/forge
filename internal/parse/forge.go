package parse

import (
	"gopkg.in/yaml.v2"
	"io"

	"mimatache/github.com/forge/internal/manifest"
)


type ForgeReader func(filePath string) (io.ReadCloser, error)

// GetForge recursively goes over the given Forge file given and the imported Forge files
// and returns a single Forge object.
// GetForge verifies that no circular actions are defined
func GetForge(filePath string, forgeReader ForgeReader) (*manifest.Forge, error) {
	forges := map[string]*manifest.Forge{}
	err := mapForges(forges, filePath, forgeReader)
	if err != nil {
		return nil, err
	}
	frg := &manifest.Forge{}
	for k, v := range forges {
		frg.Include = append(frg.Include, k)
		frg.Forgeries = append(frg.Forgeries, v.Forgeries...)
	}
	return frg, nil
}

func GetForgeries(filePath string, reader ForgeReader) (forgeries map[manifest.ForgeryName]manifest.Forgery, err error) {
	forgeries = map[manifest.ForgeryName]manifest.Forgery{}
	frg, err := GetForge(filePath, reader)
	if err != nil {
		return nil, err
	}
	for _, v := range frg.Forgeries {
		forgeries[v.Name] = v
	}
	return forgeries, nil
}


// addToForge adds a forge to a map only if the file is not already present in the map
func addToForges(forges map[string]*manifest.Forge, filePath string, reader ForgeReader) (bool, error) {
	_, ok := forges[filePath]
	if ok {
		return false, nil
	}
	contents, err := reader(filePath)
	defer func () {
		_ = contents.Close()
	} ()
	if err != nil {
		return false, err
	}
	frg := &manifest.Forge{}
	err = yaml.NewDecoder(contents).Decode(frg)
	if err != nil {
		return false, err
	}
	if err = frg.Forgeries.Validate(); err != nil {
		return false, err
	}
	forges[filePath] = frg
	return true, nil
}

// mapForges adds all forge files to a map where the key is the file location and the value is the individual forge
func mapForges(forges map[string]*manifest.Forge, filePath string, forgeReader ForgeReader) error{
	ok, err := addToForges(forges, filePath, forgeReader)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	for _, v := range forges[filePath].Include {
		err = mapForges(forges, v, forgeReader)
		if err != nil {
			return err
		}
	}
	return nil
}