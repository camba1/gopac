# goPac

Simple pacman clone created to explore the ```pixel``` library in Go. It was a fun project to tackel during the COV-19 shutdown. 
The code is base on the tutorial found in the coding dream website at https://www.codingdream.com/index.php/simple-pacman-in-using-go-and-pixelgl-part-1 .

It is not a full featured pacman by any means, but good enough to enjoy getting chased by the ghosts :)

### Note:

To get the pixel library properly, I have to first install the followingin my Mac:

- github.com/faiface/glhf
- github.com/go-gl/glfw/v3.2/glfw 

### Running

To build the application, execute the following command:

```
go build
```

then run can run the application by running:

```
./goPac
```

Note that I have only run this in my Mac

### Testing

Created a few tests just o play with the go default engine. The tests can be run using:

```
 go test
```