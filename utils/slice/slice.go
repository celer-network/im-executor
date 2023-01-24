package slice

// Map returns a new slice populated with the result of calling the provided function
// on every element in the provided input slice.
func Map[T1, T2 any](input []T1, f func(T1) T2) (output []T2) {
	output = make([]T2, 0, len(input))
	for _, v := range input {
		output = append(output, f(v))
	}
	return output
}

// Filter returns a new slice with all elements from the from the
// input elements for which the provided predicate function returns true.
func Filter[T any](input []T, pred func(T) bool) (output []T) {
	for _, v := range input {
		if pred(v) {
			output = append(output, v)
		}
	}
	return output
}

// Reduce executes a provided function on each element of the slice
// in order, passing the return value of the previous function call
// on the preceding slice element. The final result of running the
// provided function across all slice elements is returned.
func Reduce[T1, T2 any](input []T1, f func(T2, T1) T2) T2 {
	var acc T2

	for _, v := range input {
		acc = f(acc, v)
	}
	return acc
}

// ReduceWithInitialValue executes a provided function on each element of the slice
// in order, passing the return value of the previous function call
// on the preceding slice element. Unlike Reduce, the initial value of
// the accumulator can be provided as an argument. The final result of running the
// provided function across all slice elements is returned.
func ReduceWithInitialValue[T1, T2 any](input []T1, acc T2, f func(T2, T1) T2) T2 {
	for _, v := range input {
		acc = f(acc, v)
	}
	return acc
}
