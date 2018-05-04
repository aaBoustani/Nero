package main

import (
	"strconv"
	"log"
)

var NEROLIMIT = 7

func AddNero(user string, amount int) (int, error) {
	if amount > NEROLIMIT {
		log.Panic("The amount requested is over the limit")
		return -2, nil
	}

	hit, err := Find(user)
	if err != nil {
		log.Panic(err)
		return -1, err
	}

	if len(hit) < 1 {
		Add(user, strconv.Itoa(amount))
		return amount, nil
	}

	i, err := strconv.Atoi(hit)
	if err != nil {
		i = 0
	}

	total := i + amount
	Add(user, strconv.Itoa(total))
	return total, nil
}

func GetNero(user string) (int, error) {
	hit, err := Find(user)
	if err != nil {
		log.Panic(err)
		return -1, err
	}
	if hit == "" {
		return 0, nil
	}
	return strconv.Atoi(hit)
}
