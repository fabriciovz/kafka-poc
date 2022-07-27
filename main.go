package main

import (
	"context"
	"encoding/json"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"

	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ItemFileCreatedTopic = "nifi.events.item.created"
	ItemFileUpdatedTopic = "nifi.events.item.updated"
)

var kafkaBrokers = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")

func main() {

	//item_assortment()

	//item_created()

	item_updated()
}

func item_assortment() {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:          kafkaBrokers,
		Topic:            "storesystems-topic",
		Balancer:         &kafka.LeastBytes{},
		CompressionCodec: snappy.NewCompressionCodec(),
		BatchSize:        1,
		BatchTimeout:     10 * time.Millisecond,
	})

	for i := 1; i < 10; i++ {
		item := &ItemAssortment{
			Store:                   i + 8,
			Upc:                     int64(i * 1000),
			Item:                    i + 8,
			Description:             "con el producer: " + strconv.Itoa(i),
			Department:              13,
			Cost:                    12,
			RetailWithoutTax:        11,
			RetailPrice:             100,
			Tax:                     1,
			TaxRateCode:             2,
			AddTax:                  3,
			AddTAXRateCode:          4,
			TemporaryPriceAction:    5,
			OtherIncome:             6,
			Filler:                  7,
			CommercialStructureCode: 8,
			QuantityMinimum:         9,
			QuantityMultiply:        10,
		}

		infItem, _ := json.Marshal(item)
		//fmt.Println(string(infItem))

		key := fmt.Sprintf("Key-%d", i)

		kafkaMessage := &kafka.Message{
			Key:   []byte(key),
			Value: infItem,
		}

		err := w.WriteMessages(context.Background(), *kafkaMessage)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("produced", key)
		}

	}
}

