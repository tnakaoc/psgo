package main
import "fmt"
import "os"
import "bufio"
import "strings"
import "math"
import "strconv"
func main(){
	fd,err := os.Open(func()string{
		if len(os.Args)==1 { return "/dev/stdin" }
		return os.Args[1]
	}())
	interactive := func()bool{
		if len(os.Args)==1 { return true }
		return false
	}()
	if err != nil {
		fmt.Println("file open error")
		fmt.Println(err)
		return
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	counter:=0
	echo:=false
	if interactive {
		fmt.Println(" ")
		fmt.Println("\tP S C G O")
		fmt.Println("\ta postscript like calculator")
		fmt.Println("\tpowered by Go language")
		fmt.Println(" ")
		fmt.Print("pscgo[",counter,"]>")
	}
	stack := make([]float64,0,30)
	for scanner.Scan() {
		if scanner.Text()=="exit"||scanner.Text()=="q" { break }
		ind := strings.Index(scanner.Text(),"%")
		buf := strings.Fields(
			func() string{
				if ind!=-1 {
					return scanner.Text()[0:ind]
				}else{
					return scanner.Text()
				}
			}())
		for _,val := range(buf) {
			switch val {
				case "exit","q":
					os.Exit(1)
				case "noecho":
					echo = false
				case "echo":
					echo = true
				case "p",".","i":
					if len(stack) != 0 {
						for _i,_v := range(stack) {
							fmt.Println(_i,"\t:\t",_v)
						}
					} else {
						fmt.Println("stack is empty.")
					}
				case "clear":
					stack=nil
				case "pop":
					if len(stack)!=0 {
						stack=stack[:len(stack)-1]
					}
				case "exch":
					if len(stack)>=2 {
						stack[len(stack)-1],stack[len(stack)-2]=stack[len(stack)-2],stack[len(stack)-1]
					}
				case "div":
					if len(stack)>=2 {
						stack[len(stack)-2]=(stack[len(stack)-2]/stack[len(stack)-1])
						stack = stack[:len(stack)-1]
					}
				case "mul":
					if len(stack)>=2 {
						stack[len(stack)-2]=(stack[len(stack)-2]*stack[len(stack)-1])
						stack = stack[:len(stack)-1]
					}
				case "add":
					if len(stack)>=2 {
						stack[len(stack)-2]=(stack[len(stack)-2]+stack[len(stack)-1])
						stack = stack[:len(stack)-1]
					}
				case "sub":
					if len(stack)>=2 {
						stack[len(stack)-2]=(stack[len(stack)-2]-stack[len(stack)-1])
						stack = stack[:len(stack)-1]
					}
				case "pow":
					if len(stack)>=2 {
						stack[len(stack)-2]=math.Pow(stack[len(stack)-2],stack[len(stack)-1])
						stack = stack[:len(stack)-1]
					}
				case "hypot":
					if len(stack)>=2 {
						stack[len(stack)-2]=math.Hypot(stack[len(stack)-2],stack[len(stack)-1])
						stack = stack[:len(stack)-1]
					}
				case "neg":
					if len(stack)>=1 {
						stack[len(stack)-1]=-stack[len(stack)-1]
					}
				case "sin":
					if len(stack)>=1 {
						stack[len(stack)-1]=math.Sin(stack[len(stack)-1])
					}
				case "cos":
					if len(stack)>=1 {
						stack[len(stack)-1]=math.Cos(stack[len(stack)-1])
					}
				case "tan":
					if len(stack)>=1 {
						stack[len(stack)-1]=math.Tan(stack[len(stack)-1])
					}
				case "asin":
					if len(stack)>=1 {
						stack[len(stack)-1]=math.Asin(stack[len(stack)-1])
					}
				case "acos":
					if len(stack)>=1 {
						stack[len(stack)-1]=math.Acos(stack[len(stack)-1])
					}
				case "atan":
					if len(stack)>=1 {
						stack[len(stack)-1]=math.Atan(stack[len(stack)-1])
					}
				case "sqrt":
					if len(stack)>=1 {
						stack[len(stack)-1]=math.Sqrt(stack[len(stack)-1])
					}
				case "int":
					if len(stack)>=1 {
						stack[len(stack)-1]=float64(int(stack[len(stack)-1]))
					}
				case "abs":
					if len(stack)>=1 {
						stack[len(stack)-1]=math.Abs(stack[len(stack)-1])
					}
				case "exp2":
					if len(stack)>=1 {
						stack[len(stack)-1]=math.Exp2(stack[len(stack)-1])
					}
				case "exp":
					if len(stack)>=1 {
						stack[len(stack)-1]=math.Exp(stack[len(stack)-1])
					}
				case "log":
					if len(stack)>=1 {
						stack[len(stack)-1]=math.Log(stack[len(stack)-1])
					}
				case "log10":
					if len(stack)>=1 {
						stack[len(stack)-1]=math.Log10(stack[len(stack)-1])
					}
				case "log2":
					if len(stack)>=1 {
						stack[len(stack)-1]=math.Log2(stack[len(stack)-1])
					}
				case "t",",":
					if len(stack) != 0 {
						fmt.Println(stack[len(stack)-1])
					} else {
						fmt.Println("stack is empty.")
					}
				case "count","$":
					stack=append(stack,(float64(len(stack))))
				case "dup":
					if len(stack)>=1 {
						stack=append(stack,stack[len(stack)-1])
					}
				case "len","l":
					fmt.Println(len(stack))
				default:
					valu,_e:=strconv.ParseFloat(val,64)
					if _e != nil {
						fmt.Printf("syntax error around \"%s\"\n",val)
					} else {
						stack=append(stack,valu)
					}
			}
			if echo { fmt.Println(val) }
		}
		if interactive {
			if len(buf) != 0 { counter++ }
			fmt.Print("pscgo[",counter,"]>")
		}
	}
	return
}
