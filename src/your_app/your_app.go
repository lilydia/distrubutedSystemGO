package your_app

import (
  "github.com/tjarratt/babble"
)

func main() {
  babbler := babble.NewBabbler()
  println(babbler.Babble()) // excessive-yak-shaving (or some other phrase)

  // optionally set your own separator
  babbler.Separator = " "
  println(babbler.Babble()) // "hello from nowhere" (or some other phrase)

  // optionally set the number of words you want
  babbler.Count = 1
  println(babbler.Babble()) // antibiomicrobrial (or some other word)

  return
}