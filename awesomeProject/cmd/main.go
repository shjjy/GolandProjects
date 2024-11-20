package main

import "awesomeProject/internal"

func main() {
	//basics()
	//flowControl()
	//moreTypes()
	//method()
	//interfacePrac()
	//generics()
	concurrent()
}

//base:pkg, var, func
func basics() {
	internal.Hello()
	internal.MathRand()
	internal.Add()
	internal.Swap()
	internal.Split()
	internal.Var()
	internal.VarInit()
	internal.ShortVar()
	internal.BType()
	internal.Zero()
	internal.Conversion()
	internal.AutoVar()
	internal.ConstExample()
}

func flowControl() {
	internal.ForLoop()
	internal.WhileLoop()
	internal.InfiniteLoop()
	internal.IfExample()
	internal.IfCondExample()
	internal.IfElse()
	internal.LoopPractice()
	internal.Switch()
	internal.SwitchOrder()
	internal.SwitchWithoutCond()
	internal.Defer()
	internal.DeferFor()
}

func moreTypes() {
	internal.Pointers()
	internal.Struct()
	internal.Literals()
	internal.Array()
	internal.Slice()
	internal.MakeSlice()
	internal.SliceOfSlice()
	internal.AppendSlice()
	internal.Range()
	internal.SlicePractice()
	internal.MapParc()
	internal.MutatingMap()
	internal.MapExercise()
	internal.FunValue()
	internal.Closure()
	internal.Fibonacci()
}

func method() {
	internal.Method()
	internal.MethodAreFunc()
	internal.NonStr()
	internal.PointRec()
	internal.PointerAndFunc()
	internal.MedPointerIndirect()
	internal.Choosing()
}

func interfacePrac() {
	internal.MyInterface()
	internal.Implement()
	internal.InterfaceValue()
	internal.InterfaceNilValue()
	//internal.InterfaceNil()
	internal.InterfaceEmpty()
	internal.Assert()
	internal.TypeSwitch()
	internal.Stringer()
	internal.StringerPrac()
	internal.ErrorFunc()
	internal.ErrorPrac()
	internal.ReadPrac()
	internal.ReadPrac2()
	internal.Rot13Reader()
	internal.ImageT()
	//internal.ImagePrac()
}

func generics() {
	internal.Type()
	internal.ListTest()
}

func concurrent() {
	internal.ConcurrentTest()
	internal.Channels()
	internal.BuffChannel()
	internal.ChannelRC()
	internal.ChannelSelect()
	internal.DefaultSelect()
	internal.BinaryTree()
	internal.Mutex()
	internal.WebCrawler()
}
