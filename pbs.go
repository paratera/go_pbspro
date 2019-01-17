// Package pbs provides an interface to the C-based TORQUE library.
// The functions present in this package are a thin wrapper around these
// library functions and as such the TORQUE library provides documentation on
// their usage.
//
// The TORQUE library is not thread safe, particulary when it comes to
// reporting errors, and therefore problems *might* arise if you use this
// package with goroutines.
//
// The following functions have not yet been implemented:
/*
   pbs_alterjob      - untested
   pbs_alterjobasync
   pbs_manager       - untested
   pbs_rescreserve
   pbs_asyncrunjob
   pbs_sigjobasync
*/
package pbspro

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
)

func getLastError() error {
	return errors.New(Pbs_strerror(int(C.pbs_errno)))
}

func attrib2attribl(attribs []Attrib) *C.struct_attrl {
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

func freeattribl(attrl *C.struct_attrl) {
	for p := attrl; p != nil; p = p.next {
		C.free(unsafe.Pointer(p.name))
		C.free(unsafe.Pointer(p.value))
		C.free(unsafe.Pointer(p.resource))
	}
}

func get_pbs_batch_status(batch_status *_Ctype_struct_batch_status) (batch []BatchStatus) {
	for batch_status != nil {
		temp := []Attrib{}
		for attr := batch_status.attribs; attr != nil; attr = attr.next {
			temp = append(temp, Attrib{
				Name:     C.GoString(attr.name),
				Resource: C.GoString(attr.resource),
				Value:    C.GoString(attr.value),
			})
		}

		batch = append(batch, BatchStatus{
			Name:       C.GoString(batch_status.name),
			Text:       C.GoString(batch_status.text),
			Attributes: temp,
		})

		batch_status = batch_status.next
	}
	return batch
}

func sptr(p uintptr) *C.char {
	return *(**C.char)(unsafe.Pointer(p))
}

func cstrings(x **C.char) []string {
	var s []string
	for p := uintptr(unsafe.Pointer(x)); sptr(p) != nil; p += unsafe.Sizeof(uintptr(0)) {
		s = append(s, C.GoString(sptr(p)))
	}
	return s
}

func freeCstrings(x **C.char) {
	for p := uintptr(unsafe.Pointer(x)); sptr(p) != nil; p += unsafe.Sizeof(uintptr(0)) {
		C.free(unsafe.Pointer(sptr(p)))
	}
}

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

// Pbs_connect makes a connection to server, or if server is an empty string, the default server. The returned handle is used by subsequent calls to the functions in this package to identify the server.
func Pbs_connect(server string) (int, error) {
	str := C.CString(server)
	defer C.free(unsafe.Pointer(str))

	handle := C.pbs_connect(str)
	if handle < 0 {
		return 0, errors.New(Pbs_strerror(int(C.pbs_errno)))
	}

	return int(handle), nil
}

// Pbs_default reports the default torque server
func Pbs_default() string {
	// char* from pbs_default is statically allocated, so can't be freed
	return C.GoString(C.pbs_default())
}

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

func Pbs_disconnect(handle int) error {
	ret := C.pbs_disconnect(C.int(handle))
	if ret != 0 {
		return errors.New(Pbs_strerror(int(C.pbs_errno)))
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

func Pbs_geterrmsg(handle int) string {
	s := C.pbs_geterrmsg(C.int(handle))
	// char* from pbs_geterrmsg is statically allocated, so can't be freed
	return C.GoString(s)
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

func Pbs_selstat(handle int, attribs []Attrib, extend string) ([]BatchStatus, error) {
	a := attrib2attribl(attribs)
	defer freeattribl(a)

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	batch_status := C.pbs_selstat(C.int(handle), (*C.struct_attropl)(unsafe.Pointer(a)), a, e)

	// FIXME: nil also indicates no jobs matched selection criteria...
	if batch_status == nil {
		return nil, errors.New(Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)
	batch := get_pbs_batch_status(batch_status)

	return batch, nil
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

func Pbs_msgjob(handle int, id string, file MessageStream, message string, extend string) error {
	s := C.CString(id)
	defer C.free(unsafe.Pointer(s))

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	m := C.CString(message)
	defer C.free(unsafe.Pointer(m))

	ret := C.pbs_msgjob(C.int(handle), s, C.int(file), m, e)
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

func cstringArray(strings []string) **C.char {
	c := C.mkStringArray(C.uint(len(strings)))
	for i, str := range strings {
		C.addStringToArray(c, C.CString(str), C.uint(i))
	}
	return c
}

/*
func Pbs_rescquery(handle int, resources []string) (int, int, int, int, error) {
	var avail, alloc, reserv, down C.int

	rl := cstringArray(resources)
	defer C.freeCstringsN(rl, C.uint(len(resources)))

	ret := C.pbs_rescquery(C.int(handle), rl, C.int(len(resources)), &avail, &alloc, &reserv, &down)
	if ret != 0 {
		return 0, 0, 0, 0, errors.New(Pbs_strerror(int(C.pbs_errno)))
	}

	return int(avail), int(alloc), int(reserv), int(down), nil
}
*/

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

/*
func Avail(handle int, resc string) string {
	r := C.CString(resc)
	defer C.free(unsafe.Pointer(r))

	c := C.avail(C.int(handle), r)
	//defer C.free(unsafe.Pointer(c))

	return C.GoString(c)
}
*/

/*
func Totpool(handle int, update int) (int, error) {
	ret := int(C.totpool(C.int(handle), C.int(update)))
	if ret < 0 {
		return ret, errors.New(Pbs_strerror(int(C.pbs_errno)))
	}
	return ret, nil
}
*/

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

func Pbs_statjob(handle int, id string, attribs []Attrib, extend string) ([]BatchStatus, error) {
	i := C.CString(id)
	defer C.free(unsafe.Pointer(i))

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	a := attrib2attribl(attribs)
	defer freeattribl(a)

	batch_status := C.pbs_statjob(C.int(handle), i, a, e)

	if batch_status == nil {
		return nil, errors.New(Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)

	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

func Pbs_statnode(handle int, id string, attribs []Attrib, extend string) ([]BatchStatus, error) {
	i := C.CString(id)
	defer C.free(unsafe.Pointer(i))

	a := attrib2attribl(attribs)
	defer freeattribl(a)

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	batch_status := C.pbs_statnode(C.int(handle), i, a, e)

	if batch_status == nil {
		return nil, errors.New(Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)

	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

func Pbs_statque(handle int, id string, attribs []Attrib, extend string) ([]BatchStatus, error) {
	i := C.CString(id)
	defer C.free(unsafe.Pointer(i))

	a := attrib2attribl(attribs)
	defer freeattribl(a)

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	batch_status := C.pbs_statque(C.int(handle), i, a, e)

	if batch_status == nil {
		return nil, errors.New(Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)

	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

func Pbs_statserver(handle int, attribs []Attrib, extend string) ([]BatchStatus, error) {
	a := attrib2attribl(attribs)
	defer freeattribl(a)

	e := C.CString(extend)
	defer C.free(unsafe.Pointer(e))

	batch_status := C.pbs_statserver(C.int(handle), a, e)

	/*
		if batch_status == nil {
			return nil, errors.New(Pbs_strerror(int(C.pbs_errno)))
		}
	*/
	defer C.pbs_statfree(batch_status)

	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

func Pbs_strerror(errno int) string {
	// char* from pbs_strerror is statically allocated, so can't be freed
	return C.GoString(C.pbse_to_txt(C.int(errno)))
}

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
