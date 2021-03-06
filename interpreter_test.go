package go_pascal

import "testing"

func TestInterpreter(t *testing.T) {
	program := `
PROGRAM test;
BEGIN
	begin
		a := 2;
		b := A * 2;
		c := a + B
	end;
	x := 10;
	_num := 5
END.
`
	i := newInterpreter(program)
	if err := i.walk(); err != nil {
		t.Fatal(err)
	}

	expect := map[string]interface{}{
		"a":    2,
		"b":    4,
		"c":    6,
		"x":    10,
		"_num": 5,
	}
	for id, value := range expect {
		if got, ok := i.globalScope[id]; !ok {
			t.Fatalf("%v is expected to be in the symbol table", id)
		} else if got != value {
			t.Fatalf("expected to get %v for %v; got %v", value, id, got)
		}
	}
}
