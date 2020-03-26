package simple

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

//https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
func (*Ref)FindMid()  {
	nums1, nums2 := []int{3},[2]int{-2,-1}

	nums1C  := nums1[:]
	nums1CL:= len(nums1)
	i1     := 0
	for _, v2 := range nums2 {
		for ; i1 < nums1CL; i1++ {
			if v2 < nums1C[i1] {
				left   := make([]int,i1+1)
				copy(left,nums1C[:i1])
				left[i1] = v2
				left     = append(left,nums1C[i1:]...)
				nums1C   = left
				nums1CL  = len(nums1C)
				i1++
				break
			} else if i1 == nums1CL-1 {
				nums1C  = append(nums1C,v2)
				nums1CL = len(nums1C)
				i1++
				break
			}
		}
	}

	if nums1CL == 0 && len(nums2) > 0 {
		nums1C  = nums2[:]
		nums1CL = len(nums2)
	}

	var mid float64
	if nums1CL % 2 == 0 {
		index := nums1CL/2
		mid = (float64(nums1C[index]) + float64(nums1C[index-1])) / 2
	} else {
		mid = float64(nums1C[nums1CL / 2])
	}

	fmt.Println(mid)
}

// https://leetcode-cn.com/problems/greatest-common-divisor-of-strings/
func (*Ref)MaxDiv()  {
	//假设str1 为长的那个 str2为短的那个
	str1 := "LEET"
	str2 := "CODE"

	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	
	str3 := ""
	maxL := 0

	for i := range str2 {
		if  str2[:i+1] == str1[:i+1] &&
			strings.ReplaceAll(str1,str1[:i+1],"") == "" &&
			strings.ReplaceAll(str2,str1[:i+1],"") == "" {
			if maxL < len(str1[:i+1]) {
				str3 = str1[:i+1]
				maxL = len(str1[:i+1])
			}
		}
	}

	fmt.Println(str3)
	fmt.Println(maxL)
}

// https://leetcode-cn.com/problems/longest-palindromic-substring/
func (*Ref)LPalidromic() {
	s := "cyyoacmjwjubfkzrrbvquqkwhsxvmytmjvbborrtoiyotobzjmohpadfrvmxuagbdczsjuekjrmcwyaovpiogspbslcppxojgbfxhtsxmecgqjfuvahzpgprscjwwutwoiksegfreortttdotgxbfkisyakejihfjnrdngkwjxeituomuhmeiesctywhryqtjimwjadhhymydlsmcpycfdzrjhstxddvoqprrjufvihjcsoseltpyuaywgiocfodtylluuikkqkbrdxgjhrqiselmwnpdzdmpsvbfimnoulayqgdiavdgeiilayrafxlgxxtoqskmtixhbyjikfmsmxwribfzeffccczwdwukubopsoxliagenzwkbiveiajfirzvngverrbcwqmryvckvhpiioccmaqoxgmbwenyeyhzhliusupmrgmrcvwmdnniipvztmtklihobbekkgeopgwipihadswbqhzyxqsdgekazdtnamwzbitwfwezhhqznipalmomanbyezapgpxtjhudlcsfqondoiojkqadacnhcgwkhaxmttfebqelkjfigglxjfqegxpcawhpihrxydprdgavxjygfhgpcylpvsfcizkfbqzdnmxdgsjcekvrhesykldgptbeasktkasyuevtxrcrxmiylrlclocldmiwhuizhuaiophykxskufgjbmcmzpogpmyerzovzhqusxzrjcwgsdpcienkizutedcwrmowwolekockvyukyvmeidhjvbkoortjbemevrsquwnjoaikhbkycvvcscyamffbjyvkqkyeavtlkxyrrnsmqohyyqxzgtjdavgwpsgpjhqzttukynonbnnkuqfxgaatpilrrxhcqhfyyextrvqzktcrtrsbimuokxqtsbfkrgoiznhiysfhzspkpvrhtewthpbafmzgchqpgfsuiddjkhnwchpleibavgmuivfiorpteflholmnxdwewj"

	isPa := func(str string, le int) bool {
		for i := 0; i < le/2; i++ {
			if str[i] != str[le-i-1] {
				return false
			}
		}
		return true
	}

	m := ""
	l := 0
	sL := len(s)
	for i := 0; i < sL; i++ {
		for j := i + 1; j <= sL; j++ {
			Le := len(s[i:j])
			if Le > l && isPa(s[i:j], Le) {
				l = len(s[i:j])
				m = s[i:j]
			}
		}
	}
	fmt.Println(m)
}
func (*Ref)FindMid2()  {
}

