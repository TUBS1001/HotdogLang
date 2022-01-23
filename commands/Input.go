package commands


import (
  "fmt"
	"reflect"
	"strconv"
  "bufio"
  "os"
  "strings"
)

func Input(comand string){
  splitedCmd := strings.Split(command, " ")
  
  text, _ := reader.ReadString('\n')
  variables[splitedCmd[1]] = text;
}

