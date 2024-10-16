// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// ApiToken is the client for interacting with the ApiToken builders.
	ApiToken *ApiTokenClient
	// Artwork is the client for interacting with the Artwork builders.
	Artwork *ArtworkClient
	// ArtworkLike is the client for interacting with the ArtworkLike builders.
	ArtworkLike *ArtworkLikeClient
	// Bill is the client for interacting with the Bill builders.
	Bill *BillClient
	// CDKInfo is the client for interacting with the CDKInfo builders.
	CDKInfo *CDKInfoClient
	// Campaign is the client for interacting with the Campaign builders.
	Campaign *CampaignClient
	// CampaignOrder is the client for interacting with the CampaignOrder builders.
	CampaignOrder *CampaignOrderClient
	// ClientVersion is the client for interacting with the ClientVersion builders.
	ClientVersion *ClientVersionClient
	// CloudFile is the client for interacting with the CloudFile builders.
	CloudFile *CloudFileClient
	// Collect is the client for interacting with the Collect builders.
	Collect *CollectClient
	// CostAccount is the client for interacting with the CostAccount builders.
	CostAccount *CostAccountClient
	// CostBill is the client for interacting with the CostBill builders.
	CostBill *CostBillClient
	// Device is the client for interacting with the Device builders.
	Device *DeviceClient
	// DeviceGpuMission is the client for interacting with the DeviceGpuMission builders.
	DeviceGpuMission *DeviceGpuMissionClient
	// DeviceOfflineRecord is the client for interacting with the DeviceOfflineRecord builders.
	DeviceOfflineRecord *DeviceOfflineRecordClient
	// DeviceRebootTime is the client for interacting with the DeviceRebootTime builders.
	DeviceRebootTime *DeviceRebootTimeClient
	// DeviceState is the client for interacting with the DeviceState builders.
	DeviceState *DeviceStateClient
	// EarnBill is the client for interacting with the EarnBill builders.
	EarnBill *EarnBillClient
	// EnumCondition is the client for interacting with the EnumCondition builders.
	EnumCondition *EnumConditionClient
	// EnumMissionStatus is the client for interacting with the EnumMissionStatus builders.
	EnumMissionStatus *EnumMissionStatusClient
	// ExtraService is the client for interacting with the ExtraService builders.
	ExtraService *ExtraServiceClient
	// ExtraServiceOrder is the client for interacting with the ExtraServiceOrder builders.
	ExtraServiceOrder *ExtraServiceOrderClient
	// ExtraServicePrice is the client for interacting with the ExtraServicePrice builders.
	ExtraServicePrice *ExtraServicePriceClient
	// FrpcInfo is the client for interacting with the FrpcInfo builders.
	FrpcInfo *FrpcInfoClient
	// FrpsInfo is the client for interacting with the FrpsInfo builders.
	FrpsInfo *FrpsInfoClient
	// GiftMissionConfig is the client for interacting with the GiftMissionConfig builders.
	GiftMissionConfig *GiftMissionConfigClient
	// Gpu is the client for interacting with the Gpu builders.
	Gpu *GpuClient
	// GpuPeak is the client for interacting with the GpuPeak builders.
	GpuPeak *GpuPeakClient
	// HmacKeyPair is the client for interacting with the HmacKeyPair builders.
	HmacKeyPair *HmacKeyPairClient
	// IncomeManage is the client for interacting with the IncomeManage builders.
	IncomeManage *IncomeManageClient
	// InputLog is the client for interacting with the InputLog builders.
	InputLog *InputLogClient
	// Invite is the client for interacting with the Invite builders.
	Invite *InviteClient
	// InvokeModelOrder is the client for interacting with the InvokeModelOrder builders.
	InvokeModelOrder *InvokeModelOrderClient
	// LoginRecord is the client for interacting with the LoginRecord builders.
	LoginRecord *LoginRecordClient
	// Lotto is the client for interacting with the Lotto builders.
	Lotto *LottoClient
	// LottoChanceRule is the client for interacting with the LottoChanceRule builders.
	LottoChanceRule *LottoChanceRuleClient
	// LottoGetCountRecord is the client for interacting with the LottoGetCountRecord builders.
	LottoGetCountRecord *LottoGetCountRecordClient
	// LottoPrize is the client for interacting with the LottoPrize builders.
	LottoPrize *LottoPrizeClient
	// LottoRecord is the client for interacting with the LottoRecord builders.
	LottoRecord *LottoRecordClient
	// LottoUserCount is the client for interacting with the LottoUserCount builders.
	LottoUserCount *LottoUserCountClient
	// Mission is the client for interacting with the Mission builders.
	Mission *MissionClient
	// MissionBatch is the client for interacting with the MissionBatch builders.
	MissionBatch *MissionBatchClient
	// MissionCategory is the client for interacting with the MissionCategory builders.
	MissionCategory *MissionCategoryClient
	// MissionConsumeOrder is the client for interacting with the MissionConsumeOrder builders.
	MissionConsumeOrder *MissionConsumeOrderClient
	// MissionExtraService is the client for interacting with the MissionExtraService builders.
	MissionExtraService *MissionExtraServiceClient
	// MissionFailedFeedback is the client for interacting with the MissionFailedFeedback builders.
	MissionFailedFeedback *MissionFailedFeedbackClient
	// MissionKeyPair is the client for interacting with the MissionKeyPair builders.
	MissionKeyPair *MissionKeyPairClient
	// MissionKind is the client for interacting with the MissionKind builders.
	MissionKind *MissionKindClient
	// MissionLoadBalance is the client for interacting with the MissionLoadBalance builders.
	MissionLoadBalance *MissionLoadBalanceClient
	// MissionLoadBalanceAccess is the client for interacting with the MissionLoadBalanceAccess builders.
	MissionLoadBalanceAccess *MissionLoadBalanceAccessClient
	// MissionOrder is the client for interacting with the MissionOrder builders.
	MissionOrder *MissionOrderClient
	// MissionProduceOrder is the client for interacting with the MissionProduceOrder builders.
	MissionProduceOrder *MissionProduceOrderClient
	// MissionProduction is the client for interacting with the MissionProduction builders.
	MissionProduction *MissionProductionClient
	// Model is the client for interacting with the Model builders.
	Model *ModelClient
	// ModelPrice is the client for interacting with the ModelPrice builders.
	ModelPrice *ModelPriceClient
	// OutputLog is the client for interacting with the OutputLog builders.
	OutputLog *OutputLogClient
	// PlatformAccount is the client for interacting with the PlatformAccount builders.
	PlatformAccount *PlatformAccountClient
	// Price is the client for interacting with the Price builders.
	Price *PriceClient
	// ProfitAccount is the client for interacting with the ProfitAccount builders.
	ProfitAccount *ProfitAccountClient
	// ProfitSetting is the client for interacting with the ProfitSetting builders.
	ProfitSetting *ProfitSettingClient
	// RechargeCampaignRule is the client for interacting with the RechargeCampaignRule builders.
	RechargeCampaignRule *RechargeCampaignRuleClient
	// RechargeCampaignRuleOversea is the client for interacting with the RechargeCampaignRuleOversea builders.
	RechargeCampaignRuleOversea *RechargeCampaignRuleOverseaClient
	// RechargeOrder is the client for interacting with the RechargeOrder builders.
	RechargeOrder *RechargeOrderClient
	// RenewalAgreement is the client for interacting with the RenewalAgreement builders.
	RenewalAgreement *RenewalAgreementClient
	// Survey is the client for interacting with the Survey builders.
	Survey *SurveyClient
	// SurveyAnswer is the client for interacting with the SurveyAnswer builders.
	SurveyAnswer *SurveyAnswerClient
	// SurveyQuestion is the client for interacting with the SurveyQuestion builders.
	SurveyQuestion *SurveyQuestionClient
	// SurveyResponse is the client for interacting with the SurveyResponse builders.
	SurveyResponse *SurveyResponseClient
	// Symbol is the client for interacting with the Symbol builders.
	Symbol *SymbolClient
	// TransferOrder is the client for interacting with the TransferOrder builders.
	TransferOrder *TransferOrderClient
	// TroubleDeduct is the client for interacting with the TroubleDeduct builders.
	TroubleDeduct *TroubleDeductClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserDevice is the client for interacting with the UserDevice builders.
	UserDevice *UserDeviceClient
	// UserModel is the client for interacting with the UserModel builders.
	UserModel *UserModelClient
	// VXAccount is the client for interacting with the VXAccount builders.
	VXAccount *VXAccountClient
	// VXSocial is the client for interacting with the VXSocial builders.
	VXSocial *VXSocialClient
	// Wallet is the client for interacting with the Wallet builders.
	Wallet *WalletClient
	// WithdrawAccount is the client for interacting with the WithdrawAccount builders.
	WithdrawAccount *WithdrawAccountClient
	// WithdrawRecord is the client for interacting with the WithdrawRecord builders.
	WithdrawRecord *WithdrawRecordClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once
	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Commit method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	txDriver.mu.Lock()
	hooks := append([]CommitHook(nil), txDriver.onCommit...)
	txDriver.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	txDriver := tx.config.driver.(*txDriver)
	txDriver.mu.Lock()
	txDriver.onCommit = append(txDriver.onCommit, f)
	txDriver.mu.Unlock()
}

