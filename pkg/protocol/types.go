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

package protocol

import (
	"mosn.io/pkg/protocol"
)

// ProtocolName type definition
const (
	Auto      = protocol.Auto
	SofaRPC   = protocol.SofaRPC
	HTTP1     = protocol.HTTP1
	HTTP2     = protocol.HTTP2
	Xprotocol = protocol.Xprotocol
)

// header direction definition
const (
	Request  = protocol.Request
	Response = protocol.Response
)

// Host key for routing in MOSN Header
const (
	MosnHeaderDirection       = protocol.MosnHeaderDirection
	MosnHeaderScheme          = protocol.MosnHeaderScheme
	MosnHeaderHostKey         = protocol.MosnHeaderHostKey
	MosnHeaderPathKey         = protocol.MosnHeaderPathKey
	MosnHeaderQueryStringKey  = protocol.MosnHeaderQueryStringKey
	MosnHeaderMethod          = protocol.MosnHeaderMethod
	MosnOriginalHeaderPathKey = protocol.MosnOriginalHeaderPathKey
)

// Hseader with special meaning in istio
// todo maybe use ":authority"
const (
	IstioHeaderHostKey = protocol.IstioHeaderHostKey
)

// TODO: move CommonHeader to common, not only in protocol

// CommonHeader wrapper for map[string]string
type CommonHeader = protocol.CommonHeader
