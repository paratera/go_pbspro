package qstat

import (
	"fmt"
	"testing"
)

func TestServerStat(t *testing.T) {

	qstat, err := NewQstat("172.18.7.10")
	if err != nil {
		t.Error(err)
	}

	qstat.SetAttribs(nil)
	qstat.SetExtend("")

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
	qstat, err := NewQstat("172.18.7.10")
	if err != nil {
		t.Error(err)
	}

	qstat.SetId("pc01")
	qstat.SetAttribs(nil)
	qstat.SetExtend("")

	err = qstat.ConnectPBS()
	if err != nil {
		fmt.Println("ConnectPBS Error")
		t.Error(err)
	}

	bs, err := qstat.Pbs_statnode()
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

func TestQueueStat(t *testing.T) {
	qstat, err := NewQstat("172.18.7.10")
	if err != nil {
		t.Error(err)
	}

	qstat.SetId("workq")
	qstat.SetAttribs(nil)
	qstat.SetExtend("")

	err = qstat.ConnectPBS()
	if err != nil {
		fmt.Println("ConnectPBS Error")
		t.Error(err)
	}

	bs, err := qstat.Pbs_statque()
	if err != nil {
		fmt.Println(err.Error())
	}
	//Print Queue State Informations.
	fmt.Println(bs)

	err = qstat.DisconnectPBS()
	if err != nil {
		fmt.Println("DisconnectPBS Error")
		t.Error(err)
	}

}

func TestJobStat(t *testing.T) {
	qstat, err := NewQstat("172.18.7.10")
	if err != nil {
		t.Error(err)
	}

	qstat.SetId("1045")
	qstat.SetAttribs(nil)
	qstat.SetExtend("")

	err = qstat.ConnectPBS()
	if err != nil {
		fmt.Println("ConnectPBS Error")
		t.Error(err)
	}

	bs, err := qstat.Pbs_statjob()
	if err != nil {
		fmt.Println(err.Error())
	}

	//Print Job State Informations.
	fmt.Println(bs)

}
