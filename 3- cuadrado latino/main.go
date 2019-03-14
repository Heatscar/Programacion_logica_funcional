package main

import (
	"fmt"
)

func main() {
	var n, i, j, k int
	fmt.Print("n: ")
	fmt.Scanln(&n)

	for i=1; i<=n; i++ {
		j=i
		k=1
		for ; k<=n ; {
            if j>n {
                for j=1; k <= n ;  {
                    fmt.Printf(" %d", j)
                    k++
                	j++
                }
            } else {
            	fmt.Printf(" %d", j)
            }
            k++
			j++
        }
        fmt.Println("");
    }
	
}