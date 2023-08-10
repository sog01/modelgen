package main

import (
	"fmt"

	"github.com/sog01/modelgen/parser"
)

func main() {
	p, _ := parser.ParseTable(`CREATE TABLE IF NOT EXISTS tasks (
		task_id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		start_date DATE,
		due_date DATE,
		status TINYINT NOT NULL,
		priority TINYINT NOT NULL,
		description TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)  ENGINE=INNODB;`)
	s := p.Struct()
	fmt.Println(s)
}
