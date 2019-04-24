package classfile

import "fmt"

/*
ClassFile {
	u4		maglic;
	u2		minor_version;
	u2		major_version;
	u2		constant_pool_count;
	cp_info	constant_pool[constant_pool_count];
	u2		access_flags;
	u2		this_class;
	u2		super_class;
	u2		interfaces_count;
	u2		interfaces[interface_count];
	u2		fields_count;
	field_info		fileld[fields_count];
	u2		methods_count;
	method_info		methods[methods_count];
	u2		attributes_count;
	attribute_info 		attributes[attributes_count]；
}
*/

type ClassFile struct {
	magic 			uint32
	minorVersion 	uint16
	majorVersion 	uint16
	constantPool  	ConstantPool
	accessFlags 	uint16
	thisClass 		uint16
	superClass 		uint16
	interfaces 		[]uint16
	fields 			[]*MemberInfo
	methods 		[]*MemberInfo
	attributes 		[]AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
		}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return 
}

func (self *ClassFile) read(read  *ClassReader) {
	self.readAndCheckMagic(read)
}

func (self *ClassFile) readAndCheckMagic(read  *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	} 
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0{
			return
		}
	}

	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) Fileds() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

func (self *ClassFile) InterfaceNames() []string {
	interfacesNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfacesNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfacesNames
}
