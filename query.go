package ts3

import (
	"strconv"
)

func Login(user, pass string) (c Command) {
	return Command{
		Command: "login",
		Params: map[string][]string{
			"client_login_name":     []string{user},
			"client_login_password": []string{pass}},
	}
}

func Select(index int) Command {
	return Command{
		Command: "use",
		Params: map[string][]string{
			"sid": []string{strconv.Itoa(index)}},
	}
}

func SelectByPort(index int) Command {
	return Command{
		Command: "use",
		Params: map[string][]string{
			"port": []string{strconv.Itoa(index)}},
	}
}

func Version() (c Command) {
	return Command{
		Command: "version",
	}
}

func Nickname(name string) Command {
	return Command{
		Command: "clientupdate",
		Params: map[string][]string{
			"client_nickname": []string{name},
		},
	}
}
