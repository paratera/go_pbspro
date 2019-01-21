package qstat

/*
#cgo CFLAGS: -g
#cgo LDFLAGS: -L/opt/pbspro/lib -lpbs
#include <stdlib.h>
#include "/opt/pbspro/include/pbs_error.h"
#include "/opt/pbspro/include/pbs_ifl.h"
*/
import "C"
import (
	"unsafe"

	"github.com/juju/errors"

	"github.com/taylor840326/go_pbspro/utils"
)

type (
	//定义PBS结构体
	Qstat struct {
		Server        string         `json:"server"`
		Handle        int            `json:"handle"`
		DefaultServer string         `json:"default_server"`
		IsClosed      bool           `json:"is_closed"`
		Attribs       []utils.Attrib `json:"attribs"`
		Extend        string         `json:"extend"`
	}
)

//新建一个Qstat实例
func NewQstat(server string) (qs *Qstat, err error) {
	qstat := new(Qstat)

	qstat.Server = server
	qstat.Handle = 0
	qstat.DefaultServer = nil
	qstat.IsClosed = 0
	qstat.Attribs = nil
	qstat.Extend = ""

	return qstat, nil
}

//设定服务名称
func (qs *Qstat) SetServerName(server string) {
	qs.Server = server
}

//设定handle号，>= 0
func (qs *Qstat) SetHandle(handle int) {
	qs.Handle = handle
}

//设定属性列表
func (qs *Qstat) Attribs(attribs []utils.Attrib) {
	qs.Attribs = attribs
}

//设定扩展信息列表.
func (qs *Qstat) SetExtend(extend string) {
	qs.Extend = extend
}

//创建一个新的连接
func (qs *Qstat) ConnectPBS() error {
	var err error
	qs.Handle, err = utils.Pbs_connect(qs.Server)
	if err != nil {
		return errors.NewBadRequest(err, "Cann't connect PBSpro Server")
	}

	return nil
}

//断开连接
func (qs *Qstat) DisconnectPBS() error {
	err := utils.Pbs_disconnect(qs.Handle)
	if err != nil {
		return errors.NewBadRequest(err, "Can't disconnect PBSpro Server")
	}
	return nil
}

func Pbs_attrib2attribl(attribs []utils.Attrib) *C.struct_attrl {
	// Empty array returns null pointer
	if len(attribs) == 0 {
		return nil
	}

	first := &C.struct_attrl{
		value:    C.CString(attribs[0].Value),
		resource: C.CString(attribs[0].Resource),
		name:     C.CString(attribs[0].Name),
		op:       uint32(attribs[0].Op),
	}
	tail := first

	for _, attr := range attribs[1:len(attribs)] {
		tail.next = &C.struct_attrl{
			value:    C.CString(attr.Value),
			resource: C.CString(attr.Resource),
			name:     C.CString(attr.Name),
			op:       uint32(attribs[0].Op),
		}
	}

	return first
}

func Pbs_freeattribl(attrl *C.struct_attrl) {
	for p := attrl; p != nil; p = p.next {
		C.free(unsafe.Pointer(p.name))
		C.free(unsafe.Pointer(p.value))
		C.free(unsafe.Pointer(p.resource))
	}
}

func Pbs_statjob(handle int, id string, attribs []utils.Attrib, extend string) ([]utils.BatchStatus, error) {
	i := C.CString(id)
	defer C.free(unsafe.Pointer(i))

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	a := Pbs_attrib2attribl(attribs)
	defer Pbs_freeattribl(a)

	batch_status := C.pbs_statjob(C.int(handle), i, a, e)

	if batch_status == nil {
		return nil, errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)

	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

func Pbs_statnode(handle int, id string, attribs []utils.Attrib, extend string) ([]utils.BatchStatus, error) {
	i := C.CString(id)
	defer C.free(unsafe.Pointer(i))

	a := Pbs_attrib2attribl(attribs)
	defer Pbs_freeattribl(a)

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	batch_status := C.pbs_statnode(C.int(handle), i, a, e)

	if batch_status == nil {
		return nil, errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)

	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

func Pbs_statque(handle int, id string, attribs []utils.Attrib, extend string) ([]utils.BatchStatus, error) {
	i := C.CString(id)
	defer C.free(unsafe.Pointer(i))

	a := Pbs_attrib2attribl(attribs)
	defer Pbs_freeattribl(a)

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	batch_status := C.pbs_statque(C.int(handle), i, a, e)

	if batch_status == nil {
		return nil, errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)

	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

func (qstat *Qstat) Pbs_statserver() ([]utils.BatchStatus, error) {
	a := Pbs_attrib2attribl(qstat.Attribs)
	defer Pbs_freeattribl(a)

	e := C.CString(qstat.Extend)
	defer C.free(unsafe.Pointer(e))

	batch_status := C.pbs_statserver(C.int(qstat.Handle), a, e)

	if batch_status == nil {
		return nil, errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)

	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

func Pbs_selstat(handle int, attribs []utils.Attrib, extend string) ([]utils.BatchStatus, error) {
	a := Pbs_attrib2attribl(attribs)
	defer Pbs_freeattribl(a)

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	batch_status := C.pbs_selstat(C.int(handle), (*C.struct_attropl)(unsafe.Pointer(a)), a, e)

	// FIXME: nil also indicates no jobs matched selection criteria...
	if batch_status == nil {
		return nil, errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)
	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

func Pbs_msgjob(handle int, id string, file utils.MessageStream, message string, extend string) error {
	s := C.CString(id)
	defer C.free(unsafe.Pointer(s))

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	m := C.CString(message)
	defer C.free(unsafe.Pointer(m))

	ret := C.pbs_msgjob(C.int(handle), s, C.int(file), m, e)
	if ret != 0 {
		return errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	return nil
}

func get_pbs_batch_status(batch_status *_Ctype_struct_batch_status) (batch []utils.BatchStatus) {

	for batch_status != nil {
		temp := []utils.Attrib{}
		for attr := batch_status.attribs; attr != nil; attr = attr.next {
			temp = append(temp, utils.Attrib{
				Name:     C.GoString(attr.name),
				Resource: C.GoString(attr.resource),
				Value:    C.GoString(attr.value),
			})
		}

		batch = append(batch, utils.BatchStatus{
			Name:       C.GoString(batch_status.name),
			Text:       C.GoString(batch_status.text),
			Attributes: temp,
		})

		batch_status = batch_status.next
	}
	return batch
}
