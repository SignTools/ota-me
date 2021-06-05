# OTA ME

This is a very simple project that lets you OTA install iOS apps over HTTP.

It works by offering valid HTTPS OTA manifests, hosted here on GitHub, that install any app from localhost over HTTP. You can host your IPA locally on any port, then use the corresponding manifest from here to install that app. The app download remains completely local - the manifest is only used to resolve the app's local URL. No information is ever transmitted to the internet. This method works because Apple enforces the manifest to be served over HTTPS, but not the apps within the manifest.

## Usage

Open the following link on your iOS device:

https://signtools.github.io/ota-me/data/8080

It will instruct your device to install the app found under http://localhost:8080

You can replace the trailing `8080` to any number between `0-10000`.

## Development

The project has two branches:

- [master](https://github.com/SignTools/ota-me/tree/master) - manifest generator
- [data](https://github.com/SignTools/ota-me/tree/data) - the served manifests
