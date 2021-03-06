// Copyright © 2009--2013 The Web.go Authors
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/hraban/web"
)

// Compress data if client supports it

const roothtml = `<html>
<body>compression used for this page: <b id=content></b>
<p>(detection doesn't work in MSIE unfortunately)
<script>
var req = new XMLHttpRequest();
req.open('GET', document.location, false);
req.send(null);
var headers = req.getAllResponseHeaders().toLowerCase();
var compress = req.getResponseHeader("content-encoding");
document.getElementById("content").innerHTML = "" + compress;
</script>`

func root() string {
	return roothtml
}

func main() {
	web.AddWrapper(web.CompressWrapper)
	web.Get("/", root)
	web.Run("127.0.0.1:9999")
}
