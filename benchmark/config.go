package benchmark

type Config struct {
	URL         string
	Concurrent  int
	Limit       int
	Number      int
	Method      string
	Timeout     int
	Headers     []string
	RequestBody []byte
}
