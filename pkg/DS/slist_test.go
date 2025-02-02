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
		t.Fatal("Expected ErrIndexOutOfBound, got nil")
	}

	_, err = list.Get(1000)
	if err == nil {
		t.Fatal("Expected ErrIndexOutOfBound, got nil")
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
	if val1 != 10 {
		t.Fatalf("Expected 10, got %d", val1)
	}

	val2, err := list.Get(1)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if val1 != 10 {
		t.Fatalf("Expected 10, got %d", val2)
	}

	_, err = list.Get(2)
	assert.ErrorIs(t, err, ErrIndexOutOfBound, "Expected error for not used index")

}
