package data

type InterlocutorQueue interface {
	Push(user *Interlocutor)
	Pop() *Interlocutor
}

type InterlocutorNode struct {
	Val  *Interlocutor
	Next *InterlocutorNode
}

type interlocutorQueue struct {
	first *InterlocutorNode
	last  *InterlocutorNode
	ln    int
}

func NewInterlocutorQueue() InterlocutorQueue {
	return &interlocutorQueue{
		first: nil,
		last:  nil,
		ln:    0,
	}
}

func (q *interlocutorQueue) Push(user *Interlocutor) {
	q.ln++

	if q.ln == 1 {
		q.first = &InterlocutorNode{
			Val:  user,
			Next: nil,
		}

		q.last = q.first
		return
	}

	q.last.Next = &InterlocutorNode{
		Val:  user,
		Next: nil,
	}

	q.last = q.last.Next
}

func (q *interlocutorQueue) Pop() *Interlocutor {
	if q.ln == 0 {
		return nil
	}

	q.ln--

	user := q.first.Val
	q.first = q.first.Next

	return user
}
