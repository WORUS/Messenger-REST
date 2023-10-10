package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[req.ID] = &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}

	c.JSON(http.StatusOK, req)

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		//origin := r.Header.Get("Origin")
		return true
	},
}

// func (h *Handler) JoinRoom(c *gin.Context) {
// 	conn, err := upgrader.Upgrader(c.Writer, c.Request, nil)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Eror()})
// 		return
// 	}

// 	roomID := c.Param("roomId")
// 	clientID := c.Query("userId")
// 	username := c.Query("username")
// 	///ws/JoinRoom/:roomId?userId=1&username=user
// 	cl := &Client{
// 		Conn:     conn,
// 		Message:  make(chan *Message, 10),
// 		ID:       clientID,
// 		RoomID:   roomID,
// 		Username: username,
// 	}

// 	m := &Message{
// 		Content:  "A new user has joined the room",
// 		RoomID:   roomID,
// 		Username: username,
// 	}

// 	h.hub.Register <- cl

// 	h.hub.Broadcast <- m

// }
