package gosernit

// NitMsg : function that send a message to a nit.
type NitMsg func(string, string) string

// SendFunc : sends the message to web clients
type SendFunc func(string)

// HandleFunc :
type HandleFunc func(SendFunc, NitMsg)

// DoneFunc : Exits script with an error
type DoneFunc func(string)
