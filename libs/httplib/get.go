package httplib

type GetOptions struct {
	timeout int
	header  map[string]string
	cookie  string
}

func HttpGet(url string, options GetOptions) {

}
