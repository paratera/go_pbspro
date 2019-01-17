package qstat

/*
#cgo CFLAGS: -g
#cgo LDFLAGS: -L/opt/pbspro/lib -lpbs
#include <stdlib.h>
#include "/opt/pbspro/include/pbs_error.h"
#include "/opt/pbspro/include/pbs_ifl.h"

// I gave up getting the CGO functions for these right, casting was killing me
static char** mkStringArray (unsigned int len) {
  return (char **) malloc(sizeof(char *) * len);
}

static void freeCstringsN (char **array, unsigned int len) {
    unsigned int i = 0;
    for (i = 0; i < len; i++) {
        free(array[i]);
    }
    free(array);
}

static void addStringToArray (char **array, char *str, unsigned int offset) {
  array[offset] = str;
}
*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/taylor840326/go_pbspro/utils"
)

/*
func Pbs_rescquery(handle int, resources []string) (int, int, int, int, error) {
	var avail, alloc, reserv, down C.int

	rl := cstringArray(resources)
	defer C.freeCstringsN(rl, C.uint(len(resources)))

	ret := C.pbs_rescquery(C.int(handle), rl, C.int(len(resources)), &avail, &alloc, &reserv, &down)
	if ret != 0 {
		return 0, 0, 0, 0, errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}

	return int(avail), int(alloc), int(reserv), int(down), nil
}
*/

/*
func Avail(handle int, resc string) string {
	r := C.CString(resc)
	defer C.free(unsafe.Pointer(r))

	c := C.avail(C.int(handle), r)
	//defer C.free(unsafe.Pointer(c))

	return C.GoString(c)
}
*/

func Pbs_statjob(handle int, id string, attribs []utils.Attrib, extend string) ([]utils.BatchStatus, error) {
	i := C.CString(id)
	defer C.free(unsafe.Pointer(i))

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	a := utils.Pbs_attrib2attribl(attribs)
	defer utils.Pbs_freeattribl(a)

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

	a := utils.Pbs_attrib2attribl(attribs)
	defer utils.Pbs_freeattribl(a)

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

	a := utils.Pbs_attrib2attribl(attribs)
	defer utils.Pbs_freeattribl(a)

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

func Pbs_statserver(handle int, attribs []utils.Attrib, extend string) ([]utils.BatchStatus, error) {
	a := utils.Pbs_attrib2attribl(attribs)
	defer utils.Pbs_freeattribl(a)

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	batch_status := C.pbs_statserver(C.int(handle), a, e)

	/*
		if batch_status == nil {
			return nil, errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
		}
	*/
	defer C.pbs_statfree(batch_status)

	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

func Pbs_selstat(handle int, attribs []utils.Attrib, extend string) ([]utils.BatchStatus, error) {
	a := utils.Pbs_attrib2attribl(attribs)
	defer utils.Pbs_freeattribl(a)

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

/*
func Pbs_fbserver() string {
	// char* from pbs_fbserver is statically allocated, so can't be freed
	return C.GoString(C.pbs_fbserver())
}
*/

/*
func Pbs_get_server_list() string {
	// char* from pbs_get_server_list is statically allocated, so can't be freed
	return C.GoString(C.pbs_get_server_list())
}
*/

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

/*
func Totpool(handle int, update int) (int, error) {
	ret := int(C.totpool(C.int(handle), C.int(update)))
	if ret < 0 {
		return ret, errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	return ret, nil
}
*/
