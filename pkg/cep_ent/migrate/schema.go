// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BillsColumns holds the columns for the "bills" table.
	BillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Comment: "次级账单流水的类型，充值或者任务消耗或任务收益", Enums: []string{"recharge", "mission_consume", "mission_produce"}, Default: "mission_consume"},
		{Name: "is_add", Type: field.TypeBool, Comment: "是否增加余额，布尔值默认为否", Default: false},
		{Name: "serial_number", Type: field.TypeString, Comment: "账单序列号", Default: ""},
		{Name: "cep", Type: field.TypeInt64, Comment: "消费或收益多少 cep", Default: 0},
		{Name: "status", Type: field.TypeEnum, Comment: "消耗流水一开始不能直接生效，确定关联的消耗时间成功后才可以扣费", Enums: []string{"pending", "canceled", "done"}, Default: "pending"},
		{Name: "market_bill_id", Type: field.TypeInt64, Comment: "营销流水 id", Default: 0},
		{Name: "platform_cep", Type: field.TypeInt64, Comment: "平台收取多少本金余额", Default: 0},
		{Name: "reason_id", Type: field.TypeInt64, Nullable: true, Comment: "关联消耗产生的来源外键 id，比如 type 为 mission 时关联任务订单", Default: 0},
		{Name: "platform_wallet_id", Type: field.TypeInt64, Comment: "平台分润钱包 id", Default: 0},
		{Name: "user_id", Type: field.TypeInt64, Comment: "外键用户 id", Default: 0},
		{Name: "wallet_id", Type: field.TypeInt64, Comment: "外键钱包 id", Default: 0},
	}
	// BillsTable holds the schema information for the "bills" table.
	BillsTable = &schema.Table{
		Name:       "bills",
		Comment:    "记录额度账户的变动的主流水账单记录",
		Columns:    BillsColumns,
		PrimaryKey: []*schema.Column{BillsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bills_mission_consume_orders_bills",
				Columns:    []*schema.Column{BillsColumns[13]},
				RefColumns: []*schema.Column{MissionConsumeOrdersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "bills_mission_produce_orders_bills",
				Columns:    []*schema.Column{BillsColumns[13]},
				RefColumns: []*schema.Column{MissionProduceOrdersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "bills_platform_wallets_bills",
				Columns:    []*schema.Column{BillsColumns[14]},
				RefColumns: []*schema.Column{PlatformWalletsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "bills_recharge_orders_bills",
				Columns:    []*schema.Column{BillsColumns[13]},
				RefColumns: []*schema.Column{RechargeOrdersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "bills_users_bills",
				Columns:    []*schema.Column{BillsColumns[15]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "bills_wallets_bills",
				Columns:    []*schema.Column{BillsColumns[16]},
				RefColumns: []*schema.Column{WalletsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// CollectionsColumns holds the columns for the "collections" table.
	CollectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "url", Type: field.TypeString, Comment: "路径", Default: ""},
		{Name: "picture_name", Type: field.TypeString, Comment: "照片名称", Default: ""},
		{Name: "user_id", Type: field.TypeInt64, Comment: "外键用户 id", Default: 0},
	}
	// CollectionsTable holds the schema information for the "collections" table.
	CollectionsTable = &schema.Table{
		Name:       "collections",
		Comment:    "资源收藏，与用户一对多",
		Columns:    CollectionsColumns,
		PrimaryKey: []*schema.Column{CollectionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "collections_users_collections",
				Columns:    []*schema.Column{CollectionsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// DevicesColumns holds the columns for the "devices" table.
	DevicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Comment: "设备状态", Enums: []string{"online", "offline", "busy", "free"}, Default: "online"},
		{Name: "binding_status", Type: field.TypeEnum, Comment: "设备的绑定状态", Enums: []string{"init", "bound", "unbound", "rebinding"}, Default: "init"},
		{Name: "user_id", Type: field.TypeInt64, Comment: "外键用户 id", Default: 0},
	}
	// DevicesTable holds the schema information for the "devices" table.
	DevicesTable = &schema.Table{
		Name:       "devices",
		Comment:    "设备表，与用户一对多",
		Columns:    DevicesColumns,
		PrimaryKey: []*schema.Column{DevicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "devices_users_devices",
				Columns:    []*schema.Column{DevicesColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// HmacKeyPairsColumns holds the columns for the "hmac_key_pairs" table.
	HmacKeyPairsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "key", Type: field.TypeString, Comment: "密钥对的 key 值，用于检索密钥", Default: ""},
		{Name: "secret", Type: field.TypeString, Comment: "加密密钥", Default: ""},
		{Name: "caller", Type: field.TypeString, Comment: "请求方", Default: ""},
		{Name: "user_id", Type: field.TypeInt64, Comment: "外键用户 ID", Default: 0},
		{Name: "user_hmac_key_pair", Type: field.TypeInt64, Unique: true},
	}
	// HmacKeyPairsTable holds the schema information for the "hmac_key_pairs" table.
	HmacKeyPairsTable = &schema.Table{
		Name:       "hmac_key_pairs",
		Comment:    "Hmac 密钥对，用于没有登录态时安全调用任务相关接口的场景",
		Columns:    HmacKeyPairsColumns,
		PrimaryKey: []*schema.Column{HmacKeyPairsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "hmac_key_pairs_users_hmac_key_pair",
				Columns:    []*schema.Column{HmacKeyPairsColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "hmackeypair_key",
				Unique:  false,
				Columns: []*schema.Column{HmacKeyPairsColumns[6]},
			},
		},
	}
	// InputLogsColumns holds the columns for the "input_logs" table.
	InputLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "trace_id", Type: field.TypeInt64, Comment: "请求追踪 id", Default: 0},
		{Name: "headers", Type: field.TypeString, Comment: "请求头", Default: ""},
		{Name: "body", Type: field.TypeString, Nullable: true, Comment: "请求体", Default: ""},
		{Name: "query", Type: field.TypeString, Nullable: true, Comment: "Query 参数", Default: ""},
		{Name: "url", Type: field.TypeString, Comment: "请求地址", Default: ""},
		{Name: "ip", Type: field.TypeString, Comment: "客户端 IP", Default: ""},
		{Name: "caller", Type: field.TypeString, Comment: "调用方", Default: "general"},
		{Name: "method", Type: field.TypeEnum, Comment: "请求方式", Enums: []string{"GET", "POST", "PUT", "DELETE", "PATCH"}, Default: "GET"},
		{Name: "hmac_key", Type: field.TypeString, Comment: "记录调用者的密钥对", Default: ""},
	}
	// InputLogsTable holds the schema information for the "input_logs" table.
	InputLogsTable = &schema.Table{
		Name:       "input_logs",
		Comment:    "接口输入表，用于记录最老的直接任务模式",
		Columns:    InputLogsColumns,
		PrimaryKey: []*schema.Column{InputLogsColumns[0]},
	}
	// MissionsColumns holds the columns for the "missions" table.
	MissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Comment: "任务类型", Enums: []string{"sd_time", "txt2img", "img2img", "jp_time", "wt_time"}, Default: "txt2img"},
		{Name: "is_time", Type: field.TypeBool, Comment: "是否为计时类型任务", Default: false},
		{Name: "body", Type: field.TypeString, Comment: "任务的请求参数体", Default: ""},
		{Name: "call_back_url", Type: field.TypeString, Comment: "回调地址，空字符串表示不回调", Default: ""},
		{Name: "status", Type: field.TypeEnum, Comment: "任务状态", Enums: []string{"waiting", "canceled", "doing", "supplying", "closing", "succeed", "failed"}, Default: "waiting"},
		{Name: "result_urls", Type: field.TypeString, Comment: "任务结果资源位置列表序列化", Default: ""},
		{Name: "additional_result", Type: field.TypeString, Comment: "有的任务除了链接外还有其他有用的结果，都塞在这个字段，比如 sd 的实际入参", Default: ""},
		{Name: "mission_batch_number", Type: field.TypeString, Comment: "任务批次号", Default: ""},
		{Name: "hmac_key_pair_id", Type: field.TypeInt64, Comment: "外键任务创建者的密钥对 ID", Default: 0},
		{Name: "mission_batch_id", Type: field.TypeInt64, Comment: "外键任务批次", Default: 0},
		{Name: "user_id", Type: field.TypeInt64, Comment: "外键任务创建者的 ID", Default: 0},
	}
	// MissionsTable holds the schema information for the "missions" table.
	MissionsTable = &schema.Table{
		Name:       "missions",
		Columns:    MissionsColumns,
		PrimaryKey: []*schema.Column{MissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "missions_hmac_key_pairs_created_missions",
				Columns:    []*schema.Column{MissionsColumns[14]},
				RefColumns: []*schema.Column{HmacKeyPairsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "missions_mission_batches_missions",
				Columns:    []*schema.Column{MissionsColumns[15]},
				RefColumns: []*schema.Column{MissionBatchesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "missions_users_created_missions",
				Columns:    []*schema.Column{MissionsColumns[16]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MissionBatchesColumns holds the columns for the "mission_batches" table.
	MissionBatchesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "number", Type: field.TypeString, Comment: "批次号码", Default: ""},
		{Name: "user_id", Type: field.TypeInt64, Comment: "创建者", Default: 0},
	}
	// MissionBatchesTable holds the schema information for the "mission_batches" table.
	MissionBatchesTable = &schema.Table{
		Name:       "mission_batches",
		Comment:    "任务批次，和任务相关的数据一对多",
		Columns:    MissionBatchesColumns,
		PrimaryKey: []*schema.Column{MissionBatchesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "mission_batches_users_mission_batches",
				Columns:    []*schema.Column{MissionBatchesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MissionConsumeOrdersColumns holds the columns for the "mission_consume_orders" table.
	MissionConsumeOrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "mission_id", Type: field.TypeInt64, Comment: "外键任务 id，关联任务", Default: 0},
		{Name: "status", Type: field.TypeEnum, Comment: "任务订单的状态，注意不强关联任务的状态", Enums: []string{"waiting", "canceled", "doing", "supplying", "closing", "succeed", "failed"}, Default: "waiting"},
		{Name: "cep", Type: field.TypeInt64, Comment: "发布任务需消耗的 cep 量", Default: 0},
		{Name: "type", Type: field.TypeEnum, Comment: "任务类型，等于任务表的类型字段", Enums: []string{"sd_time", "txt2img", "img2img", "jp_time", "wt_time"}, Default: "txt2img"},
		{Name: "is_time", Type: field.TypeBool, Comment: "是否为计时类型任务", Default: false},
		{Name: "call_way", Type: field.TypeEnum, Comment: "调用方式，API 调用或者微信小程序调用", Enums: []string{"api", "yuan_hui", "dev_platform"}, Default: "api"},
		{Name: "serial_number", Type: field.TypeString, Comment: "订单序列号", Default: ""},
		{Name: "started_at", Type: field.TypeTime, Nullable: true, Comment: "任务开始执行时刻"},
		{Name: "finished_at", Type: field.TypeTime, Nullable: true, Comment: "任务结束执行时刻"},
		{Name: "mission_batch_number", Type: field.TypeString, Comment: "任务批次号，用于方便检索", Default: ""},
		{Name: "mission_mission_consume_order", Type: field.TypeInt64, Unique: true},
		{Name: "mission_batch_id", Type: field.TypeInt64, Comment: "外键任务批次 id", Default: 0},
		{Name: "user_id", Type: field.TypeInt64, Comment: "外键关联用户 id", Default: 0},
	}
	// MissionConsumeOrdersTable holds the schema information for the "mission_consume_orders" table.
	MissionConsumeOrdersTable = &schema.Table{
		Name:       "mission_consume_orders",
		Comment:    "任务消费订单，和任务一对一",
		Columns:    MissionConsumeOrdersColumns,
		PrimaryKey: []*schema.Column{MissionConsumeOrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "mission_consume_orders_missions_mission_consume_order",
				Columns:    []*schema.Column{MissionConsumeOrdersColumns[16]},
				RefColumns: []*schema.Column{MissionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "mission_consume_orders_mission_batches_mission_consume_orders",
				Columns:    []*schema.Column{MissionConsumeOrdersColumns[17]},
				RefColumns: []*schema.Column{MissionBatchesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "mission_consume_orders_users_mission_consume_orders",
				Columns:    []*schema.Column{MissionConsumeOrdersColumns[18]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MissionProduceOrdersColumns holds the columns for the "mission_produce_orders" table.
	MissionProduceOrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Comment: "任务订单的状态，注意不强关联任务的状态", Enums: []string{"waiting", "canceled", "doing", "supplying", "closing", "succeed", "failed"}, Default: "waiting"},
		{Name: "cep", Type: field.TypeInt64, Comment: "完成任务收益的 cep 量", Default: 0},
		{Name: "type", Type: field.TypeEnum, Comment: "任务类型，计时或者次数任务", Enums: []string{"sd_time", "txt2img", "img2img", "jp_time", "wt_time"}, Default: "txt2img"},
		{Name: "is_time", Type: field.TypeBool, Comment: "是否为计时类型任务", Default: false},
		{Name: "serial_number", Type: field.TypeString, Comment: "订单序列号", Default: ""},
		{Name: "device_id", Type: field.TypeInt64, Comment: "生产者接该任务用的设备 id", Default: 0},
		{Name: "mission_id", Type: field.TypeInt64, Comment: "外键任务 id，关联任务中枢的任务", Default: 0},
		{Name: "mission_batch_id", Type: field.TypeInt64, Comment: "外键任务批次", Default: 0},
		{Name: "mission_consume_order_id", Type: field.TypeInt64, Comment: "外键任务消费订单，一个任务消费订单可能会对应多个任务生产订单", Default: 0},
		{Name: "mission_production_id", Type: field.TypeInt64, Unique: true, Comment: "外键任务生产情况 id", Default: 0},
		{Name: "user_id", Type: field.TypeInt64, Comment: "外键关联用户 id", Default: 0},
	}
	// MissionProduceOrdersTable holds the schema information for the "mission_produce_orders" table.
	MissionProduceOrdersTable = &schema.Table{
		Name:       "mission_produce_orders",
		Comment:    "任务生产订单",
		Columns:    MissionProduceOrdersColumns,
		PrimaryKey: []*schema.Column{MissionProduceOrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "mission_produce_orders_devices_mission_produce_orders",
				Columns:    []*schema.Column{MissionProduceOrdersColumns[11]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "mission_produce_orders_missions_mission_produce_orders",
				Columns:    []*schema.Column{MissionProduceOrdersColumns[12]},
				RefColumns: []*schema.Column{MissionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "mission_produce_orders_mission_batches_mission_produce_orders",
				Columns:    []*schema.Column{MissionProduceOrdersColumns[13]},
				RefColumns: []*schema.Column{MissionBatchesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "mission_produce_orders_mission_consume_orders_mission_produce_orders",
				Columns:    []*schema.Column{MissionProduceOrdersColumns[14]},
				RefColumns: []*schema.Column{MissionConsumeOrdersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "mission_produce_orders_mission_productions_mission_produce_order",
				Columns:    []*schema.Column{MissionProduceOrdersColumns[15]},
				RefColumns: []*schema.Column{MissionProductionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "mission_produce_orders_users_mission_produce_orders",
				Columns:    []*schema.Column{MissionProduceOrdersColumns[16]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MissionProductionsColumns holds the columns for the "mission_productions" table.
	MissionProductionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "started_at", Type: field.TypeTime, Comment: "任务开始时刻"},
		{Name: "finished_at", Type: field.TypeTime, Comment: "任务完成时刻"},
		{Name: "status", Type: field.TypeEnum, Comment: "任务结果", Enums: []string{"waiting", "canceled", "doing", "supplying", "closing", "succeed", "failed"}, Default: "doing"},
		{Name: "result_urls", Type: field.TypeString, Comment: "任务结果资源位置列表序列化", Default: ""},
		{Name: "additional_result", Type: field.TypeString, Comment: "额外需要返回的结果数据，格式不定", Default: ""},
		{Name: "device_id", Type: field.TypeInt64, Comment: "领到任务的设备 ID", Default: 0},
		{Name: "hmac_key_pair_id", Type: field.TypeInt64, Comment: "密钥对 ID"},
		{Name: "mission_id", Type: field.TypeInt64, Comment: "任务 ID"},
	}
	// MissionProductionsTable holds the schema information for the "mission_productions" table.
	MissionProductionsTable = &schema.Table{
		Name:       "mission_productions",
		Comment:    "任务生产记录，任务被人接了就会产生一条记录，这一次任务完成情况就在这，同一个任务可以被多次接",
		Columns:    MissionProductionsColumns,
		PrimaryKey: []*schema.Column{MissionProductionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "mission_productions_devices_mission_productions",
				Columns:    []*schema.Column{MissionProductionsColumns[11]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "mission_productions_hmac_key_pairs_mission_productions",
				Columns:    []*schema.Column{MissionProductionsColumns[12]},
				RefColumns: []*schema.Column{HmacKeyPairsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "mission_productions_missions_mission_productions",
				Columns:    []*schema.Column{MissionProductionsColumns[13]},
				RefColumns: []*schema.Column{MissionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MissionTypesColumns holds the columns for the "mission_types" table.
	MissionTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Comment: "任务类型", Enums: []string{"sd_time", "txt2img", "img2img", "jp_time", "wt_time"}, Default: "txt2img"},
		{Name: "gpu", Type: field.TypeEnum, Comment: "显卡型号", Enums: []string{"3070", "3070Ti", "3080", "3080Ti", "3090", "3090Ti", "4070", "4070Ti", "4080", "4080Ti", "4090", "4090Ti", "A100", "V100"}, Default: "3070"},
		{Name: "cep", Type: field.TypeInt64, Comment: "单价消耗 cep", Default: 0},
		{Name: "is_time", Type: field.TypeBool, Comment: "是否计时任务", Default: false},
		{Name: "category", Type: field.TypeEnum, Comment: "任务种类，SD，Jupyter 等", Enums: []string{"SD", "Jupyter", "WeTTy"}, Default: "SD"},
	}
	// MissionTypesTable holds the schema information for the "mission_types" table.
	MissionTypesTable = &schema.Table{
		Name:       "mission_types",
		Comment:    "任务类型等信息，给任务定价归类等",
		Columns:    MissionTypesColumns,
		PrimaryKey: []*schema.Column{MissionTypesColumns[0]},
	}
	// OutputLogsColumns holds the columns for the "output_logs" table.
	OutputLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "trace_id", Type: field.TypeInt64, Comment: "请求追踪 id", Default: 0},
		{Name: "headers", Type: field.TypeString, Comment: "请求头", Default: ""},
		{Name: "body", Type: field.TypeString, Nullable: true, Comment: "请求体", Default: ""},
		{Name: "url", Type: field.TypeString, Comment: "请求地址", Default: ""},
		{Name: "ip", Type: field.TypeString, Nullable: true, Comment: "客户端 IP", Default: ""},
		{Name: "caller", Type: field.TypeString, Comment: "调用方", Default: "general"},
		{Name: "status", Type: field.TypeInt16, Nullable: true, Comment: "状态码", Default: 200},
		{Name: "hmac_key", Type: field.TypeString, Comment: "记录调用者的密钥对", Default: ""},
	}
	// OutputLogsTable holds the schema information for the "output_logs" table.
	OutputLogsTable = &schema.Table{
		Name:       "output_logs",
		Comment:    "输出表，用于最老的直连方式任务计费",
		Columns:    OutputLogsColumns,
		PrimaryKey: []*schema.Column{OutputLogsColumns[0]},
	}
	// PlatformWalletsColumns holds the columns for the "platform_wallets" table.
	PlatformWalletsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"profit"}, Default: "profit"},
		{Name: "sum_cep", Type: field.TypeInt64, Comment: "累计总余额", Default: 0},
		{Name: "cep", Type: field.TypeInt64, Comment: "剩余总余额", Default: 0},
	}
	// PlatformWalletsTable holds the schema information for the "platform_wallets" table.
	PlatformWalletsTable = &schema.Table{
		Name:       "platform_wallets",
		Comment:    "平台钱包账户，独立于用户，每种类型的 cep 余额，比如分润获取的都集中到一条数据，即一个钱包",
		Columns:    PlatformWalletsColumns,
		PrimaryKey: []*schema.Column{PlatformWalletsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "platformwallet_type_deleted_at",
				Unique:  true,
				Columns: []*schema.Column{PlatformWalletsColumns[6], PlatformWalletsColumns[5]},
			},
		},
	}
	// ProfitSettingsColumns holds the columns for the "profit_settings" table.
	ProfitSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "ratio", Type: field.TypeInt64, Comment: "分润比例", Default: 75},
		{Name: "user_id", Type: field.TypeInt64, Comment: "外键用户 id", Default: 0},
	}
	// ProfitSettingsTable holds the schema information for the "profit_settings" table.
	ProfitSettingsTable = &schema.Table{
		Name:       "profit_settings",
		Comment:    "用户的分润配置，与用户一对多，但逻辑上主要是一对一",
		Columns:    ProfitSettingsColumns,
		PrimaryKey: []*schema.Column{ProfitSettingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "profit_settings_users_profit_settings",
				Columns:    []*schema.Column{ProfitSettingsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RechargeOrdersColumns holds the columns for the "recharge_orders" table.
	RechargeOrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Comment: "充值订单的状态，比如微信发起支付后可能没完成支付", Enums: []string{"waiting", "canceled", "doing", "supplying", "closing", "succeed", "failed"}, Default: "doing"},
		{Name: "cep", Type: field.TypeInt64, Comment: "充值多少 cep", Default: 0},
		{Name: "type", Type: field.TypeEnum, Comment: "充值订单的类型", Enums: []string{"manual", "vx", "alipay"}, Default: "manual"},
		{Name: "serial_number", Type: field.TypeString, Comment: "充值订单的序列号", Default: ""},
		{Name: "third_api_resp", Type: field.TypeString, Comment: "第三方平台的返回，给到前端才能发起支付", Default: ""},
		{Name: "from_user_id", Type: field.TypeInt64, Comment: "由谁发起的充值", Default: 0},
		{Name: "out_transaction_id", Type: field.TypeString, Comment: "平台方订单号", Default: ""},
		{Name: "user_id", Type: field.TypeInt64, Comment: "充值的用户 id", Default: 0},
		{Name: "social_id", Type: field.TypeInt64, Nullable: true, Comment: "关联充值来源的身份源 id", Default: 0},
	}
	// RechargeOrdersTable holds the schema information for the "recharge_orders" table.
	RechargeOrdersTable = &schema.Table{
		Name:       "recharge_orders",
		Comment:    "充值订单，微信订单、支付宝订单、手动充值都在一张表",
		Columns:    RechargeOrdersColumns,
		PrimaryKey: []*schema.Column{RechargeOrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "recharge_orders_users_recharge_orders",
				Columns:    []*schema.Column{RechargeOrdersColumns[13]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "recharge_orders_vx_socials_recharge_orders",
				Columns:    []*schema.Column{RechargeOrdersColumns[14]},
				RefColumns: []*schema.Column{VxSocialsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Comment: "用户名"},
		{Name: "nick_name", Type: field.TypeString, Comment: "用户昵称", Default: "请设置昵称"},
		{Name: "phone", Type: field.TypeString, Comment: "手机号", Default: ""},
		{Name: "password", Type: field.TypeString, Comment: "密码", Default: ""},
		{Name: "avatar_url", Type: field.TypeString, Comment: "头像路径", Default: ""},
		{Name: "status", Type: field.TypeEnum, Comment: "用户状态", Enums: []string{"normal", "frozen"}, Default: "normal"},
		{Name: "type", Type: field.TypeEnum, Comment: "用户类型", Enums: []string{"personal", "enterprise"}, Default: "personal"},
		{Name: "platform", Type: field.TypeInt, Comment: "用户可以在什么平台登录，二进制开关数据", Default: 0},
		{Name: "hmac_key", Type: field.TypeString, Comment: "用户使用任务相关功能的密钥对的键，唯一", Default: ""},
		{Name: "hmac_secret", Type: field.TypeString, Comment: "用户使用任务相关功能的密钥对的值", Default: ""},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Comment:    "用户表，手机号唯一，用户名唯一，hmac_key 唯一",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_name_deleted_at",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[6], UsersColumns[5]},
			},
			{
				Name:    "user_phone_deleted_at",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[8], UsersColumns[5]},
			},
			{
				Name:    "user_hmac_key_deleted_at",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[14], UsersColumns[5]},
			},
		},
	}
	// UserDevicesColumns holds the columns for the "user_devices" table.
	UserDevicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "device_id", Type: field.TypeInt64, Comment: "外键设备 id", Default: 0},
		{Name: "user_id", Type: field.TypeInt64, Comment: "外键用户 id", Default: 0},
	}
	// UserDevicesTable holds the schema information for the "user_devices" table.
	UserDevicesTable = &schema.Table{
		Name:       "user_devices",
		Comment:    "用户和设备的记录，主要记录绑定时长和过程",
		Columns:    UserDevicesColumns,
		PrimaryKey: []*schema.Column{UserDevicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_devices_devices_user_devices",
				Columns:    []*schema.Column{UserDevicesColumns[6]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_devices_users_user_devices",
				Columns:    []*schema.Column{UserDevicesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// VxSocialsColumns holds the columns for the "vx_socials" table.
	VxSocialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "app_id", Type: field.TypeString, Comment: "微信应用 id", Default: ""},
		{Name: "open_id", Type: field.TypeString, Comment: "微信身份源的 open_id", Default: ""},
		{Name: "union_id", Type: field.TypeString, Comment: "微信身份源的 union_id", Default: ""},
		{Name: "scope", Type: field.TypeEnum, Comment: "账户的权限级别，一般为 base", Enums: []string{"base"}, Default: "base"},
		{Name: "session_key", Type: field.TypeString, Comment: "小程序专用的会话密钥，不可以返回给前端", Default: ""},
		{Name: "access_token", Type: field.TypeString, Comment: "微信能力访问凭证", Default: ""},
		{Name: "refresh_token", Type: field.TypeString, Comment: "刷新微信凭证的刷新凭证", Default: ""},
		{Name: "user_id", Type: field.TypeInt64, Comment: "外键用户 id", Default: 0},
	}
	// VxSocialsTable holds the schema information for the "vx_socials" table.
	VxSocialsTable = &schema.Table{
		Name:       "vx_socials",
		Comment:    "微信身份源，与用户一对多，微信订单依赖于它",
		Columns:    VxSocialsColumns,
		PrimaryKey: []*schema.Column{VxSocialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "vx_socials_users_vx_socials",
				Columns:    []*schema.Column{VxSocialsColumns[13]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// WalletsColumns holds the columns for the "wallets" table.
	WalletsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "cep", Type: field.TypeInt64, Comment: "当前余额", Default: 0},
		{Name: "sum_cep", Type: field.TypeInt64, Comment: "累计余额", Default: 0},
		{Name: "user_id", Type: field.TypeInt64, Unique: true, Comment: "外键用户 id", Default: 0},
	}
	// WalletsTable holds the schema information for the "wallets" table.
	WalletsTable = &schema.Table{
		Name:       "wallets",
		Comment:    "钱包，与用户一对一",
		Columns:    WalletsColumns,
		PrimaryKey: []*schema.Column{WalletsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "wallets_users_wallet",
				Columns:    []*schema.Column{WalletsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BillsTable,
		CollectionsTable,
		DevicesTable,
		HmacKeyPairsTable,
		InputLogsTable,
		MissionsTable,
		MissionBatchesTable,
		MissionConsumeOrdersTable,
		MissionProduceOrdersTable,
		MissionProductionsTable,
		MissionTypesTable,
		OutputLogsTable,
		PlatformWalletsTable,
		ProfitSettingsTable,
		RechargeOrdersTable,
		UsersTable,
		UserDevicesTable,
		VxSocialsTable,
		WalletsTable,
	}
)

func init() {
	BillsTable.ForeignKeys[0].RefTable = MissionConsumeOrdersTable
	BillsTable.ForeignKeys[1].RefTable = MissionProduceOrdersTable
	BillsTable.ForeignKeys[2].RefTable = PlatformWalletsTable
	BillsTable.ForeignKeys[3].RefTable = RechargeOrdersTable
	BillsTable.ForeignKeys[4].RefTable = UsersTable
	BillsTable.ForeignKeys[5].RefTable = WalletsTable
	BillsTable.Annotation = &entsql.Annotation{}
	CollectionsTable.ForeignKeys[0].RefTable = UsersTable
	CollectionsTable.Annotation = &entsql.Annotation{}
	DevicesTable.ForeignKeys[0].RefTable = UsersTable
	DevicesTable.Annotation = &entsql.Annotation{}
	HmacKeyPairsTable.ForeignKeys[0].RefTable = UsersTable
	HmacKeyPairsTable.Annotation = &entsql.Annotation{}
	InputLogsTable.Annotation = &entsql.Annotation{}
	MissionsTable.ForeignKeys[0].RefTable = HmacKeyPairsTable
	MissionsTable.ForeignKeys[1].RefTable = MissionBatchesTable
	MissionsTable.ForeignKeys[2].RefTable = UsersTable
	MissionsTable.Annotation = &entsql.Annotation{}
	MissionBatchesTable.ForeignKeys[0].RefTable = UsersTable
	MissionBatchesTable.Annotation = &entsql.Annotation{}
	MissionConsumeOrdersTable.ForeignKeys[0].RefTable = MissionsTable
	MissionConsumeOrdersTable.ForeignKeys[1].RefTable = MissionBatchesTable
	MissionConsumeOrdersTable.ForeignKeys[2].RefTable = UsersTable
	MissionConsumeOrdersTable.Annotation = &entsql.Annotation{}
	MissionProduceOrdersTable.ForeignKeys[0].RefTable = DevicesTable
	MissionProduceOrdersTable.ForeignKeys[1].RefTable = MissionsTable
	MissionProduceOrdersTable.ForeignKeys[2].RefTable = MissionBatchesTable
	MissionProduceOrdersTable.ForeignKeys[3].RefTable = MissionConsumeOrdersTable
	MissionProduceOrdersTable.ForeignKeys[4].RefTable = MissionProductionsTable
	MissionProduceOrdersTable.ForeignKeys[5].RefTable = UsersTable
	MissionProduceOrdersTable.Annotation = &entsql.Annotation{}
	MissionProductionsTable.ForeignKeys[0].RefTable = DevicesTable
	MissionProductionsTable.ForeignKeys[1].RefTable = HmacKeyPairsTable
	MissionProductionsTable.ForeignKeys[2].RefTable = MissionsTable
	MissionProductionsTable.Annotation = &entsql.Annotation{}
	MissionTypesTable.Annotation = &entsql.Annotation{}
	OutputLogsTable.Annotation = &entsql.Annotation{}
	PlatformWalletsTable.Annotation = &entsql.Annotation{}
	ProfitSettingsTable.ForeignKeys[0].RefTable = UsersTable
	ProfitSettingsTable.Annotation = &entsql.Annotation{}
	RechargeOrdersTable.ForeignKeys[0].RefTable = UsersTable
	RechargeOrdersTable.ForeignKeys[1].RefTable = VxSocialsTable
	RechargeOrdersTable.Annotation = &entsql.Annotation{}
	UsersTable.Annotation = &entsql.Annotation{}
	UserDevicesTable.ForeignKeys[0].RefTable = DevicesTable
	UserDevicesTable.ForeignKeys[1].RefTable = UsersTable
	UserDevicesTable.Annotation = &entsql.Annotation{}
	VxSocialsTable.ForeignKeys[0].RefTable = UsersTable
	VxSocialsTable.Annotation = &entsql.Annotation{}
	WalletsTable.ForeignKeys[0].RefTable = UsersTable
	WalletsTable.Annotation = &entsql.Annotation{}
}
