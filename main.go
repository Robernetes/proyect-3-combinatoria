package main

//LIBRERIAS IMPORTADAS EN ESTE PROYECTO
import (
	"fmt"
	cp "github.com/Robernetes/proyect-3-combinatoria/combinatoria"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
)


//FUNCIÓN PRINCIPAL
func main(){
    limpiar := exec.Command("pwd") // VARIABLE QUE VA A LIMPIAR LA CONSOLA DEPENDIENDO SI ES LINUX O WINDOWS
    system := runtime.GOOS // VARIABLE QUE VA CAPTURAR EN QUE OS SE ESTÁ EJECUTANDO EL PROGRAMA

    //CONDICIÓN QUE DEPENDIENDO SI ES WINDOWS O LINUX VA A RETORNAR EL METODO PARA LIMPIAR LA CONSOLA
    if "linux" == system{
      limpiar = exec.Command("clear")
      limpiar.Stdout = os.Stdout
      fmt.Println(reflect.TypeOf(limpiar))
    }
    if "windows" == system{
      limpiar = exec.Command("cmd","/c","cls")
      limpiar.Stdout = os.Stdout
    }

    //DECLARACIÓN DE VARIABLES 
    var option, n, r, m int 
    var num[] int
    var input, res string

    //MENU PRINCIPAL QUE SE MUESTRA AL USUARIO
    MEMU:
        fmt.Println("*****MENU PRINCIPAL*****")
        fmt.Println("1- Permutacion sin repeticion")
        fmt.Println("2- Permutacion con repeticion")
        fmt.Println("3- Combinacion sin repeticion")
        fmt.Println("4- Combinacion con repeticion")
        fmt.Println("5- Variaciones sin repeticion")
        fmt.Println("6- Variaciones con repeticion")
        fmt.Println("7- Emparejamientos")
        fmt.Println("8- Salir")

        fmt.Println("Selecciona una opcion: ")
        fmt.Scanln(&option)

        //CONDICIÓN PARA RE UTILIZAR LAS VARIABLES 'n' Y 'r' EN LAS FUNCIONES QUE PIDAN TALES VARIABLES
        if (option == 3) || (option == 4) || (option == 5) || (option == 6) || (option == 7) {
            fmt.Print("Ingresa n:")
            fmt.Scanln(&input)
            val , err := strconv.Atoi(input)
            if err != nil {
                fmt.Println("El valor intresado no es numerico, intenta nuevamente")
            }else{
                n = val
            }

            fmt.Print("Ingresa r:")
            fmt.Scanln(&input)
            val , err = strconv.Atoi(input)
            if err != nil {
                fmt.Println("El valor intresado no es numerico, intenta nuevamente")
            }else{
                r = val
            }

        }

        //ESTRUCTURA DE CONTROL SWITCH PARA LA SELECCION DE LAS FUNCIONES
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
                fmt.Println("Permutaciones sin repeticion:", cp.PermutacionSinRepe(n))
                fmt.Println("Presione Enter para continuar al menú principal...")
                fmt.Scanln()
                limpiar.Run()
                goto MEMU
            }
        case 2:
            ADD:
            fmt.Print("Ingresa valor de grupo: ")
            fmt.Scanln(&r)
            num = append(num, r)
            fmt.Print("Presione 'y' para add otro grupo o 'c' para continuar: ")
            fmt.Scanln(&res)
            if res == "y" {
                goto ADD
            }
            fmt.Println("Permutaciones con repeticion: ",cp.PermutacionConRepe(num...))
            fmt.Println("Presione Enter para continuar al menú principal...")
            fmt.Scanln()
            limpiar.Run()
            goto MEMU
        case 3:
            fmt.Println("Combinacion Sin Repeticion de ",n, "y", r, "es:", cp.CombinacionSinRepe(n,r))
            fmt.Println("Presione Enter para continuar al menú principal...")
            fmt.Scanln()
            limpiar.Run()
            goto MEMU
        case 4:
            fmt.Print("Ingresa n:")
            fmt.Scanln(&n)
            fmt.Print("Ingresa r:")
            fmt.Scanln(&r)
            fmt.Println(cp.CombinacionConRepe(n,r))
            fmt.Println("Presione Enter para continuar al menú principal...")
            fmt.Scanln()
            limpiar.Run()
            goto MEMU
        case 5:
            fmt.Print("Ingresa n:")
            fmt.Scanln(&n)
            fmt.Print("Ingresa r:")
            fmt.Scanln(&r)
            fmt.Println(cp.VariacionesSinRepe(n,r))
            fmt.Println("Presione Enter para continuar al menú principal...")
            fmt.Scanln()
            limpiar.Run()
            goto MEMU
        case 6:
            fmt.Print("Ingresa n:")
            fmt.Scanln(&n)
            fmt.Printf("Ingresa r:")
            fmt.Scanln(&r)
            fmt.Println(cp.VariacionesConRepe(float64(n),float64(r)))
            fmt.Println("Presione Enter para continuar al menú principal...")
            fmt.Scanln()
            limpiar.Run()
            goto MEMU
        case 7:
            fmt.Print("Ingresa n:")
            fmt.Scanln(&n)
            fmt.Print("Ingrese m:")
            fmt.Scanln(&m)
            fmt.Println(cp.Emparejamiento(n,m))
            fmt.Println("Presione Enter para continuar al menú principal...")
            fmt.Scanln()
            limpiar.Run()
            goto MEMU
        case 8:
            fmt.Println("Saliendo...")
        default:
            fmt.Println("****** Selecciona una de las opciones en el menu ******")
            fmt.Println()
            goto MEMU




        }
}


