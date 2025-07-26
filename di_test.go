package di

import "testing"

type testStructOneInterface interface {
	Name() string
}
type testStructOne struct{}

func (t *testStructOne) Name() string {
	return "testStructOne"
}

type testStructTwoInterface interface {
	Name() string
}
type testStructTwo struct{}

func (t *testStructTwo) Name() string {
	return "testStructTwo"
}

func TestBinding(t *testing.T) {
	ctr := NewContainer(false, nil)

	ctr.Bind(func() testStructOneInterface {
		return &testStructOne{}
	})

	// testStructOneInterface OK!
	if err := ctr.Invoke(func(s testStructOneInterface) {
		if s.Name() != "testStructOne" {
			t.Fatalf("expecting testStructOne, got %s", s.Name())
		}
	}); err != nil {
		t.Fatal(err)
	}

	// testStructTwoInterface FAIL!
	if err := ctr.Invoke(func(s testStructTwoInterface) {}); err == nil {
		t.Fatal("expecting error, got none")
	}

	ctrScope := ctr.Scope("two")
	ctrScope.Bind(func() testStructTwoInterface {
		return &testStructTwo{}
	})

	// testStructOneInterface OK!
	if err := ctrScope.Invoke(func(s testStructOneInterface) {
		if s.Name() != "testStructOne" {
			t.Fatalf("expecting testStructOne, got %s", s.Name())
		}
	}); err != nil {
		t.Fatal(err)
	}

	// testStructTwoInterface OK!
	if err := ctrScope.Invoke(func(s testStructTwoInterface) {
		if s.Name() != "testStructTwo" {
			t.Fatalf("expecting testStructTwo, got %s", s.Name())
		}
	}); err != nil {
		t.Fatal(err)
	}
}
