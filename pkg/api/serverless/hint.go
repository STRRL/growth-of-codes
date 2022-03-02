package serverless

func AllAvailableRepo() ([]string, error) {
	gormDB, err := getGormDB()
	if err != nil {
		return nil, err
	}

	var result []string
	if err = gormDB.Raw("SELECT DISTINCT project FROM file_complexity_snapshots").Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func AvailableLanguageForRepo(repo string) ([]string, error) {
	gormDB, err := getGormDB()
	if err != nil {
		return nil, err
	}

	var result []string
	if err = gormDB.Raw("SELECT DISTINCT language FROM file_complexity_snapshots where project = ?", repo).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
