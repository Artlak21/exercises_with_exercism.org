package clock

import "fmt"

type  Clock  struct {
    hour int
	minute int
}

func  New ( h , m  int ) Clock {
    totalMinutes := h * 60 + m
    totalMinutes = ( 1440 + totalMinutes % 1440 ) % 1440
    
    h = totalMinutes / 60
    m = totalMinutes % 60
    
	return  Clock {
        hour: h,
        minute: m,
    }
}

func ( c  Clock ) Add( m  int ) Clock {
    return New(c.hour, c.minute + m)
}

func ( c  Clock ) Subtract( m  int ) Clock {
    return New(c.hour, c.minute - m)
}

func ( c  Clock ) String() string {
    return fmt.Sprintf( "%02d:%02d" , c.hour, c.minute)
}
