package main
import ( 
	"fmt"
	"net/url"
)

func main() {
	urlStr := "https://www.google.com/search?q=golang&lang=en"

	parsedURL, err := url.Parse(urlStr)


	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Schema:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Raw Query:", parsedURL.RawQuery)

	queryParams := parsedURL.Query()

	fmt.Println("Search Query:", queryParams.Get("q"))
	fmt.Println("Language:", queryParams.Get("lang"))
}