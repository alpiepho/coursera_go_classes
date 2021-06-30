.
Question 1
What is printed when the following program is executed?

func main() {
  x := []int {4, 8, 5}
  y := -1
  for _, elt := range x {
    if elt > y {
      y = elt
    }
  }
  fmt.Print(y)
}

1 point

4


X - 8


5


-1

2.
Question 2
What is printed when the following program is executed?

func main() {
  x := [...]int {4, 8, 5}
  y := x[0:2]
  z := x[1:3]
  y[0] = 1
  z[1] = 3
  fmt.Print(x)
}

1 point

[1 3 8]


X - [1 8 3]


[4 1 3]


[4 8 5]

3.
Question 3
What is printed when the following program is executed?

func main() {
  x := [...]int {1, 2, 3, 4, 5}
  y := x[0:2]
  z := x[1:4]
  fmt.Print(len(y), cap(y), len(z), cap(z))
}

1 point

X - 2 5 3 4


2 4 3 4


2 3 3 4


2 5 3 5

4.
Question 4
What is printed when the following program is executed?

12345678910
func main() {
  x := map[string]int {
    "ian": 1, "harris": 2}
  for i, j := range x {
    if i == "harris" {
      fmt.Print(i, j)
    }
  }
}

1 point

harris1


1ian


1harris


X - harris2

5.
Question 5
What is printed when the following program is executed?

type P struct {
    x string
y int
} 
func main() {
  b := P{"x", -1}
  a := [...]P{P{"a", 10}, 
        P{"b", 2},
        P{"c", 3}}
  for _, z := range a {

1 point

X - a


b


c


x

6.
Question 6
What is printed when the following program is executed?

func main() {
  s := make([]int, 0, 3)
  s = append(s, 100)
  fmt.Println(len(s), cap(s))
}

1 point

X - 1 3


0 3


1 1


1 4
