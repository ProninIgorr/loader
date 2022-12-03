
func SaveImage(img *image.Gray, fname string) {
	f, err := os.Create(fname)
	if err != nil {
		log.Fatalf("cannot create file %s, err: %s", fname, err)
	}
	defer f.Close()

	ext := path.Ext(fname)
	if ext == ".jpg" {
		err = jpeg.Encode(f, img, nil)
	} else if ext == ".png" {
		err = png.Encode(f, img)
	} else if ext == ".bmp" {
		err = bmp.Encode(f, img)
	} else {
		log.Fatalf("%q extension not supported", ext)
	}
	if err != nil {
		log.Fatalf("cannot encode image %s, err: %s", fname, err)
	}
}
//write loader bmp with gocv
func LoadImage2 (fname string) *matrix.M {
	img := gocv.IMRead(fname, gocv.IMReadGrayScale)
	if img.Empty() {
		log.Fatalf("cannot open image %s", fname)
	}
	defer img.Close()
	resizedImg := maybeResizeImage(img)
	bounds := resizedImg.Bounds()
	gray := image.NewGray(bounds)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			c := resizedImg.At(x, y)
			gray.Set(x, y, color.GrayModel.Convert(c))
		}
	}
	return matrix.NewFromGray(gray)
}
func m() {
    //open webcam
    webcam, _ := gocv.VideoCaptureDevice(0)
    defer webcam.Close()

    //open display window
    window := gocv.NewWindow("Hello")
    defer window.Close()

    //prepare image matrix
    img := gocv.NewMat()
    defer img.Close()

    fmt.Println("start reading camera device: 0")

    for {
        if ok := webcam.Read(&img); !ok {
            fmt.Printf("cannot read device 0\n")
            return
        }
        if img.Empty() {
            continue
        }

        //show the image in the window, and wait 1 millisecond
        window.IMShow(img)
        window.WaitKey(1)
    }
}

func m2() {
	img := gocv.IMRead("test.bmp", gocv.IMReadColor)
	if img.Empty() {
		fmt.Println("Error reading image from: test.bmp")
		return
	}
	defer img.Close()

	// convert image Mat to 256x256 image
	img2 := gocv.NewMat()
	defer img2.Close()
	gocv.Resize(img, &img2, image.Pt(256, 256), 0, 0, gocv.InterpolationCubic)

	// convert image to grayscale
	img3 := gocv.NewMat()
	defer img3.Close()
	gocv.CvtColor(img2, &img3, gocv.ColorBGRToGray)

	// convert image to binary
	img4 := gocv.NewMat()
	defer img4.Close()
	gocv.Threshold(img3, &img4, 0, 255, gocv.ThresholdBinaryInv|gocv.ThresholdOtsu)

	// find contours
	contours := gocv.FindContours(img4, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	fmt.Printf("found %d contours

", len(contours))

	// draw contours
	gocv.DrawContours(&img4, contours, -1, color.RGBA{0, 0, 255, 0}, 2)

	// show the image in the window, and wait 1 millisecond
	window := gocv.NewWindow("Contours")
	defer window.Close()
	window.IMShow(img4)
	window.WaitKey(0)
}

func m3() {
	//load image
	img := helpers.LoadImage("test.jpg")
	//convert image to bmp
	bmp := gocv.NewMat()
	gocv.IMRead("test.jpg", gocv.IMReadColor)
	//convert to grayscale
	gocv.CvtColor(bmp, &bmp, gocv.ColorBGRToGray)
	//detect feature points
	featurePoints := gocv.NewMat()
	gocv.GoodFeaturesToTrack(bmp, &featurePoints, 100, 0.05, 10)
	//draw feature points
	for i := 0; i < featurePoints.Rows(); i++ {
		x := featurePoints.GetFloatAt(i, 0)
		y := featurePoints.GetFloatAt(i, 1)
		gocv.Circle(&bmp, image.Pt(int(x), int(y)), 3, color.RGBA{255, 0, 0, 0}, -1)
	}
	//save bmp
	gocv.IMWrite("test.bmp", bmp)
}

unc m4() {

	//Load image
	img := gocv.IMRead("C:/Users/Admin/Desktop/1.bmp", 1)
	if img.Empty() {
		fmt.Println("Error reading image from disk")
	}

	//Save image to file
	gocv.IMWrite("C:/Users/Admin/Desktop/2.bmp", img)

	//Show image in window
	window := gocv.NewWindow("Image")
	window.IMShow(img)
	window.WaitKey(0)

	//Convert image to grayscale
	gray := gocv.NewMat()
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	//Detect corners
	corners := gocv.NewMat()
	gocv.GoodFeaturesToTrack(gray, &corners, 100, 0.01, 10)

	//Draw corners on the original image
	for i := 0; i < corners.Rows(); i++ {
		x := corners.GetFloatAt(i, 0)
		y := corners.GetFloatAt(i, 1)
		gocv.Circle(&img, image.Pt(int(x), int(y)), 3, color.RGBA{0, 255, 0, 0}, -1)
	}

	//Show the image with corners
	window.IMShow(img)
	window.WaitKey(0)
}

