package main

import (
	"log"
	"strconv"
)

var NEROLIMIT = 7

func AddNero(user string, amount int) (int, error) {
	remNero, err := GetRemaining(user)
	if amount > NEROLIMIT || amount > remNero {
		log.Panic("The amount requested is over the limit")
		return -2, nil
	}

	go UpdateRemaining(user, remNero - amount)

	hit, err := db.FindOne(user)
	if err != nil {
		log.Panic(err)
		return -1, err
	}

	if len(hit) < 1 {
		db.UpdateTxn(user, strconv.Itoa(amount))
		return amount, nil
	}

	i, err := strconv.Atoi(hit)
	if err != nil {
		i = 0
	}

	total := i + amount
	db.UpdateTxn(user, strconv.Itoa(total))
	return total, nil
}

func GetNero(user string) (int, error) {
	hit, err := db.FindOne(user)
	if err != nil {
		log.Panic(err)
		return -1, err
	}
	if hit == "" {
		return 0, nil
	}
	return strconv.Atoi(hit)
}

func GetRemaining(user string) (int, error) {
	hit, err := rem.FindOne(user)
	if err != nil {
		log.Panic(err)
		return -1, err
	}
	if hit == "" {
		return NEROLIMIT, nil
	}
	return strconv.Atoi(hit)
}

func UpdateRemaining(user string, amount int) error {
	return rem.Update(user, strconv.Itoa(amount))
}

func ResetAllRemaining() {
	rem.ResetAll(strconv.Itoa(NEROLIMIT))
}