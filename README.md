# 2FA code generator
Generate a 2fa code for a given secret.

```bash
go install github.com/webklex/2fam
```

## Usage
```bash
Usage of 2fam:
  -secret string  Secret / Key code to be used for calculating the 2fa challenge answer
  -utc            Calculate code based on UTC time
  -version        Show version and exit
```

Example:
```bash
go run main.go -secret x53stf45ddirxafpnbpmkhuge2pe2g3f
```

## Build
```bash
git clone https://github.com/webklex/2fam
cd 2fam
go build
```
..or:
```bash
git clone https://github.com/webklex/2fam
cd 2fam
./build.sh
```

## Security
If you discover any security related issues, please email github@webklex.com instead of using the issue tracker.

## Credits
- [Webklex][link-author]
- [tilaklodha/google-authenticator](https://github.com/tilaklodha/google-authenticator)
- [All Contributors][link-contributors]

## License
The MIT License (MIT). Please see [License File](LICENSE.md) for more information.


[link-author]: https://github.com/webklex
[link-contributors]: https://github.com/webklex/2fam/graphs/contributors