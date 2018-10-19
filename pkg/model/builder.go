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

package model

import "github.com/sirupsen/logrus"

type IntegrationSketchBuilder struct {
	currentIntegration *IntegrationSketch
	currentRoute       *RouteSketch
}

func NewIntegrationSketchBuilder() *IntegrationSketchBuilder {
	return &IntegrationSketchBuilder{
		currentIntegration: &IntegrationSketch{},
	}
}

func (b *IntegrationSketchBuilder) NewRoute() {
	b.currentRoute = &RouteSketch{}
}

func (b *IntegrationSketchBuilder) InRoute() bool {
	return b.currentRoute != nil
}

func (b *IntegrationSketchBuilder) EndRoute() {
	if b.currentRoute != nil {
		if b.currentIntegration.Routes == nil {
			b.currentIntegration.Routes = make([]RouteSketch, 0)
		}
		b.currentIntegration.Routes = append(b.currentIntegration.Routes, *b.currentRoute)
		b.currentRoute = nil
	} else {
		logrus.Warnf("parsing error: no route to complete")
	}
}

func (b *IntegrationSketchBuilder) From(uri string) {
	if b.currentRoute != nil {
		b.currentRoute.From = uri
	} else {
		logrus.Warnf("parsing error: no route to attach (from) uri %s", uri)
	}
}

func (b *IntegrationSketchBuilder) AddTo(uri string) {
	if b.currentRoute != nil {
		if b.currentRoute.To == nil {
			b.currentRoute.To = make([]string, 0)
		}
		b.currentRoute.To = append(b.currentRoute.To, uri)
	} else {
		logrus.Warnf("parsing error: no route to attach (to) uri %s", uri)
	}
}

func (b *IntegrationSketchBuilder) Build() *IntegrationSketch {
	return b.currentIntegration
}
