package rester

//Payload is the essential interface for result check.
type Payload interface {
    Equal(v interface{})
}
