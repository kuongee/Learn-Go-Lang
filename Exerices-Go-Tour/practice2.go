package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    Y := make([][]uint8, dy)
    for i := range Y {
        Y[i] = make([]uint8, dx)
    }
    
    /* My solution  */
    for i, y := range Y {
        for j := range y {
         	Y[i][j] = uint8(i ^ j)   
        }
    }
    
    /* Go github solution (Go Style) */
    for i, y := range Y {
        for j := range y {
         	y[j] = uint8(i ^ j)   
        }
    }
        
    return Y
}

func main() {
    pic.Show(Pic)
}

