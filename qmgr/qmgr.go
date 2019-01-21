package qmgr

/*
#cgo CFLAGS: -g
#cgo LDFLAGS: -L/opt/pbspro/lib -lpbs
#include <stdlib.h>
#include "/opt/pbspro/include/pbs_error.h"
#include "/opt/pbspro/include/pbs_ifl.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

func Pbs_alterjob(handle int, id string, attribs []Attrib, extend string) error {
	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	s := C.CString(id)
	defer C.free(unsafe.Pointer(s))

	a := attrib2attribl(attribs)
	defer freeattribl(a)

	ret := C.pbs_alterjob(C.int(handle), s, a, e)
	if ret != 0 {
		return errors.New(Pbs_strerror(int(C.pbs_errno)))
	}

	return nil
}

/*
func Pbs_checkpointjob(handle int, id string, extend string) error {
	s := C.CString(id)
	defer C.free(unsafe.Pointer(s))

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	ret := C.pbs_checkpointjob(C.int(handle), s, e)
	if ret != 0 {
		return errors.New(Pbs_strerror(int(C.pbs_errno)))
	}

	return nil
}
*/

// Pbs_deljob deletes a job on the server
func Pbs_deljob(handle int, id string, extend string) error {
	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	s := C.CString(id)
	defer C.free(unsafe.Pointer(s))

	ret := C.pbs_deljob(C.int(handle), s, e)
	if ret != 0 {
		return errors.New(Pbs_strerror(int(C.pbs_errno)))
	}

	return nil
}

/*
func Pbs_gpumode(handle int, mom_node string, gpu_id string, gpu_mode int) error {
	m := C.CString(mom_node)
	defer C.free(unsafe.Pointer(m))

	g := C.CString(gpu_id)
	defer C.free(unsafe.Pointer(g))

	ret := C.pbs_gpumode(C.int(handle), m, g, C.int(gpu_mode))
	if ret != 0 {
		return getLastError()
	}
	return nil
}
*/

/*
// pbs_gpureset not declared in pbs_ifl.h for 3.0.0
func Pbs_gpureset (handle int, mom_node string, gpu_id int, ecc_perm int, ecc_vol int) error {
	m := C.CString(mom_node)
	defer C.free(unsafe.Pointer(m))

    ret := C.pbs_gpureset(C.int(handle), m, C.int(gpu_id), C.int(ecc_perm), C.int(ecc_vol))
    if ret != 0 {
		return getLastError()
    }
    return nil
}
*/

func Pbs_holdjob(handle int, id string, holdType Hold, extend string) error {
	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	s := C.CString(id)
	defer C.free(unsafe.Pointer(s))

	ht := C.CString(string(holdType))
	defer C.free(unsafe.Pointer(ht))

	ret := C.pbs_holdjob(C.int(handle), s, ht, e)
	if ret != 0 {
		return errors.New(Pbs_strerror(int(C.pbs_errno)))
	}

	return nil
}

