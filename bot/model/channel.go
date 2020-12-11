package model

type Channel struct {
	UUID uint32
	Id   string
	Name string
}

func CreateNewChannel(id string) error {

}

func UpdateChannelName(name string) (uint32, error) {

}

func GenerateChannelLink(uuid uint32) string {

}

func GetChannelList(userId int) ([]Channel, error) {

}
