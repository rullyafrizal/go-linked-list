package lib

// struct untuk tiap tiap node
// berisi data yang di-contain oleh node itu sendiri dan pointer data node selanjutnya
// membayangkan pointer next sebagai embedded akan lebih make sense daripada
// membayangkan pointer next seperti bersebelahan
type Node struct {
	Data int
	Next *Node
}