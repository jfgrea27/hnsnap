package utils

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

func CheckErr(err error, msg string) {
	if err != nil {
		e := fmt.Sprintf("%s - reason %s", msg, err)
		panic(e)
	}
}

func CheckReqEnv(reqEnvs []string) {
	log.Debug().Msg("Checking required Env.")

	for _, env := range reqEnvs {
		_, ok := os.LookupEnv(env)
		if !ok {
			panic(fmt.Sprintf("%s is not present", env))
		}
	}
}
