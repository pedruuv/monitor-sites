package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const count = 3
const delay = 5

func main() {
	sayHello()
	for {
		showMenu()

		chosenCommand := readCommand()

		switch chosenCommand {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo logs")
			showLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido")
			os.Exit(-1)
		}
	}

}

func sayHello() {
	name := "Pedro"
	version := 4.0
	fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("0-Sair do Programa")
	fmt.Println("1-Iniciar Monitoramento")
	fmt.Println("2-Exibir Logs")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("O comando escolhido foi", command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	// sites := []string{"https://github.com/pedruuv",
	// 	"https://www.youtube.com/playlist?list=PLfwo-jWzLTWdrDyFeApA2QOg4F2S3XEMi",
	// 	"https://www.codewars.com/dashboard", "https://pt.stackoverflow.com/users/323389/pedrovitorgl"}

	sites := readSitesFromFile()
	for i := 0; i < count; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSites(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func testSites(site string) {
	response, error := http.Get(site)

	if error != nil {
		fmt.Println("Ocorreu um erro", error)
	}
	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "não pôde ser carregado. Status Code:", response.StatusCode)
		registerLog(site, false)
	}
}

func readSitesFromFile() []string {
	var sites []string

	file, error := os.Open("sites.txt")

	if error != nil {
		fmt.Println("Ocorreu um erro", error)
	}

	reader := bufio.NewReader(file)
	for {
		line, error := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		sites = append(sites, line)

		if error == io.EOF {
			break
		}

	}
	file.Close()
	return sites
}

func registerLog(site string, status bool) {
	file, error := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if error != nil {
		fmt.Println(error)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online:" + strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLogs() {
	file, error := ioutil.ReadFile("log.txt")
	if error != nil {
		fmt.Println("error")
	}
	fmt.Println(string(file))
}
