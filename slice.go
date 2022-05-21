package nits

// sliceUtility is an empty structure that is prepared only for creating methods.
type sliceUtility struct{}

// Slice is an entity that allows the methods of SliceUtility to be executed from outside the package without initializing SliceUtility.
// nolint: gochecknoglobals
var Slice sliceUtility

// ContainsInt returns whether or not the passed slice contains the passed value.
func (sliceUtility) ContainsInt(slice []int, value int) bool {
	for _, elem := range slice {
		if value == elem {
			return true
		}
	}

	return false
}

// ContainsString returns whether or not the passed slice contains the passed value.
func (sliceUtility) ContainsString(slice []string, value string) bool {
	for _, elem := range slice {
		if value == elem {
			return true
		}
	}

	return false
}

// EqualInt will return true if the elements of the two slices are identical.
func (sliceUtility) EqualInt(sliceA, sliceB []int) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}

	for index, elem := range sliceA {
		if elem != sliceB[index] {
			return false
		}
	}

	return true
}

// EqualString will return true if the elements of the two slices are identical.
func (sliceUtility) EqualString(sliceA, sliceB []string) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}

	for index, elem := range sliceA {
		if elem != sliceB[index] {
			return false
		}
	}

	return true
}

// ExcludeString returns an array of `slice` minus `exclude`.
func (sliceUtility) ExcludeString(slice []string, exclude string) (excluded []string) {
	for _, v := range slice {
		if v == exclude {
			continue
		}

		excluded = append(excluded, v)
	}

	return excluded
}
