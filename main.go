package main

import (
	"fmt"
)

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		for j := 1; j <= y; j++ {
			if j == 1 || j == y {
				// boucle o-------o
				fmt.Print("o")
				for i := 2; i < x; i++ {
					fmt.Print("-")
				}
				if x > 1 {
					fmt.Println("o")
				} else {
					fmt.Println()
				}

			} else {
				// boucle |       |
				fmt.Print("|")
				for i := 2; i < x; i++ {
					fmt.Print(" ")
				}
				if x > 1 {
					fmt.Println("|")
				} else {
					fmt.Println()
				}

			}
		}
	}
}

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		for j := 1; j <= y; j++ {
			if j == 1 || j == y {
				if j == 1 {
					fmt.Print("/")
					for i := 2; i < x; i++ {
						fmt.Print("*")
					}
					if x > 1 {
						fmt.Println("\\")
					} else {
						fmt.Println()
					}
				}
				if y > 1 {
					if j == y {
						fmt.Print("\\")
						for i := 2; i < x; i++ {
							fmt.Print("*")
						}
						if x > 1 {
							fmt.Println("/")
						} else {
							fmt.Println()
						}
					}
				}
			} else {
				// boucle *       *
				fmt.Print("*")
				for i := 2; i < x; i++ {
					fmt.Print(" ")
				}
				if x > 1 {
					fmt.Println("*")
				} else {
					fmt.Println()
				}
			}
		}
	}
}

func QuadC(x, y int) {
	if x > 0 && y > 0 {
		for j := 1; j <= y; j++ {
			if j == 1 || j == y {
				if j == 1 {
					// boucle o-------o
					fmt.Print("A")
					for i := 2; i < x; i++ {
						fmt.Print("B")
					}
					if x > 1 {
						fmt.Println("A")
					} else {
						fmt.Println()
					}
				} else {
					// boucle o-------o
					fmt.Print("A")
					for i := 2; i < x; i++ {
						fmt.Print("B")
					}
					if x > 1 {
						fmt.Println("A")
					} else {
						fmt.Println()
					}
				}
			} else {
				// boucle |       |
				fmt.Print("B")
				for i := 2; i < x; i++ {
					fmt.Print(" ")
				}

				if x > 1 {
					fmt.Println("B")
				} else {
					fmt.Println()
				}
			}
		}
	}
}

func QuadD(x, y int) {
	if x > 0 && y > 0 {
		for j := 1; j <= y; j++ {
			if j == 1 || j == y {
				if j == 1 {
					fmt.Print("A")
					for i := 2; i < x; i++ {
						fmt.Print("B")
					}
					if x > 1 {
						fmt.Println("C")
					} else {
						fmt.Println()
					}
				}
				if y > 1 {
					if j == y {
						fmt.Print("A")
						for i := 2; i < x; i++ {
							fmt.Print("B")
						}
						if x > 1 {
							fmt.Println("C")
						} else {
							fmt.Println()
						}
					}
				}
			} else {
				// boucle *       *
				fmt.Print("B")
				for i := 2; i < x; i++ {
					fmt.Print(" ")
				}
				if x > 1 {
					fmt.Println("B")
				} else {
					fmt.Println()
				}
			}
		}
	}
}

func QuadE(x, y int) {
	if x > 0 && y > 0 {
		for j := 1; j <= y; j++ {
			if j == 1 || j == y {
				if j == 1 {
					fmt.Print("A")
					for i := 2; i < x; i++ {
						fmt.Print("B")
					}
					if x > 1 {
						fmt.Println("C")
					} else {
						fmt.Println()
					}
				}
				if y > 1 {
					if j == y {
						fmt.Print("C")
						for i := 2; i < x; i++ {
							fmt.Print("B")
						}
						if x > 1 {
							fmt.Println("A")
						} else {
							fmt.Println()
						}
					}
				}
			} else {
				// boucle *       *
				fmt.Print("B")
				for i := 2; i < x; i++ {
					fmt.Print(" ")
				}
				if x > 1 {
					fmt.Println("B")
				} else {
					fmt.Println()
				}
			}
		}
	}
}

