package benchmark

type Config struct {
	URL        string
	Concurrent int
	Limit      int
	Number     int
	Timeout    int
	Headers    []string
}
