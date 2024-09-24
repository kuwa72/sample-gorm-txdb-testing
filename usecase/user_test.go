package usecase_test

import (
	"testing"

	"kuwa72/sample-gorm-txdb-testing/usecase"

	"github.com/DATA-DOG/go-txdb"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const dsn = "file::memory:?cache=shared"

func initDB(db *gorm.DB) {
	db.AutoMigrate(usecase.User{})
}

func NewTestDB(name string) (*gorm.DB, error) {
	txdb.Register(name, "sqlite3", dsn)
	dialector := sqlite.New(sqlite.Config{
		DriverName: name,
		DSN:        dsn,
	})

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	initDB(db)
	return db, nil
}

func TestLoginUser(t *testing.T) {
	db, _ := NewTestDB("TestLoginUser")
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	db.Create(
		&usecase.User{
			Name:     "test1",
			Email:    "test1@example.com",
			Password: "test1",
		},
	)

	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *usecase.User
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				email:    "test1@example.com",
				password: "test1",
			},
			want: &usecase.User{
				Name:     "test1",
				Email:    "test1@example.com",
				Password: "test1",
			},
			wantErr: false,
		},
		{
			name: "exists",
			args: args{
				email:    "test@example.com",
				password: "test",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			got, err := usecase.LoginUser(db, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil {
				assert.Nil(got)
				return
			}
			assert.Equal(got.Name, tt.want.Name)
			assert.Equal(got.Email, tt.want.Email)
			assert.Equal(got.Password, tt.want.Password)
		})
	}
}

func TestCreateUser(t *testing.T) {
	db, _ := NewTestDB("TestCreateUser")
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	db.Create(
		&usecase.User{
			Name:     "test2",
			Email:    "test2@example.com",
			Password: "test2",
		},
	)

	type args struct {
		name     string
		email    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *usecase.User
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				name:     "test1",
				email:    "test1@example.com",
				password: "test1",
			},
			want: &usecase.User{
				Name:     "test1",
				Email:    "test1@example.com",
				Password: "test1",
			},
			wantErr: false,
		},
		{
			name: "exists",
			args: args{
				name:     "test2",
				email:    "test2@example.com",
				password: "test2",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			got, err := usecase.CreateUser(db, tt.args.name, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil {
				assert.Nil(got)
				return
			}
			assert.Equal(got.Name, tt.want.Name)
			assert.Equal(got.Email, tt.want.Email)
			assert.Equal(got.Password, tt.want.Password)
		})
	}
}
