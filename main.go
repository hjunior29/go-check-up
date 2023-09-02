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

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		command := leComando()

		switch command {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	name := "Helder"
	version := 1.1
	fmt.Println("Olá, Sr.", name)
	fmt.Println("Este programa está na versão:", version)
}

func leComando() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("O comando escolhido foi:", command)
	fmt.Println("")

	return command
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// sites := []string{"https://www.alura.com.br","https://www.google.com","https://www.youtube.com"}

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println(site, "Ok")
		registraLog(site, true)
	} else {
		fmt.Println(site, "Down", "\nStatus:", resp.StatusCode)
		registraLog(site, true)
	}
}

func leSitesDoArquivo() []string {
	sites := []string{}

	arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println(err)
	}

	leitor := bufio.NewReader(arquivo)
	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		if err == io.EOF {
			break
		}
		sites = append(sites, linha)
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:5 ") + site + "\n                    └──Status: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {
	fmt.Println("Exibindo logs...")
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}
