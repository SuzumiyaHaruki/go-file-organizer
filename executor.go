package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func ExecutePlans(plans []MovePlan) error {
	for _, plan := range plans {
		// 递归创建目标文件路径所需的目录，权限为 755 
		if err := os.MkdirAll(filepath.Dir(plan.Target), 0o755); err != nil {
			return fmt.Errorf("create target dir for %s: %w", plan.Target, err)
		}

		target := ResolveConflict(plan.Target)
		// 将文件从源路径移动到目标路径
		if err := os.Rename(plan.Source, target); err != nil {
			return fmt.Errorf("move %s to %s: %w", plan.Source, plan.Target, err)
		}
	}

	return nil
}
