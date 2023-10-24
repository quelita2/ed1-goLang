# Algoritmos de Ordenação

| Algoritmos de Ordenação | Complexidade de Tempo | 
|-------------------------|------------------------|
| Selection Sort | Pior caso: O(n^2)<br>Melhor caso: Ω(n^2)|
| Bubble Sort | Pior caso: O(n^2)<br>Melhor caso: Ω(n)|
| Insertion Sort | Pior caso: O(n^2)<br>Melhor caso: Ω(n)|
| Merge Sort | Pior caso: O(nlogn)<br>Melhor caso: Ω(nlogn)|

### 1. Selection Sort
O algoritmo varre um vetor múltiplas vezes comparando se o menor valor indicado pelo indice é de fato o menor valor entre os demais, caso não seja, trocamos a posição do atual indice que possui menor valor com o novo indice de menor valor, e assim, seguimos a para a proxima ordenação até finalizar todo o vetor. 

Vale resssaltar que a ultima posição do vetor já será o maior valor do vetor e que não precisamos compará-lo com outros valores. Na primeira vez visitamos n-1 índices do array, na segunda vez visitamos n-2 índices do array, na terceira vez visitamos n-3 índices do array, e este padrão se segue até que quase todo o array esteja ordenado, onde visitaremos 1 índice do array. 

Logo, note que teremos (n-1) + (n-2) + (n-3) + ... + 1, que na verdade é uma P.A. de razão 1. Resolvendo esta equação, vamos ter um polinômio de segundo grau na forma an² + bn + c. Por isso dizemos que o **SelectionSort é O(n²)**, que é um desempenho muito pior do que O(n). 

### 2. Bubble Sort
Este algoritmo funciona comparando repetidamente pares de elementos adjacentes em uma lista e trocando-os se estiverem na ordem errada. O processo é repetido até que nenhuma troca seja necessária, o que significa que a lista está ordenada. 

Este algoritmo recebe o nome de "Bubble Sort" porque os elementos menores "flutuam" gradualmente para a parte superior da lista, da mesma forma que as bolhas sobem na água.

### 3. Insertion Sort
O algoritmo está definido para percorrer a lista da esquerda para a direita, inserir cada elemento na posição correta na parte ordenada da lista e expandir a parte ordenada à medida que avança.

### 4. Merge Sort
O Merge Sort é um algoritmo de ordenação que segue o paradigma "divide e conquista". Ele funciona dividindo a lista original em pedaços menores, ordenando esses pedaços e, em seguida, mesclando-os para obter uma lista ordenada. A ideia fundamental é dividir a lista até que você tenha sublistas pequenas o suficiente para serem ordenadas facilmente, e depois combinar essas sublistas ordenadas para criar uma lista ordenada maior.
