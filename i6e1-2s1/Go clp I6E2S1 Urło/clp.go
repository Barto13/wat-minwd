package main

import (
        "fmt"
        "github.com/lanl/clp"
        "math"
)

func main() {
        pinf := math.Inf(1)
        ninf := math.Inf(-1)
        simp := clp.NewSimplex()
        simp.EasyLoadDenseProblem(

                []float64{1.0, 1.0, 1.0},
                [][2]float64{
                        {1, 6}, // 1 ≤ a ≤ 6
                        {1, 6}, // 1 ≤ b ≤ 6
                        {1, 6}, // 1 ≤ c ≤ 6
                },
                [][]float64{
                        {1.0, 1.0, -1.0, 0.0, pinf},
                        {1.0, 0.0, 1.0, -1.0, pinf},
                        {ninf, 1.0, -2.0, 1.0, -1.0},
                })
        simp.SetOptimizationDirection(clp.Maximize)

        simp.Primal(clp.NoValuesPass, clp.NoStartFinishOptions)
        soln := simp.PrimalColumnSolution()

        fmt.Printf("Die 1 = %.0f\n", soln[0])
        fmt.Printf("Die 2 = %.0f\n", soln[1])
        fmt.Printf("Die 3 = %.0f\n", soln[2])
		
		//Wynik:
		//Die 1 = 6
		//Die 2 = 5
		//Die 3 = 3
}