# webCrawler
A sequential web crawler for visiting all same domain links viable from a specified URL and printing any assets those pages reference. Output's onto std out in JSON format.

Note: URL input only supports absolute URI, a scheme and host must be supplied e.g. "http://www.ccochrane.com"

### Running

##### Running binary
I've supplied a windows executable binary in the releases. Please download and open the terminal in the same directory as the binary and run:

    $ ./webCrawler [your url here]

##### Running from source
First have the [go programming runtime installed](https://golang.org/). Next download the source:

    $ go get github.com/Charlesworth/webCrawler
Navigate to the webCrawler directory and build the executable:

    $ go build
Then on Linux:

    $ ./webCrawler [your url here]
Or Windows:

    $ ./webCrawler.exe [your url here]

### Testing

To run tests, in the root directory of the project please run:

    $ go test ./... -v -cover

Coverage sits at 85.9% for webCrawler and 90.9% for webCrawler/jsonPrinter.

### Limitations
- no .js rendering
- does not support switching between http and https in the crawled domain
- purely sequential by design: each Get takes time, which waterfalls to the next

### Possible future additional features
- visit the root domain of any supplied URL first
- support http/https changes on the same host as the same domain
- parallelize Get requests
