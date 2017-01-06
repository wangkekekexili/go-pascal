package go_pascal

import "testing"

func TestInterpreter(t *testing.T) {
	program := `
BEGIN
	BEGIN
		a := 2;
		b := a * 2;
		c := a + b
	END;
	x := 10
END.
`
	i := newInterpreter(program)
	if err := i.walk(); err != nil {
		t.Fatal(err)
	}

	expect := map[string]float64{
		"a": 2,
		"b": 4,
		"c": 6,
		"x": 10,
	}
	for id, value := range expect {
		if got, ok := i.globalScope[id]; !ok {
			t.Fatalf("%v is expected to be in the symbol table", id)
		} else if got != value {
			t.Fatalf("expected to get %v for %v; got %v", value, id, got)
		}
	}
}
