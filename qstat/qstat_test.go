package qstat

import (
	"fmt"
	"testing"

	"github.com/taylor840326/go_pbspro/utils"
)

func TestServerLists(t *testing.T) {
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

	sl := Pbs_get_server_list()
	//Print Server State Informations.
	fmt.Println(sl)

}

func TestServerStat(t *testing.T) {
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

	bs, err := Pbs_statserver(handle, nil, "")
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
