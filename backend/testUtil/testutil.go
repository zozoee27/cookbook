package testutil

import (
	"github.com/zozoee27/cookbook/backend/entity"
	"testing"
)

func CompareString(t *testing.T, actual, expected, err, where string) bool {
	isSame := true
	if actual != expected {
		t.Errorf("%s - %s: \nActual String:[%s]\nExpected String:[%s]\n", err, where, actual, expected)
		isSame = false
	}
	return isSame
}

func CompareInt(t *testing.T, actual, expected int, err, where string) bool {
	isSame := true
	if actual != expected {
		t.Errorf("%s - %s: \nActual Int:[%d]\nExpected Int:[%d]\n", err, where, actual, expected)
		isSame = false
	}
	return isSame
}

func CompareError(t *testing.T, actual, expected error, err, where string) bool {
	isSame := true
	if (actual == nil && expected != nil) || (expected == nil && actual != nil) {
		t.Errorf("%s - %s: \nActual Error:[%e]\nExpected Error:[%e]\n", err, where, actual, expected)
		isSame = false
	} else if actual != nil && expected != nil {
		isSame = CompareString(t, actual.Error(), expected.Error(), err, where)
	}
	return isSame
}

func CompareByte(t *testing.T, actual, expected byte, err, where string) bool {
	isSame := true
	if actual != expected {
		t.Errorf("%s - %s: \nActual Byte:[% x]\nExpected Byte:[% x]\n", err, where, actual, expected)
		isSame = false
	}

	return isSame
}

func CompareByteArray(t *testing.T, actual, expected []byte, err, where string) bool {
	isSame := true
	isSame = CompareInt(t, len(actual), len(expected), err, where)

	if CompareInt(t, len(actual), len(expected), err, where) {
		for i, a := range actual {
			if !CompareByte(t, a, expected[i], err, where) {
				t.Errorf("%s - %s: \nIndex:[%d]\nActual Byte Array:[% x]\nExpected Byte Array:[% x]\n", err, where, i, actual, expected)
				isSame = false
				break
			}
		}
	} else {
		isSame = false
	}
	return isSame
}

func CompareUserEntity(t *testing.T, actual, expected *entity.User, err, where string) bool {
	isSame := true
	if actual != nil && expected != nil && *actual != *expected {
		t.Errorf("%s - %s: \nActual User:[%s]\nExpected User:[%s]\n", err, where, actual.PrettyString(), expected.PrettyString())
		isSame = false
	} else if actual == nil && expected != nil {
		t.Errorf("%s - %s: \nActual User:[nil]\nExpected User:[%s]\n", err, where, expected.PrettyString())
		isSame = false
	} else if actual != nil && expected == nil {
		t.Errorf("%s - %s: \nActual User:[%s]\nExpected User:[nil]\n", err, where, actual.PrettyString())
		isSame = false
	}
	return isSame
}
