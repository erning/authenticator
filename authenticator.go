package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/xlzd/gotp"
)

func main() {
	var secrets []interface{}
	decoder := json.NewDecoder(os.Stdin)
	for {
		err := decoder.Decode(&secrets)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	default_color := "\033[39m"
	for {
		fmt.Print("\033[2J\033[1;1H\033[39m\033[49m")
		timeleft := 30 - time.Now().Unix()%30
		fmt.Printf("Time left: %d\n", timeleft)
		color := func() string {
			switch {
			case timeleft < 5:
				return "\033[31m"
			case timeleft < 10:
				return "\033[35m"
			default:
				return "\033[34m"
			}
		}()

		for _, v := range secrets {
			va := v.([]interface{})
			secret, desc := va[1].(string), va[0].(string)
			secret = strings.Replace(secret, " ", "", -1)
			secret = strings.ToUpper(secret)
			totp := gotp.NewDefaultTOTP(secret)
			token := totp.Now()
			fmt.Printf("%s%s%s: (%s)\n", color, token, default_color, desc)
		}
		time.Sleep(time.Second)
	}
	fmt.Printf("%v\n", secrets)
}
