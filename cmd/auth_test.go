package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"my_project/internal/auth"
	"my_project/internal/user"
	"net/http"
	"net/http/httptest"
	"os"

	"strings"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//todo1: Сделать отдельный модель для инициализации тестовой бд с методами initData removeData

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Cant123 %v", err.Error())
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		panic("Cant connect to DB in auto")
	}
	log.Println("Database connected")

	return db

}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "a35643@a.ru",
		Password: "$2a$10$4lCSnfg1F18fjgmPUntOZ.DyE5jBCcicq2xCM.yC.2VcmMm5Mmc3O",
		Name:     "oleg",
	})

}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "a35643@a.ru").
		Delete(&user.User{})
}

func TestLoginSuccsess(t *testing.T) {
	ts := httptest.NewServer(app())
	db := initDb()
	initData(db)
	defer ts.Close()

	data, err := json.Marshal(auth.LoginRequest{
		Email:    "a35643@a.ru",
		Password: "123",
	})

	log.Fatalf("cant marshal body %v", err)

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK; got %d", res.StatusCode)
	}

	defer res.Body.Close()

	doby, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("cant read body %v", err)
	}

	var resp auth.LoginResponce

	err = json.Unmarshal(doby, &resp)
	if err != nil {
		t.Fatalf("cant unmarshal body %v", err)
	}

	if resp.Token == "" {
		t.Fatalf("token is empty")
	}

	log.Println(resp)

	removeData(db)
}

func TestLoginFail(t *testing.T) {

	ts := httptest.NewServer(app())
	db := initDb()
	initData(db)
	defer ts.Close()

	data, err := json.Marshal(auth.LoginRequest{
		Email:    "a35643@a.ru",
		Password: "1233", // wrong pass
	})

	if err != nil {
		t.Fatalf("cant marshal body %v", err)
	}

	req, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))

	if err != nil {
		t.Fatalf("cant post request %v", err)

	}

	if req.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected status Unauthorized; got %d", req.StatusCode)
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("cant read body %v", err)
	}

	if strings.TrimSpace(string(body)) != auth.ErrWrongCredetials {
		t.Fatal("wrong error message")
	}
	removeData(db)
}
