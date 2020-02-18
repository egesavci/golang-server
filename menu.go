package main

import (  
    "fmt"
    
)
var usr1_arr [5]string
var usr2_arr [5]string
func main() {  
    var choice int
    //reader := bufio.NewReader(os.Stdin)
    fmt.Println("Pick 1 to enter a username\nPick 2,3 or 4.\nType '00' to quit.")
    _, err := fmt.Scanf("%d",&choice)
    if choice == 00{
	return
	}
    //var username string
    //var usr1_arr [5]string
    //  var usr2_arr [5]string
    //var usr1 string
    //var usr2 string
    var usr string
    if err !=nil {
    fmt.Println(err)
    }
    // fmt.Println(text)

    switch choice {
    case 1:
	//fmt.Println(usr1_arr[0])
        fmt.Println("Please type the username")
        _,err:=fmt.Scanf("%s",&usr)
	if len(usr1_arr[0]) <= 0{
	usr1_arr[0] = usr
	fmt.Println("user 1:",usr1_arr[0]," has created.")
	}else if usr!=usr2_arr[0]{
	usr2_arr[0] = usr
	fmt.Println("user 2:",usr2_arr[0]," has created.")
	}else{
	fmt.Println("an error occured.")
	}
	if err !=nil {
        fmt.Println(err)
        }
        main()
    case 2:
        fmt.Println("2 bebe")
        main() 
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 5:
        fmt.Println("Pinky")

    }
}

