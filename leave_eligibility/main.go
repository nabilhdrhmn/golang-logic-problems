package main

import (
	"fmt"
	"time"
)

func main() {
	var cutiBersama, durasiCuti int
	var joinDateStr, rencanaCutiStr string

	// Input for the program with specific formatting
	fmt.Println("Input:")
	fmt.Print("• Jumlah Cuti Bersama = ")
	fmt.Scan(&cutiBersama)

	fmt.Print("• Tanggal join karyawan = ")
	fmt.Scan(&joinDateStr)

	fmt.Print("• Tanggal rencana cuti = ")
	fmt.Scan(&rencanaCutiStr)

	fmt.Print("• Durasi cuti (hari) = ")
	fmt.Scan(&durasiCuti)

	// Parsing the date strings into time.Time objects
	joinDate, _ := time.Parse("2006-01-02", joinDateStr)
	rencanaCuti, _ := time.Parse("2006-01-02", rencanaCutiStr)

	// Determine eligibility for leave
	result, reason := isEligibleForLeave(cutiBersama, joinDate, rencanaCuti, durasiCuti)

	// Output result
	fmt.Println("Output:")
	if result {
		fmt.Println("• True")
	} else {
		fmt.Println("• False")
		fmt.Println("• Alasan:", reason)
	}
}

func isEligibleForLeave(cutiBersama int, joinDate, rencanaCuti time.Time, durasiCuti int) (bool, string) {
	// Calculate 180 days after joining
	minDate := joinDate.AddDate(0, 0, 180)
	if rencanaCuti.Before(minDate) {
		return false, "Karena belum 180 hari sejak tanggal join karyawan"
	}

	// Calculate remaining leave days based on join date
	endOfYear := time.Date(rencanaCuti.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	totalDays := int(endOfYear.Sub(minDate).Hours() / 24)
	pribadiLeave := (totalDays * (14 - cutiBersama)) / 365

	// Check if the employee has enough leave days and does not exceed the max duration
	if durasiCuti > pribadiLeave {
		return false, fmt.Sprintf("Hanya boleh mengambil %d hari cuti", pribadiLeave)
	}

	if durasiCuti > 3 {
		return false, "Cuti pribadi max 3 hari berturutan"
	}

	return true, ""
}
