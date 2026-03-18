package stegol

import (
	"testing"
)

func assert_checkIntegerEqual[E Integer, A Integer](t *testing.T, expected E, actual A, want bool) {
	t.Helper()
	if got, _ := checkIntegerEqual(expected, actual); got != want {
		t.Fatalf("checkIntegerEqual(%[1]T(%[1]v), %[2]T(%[2]v))=%v; want %v", expected, actual, got, want)
	}
}

func assert_checkIntegerGreater[E Integer, A Integer](t *testing.T, expected E, actual A, want bool) {
	t.Helper()
	if got, _ := checkIntegerGreater(expected, actual); got != want {
		t.Fatalf("checkIntegerGreater(%[1]T(%[1]v), %[2]T(%[2]v))=%v; want %v", expected, actual, got, want)
	}
}

func assert_checkIntegerGreaterOrEqual[E Integer, A Integer](t *testing.T, expected E, actual A, want bool) {
	t.Helper()
	if got, _ := checkIntegerGreaterOrEqual(expected, actual); got != want {
		t.Fatalf("checkIntegerGreaterOrEqual(%[1]T(%[1]v), %[2]T(%[2]v))=%v; want %v", expected, actual, got, want)
	}
}

func assert_checkIntegerLess[E Integer, A Integer](t *testing.T, expected E, actual A, want bool) {
	t.Helper()
	if got, _ := checkIntegerLess(expected, actual); got != want {
		t.Fatalf("checkIntegerLess(%[1]T(%[1]v), %[2]T(%[2]v))=%v; want %v", expected, actual, got, want)
	}
}

func assert_checkIntegerLessOrEqual[E Integer, A Integer](t *testing.T, expected E, actual A, want bool) {
	t.Helper()
	if got, _ := checkIntegerLessOrEqual(expected, actual); got != want {
		t.Fatalf("checkIntegerLessOrEqual(%[1]T(%[1]v), %[2]T(%[2]v))=%v; want %v", expected, actual, got, want)
	}
}

func assert_checkIntegerNotEqual[E Integer, A Integer](t *testing.T, expected E, actual A, want bool) {
	t.Helper()
	if got, _ := checkIntegerNotEqual(expected, actual); got != want {
		t.Fatalf("checkIntegerNotEqual(%[1]T(%[1]v), %[2]T(%[2]v))=%v; want %v", expected, actual, got, want)
	}
}

func Test_checkIntegerEqual(t *testing.T) {
	// Same-type sanity for all supported types (both signedness branches).
	for _, v := range []int64{-1, 0, 1, 42} {
		assert_checkIntegerEqual(t, int(v), int(v), true)
		assert_checkIntegerEqual(t, int8(v), int8(v), true)
		assert_checkIntegerEqual(t, int16(v), int16(v), true)
		assert_checkIntegerEqual(t, int32(v), int32(v), true)
		assert_checkIntegerEqual(t, int64(v), int64(v), true)
	}
	for _, v := range []uint64{0, 1, 42, 255} {
		assert_checkIntegerEqual(t, uint(v), uint(v), true)
		assert_checkIntegerEqual(t, uint8(v), uint8(v), true)
		assert_checkIntegerEqual(t, uint16(v), uint16(v), true)
		assert_checkIntegerEqual(t, uint32(v), uint32(v), true)
		assert_checkIntegerEqual(t, uint64(v), uint64(v), true)
		assert_checkIntegerEqual(t, uintptr(v), uintptr(v), true)
	}

	// Cross-type equality (ensures every concrete type appears on both sides).
	assert_checkIntegerEqual(t, int(0), int8(0), true)
	assert_checkIntegerEqual(t, int8(1), int16(1), true)
	assert_checkIntegerEqual(t, int16(42), int32(42), true)
	assert_checkIntegerEqual(t, int32(-1), int64(-1), true)
	assert_checkIntegerEqual(t, int64(0), int(0), true)

	assert_checkIntegerEqual(t, uint(0), uint8(0), true)
	assert_checkIntegerEqual(t, uint8(1), uint16(1), true)
	assert_checkIntegerEqual(t, uint16(42), uint32(42), true)
	assert_checkIntegerEqual(t, uint32(255), uint64(255), true)
	assert_checkIntegerEqual(t, uint64(0), uintptr(0), true)

	assert_checkIntegerEqual(t, int8(0), uint64(0), true)
	assert_checkIntegerEqual(t, uint8(1), int64(1), true)
	assert_checkIntegerEqual(t, int16(-1), uint32(0), false)
	assert_checkIntegerEqual(t, uint16(1), int8(0), false)

	// Edge/boundary values (exercise conversions and comparison corners).
	assert_checkIntegerEqual(t, int8(-128), int64(-128), true)
	assert_checkIntegerEqual(t, int8(127), uint8(127), true)
	assert_checkIntegerEqual(t, uint8(255), int16(255), true)
	assert_checkIntegerEqual(t, int16(-32768), uint64(0), false)
	assert_checkIntegerEqual(t, uint16(65535), int32(65535), true)
	assert_checkIntegerEqual(t, uint32(^uint32(0)), uint64(^uint32(0)), true)
	assert_checkIntegerEqual(t, uint64(^uint64(0)), uint64(^uint64(0)), true)
	assert_checkIntegerEqual(t, uintptr(^uintptr(0)), uintptr(^uintptr(0)), true)

	// Ensure uint64 "too large for int64" compares as not-equal when paired.
	assert_checkIntegerEqual(t, int64(^uint64(0)>>1), uint64(^uint64(0)), false)
}

