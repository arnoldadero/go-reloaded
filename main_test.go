package main

import (
	"testing"
)

func TestProcessText1(t *testing.T) {
	input := "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?"
	expected := "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"

	actual := ProcessText(input)

	if actual != expected {
		t.Errorf("Result: %s\n", actual)
		t.Errorf("Expect: %s\n", expected)
	}
}

func TestProcessText2(t *testing.T) {
	input := "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure"
	expected := "I have to pack 5 outfits. Packed 26 just to be sure"

	actual := ProcessText(input)

	if actual != expected {
		t.Errorf("Result: %s\n", actual)
		t.Errorf("Expect: %s\n", expected)
	}
}

func TestProcessText3(t *testing.T) {
	input := "Don not be sad ,because sad backwards is das . And das not good"
	expected := "Don not be sad, because sad backwards is das. And das not good"

	actual := ProcessText(input)

	if actual != expected {
		t.Errorf("Result: %s\n", actual)
		t.Errorf("Expect: %s\n", expected)
	}
}

func TestProcessText4(t *testing.T) {
	input := "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '"
	expected := "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"

	actual := ProcessText(input)

	if actual != expected {
		t.Errorf("Result: %s\n", actual)
		t.Errorf("Expect: %s\n", expected)
	}
}
