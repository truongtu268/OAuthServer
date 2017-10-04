## Oauth 2 server with echo framework
## Install and run server

<dl>
  <dt>Step 1:</dt>
  <dd>go get github.com/truongtu268/OAuthServer</dd>

  <dt>Step 2:</dt>
  <dd>go get github.com/pilu/fresh</dd>

  <dt>Step 3:</dt>
  <dd>go get github.com/tools/godep</dd>

  <dt>Step 4:</dt>
  <dd>godep restore</dd>

  <dt>Step 5:</dt>
  <dd>Install postgresql follow link https://www.postgresql.org/ and create database</dd>

  <dt>Step 6:</dt>
  <dd>Insert database config in file config.json with migrate='drop'</dd>

  <dt>Step 7:</dt>
  <dd>Run command:fresh (Start server)</dd>