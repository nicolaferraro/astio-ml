/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package java

import (
	"github.com/nicolaferraro/astio-ml/pkg/model"
	"strings"
)

type CamelJava9Listener struct {
	*BaseJava9Listener
	Builder *model.IntegrationSketchBuilder
}

func NewCamelJava9Listener() *CamelJava9Listener {
	return &CamelJava9Listener{
		Builder: model.NewIntegrationSketchBuilder(),
	}
}

func (l *CamelJava9Listener) EnterMethodInvocation(ctx *MethodInvocationContext) {
	if strings.HasPrefix(ctx.GetText(), "from(") {
		l.Builder.NewRoute()
	}
}

func (l *CamelJava9Listener) ExitMethodInvocation(ctx *MethodInvocationContext) {
	l.doExit(ctx.Identifier().GetText(), ctx.ArgumentList())

	if l.Builder.InRoute() {
		l.Builder.EndRoute()
	}
}

func (l *CamelJava9Listener) ExitMethodInvocation_lf_primary(ctx *MethodInvocation_lf_primaryContext) {
	l.doExit(ctx.Identifier().GetText(), ctx.ArgumentList())
}

func (l *CamelJava9Listener) ExitMethodInvocation_lfno_primary(ctx *MethodInvocation_lfno_primaryContext) {
	ctx.MethodName().GetText()
	l.doExit(ctx.MethodName().GetText(), ctx.ArgumentList())
}

func (l *CamelJava9Listener) doExit(methodName string, arguments IArgumentListContext) {
	if methodName == "from" {
		l.Builder.From(strings.TrimSuffix(strings.TrimPrefix(arguments.GetText(), "\""), "\""))
	} else if methodName == "to" {
		l.Builder.AddTo(strings.TrimSuffix(strings.TrimPrefix(arguments.GetText(), "\""), "\""))
	}
}
