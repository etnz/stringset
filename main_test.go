package stringset

import "testing"

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func TestClone(t *testing.T) {
	set := New("a", "b", "c")
	clone := Clone(set)

	if !Equals(set, clone) {
		t.Errorf("Invalid Clone: clones are expected to be equals %v, %v", set, clone)
	}
	set["d"] = nil // append some stuff
	if Equals(set, clone) {
		t.Errorf("Invalid Clone: clones are expected to be independent %v, %v", set, clone)
	}

}

func TestSort(t *testing.T) {
	set := New("a", "c", "b")
	keys := Sort(set)
	if len(keys) != 3 || keys[0] != "a" || keys[1] != "b" || keys[2] != "c" {
		t.Errorf("Invalid Sort: %v", keys)
	}
}

func TestContains(t *testing.T) {
	set := New(alphabet...)
	if !Contains(set, "a") {
		t.Errorf("invalid Contains: expect to contain 'a'")
	}
	if Contains(set, "&") {
		t.Errorf("invalid Contains: expect to not contain '&'")
	}
}

func TestContainsAny(t *testing.T) {
	set := New(alphabet...)
	if !ContainsAny(set, "a", "2", " ") {
		t.Errorf("invalid ContainsAny: expect to containAny of 'a' '2', ' '")
	}
	if ContainsAny(set, "&", " ", "2") {
		t.Errorf("invalid Contains: expect to not contain any  '&' ' ', '2' ")
	}
}

func TestContainsAll(t *testing.T) {
	set := New(alphabet...)
	if !ContainsAll(set, "a", "b", "c") {
		t.Errorf("invalid ContainsAll: expect to contain all of 'a' 'b', 'c'")
	}
	if ContainsAll(set, "a", "b", "2") {
		t.Errorf("invalid Contains: expect to not contain all  'a' 'b', '2' ")
	}
}

func TestUnion(t *testing.T) {
	set1 := New("a", "b", "c")
	set2 := New("a", "b", "d")
	set3 := New("a", "d", "c")
	u := Union(set1, set2, set3)

	x := New("a", "b", "c", "d")
	if !Equals(x, u) {
		t.Errorf("invalid union: got %v expecting %v", u, x)
	}

}

func TestAppend(t *testing.T) {
	u := New("z", "w")
	set1 := New("a", "b", "c")
	set2 := New("a", "b", "d")
	set3 := New("a", "d", "c")
	Append(u, set1, set2, set3)
	x := New("a", "b", "c", "d", "w", "z")
	if !Equals(x, u) {
		t.Errorf("invalid append: got %v expecting %v", u, x)
	}
}

func TestInter(t *testing.T) {
	set1 := New("a", "b", "c", "z")
	set2 := New("a", "b", "d", "z")
	set3 := New("a", "d", "c", "z")

	u := Inter(set1, set2, set3)
	x := New("a", "z")
	if !Equals(x, u) {
		t.Errorf("invalid inter: got %v expecting %v", u, x)
	}
}

func TestEquals(t *testing.T) {
	set1 := New("a", "b", "c", "z")
	set2 := New("a", "b", "d", "z")

	if Equals(set1, set2) {
		t.Errorf("invalid Equals got %v and %v expecting differents", set1, set2)
	}
	if !Equals(set1, set1) {
		t.Errorf("invalid Equals got %v and %v expecting equals", set1, set1)
	}
}

func TestSub(t *testing.T) {
	set1 := New(alphabet...)
	set2 := New(alphabet[5:]...)
	x := New(alphabet[:5]...)

	Sub(set1, set2)
	if !Equals(x, set1) {
		t.Errorf("invalid Sub: got %v, expecting %v", set1, x)
	}
}

func TestPeek(t *testing.T) {
	set1 := New(alphabet...)
	v := Peek(set1)

	if !Contains(set1, v) {
		t.Errorf("invalid Peek: got %v, not contained in %v", v, set1)
	}
	if len(set1) != len(alphabet) {
		t.Errorf("invalid Peek: src len  %v altered. Expecting %v", len(set1), len(alphabet))
	}
}

func TestPop(t *testing.T) {
	set1 := New(alphabet...)
	v := Pop(set1)

	if Contains(set1, v) {
		t.Errorf("invalid Pop: got %v, contained in %v", v, set1)
	}
	if len(set1) == len(alphabet) {
		t.Errorf("invalid Pop: src len  %v not altered. Expecting %v", len(set1), len(alphabet)-1)
	}
}

func TestNew(t *testing.T) {
	set1 := New("a", "b", "a")
	if len(set1) != 2 {
		t.Errorf("invalid New: invalid len, got %v Expecting %v", len(set1), 2)
	}
}
