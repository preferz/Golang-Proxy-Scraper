package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetSock4() {

	r, err := http.Get("https://api.openproxylist.xyz/socks4.txt")
	if err != nil {
		panic(err)
	}
	prox, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Writing To File...")

	proxies := strings.TrimSuffix(string(prox), "\n")
	file, err := os.OpenFile("Proxy/socks4.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = fmt.Fprint(file, proxies)
	if err != nil {
		panic(err)
	}
	fmt.Println("Scraped! | Saved to socks4.txt")

}

func GetSock5() {

	r, err := http.Get("https://api.openproxylist.xyz/socks5.txt")
	if err != nil {
		panic(err)
	}
	prox, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Writing To File...")

	proxies := strings.TrimSuffix(string(prox), "\n")
	file, err := os.OpenFile("Proxy/socks5.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = fmt.Fprint(file, proxies)
	if err != nil {
		panic(err)
	}
	fmt.Println("Scraped! | Saved to socks5.txt")

}

func GetHttp() {

	r, err := http.Get("https://api.openproxylist.xyz/http.txt")
	if err != nil {
		panic(err)
	}
	prox, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Writing To File...")

	proxies := strings.TrimSuffix(string(prox), "\n")
	file, err := os.OpenFile("Proxy/http.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = fmt.Fprint(file, proxies)
	if err != nil {
		panic(err)
	}
	fmt.Println("Scraped! | Saved to http.txt")

}

func main() {

	var banner string = `
	[1] Socks4
	[2] Socks5
	[3] HTTP
	`

	fmt.Println(banner)

	var choice string
	fmt.Printf("Enter Choice: ")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		GetSock4()
	case "2":
		GetSock5()
	case "3":
		GetHttp()
	default:
		fmt.Println("Invalid Choice!")

	}

}
