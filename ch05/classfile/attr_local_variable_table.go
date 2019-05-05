package classfile

type LocalVariableTableAttribute struct {
	localVariableTable   []*LocalVarialbeTableEntry
}

type LocalVarialbeTableEntry struct {
	startPc		uint16
	length		uint16
	nameIndex	uint16
	descriptorIndex 	uint16
	index 				uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader)  {
	localVariableTableLength := reader.readUint16()
	self.localVariableTable = make([]*LocalVarialbeTableEntry, localVariableTableLength)
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &LocalVarialbeTableEntry{
			startPc:	reader.readUint16(),
			length: 	reader.readUint16(),
			nameIndex: 	reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:		reader.readUint16(),
		}
 	}
}