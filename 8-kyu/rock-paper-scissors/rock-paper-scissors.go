package kata
​
​
func Rps(p1, p2 string) string {
  switch {
    case p1 == "scissors" && p2 == "paper":
    return "Player 1 won!"
    case p1 == "rock" && p2 == "scissors":
    return "Player 1 won!"
    case p1 == "paper" && p2 == "rock" :
    return "Player 1 won!"
    case p1 == p2 :
    return "Draw!"
    default :
    return "Player 2 won!"
  } 
}
​