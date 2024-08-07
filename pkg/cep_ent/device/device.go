// Code generated by ent, DO NOT EDIT.

package device

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the device type in the database.
	Label = "device"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldSerialNumber holds the string denoting the serial_number field in the database.
	FieldSerialNumber = "serial_number"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldSumCep holds the string denoting the sum_cep field in the database.
	FieldSumCep = "sum_cep"
	// FieldLinking holds the string denoting the linking field in the database.
	FieldLinking = "linking"
	// FieldBindingStatus holds the string denoting the binding_status field in the database.
	FieldBindingStatus = "binding_status"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldManageName holds the string denoting the manage_name field in the database.
	FieldManageName = "manage_name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldCoresNumber holds the string denoting the cores_number field in the database.
	FieldCoresNumber = "cores_number"
	// FieldCPU holds the string denoting the cpu field in the database.
	FieldCPU = "cpu"
	// FieldCpus holds the string denoting the cpus field in the database.
	FieldCpus = "cpus"
	// FieldMemory holds the string denoting the memory field in the database.
	FieldMemory = "memory"
	// FieldDisk holds the string denoting the disk field in the database.
	FieldDisk = "disk"
	// FieldDelay holds the string denoting the delay field in the database.
	FieldDelay = "delay"
	// FieldGpuTemperature holds the string denoting the gpu_temperature field in the database.
	FieldGpuTemperature = "gpu_temperature"
	// FieldCPUTemperature holds the string denoting the cpu_temperature field in the database.
	FieldCPUTemperature = "cpu_temperature"
	// FieldStability holds the string denoting the stability field in the database.
	FieldStability = "stability"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldFault holds the string denoting the fault field in the database.
	FieldFault = "fault"
	// FieldRank holds the string denoting the rank field in the database.
	FieldRank = "rank"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeMissionProduceOrders holds the string denoting the mission_produce_orders edge name in mutations.
	EdgeMissionProduceOrders = "mission_produce_orders"
	// EdgeUserDevices holds the string denoting the user_devices edge name in mutations.
	EdgeUserDevices = "user_devices"
	// EdgeDeviceGpuMissions holds the string denoting the device_gpu_missions edge name in mutations.
	EdgeDeviceGpuMissions = "device_gpu_missions"
	// EdgeFrpcInfos holds the string denoting the frpc_infos edge name in mutations.
	EdgeFrpcInfos = "frpc_infos"
	// EdgeMissionOrders holds the string denoting the mission_orders edge name in mutations.
	EdgeMissionOrders = "mission_orders"
	// EdgeMissionProductions holds the string denoting the mission_productions edge name in mutations.
	EdgeMissionProductions = "mission_productions"
	// EdgeDeviceRebootTimes holds the string denoting the device_reboot_times edge name in mutations.
	EdgeDeviceRebootTimes = "device_reboot_times"
	// EdgeTroubleDeducts holds the string denoting the trouble_deducts edge name in mutations.
	EdgeTroubleDeducts = "trouble_deducts"
	// EdgeDeviceStates holds the string denoting the device_states edge name in mutations.
	EdgeDeviceStates = "device_states"
	// EdgeDeviceOfflineRecords holds the string denoting the device_offline_records edge name in mutations.
	EdgeDeviceOfflineRecords = "device_offline_records"
	// EdgeMissionFailedFeedbacks holds the string denoting the mission_failed_feedbacks edge name in mutations.
	EdgeMissionFailedFeedbacks = "mission_failed_feedbacks"
	// Table holds the table name of the device in the database.
	Table = "devices"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "devices"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// MissionProduceOrdersTable is the table that holds the mission_produce_orders relation/edge.
	MissionProduceOrdersTable = "mission_produce_orders"
	// MissionProduceOrdersInverseTable is the table name for the MissionProduceOrder entity.
	// It exists in this package in order to avoid circular dependency with the "missionproduceorder" package.
	MissionProduceOrdersInverseTable = "mission_produce_orders"
	// MissionProduceOrdersColumn is the table column denoting the mission_produce_orders relation/edge.
	MissionProduceOrdersColumn = "device_id"
	// UserDevicesTable is the table that holds the user_devices relation/edge.
	UserDevicesTable = "user_devices"
	// UserDevicesInverseTable is the table name for the UserDevice entity.
	// It exists in this package in order to avoid circular dependency with the "userdevice" package.
	UserDevicesInverseTable = "user_devices"
	// UserDevicesColumn is the table column denoting the user_devices relation/edge.
	UserDevicesColumn = "device_id"
	// DeviceGpuMissionsTable is the table that holds the device_gpu_missions relation/edge.
	DeviceGpuMissionsTable = "device_gpu_missions"
	// DeviceGpuMissionsInverseTable is the table name for the DeviceGpuMission entity.
	// It exists in this package in order to avoid circular dependency with the "devicegpumission" package.
	DeviceGpuMissionsInverseTable = "device_gpu_missions"
	// DeviceGpuMissionsColumn is the table column denoting the device_gpu_missions relation/edge.
	DeviceGpuMissionsColumn = "device_id"
	// FrpcInfosTable is the table that holds the frpc_infos relation/edge.
	FrpcInfosTable = "frpc_infos"
	// FrpcInfosInverseTable is the table name for the FrpcInfo entity.
	// It exists in this package in order to avoid circular dependency with the "frpcinfo" package.
	FrpcInfosInverseTable = "frpc_infos"
	// FrpcInfosColumn is the table column denoting the frpc_infos relation/edge.
	FrpcInfosColumn = "device_id"
	// MissionOrdersTable is the table that holds the mission_orders relation/edge.
	MissionOrdersTable = "mission_orders"
	// MissionOrdersInverseTable is the table name for the MissionOrder entity.
	// It exists in this package in order to avoid circular dependency with the "missionorder" package.
	MissionOrdersInverseTable = "mission_orders"
	// MissionOrdersColumn is the table column denoting the mission_orders relation/edge.
	MissionOrdersColumn = "device_id"
	// MissionProductionsTable is the table that holds the mission_productions relation/edge.
	MissionProductionsTable = "mission_productions"
	// MissionProductionsInverseTable is the table name for the MissionProduction entity.
	// It exists in this package in order to avoid circular dependency with the "missionproduction" package.
	MissionProductionsInverseTable = "mission_productions"
	// MissionProductionsColumn is the table column denoting the mission_productions relation/edge.
	MissionProductionsColumn = "device_id"
	// DeviceRebootTimesTable is the table that holds the device_reboot_times relation/edge.
	DeviceRebootTimesTable = "device_reboot_times"
	// DeviceRebootTimesInverseTable is the table name for the DeviceRebootTime entity.
	// It exists in this package in order to avoid circular dependency with the "devicereboottime" package.
	DeviceRebootTimesInverseTable = "device_reboot_times"
	// DeviceRebootTimesColumn is the table column denoting the device_reboot_times relation/edge.
	DeviceRebootTimesColumn = "device_id"
	// TroubleDeductsTable is the table that holds the trouble_deducts relation/edge.
	TroubleDeductsTable = "trouble_deducts"
	// TroubleDeductsInverseTable is the table name for the TroubleDeduct entity.
	// It exists in this package in order to avoid circular dependency with the "troublededuct" package.
	TroubleDeductsInverseTable = "trouble_deducts"
	// TroubleDeductsColumn is the table column denoting the trouble_deducts relation/edge.
	TroubleDeductsColumn = "device_id"
	// DeviceStatesTable is the table that holds the device_states relation/edge.
	DeviceStatesTable = "device_states"
	// DeviceStatesInverseTable is the table name for the DeviceState entity.
	// It exists in this package in order to avoid circular dependency with the "devicestate" package.
	DeviceStatesInverseTable = "device_states"
	// DeviceStatesColumn is the table column denoting the device_states relation/edge.
	DeviceStatesColumn = "device_id"
	// DeviceOfflineRecordsTable is the table that holds the device_offline_records relation/edge.
	DeviceOfflineRecordsTable = "device_offline_records"
	// DeviceOfflineRecordsInverseTable is the table name for the DeviceOfflineRecord entity.
	// It exists in this package in order to avoid circular dependency with the "deviceofflinerecord" package.
	DeviceOfflineRecordsInverseTable = "device_offline_records"
	// DeviceOfflineRecordsColumn is the table column denoting the device_offline_records relation/edge.
	DeviceOfflineRecordsColumn = "device_id"
	// MissionFailedFeedbacksTable is the table that holds the mission_failed_feedbacks relation/edge.
	MissionFailedFeedbacksTable = "mission_failed_feedbacks"
	// MissionFailedFeedbacksInverseTable is the table name for the MissionFailedFeedback entity.
	// It exists in this package in order to avoid circular dependency with the "missionfailedfeedback" package.
	MissionFailedFeedbacksInverseTable = "mission_failed_feedbacks"
	// MissionFailedFeedbacksColumn is the table column denoting the mission_failed_feedbacks relation/edge.
	MissionFailedFeedbacksColumn = "device_id"
)

