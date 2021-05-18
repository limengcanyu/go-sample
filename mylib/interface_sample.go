package mylib

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func InterfaceMain() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}

// 函数实现方法就实现了接口，隐式

// File A simple File interface.
type File interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Close() error
}

type S1 struct {
}

type S2 struct {
}

func (s1 S1) Read(p []byte) (n int, err error) {
	return 1, nil
}
func (s1 S1) Write(p []byte) (n int, err error) {
	return 2, nil
}
func (s1 S1) Close() error {
	return nil
}

func (s2 S2) Read(p []byte) (n int, err error) {
	return 1, nil
}
func (s2 S2) Write(p []byte) (n int, err error) {
	return 2, nil
}
func (s2 S2) Close() error {
	return nil
}

type Locker interface {
	Lock()
	Unlock()
}

func (s1 S1) Lock()   {}
func (s1 S1) Unlock() {}

func (s2 S2) Lock()   {}
func (s2 S2) Unlock() {}

// 接口中嵌入接口

type Reader interface {
	Read(p []byte) (n int, err error)
	Close() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Close() error
}

// ReadWriter ReadWriter's methods are Read, Write, and Close.
type ReadWriter interface {
	Reader // includes methods of Reader in ReadWriter's method set
	Writer // includes methods of Writer in ReadWriter's method set
}
