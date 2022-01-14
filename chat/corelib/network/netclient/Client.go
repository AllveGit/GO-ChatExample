package netclient

var (
	_Client IClient
)

func Initialize() {
	_Client = new(ChatClient)
}

func Run() {
	_Client.Run()
}
