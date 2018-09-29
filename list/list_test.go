package list

import "testing"

func TestListInt(t *testing.T) {
	list := NewList(10)
	ints := []int{2, 3, 4, 5, 6, 7, 8, 9, 1}

	//Add
	for k := range ints {
		result, err := list.Add(ints[k])
		expected := true
		if result != expected {
			t.Errorf("List Add Failed %s", err.Error())
		}
	}

	length := list.Len()
	if 9 != length {
		t.Errorf("List Len failed %d", length)
	}

	//Get
	result, err := list.Get(5)
	if result != true {
		t.Errorf("List Get Failed %s", err.Error())
	}
	result, err = list.Get(50)
	if result != false {
		t.Errorf("List Get Failed %s", err.Error())
	}

	//Modify
	result, err = list.Modify(5, 10)
	if result != true {
		t.Errorf("List Modify Failed %s", err.Error())
	}
	result, err = list.Modify(50, 10)
	if result != false {
		t.Errorf("List Modify Failed %s", err.Error())
	}

	//Delete
	result, err = list.Delete(11)
	if result != false {
		t.Errorf("List delete Failed %s", err.Error())
	}

	result, err = list.Delete(8)
	if result != true {
		t.Errorf("List delete Failed %s", err.Error())
	}
}

func TestListStr(t *testing.T) {
	list := NewList(10)
	strs := []string{"hello", "world", "i", "am", "god", "like", "test", "worlddd"}

	//Add
	for k := range strs {
		result, err := list.Add(strs[k])
		expected := true
		if result != expected {
			t.Errorf("List Add Failed %s", err.Error())
		}
	}

	//Get
	result, err := list.Get("hello")
	if result != true {
		t.Errorf("List Get Failed %s", err.Error())
	}
	result, err = list.Get("fuck")
	if result != false {
		t.Errorf("List Get Failed %s", err.Error())
	}

	//Modify
	result, err = list.Modify("hello", "helle")
	if result != true {
		t.Errorf("List Modify Failed %s", err.Error())
	}
	result, err = list.Modify("id", "xxx")
	if result != false {
		t.Errorf("List Modify Failed %s", err.Error())
	}

	//Delete
	result, err = list.Delete("testxx")
	if result != false {
		t.Errorf("List delete Failed %s", err.Error())
	}

	result, err = list.Delete("test")
	if result != true {
		t.Errorf("List delete Failed %s", err.Error())
	}

	length := list.Len()
	if 7 != length {
		t.Errorf("List Len failed %d", length)
	}

}