// Columns holds all SQL columns for device fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldUserID,
	FieldSerialNumber,
	FieldState,
	FieldSumCep,
	FieldLinking,
	FieldBindingStatus,
	FieldStatus,
	FieldName,
	FieldManageName,
	FieldType,
	FieldCoresNumber,
	FieldCPU,
	FieldCpus,
	FieldMemory,
	FieldDisk,
	FieldDelay,
	FieldGpuTemperature,
	FieldCPUTemperature,
	FieldStability,
	FieldVersion,
	FieldFault,
	FieldRank,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedBy holds the default value on creation for the "created_by" field.
	DefaultCreatedBy int64
	// DefaultUpdatedBy holds the default value on creation for the "updated_by" field.
	DefaultUpdatedBy int64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt time.Time
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID int64
	// DefaultSerialNumber holds the default value on creation for the "serial_number" field.
	DefaultSerialNumber string
	// DefaultSumCep holds the default value on creation for the "sum_cep" field.
	DefaultSumCep int64
	// DefaultLinking holds the default value on creation for the "linking" field.
	DefaultLinking bool
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultManageName holds the default value on creation for the "manage_name" field.
	DefaultManageName string
	// DefaultCoresNumber holds the default value on creation for the "cores_number" field.
	DefaultCoresNumber int64
	// DefaultCPU holds the default value on creation for the "cpu" field.
	DefaultCPU string
	// DefaultMemory holds the default value on creation for the "memory" field.
	DefaultMemory int64
	// DefaultDisk holds the default value on creation for the "disk" field.
	DefaultDisk float32
	// DefaultDelay holds the default value on creation for the "delay" field.
	DefaultDelay float64
	// DefaultGpuTemperature holds the default value on creation for the "gpu_temperature" field.
	DefaultGpuTemperature float64
	// DefaultCPUTemperature holds the default value on creation for the "cpu_temperature" field.
	DefaultCPUTemperature float64
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
	// ValueScanner of all Device fields.
	ValueScanner struct {
		Cpus field.TypeValueScanner[[]string]
	}
)

