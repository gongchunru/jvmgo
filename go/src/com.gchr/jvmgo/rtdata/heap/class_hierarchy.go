package heap

/*
也就是说，在三种情况下，
	S类型的引用值可以赋值给T类型：
	S 和T是同一类型；T是类且S是T的子类；
	或者T是接口且S实现了T接 口。
 */
func (self *Class) isAssignableFrom(other *Class) bool {

	s, t := other, self

	if s == t {
		return true
	}

	if !s.IsArray() {

	}

}

func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.isSubClassOf(self)
}

// 判断other class 是否是self class的直接或间接超类
func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) isImplements(iface *Class) bool  {
	for c := self; c != nil; c =  c.superClass{
		for _, i := range c.interfaces{
			if  i == iface || i.isSubInterfaceOf(iface){
				return true
			}
		}
	}
	return false
}

func (self *Class) isSubInterfaceOf (iface *Class) bool{
	for _, superInterface := range  self.interfaces{
		if superInterface == iface || superInterface.isSubInterfaceOf(iface){
			return true
		}
	}
	return false
}

func (self *Class) isSuperInterfaceOf(iface *Class) bool  {
	return iface.isSubInterfaceOf(self)
}