func Pbs_locjob(handle int, id string) (string, error) {
	s := C.CString(id)
	defer C.free(unsafe.Pointer(s))

	ret := C.pbs_locjob(C.int(handle), s, nil)
	if ret == nil {
		return "", errors.New(Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.free(unsafe.Pointer(ret))

	return C.GoString(ret), nil
}

func Pbs_manager(handle int, command Command, obj_type ObjectType, obj_name string, attrib []Attrib, extend string) error {
	name := C.CString(obj_name)
	defer C.free(unsafe.Pointer(name))

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	a := attrib2attribl(attrib)
	defer freeattribl(a)

	ret := C.pbs_manager(C.int(handle), C.int(command), C.int(obj_type), name, (*C.struct_attropl)(unsafe.Pointer(a)), e)
	if ret != 0 {
		return errors.New(Pbs_strerror(int(C.pbs_errno)))
	}

	return nil
}

func Pbs_movejob(handle int, id string, destination string, extend string) error {
	i := C.CString(id)
	defer C.free(unsafe.Pointer(i))

	d := C.CString(destination)
	defer C.free(unsafe.Pointer(d))

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	ret := C.pbs_movejob(C.int(handle), i, d, e)
	if ret != 0 {
		return errors.New(Pbs_strerror(int(C.pbs_errno)))
	}

	return nil
}

func Pbs_orderjob(handle int, job_id1 string, job_id2, extend string) error {
	j1 := C.CString(job_id1)
	defer C.free(unsafe.Pointer(j1))

	j2 := C.CString(job_id1)
	defer C.free(unsafe.Pointer(j2))

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	ret := C.pbs_orderjob(C.int(handle), j1, j2, e)
	if ret != 0 {
		return errors.New(Pbs_strerror(int(C.pbs_errno)))
	}
	return nil
}

/*
func Usepool(handle int, update int) (int, error) {
	ret := int(C.usepool(C.int(handle), C.int(update)))
	if ret < 0 {
		return ret, errors.New(Pbs_strerror(int(C.pbs_errno)))
	}
	return ret, nil
}
*/

func Pbs_rlsjob(handle int, id string, holdType Hold, extend string) error {
	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	s := C.CString(id)
	defer C.free(unsafe.Pointer(s))

	ht := C.CString(string(holdType))
	defer C.free(unsafe.Pointer(ht))

	ret := C.pbs_rlsjob(C.int(handle), s, ht, e)
	if ret != 0 {
		return errors.New(Pbs_strerror(int(C.pbs_errno)))
	}

	return nil
}

func Pbs_runjob(handle int, id string, location string, extend string) error {
	i := C.CString(id)
	defer C.free(unsafe.Pointer(i))

	l := C.CString(location)
	defer C.free(unsafe.Pointer(l))

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	ret := C.pbs_runjob(C.int(handle), i, l, e)
	if ret != 0 {
		return errors.New(Pbs_strerror(int(C.pbs_errno)))
	}
	return nil
}

// Pbs_selectjob returns a list of jobs
func Pbs_selectjob(handle int, attrib []Attrib, extend string) ([]string, error) {
	// Torque's implementation of pbs_selectjob() is broken and only works
	// accidentally - they allocate a single block of memory (which is
	// oversized) for the jobids and then copy them into it.
	// Because only a single malloc() is used you only need to free() the
	// char** returned by pbs_selectjob().
	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	a := attrib2attribl(attrib)
	defer freeattribl(a)

	p := C.pbs_selectjob(C.int(handle), (*C.struct_attropl)(unsafe.Pointer(a)), e)
	if p == nil {
		return nil, errors.New(Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.free(unsafe.Pointer(p))

	jobArray := cstrings(p)
	return jobArray, nil
}

func Pbs_sigjob(handle int, id string, signal string, extend string) error {
	i := C.CString(id)
	defer C.free(unsafe.Pointer(i))

	s := C.CString(signal)
	defer C.free(unsafe.Pointer(s))

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	ret := C.pbs_sigjob(C.int(handle), i, s, e)
	if ret != 0 {
		return errors.New(Pbs_strerror(int(C.pbs_errno)))
	}
	return nil
}

/*
// pbs_stagein not declared in pbs_ifl.h 3.0.0
func Pbs_stagein(handle int, id string, location string, extend string) error {
	i := C.CString(id)
	defer C.free(unsafe.Pointer(i))

	l := C.CString(location)
	defer C.free(unsafe.Pointer(l))

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

    ret := C.pbs_stagein(C.int(handle), i, l, e)
	if ret != 0 {
		return errors.New(Pbs_strerror(int(C.pbs_errno)))
	}
	return nil
}
*/

func Pbs_submit(handle int, attribs []Attrib, script string, destination string, extend string) (string, error) {
	a := attrib2attribl(attribs)
	defer freeattribl(a)

	s := C.CString(script)
	defer C.free(unsafe.Pointer(s))

	d := C.CString(destination)
	defer C.free(unsafe.Pointer(d))

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	jobid := C.pbs_submit(C.int(handle), (*C.struct_attropl)(unsafe.Pointer(a)), s, d, e)
	if jobid == nil {
		return "", errors.New(Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.free(unsafe.Pointer(jobid))

	return C.GoString(jobid), nil
}

func Pbs_terminate(handle int, manner Manner, extend string) error {
	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	ret := C.pbs_terminate(C.int(handle), C.int(int(manner)), e)
	if ret != 0 {
		return errors.New(Pbs_strerror(int(C.pbs_errno)))
	}
	return nil
}

func Pbs_rerunjob(handle int, id string, extend string) error {
	s := C.CString(id)
	defer C.free(unsafe.Pointer(s))

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	ret := C.pbs_rerunjob(C.int(handle), s, e)

	if ret != 0 {
		return errors.New(Pbs_strerror(int(C.pbs_errno)))
	}
	return nil
}
