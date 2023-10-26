package database

import (
	"fmt"
	"github.com/supwr/fiap-pos-tech-docker/src/domain/entity"
	"github.com/supwr/fiap-pos-tech-docker/src/infra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewConnection(cfg infra.Config) (*gorm.DB, error) {
	var err error
	var conn *gorm.DB

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.DatabaseHost, cfg.DatabaseUsername, cfg.DatabasePassword, cfg.DatabaseDBName, cfg.DatabasePort)

	conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormlogger.Discard,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s.", cfg.DatabaseSchema),
			SingularTable: false,
		},
	})

	if err != nil {
		return nil, err
	}

	db, err := conn.DB()

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func CreateSchema(db *gorm.DB, cfg infra.Config) {
	db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", cfg.DatabaseSchema))
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&entity.Language{})
}

func InsertLanguages(db *gorm.DB, cfg infra.Config) {
	var languages = []*entity.Language{
		{ID: 1, Name: "Python"},
		{ID: 2, Name: "C"},
		{ID: 3, Name: "C++"},
		{ID: 4, Name: "Java"},
		{ID: 5, Name: "C#"},
		{ID: 6, Name: "Visual Basic .NET"},
		{ID: 7, Name: "JavaScript"},
		{ID: 8, Name: "SQL"},
		{ID: 9, Name: "Assembly language"},
		{ID: 10, Name: "PHP"},
		{ID: 11, Name: "R"},
		{ID: 12, Name: "Go"},
		{ID: 13, Name: "Classic Visual Basic"},
		{ID: 14, Name: "MATLAB"},
		{ID: 15, Name: "Swift"},
		{ID: 16, Name: "Delphi/Object Pascal"},
		{ID: 17, Name: "Ruby"},
		{ID: 18, Name: "Perl"},
		{ID: 19, Name: "Objective-C"},
		{ID: 20, Name: "Rust"},
		{ID: 21, Name: "Scratch"},
		{ID: 22, Name: "SAS"},
		{ID: 23, Name: "Kotlin"},
		{ID: 24, Name: "Julia"},
		{ID: 25, Name: "Lua"},
		{ID: 26, Name: "Fortran"},
		{ID: 27, Name: "COBOL"},
		{ID: 28, Name: "Lisp"},
		{ID: 29, Name: "(Visual) FoxPro"},
		{ID: 30, Name: "Ada"},
		{ID: 31, Name: "Dart"},
		{ID: 32, Name: "Scala"},
		{ID: 33, Name: "Prolog"},
		{ID: 34, Name: "D"},
		{ID: 35, Name: "PL/SQL"},
		{ID: 36, Name: "Bash"},
		{ID: 37, Name: "Powershell"},
		{ID: 38, Name: "Haskell"},
		{ID: 39, Name: "Logo"},
		{ID: 40, Name: "Transact SQL"},
	}

	db.CreateInBatches(languages, 40)
}
