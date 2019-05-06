package base

import "jvmgo/ch05/rtda"

func Branch(frame *rtda.Frame, Offset int)  {
	pc := frame.Thread().PC()
	pc += Offset
	frame.SetNextPC(pc)
}