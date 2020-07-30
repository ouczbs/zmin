package zproto

import (
	"github.com/ouczbs/Zmin/engine/zlog"
	"github.com/ouczbs/Zmin/engine/znet"
	"github.com/ouczbs/Zmin/engine/zutil"
)

func sendPbMessage(proxy *UClientProxy , message IReflectMessage, wrap *UWrapMessage , messageType TMessageType)error{
	packet := znet.NewPacket()
	packet.WriteMessageType(messageType)
	out , err := Marshal(message)
	if err != nil{
		return err
	}
	wrap.Content = out
	buf , err := Marshal(wrap)
	if err != nil{
		return err
	}
	packet.AppendBytes(buf)
	err = proxy.SendPacket(packet)
	packet.Release()
	return err
}
func SendPbMessage(proxy *UClientProxy , message IReflectMessage, request * URequest)error{
	wrap := &UWrapMessage{
		Cmd: request.Cmd,
	}
	return sendPbMessage(proxy , message , wrap, request.MessageType)
}
func RequestMessage(proxy *UClientProxy , message IReflectMessage, request * URequest , handle FRequestHandle)error{
	sequence := zutil.IncSequence()
	wrap := &UWrapMessage{
		Cmd: request.Cmd,
		Request: sequence,
		Code:request.Code,
	}
	proxy.ReqHandleMaps[sequence] = handle
	return sendPbMessage(proxy , message , wrap, request.MessageType)
}
func ResponseMessage(proxy *UClientProxy , message IReflectMessage,request * URequest)error{
	wrap := &UWrapMessage{
		Cmd: request.Cmd,
		Response: request.Request,
		Code:request.Code,
	}
	return sendPbMessage(proxy ,message , wrap , request.MessageType)
}
func MakePbMessagePacket(message IReflectMessage , request * URequest)* UPacket{
	packet := znet.NewPacket()
	wrap := &UWrapMessage{
		Cmd: request.Cmd,
		Code:request.Code,
	}
	packet.WriteMessageType(request.MessageType)
	out , err := Marshal(message)
	if err != nil{
		zlog.Error(" MakePbMessagePacket Marshal message error " ,err)
		return nil
	}
	wrap.Content = out
	buf , err := Marshal(wrap)
	if err != nil{
		zlog.Error(" MakePbMessagePacket Marshal wrap message error " ,err)
		return nil
	}
	packet.AppendBytes(buf)
	return packet
}