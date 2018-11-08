package rect_test

import "testing"

import(
  "fmt"
  "rect"
)

func TestRect(t *testing.T){ // Test prefix is required
  fmt.Println("In test rect")
  rect.Print()
}

func ExamplePrefix() {  // Example prefix is required
  fmt.Println("hello")
  rect.Print() // Below Output comments are required, else Go compiles test but skips the run
  // Output:
  // hello
  // area: 50
}