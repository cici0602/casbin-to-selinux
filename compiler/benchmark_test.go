// Copyright 2025 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions
// and limitations under the License.

package compiler

import (
	"testing"
)

// BenchmarkParser
func BenchmarkParser(b *testing.B) {
	modelPath := "../examples/httpd/httpd_model.conf"
	policyPath := "../examples/httpd/httpd_policy.csv"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parser := NewParser(modelPath, policyPath)
		_, err := parser.Parse()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkAnalyzer
func BenchmarkAnalyzer(b *testing.B) {
	parser := NewParser("../examples/httpd/httpd_model.conf", "../examples/httpd/httpd_policy.csv")
	pml, _ := parser.Parse()
	decoded, _ := parser.Decode(pml)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		analyzer := NewAnalyzer(decoded)
		if err := analyzer.Analyze(); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkGenerator 
func BenchmarkGenerator(b *testing.B) {
	parser := NewParser("../examples/httpd/httpd_model.conf", "../examples/httpd/httpd_policy.csv")
	pml, _ := parser.Parse()
	decoded, _ := parser.Decode(pml)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		generator := NewGenerator(decoded, "")
		_, err := generator.Generate()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkOptimizer 
// TODO: Optimization feature is not yet implemented
// func BenchmarkOptimizer(b *testing.B) {
// 	parser := NewParser("../examples/httpd/httpd_model.conf", "../examples/httpd/httpd_policy.csv")
// 	pml, _ := parser.Parse()
// 	decoded, _ := parser.Decode(pml)
// 	generator := NewGenerator(decoded, "")
// 	sePolicy, _ := generator.Generate()
//
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		// Optimization not implemented
// 	}
// }

// BenchmarkFullPipeline 
func BenchmarkFullPipeline(b *testing.B) {
	modelPath := "../examples/httpd/httpd_model.conf"
	policyPath := "../examples/httpd/httpd_policy.csv"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Parse
		parser := NewParser(modelPath, policyPath)
		pml, err := parser.Parse()
		if err != nil {
			b.Fatal(err)
		}

		// Decode
		decoded, err := parser.Decode(pml)
		if err != nil {
			b.Fatal(err)
		}

		// Analyze
		analyzer := NewAnalyzer(decoded)
		if err := analyzer.Analyze(); err != nil {
			b.Fatal(err)
		}

		// Generate
		generator := NewGenerator(decoded, "")
		_, err = generator.Generate()
		if err != nil {
			b.Fatal(err)
		}

		// Optimize
		// TODO: Optimization feature is not yet implemented
	}
}
