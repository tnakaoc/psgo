package main
import "fmt"
import "os"
import "bufio"
import "strings"
import "math"
import "math/rand"
import "strconv"
func main(){
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
	is_unary  := func(s string) bool{
		_,ok:=unifun[s];
		return ok
	}
	is_binary := func(s string) bool{
		_,ok:=binfun[s];
		return ok
	}
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
	scanner := bufio.NewScanner(fd)
	counter:=0
	echo:=false
	if interactive {
		fmt.Println(" ")
		fmt.Println("\tP S C G O")
		fmt.Println("\ta postscript like calculator")
		fmt.Println("\twritten by Go language")
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
			if val=="exit"||val=="q" {
				os.Exit(1)
			}
			le:=len(stack)
			if is_unary(val) {
				if le>=1 {
					stack[le-1]=unifun[val](stack[le-1])
				}
			} else if is_binary(val) {
				if le>=2 {
					stack[le-2]=binfun[val](stack[le-2],stack[le-1])
					stack = stack[:le-1]
				}
			} else {
				switch val {
					case "p",".","i":
						if le != 0 {
							for _i,_v := range(stack) {
								fmt.Println(_i,"\t:\t",_v)
							}
						} else {
							fmt.Println("stack is empty.")
						}
					case "clear":
						stack=nil
					case "pop":
						if le!=0 {
							stack=stack[:le-1]
						}
					case "npop":
						if le!=0 {
							n:=int(stack[le-1])
							if n<=0 {
								stack=stack[:le-1]
								continue
							}
							if le>n {
								stack=stack[:le-n-1]
							} else {
								stack=nil
							}
						}
					case "exch":
						if le>=2 {
							stack[le-1],stack[le-2]=stack[le-2],stack[le-1]
						}
					case "dup":
						if le>=1 {
							stack=append(stack,stack[le-1])
						}
					case "rand":
						stack=append(stack,rand.Float64())
					case "irand":
						if le>=2 {
							sta := int64(stack[le-1])
							sto := int64(stack[le-2])
							if sta>sto { sta,sto = sto,sta }
							stack=append(stack,float64(sta+rand.Int63n(sto-sta)))
						}
					case "ndup":
						if le>=2 {
							n:=int(stack[le-1])
							if n<=0 {
								stack=stack[:le-1]
								continue
							}
							v:=stack[le-2]
							stack=stack[:le-1]
							for i:=0;i<n;i++ {
								stack=append(stack,v)
							}
						}
					case "copy":
						if le>=2 {
							w:=int(stack[le-1])
							if w<=0 {
								stack=stack[:le-1]
								continue
							}
							stack=stack[:le-1]
							le = len(stack)
							if le>=w {
								s:=stack[le-w:le]
								for _,v := range(s) {
									stack = append(stack,v)
								}
							}
						}
					case "roll":
						if le>=3 {
							w:=int(stack[le-2])
							s:=int(stack[le-1])
							stack=stack[:le-2]
							le = len(stack)
							if le>=w {
								tmp:=stack[le-w:le]
								fmt.Println(tmp)
								for i:=0;i<w;i++ {
									stack[le-w+i]=tmp[(i+s+w)%w]
								}
							}
						}
					case "t",",":
						if le != 0 {
							fmt.Println(stack[le-1])
						} else {
							fmt.Println("stack is empty.")
						}
					case "count","$":
						stack=append(stack,(float64(len(stack))))
					case "len","l":
						fmt.Println(len(stack))
					case "noecho":
						echo = false
					case "echo":
						echo = true
					default:
						valu,_e:=strconv.ParseFloat(val,64)
						if _e != nil {
							fmt.Printf("syntax error around \"%s\"\n",val)
						} else {
							stack=append(stack,valu)
						}
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
