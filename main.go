package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var proxyType string
	var proxyName string

	fmt.Print("[*] Type of proxies you would like to gather >> ")
	fmt.Scanln(&proxyType)
	fmt.Print("[*] Name of proxies file where you would like to save result >> ")
	fmt.Scanln(&proxyName)

	switch proxyType {
	case "socks4":
		resp, err := http.Get("https://api.proxyscrape.com/v2/?request=getproxies&protocol=socks4&timeout=10000&country=all")
		if err != nil {
			fmt.Println("[X] Error occurred while trying to make GET request")
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("[X] Failed to read body of GET request")
			return
		}
		ioutil.WriteFile(proxyName+".txt", body, 0644)
	case "socks5":
		resp, err := http.Get("https://api.proxyscrape.com/v2/?request=getproxies&protocol=socks5&timeout=10000&country=all")
		if err != nil {
			fmt.Println("[X] Error occurred while trying to make GET request")
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("[X] Failed to read body of GET request")
			return
		}

		ioutil.WriteFile(proxyName+".txt", body, 0644)
	case "http":
		resp, err := http.Get("https://api.proxyscrape.com/v2/?request=getproxies&protocol=http&timeout=10000&country=all")
		if err != nil {
			fmt.Println("[X] Error occurred while trying to make GET request")
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("[X] Failed to read body of GET request")
			return
		}

		ioutil.WriteFile(proxyName+".txt", body, 0644)
	}
}
