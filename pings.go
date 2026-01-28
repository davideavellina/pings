package main

import "bufio"
import "os"
import "fmt"
import "strings"
import "strconv"


type Host struct{
    IP string
    pings []float32
}

func (host Host) String() string {
    var pingsReturnStr strings.Builder
    for _,p := range host.pings {
        pingsReturnStr.WriteString(strconv.FormatFloat(float64(p), 'f', 2, 32))
        pingsReturnStr.WriteString(" ")
    }
    return host.IP + " : " + pingsReturnStr.String()
}

func addPing(host *Host, pingtime float32) {
    host.pings = append(host.pings, pingtime)
}

func averageHost(host Host) float32{
    var c,sum float32
    c = 0
    sum = 0
    for _,p := range host.pings {
        sum += p
        c++
    }
    return sum/c
}

func minHost(host Host) float32{
    var min float32
    min = host.pings[0]
    for _,p := range host.pings {
        if p < min {
            min = p
        }
    }
    return min
}

func maxHost(host Host) float32{
    var max float32
    max = host.pings[0]
    for _,p := range host.pings {
        if p > max {
            max = p
        }
    }
    return max
}

func main() {
    hostMap := make(map[string]Host)
    file,err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    
        line := scanner.Text()
        lineFields := strings.Fields(line)
        
        if strings.Contains(line,"ICMP"){
            ip := lineFields[10]
            
            _,ok := hostMap[ip]
            if !ok {
                hostMap[ip] = Host{IP: ip}
            }     
        }else {
            ip := lineFields[0]
            strping := lineFields[5]
            ping64,err := strconv.ParseFloat(strping,32)
            if err != nil {
                fmt.Println(err)
            }
            ping32 := float32(ping64) 
            
            //fmt.Printf("%s %0.2f\n",ip,ping32)
            trovato := false
            host,ok := hostMap[ip]
            if ok {
                
                for _,p := range host.pings{
                    if p == ping32 {
                        trovato = true
                        break //se il ping è già presente nella slice host.pings
                    }
                }
                if !trovato {
                    localHost := host
                    addPing(&localHost,ping32)//altrimenti lo aggiunge
                    hostMap[ip] = localHost
                }
                
            }else {
                hostMap[ip] = Host{IP: ip, pings: []float32{ping32}}
                
                //fmt.Printf("Host: %s Pings: %v\n", ip, host.pings)
            }
        }      
    }
    //fmt.Println(hostMap["192.168.142.242"].String()
    c := 0
    first := true
    var maxTot , minTot float32
    var minIp , maxIp string
    var sumAverHost float32
    
    for _,h := range hostMap{
       
        c ++
        if h.pings == nil {
           fmt.Printf("%s : %s\n", h.IP, "nil") 
        }else {
            fmt.Println(h)  
            
            sumAverHost += averageHost(h)

            if first { //alla prima iterazione inzializza minTot e maxTot con i risultati delle funzioni per il primo host considerato
                maxTot = maxHost(h)
                maxIp = h.IP
                
                minTot = minHost(h)
                minIp = h.IP
                
                first = false
            }else {
                if minHost(h) < minTot {
                    minTot = minHost(h)
                    minIp = h.IP
                }
                if maxHost(h) > maxTot {
                    maxTot = maxHost(h)
                    maxIp = h.IP
                }
            }
            
        }
    }
    
    fmt.Printf("numero di host esaminati : %d\n", c)
    
    fmt.Printf("Min pingtime: %0.4f , %s \n",minTot,minIp)

    fmt.Printf("Max pingtime: %0.4f , %s \n",maxTot,maxIp)
    
    fmt.Printf("Media pingtime: %f \n", sumAverHost/float32(c))
}
