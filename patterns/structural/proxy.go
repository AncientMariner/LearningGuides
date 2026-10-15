package structural

import "fmt"

// Proxy: controls access to an object.
// Common variants: virtual (lazy init), protection (access control), logging.

type Image interface {
	Display()
}

type RealImage struct {
	filename string
}

func NewRealImage(filename string) *RealImage {
	fmt.Printf("Loading image: %s\n", filename)
	return &RealImage{filename: filename}
}

func (r *RealImage) Display() {
	fmt.Printf("Displaying: %s\n", r.filename)
}

// VirtualProxy defers loading until first use.
type VirtualProxy struct {
	filename string
	real     *RealImage
}

func NewVirtualProxy(filename string) Image {
	return &VirtualProxy{filename: filename}
}

func (p *VirtualProxy) Display() {
	if p.real == nil {
		p.real = NewRealImage(p.filename)
	}
	p.real.Display()
}

// ProtectionProxy restricts access by role.
type ProtectionProxy struct {
	real Image
	role string
}

func NewProtectionProxy(real Image, role string) Image {
	return &ProtectionProxy{real: real, role: role}
}

func (p *ProtectionProxy) Display() {
	if p.role != "admin" {
		fmt.Println("Access denied")
		return
	}
	p.real.Display()
}

func ExampleProxy() {
	img := NewVirtualProxy("photo.png")
	img.Display() // loads on first call
	img.Display() // uses cached real image

	real := NewRealImage("secret.png")
	protected := NewProtectionProxy(real, "guest")
	protected.Display()
	// Output:
	// Loading image: photo.png
	// Displaying: photo.png
	// Displaying: photo.png
	// Loading image: secret.png
	// Access denied
}
