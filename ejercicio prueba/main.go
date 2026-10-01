package main

import main

import "fmt"

var productosventas
package main

import (
	"fmt"
)

var productosVendidos []string
var subtotalesVentas []float64

func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)
	productosVendidos = append(productosVendidos, nombre)
	subtotalesVentas = append(subtotalesVentas, subtotal)
}

                                               
	if len(productosVendidos) == 0 {
		fmt.Println("\nNo existen ventas registradas aún.")
		return
	}

	totalRecaudado := 0.0
	for _, subtotal := range subtotalesVentas {
		totalRecaudado += subtotal
	}

	fmt.Println("\n=== ESTADÍSTICAS DE VENTAS ===")
	fmt.Printf("Total recaudado: $%.2f\n", totalRecaudado)
	fmt.Println("Detalle de ventas:")
	for i := 0; i < len(productosVendidos); i++ {
		fmt.Printf("- %s: $%.2f\n", productosVendidos[i], subtotalesVentas[i])
	}
}

