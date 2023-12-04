package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Node struct {
	alphabet   rune
	value      int
	summ       int
	leftChild  *Node
	rightChild *Node
}
type encodedResult struct {
	char string
	code string
}

func (n *Node) String() string {
	if n.summ == 0 {
		return fmt.Sprintf("This is a char node %s, value %d", n.alphabet, n.value)
	}
	return fmt.Sprintf("This is a summ node %d", n.summ)
}

func alphabetCount(fileName string) (map[rune]int, error) {
	file, err := os.ReadFile("testFile.txt")
	if err != nil {
		return nil, err
	}
	byteData := []byte(file)
	alphabetCount := make(map[rune]int)
	for i := 0; i < len(byteData); i++ {
		if byte('a') < []byte(strings.ToLower(string(byteData[i])))[0] && byte('z') >= []byte(strings.ToLower(string(byteData[i])))[0] {
			alphabetCount[rune(byteData[i])]++
		}
	}
	sortedKeys := make([]rune, 0, len(alphabetCount))
	for key := range alphabetCount {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Slice(sortedKeys, func(i, j int) bool {
		return sortedKeys[i] < sortedKeys[j]
	})
	sortedAlphabetCount := make(map[rune]int)
	for _, key := range sortedKeys {
		sortedAlphabetCount[key] = alphabetCount[key]
	}
	return sortedAlphabetCount, nil
}

type alphabetElement struct {
	rune rune
	occ  int
}

func buildHuffmanTree(alphabetsCount map[rune]int) *Node {
	var result []alphabetElement
	for key, value := range alphabetsCount {
		result = append(result, alphabetElement{rune: key, occ: value})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].occ < result[j].occ
	})
	var firstNode *Node
	var secodeNode *Node
	if len(result) >= 1 {
		firstNode = &Node{
			alphabet:   result[0].rune,
			value:      result[0].occ,
			summ:       0,
			leftChild:  nil,
			rightChild: nil,
		}
	}
	if len(result) >= 2 {
		secodeNode = &Node{
			alphabet:   result[1].rune,
			value:      result[1].occ,
			summ:       0,
			leftChild:  nil,
			rightChild: nil,
		}
	}
	parent := &Node{
		alphabet:   0,
		value:      0,
		summ:       firstNode.value + secodeNode.value,
		leftChild:  firstNode,
		rightChild: secodeNode,
	}
	var i int
	for i = 2; i < len(result); i++ {
		newNode := &Node{
			alphabet:   result[i].rune,
			value:      result[i].occ,
			summ:       0,
			leftChild:  nil,
			rightChild: nil,
		}
		Newparent := &Node{
			alphabet:   0,
			value:      0,
			summ:       parent.summ + newNode.value,
			leftChild:  parent,
			rightChild: newNode,
		}
		parent = Newparent
	}
	return parent
}

func DisplayPreOrderTraversal(root *Node) {
	if root != nil {
		fmt.Println(root.String())
		DisplayPreOrderTraversal(root.leftChild)
		DisplayPreOrderTraversal(root.rightChild)
	}
}
func generateFreqTable(occMap map[rune]int) {
	fmt.Println("Character | Frequency")
	fmt.Println("----------------------")

	for key, value := range occMap {
		fmt.Printf("%9s | %9d\n", string(key), value)
	}
}
func searchNodePreOrderTraversal(root *Node, charC rune, path string) string {
	if root == nil {
		return ""
	}
	if root.alphabet == charC {
		return path
	}
	leftPath := searchNodePreOrderTraversal(root.leftChild, charC, path+"0")
	if leftPath != "" {
		return leftPath
	}
	rightPath := searchNodePreOrderTraversal(root.rightChild, charC, path+"1")
	if rightPath != "" {
		return rightPath
	}
	return ""
}
func generatePrefixCodeTable(root *Node, occMap map[rune]int) []encodedResult {
	var encodedResults []encodedResult
	for key, _ := range occMap {
		var path string

		encodedResults = append(encodedResults, encodedResult{
			char: string(key),
			code: searchNodePreOrderTraversal(root, key, path),
		})
	}
	return encodedResults
}
func main() {
	chars, err := alphabetCount("testFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	generateFreqTable(chars)
	root := buildHuffmanTree(chars)
	//DisplayPreOrderTraversal(root)
	table := generatePrefixCodeTable(root, chars)
	fmt.Println("Character | Code")
	fmt.Println("----------------------")

	for _, enc := range table {
		fmt.Printf("%9s | %9s\n", string(enc.char), enc.code)
	}
}
