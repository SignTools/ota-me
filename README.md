# OTA ME

This is a very simple project that lets you OTA install iOS apps over HTTP.

It works by offering valid HTTPS OTA manifests that redirect to localhost, one for each port between `0` to `10000`. This way you can host your IPA locally on any port, then use the corresponding manifest from here to redirect to your local URL. The app download remains completely local - the manifest is only used to resolve the app's local URL. No information is ever transmitted to the internet.

## Usage

Open the following link on your iOS device:

https://signtools.github.io/ota-me/data/8080

It will instruct your device to install the app found under http://localhost:8080

You can replace the trailing `8080` to any number between `0-10000`.

## Why this works

Apple enforces the manifest to be served with HTTPS, but the apps within the manifest do not.

## Development

The project has two branches:

- [master](https://github.com/SignTools/ota-me/tree/master) - manifest generator
- [data](https://github.com/SignTools/ota-me/tree/data) - the served manifests
