package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"log"
	"os"
	"path/filepath"
)

// 配置信息
var (
	farmerContract   *gateway.Contract
	driverContract   *gateway.Contract
	materialContract *gateway.Contract
	sellContract     *gateway.Contract
	sellerContract   *gateway.Contract
	capitalContract  *gateway.Contract
	usersContract    *gateway.Contract
)

// 初始化
func init() {
	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "false")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environment variable: %v", err)
	}

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}

	ccpPath := filepath.Join(
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	farmerContract = network.GetContract("farmer")
	driverContract = network.GetContract("driver")
	materialContract = network.GetContract("material")
	sellContract = network.GetContract("sell")
	capitalContract = network.GetContract("capital")
	sellerContract = network.GetContract("seller")
	usersContract = network.GetContract("users")
}

func populateWallet(wallet *gateway.Wallet) error {
	log.Println("============ Populating wallet ============")
	credPath := filepath.Join(
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "User1@org1.example.com-cert.pem")
	// read the certificate pem
	cert, err := os.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := os.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := os.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	return wallet.Put("appUser", identity)
}

//farmer链码

func CreateCrops(CropsId, CropsName, Address, RegisterTime, Year, FarmerName, FarmerID, FarmerTel, FertilizerName, PlatMode, BaggingStatus, GrowSeedlingsCycle, IrrigationCycle, ApplyFertilizerCycle, WeedCycle, Remarks string) ([]byte, error) {
	return farmerContract.SubmitTransaction("CreateCrops", CropsId, CropsName, Address, RegisterTime, Year, FarmerName, FarmerID, FarmerTel, FertilizerName, PlatMode, BaggingStatus, GrowSeedlingsCycle, IrrigationCycle, ApplyFertilizerCycle, WeedCycle, Remarks)
}
func RecordCropsGrow(CropsGrowId, CropsBakId, RecordTime, CropsGrowPhotoUrl, Temperature, GrowStatus, WaterContent, IlluminationStatus, Remarks string) ([]byte, error) {
	return farmerContract.SubmitTransaction("RecordCropsGrow", CropsGrowId, CropsBakId, RecordTime, CropsGrowPhotoUrl, Temperature, GrowStatus, WaterContent, IlluminationStatus, Remarks)
}
func QueryCropsById(CropsID string) ([]byte, error) {
	return farmerContract.EvaluateTransaction("QueryCropsById", CropsID)
}
func QueryCropsProcessByCropsId(CropsID string) ([]byte, error) {
	return farmerContract.EvaluateTransaction("QueryCropsProcessByCropsId", CropsID)
}

//driver链码

func CreateTransport(TransportId, DriverId, DriverName, DriverTel, DriverDept, CropsId, TransportToChainTime, TransportToAddress, Remarks string) ([]byte, error) {
	return driverContract.SubmitTransaction("CreateTransport", TransportId, DriverId, DriverName, DriverTel, DriverDept, CropsId, TransportToChainTime, TransportToAddress, Remarks)
}
func QueryTransportById(transportID string) ([]byte, error) {
	return driverContract.EvaluateTransaction("QueryTransportById", transportID)
}
func QueryTransportByCropsId(cropsID string) ([]byte, error) {
	return driverContract.EvaluateTransaction("QueryTransportById", cropsID)
}

//material链码

func CreateMachining(MachiningId, Leader, CropsId, LeaderTel, FactoryName, TestingResult, InFactoryTime, OutFactoryTime, TestingPhotoUrl, Remarks string) ([]byte, error) {
	return materialContract.SubmitTransaction("CreateMachining", MachiningId, Leader, CropsId, LeaderTel, FactoryName, TestingResult, InFactoryTime, OutFactoryTime, TestingPhotoUrl, Remarks)
}
func QueryMachiningById(MachiningId string) ([]byte, error) {
	return materialContract.EvaluateTransaction("QueryMachiningById", MachiningId)
}
func QueryMachiningByCropsId(CropsID string) ([]byte, error) {
	return materialContract.EvaluateTransaction("QueryMachiningById", CropsID)
}

//sell链码

func CreateSelling(SellID, CropsID, SellerID, BuyerID, Price string) ([]byte, error) {
	return sellContract.SubmitTransaction("CreateSelling", SellID, CropsID, SellerID, BuyerID, Price)
}
func QueryBySellID(SellID string) ([]byte, error) {
	return sellContract.EvaluateTransaction("QueryBySellID", SellID)
}
func QueryBySellerID(SellerID string) ([]byte, error) {
	return sellContract.EvaluateTransaction("QueryBySellerID", SellerID)
}
func QueryByBuyerID(BuyerID string) ([]byte, error) {
	return sellContract.EvaluateTransaction("QueryByBuyerID", BuyerID)
}

//seller链码

func CreateCommodity(CropsID, SellerID, Price, State string) ([]byte, error) {
	return sellerContract.SubmitTransaction("CreateCommodity", CropsID, SellerID, Price, State)
}
func ChangeCommodity(CropsID, SellerID, Price, State string) ([]byte, error) {
	return sellerContract.SubmitTransaction("ChangeCommodity", CropsID, SellerID, Price, State)
}
func DelCommodity(CropsID string) ([]byte, error) {
	return sellerContract.EvaluateTransaction("DelCommodity", CropsID)
}
func QueryAllBySellerID(SellerID string) ([]byte, error) {
	return sellerContract.EvaluateTransaction("QueryAllBySellerID", SellerID)
}

//capital链码

func CreateBalance(UserID, Balance string) ([]byte, error) {
	return capitalContract.SubmitTransaction("CreateBalance", UserID, Balance)
}
func QueryBalance(UserID string) ([]byte, error) {
	return capitalContract.EvaluateTransaction("QueryBalance", UserID)
}
func ChangeBalance(UserID, Balance string) ([]byte, error) {
	return capitalContract.EvaluateTransaction("ChangeBalance", UserID, Balance)
}

//users链码

func CreateAccount(UserID, Passwd string) ([]byte, error) {
	return usersContract.SubmitTransaction("CreateAccount", UserID, Passwd)
}
func ChangePass(UserID, Passwd string) ([]byte, error) {
	return usersContract.SubmitTransaction("ChangePass", UserID, Passwd)
}
