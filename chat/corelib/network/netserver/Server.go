package netserver

var (
	_Server IServer
)

func Initialize() {
	_Server = new(ChatServer)
}

func Run() {
	_Server.Run()
}
