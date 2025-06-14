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

import (
	"fmt"

	"github.com/casibase/chainserver/util"
)

// GetPatients generates mock patient data based on the provided form definition.
// It returns paginated data according to limitStr and pageStr.
//
// Parameters:
// - form: The form definition containing form items used to generate data fields.
// - limitStr: A string representing the number of records per page.
// - pageStr: A string representing the current page number.
//
// Returns:
// - A slice of map[string]string representing the patient records.
// - The total number of records (always 100 for now).
// - An error if limitStr or pageStr cannot be parsed as integers.
func GetPatients(form Form, limitStr string, pageStr string) ([]map[string]string, int, error) {
	totalCount := 100
	// Initialize a slice to store all generated patient records
	allData := make([]map[string]string, 0, totalCount)

	// Generate mock data: 100 records with values based on form.FormItems
	for i := 1; i <= totalCount; i++ {
		itemMap := make(map[string]string)
		for _, item := range form.FormItems {
			// Populate each form field with a mock value like "FieldName 1"
			itemMap[item.Name] = fmt.Sprintf("%s %d", item.Name, i)
		}
		allData = append(allData, itemMap)
	}

	// If no pagination is requested, return all data
	if limitStr == "" || pageStr == "" {
		return allData, totalCount, nil
	}

	// Parse limitStr into integer
	limit, err := util.ParseIntWithError(limitStr)
	if err != nil {
		return nil, 0, err
	}

	// Parse pageStr into integer
	page, err := util.ParseIntWithError(pageStr)
	if err != nil {
		return nil, 0, err
	}

	// Calculate the offset based on page and limit
	offset := (page - 1) * limit
	end := offset + limit

	// Clamp offset and end within bounds
	if offset > totalCount {
		offset = totalCount
	}
	if end > totalCount {
		end = totalCount
	}

	// Slice the full dataset to return only the requested page
	pagedData := allData[offset:end]
	return pagedData, totalCount, nil
}
