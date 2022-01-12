package lib

import "fmt"

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Len() uint {
	if l.Head == nil {
		return 0
	}

	var length uint
	var temp *Node = l.Head

	for temp != nil {
		length++
		temp = temp.Next
	}

	return length
}

func (l LinkedList) Display() {
	for l.Head != nil {
		fmt.Printf("%v -> ", l.Head.Data)
		l.Head = l.Head.Next
	}

	fmt.Printf("\n")
}

func (l *LinkedList) PushBack(new *Node) {
	// time complexity = O(n)

	if l.Head == nil {
		l.Head = new

		return
	}

	// karena new node kita taruh di akhir, maka next dari new node adalah nil
	new.Next = nil

	var last *Node = l.Head

	for last.Next != nil {
		last = last.Next
	}

	// ubah pointer next dari last node ke new node
	last.Next = new
}

func (l *LinkedList) Push(new *Node) {
	// Time Complexity = O(1)
	// Method ini untuk insert node baru dari depan
	// l.Head adalah kepala semula, kita pindahkan value l.Head ke n.Next
	new.Next = l.Head

	// l.Head akan digantikan oleh value n yang baru
	l.Head = new
}

func (l *LinkedList) PushAfter(prev *Node, new *Node) {
	// Method ini untuk taruh node in certain place
	// Time Complexity = O(1)

	// check if the given previous node is null
	if prev == nil {
		fmt.Println("The given previous node cannot be empty")
		return
	}

	// taruh next dari previous node ke next dari new node
	new.Next = prev.Next

	// ubah next dari previous node menjadi new node
	prev.Next = new
}

func (l *LinkedList) Pop() int {
	if l.Head == nil {
		return -1
	} else if l.Head.Next == nil {
		l.Head = nil
		return -1
	}

	var last *Node = l.Head

	for last.Next.Next != nil {
		last = last.Next
	}

	output := last.Next
	last.Next = nil

	return output.Data
}

func (l *LinkedList) Shift() int {
	// time complexity = O(1)

	// cek apakah head memiliki value
	if l.Head == nil {
		return -1
	}

	// taruh data dari head untuk dijadikan return value
	output := l.Head.Data

	// taruh node next dari head ke variabel temp agar tidak hilang ketika head dihapus
	var temp *Node = l.Head.Next

	// hapus head
	l.Head = nil

	// letakkan node next tadi ke posisi head
	l.Head = temp

	return output
}

func (l *LinkedList) Delete(pos int) {
	if l.Head == nil {
		return
	}

	// taruh node head ke variabel temp
	var temp *Node = l.Head

	// jika yang dihapus adalah head,
	// maka cukup ganti head menjadi node selanjutnya
	if pos == 0 {
		l.Head = temp.Next

		return
	}

	// temp di sini adalah previous node
	// cari previous node dengan menggunakan for-loop
	for i := 0; temp != nil && i < pos-1; i++ {
		temp = temp.Next
	}

	// jika setelah dicari temp sama dengan nil, atau temp.Next sama dengan nil
	// maka artinya position offset, diluar panjang linkedlist
	if temp == nil || temp.Next == nil {
		return
	}

	// tampung next node ke variabel next
	var next *Node = temp.Next.Next

	// hubungkan previous node dengan next node
	// otomatis node di antara previous dan next akan hilang
	// karena prev dan next langsung disambungkan
	temp.Next = next
}

func (l *LinkedList) Get(pos int) (int, bool) {
	if l.Head == nil {
		return 0, false
	}

	var temp *Node = l.Head

	for i := 0; i < pos && temp != nil; i++ {
		temp = temp.Next
	}

	if temp == nil {
		return 0, false
	}

	return temp.Data, true
}
