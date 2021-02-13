package main

import (
	"fmt"
)

func main() {
	staffArr := make([]*Staff, 8)
	h := &HashTable{StaffArr: staffArr, size: len(staffArr)}
	h.Add(&Staff{id: 1, name: "Peter"})
	h.Add(&Staff{id: 105, name: "Jack"})
	h.Add(&Staff{id: 8, name: "Haris"})
	h.Add(&Staff{id: 65, name: "Jones"})

	h.List()
	staff := h.findById(105)
	fmt.Printf("%+v\n", staff)

	fmt.Println("------------------------", h.Remove(105))
	fmt.Println("------------------------", h.Remove(8))

	h.List()
}

type HashTable struct {
	StaffArr []*Staff
	size     int
}

// 根据 Staff 的 id 获取对应 StaffArr 的索引
func (ht *HashTable) Hash(id int) int {
	return id % ht.size
}

func (ht *HashTable) Add(staff *Staff) {
	indexNo := ht.Hash(staff.id)
	s := ht.StaffArr[indexNo]
	if s == nil {
		ht.StaffArr[indexNo] = staff
	} else {
		s.Add(staff)
	}
}

func (ht *HashTable) List() {
	for i := 0; i < ht.size; i++ {
		ht.StaffArr[i].List(i)
	}
}

func (ht *HashTable) findById(id int) *Staff {
	indexNo := ht.Hash(id)
	return ht.StaffArr[indexNo].findById(id)
}

func (ht *HashTable) Remove(id int) bool {
	indexNo := ht.Hash(id)
	staff := ht.StaffArr[indexNo]
	if staff == nil {
		return false
	}
	if staff.id == id {
		ht.StaffArr[indexNo] = staff.next
		return true
	} else {
		return staff.Remove(id)
	}
}

type Staff struct {
	id   int
	name string
	next *Staff
}

func (staff *Staff) Add(newStaff *Staff) {
	if staff.next == nil {
		staff.next = newStaff
	} else {
		temp := staff
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newStaff
	}
}

func (staff *Staff) List(i int) {
	fmt.Print("第", i, "个链表")
	temp := staff
	for temp != nil {
		fmt.Print("---> id#", temp.id, "name#", temp.name)
		temp = temp.next
	}
	fmt.Println()
}

func (staff *Staff) findById(id int) *Staff {
	if staff == nil {
		return nil
	} else {
		temp := staff
		for {
			if temp == nil {
				temp = nil
				break
			}

			if temp.id == id {
				break
			}
			temp = temp.next
		}
		return temp
	}
}

func (staff *Staff) Remove(id int) bool {
	var remove = false
	temp := staff
	for temp.next != nil {
		if temp.next.id == id {
			temp.next = temp.next.next
			remove = true
			break
		} else {
			temp = temp.next
		}
	}
	return remove
}
