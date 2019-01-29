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

	err = qstat.PbsServerState()
	if err != nil {
		fmt.Println(err.Error())
	}
	//Print Server State Informations.
	fmt.Println(qstat.ServerState)

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

	qstat.SetID("pc01")
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

	qstat.SetID("workq")
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

	qstat.SetAttribs(nil)
	qstat.SetExtend("x")

	err = qstat.ConnectPBS()
	if err != nil {
		fmt.Println("ConnectPBS Error")
		t.Error(err)
	}

	for i := 0; i < 100; i++ {
		qstat.SetID(fmt.Sprintf("%d", i))
		bs, err := qstat.Pbs_statjob()
		if err != nil {
			fmt.Println(err.Error())
		}
		//Print Job State Informations.
		fmt.Println(bs)
	}

}
