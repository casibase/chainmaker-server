// Copyright 2025 The Casibase Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package object

import "github.com/casibase/chainserver/util"

// GetHospitals returns 24 hardcoded hospital records based on the expected form structure.
// Supports pagination via limitStr and pageStr.
func GetHospitals(form Form, limitStr string, pageStr string) ([]map[string]string, int, error) {
	allData := []map[string]string{
		{"name": "Central Hospital", "id": "HOSP001", "city": "Beijing", "location": "Chaoyang", "size": "Large"},
		{"name": "Westside Medical", "id": "HOSP002", "city": "Shanghai", "location": "Pudong", "size": "Medium"},
		{"name": "East Health Center", "id": "HOSP003", "city": "Guangzhou", "location": "Tianhe", "size": "Medium"},
		{"name": "North Clinic", "id": "HOSP004", "city": "Shenyang", "location": "Huanggu", "size": "Small"},
		{"name": "South General", "id": "HOSP005", "city": "Chengdu", "location": "Jinjiang", "size": "Large"},
		{"name": "Hope Hospital", "id": "HOSP006", "city": "Wuhan", "location": "Wuchang", "size": "Medium"},
		{"name": "BrightCare Center", "id": "HOSP007", "city": "Hangzhou", "location": "Xihu", "size": "Small"},
		{"name": "Peace Medical", "id": "HOSP008", "city": "Nanjing", "location": "Jianye", "size": "Medium"},
		{"name": "Unity Health", "id": "HOSP009", "city": "Shenzhen", "location": "Nanshan", "size": "Large"},
		{"name": "Harmony Clinic", "id": "HOSP010", "city": "Xi'an", "location": "Yanta", "size": "Small"},
		{"name": "Grace Hospital", "id": "HOSP011", "city": "Suzhou", "location": "Gusu", "size": "Medium"},
		{"name": "Sunrise Medical", "id": "HOSP012", "city": "Tianjin", "location": "Heping", "size": "Large"},
		{"name": "HealthFirst", "id": "HOSP013", "city": "Qingdao", "location": "Shinan", "size": "Medium"},
		{"name": "River Valley Clinic", "id": "HOSP014", "city": "Ningbo", "location": "Haishu", "size": "Small"},
		{"name": "Evergreen Hospital", "id": "HOSP015", "city": "Fuzhou", "location": "Gulou", "size": "Large"},
		{"name": "CityMed", "id": "HOSP016", "city": "Changsha", "location": "Yuelu", "size": "Medium"},
		{"name": "Pinewood Health", "id": "HOSP017", "city": "Jinan", "location": "Lixia", "size": "Small"},
		{"name": "Golden Gate Medical", "id": "HOSP018", "city": "Xiamen", "location": "Siming", "size": "Medium"},
		{"name": "Oceanview Clinic", "id": "HOSP019", "city": "Haikou", "location": "Longhua", "size": "Small"},
		{"name": "Mountain Ridge Hospital", "id": "HOSP020", "city": "Kunming", "location": "Panlong", "size": "Large"},
		{"name": "Skyline Health", "id": "HOSP021", "city": "Zhengzhou", "location": "Erqi", "size": "Medium"},
		{"name": "Lakeside Medical", "id": "HOSP022", "city": "Nanchang", "location": "Donghu", "size": "Small"},
		{"name": "Greenfield Clinic", "id": "HOSP023", "city": "Hefei", "location": "Luyang", "size": "Medium"},
		{"name": "Wellbeing Center", "id": "HOSP024", "city": "Harbin", "location": "Daoli", "size": "Large"},
	}

	totalCount := len(allData)

	// If no pagination parameters are given, return full dataset
	if limitStr == "" || pageStr == "" {
		return allData, totalCount, nil
	}

	// Parse limit and page
	limit, err := util.ParseIntWithError(limitStr)
	if err != nil {
		return nil, 0, err
	}
	page, err := util.ParseIntWithError(pageStr)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	end := offset + limit
	if offset > totalCount {
		offset = totalCount
	}
	if end > totalCount {
		end = totalCount
	}

	pagedData := allData[offset:end]
	return pagedData, totalCount, nil
}