// State defines the type for the "state" enum field.
type State string

// StateInit is the default value of the State enum.
const DefaultState = StateInit

// State values.
const (
	StateOn   State = "On"
	StateDown State = "Down"
	StateInit State = "Init"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateOn, StateDown, StateInit:
		return nil
	default:
		return fmt.Errorf("device: invalid enum value for state field: %q", s)
	}
}

const DefaultBindingStatus enums.DeviceBindingStatus = "init"

// BindingStatusValidator is a validator for the "binding_status" field enum values. It is called by the builders before save.
func BindingStatusValidator(bs enums.DeviceBindingStatus) error {
	switch bs {
	case "init", "bound", "unbound", "rebinding":
		return nil
	default:
		return fmt.Errorf("device: invalid enum value for binding_status field: %q", bs)
	}
}

const DefaultStatus enums.DeviceStatus = "offline"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s enums.DeviceStatus) error {
	switch s {
	case "online", "offline", "busy", "free", "exit":
		return nil
	default:
		return fmt.Errorf("device: invalid enum value for status field: %q", s)
	}
}

const DefaultType enums.DeviceType = "ordinary"

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.DeviceType) error {
	switch _type {
	case "official", "ordinary":
		return nil
	default:
		return fmt.Errorf("device: invalid enum value for type field: %q", _type)
	}
}

