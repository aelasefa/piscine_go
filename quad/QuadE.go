package piscine

import "github.com/01-edu/z01"

func QuadE(x,y int) {

  for i := 1 ; i <= y ; i++ {
  
    if i == 1 { 

      //First line
      for i := 1 ; i <= x ; i++ {
  
        if i == 1 {
          z01.PrintRune('A')
        }
    
        if i > 1 && i < x {
          z01.PrintRune('B')
        }
    
        if i == x && i != 1 {
          z01.PrintRune('C')
          z01.PrintRune('\n')
        }
        if i == x && i == 1 {
          z01.PrintRune('\n')
        }
      }
    
    }

    if i > 1 && i < y {

      //Second line
      for i := 1 ; i <= x ; i++ {
  
        if i == 1 {
          z01.PrintRune('B')
        }
    
        if i > 1 && i < x {
          z01.PrintRune(' ')
        }
    
        if i == x && i != 1 {
          z01.PrintRune('B')
          z01.PrintRune('\n')
        }
        if i == x && i == 1 {
          z01.PrintRune('\n')
        }
        
      }

    }

      if i >= y && i != 1 {
        
        //Third line
        for i := 1 ; i <= x ; i++ {
    
          if i == 1 {
            z01.PrintRune('C')
          }
      
          if i > 1 && i < x {
            z01.PrintRune('B')
          }
      
          if i == x && i != 1 {
            z01.PrintRune('A') 
            z01.PrintRune('\n')
          }
          if i == x && i == 1 {
            z01.PrintRune('\n')
          }
        }
        
      }
  
  }

}
