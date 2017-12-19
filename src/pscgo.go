package main
import "fmt"
import "os"
import "psgo"
import "bufio"
//import "strings"
import "math"
//import "math/rand"
//import "strconv"
func main(){
	fnm,interactive := func()(string,bool){
		if len(os.Args)==1 { return "/dev/stdin",true }
		return os.Args[1],false
	}()
	fd,err := os.Open(fnm)
	if err != nil {
		fmt.Println("file open error")
		fmt.Println(err)
		return
	}
	defer fd.Close()
	unifun := map[string]func(float64)float64{
		"sqrt" :func(x float64)float64{ return math.Sqrt(x)    },
		"abs"  :func(x float64)float64{ return math.Abs(x)     },
		"cos"  :func(x float64)float64{ return math.Cos(x)     },
		"sin"  :func(x float64)float64{ return math.Sin(x)     },
		"tan"  :func(x float64)float64{ return math.Tan(x)     },
		"acos" :func(x float64)float64{ return math.Acos(x)    },
		"asin" :func(x float64)float64{ return math.Asin(x)    },
		"atan" :func(x float64)float64{ return math.Atan(x)    },
		"exp2" :func(x float64)float64{ return math.Exp2(x)    },
		"exp"  :func(x float64)float64{ return math.Exp(x)     },
		"log"  :func(x float64)float64{ return math.Log(x)     },
		"log10":func(x float64)float64{ return math.Log10(x)   },
		"log2" :func(x float64)float64{ return math.Log2(x)    },
		"int"  :func(x float64)float64{ return float64(int(x)) },
		"neg"  :func(x float64)float64{ return -x},
	}
	binfun := map[string]func(float64,float64)float64{
		"add":func(x float64,y float64)float64{ return x+y },
		"sub":func(x float64,y float64)float64{ return x-y },
		"mul":func(x float64,y float64)float64{ return x*y },
		"div":func(x float64,y float64)float64{ return x/y },
		"hypot":func(x float64,y float64)float64{ return math.Hypot(x,y)},
		"pow"  :func(x float64,y float64)float64{ return math.Pow(x,y)},
	}
	macro := map[string]string{
		"nop":"",
	}
	var psgo psgo.Psgo
	psgo.SetUnary(unifun)
	psgo.SetBinary(binfun)
	psgo.SetMacro(macro)
	scanner := bufio.NewScanner(fd)
	counter:=0
	if interactive {
		fmt.Println(" ")
		fmt.Println("\tP S C G O")
		fmt.Println("\ta postscript like calculator")
		fmt.Println("\twritten by Go language")
		fmt.Println(" ")
		fmt.Print("PSCGO[",counter,"]>")
	}
	for scanner.Scan() {
		ltext := scanner.Text()
		if !psgo.Parse(ltext) { break }
		if interactive {
			if len(ltext) != 0 { counter++ }
			fmt.Print("PSCGO[",counter,"]>")
		}
	}
	return
}
