package main

import "fmt"

var (
	layerHigh  = 4
	layerNum = 4
)

func main()  {
	//fmt.Println("     🌲")
	//fmt.Println("    🌲🌲")
	//fmt.Println("   🌲🌲🌲")
	//fmt.Println("  🌲🌲🌲🌲")
	//fmt.Println("     🌲")
	//fmt.Println("    🌲🌲")
	//fmt.Println("   🌲🌲🌲")
	//fmt.Println("  🌲🌲🌲🌲")
	//fmt.Println(" 🌲🌲🌲🌲🌲")
	//fmt.Println("     🌲")
	//fmt.Println("    🌲🌲")
	//fmt.Println("   🌲🌲🌲")
	//fmt.Println("  🌲🌲🌲🌲")
	//fmt.Println(" 🌲🌲🌲🌲🌲")
	//fmt.Println("🌲🌲🌲🌲🌲🌲")
	for l := 1; l <= layerNum; l ++ {
		for i := 1; i <= layerHigh; i++ {
			if l > 1 && i <= l-1 {
				continue
			}
			for j := layerHigh - i; j > 0; j-- {
				fmt.Print("  ")
			}
			if l >= 1 && layerNum > l {
				for n := 1; n <= layerNum - l; n++ {
					fmt.Print("  ")
				}
			}
			for k := 1; k <= i; k++ {
				fmt.Print("🎄🎄")
			}
			fmt.Println()
		}
		layerHigh++
	}
	for i := 1; i <= 5; i++ {
		for j := 1; j <= layerHigh-2; j++ {
			fmt.Print("  ")
		}
		fmt.Print("🚪")
		fmt.Println()
	}
}