package main

//LIBRERIAS IMPORTADAS EN ESTE PROYECTO
import (
	"fmt"
	"os"
	"os/exec"
    "time"
	"strconv"

	cp "github.com/Robernetes/proyect-3-combinatoria/combinatoria"
	"github.com/gookit/color"
)

var limpiarW = exec.Command("cmd", "/c", "cls")
var limpiarL = exec.Command("cmd", "/c", "cls")

//FUNCIÓN PRINCIPAL
func main() {
	//DECLARACIÓN DE VARIABLES
	var option, n, r, m int
	var num []int
	var input, res string


	//MENU PRINCIPAL QUE SE MUESTRA AL USUARIO
MEMU:
    option = 0
	fmt.Println("[*******MENU PRINCIPAL*******]")
	color.Println("<blue>[1]</> Permutacion sin repeticion")
	color.Println("<blue>[2]</> Permutacion con repeticion")
	color.Println("<blue>[3]</> Combinacion sin repeticion")
	color.Println("<blue>[4]</> Combinacion con repeticion")
	color.Println("<blue>[5]</> Variaciones sin repeticion")
	color.Println("<blue>[6]</> Variaciones con repeticion")
	color.Println("<blue>[7]</> Emparejamientos")
	color.Println("<red>[8]</> Salir")

	fmt.Print("Selecciona una opcion: ")
	fmt.Scanln(&option)

	//CONDICIÓN PARA RE UTILIZAR LAS VARIABLES 'n' Y 'r' EN LAS FUNCIONES QUE PIDAN TALES VARIABLES
	if ((option == 3) || (option == 4) || (option == 5) || (option == 6) || (option == 7)) && option != 0 {
    ENTERN:
		color.Print("Ingresa <suc>n</>:")
		fmt.Scanln(&input)
		val, err := strconv.Atoi(input)
		if err != nil {
			color.Info.Tips("El valor ingresado no es numerico, intenta nuevamente")
            goto ENTERN
		} else {
			n = val
		}
        ENTERR:
        if option == 7 {
            color.Print("Ingresa <suc>m</>:")
		    fmt.Scanln(&input)
        }else{
		    color.Print("Ingresa <suc>r</>:")
		    fmt.Scanln(&input)
        }
		val, err = strconv.Atoi(input)
		if err != nil {
			color.Info.Tips("El valor intresado no es numerico, intenta nuevamente")
            goto ENTERR
		} else {
			r = val
            m = val
		}

	}

	//ESTRUCTURA DE CONTROL SWITCH PARA LA SELECCION DE LAS FUNCIONES
	switch option {
	case 1:
	CASE1:
		color.Print("Ingresa <suc>n</>: ")
		fmt.Scanln(&input)
		val, err := strconv.Atoi(input)
		if err != nil {
			color.Info.Tips("El valor intresado no es numerico, intenta nuevamente")
			goto CASE1
		} else {
			n = val
			color.Printf("Permutaciones sin repeticion: <suc>%v</>\n", cp.PermutacionSinRepe(n))
			color.Info.Tips("Presione Enter para continuar al menú principal...")
			fmt.Scanln()
			LimpiarTerminal()
			goto MEMU
		}
	case 2:
	ADD:
		fmt.Print("Ingresa valor de un grupo: ")
		fmt.Scanln(&r)
		num = append(num, r)
		color.Print("Presione <suc>'y'</> para add otro grupo o <suc>'c'</> para continuar: ")
		fmt.Scanln(&res)
		if res == "y" {
			goto ADD
		}
        color.Printf("Permutaciones con repeticion: <suc>%v</>\n", cp.PermutacionConRepe(num...))
		color.Info.Tips("Presione Enter para continuar al menú principal...")
		fmt.Scanln()
		LimpiarTerminal()
		goto MEMU
	case 3:
		color.Printf("Combinacion Sin Repeticion: <suc>%v</>\n", cp.CombinacionSinRepe(n, r))
		color.Info.Tips("Presione Enter para continuar al menú principal...")
		fmt.Scanln()
		LimpiarTerminal()
		//LimpiarTerminal()
		goto MEMU
	case 4:
        color.Printf("Combinacion con repeticion: <suc>%v</>\n",cp.CombinacionConRepe(n, r))
		color.Info.Tips("Presione Enter para continuar al menú principal...")
		fmt.Scanln()
		LimpiarTerminal()
		goto MEMU
	case 5:
        color.Printf("Variaciones sin repeticion: <suc>%v</>\n",cp.VariacionesSinRepe(n, r))
		color.Info.Tips("Presione Enter para continuar al menú principal...")
		fmt.Scanln()
		LimpiarTerminal()
		goto MEMU
	case 6:
        color.Printf("Variaciones con repeticion: <suc>%v</>\n",cp.VariacionesConRepe(float64(n), float64(r)))
		color.Info.Tips("Presione Enter para continuar al menú principal...")
		fmt.Scanln()
		LimpiarTerminal()
		goto MEMU
	case 7:
        color.Printf("Emparejamientos: <suc>%v</>\n",cp.Emparejamiento(n, m))
		color.Info.Tips("Presione Enter para continuar al menú principal...")
		fmt.Scanln()
		LimpiarTerminal()
		goto MEMU
	case 8:
		color.Red.Println("Saliendo...")
        time.Sleep(time.Second * 1)

	default:
		color.Info.Tips("****** Selecciona una de las opciones en el menu ******")
		//fmt.Println()
        time.Sleep(time.Second * 2)
		LimpiarTerminal()
		goto MEMU

	}
}
//FUNCION QUE LIMPIA LA TERMINAL DEPENDIENDO SI SE ESTA EJECUTANDO EN LINUX O WINDOWS
func LimpiarTerminal() {
	limpiarW = exec.Command("cmd", "/c", "cls")
	limpiarW.Stdout = os.Stdout
    //limpiarL = exec.Command("clear")
	//limpiarL.Stdout = os.Stdout
	//limpiarL.Run()
	limpiarW.Run()
}
