package logistics

import (
	"encoding/json"
	"time"

	dialectsPG "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Logistic struct {
	ID   string `gorm:"type:UUID;primaryKey;default:uuid_generate_v4()"`
	Name string
}

type ArtworkOwnershipSnapshot struct {
	LogisticID   string    `json:"-" gorm:"index:logistic_token_ownership_unique,unique"`
	TokenID      string    `json:"token_id" gorm:"index:logistic_token_ownership_unique,unique"`
	Blockchain   string    `json:"-"`
	OwnerAddress string    `json:"owner_address"`
	SnapshotAt   time.Time `json:"-" gorm:"type:TIMESTAMP WITH TIME ZONE;default:now()"`
}

type ShipmentInformation struct {
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Email            string `json:"email"`
	Address1         string `json:"address1"`
	Address2         string `json:"address2"`
	Postcode         string `json:"postcode"`
	City             string `json:"city"`
	State            string `json:"state"`
	Country          string `json:"country"`
	PhoneCountryCode string `json:"phoneCountryCode"`
	PhoneNumber      string `json:"phoneNumber"`
}

type LogisticShipmentInformation struct {
	gorm.Model
	ShipmentInformation
	LogisticID   string           `gorm:"index:logistic_shipment_info_unique,unique"`
	OwnerAddress string           `gorm:"index:logistic_shipment_info_unique,unique"`
	Data         dialectsPG.Jsonb `gorm:"type:JSONB;NOT NULL;default:'{}'::json"`
}

type LogisticStore struct {
	db *gorm.DB
}

func NewLogisticStore(dsn string) *LogisticStore {
	orm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithError(err).Panic("can not open database")
	}

	return &LogisticStore{
		db: orm,
	}
}

func (s *LogisticStore) QueryOwnedArtworkCounts(address string) (int64, error) {
	var count int64
	if err := s.db.Model(ArtworkOwnershipSnapshot{}).Where("owner_address = ?", address).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (s *LogisticStore) QueryOwnedArtworks(address string) ([]ArtworkOwnershipSnapshot, error) {
	var artworks []ArtworkOwnershipSnapshot
	if err := s.db.Where("owner_address = ?", address).Find(&artworks).Error; err != nil {
		return nil, err
	}

	return artworks, nil
}

func (s *LogisticStore) ValidateOwnedArtworks(address string, tokenIDs []string) (bool, error) {
	var ownedTokenCount int64
	if err := s.db.Model(ArtworkOwnershipSnapshot{}).Where("owner_address = ? AND token_id IN ?", address, tokenIDs).Count(&ownedTokenCount).Error; err != nil {
		return false, err
	}

	return int(ownedTokenCount) == len(tokenIDs), nil
}

func (s *LogisticStore) SaveShipmentInformation(logisticID, ownerAddress string, information ShipmentInformation, data json.RawMessage) error {
	info := LogisticShipmentInformation{
		ShipmentInformation: information,
		LogisticID:          logisticID,
		OwnerAddress:        ownerAddress,
		Data:                dialectsPG.Jsonb{RawMessage: data},
	}

	return s.db.Create(&info).Error
}

func (s *LogisticStore) GetShipmentInformation(logisticID, ownerAddress string) (LogisticShipmentInformation, error) {
	var info LogisticShipmentInformation

	err := s.db.Where(map[string]string{
		"logistic_id":   logisticID,
		"owner_address": ownerAddress,
	}).First(&info).Error

	return info, err
}
