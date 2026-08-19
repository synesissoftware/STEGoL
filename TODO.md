# STEGoL - TODO <!-- omit in toc -->


## Table of Contents <!-- omit in toc -->

- [Functional improvements](#functional-improvements)
- [Performance improvements](#performance-improvements)
- [Packaging improvements](#packaging-improvements)


## Functional improvements

* [ ] strict same-type, same-value comparisons;
* [ ] same-value, somewhat loose types' comparison (e.g. a `uint32` value compared with an `int64` is accepted, but `uint32` with `int32` is conditional on the actual `uint32` value);
* [ ] fully laissez-faire (e.g. comparing a `uint32` to a string-containing-integer-value) -> should only allow at-most-one type conversion;


## Performance improvements

* \<none>


## Packaging improvements

* [ ] Before the next official release: confirm **`go.mod`** (`go 1.21`) and the CI Go-version matrix, bump Synesis `require`s to newly published tags, then run **`go mod tidy`** (not against currently published tags). Prior Synesis Go releases, in order:
  * **ver2go**;


<!-- ########################### end of file ########################### -->
