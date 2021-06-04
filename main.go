package main

import (
    "fmt"
    "math"
    "time"
    "os"
    "os/exec"
    "strconv"
)

//FUNCIÓN PARA CALCULAR FACTORIAL
func factorial(num int) int{
    var result int = 1
    if num < 0{
        fmt.Println("Numero debe ser mayor que cero")
    }else{
        for i:=1; i<= num; i++{
            result *= i
        }
    }
    return result
}

//FUNCION PARA PERMUTACIÓN SIN REPETICIÓN
func permutacionSinRepe(num int) int{
    result:=factorial(num)
    return result
}

func permutacionConRepe(cantidad ...int) float64{
    result:=0.0
    var n int = 5
    multFact:=1
    for _, num:= range cantidad{
        multFact = multFact * factorial(num)
    }
    result = float64(factorial(n) / multFact)
    return result
}

//FUNCIÓN PARA COMBINACIÓN SIN REPETECIÓN
func combinacionSinRepe(n,r int) int{
    result:=factorial(n)/(factorial(n-r) * factorial(r))
    return result
}

//FUNCIÓN PARA COMBINACIÓN CON REPETICIÓN
func combinacionConRepe(n,r int) int {
    result:=factorial(n+r-1)/(factorial(r)*factorial(n-1))
    return result
}

//FUNCIÓN PARA VARIACIONES SIN REPETICIÓN
func variacionesSinRepe(n,r int) int {
    result:=factorial(n)/factorial(n-r)
    return result
}

//FUNCIÓN PARA VARIACIONES CON REPETICIÓN
func variacionesConRepe(n,r float64) float64 {
    result:=math.Pow(n, r)
    return result
}

//FUNCIÓN PARA EMPAREJAMIENTOS
func emparejamiento(n,m int) int {
    mf:=factorial(m)
    nf:=factorial(n)
    result:=factorial(m*n) / int(math.Pow(float64(mf),float64(n)) * float64(nf))
    return result
}

//FUNCIÓN PRINCIPAL
func main(){
    cmd:= exec.Command("cmd","/c","cls")
    cmd.Stdout = os.Stdout
    clear := exec.Command("clear")
    clear.Stdout = os.Stdout

    var option, n, r, m int
    var input string
    MEMU:
        fmt.Println("1- Permutacion sin repeticion")
        fmt.Println("2- Permutacion con repeticion")
        fmt.Println("3- Combinacion sin repeticion")
        fmt.Println("4- Combinacion con repeticion")
        fmt.Println("5- Variaciones sin repeticion")
        fmt.Println("6- Variaciones con repeticion")
        fmt.Println("7- Emparejamientos")
        fmt.Println("8 - Salir")

        fmt.Println("Selecciona una opcion: ")
        fmt.Scanln(&option)

        switch option{
        case 1:
            CASE1:
            fmt.Print("Ingresa n: ")
            fmt.Scanln(&input)
            val , err := strconv.Atoi(input)
            if err != nil {
                fmt.Println("El valor ingresado no es un numero, intenta nuevamente") 
                goto CASE1
            }else{
                n = val
                fmt.Println(permutacionSinRepe(n))
                time.Sleep(3 * time.Second)
                cmd.Run()
                clear.Run()
                goto MEMU
            }
        case 2:
            fmt.Println(permutacionConRepe(3,4,5))
            cmd.Run()
            goto MEMU
        case 3:
            fmt.Print("Ingresa n:")
            fmt.Scanln(&n)
            fmt.Print("Ingresa r:")
            fmt.Scanln(&r)
            fmt.Println(combinacionSinRepe(n,r))
            cmd.Run()
            goto MEMU
        case 4:
            fmt.Print("Ingresa n:")
            fmt.Scanln(&n)
            fmt.Print("Ingresa r:")
            fmt.Scanln(&r)
            fmt.Println(combinacionConRepe(n,r))
            cmd.Run()
            goto MEMU
        case 5:
            fmt.Print("Ingresa n:")
            fmt.Scanln(&n)
            fmt.Print("Ingresa r:")
            fmt.Scanln(&r)
            fmt.Println(variacionesSinRepe(n,r))
            cmd.Run()
            goto MEMU
        case 6:
            fmt.Print("Ingresa n:")
            fmt.Scanln(&n)
            fmt.Print("Ingresa r:")
            fmt.Scanln(&r)
            fmt.Println(variacionesConRepe(float64(n),float64(r)))
            cmd.Run()
            goto MEMU
        case 7:
            fmt.Print("Ingresa n:")
            fmt.Scanln(&n)
            fmt.Print("Ingrese m:")
            fmt.Scanln(&m)
            fmt.Println(emparejamiento(n,m))
            cmd.Run()
            goto MEMU
        case 8:
            fmt.Println("Saliendo...")
        default:
            fmt.Println("****** Selecciona una de las opciones en el menu ******")
            fmt.Println()
            goto MEMU




        }
}


