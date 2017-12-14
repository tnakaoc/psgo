package main
import "fmt"
import "os"
import "bufio"
import "strings"
import "math/big"
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
	var prec uint = 256
	if interactive {
		fmt.Print("psgo[",counter,"]>")
	}
	stack := make([]*big.Float,0,30)
	for scanner.Scan() {
		if scanner.Text()=="exit"||scanner.Text()=="q" { break }
		buf := strings.Fields(scanner.Text())
		for _,val := range(buf) {
			if val=="noecho" {
				echo = false
			}
			if val=="echo" {
				echo = true
			}
			if val=="p"||val=="." {
				if len(stack) != 0 {
					for _i,_v := range(stack) {
						fmt.Println(_i,"\t:\t",_v)
					}
				} else {
					fmt.Println("stack is empty.")
				}
			}
			if val=="clear" {
				stack=nil
			}
			if val=="pop" {
				if len(stack)!=0 {
					stack=stack[:len(stack)-1]
				}
			}
			if val=="exch" {
				if len(stack)>=2 {
					stack[len(stack)-1],stack[len(stack)-2]=stack[len(stack)-2],stack[len(stack)-1]
				}
			}
			if val=="setprec" {
				if len(stack)!=0 {
					prec_,_ := (stack[len(stack)-1]).Int64()
					prec=uint(prec_)
					stack = stack[:len(stack)-1]
				}
			}
			if val=="t"||val=="," {
				if len(stack) != 0 {
					fmt.Println(stack[len(stack)-1])
				} else {
					fmt.Println("stack is empty.")
				}
			}
			if val=="count"||val=="$" {
				stack=append(stack,big.NewFloat(float64(len(stack))))
			}
			if val=="len"||val=="l" {
				fmt.Println(len(stack))
			}
			if echo { fmt.Println(val) }
			valu,_e:=strconv.ParseFloat(val,64)
			if _e == nil {
				stack=append(stack,big.NewFloat(valu))
			}
		}
		if interactive {
			if len(buf) != 0 { counter++ }
			fmt.Print("psgo[",counter,"]>")
		}
	}
	return
}
