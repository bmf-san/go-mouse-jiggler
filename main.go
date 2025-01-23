package main

/*
#cgo LDFLAGS: -framework ApplicationServices
#include <ApplicationServices/ApplicationServices.h>

void moveCursor(int x, int y) {
    CGEventRef move = CGEventCreateMouseEvent(NULL, kCGEventMouseMoved, CGPointMake(x, y), 0);
    CGEventPost(kCGHIDEventTap, move);
    CFRelease(move);
}
*/
import "C"
import (
	"time"
)

func moveCursor(x, y int) {
	C.moveCursor(C.int(x), C.int(y))
}

func main() {
	for {
		moveCursor(100, 100)
		time.Sleep(1 * time.Second)
		moveCursor(200, 200)
		time.Sleep(1 * time.Second)
	}
}
