package klogger

import (
	"errors"
	"fmt"

	"kutil"
)

type IKLogOpt interface {
	VerifyAndSetDefault()
	Verify()
	SetDefault()
}

type KLogFileOpt struct {
	ShiftType			KLogFileShiftType
	RootDirectoryName	string
	Prefix				string
}

func (m *KLogFileOpt) SetDefault() {
	m.ShiftType			= KLogFileShiftType_Day
	m.RootDirectoryName	= "log"
	m.Prefix			= "default"
}

func (m *KLogFileOpt) VerifyAndSetDefault() {
	if err := m.Verify() ; nil != err {
		m.SetDefault()
		println(fmt.Sprintf("!!!---> KLogFileOpt Verify failed and set default : %s", err.Error()))
	}

}

func (m *KLogFileOpt) Verify() (err error) {

	if 0 > m.ShiftType || KLogFileShiftType_Max <= m.ShiftType {
		err = errors.New("KLogFileOpt.Verify() Unknown LogFileSwapType")
		return
	}

	if 0 < len(m.RootDirectoryName) && false == kutil.CheckStringAlphabetOnly(m.RootDirectoryName) {
		err = errors.New("KLogFileOpt.Verify() Set the directory name alphabet only")
		return
	}

	if 0 >= len(m.Prefix) || false == kutil.CheckStringAlphabetOnly(m.Prefix) {
		err = errors.New("KLogFileOpt.Verify() Set the prefix name alphabet only")
		return
	}

	return
}



type KDefaultLoggerOpt struct {
	LogTypeDepth 		KLogType
	LoggerName 			string
	RootDirectoryName	string
	UseQueue			bool
}

func (m *KDefaultLoggerOpt) SetDefault() {
	m.LogTypeDepth			= KLogType_Fatal
	m.LoggerName			= "default"
	m.RootDirectoryName		= "log"
	m.UseQueue				= false
}

func (m *KDefaultLoggerOpt) VerifyAndSetDefault() {
	if err := m.Verify() ; nil != err {
		m.SetDefault()
		println(fmt.Sprintf("!!!---> KDefaultLoggerOpt Verify failed and set default : %s", err.Error()))
	}
}

func (m *KDefaultLoggerOpt) Verify() ( err error ) {

	if 0 < len(m.RootDirectoryName) && false == kutil.CheckStringAlphabetOnly(m.RootDirectoryName) {
		err = errors.New("KDefaultLoggerOpt.Verify() Set the directory name alphabet only")
		return
	}

	if 0 >= len(m.LoggerName) || false == kutil.CheckStringAlphabetOnly(m.LoggerName){
		err = errors.New("KDefaultLoggerOpt.Verify() Set the logger name alphabet only")
		return
	}

	if 0 > m.LogTypeDepth || KLogType_Max <= m.LogTypeDepth {
		err = errors.New("KDefaultLoggerOpt.Verify() Undefined KLogType for LogTypeDepth")
		return
	}

	return
}

