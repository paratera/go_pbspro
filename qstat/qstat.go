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
	"fmt"
	"strconv"
	"strings"
	"unsafe"

	"github.com/juju/errors"

	"github.com/taylor840326/go_pbspro/utils"
)

type (
	// qstat gather server state information.
	QstatServerInfo struct {
		ServerName              string `json:"server_name" db:"server_name"`
		ServerState             string `json:"server_state" db:"server_state"`
		ServerHost              string `json:"server_host" db:"server_host"`
		ServerScheduling        string `json:"server_scheduling" db:"server_scheduling"`
		TotalJobs               int64  `json:"total_jobs" db:"total_jobs"`
		StateCountTransit       int64  `json:"state_count_transit" db:"state_count_transit"`
		StateCountQueued        int64  `json:"state_count_queued" db:"state_count_queued"`
		StateCountHeld          int64  `json:"state_count_held" db:"state_count_held"`
		StateCountWaiting       int64  `json:"state_count_waiting" db:"state_count_waiting"`
		StateCountRunning       int64  `json:"state_count_running" db:"state_count_running"`
		StateCountExiting       int64  `json:"state_count_exiting" db:"state_count_exiting"`
		StateCountBegun         int64  `json:"state_count_begun" db:"state_count_begun"`
		DefaultQueue            string `json:"default_queue" db:"default_queue"`
		LogEvents               int64  `json:"log_events" db:"log_events"`
		MailFrom                string `json:"mail_from" db:"mail_from"`
		QueryOtherJobs          string `json:"query_other_jobs" db:"query_other_jobs"`
		ResourcesDefaultNcpus   int64  `json:"resources_default_ncpus" db:"resources_default_ncpus"`
		DefaultChunkNcpus       int64  `json:"default_chunk_ncpus" db:"default_chunk_ncpus"`
		ResourcesAssignedNcpus  int64  `json:"resources_assigned_ncpus" db:"resources_assigned_ncpus"`
		ResourcesAssignedNodect int64  `json:"resources_assigned_nodect" db:"resources_assigned_nodect"`
		SchedulerIteration      int64  `json:"scheduler_iteration" db:" scheduler_iteration"`
		Flicenses               int64  `json:"flicenses" db:"flicenses"`
		ResvEnable              string `json:"resv_enable" db:"resv_enable"`
		NodeFailRequeue         int64  `json:"node_fail_requeue" db:"node_fail_requeue"`
		MaxArraySize            int64  `json:"max_array_size" db:"max_array_size"`
		PBSLicenseMin           int64  `json:"pbs_license_min" db:"pbs_license_min"`
		PBSLicenseMax           int64  `json:"pbs_license_max" db:"pbs_license_max"`
		PBSLicenseLingerTime    int64  `json:"pbs_license_linger_time" db:"pbs_license_linger_time"`
		LicenseCountAvailGlobal int64  `json:"license_count_avail_global" db:"license_count_avail_global"`
		LicenseCountAvailLocal  int64  `json:"license_count_avail_local" db:"license_count_avail_local"`
		LicenseCountUsed        int64  `json:"license_count_used" db:"license_count_used"`
		LicenseCountHighUse     int64  `json:"license_count_high_use" db:"license_count_high_use"`
		PBSVersion              string `json:"pbs_version" db:"pbs_version"`
		EligibleTimeEnable      string `json:"eligible_time_enable" db:"eligible_time_enable"`
		JobHistoryEnable        string `json:"job_history_enable" db:"job_history_enable"`
		JobHistoryDuration      string `json:"job_history_duration" db:"job_history_duration"`
		MaxConcurrentProvision  int64  `json:"max_concurrent_provision" db:"max_concurrent_provision"`
		PowerProvisioning       string `json:"power_provisioning" db:"power_provisioning"`
	}

	// qstat gather queue information.
	QstatQueueInfo struct {
		QueueName               string `json:"queue_name" db:"queue_name"`
		QueueType               string `json:"queue_type" db:"queue_type"`
		TotalJobs               int64  `json:"total_jobs" db:"total_jobs"`
		StateCountTransit       int64  `json:"state_count_transit" db:"state_count_transit"`
		StateCountQueued        int64  `json:"state_count_queued" db:"state_count_queued"`
		StateCountHeld          int64  `json:"state_count_held" db:"state_count_held"`
		StateCountWaiting       int64  `json:"state_count_waiting" db:"state_count_waiting"`
		StateCountRunning       int64  `json:"state_count_running" db:"state_count_running"`
		StateCountExiting       int64  `json:"state_count_exiting" db:"state_count_exiting"`
		StateCountBegun         int64  `json:"state_count_begun" db:"state_count_begun"`
		ResourcesAssignedNcpus  int64  `json:"resources_assigned_ncpus" db:"resources_assigned_ncpus"`
		ResourcesAssignedNodect int64  `json:"resources_assigned_nodect" db:"resources_assigned_nodect"`
		Enable                  string `json:"enable" db:"enable"`
		Started                 string `json:"started" db:"started"`
	}

	//qstat gather node information.
	QstatNodeInfo struct {
		NodeName                           string `json:"node_name" db:"node_name"`
		Mom                                string `json:"mom" db:"mom"`
		Ntype                              string `json:"ntype" db:"ntype"`
		State                              string `json:'state" db:"state"`
		Pcpus                              int64  `json:"pcpus" db:"pcpus"`
		Jobs                               string `json:"jobs" db:"jobs"`
		ResourcesAvailableArch             string `json:"resources_available_arch" db:"resources_available_arch"`
		ResourcesAvailableHost             string `json:"resources_available_host" db:"resources_available_host"`
		ResourcesAvailableMem              string `json:"resources_available_mem" db:"resources_available_mem"`
		ResourcesAvailableNcpus            int64  `json:"resources_available_ncpus" db:"resources_available_ncpus"`
		ResourcesAvailableApplications     string `json:"resources_available_pas_applications_enabled" db:"resources_available_pas_applications_enabled"`
		ResourcesAvailablePlatform         string `json:"resources_available_platform" db:"resources_available_platform"`
		ResourcesAvailableSoftware         string `json:"resources_availabled_software" db:"resources_available_software"`
		ResourcesAvailableVnodes           string `json:"resources_available_vnodes" db:"resources_available_vnodes"`
		ResourcesAssignedAcceleratorMemory string `json:"resources_assigned_accelerator_memory" db:"resources_assigned_accelerator_memory"`
		ResourcesAssignedHbmem             string `json:"resources_assigned_hbmem" db:"resources_assigned_hbmem"`
		ResourcesAssignedMem               string `json:"resources_assigned_mem" db:"resources_assigned_mem"`
		ResourcesAssignedNaccelerators     int64  `json:"resources_assigned_naccelerators" db:"resources_assigned_naccelerators"`
		ResourcesAssignedNcpus             int64  `json:"resources_assigned_ncpus" db:"resources_assigned_ncpus"`
		ResourcesAssignedVmem              string `json:"resources_assigned_vmem" db:"resources_assigned_vmem"`
		ResvEnable                         string `json:"resv_enable" db:"resv_enable"`
		Sharing                            string `json:"sharing" db:"sharing"`
		LastStateChangeTime                int64  `json:"last_state_change_time" db:"last_state_change_time"`
		LastUsedTime                       int64  `json:"last_used_time" db:"last_used_time"`
	}

	//qstat gather jobs information.
	QstatJobsInfo struct {
		JobName                 string  `json:"job_name" db:"job_name"`
		JobOwner                string  `json:"job_owner" db:"job_owner"`
		ResourcesUsedCpuPercent float64 `json:"resources_used_cpupercent" db:"resources_used_cpupercent"`
		ResourcesUsedCput       string  `json:"resources_used_cput" db:"resources_used_cput"`
		ResourcesUsedMem        string  `json:"resources_used_mem" db:"resources_used_mem"`
		ResourcesUsedNcpus      int64   `json:"resources_used_ncpus" db:"resources_used_ncpus"`
		ResourcesUsedVmem       string  `json:"resources_used_vmem" db:"resources_used_vmem"`
		ResourcesUsedWallTime   string  `json:"resources_used_walltime" db:"resources_used_walltime"`
		JobState                string  `json:"job_state" db:"job_state"`
		Queue                   string  `json:"queue" db:"queue"`
		Server                  string  `json:"server" db:"server"`
		CheckPoint              string  `json:"checkpoint" db:"checkpoint"`
		Ctime                   int64   `json:"ctime" db:"ctime"`
		ErrorPath               string  `json:"error_path" db:"error_path"`
		ExecHost                string  `json:"exec_host" db:"exec_host"`
		ExecVnode               string  `json:"exec_vnode" db:"exec_vnode"`
		HoldType                string  `json:"hold_type" db:"hold_type"`
		JoinPath                string  `json:"join_path" db:"join_path"`
		KeepFiles               string  `json:"keep_files" db:"keep_files"`
		MailPoints              string  `json:"mail_points" db:"mail_points"`
		Mtime                   int64   `json:"mtime" db:"mtime"`
		OutputPath              string  `json:"output_path" db:"output_path"`
		Priorty                 int64   `json:"priorty" db:"priorty"`
		Qtime                   int64   `json:"qtime" db:"qtime"`
		Rerunable               string  `json:"rerunable" db:"rerunable"`
		ResourceListNcpus       int64   `json:"resource_list_ncpus" db:"resource_list_ncpus"`
		ResourceListNodect      int64   `json:"resource_list_nodect" db:"resource_list_nodect"`
		ResourceListPlace       string  `json:"resource_list_place" db:"resource_list_place"`
		ResourceListSelect      string  `json:"resource_list_select" db:"resource_list_select"`
		ResourceListSoftware    string  `json:"resource_list_software" db:"resource_list_software"`
		ResourceListWallTime    string  `json:"resource_list_walltime" db:"resource_list_walltime"`
		Stime                   int64   `json:"stime" db:"stime"`
		SessionID               int64   `json:"session_id" db:"session_id"`
		JobDir                  string  `json:"jobdir" db:"jobdir"`
		SubState                int64   `json:"substate" db:"substate"`
		VariableList            string  `json:"variable_list" db:"variable_list"`
		Comment                 string  `json:"comment" db:"comment"`
		Etime                   int64   `json:"etime" db:"etime"`
		RunCount                int64   `json:"run_count" db:"run_count"`
		SubmitArguments         string  `json:"submit_arguments" db:"submit_arguments"`
		Project                 string  `json:"project" db:"project"`
	}

	//定义PBS结构体
	Qstat struct {
		Server  string         `json:"server"`
		Handle  int            `json:"handle"`
		Attribs []utils.Attrib `json:"attribs"`
		Extend  string         `json:"extend"`
		ID      string         `json:"id"`

		// server state information.
		ServerState []QstatServerInfo
		// queue state information.
		QueueState []QstatQueueInfo
		// node state information.
		NodeState []QstatNodeInfo
		// jobs state information.
		JobsState []QstatJobsInfo
	}
)

