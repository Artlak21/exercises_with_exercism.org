package collatzconjecture

import "errors"

var stepsGlobal int = 0

func CollatzConjecture(n int) (int, error) {
    x := n
    
    if n <= 0 {
        return 0, errors.New("error")
    }
    
	if x != 1 && x % 2 == 0 {
        x = x / 2
    } else if x != 1 {
        x = x * 3 + 1
    } else {
        steps := stepsGlobal
        stepsGlobal = 0
        return steps, nil
    }
    
    stepsGlobal++
    return CollatzConjecture(x)
}
