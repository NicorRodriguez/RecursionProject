package main
import "fmt"
import "time"

// Función que utliza la recursio para calcular la siguiente sucesión:
// n -> n/2 (n es par)
// n -> 3n + 1 (n es impar)
func suc(n int) int {

	i := 1

    if n != 1 && n % 2 == 0 {
        i = suc(int(n/2)) + 1
	} else if n != 1 {
        i = suc(int(n * 3 +1)) + 1
	}

    return i

}

func main() {
	t1 := time.Now()

	total := 0
	max := 2

	for i := 2; i <= 1000000; i++ {
		mysuc := suc(i)
		if mysuc > total {
			total = mysuc
			max = i
		}
	}

	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println(diff)

	fmt.Println("Cantidad de digitos:", total)
	fmt.Println("Numero en cuestion::", max)
}