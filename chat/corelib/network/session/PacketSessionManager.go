package session

import (
	"chat/corelib/logger"
	. "chat/types"
	"net"
	"sync"
	"sync/atomic"
)

type ISessionManager interface {
	GenerateSessionID() int
	GenerateSession(InConn net.Conn) (ISession, error)
	Shutdown()
	BroadCast(InSessionID int, InMessagePacket MessagePacket)
}

type PacketSessionManager struct {
	sessions      map[int]*PacketSession
	locker        *sync.RWMutex
	nextSessionID int32
}

func NewSessionManager() *PacketSessionManager {
	sessionMgr := &PacketSessionManager{
		sessions:      make(map[int]*PacketSession),
		locker:        &sync.RWMutex{},
		nextSessionID: 0,
	}

	return sessionMgr
}

func (self *PacketSessionManager) GenerateSessionID() int {
	generateID := atomic.AddInt32(&self.nextSessionID, 1)
	return int(generateID)
}

func (self *PacketSessionManager) GenerateSession(InConn net.Conn) (ISession, error) {
	sessionID := self.GenerateSessionID()

	session, makeSessionErr := NewPacketSession(sessionID, InConn)
	if makeSessionErr == nil {
		session.SessionManager = self
	}

	self.locker.Lock()
	defer self.locker.Unlock()

	self.sessions[sessionID] = session

	return session, makeSessionErr
}

func (self *PacketSessionManager) Remove(InSessionID int) (*PacketSession, bool) {
	self.locker.Lock()
	defer self.locker.Unlock()

	session, isFind := self.sessions[InSessionID]
	if isFind == false {
		return nil, false
	}

	session.Shutdown()
	delete(self.sessions, InSessionID)

	logger.Text("session[%d] removed.", InSessionID)
	return session, true
}

func (self *PacketSessionManager) Shutdown() {
	self.locker.Lock()
	defer self.locker.Unlock()

	for sessionID, session := range self.sessions {
		session.Shutdown()
		delete(self.sessions, sessionID)
	}
}

func (self *PacketSessionManager) GetAll() map[int]*PacketSession {
	return self.sessions
}

func (self *PacketSessionManager) BroadCast(InSessionID int, InMessagePacket MessagePacket) {
	for sessionID, session := range self.sessions {
		if sessionID == InSessionID {
			continue
		}
		if session == nil {
			continue
		}

		session.Send(InMessagePacket)
	}
}
