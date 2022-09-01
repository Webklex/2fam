package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Secret string
	UTC    bool // use UTC for the timestamp instead of local time
}

var buildNumber string
var buildVersion string

func main() {
	c := &Config{
		Secret: "",
	}

	flag.CommandLine.StringVar(&c.Secret, "secret", c.Secret, "Secret / Key code to be used for calculating the 2fa challenge answer")
	flag.CommandLine.BoolVar(&c.UTC, "utc", c.UTC, "Calculate code based on UTC time")

	sv := flag.Bool("version", false, "Show version and exit")
	flag.Parse()

	if *sv {
		fmt.Printf("version: %s\n", buildVersion)
		fmt.Printf("build number: %s\n", buildNumber)
		os.Exit(0)
	}

	fmt.Println(c.getTOTPToken())
}

func (c *Config) getTOTPToken() string {
	var interval int64
	if c.UTC {
		interval = time.Now().UTC().Unix()
	} else {
		interval = time.Now().Unix()
	}
	return c.getHOTPToken(interval / 30)
}

func (c *Config) getHOTPToken(interval int64) string {
	//Converts secret to base32 Encoding. Base32 encoding desires a 32-character
	//subset of the twenty-six letters A–Z and ten digits 0–9
	key, err := base32.StdEncoding.DecodeString(strings.ToUpper(c.Secret))
	check(err)
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(interval))

	//Signing the value using HMAC-SHA1 Algorithm
	hash := hmac.New(sha1.New, key)
	hash.Write(bs)
	h := hash.Sum(nil)

	// We're going to use a subset of the generated hash.
	// Using the last nibble (half-byte) to choose the index to start from.
	// This number is always appropriate as it's maximum decimal 15, the hash will
	// have the maximum index 19 (20 bytes of SHA1) and we need 4 bytes.
	o := (h[19] & 15)

	var header uint32
	//Get 32 bit chunk from hash starting at the o
	r := bytes.NewReader(h[o : o+4])
	err = binary.Read(r, binary.BigEndian, &header)

	check(err)
	//Ignore most significant bits as per RFC 4226.
	//Takes division from one million to generate a remainder less than < 7 digits
	h12 := (int(header) & 0x7fffffff) % 1000000

	//Converts number as a string
	otp := strconv.Itoa(int(h12))

	return prefix0(otp)
}

//Panic if error is not nil
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Append extra 0s if the length of otp is less than 6
//If otp is "1234", it will return it as "001234"
func prefix0(otp string) string {
	if len(otp) == 6 {
		return otp
	}
	for i := (6 - len(otp)); i > 0; i-- {
		otp = "0" + otp
	}
	return otp
}