func main() {
	// Test de la fonction QuadA avec tous les cas possibles
	fmt.Println("=====================================================")
	fmt.Println("Test de la fonction QuadA avec tous les cas possibles")
	fmt.Println("=====================================================")
	// Valeurs négatives: 3 cas
	fmt.Println("Valeurs négatives: 3 cas")
	// 1) x<0  et y>0
	fmt.Println("1) x<0  et y>0 : QuadA(-1, 4)")
	QuadA(-1, 4)
	// 2) x>0  et y<0
	fmt.Println("2) x>0  et y<0 : QuadA(4, -1)")
	QuadA(4, -1)
	// 3) x<0  et y<0
	fmt.Println("3) x<0  et y<0 : QuadA(-4, -1)")
	QuadA(-4, -1)
	fmt.Println("Rien affiché  pour toute valeur négative  :-)")
	// Valeurs nulles
	fmt.Println()
	fmt.Println("Valeurs nulles: 3 cas")
	// 1) x=0  et y>0
	fmt.Println("1) x=0  et y>0 : QuadA(0, 4)")
	QuadA(0, 4)
	// 2) x>0  et y=0
	fmt.Println("2) x>0  et y=0 : QuadA(4, 0)")
	QuadA(4, 0)
	// 3) x=0  et y=0
	fmt.Println("3) x=0  et y=0 : QuadA(0, 0)")
	QuadA(0, 0)
	fmt.Println()
	fmt.Println("Rien affiché pour toutes les valeurs nulles :-)")
	// Valeurs strictement positives

	fmt.Println()
	fmt.Println()
	fmt.Println("strictement positives: 6 cas")
	// 1) x=1  et y=1
	fmt.Println("1) x=1  et y=1 : QuadA(1, 1)")
	QuadA(1, 1)
	// 2) x=1  et y=2
	fmt.Println("2) x=1  et y=2 : QuadA(1, 2)")
	QuadA(1, 2)
	// 3) x=1  et y>2
	fmt.Println("3) x=1  et y>2 : QuadA(1, 4)")
	QuadA(1, 4)
	// 4) x=2  et y=1
	fmt.Println("4) x=2  et y=1 : QuadA(2, 1)")
	QuadA(2, 1)
	// 5) x>2  et y=2
	fmt.Println("5) x>2  et y=2 : QuadA(4, 2)")
	QuadA(4, 2)
	// 6) x>2  et y>2
	fmt.Println("6) x>2  et y>2 : QuadA(4, 4)")
	QuadA(4, 4)
	fmt.Println()
	fmt.Println()
	// Test de la fonction QuadB avec tous les cas possibles
	fmt.Println("=====================================================")
	fmt.Println("Test de la fonction QuadB avec tous les cas possibles")
	fmt.Println("=====================================================")
	// Valeurs négatives: 3 cas	fmt.Println("Valeurs négatives: 3 cas")
	fmt.Println()
	// 1) x<0  et y>0
	fmt.Println("1) x<0  et y>0 : QuadB(-1, 4)")
	QuadB(-1, 4)
	// 2) x>0  et y<0
	fmt.Println("2) x>0  et y<0 : QuadB(4, -1)")
	QuadB(4, -1)
	// 3) x<0  et y<0
	fmt.Println("3) x<0  et y<0 : QuadB(-4, -1)")
	QuadB(-4, -1)
	// Valeurs nulles
	fmt.Println()
	fmt.Println()
	fmt.Println("Valeurs nulles: 3 cas")
	// 1) x=0  et y>0
	fmt.Println("1) x=0  et y>0 : QuadB(0, 4)")
	QuadB(0, 4)
	// 2) x>0  et y=0
	fmt.Println("2) x>0  et y=0 : QuadB(4, 0)")
	QuadB(4, 0)
	// 3) x=0  et y=0
	fmt.Println("3) x=0  et y=0 : QuadB(0, 0)")
	QuadB(0, 0)

	// Valeurs strictement positives
	fmt.Println()
	fmt.Println()
	fmt.Println("strictement positives: 6 cas")
	// 1) x=1  et y=1
	fmt.Println("1) x=1  et y=1 : QuadB(1, 1)")
	QuadB(1, 1)
	// 2) x=1  et y=2
	fmt.Println("2) x=1  et y=2 : QuadB(1, 2)")
	QuadB(1, 2)
	// 3) x=1  et y>2
	fmt.Println("3) x=1  et y>2 : QuadB(1, 4)")
	QuadB(1, 4)
	// 4) x=2  et y=1
	fmt.Println("4) x=2  et y=1 : QuadB(2, 1)")
	QuadB(2, 1)
	// 5) x>2  et y=2
	fmt.Println("5) x>2  et y=2 : QuadB(4, 2)")
	QuadB(4, 2)
	// 6) x>2  et y>2
	fmt.Println("6) x>2  et y>2 : QuadB(4, 4)")
	QuadB(4, 4)
	fmt.Println()
	fmt.Println()
	// Test de la fonction QuadC avec tous les cas possibles
	fmt.Println("=====================================================")
	fmt.Println("Test de la fonction QuadC avec tous les cas possibles")
	fmt.Println("=====================================================")
	// Valeurs négatives: 3 cas
	fmt.Println()
	fmt.Println("Valeurs négatives: 3 cas")
	// 1) x<0  et y>0
	fmt.Println("1) x<0  et y>0 : QuadC(-1, 4)")
	QuadC(-1, 4)
	// 2) x>0  et y<0
	fmt.Println("2) x>0  et y<0 : QuadC(4, -1)")
	QuadC(4, -1)
	// 3) x<0  et y<0
	fmt.Println("3) x<0  et y<0 : QuadC(-4, -1)")
	QuadC(-4, -1)
	// Valeurs nulles
	fmt.Println()
	fmt.Println()
	fmt.Println("Valeurs nulles: 3 cas")
	// 1) x=0  et y>0
	fmt.Println("1) x=0  et y>0 : QuadC(0, 4)")
	QuadC(0, 4)
	// 2) x>0  et y=0
	fmt.Println("2) x>0  et y=0 : QuadC(4, 0)")
	QuadC(4, 0)
	// 3) x=0  et y=0
	fmt.Println("3) x=0  et y=0 : QuadC(0, 0)")
	QuadC(0, 0)
	// Valeurs strictement positives
	fmt.Println()
	fmt.Println()
	fmt.Println("strictement positives: 6 cas")
	// 1) x=1  et y=1
	fmt.Println("1) x=1  et y=1 : QuadC(1, 1)")
	QuadC(1, 1)
	// 2) x=1  et y=2
	fmt.Println("2) x=1  et y=2 : QuadC(1, 2)")
	QuadC(1, 2)
	// 3) x=1  et y>2
	fmt.Println("3) x=1  et y>2 : QuadC(1, 4)")
	QuadC(1, 4)
	// 4) x=2  et y=1
	fmt.Println("4) x=2  et y=1 : QuadC(2, 1)")
	QuadC(2, 1)
	// 5) x>2  et y=2
	fmt.Println("5) x>2  et y=2 : QuadC(4, 2)")
	QuadC(4, 2)
	// 6) x>2  et y>2
	fmt.Println("6) x>2  et y>2 : QuadC(4, 4)")
	QuadC(4, 4)
	fmt.Println()
	fmt.Println()
	// Test de la fonction QuadD avec tous les cas possibles
	fmt.Println("=====================================================")
	fmt.Println("Test de la fonction QuadD avec tous les cas possibles")
	fmt.Println("=====================================================")
	// Valeurs négatives: 3 cas
	fmt.Println()
	fmt.Println()
	fmt.Println("Valeurs négatives: 3 cas")
	// 1) x<0  et y>0
	fmt.Println("1) x<0  et y>0 : QuadD(-1, 4)")
	QuadD(-1, 4)
	// 2) x>0  et y<0
	fmt.Println("2) x>0  et y<0 : QuadD(4, -1)")
	QuadD(4, -1)
	// 3) x<0  et y<0
	fmt.Println("3) x<0  et y<0 : QuadD(-4, -1)")
	QuadD(-4, -1)
	// Valeurs nulles
	fmt.Println()
	fmt.Println()
	fmt.Println("Valeurs nulles: 3 cas")
	// 1) x=0  et y>0
	fmt.Println("1) x=0  et y>0 : QuadD(0, 4)")
	QuadD(0, 4)
	// 2) x>0  et y=0
	fmt.Println("2) x>0  et y=0 : QuadD(4, 0)")
	QuadD(4, 0)
	// 3) x=0  et y=0
	fmt.Println("3) x=0  et y=0 : QuadD(0, 0)")
	QuadD(0, 0)
	// Valeurs strictement positives
	fmt.Println()
	fmt.Println("strictement positives: 6 cas")
	// 1) x=1  et y=1
	fmt.Println("1) x=1  et y=1 : QuadD(1, 1)")
	QuadD(1, 1)
	// 2) x=1  et y=2
	fmt.Println("2) x=1  et y=2 : QuadD(1, 2)")
	QuadD(1, 2)
	// 3) x=1  et y>2
	fmt.Println("3) x=1  et y>2 : QuadD(1, 4)")
	QuadD(1, 4)
	// 4) x=2  et y=1
	fmt.Println("4) x=2  et y=1 : QuadD(2, 1)")
	QuadD(2, 1)
	// 5) x>2  et y=2
	fmt.Println("5) x>2  et y=2 : QuadD(4, 2)")
	QuadD(4, 2)
	// 6) x>2  et y>2
	fmt.Println("6) x>2  et y>2 : QuadD(4, 4)")
	QuadD(4, 4)
	fmt.Println()
	fmt.Println()
	// Test de la fonction QuadE avec tous les cas possibles
	fmt.Println("=====================================================")
	fmt.Println("Test de la fonction QuadE avec tous les cas possibles")
	fmt.Println("=====================================================")
	// Valeurs négatives: 3 cas
	fmt.Println()
	fmt.Println("Valeurs négatives: 3 cas")
	// 1) x<0  et y>0
	fmt.Println("1) x<0  et y>0 : QuadE(-1, 4)")
	QuadE(-1, 4)
	// 2) x>0  et y<0
	fmt.Println("2) x>0  et y<0 : QuadE(4, -1)")
	QuadE(4, -1)
	// 3) x<0  et y<0
	fmt.Println("3) x<0  et y<0 : QuadE(-4, -1)")
	QuadE(-4, -1)
	// Valeurs nulles
	fmt.Println()
	fmt.Println("Valeurs nulles: 3 cas")
	// 1) x=0  et y>0
	fmt.Println("1) x=0  et y>0 : QuadE(0, 4)")
	QuadE(0, 4)
	// 2) x>0  et y=0
	fmt.Println("2) x>0  et y=0 : QuadE(4, 0)")
	QuadE(4, 0)
	// 3) x=0  et y=0
	fmt.Println("3) x=0  et y=0 : QuadE(0, 0)")
	QuadE(0, 0)
	// Valeurs strictement positives
	fmt.Println()
	fmt.Println("strictement positives: 6 cas")
	// 1) x=1  et y=1
	fmt.Println("1) x=1  et y=1 : QuadE(1, 1)")
	QuadE(1, 1)
	// 2) x=1  et y=2
	fmt.Println("2) x=1  et y=2 : QuadE(1, 2)")
	QuadE(1, 2)
	// 3) x=1  et y>2
	fmt.Println("3) x=1  et y>2 : QuadE(1, 4)")
	QuadE(1, 4)
	// 4) x=2  et y=1
	fmt.Println("4) x=2  et y=1 : QuadE(2, 1)")
	QuadE(2, 1)
	// 5) x>2  et y=2
	fmt.Println("5) x>2  et y=2 : QuadE(4, 2)")
	QuadE(4, 2)
	// 6) x>2  et y>2
	fmt.Println("6) x>2  et y>2 : QuadE(4, 4)")
	QuadE(4, 4)
}
