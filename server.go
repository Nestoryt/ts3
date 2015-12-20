package ts3

import (
	"strconv"
)

func GlobalMessage(msg string) Command {
	return Command{
		Command: "gm",
		Params: map[string][]string{
			"msg": []string{msg},
		},
	}

}

func ServerInfo() Command {
	return Command{
		Command: "serverinfo",
	}
}

func ServerProcessStop() Command {
	return Command{
		Command: "serverprocessstop",
	}
}

func ServerStop(sid int) Command {
	return Command{
		Command: "serverstop",
		Params: map[string][]string{
			"sid": []string{strconv.Itoa(sid)},
		},
	}

}

func ServerStart(sid int) Command {
	return Command{
		Command: "serverstart",
		Params: map[string][]string{
			"sid": []string{strconv.Itoa(sid)},
		},
	}

}

func LogAdd(msg string, loglevel int) Command {
	return Command{
		Command: "logadd",
		Params: map[string][]string{
			"loglevel": []string{strconv.Itoa(loglevel)},
			"msg":      []string{msg},
		},
	}

}
