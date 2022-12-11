package main

import "testing"

var sample []string = []string{
	`Monkey 1:
  Starting items: 82
  Operation: new = old + 7
  Test: divisible by 13
    If true: throw to monkey 4
    If false: throw to monkey 3`,
	`Monkey 4:
  Starting items: 64, 57, 81, 95, 52, 71, 58
  Operation: new = old * old
  Test: divisible by 11
  If true: throw to monkey 7
  If false: throw to monkey 3`,
}

func TestParseMonkey(t *testing.T) {
	m1 := parseMonkey(sample[0])
	assertMonkeyIsCorrect(t, m1, []int{82}, 5, 12, 12, 3)

	m2 := parseMonkey(sample[1])
	assertMonkeyIsCorrect(t, m2, []int{64, 57, 81, 95, 52, 71, 58}, 5, 25, 11, 7)
}

func assertMonkeyIsCorrect(t *testing.T, m *Monkey, items []int, opin, opres, testin, testres int) {
	if len(m.Items) != len(items) {
		t.Errorf("number of items incorrect: %v should be %v\n", m.Items, items)
	}
	for i := range m.Items {
		if m.Items[i] != items[i] {
			t.Errorf("item has wrong value: %d should be %d\n", m.Items[i], items[i])
		}
	}

	op := m.Operation(opin)
	if op != opres {
		t.Errorf("operation output is wrong: %d should be %d\n", op, opres)
	}

	test := m.Test(testin)
	if test != testres {
		t.Errorf("test output is wrong: %d should be %d\n", test, testres)
	}
}
