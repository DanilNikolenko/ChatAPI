package server

import (
	"ChatAPI/chats"
	"ChatAPI/chats/delivery/httpChats"
	chatsPSQL "ChatAPI/chats/repository/postgresql"
	chatsUC "ChatAPI/chats/usecase"
	"ChatAPI/messages"
	"ChatAPI/messages/delivery/httpMessages"
	messagesPSQL "ChatAPI/messages/repository/postgresql"
	messagesUC "ChatAPI/messages/usecase"
	"ChatAPI/services"
	"ChatAPI/users"
	"ChatAPI/users/delivery/httpUsers"
	"ChatAPI/users/repository/postgresql"
	"ChatAPI/users/usecase"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
)

type App struct {
	httpServer *gin.Engine

	UsersUC    users.UseCaseUsers
	ChatsUC    chats.UseCaseChats
	MessagesUC messages.UseCaseMessages
}

func NewApp() *App {
	Server := gin.Default()
	connDB := services.ConnToPSQL()
	repoUsers := postgresql.NewPSQLRepository(connDB)
	repoChats := chatsPSQL.NewPSQLRepositoryChats(connDB)
	repoMessages := messagesPSQL.NewMessagesPSQLRepository(connDB)

	return &App{
		httpServer: Server,
		UsersUC:    usecase.NewUsersUseCase(repoUsers),
		ChatsUC:    chatsUC.NewChatsUseCase(repoChats),
		MessagesUC: messagesUC.NewMessagesUseCase(repoMessages),
	}
}

func (a *App) Run(port string) error {
	httpUsers.RegisterHTTPEndpointsUsers(a.httpServer, a.UsersUC)
	httpChats.RegisterHTTPEndpointsChats(a.httpServer, a.ChatsUC)
	httpMessages.RegisterHTTPEndpointsMessages(a.httpServer, a.MessagesUC)

	err := a.httpServer.Run(port)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