//新建一个Qstat实例
func NewQstat(server string) (qs *Qstat, err error) {
	qstat := new(Qstat)

	qstat.Server = server
	qstat.Handle = 0
	qstat.Attribs = nil
	qstat.Extend = ""
	qstat.ID = ""

	qstat.ServerState = []QstatServerInfo{}
	qstat.QueueState = []QstatQueueInfo{}
	qstat.NodeState = []QstatNodeInfo{}
	qstat.JobsState = []QstatJobsInfo{}

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
func (qs *Qstat) SetAttribs(attribs []utils.Attrib) {
	qs.Attribs = attribs
}

//设定扩展信息列表.
func (qs *Qstat) SetExtend(extend string) {
	qs.Extend = extend
}

//设定Id值
func (qs *Qstat) SetID(id string) {
	qs.ID = id
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

//查询指定作业的信息
func (qs *Qstat) Pbs_statjob() ([]utils.BatchStatus, error) {
	i := C.CString(qs.ID)
	defer C.free(unsafe.Pointer(i))

	e := C.CString(qs.Extend)
	defer C.free(unsafe.Pointer(e))

	a := Pbs_attrib2attribl(qs.Attribs)
	defer Pbs_freeattribl(a)

	batch_status := C.pbs_statjob(C.int(qs.Handle), i, a, e)

	if batch_status == nil {
		return nil, errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)

	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

//查询指定节点状态
func (qs *Qstat) PbsNodeState() error {
	i := C.CString(qs.ID)
	defer C.free(unsafe.Pointer(i))

	a := Pbs_attrib2attribl(qs.Attribs)
	defer Pbs_freeattribl(a)

	e := C.CString(qs.Extend)
	defer C.free(unsafe.Pointer(e))

	batch_status := C.pbs_statnode(C.int(qs.Handle), i, a, e)

	if batch_status == nil {
		return errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)

	batch := get_pbs_batch_status(batch_status)

	for _, bs := range batch {
		var tmpServerNodeState QstatNodeInfo
		tmpServerNodeState.NodeName = bs.Name
		for _, attr := range bs.Attributes {
			switch attr.Name {
			case "Mom":
				tmpServerNodeState.Mom = attr.Value
			case "ntype":
				tmpServerNodeState.Ntype = attr.Value
			case "state":
				tmpServerNodeState.State = attr.Value
			case "pcpus":
				tmpServerNodeState.Pcpus, _ = strconv.ParseInt(attr.Value, 10, 64)
			case "jobs":
				tmpServerNodeState.Jobs = attr.Value
			case "resources_available":
				if len(attr.Resource) == 0 {
					break
				}
				switch attr.Resource {
				case "arch":
					tmpServerNodeState.ResourcesAvailableArch = attr.Value
				case "host":
					tmpServerNodeState.ResourcesAvailableHost = attr.Value
				case "mem":
					tmpServerNodeState.ResourcesAssignedMem = attr.Value
				case "ncpus":
					tmpServerNodeState.ResourcesAssignedNcpus, _ = strconv.ParseInt(attr.Value, 10, 64)
				case "pas_applications_enabled":
					tmpServerNodeState.ResourcesAvailableApplications = attr.Value
				case "platform":
					tmpServerNodeState.ResourcesAvailablePlatform = attr.Value
				case "software":
					tmpServerNodeState.ResourcesAvailableSoftware = attr.Value
				case "vnode":
					tmpServerNodeState.ResourcesAvailableVnodes = attr.Value
				default:
					fmt.Println("other node resources avaiable", attr.Name)
				}
			case "resources_assigned":
				if len(attr.Resource) == 0 {
					break
				}
				switch attr.Resource {
				case "accelerator_memory":
					tmpServerNodeState.ResourcesAssignedAcceleratorMemory = attr.Value
				case "hbmem":
					tmpServerNodeState.ResourcesAssignedHbmem = attr.Value
				case "mem":
					tmpServerNodeState.ResourcesAssignedMem = attr.Value
				case "naccelerators":
					tmpServerNodeState.ResourcesAssignedNaccelerators, _ = strconv.ParseInt(attr.Value, 10, 64)
				case "ncpus":
					tmpServerNodeState.ResourcesAssignedNcpus, _ = strconv.ParseInt(attr.Value, 10, 64)
				case "vmem":
					tmpServerNodeState.ResourcesAssignedVmem = attr.Value
				default:
					fmt.Println("other node resources assigned", attr.Name)
				}
			case "resv_enable":
				tmpServerNodeState.ResvEnable = attr.Value
			case "sharing":
				tmpServerNodeState.Sharing = attr.Value
			case "last_state_change_time":
				tmpServerNodeState.LastStateChangeTime, _ = strconv.ParseInt(attr.Value, 10, 64)
			case "last_used_time":
				tmpServerNodeState.LastUsedTime, _ = strconv.ParseInt(attr.Value, 10, 64)
			default:
				fmt.Println("other node state", attr.Name)
			}
		}
		qs.NodeState = append(qs.NodeState, tmpServerNodeState)
	}

	return nil
}

//查询指定队列信息
func (qs *Qstat) PbsQueueState() error {
	i := C.CString(qs.ID)
	defer C.free(unsafe.Pointer(i))

	a := Pbs_attrib2attribl(qs.Attribs)
	defer Pbs_freeattribl(a)

	e := C.CString(qs.Extend)
	defer C.free(unsafe.Pointer(e))

	batch_status := C.pbs_statque(C.int(qs.Handle), i, a, e)

	if batch_status == nil {
		return errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)

	batch := get_pbs_batch_status(batch_status)

	for _, bs := range batch {
		var tmpServerQueueState QstatQueueInfo
		tmpServerQueueState.QueueName = bs.Name
		for _, attr := range bs.Attributes {
			switch attr.Name {
			case "queue_type":
				tmpServerQueueState.QueueType = attr.Value
			case "total_jobs":
				tmpServerQueueState.TotalJobs, _ = strconv.ParseInt(attr.Value, 10, 64)
			case "state_count":
				attrArray := strings.Split(attr.Value, " ")
				for _, sc_valu := range attrArray {
					if len(sc_valu) == 0 {
						break
					}
					scname := strings.Split(sc_valu, ":")[0]
					scval := strings.Split(sc_valu, ":")[1]
					switch scname {
					case "Transit":
						tmpServerQueueState.StateCountTransit, _ = strconv.ParseInt(scval, 10, 64)
					case "Queued":
						tmpServerQueueState.StateCountQueued, _ = strconv.ParseInt(scval, 10, 64)
					case "Held":
						tmpServerQueueState.StateCountHeld, _ = strconv.ParseInt(scval, 10, 64)
					case "Waiting":
						tmpServerQueueState.StateCountWaiting, _ = strconv.ParseInt(scval, 10, 64)
					case "Running":
						tmpServerQueueState.StateCountRunning, _ = strconv.ParseInt(scval, 10, 64)
					case "Exiting":
						tmpServerQueueState.StateCountExiting, _ = strconv.ParseInt(scval, 10, 64)
					case "Begun":
						tmpServerQueueState.StateCountBegun, _ = strconv.ParseInt(scval, 10, 64)
					default:
						fmt.Println("other server_state_count")
					}
				}
			case "resources_assigned":
				if attr.Resource == "ncpus" {
					tmpServerQueueState.ResourcesAssignedNcpus, _ = strconv.ParseInt(attr.Value, 10, 64)
				}
				if attr.Resource == "nodect" {
					tmpServerQueueState.ResourcesAssignedNodect, _ = strconv.ParseInt(attr.Value, 10, 64)
				}
			case "enabled":
				tmpServerQueueState.Enable = attr.Value
			case "started":
				tmpServerQueueState.Started = attr.Value
			default:
				fmt.Println("other queue state", attr.Name)
			}
		}
		qs.QueueState = append(qs.QueueState, tmpServerQueueState)
	}

	return nil
}

//查询服务信息
func (qs *Qstat) PbsServerState() error {
	a := Pbs_attrib2attribl(qs.Attribs)
	defer Pbs_freeattribl(a)

	e := C.CString(qs.Extend)
	defer C.free(unsafe.Pointer(e))

	batchStatus := C.pbs_statserver(C.int(qs.Handle), a, e)

	if batchStatus == nil {
		return errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batchStatus)

	batch := get_pbs_batch_status(batchStatus)

	for _, value := range batch {
		var tmp_server_state_info QstatServerInfo
		tmp_server_state_info.ServerName = value.Name
		for _, attr := range value.Attributes {
			switch attr.Name {
			case "server_state":
				tmp_server_state_info.ServerState = attr.Value
			case "server_host":
				tmp_server_state_info.ServerHost = attr.Value
			case "scheduling":
				tmp_server_state_info.ServerScheduling = attr.Value
			case "total_jobs":
				tmp_server_state_info.TotalJobs, _ = strconv.ParseInt(attr.Value, 10, 64)
			case "state_count":
				attr_array := strings.Split(attr.Value, " ")
				for _, sc_valu := range attr_array {
					if len(sc_valu) == 0 {
						break
					}
					scname := strings.Split(sc_valu, ":")[0]
					scval := strings.Split(sc_valu, ":")[1]
					switch scname {
					case "Transit":
						tmp_server_state_info.StateCountTransit, _ = strconv.ParseInt(scval, 10, 64)
					case "Queued":
						tmp_server_state_info.StateCountQueued, _ = strconv.ParseInt(scval, 10, 64)
					case "Held":
						tmp_server_state_info.StateCountHeld, _ = strconv.ParseInt(scval, 10, 64)
					case "Waiting":
						tmp_server_state_info.StateCountWaiting, _ = strconv.ParseInt(scval, 10, 64)
					case "Running":
						tmp_server_state_info.StateCountRunning, _ = strconv.ParseInt(scval, 10, 64)
					case "Exiting":
						tmp_server_state_info.StateCountExiting, _ = strconv.ParseInt(scval, 10, 64)
					case "Begun":
						tmp_server_state_info.StateCountBegun, _ = strconv.ParseInt(scval, 10, 64)
					default:
						fmt.Println("other server_state_count")
					}
				}
			case "default_queue":
				tmp_server_state_info.DefaultQueue = attr.Value
			case "log_events":
				tmp_server_state_info.LogEvents, _ = strconv.ParseInt(attr.Value, 10, 64)
			case "mail_from":
				tmp_server_state_info.MailFrom = attr.Value
			case "query_other_jobs":
				tmp_server_state_info.QueryOtherJobs = attr.Value
			case "resources_default":
				tmp_server_state_info.ResourcesDefaultNcpus, _ = strconv.ParseInt(attr.Value, 10, 64)
			case "default_chunk":
				tmp_server_state_info.DefaultChunkNcpus, _ = strconv.ParseInt(attr.Value, 10, 64)
			case "resources_assigned":
				if attr.Resource == "ncpus" {
					tmp_server_state_info.ResourcesAssignedNcpus, _ = strconv.ParseInt(attr.Value, 10, 64)
				}
				if attr.Resource == "nodect" {
					tmp_server_state_info.ResourcesAssignedNodect, _ = strconv.ParseInt(attr.Value, 10, 64)
				}
			case "scheduler_iteration":
				tmp_server_state_info.SchedulerIteration, _ = strconv.ParseInt(attr.Value, 10, 64)
			case "FLicenses":
				tmp_server_state_info.Flicenses, _ = strconv.ParseInt(attr.Value, 10, 64)
			case "resv_enable":
				tmp_server_state_info.ResvEnable = attr.Value
			case "node_fail_requeue":
				tmp_server_state_info.NodeFailRequeue, _ = strconv.ParseInt(attr.Value, 10, 64)
			case "max_array_size":
				tmp_server_state_info.MaxArraySize, _ = strconv.ParseInt(attr.Value, 10, 64)
			case "pbs_license_min":
				tmp_server_state_info.PBSLicenseMin, _ = strconv.ParseInt(attr.Value, 10, 64)
			case "pbs_license_max":
				tmp_server_state_info.PBSLicenseMax, _ = strconv.ParseInt(attr.Value, 10, 64)
			case "pbs_license_linger_time":
				tmp_server_state_info.PBSLicenseLingerTime, _ = strconv.ParseInt(attr.Value, 10, 64)
			case "license_count":
				attrArray := strings.Split(attr.Value, " ")
				for _, scValList := range attrArray {
					scname := strings.Split(scValList, ":")[0]
					scval := strings.Split(scValList, ":")[1]
					switch scname {
					case "Avail_Global":
						tmp_server_state_info.LicenseCountAvailGlobal, _ = strconv.ParseInt(scval, 10, 64)
					case "Avail_Local":
						tmp_server_state_info.LicenseCountAvailLocal, _ = strconv.ParseInt(scval, 10, 64)
					case "Used":
						tmp_server_state_info.LicenseCountUsed, _ = strconv.ParseInt(scval, 10, 64)
					case "High_Use":
						tmp_server_state_info.LicenseCountHighUse, _ = strconv.ParseInt(scval, 10, 64)
					default:
						fmt.Println("other license_count")
					}
				}
			case "pbs_version":
				tmp_server_state_info.PBSVersion = attr.Value
			case "eligible_time_enable":
				tmp_server_state_info.EligibleTimeEnable = attr.Value
			case "job_history_enable":
				tmp_server_state_info.JobHistoryEnable = attr.Value
			case "job_history_duration":
				tmp_server_state_info.JobHistoryDuration = attr.Value
			case "max_concurrent_provision":
				tmp_server_state_info.MaxConcurrentProvision, _ = strconv.ParseInt(attr.Value, 10, 64)
			case "power_provisioning":
				tmp_server_state_info.PowerProvisioning = attr.Value
			default:
				fmt.Println("other server state info.", attr.Name)
			}
		}
		qs.ServerState = append(qs.ServerState, tmp_server_state_info)
	}

	return nil
}

//返回所有作业信息，如果Extend设为x，则返回所有历史信息。
func (qs *Qstat) Pbs_selstat() ([]utils.BatchStatus, error) {
	a := Pbs_attrib2attribl(qs.Attribs)
	defer Pbs_freeattribl(a)

	e := C.CString(qs.Extend)
	defer C.free(unsafe.Pointer(e))

	batch_status := C.pbs_selstat(C.int(qs.Handle), (*C.struct_attropl)(unsafe.Pointer(a)), a, e)

	// FIXME: nil also indicates no jobs matched selection criteria...
	if batch_status == nil {
		return nil, errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)
	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

//获取信息
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
