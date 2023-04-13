package main

/*    FAZER:
- Método p/ Alocar e incializar os valores da LinkedList
- Método p/ Adicionar valores no fim LinkedList
- Método p/ Adicionar um valor no início
- Método p/ Adicionar um valor em uma posição específica da LinkedList
- Método p/ Remover um valor da LinkedList (final)
- Método p/ Remover um valor numa posição específica
- Método p/ Acessar item do ArrayList (Get)
- Método p/ Modificar item do ArrayList (Set)
- Método p/ Retorna o tamanho do ArrayList
- Método p/ Mostrar a lista ligada
*/

/*  LinkedList: 
os arrays são sempre armazenados em blocos contíguos de memória, e isto traz algumas implicações, por isso para diminuir a quantidade de espaços de memória pequenos desalocados é mais viável usar listas ligadas (Linked List).

LinkedLists (ou listas ligadas) são listas implementadas com nós, e cada nó possui um espaço de memória para armazenar o elemento e outro espaço de memória para armazenar o ponteiro (endereço de memória) para o nó seguinte.
*/

type LinkedList struct {
  head *No
  tam int
}

type No struct {
  value int
  next *No
}

// Alocar e incializar os valores da LinkedList
func (linkedList *LinkedList) Init(){
  // Ao contrário de ArrayLists, em que precisamos inicializar um array, LinkedLists podem inicializar com o primeiro nó nulo, pois sempre que adicionarmos um novo nó alocaremos espaço na RAM para o novo nó. Para a variável tam, sempre que eu adicionar um novo nó eu incrementarei esta variável, por isso ela iniciará com 0.
  linkedList.head = nil
  linkedList.tam = 0
}

// Adicionar valores à LinkedList
func (linkedList *LinkedList) Add(value int){
  if linkedList.head == nil{
    linkedList.head = &No{value: value, next: nil}
  }else{
    // lasItem é o valor anterior ao nil que é o último dado da lista
    lastItem := linkedList.head
    // Procurando o espaço de memória que tem nil como valor, ao encontrar o for encerra
    for lastItem.next != nil{
      lastItem = lastItem.next
    }
    // Aplicando valor para o ultimo item, o próximo dado depois do lastItem será o novo nó
    lastItem.next = &No{value: value, next: nil}
  }
  
    linkedList.tam++
}

// Adicionar um valor no inicio da LinkedList
func (linkedList *LinkedList) AddOnStart(value int){
    // se a lista estiver vazia, ele add apenas um nó
    if linkedList.head == nil{
        linkedList.head = &No{value: value, next: nil}
    }else{
        // caso haja espaço, receberemos um novo head
        noAdd := &No{value: value, next: linkedList.head}
        linkedList.head = noAdd
    }
    linkedList.tam++
}

// Adicionar um valor em uma posição específica da LinkedList
func (linkedList *LinkedList) AddOnIndex (value int, index int){
    //o indice pode ser tanto o inicio como uma posição específica da lista ligada
	if index == 0{
		linkedList.AddOnStart(value)
	}else{
		aux := linkedList.head
		for i:=0; i<index-1; i++{
			aux = aux.next
		}

		noAdd := &No{value: value, next: aux.next}
		aux.next = noAdd
	}
	linkedList.tam++
}

// Remover um valor da LinkedList (final)
func (linkedList *LinkedList) Remove(){
	linkedList.tam--
}

// Remover um valor numa posição específica
func (linkedList *LinkedList) RemoveOnIndex(index int){
	
}
// Acessar item do ArrayList (Get)
// Modificar item do ArrayList (Set)
// Retorna o tamanho do ArrayList
// Mostrar a lista ligada