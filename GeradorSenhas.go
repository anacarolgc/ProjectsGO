package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func gerarSenha(comprimento int, usarMaiusculas, usarNumeros, usarSimbolos bool) string {
	var caracteres []rune
	caracteres = append(caracteres, 'a', 'z')

	if usarMaiusculas {
		for i := 'A'; i <= 'Z'; i++ {
			caracteres = append(caracteres, i)
		}
	}

	if usarNumeros {
		for i := '0'; i <= '9'; i++ {
			caracteres = append(caracteres, i)
		}
	}

	if usarSimbolos {
		simbolos := []rune{'!', '"', '#', '$', '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~'}
		caracteres = append(caracteres, simbolos...)
	}

	rand.Seed(time.Now().UnixNano())
	var senha []rune
	for i := 0; i < comprimento; i++ {
		senha = append(senha, caracteres[rand.Intn(len(caracteres))])
	}

	return string(senha)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o comprimento da senha: ")
	comprimentoStr, _ := reader.ReadString('\n')
	comprimentoStr = strings.TrimSpace(comprimentoStr)

	var comprimento int
	fmt.Sscanf(comprimentoStr, "%d", &comprimento)

	fmt.Print("Deseja incluir letras maiúsculas? (s/n): ")
	incluirMaiusculas, _ := reader.ReadString('\n')
	incluirMaiusculas = strings.TrimSpace(incluirMaiusculas)

	fmt.Print("Deseja incluir números? (s/n): ")
	incluirNumeros, _ := reader.ReadString('\n')
	incluirNumeros = strings.TrimSpace(incluirNumeros)

	fmt.Print("Deseja incluir símbolos? (s/n): ")
	incluirSimbolos, _ := reader.ReadString('\n')
	incluirSimbolos = strings.TrimSpace(incluirSimbolos)

	usarMaiusculas := incluirMaiusculas == "s"
	usarNumeros := incluirNumeros == "s"
	usarSimbolos := incluirSimbolos == "s"

	senha := gerarSenha(comprimento, usarMaiusculas, usarNumeros, usarSimbolos)
	fmt.Printf("\nSua senha gerada é: %s\n", senha)
}