const DefaultStability enums.DeviceStabilityType = "good"

// StabilityValidator is a validator for the "stability" field enum values. It is called by the builders before save.
func StabilityValidator(s enums.DeviceStabilityType) error {
	switch s {
	case "great", "good", "ok", "bad":
		return nil
	default:
		return fmt.Errorf("device: invalid enum value for stability field: %q", s)
	}
}

const DefaultFault enums.DeviceFaultType = "ok"

// FaultValidator is a validator for the "fault" field enum values. It is called by the builders before save.
func FaultValidator(f enums.DeviceFaultType) error {
	switch f {
	case "ok", "drive_abnormal", "no_disk", "task_failure", "recover":
		return nil
	default:
		return fmt.Errorf("device: invalid enum value for fault field: %q", f)
	}
}

const DefaultRank enums.DeviceRankType = "normal"

// RankValidator is a validator for the "rank" field enum values. It is called by the builders before save.
func RankValidator(r enums.DeviceRankType) error {
	switch r {
	case "blask", "normal":
		return nil
	default:
		return fmt.Errorf("device: invalid enum value for rank field: %q", r)
	}
}

// OrderOption defines the ordering options for the Device queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// BySerialNumber orders the results by the serial_number field.
func BySerialNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSerialNumber, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// BySumCep orders the results by the sum_cep field.
func BySumCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSumCep, opts...).ToFunc()
}

// ByLinking orders the results by the linking field.
func ByLinking(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLinking, opts...).ToFunc()
}

// ByBindingStatus orders the results by the binding_status field.
func ByBindingStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBindingStatus, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByManageName orders the results by the manage_name field.
func ByManageName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldManageName, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByCoresNumber orders the results by the cores_number field.
func ByCoresNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoresNumber, opts...).ToFunc()
}

// ByCPU orders the results by the cpu field.
func ByCPU(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCPU, opts...).ToFunc()
}

// ByCpus orders the results by the cpus field.
func ByCpus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCpus, opts...).ToFunc()
}

// ByMemory orders the results by the memory field.
func ByMemory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemory, opts...).ToFunc()
}

// ByDisk orders the results by the disk field.
func ByDisk(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisk, opts...).ToFunc()
}

// ByDelay orders the results by the delay field.
func ByDelay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDelay, opts...).ToFunc()
}

// ByGpuTemperature orders the results by the gpu_temperature field.
func ByGpuTemperature(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGpuTemperature, opts...).ToFunc()
}

// ByCPUTemperature orders the results by the cpu_temperature field.
func ByCPUTemperature(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCPUTemperature, opts...).ToFunc()
}

// ByStability orders the results by the stability field.
func ByStability(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStability, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByFault orders the results by the fault field.
func ByFault(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFault, opts...).ToFunc()
}

// ByRank orders the results by the rank field.
func ByRank(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRank, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByMissionProduceOrdersCount orders the results by mission_produce_orders count.
func ByMissionProduceOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMissionProduceOrdersStep(), opts...)
	}
}

// ByMissionProduceOrders orders the results by mission_produce_orders terms.
func ByMissionProduceOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionProduceOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserDevicesCount orders the results by user_devices count.
func ByUserDevicesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserDevicesStep(), opts...)
	}
}

