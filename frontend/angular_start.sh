#!/bin/bash

# /codeがDockerfileのWORKDIR
cd frontend
# node/modules内のアプリケーションをインストールする
npm install

# 起動
ng serve --host=0.0.0.0 --poll=2000