func item_created() {

	item_create := ItemFile{
		WhseMaxOrderQty:        "1",
		LastUpdatePgmID:        "2",
		StoreFormatCode:        3,
		LicenseCode:            "4",
		ItemScannableInd:       "5",
		MdseProgramID:          "6",
		DeptNbr:                7,
		VendorDeptNbr:          8,
		SigningDesc:            "9",
		SeasonCode:             "10",
		BaseRetailUomCd:        "11",
		VndrMinordUomCd:        "12",
		CompPackageQty:         "13",
		VnpkHeightQty:          "14",
		UpcDesc:                "15",
		ItemExpireDate:         "16",
		ExclusSplyDcNbr:        "17",
		ShelfLblRqmtInd:        "18",
		ItemOrdEffDate:         "19",
		CntrlSbsncInd:          "20",
		InfrmReordTypCd:        "21",
		PriceCompQty:           "22",
		AcctNbrTypeCode:        "23",
		SecurityTagInd:         "24",
		TempSensitiveInd:       "25",
		ShlfLifeDaysQty:        26,
		WhpkQty:                27,
		IdealTempLoQty:         28,
		PresnUnitQty:           29,
		ItemStatusChgDt:        "30",
		SellUomCode:            "31",
		MdsFamID:               32,
		ItemfileSourceNm:       "34",
		VndrFirstShipDt:        35,
		VndrFirstOrdDt:         36,
		PriceCompUomCd:         37,
		ShelfRotationInd:       "38",
		VnpkCubeUomCd:          "39",
		VndrMinOrdQty:          "40",
		CrushFactorCode:        "41",
		ReplGroupNbr:           "42",
		TemprUomCode:           "43",
		CustBaseRtlAmt:         "44",
		BuyingRegionCode:       45,
		RppcInd:                46,
		WhseRotationCode:       47,
		CompUnitQty:            48,
		VnpkLengthQty:          "49",
		AssortmentTypeCd:       50,
		OriginCountryCd:        "51",
		PalletHiQty:            52,
		FsaInd:                 "53",
		SeasonYear:             54,
		ItemWeightUomCd:        "55",
		WhpkCubeQty:            "56",
		FppPrepnHrQty:          57,
		CannedOrderInd:         "58",
		SendStoreDate:          "59",
		WhpkUpcNbr:             "60",
		PesticideInd:           "61",
		PerformRatingCd:        62,
		VnpkNetwgtUomCd:        "63",
		VnpkCostAmt:            "64",
		WhpkDimUomCode:         "65",
		FinelineNbr:            66,
		WhpkSellAmt:            "67",
		AcctgDeptNbr:           68,
		FppTrgtThrwyPct:        69,
		PromptPriceInd:         "70",
		LastUpdateUserid:       "71",
		WhpkCalcMthdCd:         "72",
		ItemHeightQty:          "73",
		WhseMinOrderQty:        74,
		IdealTempHiQty:         75,
		PalletUpcFmtCd:         76,
		EstOutOfStckDt:         77,
		DiscountInd:            "78",
		BaseUnitRtlAmt:         "79",
		MasterCartonInd:        "80",
		ConsumerItemNbr:        81,
		WhseAreaCode:           82,
		ProjYrSaleQty:          83,
		DcDeaReportInd:         "84",
		MdseSubcatgNbr:         85,
		MbmCode:                "86",
		ItemImportInd:          "87",
		WhpkNetwgtQty:          88,
		DietTypeCode:           89,
		SegregationCode:        90,
		VnpkWidthQty:           "91",
		Item1Desc:              "92",
		MinWhseLifeQty:         93,
		CaseUpcNbr:             "94",
		CompUnitUomCode:        95,
		WhpkWeightQty:          "96",
		Shlflbl2SizeDesc:       "97",
		ItemNbr:                98,
		UpcNbr:                 "40000738738",
		MfgrSugdRtlAmt:         100,
		SellUnitUomCode:        101,
		VnpkWeightQty:          "102",
		PalletTiQty:            103,
		AltChnnlSrcInd:         "104",
		SanitRegtnExpDt:        105,
		ReplenishUnitInd:       "106",
		ItemCubeUomCd:          "107",
		SubclassNbr:            108,
		OrdSizngFctrQty:        109,
		CreatedAt:              "110",
		PalletSizeCode:         111,
		VendorStockID:          "112",
		ItemStatusCode:         "113",
		AbcPharmaNbr:           114,
		StrShlfLfHrQty:         115,
		HealthIncentiveEligInd: 116,
		VnpkWeightUomCd:        "117",
		VendorSeqNbr:           118,
		VnpkQty:                119,
		FppVolumeUomCd:         120,
		BrandID:                121,
		BrandName:              "122",
		NonMbrUpchrgInd:        "123",
		FppMinVolumeQty:        124,
		ItemLengthQty:          "125",
		MdseCatgNbr:            126,
		VndrLeadTmQty:          127,
		AcceptTempLoQty:        128,
		CancelWhnOutInd:        "129",
		VnpkDimUomCode:         "130",
		CountryCode:            "131",
		ItemDimUomCode:         "132",
		SellUnitQty:            "133",
		WhseAlignTypeCd:        "134",
		FhsDcSlotCode:          "135",
		Item2Desc:              "136",
		ReserveMdseInd:         "137",
		ChemicalInd:            138,
		VnpkWeightFmtCd:        "139",
		ItemTypeCode:           140,
		GuarSalesInd:           "141",
		VndrIncrmOrdQty:        142,
		FppEstShrinkPct:        143,
		AerosalInd:             144,
		WhpkWeightUomCd:        "145",
		WhpkCubeUomCd:          "146",
		FppVolumeQty:           147,
		MarshalID:              148,
		Shlflbl1ColrDesc:       "149",
		PalletUpcNbr:           150,
		ItemWidthQty:           "151",
		VariableCompInd:        "152",
		LastOrderDate:          153,
		VnpkCubeQty:            "154",
		VndrLastShipDt:         155,
		CnsumableDivNbr:        156,
		VendorNbr:              157,
		ReplSubtypeCode:        158,
		QualityCntrlCd:         159,
		LegacyPluNbr:           200,
		UpdatedAt:              "201",
		RfidInd:                "202",
		WhpkUpcFormatCd:        203,
		CatchWeightInd:         "204",
		SuppDisplayInd:         "205",
		VnpkCspkCode:           "206",
		VnpkNetwgtQty:          207,
		ProductNbr:             208,
		ItemWeightQty:          "209",
		AcceptTempHiQty:        210,
		NeverOutInd:            "211",
		PluNbr:                 212,
		FppBtchShrnkQty:        213,
		ShopDesc:               214,
		CaseUpcFormatCd:        215,
		RtlNotfyStrInd:         "216",
		AccountNbr:             217,
		SellPackageQty:         218,
		FppRtrdRangeInd:        "219",
		WhpkWidthQty:           "220",
		VndrFrstAvailDt:        "221",
		SanitaryRegtnNbr:       222,
		PalletRoundPct:         223,
		ItemCubeQty:            "224",
		VariableWtInd:          "225",
		MinRcvngDaysQty:        226,
		MfgrPrePriceAmt:        227,
		BackrmScaleInd:         "228",
		WhpkHeightQty:          "229",
		ConveyableInd:          "230",
		CommodityID:            231,
		UpcFormatCode:          232,
		SellQty:                "233",
		DestinationCode:        234,
		WhpkLengthQty:          "235",
		ItemRplnshblInd:        "236",
		ReserveMdseCode:        237,
		PromoOrdBookCd:         238,
		DeptDesc:               "239",
		FinelineDesc:           "240",
		DeptSubcatgNbr:         241,
		DeptSubcatgDesc:        "242",
		DeptCategoryNbr:        243,
		DeptCategoryDesc:       "244",
		MdseSegmentNbr:         245,
		MdseSegmentDesc:        "246",
	}

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:          kafkaBrokers,
		Topic:            ItemFileCreatedTopic,
		Balancer:         &kafka.LeastBytes{},
		CompressionCodec: snappy.NewCompressionCodec(),
		BatchSize:        1,
		BatchTimeout:     10 * time.Millisecond,
	})

	for i := 1; i < 2000; i++ {

		item := &ItemFileCreated{
			UUID:       "11111",
			Timestamp:  "222222222",
			Issuer:     "333333333",
			StatusCode: "ok",
			ReplyTo:    "",
			Content:    item_create,
		}

		item.Content.ItemNbr = i + 1
		item.Content.PluNbr = 0

		infItem, _ := json.Marshal(item)
		//fmt.Println(string(infItem))

		key := fmt.Sprintf("Key-%d", i)

		kafkaMessage := &kafka.Message{
			Key:   []byte(key),
			Value: infItem,
		}

		err := w.WriteMessages(context.Background(), *kafkaMessage)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("produced", key)
		}

	}
}

