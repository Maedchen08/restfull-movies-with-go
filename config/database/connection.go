package database
import(
	"log"
	"time"
	"homework-rakamin-go-sql/models"
	"homework-rakamin-go-sql/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	maxIdleConn := config.GetInt("DB_MAX_IDLE_CONNECTIONS")
	maxConn := config.GetInt("DB_MAX_CONNECTIONS")
	maxLifetimeConn := config.GetInt("DB_MAX_LIFETIME_CONNECTIONS")

	dsn := config.GetString("DB_SERVER_URL")
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	InitCreateTable(db)
	log.Println("database connect success")
	return db
	
}

func InitCreateTable(db *gorm.DB){
	
	db.AutoMigrate(&models.Movies{})
}

