package database

import(
	"fmt"

	"go-db-learn/domain/entity"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type Repositories struct{
	db	*gorm.DB
	Name string
}

func NewRepository(DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error){

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName, DbPort )
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		return nil, err
	}

	return &Repositories{
		db: db,
		Name: "Masa",
	}, nil
}

func (s *Repositories) Close() error{
	sqlDb, err := s.db.DB()
	if err != nil{
		return sqlDb.Close()
	}

	return err
}

func (s *Repositories) Automigrate() error{
	return s.db.AutoMigrate(&entity.Country{}, &entity.Currency{})
}