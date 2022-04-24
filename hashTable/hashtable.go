//Hashtable implementa os métodos para geração de hash e de inserção/busca/exclusão na lista.
package main

import (
	"fmt"
	"sync"
)

type Hashtable struct {
	itens map[int][]Pessoa
	lock  sync.RWMutex
}

//Hash gera e retorna código hashcode com base no parâmetro nome informado.
func hash(nome string) (key int) {
	for _, letra := range nome {
		key = 31*key + int(letra)
	}
	return
}

//Put insere pessoa na tabela com base no código hash dela.
func (ht *Hashtable) Put(pessoa Pessoa) int {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	key := hash(pessoa.Nome)
	if ht.itens == nil {
		ht.itens = make(map[int][]Pessoa)
	}
	ht.itens[key] = append(ht.itens[key], pessoa)
	return key
}

//Remove busca pela chave hash e remove pessoa da lista.
func (ht *Hashtable) Remove(nome string) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	key := hash(nome)
	delete(ht.itens, key)
	fmt.Println("\nExcluindo registro: ", nome)
}

//Get retorna o registro tipo pessoa para o parâmetro chave.
func (ht *Hashtable) Get(key int) []Pessoa {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	return ht.itens[key]
}

//Busca procura por parâmetro nome na lista e retorna registro pessoa. 
func (ht *Hashtable) Busca(nome string) []Pessoa {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	key := hash(nome)
	return ht.itens[key]
}