func item_updated() {

	item_create := ItemFile{
		WhseMaxOrderQty:        "1",
		LastUpdatePgmID:        "2",
		StoreFormatCode:        3,
		LicenseCode:            "4",
		ItemScannableInd:       "5",
		MdseProgramID:          "6",
		DeptNbr:                7,
		VendorDeptNbr:          8,
		SigningDesc:            "9",
		SeasonCode:             "10",
		BaseRetailUomCd:        "11",
		VndrMinordUomCd:        "12",
		CompPackageQty:         "13",
		VnpkHeightQty:          "14",
		UpcDesc:                "15",
		ItemExpireDate:         "16",
		ExclusSplyDcNbr:        "17",
		ShelfLblRqmtInd:        "18",
		ItemOrdEffDate:         "19",
		CntrlSbsncInd:          "20",
		InfrmReordTypCd:        "21",
		PriceCompQty:           "22",
		AcctNbrTypeCode:        "23",
		SecurityTagInd:         "24",
		TempSensitiveInd:       "25",
		ShlfLifeDaysQty:        26,
		WhpkQty:                27,
		IdealTempLoQty:         28,
		PresnUnitQty:           29,
		ItemStatusChgDt:        "30",
		SellUomCode:            "31",
		MdsFamID:               32,
		ItemfileSourceNm:       "34",
		VndrFirstShipDt:        35,
		VndrFirstOrdDt:         36,
		PriceCompUomCd:         37,
		ShelfRotationInd:       "38",
		VnpkCubeUomCd:          "39",
		VndrMinOrdQty:          "40",
		CrushFactorCode:        "41",
		ReplGroupNbr:           "42",
		TemprUomCode:           "43",
		CustBaseRtlAmt:         "44",
		BuyingRegionCode:       45,
		RppcInd:                46,
		WhseRotationCode:       47,
		CompUnitQty:            48,
		VnpkLengthQty:          "49",
		AssortmentTypeCd:       50,
		OriginCountryCd:        "51",
		PalletHiQty:            52,
		FsaInd:                 "53",
		SeasonYear:             54,
		ItemWeightUomCd:        "55",
		WhpkCubeQty:            "56",
		FppPrepnHrQty:          57,
		CannedOrderInd:         "58",
		SendStoreDate:          "59",
		WhpkUpcNbr:             "60",
		PesticideInd:           "61",
		PerformRatingCd:        62,
		VnpkNetwgtUomCd:        "63",
		VnpkCostAmt:            "64",
		WhpkDimUomCode:         "65",
		FinelineNbr:            66,
		WhpkSellAmt:            "67",
		AcctgDeptNbr:           68,
		FppTrgtThrwyPct:        69,
		PromptPriceInd:         "70",
		LastUpdateUserid:       "71",
		WhpkCalcMthdCd:         "72",
		ItemHeightQty:          "73",
		WhseMinOrderQty:        74,
		IdealTempHiQty:         75,
		PalletUpcFmtCd:         76,
		EstOutOfStckDt:         77,
		DiscountInd:            "78",
		BaseUnitRtlAmt:         "79",
		MasterCartonInd:        "80",
		ConsumerItemNbr:        81,
		WhseAreaCode:           82,
		ProjYrSaleQty:          83,
		DcDeaReportInd:         "84",
		MdseSubcatgNbr:         85,
		MbmCode:                "86",
		ItemImportInd:          "87",
		WhpkNetwgtQty:          88,
		DietTypeCode:           89,
		SegregationCode:        90,
		VnpkWidthQty:           "91",
		Item1Desc:              "92",
		MinWhseLifeQty:         93,
		CaseUpcNbr:             "94",
		CompUnitUomCode:        95,
		WhpkWeightQty:          "96",
		Shlflbl2SizeDesc:       "97",
		ItemNbr:                98,
		UpcNbr:                 "40000738738",
		MfgrSugdRtlAmt:         100,
		SellUnitUomCode:        101,
		VnpkWeightQty:          "102",
		PalletTiQty:            103,
		AltChnnlSrcInd:         "104",
		SanitRegtnExpDt:        105,
		ReplenishUnitInd:       "106",
		ItemCubeUomCd:          "107",
		SubclassNbr:            108,
		OrdSizngFctrQty:        109,
		CreatedAt:              "110",
		PalletSizeCode:         111,
		VendorStockID:          "112",
		ItemStatusCode:         "113",
		AbcPharmaNbr:           114,
		StrShlfLfHrQty:         115,
		HealthIncentiveEligInd: 116,
		VnpkWeightUomCd:        "117",
		VendorSeqNbr:           118,
		VnpkQty:                119,
		FppVolumeUomCd:         120,
		BrandID:                121,
		BrandName:              "122",
		NonMbrUpchrgInd:        "123",
		FppMinVolumeQty:        124,
		ItemLengthQty:          "125",
		MdseCatgNbr:            126,
		VndrLeadTmQty:          127,
		AcceptTempLoQty:        128,
		CancelWhnOutInd:        "129",
		VnpkDimUomCode:         "130",
		CountryCode:            "131",
		ItemDimUomCode:         "132",
		SellUnitQty:            "133",
		WhseAlignTypeCd:        "134",
		FhsDcSlotCode:          "135",
		Item2Desc:              "136",
		ReserveMdseInd:         "137",
		ChemicalInd:            138,
		VnpkWeightFmtCd:        "139",
		ItemTypeCode:           140,
		GuarSalesInd:           "141",
		VndrIncrmOrdQty:        142,
		FppEstShrinkPct:        143,
		AerosalInd:             144,
		WhpkWeightUomCd:        "145",
		WhpkCubeUomCd:          "146",
		FppVolumeQty:           147,
		MarshalID:              148,
		Shlflbl1ColrDesc:       "149",
		PalletUpcNbr:           150,
		ItemWidthQty:           "151",
		VariableCompInd:        "152",
		LastOrderDate:          153,
		VnpkCubeQty:            "154",
		VndrLastShipDt:         155,
		CnsumableDivNbr:        156,
		VendorNbr:              157,
		ReplSubtypeCode:        158,
		QualityCntrlCd:         159,
		LegacyPluNbr:           200,
		UpdatedAt:              "201",
		RfidInd:                "202",
		WhpkUpcFormatCd:        203,
		CatchWeightInd:         "204",
		SuppDisplayInd:         "205",
		VnpkCspkCode:           "206",
		VnpkNetwgtQty:          207,
		ProductNbr:             208,
		ItemWeightQty:          "209",
		AcceptTempHiQty:        210,
		NeverOutInd:            "211",
		PluNbr:                 212,
		FppBtchShrnkQty:        213,
		ShopDesc:               214,
		CaseUpcFormatCd:        215,
		RtlNotfyStrInd:         "216",
		AccountNbr:             217,
		SellPackageQty:         218,
		FppRtrdRangeInd:        "219",
		WhpkWidthQty:           "220",
		VndrFrstAvailDt:        "221",
		SanitaryRegtnNbr:       222,
		PalletRoundPct:         223,
		ItemCubeQty:            "224",
		VariableWtInd:          "225",
		MinRcvngDaysQty:        226,
		MfgrPrePriceAmt:        227,
		BackrmScaleInd:         "228",
		WhpkHeightQty:          "229",
		ConveyableInd:          "230",
		CommodityID:            231,
		UpcFormatCode:          232,
		SellQty:                "233",
		DestinationCode:        234,
		WhpkLengthQty:          "235",
		ItemRplnshblInd:        "236",
		ReserveMdseCode:        237,
		PromoOrdBookCd:         238,
		DeptDesc:               "239",
		FinelineDesc:           "240",
		DeptSubcatgNbr:         241,
		DeptSubcatgDesc:        "242",
		DeptCategoryNbr:        243,
		DeptCategoryDesc:       "244",
		MdseSegmentNbr:         245,
		MdseSegmentDesc:        "246",
	}

	content := Content{
		Actual:   item_create,
		Previous: item_create,
	}

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:          kafkaBrokers,
		Topic:            ItemFileUpdatedTopic,
		Balancer:         &kafka.LeastBytes{},
		CompressionCodec: snappy.NewCompressionCodec(),
		BatchSize:        1,
		BatchTimeout:     10 * time.Millisecond,
	})

	for i := 1; i < 2000; i++ {

		item := &ItemFileUpdated{
			UUID:       "11111",
			Timestamp:  "222222222",
			Issuer:     "333333333",
			StatusCode: "ok",
			ReplyTo:    "",
			Content:    content,
		}

		fmt.Println("upc number: ", item.Content.Actual.UpcNbr)
		item.Content.Actual.ItemNbr = i + 1
		item.Content.Previous.ItemNbr = i + 1
		item.Content.Actual.PluNbr = 5

		infItem, _ := json.Marshal(item)
		//fmt.Println(string(infItem))

		key := fmt.Sprintf("Key-%d", i)

		kafkaMessage := &kafka.Message{
			Key:   []byte(key),
			Value: infItem,
		}

		err := w.WriteMessages(context.Background(), *kafkaMessage)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("produced", key)
		}

	}
}

