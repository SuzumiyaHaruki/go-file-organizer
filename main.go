package main

import (
	"flag"
	"fmt"
	"log"
)

func main(){
	// 定义 --dir 参数，用于指定要整理的目录，默认值为 .（当前目录）
	dir := flag.String("dir", ".", "directory to organize")
	// 定义 --apply 参数，控制是否实际执行更改，默认值为 false
	apply := flag.Bool("apply", false, "apply changes")
	flag.Parse()

	plans, err := BuildPlans(*dir)
	if err != nil{
		log.Fatal(err)
	}

	if len(plans) == 0{
		fmt.Println("No files to organize")
		return
	}

	fmt.Println("planned operations:")
	for _, p := range plans{
		fmt.Printf("%s --> %s\n", p.Source, p.Target)
	}

	if !*apply{
		fmt.Println("Dry-Plan Only.Use --apply to execute.")
		return
	}

	if err := ExecutePlans(plans); err != nil {
		log.Fatal(err)
	}

	fmt.Println("done.")
}	
