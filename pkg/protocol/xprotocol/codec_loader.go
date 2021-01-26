/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
			log.DefaultLogger.Infof("load go plugin codec succeed: %+v", dir+"/"+fileName)
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