type ItemAssortment struct {
	Store                   int    `json:"store"`
	Upc                     int64  `json:"upc"`
	Item                    int    `json:"item"`
	Description             string `json:"description"`
	Department              int    `json:"department"`
	Cost                    int    `json:"cost"`
	RetailWithoutTax        int    `json:"retailWithoutTax"`
	RetailPrice             int    `json:"retailPrice"`
	Tax                     int    `json:"tax"`
	TaxRateCode             int    `json:"taxRateCode"`
	AddTax                  int    `json:"addTax"`
	AddTAXRateCode          int    `json:"addTAXRateCode"`
	TemporaryPriceAction    int    `json:"temporaryPriceAction"`
	OtherIncome             int    `json:"otherIncome"`
	Filler                  int    `json:"filler"`
	CommercialStructureCode int64  `json:"commercialStructureCode"`
	QuantityMinimum         int    `json:"quantityMinimum"`
	QuantityMultiply        int    `json:"quantityMultiply"`
}

//items

type ItemFileUpdated struct {
	UUID       string  `json:"uuid"`
	Timestamp  string  `json:"timestamp"`
	Issuer     string  `json:"issuer"`
	StatusCode string  `json:"statusCode"`
	ReplyTo    string  `json:"replyTo"`
	Content    Content `json:"content"`
}

