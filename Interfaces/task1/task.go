package main

import (
	"fmt"           
	"os"            
)

func readTask() (value1, value2, operation interface{}) {

	return 5.0, 5.6, "/" 
	
	}

func main() {
	value1, value2, operation := readTask() 
    
    switch v := value1.(type) {
    case float64:
        
        break
    case bool:
        fmt.Printf("value=%t: %T",v,v)
        os.Exit(0)
    default:
        fmt.Printf("value=%v: %T",v,v)
        os.Exit(0)
        
    }
        switch v := value2.(type) {
        case float64:
            
            break
        case bool:
            fmt.Printf("value=%t: %T",v,v)
            os.Exit(0)
        default:
            fmt.Printf("value=%v: %T",v,v)
            os.Exit(0)
        }
    
    
        switch v := operation.(type) {
        case string:
            if v == "+" {
                fmt.Printf("%.4f",(value1.(float64)+value2.(float64)))
                 
            } else if v == "-" {
                fmt.Printf("%.4f",(value1.(float64)-value2.(float64)))
                
            } else if v == "*" {
                fmt.Printf("%.4f",(value1.(float64)*value2.(float64)))
                
            } else if v == "/" {
                fmt.Printf("%.4f",(value1.(float64)/value2.(float64)))
                
            }
        case bool :
            fmt.Printf("%t",v)
            os.Exit(0)
        default:
            fmt.Printf("%v",v)
            os.Exit(0)
        }
	}
    