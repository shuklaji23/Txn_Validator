package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	var NewTxn []map[string]Data
	err := c.ShouldBindJSON(&NewTxn)
	if  err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, mp2 := range NewTxn {
		Validator(mp2)
	}
}

func GetAllBlocks(c *gin.Context) {
	blocks := readFile()
	blocks = blocks[1:]
	c.JSON(http.StatusOK, blocks)
}

func readFile() []Block {
	var blocks []Block
	file, err := os.Open("ledger.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var tmp Block
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println("error while unmarshalling", err)
		}
		blocks = append(blocks, tmp)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error while Scanning", err)
	}
	return blocks
}

func GetBlock(c *gin.Context) {
	blockNumber := c.Param("blockNumber")
	Block, _ := strconv.Atoi(blockNumber)
	line, err := readLine(Block+1, "ledger.txt")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := fetchBlockDetails(line)
	c.JSON(http.StatusOK, response)
}

func readLine(lineNumber int, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	currentLine := 1
	for scanner.Scan() {
		if lineNumber == int(currentLine) {
			return scanner.Text(), nil
		}
		currentLine++
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("line not found")
}

func fetchBlockDetails(blockData string) Block {
	var tmp Block
	err := json.Unmarshal([]byte(blockData), &tmp)
	if err != nil {
		fmt.Println(err)
	}
	return tmp
}
