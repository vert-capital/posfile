package pad

import (
	"errors"
	"strings"
)

// Right é uma função que preenche uma string à direita com um caractere de preenchimento especificado até atingir um tamanho desejado.
func Right(str string, size int, pad string) (string, error) {
	// Verifica se o tamanho desejado é negativo. Se for, retorna um erro.
	if size < 0 {
		return "", errors.New("size must be non-negative")
	}

	// Verifica se a string de preenchimento está vazia. Se estiver, retorna um erro.
	if pad == "" {
		return "", errors.New("pad must not be empty")
	}

	// Converte a string de entrada em uma fatia de runas para poder manipular os caracteres individualmente.
	content := []rune(str)

	// Se a string de entrada for maior que o tamanho desejado, corta a string para o tamanho desejado.
	if len(content) > size {
		content = content[:size]
	}

	// Retorna a string de entrada, preenchida à direita com o caractere de preenchimento até atingir o tamanho desejado.
	return string(content) + strings.Repeat(pad, size-len(content)), nil
}

// Left é uma função que preenche uma string à esquerda com um caractere de preenchimento especificado até atingir um tamanho desejado.
func Left(str string, size int, pad string) (string, error) {
	// Verifica se o tamanho desejado é negativo. Se for, retorna um erro.
	if size < 0 {
		return "", errors.New("size must be non-negative")
	}

	// Verifica se a string de preenchimento está vazia. Se estiver, retorna um erro.
	if pad == "" {
		return "", errors.New("pad must not be empty")
	}

	// Converte a string de entrada em uma fatia de runas para poder manipular os caracteres individualmente.
	content := []rune(str)

	// Se a string de entrada for maior que o tamanho desejado, corta a string para o tamanho desejado.
	if len(content) > size {
		content = content[:size]
	}

	// Retorna a string de entrada, preenchida à esquerda com o caractere de preenchimento até atingir o tamanho desejado.
	return strings.Repeat(pad, size-len(content)) + string(content), nil
}
