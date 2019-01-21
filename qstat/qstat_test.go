package qstat

import (
	"fmt"
	"testing"

	"github.com/taylor840326/go_pbspro/utils"
)

func TestServerStat(t *testing.T) {

	qstat, err := NewQstat("172.18.7.10")
	if err != nil {
		t.Error(err)
	}

	qstat.Attribs = nil
	qstat.Extend = ""

	err = qstat.ConnectPBS()
	if err != nil {
		fmt.Println("ConnectPBS Error")
		t.Error(err)
	}

	bs, err := qstat.Pbs_statserver()
	if err != nil {
		fmt.Println(err.Error())
	}
	//Print Server State Informations.
	fmt.Println(bs)

	err = qstat.DisconnectPBS()
	if err != nil {
		fmt.Println("DisconnectPBS Error")
		t.Error(err)
	}
}

func TestNodeStat(t *testing.T) {
	handle, err := utils.Pbs_connect("172.18.7.10")
	if err != nil {
		t.Error(err)
	}

	defer func() {
		err = utils.Pbs_disconnect(handle)
		if err != nil {
			t.Error(err)
		}
	}()

	bs, err := Pbs_statnode(handle, "pc01", nil, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	//Print Server State Informations.
	fmt.Println(bs)

}

func TestQueueStat(t *testing.T) {
	handle, err := utils.Pbs_connect("172.18.7.10")
	if err != nil {
		t.Error(err)
	}

	defer func() {
		err = utils.Pbs_disconnect(handle)
		if err != nil {
			t.Error(err)
		}
	}()

	bs, err := Pbs_statque(handle, "workq", nil, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	//Print Queue State Informations.
	fmt.Println(bs)

}

func TestJobStat(t *testing.T) {
	handle, err := utils.Pbs_connect("172.18.7.10")
	if err != nil {
		t.Error(err)
	}

	defer func() {
		err = utils.Pbs_disconnect(handle)
		if err != nil {
			t.Error(err)
		}
	}()

	bs, err := Pbs_statjob(handle, "1045", nil, "")
	if err != nil {
		fmt.Println(err.Error())
	}

	//Print Job State Informations.
	fmt.Println(bs)

}
