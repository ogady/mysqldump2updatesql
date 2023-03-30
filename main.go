package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// コマンドライン引数でインプットファイルとアウトプットファイルを指定
	var inputFile, outputFile string
	flag.StringVar(&inputFile, "input", "", "CSV input file")
	flag.StringVar(&outputFile, "output", "", "Output file with update statements")
	flag.Parse()

	if inputFile == "" || outputFile == "" {
		fmt.Println("Error: both -input and -output flags are required")
		return
	}

	// CSVファイルを開く
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// 出力ファイルを作成
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer outFile.Close()

	// CSVリーダーを作成する
	reader := csv.NewReader(file)

	// ヘッダー行を読み取る
	header, err := reader.Read()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// CSVファイルを行ごとに読み取り、UPDATE文に変換する
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		primaryKey := record[0]
		updateFields := []string{}
		for i := 1; i < len(record); i++ {
			if record[i] == "NULL" {
				updateFields = append(updateFields, fmt.Sprintf("%s = %s", header[i], record[i]))
			} else {
				updateFields = append(updateFields, fmt.Sprintf("%s = '%s'", header[i], record[i]))
			}
		}
		updateStatement := fmt.Sprintf("UPDATE eu_pairs_jp.user_vaccination_status SET %s WHERE %s = %s;\n", strings.Join(updateFields, ", "), header[0], primaryKey)
		outFile.WriteString(updateStatement)
	}
}
