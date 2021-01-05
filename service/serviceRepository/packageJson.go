package serviceRepository

import (
	"MareWood/config"
	"MareWood/helper"
	"MareWood/models"
	"encoding/json"
	"errors"
	"io/ioutil"
)

func ParsePackageJson(repositoryId string) (models.PackageJson , error)  {


	var packageJson models.PackageJson

	packageJsonFile := config.Cfg.RepositoryDir + "/" + repositoryId + "/package.json"

	if !helper.FileIsExisted(packageJsonFile){
		return packageJson, errors.New("Json file not found=>" + packageJsonFile)
	}

	file, err := ioutil.ReadFile(packageJsonFile)

	if err != nil {
		return packageJson, err
	}

	err = json.Unmarshal(file, &packageJson)

	if err != nil {
		return packageJson, err
	}

	return packageJson ,nil

}