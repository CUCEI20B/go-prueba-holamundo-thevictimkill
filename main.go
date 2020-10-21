package main


func burbuja(Lista []int){
 var auxiliar int
 for i := 0; i < len(Lista); i++ {
  for j := 0; j < len(Lista); j++ {
   if Lista[i] < Lista[j] {
    auxiliar = Lista[i]
    Lista[i] = Lista[j]
    Lista[j] = auxiliar
    }
  }
}
}

func main() {
  lista := []int{1,6,3,5,4,9,12,34}
  lista_slice := lista[0:6]
  burbuja(lista_slice)
}