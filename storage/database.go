package storage

import (
	"fmt"
	"github.com/pkg/errors"
	jww "github.com/spf13/jwalterweatherman"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

// database interface holds function definitions for storage
type database interface {
	CheckUser(id string) (string, error)
	GetCouponCode(trigger string) (string, int, error)
	InsertCoupon(c Coupon) error
	UseCode(id, trigger string) error
}

// DatabaseImpl struct implements the database interface with an underlying DB
type DatabaseImpl struct {
	db *gorm.DB // Stored database connection
}

// MapImpl struct implements the database interface with an underlying Map
type MapImpl struct {
	coupons map[string]*Coupon
	users   map[string]*Coupon
	sync.RWMutex
}

type User struct {
	ID      string `gorm:"primary_key"`
	Trigger string
	Coupon  Coupon `gorm:"foreignkey:trigger;references:trigger"`
}

// Coupon struct defines coupons in the db
type Coupon struct {
	Trigger string `gorm:"primary_key"`
	Code    string `gorm:"NOT NULL;"`
	Uses    int
}

// newDatabase initializes the database interface
// Returns a database interface and error
func newDatabase(username, password, dbName, address,
	port string) (database, error) {

	var err error
	var db *gorm.DB
	// Connect to the database if the correct information is provided
	if address != "" && port != "" {
		// Create the database connection
		connectString := fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable",
			address, port, username, dbName)
		// Handle empty database password
		if len(password) > 0 {
			connectString += fmt.Sprintf(" password=%s", password)
		}
		db, err = gorm.Open(postgres.Open(connectString), &gorm.Config{
			Logger: logger.New(jww.TRACE, logger.Config{LogLevel: logger.Info}),
		})
	}

	// Return the map-backend interface
	// in the event there is a database error or information is not provided
	if (address == "" || port == "") || err != nil {

		if err != nil {
			jww.WARN.Printf("Unable to initialize database backend: %+v", err)
		} else {
			jww.WARN.Printf("Database backend connection information not provided")
		}

		defer jww.INFO.Println("Map backend initialized successfully!")

		mapImpl := &MapImpl{
			coupons: map[string]*Coupon{},
			users:   map[string]*Coupon{},
		}

		return database(mapImpl), nil
	}

	// Get and configure the internal database ConnPool
	sqlDb, err := db.DB()
	if err != nil {
		return database(&DatabaseImpl{}), errors.Errorf("Unable to configure database connection pool: %+v", err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDb.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the Database.
	sqlDb.SetMaxOpenConns(50)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be idle.
	sqlDb.SetConnMaxIdleTime(10 * time.Minute)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDb.SetConnMaxLifetime(12 * time.Hour)

	// Initialize the database schema
	// WARNING: Order is important. Do not change without database testing
	models := []interface{}{Coupon{}, User{}}
	for _, model := range models {
		err = db.AutoMigrate(model)
		if err != nil {
			return database(&DatabaseImpl{}), err
		}
	}

	jww.INFO.Println("Database backend initialized successfully!")
	return &DatabaseImpl{db: db}, nil
}
