// Package to identify crawlers based on their User Agent.
//  BenchmarkExactMatch/case0-8             10000000               182 ns/op
//  BenchmarkExactMatch/case101-8           10000000               128 ns/op
//  BenchmarkExactMatch/case200-8           10000000               137 ns/op
//  BenchmarkExactMatch/case300-8           10000000               124 ns/op
//  BenchmarkExactMatch/case400-8           20000000               113 ns/op
//  BenchmarkExactMatch/miss-8              200000000                8.03 ns/op
//  BenchmarkRegExpMatch/case0-8             5000000               292 ns/op
//  BenchmarkRegExpMatch/case101-8            200000              7335 ns/op
//  BenchmarkRegExpMatch/case200-8            100000             12866 ns/op
//  BenchmarkRegExpMatch/case300-8            100000             21898 ns/op
//  BenchmarkRegExpMatch/case400-8             50000             26515 ns/op
//  BenchmarkRegExpMatch/miss-8                50000             23963 ns/op
package crawlerflagger
