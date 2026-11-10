package luhn

//Valid is a check to see if a number is valid according to the Luhn algorithm
func Valid(id string) bool {
    var sum int
    digits := 0
    for i := len(id) - 1; i >= 0; i-- {

        if id[i] == ' ' {
            continue
        }

        if (id[i] < '0' || id[i] > '9') {
            return false
        }

        num := int(id[i] - '0')

        if digits  % 2 != 0 {
            num *= 2

            if num > 9 {
                num -= 9
            }
        }
        sum += num
        digits ++
    }

    return digits  > 1 && sum % 10 == 0
}
