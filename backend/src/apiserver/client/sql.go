// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type postgresConfig struct {
	DSN string
}

func CreateMySQLConfig(user, password string, mysqlServiceHost string,
	mysqlServicePort string, dbName string, mysqlGroupConcatMaxLen string, mysqlExtraParams map[string]string) *mysql.Config {

	params := map[string]string{
		"charset":              "utf8",
		"parseTime":            "True",
		"loc":                  "Local",
		"group_concat_max_len": mysqlGroupConcatMaxLen,
	}

	for k, v := range mysqlExtraParams {
		params[k] = v
	}

	return &mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", mysqlServiceHost, mysqlServicePort),
		Params:               params,
		DBName:               dbName,
		AllowNativePasswords: true,
	}
}

func CreatePostgresConfig(user, password string, postgresServiceHost string, postgresServicePort string, dbName string, postgresExtraParams map[string]string) *postgresConfig {
	// There doesn't seem to be a widely-used package that constructs a DSN from a Config
	// or any other more programatic method than simply creating the string ourselves.
	// For now, we'll use a simple struct to emulate this configuration
	// Eventually, we should find a better postgres connection handlers  here.
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", postgresServiceHost, postgresServicePort, user, password, dbName)

	for k, v := range postgresExtraParams {
		dsn = fmt.Sprintf("%s %s=%s", dsn, k, v)
	}

	return &postgresConfig{
		DSN: dsn,
	}
}
