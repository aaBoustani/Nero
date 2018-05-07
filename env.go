package main

type Env struct {
	SlackAccessToken string
	SlackVerToken string
}

func InitEnv() *Env {
	return &Env{
		SlackAccessToken: "TOKEN-STUFF",
		SlackVerToken: "TOKEN-STUFF",
	}
}
