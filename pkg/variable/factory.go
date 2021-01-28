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

package variable

import (
	"mosn.io/pkg/variable"
)

// AddVariable is used to check variable name exists. Typical usage is variables used in access logs.
// Deprecated: use mosn.io/pkg/variable/factory.go:GetVariableValue instead
var AddVariable = variable.AddVariable

// Deprecated: use mosn.io/pkg/variable/factory.go:RegisterVariable instead
var RegisterVariable = variable.RegisterVariable

// Deprecated: use mosn.io/pkg/variable/factory.go:RegisterPrefixVariable instead
var RegisterPrefixVariable = variable.RegisterPrefixVariable

// Deprecated: use mosn.io/pkg/variable/factory.go:NewVariableContext instead
var NewVariableContext = variable.NewVariableContext
