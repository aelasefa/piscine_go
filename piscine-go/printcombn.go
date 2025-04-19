package piscine

import "github.com/01-edu/z01"

func printcomb1() {
	for a := '0'; a <= '9'; a++ {
		z01.PrintRune(a)
		if a != '9' {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func printcomb2() {
	for a := '0'; a <= '8'; a++ {
		for b := a + 1; b <= '9'; b++ {
			z01.PrintRune(a)
			z01.PrintRune(b)
			if a != '8' || b != '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

func printcomb3() {
	for a := '0'; a <= '7'; a++ {
		for b := a + 1; b <= '8'; b++ {
			for c := b + 1; c <= '9'; c++ {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)
				if a != '7' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func printcomb4() {
	for a := '0'; a <= '6'; a++ {
		for b := a + 1; b <= '7'; b++ {
			for c := b + 1; c <= '8'; c++ {
				for d := c + 1; d <= '9'; d++ {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
					z01.PrintRune(d)
					if a != '6' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func printcomb5() {
	for a := '0'; a <= '5'; a++ {
		for b := a + 1; b <= '6'; b++ {
			for c := b + 1; c <= '7'; c++ {
				for d := c + 1; d <= '8'; d++ {
					for e := d + 1; e <= '9'; e++ {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(c)
						z01.PrintRune(d)
						z01.PrintRune(e)
						if a != '5' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func printcomb6() {
	for a := '0'; a <= '4'; a++ {
		for b := a + 1; b <= '5'; b++ {
			for c := b + 1; c <= '6'; c++ {
				for d := c + 1; d <= '7'; d++ {
					for e := d + 1; e <= '8'; e++ {
						for f := e + 1; f <= '9'; f++ {
							z01.PrintRune(a)
							z01.PrintRune(b)
							z01.PrintRune(c)
							z01.PrintRune(d)
							z01.PrintRune(e)
							z01.PrintRune(f)
							if a != '4' {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func printcomb7() {
	for a := '0'; a <= '3'; a++ {
		for b := a + 1; b <= '4'; b++ {
			for c := b + 1; c <= '5'; c++ {
				for d := c + 1; d <= '6'; d++ {
					for e := d + 1; e <= '7'; e++ {
						for f := e + 1; f <= '8'; f++ {
							for g := f + 1; g <= '9'; g++ {
								z01.PrintRune(a)
								z01.PrintRune(b)
								z01.PrintRune(c)
								z01.PrintRune(d)
								z01.PrintRune(e)
								z01.PrintRune(f)
								z01.PrintRune(g)
								if a != '3' {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func printcomb8() {
	for a := '0'; a <= '2'; a++ {
		for b := a + 1; b <= '3'; b++ {
			for c := b + 1; c <= '4'; c++ {
				for d := c + 1; d <= '5'; d++ {
					for e := d + 1; e <= '6'; e++ {
						for f := e + 1; f <= '7'; f++ {
							for g := f + 1; g <= '8'; g++ {
								for h := g + 1; h <= '9'; h++ {
									z01.PrintRune(a)
									z01.PrintRune(b)
									z01.PrintRune(c)
									z01.PrintRune(d)
									z01.PrintRune(e)
									z01.PrintRune(f)
									z01.PrintRune(g)
									z01.PrintRune(h)
									if a != '2' {
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func printcomb9() {
	for a := '0'; a <= '1'; a++ {
		for b := a + 1; b <= '2'; b++ {
			for c := b + 1; c <= '3'; c++ {
				for d := c + 1; d <= '4'; d++ {
					for e := d + 1; e <= '5'; e++ {
						for f := e + 1; f <= '6'; f++ {
							for g := f + 1; g <= '7'; g++ {
								for h := g + 1; h <= '8'; h++ {
									for i := h + 1; i <= '9'; i++ {
										z01.PrintRune(a)
										z01.PrintRune(b)
										z01.PrintRune(c)
										z01.PrintRune(d)
										z01.PrintRune(e)
										z01.PrintRune(f)
										z01.PrintRune(g)
										z01.PrintRune(h)
										z01.PrintRune(i)
										if a != '1' {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintCombN(n int) {
	if n == 1 {
		printcomb1()
	}
	if n == 2 {
		printcomb2()
	}
	if n == 3 {
		printcomb3()
	}
	if n == 4 {
		printcomb4()
	}
	if n == 5 {
		printcomb5()
	}
	if n == 6 {
		printcomb6()
	}
	if n == 7 {
		printcomb7()
	}
	if n == 8 {
		printcomb8()
	}
	if n == 9 {
		printcomb9()
	}
}
