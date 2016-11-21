req_mirror
==

An HTTP server that parrots a received request as a response.

Usage
--

```
Usage of req_mirror:
  -port uint
        Port to listen (default 22222)
```

Example
--

### Run server

```
$ req_mirror [--port=22222]
```

### Client

```
$ curl -s -XPOST -F "foo=bar" -F "buz=qux" 127.0.0.1:22222/hoge/fuga | jq .
{
  "Method": "POST",
  "URL": {
    "Scheme": "",
    "Opaque": "",
    "User": null,
    "Host": "",
    "Path": "/hoge/fuga",
    "RawPath": "",
    "ForceQuery": false,
    "RawQuery": "",
    "Fragment": ""
  },
  "Proto": "HTTP/1.1",
  "Header": {
    "Accept": [
      "*/*"
    ],
    "Content-Length": [
      "236"
    ],
    "Content-Type": [
      "multipart/form-data; boundary=------------------------74755673fc9beca6"
    ],
    "Expect": [
      "100-continue"
    ],
    "User-Agent": [
      "curl/7.43.0"
    ]
  },
  "Body": "--------------------------74755673fc9beca6\r\nContent-Disposition: form-data; name=\"foo\"\r\n\r\nbar\r\n--------------------------74755673fc9beca6\r\nContent-Disposition: form-data; name=\"buz\"\r\n\r\nqux\r\n--------------------------74755673fc9beca6--\r\n",
  "TransferEncoding": null,
  "Host": "127.0.0.1:22222"
}
```

Author
--

moznion (<moznion@gmail.com>)

License
--

```
The MIT License (MIT)
Copyright © 2016 moznion, http://moznion.net/ <moznion@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the “Software”), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```

