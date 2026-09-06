package debug

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/jundaychan/spotlight-mapi/util"
)

// PrintError print error with debug
func PrintError(err error, debug bool) {
	if !debug {
		return
	}
	log.Println("[DEBUG] [ERROR]", err)
}

// PrintStringResponse print string response with debug
func PrintStringResponse(str string, debug bool) {
	if !debug {
		return
	}
	log.Println("[DEBUG] [RESPONSE]", str)
}

// PrintGetRequest print get request with debug
func PrintGetRequest(url string, debug bool) {
	if !debug {
		return
	}
	log.Println("[DEBUG] [API] GET", url)
}

// PrintPostJSONRequest print json request with debug
func PrintPostJSONRequest(url string, body []byte, debug bool) {
	if !debug {
		return
	}
	const format = "[DEBUG] [API] JSON POST %s\n" +
		"http request body:\n%s\n"

	buf := util.NewBufferPool()
	defer util.ReleaseBufferPool(buf)
	json.Indent(buf, body, "", "    ")
	log.Printf(format, url, buf.String())
}

// PrintJSONRequest print json request with debug
func PrintJSONRequest(method string, url string, header http.Header, body []byte, debug bool) {
	if !debug {
		return
	}

	const format = "[DEBUG] [API] JSON %s %s\n" +
		"http request header:\n%s\n" +
		"http request body:\n%s\n"

	buf := util.NewBufferPool()
	defer util.ReleaseBufferPool(buf)
	json.Indent(buf, body, "", "\t")
	headers := make([]string, 0, len(header))
	for k := range header {
		headers = append(headers, util.StringsJoin(k, ": ", header.Get(k)))
	}

	log.Printf(format, method, url, strings.Join(headers, "\n"), buf.String())
}

// PrintPostMultipartRequest print multipart/form-data post request with debug
func PrintPostMultipartRequest(url string, mp map[string]string, debug bool) {
	if !debug {
		return
	}
	const format = "[DEBUG] [API] multipart/form-data POST %s\n" +
		"http request body:\n%s\n"

	bs, _ := json.MarshalIndent(mp, "", "\t")
	log.Printf(format, url, bs)
}

// DecodeJSONHttpResponse decode json response with debug
// DecodeJSONHttpResponse 解 JSON 到 v，并**始终**把原始 body 拷贝一份返回。
//
// 以前只在 debug 开着时才返回 body，关着就回 nil——于是 client 想把「-1 且 msg 为空」
// 这类错误的原始响应带进错误串时，拿到的永远是空的，排查者对着 body= 什么都看不到。
// 另外旧实现直接返回 buf.Bytes()，而 buf 在 defer 里归还给池子：返回的切片别名了
// 一块马上会被复用的内存，调用方读到的可能是别的请求的字节。这里一律 copy 出来。
func DecodeJSONHttpResponse(r io.Reader, v interface{}, debug bool) ([]byte, error) {
	buf := util.NewBufferPool()
	defer util.ReleaseBufferPool(buf)
	tee := io.TeeReader(r, buf)
	decErr := json.NewDecoder(tee).Decode(v)
	// Decode 只读到第一个完整 JSON 值就停，把剩余字节也吞进 buf，body 才是完整的
	_, _ = io.Copy(io.Discard, tee)
	bs := append([]byte(nil), buf.Bytes()...)
	if decErr != nil {
		return bs, decErr
	}
	if debug {
		debugBuf := util.NewBufferPool()
		defer util.ReleaseBufferPool(debugBuf)
		if err := json.Indent(debugBuf, bs, "", "\t"); err == nil {
			log.Println(util.StringsJoin("[DEBUG] [API] http response body:\n", debugBuf.String()))
		}
	}
	return bs, nil
}
