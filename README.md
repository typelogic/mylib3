## Steps

Get gomobile, this is one time:
- go install golang.org/x/mobile/cmd/gomobile@latest

Build the Android library:
- go get -d golang.org/x/mobile/cmd/gobind
- gomobile bind -v -o android.aar -target=android .

An Android *android.aar* library will be generated.
