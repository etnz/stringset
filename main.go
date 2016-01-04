// Package stringset offers various tools to manipulate set of strings
//
// A stringset is represented as a map[string]interface{} where we do not make any assumption on the content of interface{}
package stringset

import "sort"

//Clone a set
//
func Clone(set map[string]interface{}) (clone map[string]interface{}) {
	clone = make(map[string]interface{})
	for k, v := range set {
		clone[k] = v
	}
	return
}

//Sort returns a sorted slice of all elements in the set
func Sort(set map[string]interface{}) (sorted []string) {
	sorted = make([]string, 0, len(set))
	for k := range set {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	return
}

//Contains returns true if the 'element' is contained in the set
func Contains(set map[string]interface{}, element string) bool {
	_, exists := set[element]
	return exists
}

//ContainsAny returns true if any of the 'elements' are contained in the set
func ContainsAny(set map[string]interface{}, elements ...string) bool {
	for _, e := range elements {
		if _, exists := set[e]; exists {
			return true
		}
	}
	return false
}

//ContainsAll returns true if all  'elements' are contained in the set
func ContainsAll(set map[string]interface{}, elements ...string) bool {
	for _, e := range elements {
		if _, exists := set[e]; !exists {
			return false
		}
	}
	return true
}

//Union all 'src' together into 'dst'
func Union(src ...map[string]interface{}) (dst map[string]interface{}) {
	dst = make(map[string]interface{})
	Append(dst, src...)
	return
}

//Append the Union of 'src' into 'dst'.
func Append(dst map[string]interface{}, src ...map[string]interface{}) {
	for _, source := range src {
		for k, v := range source {
			dst[k] = v
		}
	}
}

//Inter computes the intersection of 'src' into 'inter'
func Inter(src ...map[string]interface{}) (inter map[string]interface{}) {
	// for each source check if exists the other one
	inter = make(map[string]interface{})
	// peak one set (the first one at random)
	if len(src) == 0 {
		return // empty inter
	}
	scanner := src[0] // a better solution would be to find the smaller one

Scanning: // label to continue the scan loop
	for element := range scanner { // for each element in one of the sets

		//check for each other set if  element is contained
		for _, source := range src {
			if !Contains(source, element) {
				continue Scanning // skipping element from entering inter
			}
		}
		// the element obviously exists in all the sources !
		inter[element] = nil // add it
	}
	return
}

// Equals return true if all sets are equals.
// if there are no sets, return true too.
func Equals(sets ...map[string]interface{}) bool {
	if len(sets) == 0 {
		return true
	}

	size := len(sets[0]) // to start
	for _, set := range sets {
		if len(set) != size {
			return false // set with different size cannot be equal, never
		}
	}

	// now I now, that I have to scan one set (let's select the first)
	for key := range sets[0] {

		// this key must be in all the other one
		for _, set := range sets {
			if _, exists := set[key]; !exists {
				// this one is missing from one of the set ! therefore they are not equals
				return false
			}
		}
	}
	return true
}

//Sub remove all element in 'diff' out of 'src'
func Sub(src, diff map[string]interface{}) {
	for v := range diff {
		delete(src, v)
	}
}

//Peek select one random value from 'set'
func Peek(set map[string]interface{}) string {
	for v := range set {
		return v
	}
	return ""
}

//Pop select one random value from  'set', and remove it
func Pop(set map[string]interface{}) string {
	for v := range set {
		delete(set, v)
		return v
	}
	return ""
}

//New creates a new set from a list of values
func New(val ...string) map[string]interface{} {
	res := make(map[string]interface{}, len(val))
	for _, v := range val {
		res[v] = nil
	}
	return res
}
