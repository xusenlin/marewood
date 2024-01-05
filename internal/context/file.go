package context

import "net/http"

func (r *Context) SendFile(name string, file []byte) {
	//name = example.txt
	r.c.Header("Content-Disposition", "attachment; filename="+name)
	r.c.Data(http.StatusOK, "application/octet-stream", file)
}
