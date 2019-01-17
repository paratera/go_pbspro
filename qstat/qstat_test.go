package main

import (
	"fmt"
	"log"

	"github.com/taylor840326/go_pbspro/qstat"
	"github.com/taylor840326/go_pbspro/utils"
)

func main() {
	handle, err := utils.Pbs_connect("172.18.7.10")
	if err != nil {
		log.Fatal("Couldn't connect to server: %s", err)
	}

	defer func() {
		err = utils.Pbs_disconnect(handle)
		if err != nil {
			log.Fatal("Disconnect failed: %s\n", err)
		}
	}()

	/*
	   _, err = pbs.Pbs_submit(handle, nil, "test.sh", "","")
	*/
	bs, err := qstat.Pbs_statjob(handle, "1045", nil, "")
	if err != nil {
		log.Fatal("Job submission failed: %s\n", err)
	}

	fmt.Println(bs)
	//fmt.Fprintf(os.Stdout,"bs.Name=%s,bs.Text=%s",bs.Name,bs.Text)

}