type Content struct {
	Actual   ItemFile `json:"actual"`
	Previous ItemFile `json:"previous"`
}

type ItemFileCreated struct {
	UUID       string   `json:"uuid"`
	Timestamp  string   `json:"timestamp"`
	Issuer     string   `json:"issuer"`
	StatusCode string   `json:"statusCode"`
	ReplyTo    string   `json:"replyTo"`
	Content    ItemFile `json:"content"`
}

type ItemFile struct {
	WhseMaxOrderQty        interface{} `json:"whseMaxOrderQty"`
	LastUpdatePgmID        string      `json:"lastUpdatePgmId"`
	StoreFormatCode        int         `json:"storeFormatCode"`
	LicenseCode            interface{} `json:"licenseCode"`
	ItemScannableInd       string      `json:"itemScannableInd"`
	MdseProgramID          interface{} `json:"mdseProgramId"`
	DeptNbr                int         `json:"deptNbr"`
	VendorDeptNbr          int         `json:"vendorDeptNbr"`
	SigningDesc            string      `json:"signingDesc"`
	SeasonCode             interface{} `json:"seasonCode"`
	BaseRetailUomCd        string      `json:"baseRetailUomCd"`
	VndrMinordUomCd        interface{} `json:"vndrMinordUomCd"`
	CompPackageQty         interface{} `json:"compPackageQty"`
	VnpkHeightQty          string      `json:"vnpkHeightQty"`
	UpcDesc                string      `json:"upcDesc"`
	ItemExpireDate         string      `json:"itemExpireDate"`
	ExclusSplyDcNbr        interface{} `json:"exclusSplyDcNbr"`
	ShelfLblRqmtInd        string      `json:"shelfLblRqmtInd"`
	ItemOrdEffDate         string      `json:"itemOrdEffDate"`
	CntrlSbsncInd          string      `json:"cntrlSbsncInd"`
	InfrmReordTypCd        interface{} `json:"infrmReordTypCd"`
	PriceCompQty           interface{} `json:"priceCompQty"`
	AcctNbrTypeCode        interface{} `json:"acctNbrTypeCode"`
	SecurityTagInd         string      `json:"securityTagInd"`
	TempSensitiveInd       string      `json:"tempSensitiveInd"`
	ShlfLifeDaysQty        int         `json:"shlfLifeDaysQty"`
	WhpkQty                int         `json:"whpkQty"`
	IdealTempLoQty         int         `json:"idealTempLoQty"`
	PresnUnitQty           interface{} `json:"presnUnitQty"`
	ItemStatusChgDt        string      `json:"itemStatusChgDt"`
	SellUomCode            string      `json:"sellUomCode"`
	MdsFamID               int         `json:"mdsFamId"`
	ItemfileSourceNm       string      `json:"itemfileSourceNm"`
	VndrFirstShipDt        interface{} `json:"vndrFirstShipDt"`
	VndrFirstOrdDt         interface{} `json:"vndrFirstOrdDt"`
	PriceCompUomCd         interface{} `json:"priceCompUomCd"`
	ShelfRotationInd       string      `json:"shelfRotationInd"`
	VnpkCubeUomCd          string      `json:"vnpkCubeUomCd"`
	VndrMinOrdQty          interface{} `json:"vndrMinOrdQty"`
	CrushFactorCode        interface{} `json:"crushFactorCode"`
	ReplGroupNbr           string      `json:"replGroupNbr"`
	TemprUomCode           interface{} `json:"temprUomCode"`
	CustBaseRtlAmt         string      `json:"custBaseRtlAmt"`
	BuyingRegionCode       int         `json:"buyingRegionCode"`
	RppcInd                interface{} `json:"rppcInd"`
	WhseRotationCode       interface{} `json:"whseRotationCode"`
	CompUnitQty            interface{} `json:"compUnitQty"`
	VnpkLengthQty          string      `json:"vnpkLengthQty"`
	AssortmentTypeCd       interface{} `json:"assortmentTypeCd"`
	OriginCountryCd        string      `json:"originCountryCd"`
	PalletHiQty            int         `json:"palletHiQty"`
	FsaInd                 string      `json:"fsaInd"`
	SeasonYear             interface{} `json:"seasonYear"`
	ItemWeightUomCd        string      `json:"itemWeightUomCd"`
	WhpkCubeQty            string      `json:"whpkCubeQty"`
	FppPrepnHrQty          interface{} `json:"fppPrepnHrQty"`
	CannedOrderInd         string      `json:"cannedOrderInd"`
	SendStoreDate          string      `json:"sendStoreDate"`
	WhpkUpcNbr             string      `json:"whpkUpcNbr"`
	PesticideInd           interface{} `json:"pesticideInd"`
	PerformRatingCd        int         `json:"performRatingCd"`
	VnpkNetwgtUomCd        string      `json:"vnpkNetwgtUomCd"`
	VnpkCostAmt            string      `json:"vnpkCostAmt"`
	WhpkDimUomCode         string      `json:"whpkDimUomCode"`
	FinelineNbr            int         `json:"finelineNbr"`
	WhpkSellAmt            string      `json:"whpkSellAmt"`
	AcctgDeptNbr           int         `json:"acctgDeptNbr"`
	FppTrgtThrwyPct        interface{} `json:"fppTrgtThrwyPct"`
	PromptPriceInd         string      `json:"promptPriceInd"`
	LastUpdateUserid       string      `json:"lastUpdateUserid"`
	WhpkCalcMthdCd         string      `json:"whpkCalcMthdCd"`
	ItemHeightQty          string      `json:"itemHeightQty"`
	WhseMinOrderQty        int         `json:"whseMinOrderQty"`
	IdealTempHiQty         int         `json:"idealTempHiQty"`
	PalletUpcFmtCd         interface{} `json:"palletUpcFmtCd"`
	EstOutOfStckDt         interface{} `json:"estOutOfStckDt"`
	DiscountInd            string      `json:"discountInd"`
	BaseUnitRtlAmt         string      `json:"baseUnitRtlAmt"`
	MasterCartonInd        string      `json:"masterCartonInd"`
	ConsumerItemNbr        int         `json:"consumerItemNbr"`
	WhseAreaCode           int         `json:"whseAreaCode"`
	ProjYrSaleQty          int         `json:"projYrSaleQty"`
	DcDeaReportInd         string      `json:"dcDeaReportInd"`
	MdseSubcatgNbr         interface{} `json:"mdseSubcatgNbr"`
	MbmCode                string      `json:"mbmCode"`
	ItemImportInd          string      `json:"itemImportInd"`
	WhpkNetwgtQty          interface{} `json:"whpkNetwgtQty"`
	DietTypeCode           interface{} `json:"dietTypeCode"`
	SegregationCode        int         `json:"segregationCode"`
	VnpkWidthQty           string      `json:"vnpkWidthQty"`
	Item1Desc              string      `json:"item1Desc"`
	MinWhseLifeQty         int         `json:"minWhseLifeQty"`
	CaseUpcNbr             string      `json:"caseUpcNbr"`
	CompUnitUomCode        interface{} `json:"compUnitUomCode"`
	WhpkWeightQty          string      `json:"whpkWeightQty"`
	Shlflbl2SizeDesc       string      `json:"shlflbl2SizeDesc"`
	ItemNbr                int         `json:"itemNbr"`
	UpcNbr                 string      `json:"upcNbr"`
	MfgrSugdRtlAmt         interface{} `json:"mfgrSugdRtlAmt"`
	SellUnitUomCode        interface{} `json:"sellUnitUomCode"`
	VnpkWeightQty          string      `json:"vnpkWeightQty"`
	PalletTiQty            int         `json:"palletTiQty"`
	AltChnnlSrcInd         string      `json:"altChnnlSrcInd"`
	SanitRegtnExpDt        interface{} `json:"sanitRegtnExpDt"`
	ReplenishUnitInd       string      `json:"replenishUnitInd"`
	ItemCubeUomCd          string      `json:"itemCubeUomCd"`
	SubclassNbr            int         `json:"subclassNbr"`
	OrdSizngFctrQty        interface{} `json:"ordSizngFctrQty"`
	CreatedAt              string      `json:"createdAt"`
	PalletSizeCode         interface{} `json:"palletSizeCode"`
	VendorStockID          string      `json:"vendorStockId"`
	ItemStatusCode         string      `json:"itemStatusCode"`
	AbcPharmaNbr           interface{} `json:"abcPharmaNbr"`
	StrShlfLfHrQty         interface{} `json:"strShlfLfHrQty"`
	HealthIncentiveEligInd interface{} `json:"healthIncentiveEligInd"`
	VnpkWeightUomCd        string      `json:"vnpkWeightUomCd"`
	VendorSeqNbr           int         `json:"vendorSeqNbr"`
	VnpkQty                int         `json:"vnpkQty"`
	FppVolumeUomCd         interface{} `json:"fppVolumeUomCd"`
	BrandID                int         `json:"brandId"`
	BrandName              string      `json:"brandName"`
	NonMbrUpchrgInd        string      `json:"nonMbrUpchrgInd"`
	FppMinVolumeQty        interface{} `json:"fppMinVolumeQty"`
	ItemLengthQty          string      `json:"itemLengthQty"`
	MdseCatgNbr            interface{} `json:"mdseCatgNbr"`
	VndrLeadTmQty          int         `json:"vndrLeadTmQty"`
	AcceptTempLoQty        interface{} `json:"acceptTempLoQty"`
	CancelWhnOutInd        string      `json:"cancelWhnOutInd"`
	VnpkDimUomCode         string      `json:"vnpkDimUomCode"`
	CountryCode            string      `json:"countryCode"`
	ItemDimUomCode         string      `json:"itemDimUomCode"`
	SellUnitQty            interface{} `json:"sellUnitQty"`
	WhseAlignTypeCd        string      `json:"whseAlignTypeCd"`
	FhsDcSlotCode          interface{} `json:"fhsDcSlotCode"`
	Item2Desc              string      `json:"item2Desc"`
	ReserveMdseInd         string      `json:"reserveMdseInd"`
	ChemicalInd            interface{} `json:"chemicalInd"`
	VnpkWeightFmtCd        string      `json:"vnpkWeightFmtCd"`
	ItemTypeCode           int         `json:"itemTypeCode"`
	GuarSalesInd           string      `json:"guarSalesInd"`
	VndrIncrmOrdQty        interface{} `json:"vndrIncrmOrdQty"`
	FppEstShrinkPct        interface{} `json:"fppEstShrinkPct"`
	AerosalInd             interface{} `json:"aerosalInd"`
	WhpkWeightUomCd        string      `json:"whpkWeightUomCd"`
	WhpkCubeUomCd          string      `json:"whpkCubeUomCd"`
	FppVolumeQty           interface{} `json:"fppVolumeQty"`
	MarshalID              interface{} `json:"marshalId"`
	Shlflbl1ColrDesc       string      `json:"shlflbl1ColrDesc"`
	PalletUpcNbr           interface{} `json:"palletUpcNbr"`
	ItemWidthQty           string      `json:"itemWidthQty"`
	VariableCompInd        string      `json:"variableCompInd"`
	LastOrderDate          interface{} `json:"lastOrderDate"`
	VnpkCubeQty            string      `json:"vnpkCubeQty"`
	VndrLastShipDt         interface{} `json:"vndrLastShipDt"`
	CnsumableDivNbr        interface{} `json:"cnsumableDivNbr"`
	VendorNbr              int         `json:"vendorNbr"`
	ReplSubtypeCode        int         `json:"replSubtypeCode"`
	QualityCntrlCd         interface{} `json:"qualityCntrlCd"`
	LegacyPluNbr           int         `json:"legacyPluNbr"`
	UpdatedAt              string      `json:"updatedAt"`
	RfidInd                string      `json:"rfidInd"`
	WhpkUpcFormatCd        interface{} `json:"whpkUpcFormatCd"`
	CatchWeightInd         string      `json:"catchWeightInd"`
	SuppDisplayInd         string      `json:"suppDisplayInd"`
	VnpkCspkCode           string      `json:"vnpkCspkCode"`
	VnpkNetwgtQty          interface{} `json:"vnpkNetwgtQty"`
	ProductNbr             int         `json:"productNbr"`
	ItemWeightQty          string      `json:"itemWeightQty"`
	AcceptTempHiQty        interface{} `json:"acceptTempHiQty"`
	NeverOutInd            string      `json:"neverOutInd"`
	PluNbr                 int         `json:"pluNbr"`
	FppBtchShrnkQty        interface{} `json:"fppBtchShrnkQty"`
	ShopDesc               interface{} `json:"shopDesc"`
	CaseUpcFormatCd        int         `json:"caseUpcFormatCd"`
	RtlNotfyStrInd         string      `json:"rtlNotfyStrInd"`
	AccountNbr             interface{} `json:"accountNbr"`
	SellPackageQty         interface{} `json:"sellPackageQty"`
	FppRtrdRangeInd        string      `json:"fppRtrdRangeInd"`
	WhpkWidthQty           string      `json:"whpkWidthQty"`
	VndrFrstAvailDt        string      `json:"vndrFrstAvailDt"`
	SanitaryRegtnNbr       interface{} `json:"sanitaryRegtnNbr"`
	PalletRoundPct         interface{} `json:"palletRoundPct"`
	ItemCubeQty            string      `json:"itemCubeQty"`
	VariableWtInd          string      `json:"variableWtInd"`
	MinRcvngDaysQty        int         `json:"minRcvngDaysQty"`
	MfgrPrePriceAmt        interface{} `json:"mfgrPrePriceAmt"`
	BackrmScaleInd         string      `json:"backrmScaleInd"`
	WhpkHeightQty          string      `json:"whpkHeightQty"`
	ConveyableInd          string      `json:"conveyableInd"`
	CommodityID            interface{} `json:"commodityId"`
	UpcFormatCode          int         `json:"upcFormatCode"`
	SellQty                string      `json:"sellQty"`
	DestinationCode        int         `json:"destinationCode"`
	WhpkLengthQty          string      `json:"whpkLengthQty"`
	ItemRplnshblInd        string      `json:"itemRplnshblInd"`
	ReserveMdseCode        int         `json:"reserveMdseCode"`
	PromoOrdBookCd         interface{} `json:"promoOrdBookCd"`
	DeptDesc               string      `json:"deptDesc"`
	FinelineDesc           string      `json:"finelineDesc"`
	DeptSubcatgNbr         int         `json:"deptSubcatgNbr"`
	DeptSubcatgDesc        string      `json:"deptSubcatgDesc"`
	DeptCategoryNbr        int         `json:"deptCategoryNbr"`
	DeptCategoryDesc       string      `json:"deptCategoryDesc"`
	MdseSegmentNbr         int         `json:"segmentNbr"`
	MdseSegmentDesc        string      `json:"segmentDesc"`
}
