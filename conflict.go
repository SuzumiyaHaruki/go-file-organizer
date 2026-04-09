package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func ResolveConflict(target string) string {
	if _, err := os.Stat(target); os.IsNotExist(err) {
		return target
	}

	// 获取 target的目录部分
	dir := filepath.Dir(target)
	// 提取文件扩展名
	ext := filepath.Ext(target)
	// 获取完整文件名（不含扩展名）
	name := target[:len(target)-len(ext)]
	// 提取基本文件名（不含扩展名）
	base := filepath.Base(name)

	// 修改文件名为 base(i).ext 形式
	for i := 1; ; i++ {
		candidate := filepath.Join(dir, fmt.Sprintf("%s (%d)%s", base, i, ext))
		if _, err := os.Stat(candidate); os.IsNotExist(err) {
			return candidate
		}
	}
}
