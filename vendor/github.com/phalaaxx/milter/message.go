package milter

// Message represents a command sent from milter client
type Message struct {
	Code byte
	Data []byte
}

// Define milter response codes
const (
	acceptMilter   = 'a'
	continueMilter = 'c'
	discardMilter  = 'd'
	rejectMilter   = 'r'
	tempFailMilter = 't'
)
