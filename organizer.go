package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type MovePlan struct{
	Source string
	Target string
	Category string
}

func BuildPlans(dir string) ([]MovePlan, error){
	// 获取 dir 路径中的文件
	entities, err := os.ReadDir(dir)
	if err != nil{
		return nil, fmt.Errorf("read dir: %w", err)
	}

	var plans []MovePlan
	for _, entity := range entities{
		if entity.IsDir(){
			continue
		}

		name := entity.Name()
		ext := strings.ToLower(filepath.Ext(name))
		category := CategoryByExt(ext)

		plans = append(plans, MovePlan{
			Source:   filepath.Join(dir, name),
			Target:   filepath.Join(dir, category, name),
			Category: category,
		})
	}

	return plans, nil
}
