package zutil

var (
	seq TSequence = 0
)

type USequence struct {
	id TSequence
}

func NewSequence() *USequence {
	return &USequence{0}
}
func (sequence *USequence) Reset() {
	sequence.id = 0
}
func (sequence *USequence) Inc() TSequence {
	sequence.id++
	return sequence.id
}
func IncSequence() TSequence {
	seq++
	return seq
}
