package _interface

//战法锁定
type TacticsLocking interface {
	//锁定主将
	IsLockingMaster() bool
	//锁定副将
	IsLockingVice() bool
	//锁定武将
	LockingGeneral() int64
}
