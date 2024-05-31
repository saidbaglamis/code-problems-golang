package main
import "fmt"

func QuestionsMarks(str string) string {

var counter int  = 0
var opened = false
var firstNumber = 0
  for _, ch := range (str) {

    if(ch >= 48 && ch <= 57 && !opened){
        opened = true
        firstNumber = int(ch - 48)
        continue
    }
    if (opened){
      if (ch == '?'){
        counter++
      } else if (ch >= 48 && ch <= 57){
        if (counter >= 3 && firstNumber + int(ch - 48) == 10){
          return "true"
        } else {
          counter = 0
          firstNumber = 0
          opened = false
          continue
        }
      }
    }
  }  
  return "false"

}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(QuestionsMarks(readline()))

}
