#!/bin/bash

# /codeがDockerfileのWORKDIR
cd frontend
# node/modules内のアプリケーションをインストールする
npm install --location=global
npm install --dev-server @angular-devkit/build-angular
# 起動
ng serve --host=0.0.0.0 --poll=1000
