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

	err = qstat.PbsNodeState()
	if err != nil {
		fmt.Println(err.Error())
	}
	//Print Server State Informations.
	fmt.Println(qstat.NodeState)

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

	err = qstat.PbsQueueState()
	if err != nil {
		fmt.Println(err.Error())
	}
	//Print Queue State Informations.
	fmt.Println(qstat.QueueState)

	err = qstat.DisconnectPBS()
	if err != nil {
		fmt.Println("DisconnectPBS Error")
		t.Error(err)
	}

}

func TestOneJobStat(t *testing.T) {
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
		_, err := qstat.Pbs_statjob()
		if err != nil {
			fmt.Println(err.Error())
		}
		//Print Job State Informations.
		//fmt.Println(bs)
	}

}

func TestAllJobState(t *testing.T) {
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

	err = qstat.PbsJobsState()
	if err != nil {
		fmt.Println(err.Error())
	}
	//Print Job State Informations.
	fmt.Println(qstat.JobsState)

}
