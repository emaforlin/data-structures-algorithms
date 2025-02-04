package DS

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSList_NewSList(t *testing.T) {
	list := NewSList[any]()
	if list == nil {
		t.Fatal("NewSList() returned nil")
	}

	if list.Len() != 0 {
		t.Errorf("Expected size %d, got %v", 0, list.Len())
	}
}

func TestSList_Add(t *testing.T) {
	var list = NewSList[int]()

	// Add some values
	list.Add(10)
	list.Add(20)

	// Check if the length is correct
	if list.Len() != 2 {
		t.Errorf("Expected length 2, got %d", list.Len())
	}

	// Try getting valid indexes
	for i := 0; i < list.Len(); i++ {
		_, err := list.Get(i)
		if err != nil {
			t.Errorf("Expected error nil, got: %v", err)
		}
	}

	// Try getting invalid indexes
	_, err := list.Get(-1)
	if err == nil {
		t.Error("Expected ErrIndexOutOfBound, got nil")
	}

	_, err = list.Get(1000)
	if err == nil {
		t.Error("Expected ErrIndexOutOfBound, got nil")
	}

}

func TestSList_Get(t *testing.T) {
	var list = NewSList[int]()

	_, err := list.Get(0)
	assert.ErrorIs(t, err, ErrIndexOutOfBound, "Expected error for empty list")

	list.Add(10)
	list.Add(20)

	val1, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if *val1 != 10 {
		t.Errorf("Expected 10, got %d", *val1)
	}

	val2, err := list.Get(1)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if *val1 != 10 {
		t.Errorf("Expected 10, got %d", *val2)
	}

	_, err = list.Get(2)
	assert.ErrorIs(t, err, ErrIndexOutOfBound, "Expected error for not used index")

}

func TestSList_Insert(t *testing.T) {
	var list = NewSList[string]()
	list.Add("1")
	list.Add("3")

	err := list.Insert(-1, "invalid insertion")
	assert.ErrorIs(t, err, ErrIndexOutOfBound, "Expected error for invalid index")

	err = list.Insert(0, "prepend this string")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	err = list.Insert(2, "insert this between '1' and '3'")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	err = list.Insert(list.Len(), "append this string")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	val1, err := list.Get(0)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	if *val1 != "prepend this string" {
		t.Errorf("Unexpected string: %s", *val1)
	}

	val2, err := list.Get(1)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	if *val2 != "1" {
		t.Errorf("Unexpected string: %s", *val2)
	}

	val3, err := list.Get(2)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	if *val3 != "insert this between '1' and '3'" {
		t.Errorf("Unexpected string: %s", *val3)
	}

	val4, err := list.Get(3)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	if *val4 != "3" {
		t.Errorf("Unexpected string: %s", *val4)
	}
	list.Print()

}

func TestSList_Find(t *testing.T) {
	list := SListImpl[int]{}
	val1, val2, val3 := 10, 20, 30
	list.head = &cell[int]{data: &val1, next: &cell[int]{data: &val2, next: &cell[int]{data: &val3, next: nil}}}
	list.size = 3

	tests := []struct {
		value    int
		expected int
	}{
		{10, 0},
		{20, 1},
		{30, 2},
		{40, -1},
	}

	for _, test := range tests {
		result, _ := list.Find(test.value)
		if (result == nil && test.expected != -1) || (result != nil && *result != test.value) {
			t.Errorf("Find(%d) = %v; want %d", test.value, result, test.expected)
		}
	}
}

func TestSList_RemoveAt(t *testing.T) {
	list := NewSList[int]()
	list.Add(10)
	list.Add(20)
	list.Add(30)
	tests := []struct {
		index        int
		expectedSize int
		error        error
	}{
		{1, 2, nil},                // Removing second element
		{5, 2, ErrIndexOutOfBound}, // Out of bounds index
		{0, 1, nil},
	}

	for _, test := range tests {
		err := list.RemoveAt(test.index)
		if err != test.error {
			t.Errorf("RemoveAt(%d) returned error %v; want %v", test.index, err, test.error)
		}
		if list.Len() != test.expectedSize {
			t.Errorf("After RemoveAt(%d), list size = %d; want %d", test.index, list.Len(), test.expectedSize)
		}
	}
}

func TestSList_Remove(t *testing.T) {
	list := NewSList[int]()
	list.Add(10)
	list.Add(20)
	list.Add(30)
	tests := []struct {
		value        int
		expectedSize int
		error        error
	}{
		{10, 2, nil},               // Removing second element
		{560, 2, ErrValueNotFound}, // Out of bounds index
		{30, 1, nil},
	}

	for _, test := range tests {
		err := list.Remove(test.value)
		if err != test.error {
			t.Errorf("Remove(%d) returned error %v; want %v", test.value, err, test.error)
		}
		if list.Len() != test.expectedSize {
			t.Errorf("After Remove(%d), list size = %d; want %d", test.value, list.Len(), test.expectedSize)
		}
	}
	list.Print()
}
