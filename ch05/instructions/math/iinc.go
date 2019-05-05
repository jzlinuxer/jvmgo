package math

import "jvmgo/ch05/rtda"
import (
	"jvmgo/ch05/instructions/base"
)

type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FecthOperands(reader *base.BytecodeReader)  {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame)  {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}
