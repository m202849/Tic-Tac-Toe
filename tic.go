package main

import (
	"fmt"
	"strconv"
)

type Board struct {
	tokens [9]int
}

func (b *Board) InitialBoard() {
    var i int
    
	for i = 0; i < 9; i++ {
		b.tokens[i] = 0
	}
}


func (b *Board) put(x, y, u int) {
	if u == 1 || u == 2 {
		b.tokens[3*x+y] = u
	} else {
		fmt.Println("bad input\n")
	}
}

func (b *Board) get(x, y int) string {
	if b.tokens[3*x+y] == 1 {
		return "o"
	} else if b.tokens[3*x+y] == 2 {
		return "x"
	} else {
		return "."
	}
}

func (b *Board) print() {
    var i int
    var j int
    
	for i = 0; i < 3; i++ {
	    for j = 0; j < 3; j++ {
            fmt.Printf(b.get(i,j))
	    }
	    fmt.Printf("\n")
	}
}

func (b *Board) CheckWin() int {
    if (b.tokens[0]==b.tokens[1]&&b.tokens[1]==b.tokens[2]) || (b.tokens[0]==b.tokens[3]&&b.tokens[3]==b.tokens[6]) || (b.tokens[0]==b.tokens[4]&&b.tokens[4]==b.tokens[8]){
	    return b.tokens[0]
    }else if (b.tokens[3]==b.tokens[4]&&b.tokens[4]==b.tokens[5]) || (b.tokens[1]==b.tokens[4]&&b.tokens[4]==b.tokens[7]) || (b.tokens[2]==b.tokens[4]&&b.tokens[4]==b.tokens[6]){
        return b.tokens[4]
    }else if (b.tokens[6]==b.tokens[7]&&b.tokens[7]==b.tokens[8]) || (b.tokens[2]==b.tokens[5]&&b.tokens[5]==b.tokens[8]){
        return b.tokens[8]
    }else{
        return 0
    }
}

func main(){
    var sx string
    var sy string
    
    var b Board
    var temp int
    var i int
    
    b.InitialBoard()
    for i = 0; i < 9; i++ {
        temp = b.CheckWin()
        if temp == 0{
            fmt.Println("Player ",i%2+1,": Input (x,y) ")
            fmt.Scanln(&sx, &sy)
            index_x, err := strconv.Atoi(sx)
            if err != nil {
                fmt.Println("input err\n")
                return
            }
            index_y, err := strconv.Atoi(sy)
            if err != nil {
                fmt.Println("input err\n")
                return
            }
            b.put(index_x,index_y,i%2+1)
            b.print()
        }else if temp == 1{
            fmt.Println("Player 1 \"o\" won\n")
            return
        }else if temp == 2{
            fmt.Println("Player 2 \"x\" won\n")
            return
        }
    }
    
    fmt.Println("draw\n")
    return
}