func Test_CheckIntegerEqual(t *testing.T) {
	// Same-type sanity for all supported types (both signedness branches).
	for _, v := range []int64{-1, 0, 1, 42} {
		CheckIntegerEqual(t, int(v), int(v))
		CheckIntegerEqual(t, int8(v), int8(v))
		CheckIntegerEqual(t, int16(v), int16(v))
		CheckIntegerEqual(t, int32(v), int32(v))
		CheckIntegerEqual(t, int64(v), int64(v))
	}
	for _, v := range []uint64{0, 1, 42, 255} {
		CheckIntegerEqual(t, uint(v), uint(v))
		CheckIntegerEqual(t, uint8(v), uint8(v))
		CheckIntegerEqual(t, uint16(v), uint16(v))
		CheckIntegerEqual(t, uint32(v), uint32(v))
		CheckIntegerEqual(t, uint64(v), uint64(v))
		CheckIntegerEqual(t, uintptr(v), uintptr(v))
	}

	// Cross-type equality (ensures every concrete type appears on both sides).
	CheckIntegerEqual(t, int(0), int8(0))
	CheckIntegerEqual(t, int8(1), int16(1))
	CheckIntegerEqual(t, int16(42), int32(42))
	CheckIntegerEqual(t, int32(-1), int64(-1))
	CheckIntegerEqual(t, int64(0), int(0))

	CheckIntegerEqual(t, uint(0), uint8(0))
	CheckIntegerEqual(t, uint8(1), uint16(1))
	CheckIntegerEqual(t, uint16(42), uint32(42))
	CheckIntegerEqual(t, uint32(255), uint64(255))
	CheckIntegerEqual(t, uint64(0), uintptr(0))

	CheckIntegerEqual(t, int8(0), uint64(0))
	CheckIntegerEqual(t, uint8(1), int64(1))

	// Edge/boundary values (exercise conversions and comparison corners).
	CheckIntegerEqual(t, int8(-128), int64(-128))
	CheckIntegerEqual(t, int8(127), uint8(127))
	CheckIntegerEqual(t, uint8(255), int16(255))
	CheckIntegerEqual(t, uint16(65535), int32(65535))
	CheckIntegerEqual(t, uint32(^uint32(0)), uint64(^uint32(0)))
	CheckIntegerEqual(t, uint64(^uint64(0)), uint64(^uint64(0)))
	CheckIntegerEqual(t, uintptr(^uintptr(0)), uintptr(^uintptr(0)))
}

