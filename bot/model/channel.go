package model

type Channel struct {
	UUID    uint32
	Id      string
	Name    string
	Deleted bool
}

func CreateNewChannel(id string) error {

}

func UpdateChannelName(name string) (uint32, error) {

}

func GenerateChannelLink(uuid uint32) string {

}

func GetChannelList(userId int) ([]Channel, error) {

}

func GetChannel(uuid uint32) (Channel, error) {

}

func RevokeChannel(uuid uint32) (Channel, error) {

}

func DeleteChannel(uuid uint32) error {

}
