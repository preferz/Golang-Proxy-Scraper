package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
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

func CheckProxies() {

	defer wg.Done()

	fmt.Printf("Enter file path: ")
	var path string
	fmt.Scanln(&path)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("File does not exist!")
		return
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var proxy string
	var proxyList []string

	fmt.Println("Checking Proxies...")

	for {
		_, err := fmt.Fscanln(file, &proxy)
		if err != nil {
			break
		}
		proxyList = append(proxyList, proxy)
	}

	var GoodProxies []string
	var BadProxies []string

	for _, proxy := range proxyList {

		proxyURL, _ := url.Parse(proxy)
		Client := &http.Client{
			Timeout: 5 * time.Second,
			Transport: &http.Transport{
				DisableKeepAlives: true,
				Proxy:             http.ProxyURL(proxyURL),
			},
		}
		resp, err := Client.Get("https://api.ipify.org?format=json")
		if err != nil {
			fmt.Printf("%s is not working\n", proxy)
			BadProxies = append(BadProxies, proxy)
			continue
		}
		defer resp.Body.Close()
		fmt.Printf("%s is working\n", proxy)
		GoodProxies = append(GoodProxies, proxy)
	}

	fmt.Println("Writing To File...")

	file, err = os.OpenFile("Proxy/good.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = fmt.Fprint(file, GoodProxies)
	if err != nil {
		panic(err)
	}
	file, err = os.OpenFile("Proxy/bad.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprint(file, BadProxies)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scraped! | Saved to good.txt & bad.txt")

}

func main() {

	var banner string = `
	[1] Socks4
	[2] Socks5
	[3] HTTP
	[4] Check Proxy
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
	case "4":
		wg.Add(1)
		CheckProxies()
	default:
		fmt.Println("Invalid Choice!")

	}

}