func Test_checkIntegerGreater(t *testing.T) {
	// Representative values across both signedness branches (pass + fail cases).
	assert_checkIntegerGreater(t, int8(-1), int8(0), true)
	assert_checkIntegerGreater(t, int8(0), int8(0), false)
	assert_checkIntegerGreater(t, int8(1), int8(0), false)

	assert_checkIntegerGreater(t, int16(-10), int32(10), true)
	assert_checkIntegerGreater(t, int32(-123_456_789), int64(123_456_789), true)
	assert_checkIntegerGreater(t, int16(10), int32(-10), false)
	assert_checkIntegerGreater(t, int32(123_456_789), int64(-123_456_789), false)

	assert_checkIntegerGreater(t, int8(-1), uint8(0), true) // 0 > -1
	assert_checkIntegerGreater(t, uint8(0), int8(-1), false) // -1 not > 0
	assert_checkIntegerGreater(t, uint8(1), int8(2), true)   // 2 > 1
	assert_checkIntegerGreater(t, int16(1), uint8(0), false)  // 0 !> 1

	assert_checkIntegerGreater(t, uintptr(0), uint(1), true)
	assert_checkIntegerGreater(t, uint(1), uintptr(0), false)
	assert_checkIntegerGreater(t, uint32(42), uintptr(42), false)
	assert_checkIntegerGreater(t, uintptr(43), uint32(42), false)

	// Large-boundary comparisons involving uint64/int64 interaction.
	assert_checkIntegerGreater(t, int64(^uint64(0)>>1), uint64(^uint64(0)), true) // MaxUint64 > MaxInt64
	assert_checkIntegerGreater(t, uint64(^uint64(0)), int64(^uint64(0)>>1), false)
}

func Test_CheckIntegerGreater(t *testing.T) {
	// Representative values across both signedness branches.
	CheckIntegerGreater(t, int8(-1), int8(0))
	CheckIntegerGreater(t, int16(-10), int32(10))
	CheckIntegerGreater(t, uint8(1), int8(2))
	CheckIntegerGreater(t, uintptr(0), uint(1))
	CheckIntegerGreater(t, int64(^uint64(0)>>1), uint64(^uint64(0)))
}

func Test_checkIntegerGreaterOrEqual(t *testing.T) {
	// Pass cases.
	assert_checkIntegerGreaterOrEqual(t, int8(0), int8(0), true)
	assert_checkIntegerGreaterOrEqual(t, int8(0), int8(1), true)
	assert_checkIntegerGreaterOrEqual(t, uint8(42), int16(42), true)
	assert_checkIntegerGreaterOrEqual(t, uint8(42), int16(43), true)

	// Fail cases.
	assert_checkIntegerGreaterOrEqual(t, int8(1), int8(0), false)
	assert_checkIntegerGreaterOrEqual(t, uint8(1), int8(-1), false)

	// Large-boundary comparisons.
	assert_checkIntegerGreaterOrEqual(t, uint64(^uint64(0)), int64(^uint64(0)>>1), false)
	assert_checkIntegerGreaterOrEqual(t, int64(^uint64(0)>>1), uint64(^uint64(0)), true)
}

func Test_CheckIntegerGreaterOrEqual(t *testing.T) {
	CheckIntegerGreaterOrEqual(t, int8(0), int8(0))
	CheckIntegerGreaterOrEqual(t, int8(0), int8(1))
	CheckIntegerGreaterOrEqual(t, uint8(42), int16(42))
	CheckIntegerGreaterOrEqual(t, int64(^uint64(0)>>1), uint64(^uint64(0)))
}

