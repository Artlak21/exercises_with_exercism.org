package raindrops

import "strconv"

func Convert(number int) string {
    if number % 3 != 0 && number % 5 != 0 && number % 7 != 0  {
        return strconv.Itoa(number)
    }
    
    answer := ""
    if number % 3 == 0 {answer += "Pling"}
    if number % 5 == 0 {answer += "Plang"}
    if number % 7 == 0 {answer += "Plong"}
    return answer
}
