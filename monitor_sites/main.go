package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Site struct {
	Rank               int    `json:"rank"`
	RootDomain         string `json:"rootDomain"`
	LinkingRootDomains int    `json:"linkingRootDomains"`
	DomainAuthority    int    `json:"domainAuthority"`
}

func main() {

	showMenu()

	cmd := commandInput()

	switch cmd {
	case 0:
		os.Exit(0)
	case 1:
		monitoring()
	case 2:
		fmt.Println("two")
	default:
		os.Exit(-1)
	}
}

func commandInput() int {
	var cmd int

	fmt.Scan(&cmd)
	return cmd
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("0- Sair do Programa")
}

func monitoring() {
	sites := jsonLoad()
	responseChannel := make(chan string)
	timeStamp := time.Now()
	file, err := os.Create(fmt.Sprintf("result%s.txt", timeStamp.String()))

	defer file.Close()

	buffer := bufio.NewWriter(file)

	if err != nil {
		log.Fatal("[Error] ", err.Error())
	}

	for i := 0; i < len(sites); i++ {
		go callHttp(sites[i], responseChannel)
	}

	for i := 0; i < len(sites); i++ {
		result := <-responseChannel
		position := i + 1
		writeFile(buffer, result, position)
		log.Printf("%d - %s", position, result)
	}
	close(responseChannel)
}

func jsonLoad() []Site {
	jsonFile, err := os.Open("top-sites.json")

	if err != nil {
		log.Fatal("[Error] ", err.Error())
	}

	log.Printf("[Info] Successfully Opened file: %s ", jsonFile.Name())

	defer jsonFile.Close()

	byteJson, _ := io.ReadAll(jsonFile)

	var sites []Site

	json.Unmarshal(byteJson, &sites)
	return sites
}

func callHttp(site Site, responseChannel chan<- string) {

	var feedback string

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", "https://"+site.RootDomain, nil)
	if err != nil {
		responseChannel <- fmt.Sprintf("[Error] Problema ao encontrar [URL: %s] [Msg: %s]", site.RootDomain, err.Error())
		return
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		responseChannel <- fmt.Sprintf("[Error] Problema ao encontrar [URL: %s] [Msg: %s]", site.RootDomain, err.Error())
		return
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		feedback = fmt.Sprintf("[Info] [URL: %s] [StatusCode: %d]", site.RootDomain, response.StatusCode)
	} else {
		feedback = fmt.Sprintf("[Warning] [URL: %s] [StatusCode: %d]", site.RootDomain, response.StatusCode)
	}

	responseChannel <- feedback
}

func writeFile(buffer *bufio.Writer, result string, position int) {
	_, err := buffer.WriteString(fmt.Sprintf("%d - %s", position, result) + "\n")
	if err != nil {
		log.Fatal(err)
	}
}
