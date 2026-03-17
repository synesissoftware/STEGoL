# STEGoL <!-- omit in toc -->

[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GitHub release](https://img.shields.io/github/v/release/synesissoftware/STEGoL.svg)](https://github.com/synesissoftware/STEGoL/releases/latest)
[![Last Commit](https://img.shields.io/github/last-commit/synesissoftware/STEGoL)](https://github.com/synesissoftware/STEGoL/commits/master)
[![Go](https://github.com/synesissoftware/STEGoL/actions/workflows/go.yml/badge.svg)](https://github.com/synesissoftware/STEGoL/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/synesissoftware/STEGoL)](https://goreportcard.com/report/github.com/synesissoftware/STEGoL)
[![Go Reference](https://pkg.go.dev/badge/github.com/synesissoftware/STEGoL.svg)](https://pkg.go.dev/github.com/synesissoftware/STEGoL)

**S**imple **T**esting **E**nhancements for the **Go** **L**anguage


## Introduction

T.B.C.


## Table of Contents <!-- omit in toc -->

- [Introduction](#introduction)
- [Installation](#installation)
- [Components](#components)
- [Examples](#examples)
- [Project Information](#project-information)
	- [Where to get help](#where-to-get-help)
	- [Contribution guidelines](#contribution-guidelines)
	- [Dependencies](#dependencies)
		- [Development/Testing Dependencies](#developmenttesting-dependencies)
	- [Related projects](#related-projects)
	- [License](#license)


## Installation

```Go
import stegol "github.com/synesissoftware/STEGoL"
```


## Components

```Go
// Evaluates two strings for equality, calling testing.T.Errorf() if the
// evaluation fails. The evaluation is done by equality comparison.
func CheckStringEqual(t *testing.T, expected, actual string, options any)

// Evaluates two strings for inequality, calling testing.T.Errorf() if the
// evaluation fails. The evaluation is done by equality comparison.
func CheckStringNotEqual(t *testing.T, expected, actual string, options any)

// Evaluates two strings for equality, calling testing.T.Errorf() if the
// evaluation fails. The evaluation is done by equality comparison after
// chomping the actual value.
func CheckStringEqualChomped(t *testing.T, expected, actual string, options any)

// Evaluates two strings for inequality, calling testing.T.Errorf() if the
// evaluation fails. The evaluation is done by equality comparison after
// chomping the actual value.
func CheckStringNotEqualChomped(t *testing.T, expected, actual string, options any)

// Evaluates two strings for equality, calling testing.T.Errorf() if the
// evaluation fails. The evaluation is done by equality comparison after
// whitespace-trimming the actual value.
func CheckStringEqualTrimmed(t *testing.T, expected, actual string, options any)

// Evaluates two strings for inequality, calling testing.T.Errorf() if the
// evaluation fails. The evaluation is done by equality comparison after
// whitespace-trimming the actual value.
func CheckStringNotEqualTrimmed(t *testing.T, expected, actual string, options any)

// Evaluates two strings for equality, calling testing.T.Errorf() if the
// evaluation fails. The evaluation is done by equality comparison ignoring
// the case of the strings, via the strings.EqualFold() standard library
// function.
func CheckStringEqualIgnoreCase(t *testing.T, expected, actual string, options any)

// Evaluates two strings for inequality, calling testing.T.Errorf() if the
// evaluation fails. The evaluation is done by equality comparison ignoring
// the case of the strings, via the strings.EqualFold() standard library
// function.
func CheckStringNotEqualIgnoreCase(t *testing.T, expected, actual string, options any)

func CheckStringByStringMatch(t *testing.T, pattern string, actual string, options any)

// Evaluates two strings, calling testing.T.Errorf() if the evaluation
// fails. The evaluation is done by the given caller-supplied comparison
// function, whose brief descriptor comparison_type, e.g.
// "regular expression" will be prefixed with the string
// "when compared by ".
func CheckStringCompare(t *testing.T, expected, actual string, fn StringCompareFunc, comparison_type string, options any)

// Evaluates a string for equality against an array of string values,
// calling testing.T.Errorf() if every evaluation fails. Each evaluation is
// done by equality comparison.
func CheckStringEqualAny(t *testing.T, expecteds []string, actual string, options any)
```


## Examples

Examples are provided in the `examples` directory, along with a markdown description for each. A detailed list TOC of them is provided in [EXAMPLES.md](./EXAMPLES.md).


## Project Information


### Where to get help

[GitHub Page](https://github.com/synesissoftware/STEGoL "GitHub Page")


### Contribution guidelines

Defect reports, feature requests, and pull requests are welcome on https://github.com/synesissoftware/STEGoL.


### Dependencies

* [**ver2go**](https://github.com/synesissoftware/ver2go/);


#### Development/Testing Dependencies

* [**require**](https://github.com/stretchr/testify/);


### Related projects

T.B.C.


### License

**STEGoL** is released under the 3-clause BSD license. See [LICENSE](./LICENSE) for details.