func (*Ref)CompressString()  {
	S := "aabcccccaaa"
	SL := len(S)
	sc := ""
	var v1 rune
	v1L := 0
	for i:= 0; i < SL + 1; i++ {
		if (i == SL) || (v1 != rune(S[i]) && v1L > 0) {
			sc += string(v1) + strconv.Itoa(v1L)
			v1L = 0
		}

		if i == SL {
			continue
		}
		v1 = rune(S[i])
		v1L++
	}


	if SL <= len(sc) {
		sc = S
	}
	fmt.Println(sc)
}

// https://leetcode-cn.com/problems/surface-area-of-3d-shapes/
func (*Ref)SurfaceArea()  {
	grid := [][]int{
		{2,2,2},{2,1,2},{2,2,2},
	}

	s     := 0
	for i := range grid {
		for j,v := range grid[i]  {
			curS := 0
			if v > 0 {
				curS = 6*v - 2*(v-1)
			}
			s = s + curS

			//情况1 j相邻
			prevj := j-1
			previ := i
			if prevj >= 0 {
				min := grid[previ][prevj]
				if v < min {
					min = v
				}
				s = s - min * 2
			}

			//情况2 i相邻
			prevj = j
			previ = i-1
			if previ >= 0 {
				min := grid[previ][prevj]
				if v < min {
					min = v
				}
				s = s - min * 2
			}
		}
	}

	log.Println(s)
	//log.Println(prevAll)
}

// https://leetcode-cn.com/problems/available-captures-for-rook/
func (*Ref)NRC()  {

	//R : [2,3]  //y轴位置 [7,3] ~ [0,3]
	board := [][]byte{
		{'.','.','.','.','.','.','.','.'},
		{'.','.','.','p','.','.','.','.'},
		{'.','.','.','R','.','.','.','p'},
		{'.','.','.','.','.','.','.','.'},
		{'.','.','.','.','.','.','.','.'},
		{'.','.','.','p','.','.','.','.'},
		{'.','.','.','.','.','.','.','.'},
		{'.','.','.','.','.','.','.','.'},
	}
	//是不是白车
	ifWhiteRook := func(b byte) bool {
		return b == 'R'
	}
	//是不是黑卒
	ifBlackPawn := func(b byte) bool {
		return b == 'p'
	}
	//是不是空格
	ifEmpty := func(b byte) bool {
		return b == '.'
	}
	//白车一次吃的黑卒的数量
	nums := 0

	//1.找白车的位置[x1,y1]
	x1, y1 := 0, 0
	out:for x := range board  {
		for y := range board[x]  {
			if ifWhiteRook(board[x][y]) {
				x1, y1 = x, y
				break out
			}
		}
	}

	//找x轴上的黑卒的位置[x2,y2] [0,y1] ~ [8,y1]
	for i := 0; i < 8; i++ {
		if ifBlackPawn(board[i][y1]) {

			x2,y2 := i, y1
			leftX  := x2
			rightX := x1
			if rightX < leftX {
				leftX, rightX = rightX, leftX
			}

			//[x2,y2] ~ [x1,y1]之间必须全为'.',否则不能吃到黑卒
			for i1 := leftX + 1; i1 < rightX ; i1++ {
				if !ifEmpty(board[i1][y2]) {
					nums--
					break
				}
			}
			//假设找到了
			nums++
		}
	}

	//找y轴上的黑卒的位置[x2,y2] [x1,0] ~ [x1,8]
	for j := 0; j < 8; j++ {
		if ifBlackPawn(board[x1][j]) {
			x2,y2 := x1,j

			topY    := y2
			bottomY := y1
			if topY < bottomY {
				topY, bottomY = bottomY,topY
			}

			//[x2,y2] ~ [x1,y1]之间必须全为'.',否则不能吃到黑卒
			for j1 := bottomY + 1; j1 < topY ; j1++ {
				if !ifEmpty(board[x2][j1]) {
					nums--
					break
				}
			}
			//假设找到了
			nums++
		}
	}


	log.Println(nums)
}