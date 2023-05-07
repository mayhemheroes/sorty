package fuzz_sorty

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/jfcg/sorty/v2"
)

func test(n int) bool {

    if n > 0 {
        return true
    } else {
        return false
    }
}

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            case 0:
                var testSlice []int
                fuzzConsumer.CreateSlice(&testSlice)

                sorty.SortSlice(testSlice)
                return 0

            case 1:
                var testSlice []string
                fuzzConsumer.CreateSlice(&testSlice)

                sorty.SortSlice(testSlice)
                return 0

            case 2:
                var stringSlice []string
                fuzzConsumer.CreateSlice(&stringSlice)

                sorty.SortLen(stringSlice)
                return 0
                
            case 3:
                var stringSlice [][]byte
                fuzzConsumer.CreateSlice(&stringSlice)

                sorty.SortLen(stringSlice)
                return 0

            case 4:
                testInt, _ := fuzzConsumer.GetInt()

                sorty.Search(testInt, test)
                return 0
            
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}