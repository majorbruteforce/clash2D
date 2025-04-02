package exception

import "log"

func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
