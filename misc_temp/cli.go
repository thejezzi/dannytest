package main

import (
 "flag"
 "fmt"
)

func printf(msg string, values...interface{}) {
  str := fmt.Sprintf(msg, values...)
  fmt.Println(str)
}

func main() {
  piece := flag.String("piece", "", "The chess piece to move")
  pos := flag.String("pos", "", "The position where to put the chess piece FROM:TO")

  flag.Parse()

  printf("%v, %v", *piece, *pos)
}
