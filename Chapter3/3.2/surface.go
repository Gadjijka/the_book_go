package main

import (
    "fmt"
    "math"
    "net/http"
    "log"
)

const (
    width, height = 600, 320
    cells = 100
    xyrange = 30.0
    xyscale = width/2/xyrange
    zscale = height * 0.4
    angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30), cos(30)

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:900", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("ContentType", "image/svg+xml")
    polygon()
}

func polygon() {
    fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg'> "+
        "style='stroke:grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay, err := corner(i+1, j)
            if err != nil {
                continue
            }
            bx, by, err := corner(i, j)
            if err != nil {
                continue
            }
            cx, cy, err := corner(i, j+1)
            if err != nil {
                continue
            }
            dx, dy, err := corner(i+1, j+1)
            if err != nil {
                continue
            }
            fmt.Printf("<polygon points='%g,%g, %g,%g, %g,%g, %g,%g'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy)
        }
    }
    fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, error) {
    x := xyrange * (float64(i)/cells -0.5)
    y := xyrange * (float64(j)/cells - 0.5)
    z := f(x, y)
    if math.IsInf(z, 0) || math.IsNaN(z) {
	return 0, 0, fmt.Errorf("invalid value")
    }
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy, nil
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}

