1.
Question 1
Which of the following expressions does NOT compute the
average of two integers a and b?

1 point

X - avg := 2 % (a + b)


avg := float64(a + b) / 2


avg := float64(a + b) / 2.0


avg := float64(float64(a + b) / 2.0)

2.
Question 2
What is printed when the following program is executed?

123456
func main() {
  i, _ := strconv.Atoi("10")
  y := i * 2
  fmt.Println(y)
}

1 point

1010


10


102


X - 20

3.
Question 3
What is printed when the following program is executed?

5
func main() {
  s := strings.Replace("ianianian", "ni", "in", 2)
  fmt.Println(s)
}

1 point

ianianian


X - iainainan


iainanian


nianiania

4.
Question 4
What is printed by this code?

14
func main() {
  x:=7
  switch {
    case x>3:
      fmt.Printf("1")
    case x>5:
      fmt.Printf("2")
    case x==7:
      fmt.Printf("3")
    default: 

1 point

X - 1


2


3


4

5.
Question 5
What is printed by this code?

func main() {
  var xtemp int
  x1 := 0
  x2 := 1
  for x:=0; x<5; x++ {
    xtemp = x2
    x2 = x2 + x1
    x1 = xtemp
  }
  fmt.Println(x2)

1 point

5


13


X - 8


4

6.
Question 6
True or False:

This code compiles correctly.

func main() {
  var x int
  var y *int
  z := 3
  y = &z
  x = &y
}

1 point

True


X - False

7.
Question 7
Which integer type provides higher accuracy?

1 point

int16


int32


int


X - All of these types provide the same accuracy