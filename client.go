package ts3

import "strconv"

func SendMsg(msg string, targetmode int, target int) Command {
	return Command{
		Command: "sendtextmessage",
		Params: map[string][]string{
			"targetmode": []string{strconv.Itoa(targetmode)},
			"target":     []string{strconv.Itoa(target)},
			"msg":        []string{msg},
		},
	}

}

func ClientList() Command {
	return Command{
		Command: "clientlist",
	}
}

func Kick(clients []string, reasonmsg string) Command {
	return Command{
		Command: "clientkick",
		Params: map[string][]string{
			"clid":      clients,
			"reasonid":  []string{"5"},
			"reasonmsg": []string{reasonmsg},
		},
	}
}

func ClientGetDbIdFromUid(uid string) Command {
	return Command{
		Command: "´clientgetdbidfromuid",
		Params: map[string][]string{
			"cluid": []string{uid},
		},
	}
}

func ClientGetIds(uid string) Command {
	return Command{
		Command: "´clientgetids",
		Params: map[string][]string{
			"cluid": []string{uid},
		},
	}
}

func ClientGetNameFromUid(uid string) Command {
	return Command{
		Command: "´clientgetnamefromuid",
		Params: map[string][]string{
			"cluid": []string{uid},
		},
	}
}

func ClientGetNameFromDbId(dbid int) Command {
	return Command{
		Command: "´clientgetnamefromdbid",
		Params: map[string][]string{
			"cldbid": []string{dbid},
		},
	}
}

func BanList() Command {
	return Command{
		Command: "banlist",
	}
}

func BanAdd(m map[string][]string) Command {
	return Command{
		Command: "banadd",
		Params:  m,
	}
}

func BanRemove(banid int) Command {
	return Command{
		Command: "bandel",
		Params: map[string][]string{
			"banid": []string{strconv.Itoa(banid)},
		},
	}
}

fucn BansRemoveAll() Command {
	return Command {
		Command: "bandelall",
	}
}

func (client *Client) WalkClients(step func(int, map[string]string)) (err error) {

	var res Response

	res, err = client.Exec(ClientList())
	if err != nil {
		return
	}

	for i := range res.Params {
		step(i, res.Params[i])
	}

	return
}

func (client *Client) WalkChannels(step func(int, map[string]string)) (err error) {

	var res Response

	res, err = client.Exec(ChannelList())
	if err != nil {
		return
	}

	for i := range res.Params {
		step(i, res.Params[i])
	}

	return
}
