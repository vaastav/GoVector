package brokervec

import (

)

// Message abstraction to hold different types of message
type Message interface {
	GetMessage()	string
}

type LocalMessage struct {
	pid string // pid
	vclock string
	message string
}

func (lm *LocalMessage) GetMessage() string {
		
	return lm.pid + " " + lm.vclock + "\n" + lm.message + "\n"
}

type NetworkMessage struct {
	pid string // pid
	vclock string
	message string
}

func (nm *NetworkMessage) GetMessage() string {
		
	return nm.pid + " " + nm.vclock + "\n" + nm.message + "\n"
}