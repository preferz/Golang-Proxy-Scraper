package main

import( 
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"os"
)

func main() {

	fmt.Println("Starting...")

	r, err := http.Get("https://api.proxyscrape.com/?request=displayproxies&proxytype=http&timeout=10000&country=all&anonymity=all&ssl=no")
	if err != nil {
		fmt.Println(err)
	}
	prox, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Writing To File...")

	proxies := strings.TrimSuffix(string(prox), "\n")
	file, err := os.OpenFile("Proxy/proxies.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("File Not found!")
		os.Exit(0)
	}
	defer file.Close()

	_, err = fmt.Fprint(file, proxies)
	if err != nil {
		file.Close()
		fmt.Println(err)
	}
	fmt.Println("Scraped! | Saved to proxies.txt")

}