// ByUserDevices orders the results by user_devices terms.
func ByUserDevices(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserDevicesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDeviceGpuMissionsCount orders the results by device_gpu_missions count.
func ByDeviceGpuMissionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDeviceGpuMissionsStep(), opts...)
	}
}

// ByDeviceGpuMissions orders the results by device_gpu_missions terms.
func ByDeviceGpuMissions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeviceGpuMissionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFrpcInfosCount orders the results by frpc_infos count.
func ByFrpcInfosCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFrpcInfosStep(), opts...)
	}
}

// ByFrpcInfos orders the results by frpc_infos terms.
func ByFrpcInfos(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFrpcInfosStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMissionOrdersCount orders the results by mission_orders count.
func ByMissionOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMissionOrdersStep(), opts...)
	}
}

// ByMissionOrders orders the results by mission_orders terms.
func ByMissionOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMissionProductionsCount orders the results by mission_productions count.
func ByMissionProductionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMissionProductionsStep(), opts...)
	}
}

// ByMissionProductions orders the results by mission_productions terms.
func ByMissionProductions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionProductionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDeviceRebootTimesCount orders the results by device_reboot_times count.
func ByDeviceRebootTimesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDeviceRebootTimesStep(), opts...)
	}
}

// ByDeviceRebootTimes orders the results by device_reboot_times terms.
func ByDeviceRebootTimes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeviceRebootTimesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTroubleDeductsCount orders the results by trouble_deducts count.
func ByTroubleDeductsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTroubleDeductsStep(), opts...)
	}
}

// ByTroubleDeducts orders the results by trouble_deducts terms.
func ByTroubleDeducts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTroubleDeductsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDeviceStatesCount orders the results by device_states count.
func ByDeviceStatesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDeviceStatesStep(), opts...)
	}
}

// ByDeviceStates orders the results by device_states terms.
func ByDeviceStates(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeviceStatesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDeviceOfflineRecordsCount orders the results by device_offline_records count.
func ByDeviceOfflineRecordsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDeviceOfflineRecordsStep(), opts...)
	}
}

// ByDeviceOfflineRecords orders the results by device_offline_records terms.
func ByDeviceOfflineRecords(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeviceOfflineRecordsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMissionFailedFeedbacksCount orders the results by mission_failed_feedbacks count.
func ByMissionFailedFeedbacksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMissionFailedFeedbacksStep(), opts...)
	}
}

// ByMissionFailedFeedbacks orders the results by mission_failed_feedbacks terms.
func ByMissionFailedFeedbacks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionFailedFeedbacksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newMissionProduceOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionProduceOrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MissionProduceOrdersTable, MissionProduceOrdersColumn),
	)
}
func newUserDevicesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserDevicesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserDevicesTable, UserDevicesColumn),
	)
}
func newDeviceGpuMissionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeviceGpuMissionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DeviceGpuMissionsTable, DeviceGpuMissionsColumn),
	)
}
func newFrpcInfosStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FrpcInfosInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FrpcInfosTable, FrpcInfosColumn),
	)
}
func newMissionOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionOrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MissionOrdersTable, MissionOrdersColumn),
	)
}
func newMissionProductionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionProductionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MissionProductionsTable, MissionProductionsColumn),
	)
}
func newDeviceRebootTimesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeviceRebootTimesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DeviceRebootTimesTable, DeviceRebootTimesColumn),
	)
}
func newTroubleDeductsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TroubleDeductsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TroubleDeductsTable, TroubleDeductsColumn),
	)
}
func newDeviceStatesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeviceStatesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DeviceStatesTable, DeviceStatesColumn),
	)
}
func newDeviceOfflineRecordsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeviceOfflineRecordsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DeviceOfflineRecordsTable, DeviceOfflineRecordsColumn),
	)
}
func newMissionFailedFeedbacksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionFailedFeedbacksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MissionFailedFeedbacksTable, MissionFailedFeedbacksColumn),
	)
}
