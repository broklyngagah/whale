package com

import (
	"testing"
	"fmt"
	"encoding/json"
)


var req = `{"device_type":"and17#$ro*id","en_data":"JpDLEtNPLiA3nGHlVjpfVPNcTtWA9HRoaTvBWZ8lkx9qgBHrx8kb5Efm5mE2owEeMNgheOBFZvI70vJ3UwwAXjEVdzsZOWKvG/CK7O7LMhlS5eFTc0CdGeYTPxwL5p9AWDKceX8b4KwKwC5ktVS5aAc8chMvpCkV+nkiRWy3EacLsmv3shekMdyf3Wdonhtpzc7zVCzLyQM5WzyZ7+0wrKE9Y3QyrBZU94TXX9wQR7Y7ddNUcBo6MNwFRgC4pCaa847JkCQaRF8gUKd2OHBkJi8qca+Cedh9GzMFnTDapPM=","en_key":"anj#*yud","opact":"article/relatedToMe","version":"1_0_0"}`

func TestBaseArg_GetVersion(t *testing.T) {
	if  getVersion("1_2_*") != "v1_2" {
		t.Error("1_2_*")
	}
	if  getVersion("1_1_*") != "v1" {
		t.Error("1_1_*")
	}
	if  getVersion("1_2_3") != "v1_2" {
		t.Error("1_2_3")
	}

	if  getVersion("2_2_*") != "" {
		t.Error("2_2_*")
	}

	if  getVersion("1_1_*") != "v1" {
		t.Error("1_1_*")
	}
}

func TestBaseArg_GetData(t *testing.T) {

	var baseArg *BaseArg

	json.Unmarshal([]byte(req), &baseArg)
	fmt.Println(baseArg)

	buf, _ := baseArg.GetData()
	fmt.Println(buf)
	var sysArg *SysArg
	err := json.Unmarshal(buf, &sysArg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sysArg.Tms)
	ti, err := sysArg.GetTime()
	fmt.Println(ti, err)
}

var resp = `MAkhQdOtfHOOHfoConpwXwcrT0aixVS9tRix26vwm0K4GK4Y9r2d1ekURjNaxEUPiCl8GErYh4WzvMFEtnuNpII9Nfq0Sdsswz8c2mXhZMTVlVpDfwRuXnqtlHzpekPPEFKK40hcqnxXll4rzfYwgIhXIregglo2MoCLkKu8ZX8xmT1Ajlfp7CQGtSglmo5QjDpuQLWTF9EcQi1zWpG2XDyBRg2tKSzgVEnmFsV3NduG3Dg5vgytt5KlnYVqo9Bbg4ePF8AhcqXMEkVATEvSIQ2iVw4epuBv`
func TestSysArg_GetTime(t *testing.T) {
	var baseArg *BaseArg

	json.Unmarshal([]byte(req), &baseArg)
	fmt.Println(baseArg)
	baseArg.EnData = resp
	buf, _ := baseArg.GetData()
	fmt.Println(string(buf))
}



