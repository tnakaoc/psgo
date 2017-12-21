package psgo
import "fmt"
import "strings"
import "math/rand"
import "strconv"
type Psgo struct{
	unary_func  map[string]func(float64)float64
	binary_func map[string]func(float64,float64)float64
	macro map[string]string
	stack []float64
	rindex int
	is_mac bool
	sc_dep uint
}
func (p *Psgo)SetUnary(m map[string](func(float64)float64)){
	p.unary_func=m
}
func (p *Psgo)SetBinary(m map[string]func(float64,float64)float64){
	p.binary_func=m
}
func (p *Psgo)SetMacro(m map[string]string){
	is_mac=false
	p.macro=m
}
func (p *Psgo)is_unary(s string) bool{
	_,ok:=p.unary_func[s]
	return ok
}
func (p *Psgo)is_binary(s string) bool{
	_,ok:=p.binary_func[s]
	return ok
}
func (p *Psgo)is_macro(s string) bool{
	_,ok:=p.macro[s]
	return ok
}
func (p *Psgo) Parse(text string) bool{
	if text=="exit"||text=="q" { return false }
	ind := strings.Index(text,"%")
	strings.Replace(text,"{","\t{\t",-1)
	strings.Replace(text,"}","\t}\t",-1)
	strings.Replace(text,"[","\t[\t",-1)
	strings.Replace(text,"]","\t]\t",-1)
	buf := strings.Fields(
		func() string{
			if ind!=-1 {
				return text[0:ind]
			}else{
				return text
			}
		}())
	for _,val := range(buf) {
		if val=="exit"||val=="q" {
			return false
		}
		if val[0]=='/' {
			is_mac=true
		}
		le:=len(p.stack)
		if p.is_unary(val) {
			if le>=1 {
				p.stack[le-1]=p.unary_func[val](p.stack[le-1])
			}
		} else if p.is_binary(val) {
			if le>=2 {
				p.stack[le-2]=p.binary_func[val](p.stack[le-2],p.stack[le-1])
				p.stack = p.stack[:le-1]
			}
		} else if p.is_macro(val) {
				p.Parse(p.macro[val])
		} else {
			switch val {
				case "p",".","i":
					if le != 0 {
						fmt.Println("#index\tvalue")
						for _i,_v := range(p.stack) {
							fmt.Println(int(2*(float64(p.rindex)-0.5))*(le*p.rindex-_i),"\t",_v)
						}
					} else {
						fmt.Println("stack is empty.")
					}
				case "rorder":
					p.rindex=1
				case "order":
					p.rindex=0
				case "clear":
					p.stack=nil
				case "pop":
					if le!=0 {
						p.stack=p.stack[:le-1]
					}
				case "npop":
					if le!=0 {
						n:=int(p.stack[le-1])
						if n<=0 {
							p.stack=p.stack[:le-1]
							continue
						}
						if le>n {
							p.stack=p.stack[:le-n-1]
						} else {
							p.stack=nil
						}
					}
				case "exch":
					if le>=2 {
						p.stack[le-1],p.stack[le-2]=p.stack[le-2],p.stack[le-1]
					}
				case "dup":
					if le>=1 {
						p.stack=append(p.stack,p.stack[le-1])
					}
				case "index":
					if le!=0 {
						ind := int(p.stack[le-1])
						if ind>0&&ind<le+1 {
							p.stack[le-1]=p.stack[le-ind-1]
						} else {
							p.stack=p.stack[:le-1]
						}
					}
				case "seq":
					if le>=3 {
						sta := p.stack[le-3]
						inc := p.stack[le-2]
						sto := p.stack[le-1]
						cnt :=1+int((sto-sta)/inc)
						if cnt<0 {
							fmt.Println("invalid syntax for \"seq\" operator.")
						}
						p.stack=p.stack[:le-3]
						for i:=0;i<cnt;i++ {
							p.stack=append(p.stack,sta+float64(i)*inc)
						}
					}
				case "rand":
					p.stack=append(p.stack,rand.Float64())
				case "irand":
					if le>=2 {
						sta := int64(p.stack[le-1])
						sto := int64(p.stack[le-2])
						if sta>sto { sta,sto = sto,sta }
						if sta==sto { continue }
						p.stack=append(p.stack,float64(sta+rand.Int63n(sto-sta)))
					}
				case "ndup":
					if le>=2 {
						n:=int(p.stack[le-1])
						if n<=0 {
							p.stack=p.stack[:le-1]
							continue
						}
						v:=p.stack[le-2]
						p.stack=p.stack[:le-1]
						for i:=0;i<n;i++ {
							p.stack=append(p.stack,v)
						}
					}
				case "copy":
					if le>=2 {
						w:=int(p.stack[le-1])
						if w<=0 {
							p.stack=p.stack[:le-1]
							continue
						}
						p.stack=p.stack[:le-1]
						le = len(p.stack)
						if le>=w {
							s:=p.stack[le-w:le]
							for _,v := range(s) {
								p.stack = append(p.stack,v)
							}
						}
					}
				case "roll":
					if le>=3 {
						w:=int(p.stack[le-2])
						s:=int(p.stack[le-1])
						p.stack=p.stack[:le-2]
						le = len(p.stack)
						if w>0&&le>=w {
							tmp:=p.stack[le-w:le]
							fmt.Println(tmp)
							for i:=0;i<w;i++ {
								p.stack[le-w+i]=tmp[(i+s+w)%w]
							}
						}
					}
				case "t",",":
					if le != 0 {
						fmt.Println(p.stack[le-1])
					} else {
						fmt.Println("stack is empty.")
					}
				case "count","$":
					p.stack=append(p.stack,(float64(len(p.stack))))
				case "len","l":
					fmt.Println(len(p.stack))
				default:
					valu,_e:=strconv.ParseFloat(val,64)
					if _e != nil {
						fmt.Printf("syntax error around \"%s\"\n",val)
					} else {
						p.stack=append(p.stack,valu)
					}
			}
		}
	}
	return true
}
