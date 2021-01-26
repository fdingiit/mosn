package xprotocol

import (
	"io/ioutil"
	"plugin"
	"strings"

	"mosn.io/mosn/pkg/log"
)

const (
	LoaderFunctionName string = "LoadCodec"
)

func InitCodec(dir string) error {
	fileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, info := range fileInfo {
		fileName := info.Name()
		if strings.Contains(fileName, ".so") {
			if err := readProtocolPlugin(dir + "/" + fileName); err != nil {
				log.DefaultLogger.Warnf("load go plugin codec failed: %+v, err: %+v", dir+"/"+fileName, err)
			}
		}
	}

	return nil
}

func readProtocolPlugin(path string) error {
	p, err := plugin.Open(path)
	if err != nil {
		return err
	}

	sym, err := p.Lookup(LoaderFunctionName)
	if err != nil {
		return err
	}

	loadFunc := sym.(func() error)
	return loadFunc()
}
