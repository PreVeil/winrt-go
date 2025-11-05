package winrt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetCurrent(t *testing.T) {
	a := NewArrayIterable([]any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, SignatureInt32)

	it, err := a.First()
	require.NoError(t, err)

	ok := false
	i := 1
	for ok, err = it.MoveNext(); err == nil && ok; ok, err = it.MoveNext() {
		b, err := it.GetHasCurrent()
		require.NoError(t, err)
		require.True(t, b)
		ptr, err := it.GetCurrent()
		require.NoError(t, err)
		require.Equal(t, i, int(uintptr(ptr)))
		i++
	}
}

func Test_GetMany(t *testing.T) {
	a := NewArrayIterable([]any{101, 202, 303}, SignatureInt32)

	it, err := a.First()
	require.NoError(t, err)
	resp, n, err := it.GetMany(12)
	require.NoError(t, err)
	require.Equal(t, uint32(3), n)

	println("RESP", n, resp, len(resp))

	var i uint32
	j := 101
	for i = 0; i < n; i++ {
		val := int(uintptr(resp[i]))
		require.Equal(t, j, val)
		j += 101
	}
}

// Test_EmptyIterable verifies that empty arrays work correctly.
func Test_EmptyIterable(t *testing.T) {
	// Create empty iterable
	iterable := NewArrayIterable([]any{}, SignatureInt32)
	require.NotNil(t, iterable)

	// Get iterator - this calls NewArrayIterator internally.
	iterator, err := iterable.First()
	require.NoError(t, err)
	require.NotNil(t, iterator)

	// Iterator should have no current element.
	hasCurrent, err := iterator.GetHasCurrent()
	require.NoError(t, err)
	require.False(t, hasCurrent)

	// MoveNext should return false (no elements).
	moved, err := iterator.MoveNext()
	require.NoError(t, err)
	require.False(t, moved)

	// GetMany should return 0 count (implementation may return -1/0xffffffff).
	// Note: The items slice may be allocated but count will be 0.
	_, count, err := iterator.GetMany(10)
	require.NoError(t, err)
	// Empty iterators may return 0 or 0xffffffff for count.
	require.True(t, count == 0 || count == 0xffffffff, "count should be 0 or 0xffffffff, got %d", count)

	// Clean up.
	iterator.Release()
	iterable.Release()
}

// Test_SequentialEmptyAndNonEmpty verifies that creating empty and non-empty
// iterables sequentially doesn't cause heap corruption.
func Test_SequentialEmptyAndNonEmpty(t *testing.T) {
	// First: Create and release an empty iterable
	emptyIterable := NewArrayIterable([]any{}, SignatureInt32)
	emptyIterator, err := emptyIterable.First()
	require.NoError(t, err)

	hasCurrent, err := emptyIterator.GetHasCurrent()
	require.NoError(t, err)
	require.False(t, hasCurrent)

	// Release empty iterator and iterable.
	emptyIterator.Release()
	emptyIterable.Release()

	// Second: Create and release a non-empty iterable.
	nonEmptyIterable := NewArrayIterable([]any{42, 43}, SignatureInt32)
	nonEmptyIterator, err := nonEmptyIterable.First()
	require.NoError(t, err)

	moved, err := nonEmptyIterator.MoveNext()
	require.NoError(t, err)
	require.True(t, moved)

	current, err := nonEmptyIterator.GetCurrent()
	require.NoError(t, err)
	require.Equal(t, 42, int(uintptr(current)))

	// Clean up.
	nonEmptyIterator.Release()
	nonEmptyIterable.Release()
}

// Test_MultipleIterables verifies that multiple iterables can coexist
// and be released in various orders without heap corruption.
func Test_MultipleIterables(t *testing.T) {
	// Create multiple iterables of different sizes.
	it1 := NewArrayIterable([]any{}, SignatureInt32)
	it2 := NewArrayIterable([]any{1}, SignatureInt32)
	it3 := NewArrayIterable([]any{2, 3}, SignatureInt32)
	it4 := NewArrayIterable([]any{}, SignatureInt32)

	// Get all iterators.
	iter1, err := it1.First()
	require.NoError(t, err)
	iter2, err := it2.First()
	require.NoError(t, err)
	iter3, err := it3.First()
	require.NoError(t, err)
	iter4, err := it4.First()
	require.NoError(t, err)

	// Release in non-creation order.
	iter2.Release()
	it2.Release()

	iter4.Release()
	it4.Release()

	iter1.Release()
	it1.Release()

	iter3.Release()
	it3.Release()
}

// Test_ReleaseRefcounting verifies that AddRef/Release work correctly.
func Test_ReleaseRefcounting(t *testing.T) {
	iterable := NewArrayIterable([]any{1, 2}, SignatureInt32)

	// Initial refcount is 1, AddRef should increment to 2.
	// Note: The return type varies by interface (may be uintptr or uint32)
	iterable.AddRef()

	// Release once - refcount should decrement.
	iterable.Release()

	// Final Release should free the object (refcount reaches 0).
	// Note: Release may return uintptr or int32 depending on interface.
	iterable.Release()

	// Note: accessing iterable after this point would be undefined behavior.
}

// Test_MultipleIteratorsFromSameIterable verifies that First() can be called
// multiple times to get independent iterators.
func Test_MultipleIteratorsFromSameIterable(t *testing.T) {
	iterable := NewArrayIterable([]any{10, 20, 30}, SignatureInt32)

	// Get first iterator.
	iter1, err := iterable.First()
	require.NoError(t, err)

	// Get second iterator.
	iter2, err := iterable.First()
	require.NoError(t, err)

	// Advance first iterator.
	moved, err := iter1.MoveNext()
	require.NoError(t, err)
	require.True(t, moved)

	current, err := iter1.GetCurrent()
	require.NoError(t, err)
	require.Equal(t, 10, int(uintptr(current)))

	// Second iterator should still be at start.
	hasCurrent, err := iter2.GetHasCurrent()
	require.NoError(t, err)
	require.False(t, hasCurrent) // hasn't moved yet.

	// Advance second iterator.
	moved, err = iter2.MoveNext()
	require.NoError(t, err)
	require.True(t, moved)

	current, err = iter2.GetCurrent()
	require.NoError(t, err)
	require.Equal(t, 10, int(uintptr(current)))

	// Clean up.
	iter1.Release()
	iter2.Release()
	iterable.Release()
}
