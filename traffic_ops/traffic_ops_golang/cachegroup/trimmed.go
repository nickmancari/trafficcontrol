package cachegroup

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/apache/trafficcontrol/lib/go-tc"
	"github.com/apache/trafficcontrol/traffic_ops/traffic_ops_golang/api"
)

func GetTrimmed(w http.ResponseWriter, r *http.Request) {
	inf, userErr, sysErr, errCode := api.NewInfo(r, nil, nil)
	if userErr != nil || sysErr != nil {
		api.HandleErr(w, r, inf.Tx, errCode, userErr, sysErr)
		return
	}

	defer inf.Close()
	api.RespWriter(w, r, inf.Tx)(getCachegroupsTrimmed(inf.Tx))
}

func getCachegroupsTrimmed(tx *sql.Tx) ([]tc.CachegroupTrimmedName, error) {
	names := []tc.CachegroupTrimmedName{}
	rows, err := tx.Query(`SELECT name from cachegroup`)
	if err != nil {
		return nil, errors.New("selecting cachegroup names: " + err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		name := ""
		if err := rows.Scan(&name); err != nil {
			return nil, errors.New("scanning cachegroup names: " + err.Error())
		}
		names = append(names, tc.CachegroupTrimmedName{Name: name})
	}
	return names, nil
}