type (
	// Rollbacker is the interface that wraps the Rollback method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	txDriver.mu.Lock()
	hooks := append([]RollbackHook(nil), txDriver.onRollback...)
	txDriver.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	txDriver := tx.config.driver.(*txDriver)
	txDriver.mu.Lock()
	txDriver.onRollback = append(txDriver.onRollback, f)
	txDriver.mu.Unlock()
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.ApiToken = NewApiTokenClient(tx.config)
	tx.Artwork = NewArtworkClient(tx.config)
	tx.ArtworkLike = NewArtworkLikeClient(tx.config)
	tx.Bill = NewBillClient(tx.config)
	tx.CDKInfo = NewCDKInfoClient(tx.config)
	tx.Campaign = NewCampaignClient(tx.config)
	tx.CampaignOrder = NewCampaignOrderClient(tx.config)
	tx.ClientVersion = NewClientVersionClient(tx.config)
	tx.CloudFile = NewCloudFileClient(tx.config)
	tx.Collect = NewCollectClient(tx.config)
	tx.CostAccount = NewCostAccountClient(tx.config)
	tx.CostBill = NewCostBillClient(tx.config)
	tx.Device = NewDeviceClient(tx.config)
	tx.DeviceGpuMission = NewDeviceGpuMissionClient(tx.config)
	tx.DeviceOfflineRecord = NewDeviceOfflineRecordClient(tx.config)
	tx.DeviceRebootTime = NewDeviceRebootTimeClient(tx.config)
	tx.DeviceState = NewDeviceStateClient(tx.config)
	tx.EarnBill = NewEarnBillClient(tx.config)
	tx.EnumCondition = NewEnumConditionClient(tx.config)
	tx.EnumMissionStatus = NewEnumMissionStatusClient(tx.config)
	tx.ExtraService = NewExtraServiceClient(tx.config)
	tx.ExtraServiceOrder = NewExtraServiceOrderClient(tx.config)
	tx.ExtraServicePrice = NewExtraServicePriceClient(tx.config)
	tx.FrpcInfo = NewFrpcInfoClient(tx.config)
	tx.FrpsInfo = NewFrpsInfoClient(tx.config)
	tx.GiftMissionConfig = NewGiftMissionConfigClient(tx.config)
	tx.Gpu = NewGpuClient(tx.config)
	tx.GpuPeak = NewGpuPeakClient(tx.config)
	tx.HmacKeyPair = NewHmacKeyPairClient(tx.config)
	tx.IncomeManage = NewIncomeManageClient(tx.config)
	tx.InputLog = NewInputLogClient(tx.config)
	tx.Invite = NewInviteClient(tx.config)
	tx.InvokeModelOrder = NewInvokeModelOrderClient(tx.config)
	tx.LoginRecord = NewLoginRecordClient(tx.config)
	tx.Lotto = NewLottoClient(tx.config)
	tx.LottoChanceRule = NewLottoChanceRuleClient(tx.config)
	tx.LottoGetCountRecord = NewLottoGetCountRecordClient(tx.config)
	tx.LottoPrize = NewLottoPrizeClient(tx.config)
	tx.LottoRecord = NewLottoRecordClient(tx.config)
	tx.LottoUserCount = NewLottoUserCountClient(tx.config)
	tx.Mission = NewMissionClient(tx.config)
	tx.MissionBatch = NewMissionBatchClient(tx.config)
	tx.MissionCategory = NewMissionCategoryClient(tx.config)
	tx.MissionConsumeOrder = NewMissionConsumeOrderClient(tx.config)
	tx.MissionExtraService = NewMissionExtraServiceClient(tx.config)
	tx.MissionFailedFeedback = NewMissionFailedFeedbackClient(tx.config)
	tx.MissionKeyPair = NewMissionKeyPairClient(tx.config)
	tx.MissionKind = NewMissionKindClient(tx.config)
	tx.MissionLoadBalance = NewMissionLoadBalanceClient(tx.config)
	tx.MissionLoadBalanceAccess = NewMissionLoadBalanceAccessClient(tx.config)
	tx.MissionOrder = NewMissionOrderClient(tx.config)
	tx.MissionProduceOrder = NewMissionProduceOrderClient(tx.config)
	tx.MissionProduction = NewMissionProductionClient(tx.config)
	tx.Model = NewModelClient(tx.config)
	tx.ModelPrice = NewModelPriceClient(tx.config)
	tx.OutputLog = NewOutputLogClient(tx.config)
	tx.PlatformAccount = NewPlatformAccountClient(tx.config)
	tx.Price = NewPriceClient(tx.config)
	tx.ProfitAccount = NewProfitAccountClient(tx.config)
	tx.ProfitSetting = NewProfitSettingClient(tx.config)
	tx.RechargeCampaignRule = NewRechargeCampaignRuleClient(tx.config)
	tx.RechargeCampaignRuleOversea = NewRechargeCampaignRuleOverseaClient(tx.config)
	tx.RechargeOrder = NewRechargeOrderClient(tx.config)
	tx.RenewalAgreement = NewRenewalAgreementClient(tx.config)
	tx.Survey = NewSurveyClient(tx.config)
	tx.SurveyAnswer = NewSurveyAnswerClient(tx.config)
	tx.SurveyQuestion = NewSurveyQuestionClient(tx.config)
	tx.SurveyResponse = NewSurveyResponseClient(tx.config)
	tx.Symbol = NewSymbolClient(tx.config)
	tx.TransferOrder = NewTransferOrderClient(tx.config)
	tx.TroubleDeduct = NewTroubleDeductClient(tx.config)
	tx.User = NewUserClient(tx.config)
	tx.UserDevice = NewUserDeviceClient(tx.config)
	tx.UserModel = NewUserModelClient(tx.config)
	tx.VXAccount = NewVXAccountClient(tx.config)
	tx.VXSocial = NewVXSocialClient(tx.config)
	tx.Wallet = NewWalletClient(tx.config)
	tx.WithdrawAccount = NewWithdrawAccountClient(tx.config)
	tx.WithdrawRecord = NewWithdrawRecordClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: ApiToken.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
	// completion hooks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v any) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v any) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)