func Test_checkIntegerLess(t *testing.T) {
	// Representative values across both signedness branches (pass + fail cases).
	assert_checkIntegerLess(t, int8(-1), int8(0), false)
	assert_checkIntegerLess(t, int8(0), int8(0), false)
	assert_checkIntegerLess(t, int8(1), int8(0), true)

	assert_checkIntegerLess(t, int16(-10), int32(10), false)
	assert_checkIntegerLess(t, int32(-123_456_789), int64(123_456_789), false)
	assert_checkIntegerLess(t, int16(10), int32(-10), true)
	assert_checkIntegerLess(t, int32(123_456_789), int64(-123_456_789), true)

	assert_checkIntegerLess(t, int8(-1), uint8(0), false)  // unsigned cannot be < negative expected
	assert_checkIntegerLess(t, uint8(0), int8(-1), true)   // negative actual always < unsigned expected
	assert_checkIntegerLess(t, uint8(1), int8(2), false)   // 2 !< 1
	assert_checkIntegerLess(t, int16(1), uint8(0), true)   // 0 < 1

	assert_checkIntegerLess(t, uintptr(0), uint(1), false)
	assert_checkIntegerLess(t, uint(1), uintptr(0), true)
	assert_checkIntegerLess(t, uint32(42), uintptr(42), false)
	assert_checkIntegerLess(t, uintptr(43), uint32(42), true)

	// Large-boundary comparisons involving uint64/int64 interaction.
	assert_checkIntegerLess(t, int64(^uint64(0)>>1), uint64(^uint64(0)), false) // MaxUint64 !< MaxInt64 (as actual)
	assert_checkIntegerLess(t, uint64(^uint64(0)), int64(^uint64(0)>>1), true)  // MaxInt64 < MaxUint64 (as actual)
}

func Test_CheckIntegerLess(t *testing.T) {
	// Representative values across both signedness branches (pass + fail cases).
	CheckIntegerLess(t, int8(1), int8(0))

	CheckIntegerLess(t, int16(10), int32(-10))
	CheckIntegerLess(t, int32(123_456_789), int64(-123_456_789))

	CheckIntegerLess(t, uint8(0), int8(-1))   // negative always < unsigned expected
	CheckIntegerLess(t, int16(1), uint8(0))   // 0 < 1

	CheckIntegerLess(t, uint(1), uintptr(0))

	// Large-boundary comparisons involving uint64/int64 interaction.
	CheckIntegerLess(t, uint64(^uint64(0)), int64(^uint64(0)>>1)) // MaxInt64 < MaxUint64 (as actual)
}

func Test_checkIntegerLessOrEqual(t *testing.T) {
	// Pass cases.
	assert_checkIntegerLessOrEqual(t, int8(0), int8(0), true)
	assert_checkIntegerLessOrEqual(t, int8(1), int8(0), true)
	assert_checkIntegerLessOrEqual(t, uint8(42), int16(42), true)
	assert_checkIntegerLessOrEqual(t, uint8(43), int16(42), true)

	// Fail cases.
	assert_checkIntegerLessOrEqual(t, int8(0), int8(1), false)
	assert_checkIntegerLessOrEqual(t, int8(-1), uint8(0), false) // 0 <= -1 is false

	// Large-boundary comparisons.
	assert_checkIntegerLessOrEqual(t, uint64(^uint64(0)), int64(^uint64(0)>>1), true)
	assert_checkIntegerLessOrEqual(t, int64(^uint64(0)>>1), uint64(^uint64(0)), false)
}

func Test_CheckIntegerLessOrEqual(t *testing.T) {
	CheckIntegerLessOrEqual(t, int8(0), int8(0))
	CheckIntegerLessOrEqual(t, int8(1), int8(0))
	CheckIntegerLessOrEqual(t, uint8(42), int16(42))
	CheckIntegerLessOrEqual(t, uint64(^uint64(0)), int64(^uint64(0)>>1))
}

func Test_checkIntegerNotEqual(t *testing.T) {
	assert_checkIntegerNotEqual(t, int8(0), int8(0), false)
	assert_checkIntegerNotEqual(t, int8(0), int8(1), true)
	assert_checkIntegerNotEqual(t, uint8(42), int16(42), false)
	assert_checkIntegerNotEqual(t, uint8(42), int16(43), true)
	assert_checkIntegerNotEqual(t, int16(-1), uint32(0), true)
	assert_checkIntegerNotEqual(t, uint8(1), int64(1), false)

	// Large-boundary comparisons.
	assert_checkIntegerNotEqual(t, int64(^uint64(0)>>1), uint64(^uint64(0)), true)
}

func Test_CheckIntegerNotEqual(t *testing.T) {
	CheckIntegerNotEqual(t, int8(0), int8(1))
	CheckIntegerNotEqual(t, uint8(42), int16(43))
	CheckIntegerNotEqual(t, int16(-1), uint32(0))
	CheckIntegerNotEqual(t, int64(^uint64(0)>>1), uint64(^uint64(0)))
}
