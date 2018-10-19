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

package grammars

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	_ "github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/nicolaferraro/astio-ml/pkg/grammars/java"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJava(t *testing.T) {

	ttt(t, "../../test/data/CaffeineCacheSample.java")
	ttt(t, "../../test/data/CaffeineCacheSample2.java")

}

func ttt(t *testing.T, n string) {
	charStream, err := antlr.NewFileStream(n)
	assert.Nil(t, err)
	lexer := java.NewJava9Lexer(charStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := java.NewJava9Parser(stream)

	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	tree := parser.CompilationUnit()
	antlr.ParseTreeWalkerDefault.Walk(&TreeShapeListener{}, tree)
}

type TreeShapeListener struct {
	*java.BaseJava9Listener
}

func (s *TreeShapeListener) EnterClassType(ctx *java.ClassTypeContext) {
	println("Cippalippa " + ctx.GetText())
}
